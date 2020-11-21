package lb

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alexrudd/lb-teams/domain"

	oldproto "github.com/golang/protobuf/proto" // nolint
	lift "github.com/liftbridge-io/go-liftbridge/v2"
	liftapi "github.com/liftbridge-io/liftbridge-api/go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// LiftBridgeEventStore implements domain.EventStore using LiftBridge.
type LiftBridgeEventStore struct {
	lbc            lift.Client
	cancelIndexing context.CancelFunc
	streamIndex    map[string]map[string]struct{}
}

// NewLiftBridgeEventStore returns a new LiftBridgeEventStore that is
// connected to a LiftBridge cluster.
func NewLiftBridgeEventStore(lbc lift.Client) *LiftBridgeEventStore {
	return &LiftBridgeEventStore{
		lbc:         lbc,
		streamIndex: map[string]map[string]struct{}{},
	}
}

// Close closes the internal Liftbridge client.
func (es *LiftBridgeEventStore) Close() error {
	es.cancelIndexing()
	return es.lbc.Close()
}

// GetStream attempts to get a stream using the provided ID and returns all
// events in that stream.
func (es *LiftBridgeEventStore) GetStream(ctx context.Context, aggregate, id string) (domain.Stream, error) {
	streamID := aggregate + "." + id

	err := es.lbc.CreateStream(ctx, streamID, streamID)
	if err != nil {
		if err != lift.ErrStreamExists {
			return nil, fmt.Errorf("creating stream: %w", err)
		}
	}

	event := &Event{}
	eventChan := make(chan *Event)

	err = es.lbc.Subscribe(ctx, streamID, func(msg *lift.Message, err error) {
		if err == lift.ErrReadonlyPartition {
			close(eventChan)
			return
		} else if err != nil {
			return
		}

		err = proto.Unmarshal(msg.Value(), event)
		if err != nil {
			return
		}

		eventChan <- event
	}, lift.StartAtEarliestReceived(), lift.StopAtLatestReceived())
	if err == lift.ErrReadonlyPartition {
		close(eventChan)
	} else if err != nil {
		return nil, fmt.Errorf("subscribing to stream: %w", err)
	}

	return &stream{
		client:   es.lbc,
		streamID: streamID,
		events:   eventChan,
	}, nil
}

type stream struct {
	client      lift.Client
	streamID    string
	events      <-chan *Event
	latestOffet int64
}

func (s *stream) Events() []domain.Event {
	var events []domain.Event

	for e := range s.events {
		s.latestOffet = e.GetOffset()
		events = append(events, e)
	}

	return events
}

// Publish publishes a new event to the stream.
func (s *stream) Publish(ctx context.Context, event proto.Message) error {
	data, err := anypb.New(event)
	if err != nil {
		return fmt.Errorf("marshalling data: %w", err)
	}

	value, err := proto.Marshal(&Event{
		Offset:      s.latestOffet + 1,
		PubTimstamp: timestamppb.Now(),
		StreamName:  s.streamID,
		RawData:     data,
	})
	if err != nil {
		return fmt.Errorf("marshalling event: %w", err)
	}

	ack, err := s.client.Publish(ctx, s.streamID, value)
	if err != nil {
		return fmt.Errorf("publishing message: %w", err)
	}

	s.latestOffet = ack.Offset()

	return nil
}

func (es *LiftBridgeEventStore) GetAggregateStream(ctx context.Context, aggregateType string) (domain.AggregateStream, error) {

	return &aggregateStream{
		ctx:           ctx,
		lbc:           es.lbc,
		aggregateType: aggregateType,
		cancels:       map[string]context.CancelFunc{},
	}, nil
}

type aggregateStream struct {
	ctx           context.Context
	lbc           lift.Client
	handler       domain.EventHandler
	aggregateType string
	cancels       map[string]context.CancelFunc
}

func (as *aggregateStream) Bind(handler domain.EventHandler) error {
	as.handler = handler
	ctx, cancel := context.WithCancel(as.ctx)
	as.cancels["__activity"] = cancel

	err := as.lbc.Subscribe(ctx, "__activity", as.handleActivityMsg, lift.StartAtEarliestReceived())
	if err != nil {
		return fmt.Errorf("subscribing to activity stream: %w", err)
	}

	return nil
}

func (as *aggregateStream) Close() {
	for _, c := range as.cancels {
		c()
	}
}

func (as *aggregateStream) handlerEvent(msg *lift.Message, err error) {
	if err != nil {
		return
	}

	event := &Event{}

	err = proto.Unmarshal(msg.Value(), event)
	if err != nil {
		return
	}

	as.handler(context.Background(), event)
}

func (as *aggregateStream) handleActivityMsg(msg *lift.Message, err error) {
	if err != nil {
		return
	}

	evt := &liftapi.ActivityStreamEvent{}

	err = oldproto.Unmarshal(msg.Value(), evt)
	if err != nil {
		return
	}

	switch evt.GetOp() {
	case liftapi.ActivityStreamOp_CREATE_STREAM:
		stream := evt.CreateStreamOp.GetStream()
		aggregateType := strings.Split(stream, ".")[0]

		if aggregateType != as.aggregateType {
			return
		}

		ctx, cancel := context.WithCancel(as.ctx)
		_ = as.lbc.Subscribe(ctx, stream, as.handlerEvent, lift.StartAtEarliestReceived())

		as.cancels[stream] = cancel
	case liftapi.ActivityStreamOp_DELETE_STREAM:
		stream := evt.CreateStreamOp.GetStream()
		aggregateType := strings.Split(stream, ".")[0]

		if aggregateType != as.aggregateType {
			return
		}

		if c, ok := as.cancels[stream]; ok {
			c()
			delete(as.cancels, stream)
		}
	}
}

func (e *Event) Clock() uint64 {
	return uint64(e.GetOffset())
}

func (e *Event) Timestamp() time.Time {
	return e.GetPubTimstamp().AsTime()
}

func (e *Event) Data() proto.Message {
	d, _ := e.GetRawData().UnmarshalNew()
	return d
}
