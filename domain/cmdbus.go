package domain

import (
	"context"
	"errors"

	"google.golang.org/protobuf/proto"
)

var ErrNoHandlerForCommand = errors.New("no handler for command")

type CommandBus interface {
	RegisterHandler(proto.Message, CommandHandler)
	SubmitCommand(context.Context, proto.Message) error
	AsyncSubmitCommand(proto.Message) error
}

type CommandHandler func(context.Context, proto.Message) error
