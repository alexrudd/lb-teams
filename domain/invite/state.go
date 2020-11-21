package invite

import (
	"context"
	"fmt"

	"github.com/alexrudd/lb-teams/domain"
	"github.com/google/uuid"
)

const Name = "invite"

type status int8

const (
	pending status = iota
	declined
	cancelled
	expired
	accepted
)

type state struct {
	id            string
	status        status
	teamID        string
	inviterUserID string
	inviteeUserID string
}

// NewFactoryHandler handler invite creation.
func NewFactoryHandler(store domain.EventStore) func(context.Context, Factory) error {
	return func(ctx context.Context, fac Factory) error {
		if err := fac.Validate(); err != nil {
			return fmt.Errorf("validating factory command for invite: %w", err)
		}

		inviteID := uuid.New().String()

		event := fac.InitialEvent(inviteID)

		stream, err := store.GetStream(ctx, Name, inviteID)
		if err != nil {
			return fmt.Errorf("getting stream for invite %s: %w", inviteID, err)
		}

		if err := stream.Publish(ctx, event); err != nil {
			return fmt.Errorf("publishing event for invite %s: %w", inviteID, err)
		}

		return nil
	}
}

// NewCommandHandler handles invite commands.
func NewCommandHandler(store domain.EventStore) func(context.Context, Command) error {
	return func(ctx context.Context, cmd Command) error {
		if err := cmd.Validate(); err != nil {
			return fmt.Errorf("validating command for invite %s: %w", cmd.GetInviteId(), err)
		}

		stream, err := store.GetStream(ctx, Name, cmd.GetInviteId())
		if err != nil {
			return fmt.Errorf("getting stream for invite %s: %w", cmd.GetInviteId(), err)
		}

		// rehydrate state
		invite := RehydrateState(stream.Events())
		event, err := cmd.Execute(invite)
		if err != nil {
			return fmt.Errorf("executing command for invite %s: %w", cmd.GetInviteId(), err)
		}

		if !event.ProtoReflect().IsValid() {
			return nil
		}

		if err := stream.Publish(ctx, event); err != nil {
			return fmt.Errorf("publishing event for invite %s: %w", cmd.GetInviteId(), err)
		}

		return nil
	}
}

// RehydrateState rehydrates an invite's state.
func RehydrateState(events []domain.Event) *state {
	s := &state{}

	for _, e := range events {
		if ue, ok := e.Data().(Event); ok {
			ue.Apply(s)
		}
	}

	return s
}
