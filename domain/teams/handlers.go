package teams

import (
	"context"
	"fmt"

	"github.com/alexrudd/lb-teams/domain"
	"google.golang.org/protobuf/proto"
)

// UserCommand is a command that can be executed on a User.
type UserCommand interface {
	GetUserId() string
	Execute(*User) (proto.Message, error)
	proto.Message
}

// UserCommandHandler is a function that handles User commands
type UserCommandHandler func(context.Context, UserCommand) error

// NewUserCommandHandler handler user commands.
func NewUserCommandHandler(store domain.EventStore) UserCommandHandler {
	return func(ctx context.Context, cmd UserCommand) error {
		stream, err := store.GetStream(ctx, UserAggregate, cmd.GetUserId())
		if err != nil {
			return fmt.Errorf("getting user stream for user %s: %w", cmd.GetUserId(), err)
		}

		// rehydrate user
		user := RehydrateUser(stream.Events(), cmd.GetUserId())
		event, err := cmd.Execute(user)
		if err != nil {
			return fmt.Errorf("executing user command for user %s: %w", cmd.GetUserId(), err)
		}

		if err := stream.Publish(ctx, event); err != nil {
			return fmt.Errorf("publishing user event for user %s: %w", cmd.GetUserId(), err)
		}

		return nil
	}
}

func SetupEventHandlers(store domain.EventStore) error {
	if err := store.Subscribe(UserAggregate, newOwnerChangedHandler(store)); err != nil {
		return err
	}

	return nil
}

// NewOwnerChangedHandler handler for handling OwnerChanged events.
func newOwnerChangedHandler(store domain.EventStore) domain.EventHandler {
	return func(ctx context.Context, evt proto.Message) {
		ownerChanged, ok := evt.(*OwnerChanged)
		if !ok {
			return
		}

		stream, err := store.GetStream(ctx, UserAggregate, ownerChanged.GetNewOwnerUserId())
		if err != nil {
			return
		}

		// rehydrate user
		user := RehydrateUser(stream.Events(), ownerChanged.GetNewOwnerUserId())
		event := user.OnLeaderChanged(ownerChanged)

		if err := stream.Publish(ctx, event); err != nil {
			return
		}
	}
}
