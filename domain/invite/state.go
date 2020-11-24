package invite

import (
	"context"
	"fmt"
	"log"

	"github.com/alexrudd/lb-teams/domain"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

const Name = "invite"

// status represents the posible states an invite can be in.
type status int8

const (
	pending status = iota
	declined
	cancelled
	expired
	accepted
)

// state stores the aggregate for a particular invite stream.
type state struct {
	id            string
	status        status
	teamID        string
	inviterUserID string
	inviteeUserID string
}

func RegisterWithCommandBus(bus domain.CommandBus, store domain.EventStore) {
	var (
		facHandler = NewFactoryHandler(store)
		cmdHandler = NewCommandHandler(store)
	)

	bus.RegisterHandler(&SendTeamFormationInvite{}, facHandler)
	bus.RegisterHandler(&SendTeamInvite{}, facHandler)

	bus.RegisterHandler(&DeclineTeamFormationInvite{}, cmdHandler)
	bus.RegisterHandler(&CancelTeamFormationInvite{}, cmdHandler)
	bus.RegisterHandler(&ExpireTeamFormationInvite{}, cmdHandler)
	bus.RegisterHandler(&AcceptTeamFormationInvite{}, cmdHandler)
	bus.RegisterHandler(&DeclineTeamInvite{}, cmdHandler)
	bus.RegisterHandler(&CancelTeamInvite{}, cmdHandler)
	bus.RegisterHandler(&ExpireTeamInvite{}, cmdHandler)
	bus.RegisterHandler(&AcceptTeamInvite{}, cmdHandler)
}

// NewFactoryHandler handler invite foactor commands.
func NewFactoryHandler(store domain.EventStore) domain.CommandHandler {
	return func(ctx context.Context, msg proto.Message) error {
		fac, ok := msg.(Factory)
		if !ok {
			return domain.ErrNoHandlerForCommand
		}

		log.Printf("[Invite] Received factory command of type %T", fac)

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
func NewCommandHandler(store domain.EventStore) domain.CommandHandler {
	return func(ctx context.Context, msg proto.Message) error {
		cmd, ok := msg.(Command)
		if !ok {
			return domain.ErrNoHandlerForCommand
		}

		log.Printf("[Invite] Received command of type %T", cmd)

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
	log.Printf("[Invite] Rehydrating state with %d events", len(events))
	s := &state{}

	for _, e := range events {
		if ue, ok := e.Data().(Event); ok {
			ue.Apply(s)
		}
	}

	return s
}
