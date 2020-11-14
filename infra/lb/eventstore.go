package lb

import (
	"context"
	"fmt"
	"log"

	"github.com/alexrudd/lb-teams/domain"
	"github.com/nats-io/nats.go"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// LiftBridgeEventStore implements domain.EventStore using
// LiftBridge.
type LiftBridgeEventStore struct {
	nc  *nats.Conn
	lbc lift.Client
}

// NewLiftBridgeEventStore returns a new LiftBridgeEventStore that is
// connected to a LiftBridge cluster.
func NewLiftBridgeEventStore(nc *nats.Conn, lbc lift.Client) *LiftBridgeEventStore {
	return &LiftBridgeEventStore{
		nc:  nc,
		lbc: lbc,
	}
}

// Close closes the internal Liftbridge client.
func (es *LiftBridgeEventStore) Close() error {
	return es.lbc.Close()
}

// GetStream attempts to get a stream using the provided ID
// and returns all events in that stream.
func (es *LiftBridgeEventStore) GetStream(ctx context.Context, aggregate, id string) (domain.Stream, error) {
	streamID := aggregate + "." + id

	err := es.lbc.CreateStream(ctx, streamID, streamID)
	if err != nil {
		if err != lift.ErrStreamExists {
			return nil, fmt.Errorf("creating stream: %w", err)
		}
	}

	md, err := es.lbc.FetchPartitionMetadata(ctx, streamID, 0)
	if err != nil {
		return nil, fmt.Errorf("fetching partition metadata: %w", err)
	}

	log.Printf("got stream for ID: %s\n", id)

	if md.NewestOffset() == -1 {
		return &stream{
			client:   es.lbc,
			streamID: streamID,
			events:   nil,
		}, nil
	}

	envelope := &Message{}
	eventChan := make(chan proto.Message)
	ctx, cancel := context.WithCancel(ctx)

	err = es.lbc.Subscribe(ctx, streamID, func(msg *lift.Message, err error) {
		if err != nil {
			log.Printf("received error from stream: %s", err)
			return
		}

		err = proto.Unmarshal(msg.Value(), envelope)
		if err != nil {
			log.Printf("error unmarshalling message from stream %s: %s", streamID, err)
			return
		}

		event, err := envelope.GetPayload().UnmarshalNew()
		if err != nil {
			log.Printf("error unmarshalling payload from stream %s: %s", streamID, err)
			return
		}

		eventChan <- event

		if msg.Offset() == md.NewestOffset() {
			close(eventChan)
			cancel()
		}
	}, lift.StartAtEarliestReceived())
	if err != nil {
		return nil, fmt.Errorf("subscribing to stream: %w", err)
	}

	return &stream{
		client:   es.lbc,
		streamID: streamID,
		events:   eventChan,
	}, nil
}

func (es *LiftBridgeEventStore) Subscribe(aggregate string, handler domain.EventHandler) error {
	_, err := es.nc.Subscribe(aggregate+".>", func(msg *nats.Msg) {
		envelope := &Message{}
		err := proto.Unmarshal(msg.Data, envelope)
		if err != nil {
			return
		}

		e, err := envelope.GetPayload().UnmarshalNew()
		if err != nil {
			return
		}

		handler(context.Background(), e)
	})

	return err
}

type stream struct {
	client   lift.Client
	streamID string
	events   <-chan proto.Message
}

func (s *stream) Events() <-chan proto.Message {
	return s.events
}

// Publish publishes a new event to the stream.
func (s *stream) Publish(ctx context.Context, event proto.Message) error {
	pl, err := anypb.New(event)
	if err != nil {
		return fmt.Errorf("marshalling payload: %w", err)
	}

	value, err := proto.Marshal(&Message{Payload: pl})
	if err != nil {
		return fmt.Errorf("marshalling payload: %w", err)
	}

	_, err = s.client.Publish(ctx, s.streamID, value)
	if err != nil {
		return fmt.Errorf("publishing message: %w", err)
	}

	return nil
}
