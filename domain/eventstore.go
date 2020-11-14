package domain

import (
	"context"

	"google.golang.org/protobuf/proto"
)

// EventStore persists streams of aggregate events.
type EventStore interface {
	GetStream(ctx context.Context, aggregate, id string) (Stream, error)
	Subscribe(aggregate string, handler EventHandler) error
}

// Stream contains all events published by an aggregate instance.
type Stream interface {
	Events() <-chan proto.Message
	Publish(ctx context.Context, event proto.Message) error
}

// EventHandler is a function that handles events
type EventHandler func(context.Context, proto.Message)
