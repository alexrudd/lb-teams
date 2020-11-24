package im

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/alexrudd/lb-teams/domain"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Bus struct {
	commands []proto.Message
	handlers map[protoreflect.FullName]domain.CommandHandler
	mtx      sync.Mutex
}

func NewBus() *Bus {
	return &Bus{
		handlers: map[protoreflect.FullName]domain.CommandHandler{},
	}
}

func (b *Bus) Run(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Millisecond)

	go func() {
		for {
			select {
			case <-ticker.C:
				if len(b.commands) != 0 {
					b.sendCommands()
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (b *Bus) RegisterHandler(msg proto.Message, h domain.CommandHandler) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	name := msg.ProtoReflect().Descriptor().FullName()

	if _, ok := b.handlers[name]; ok {
		panic("handler already registered for: " + string(name))
	}

	b.handlers[name] = h
}

func (b *Bus) AsyncSubmitCommand(cmd proto.Message) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.commands = append(b.commands, cmd)

	return nil
}

func (b *Bus) SubmitCommand(ctx context.Context, cmd proto.Message) error {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if h, ok := b.handlers[cmd.ProtoReflect().Descriptor().FullName()]; ok {
		return h(ctx, cmd)
	}

	return domain.ErrNoHandlerForCommand
}

func (b *Bus) sendCommands() {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	for _, cmd := range b.commands {
		if h, ok := b.handlers[cmd.ProtoReflect().Descriptor().FullName()]; ok {
			if err := h(context.Background(), cmd); err != nil {
				log.Printf("error executing command %T: %s", cmd, err)
			}
		} else {
			log.Printf("no handler for command %T", cmd)
		}
	}
}
