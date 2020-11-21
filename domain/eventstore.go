package domain

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
)

// EventStore persists streams of aggregate events.
type EventStore interface {
	GetStream(ctx context.Context, aggregateType, aggregateID string) (Stream, error)
	GetAggregateStream(ctx context.Context, aggregateType string) (AggregateStream, error)
}

// Stream contains all events published by an aggregate instance.
type Stream interface {
	Events() []Event
	Publish(ctx context.Context, event proto.Message) error
}

// AggregateStream contains all events published by an aggregate instance.
type AggregateStream interface {
	Bind(EventHandler) error
	Close()
}

type Event interface {
	Clock() uint64
	Timestamp() time.Time
	Data() proto.Message
}

// EventHandler is a function that handles events
type EventHandler func(context.Context, Event)
