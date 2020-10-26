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

	if md.NewestOffset() == -1 {
		return &stream{
			client:   es.lbc,
			streamID: streamID,
			events:   nil,
		}, nil
	}

	msgChan := make(chan *lift.Message)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = es.lbc.Subscribe(ctx, streamID, func(msg *lift.Message, err error) {
		if err != nil {
			log.Printf("received error from stream: %s", err)
			return
		}

		msgChan <- msg
	}, lift.StartAtEarliestReceived())
	if err != nil {
		return nil, fmt.Errorf("subscribing to stream: %w", err)
	}

	var (
		events   = []proto.Message{}
		envelope = &Message{}
	)

	for m := range msgChan {
		err = proto.Unmarshal(m.Value(), envelope)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling message from stream %s: %s", streamID, err)
		}

		e, err := envelope.GetPayload().UnmarshalNew()
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling payload from stream %s: %s", streamID, err)
		}

		events = append(events, e)

		if m.Offset() == md.NewestOffset() {
			break
		}
	}

	return &stream{
		client:   es.lbc,
		streamID: streamID,
		events:   events,
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
	events   []proto.Message
}

func (s *stream) Events() []proto.Message {
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
