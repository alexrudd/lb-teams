package user

import (
	"context"
	"fmt"
	"log"

	"github.com/alexrudd/lb-teams/domain"
	"google.golang.org/protobuf/proto"
)

const UserAggregate = "user"

// User is a user who is part of a team.
type User struct {
	userID  string
	teamID  string
	invites map[string]struct{}
}

// CommandHandler is a function that handles User commands
type CommandHandler func(context.Context, Command) error

// NewUserCommandHandler handler user commands.
func NewCommandHandler(store domain.EventStore) CommandHandler {
	return func(ctx context.Context, cmd Command) error {
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

// RehydrateUser rehydrates a User aggregate from an event
// stream
func RehydrateUser(events <-chan proto.Message, userID string) *User {
	u := &User{
		userID:  userID,
		invites: map[string]struct{}{},
	}

	if events == nil {
		return u
	}

	for e := range events {
		if ue, ok := e.(Event); ok {
			log.Printf("applying event: %T\n", ue)
			ue.Apply(u)
		}
	}

	return u
}
