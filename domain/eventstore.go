package domain

import (
	"context"

	"google.golang.org/protobuf/proto"
)

// EventStore manages streams of aggregate events.
type EventStore interface {
	GetStream(ctx context.Context, id string) (Stream, error)
}

// Stream represents all events published by an aggregate instance.
type Stream interface {
	Events() []proto.Message
	Publish(ctx context.Context, event proto.Message) error
}
