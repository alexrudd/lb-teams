package team

import (
	"context"
	"fmt"
	"log"

	"github.com/alexrudd/lb-teams/domain"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

const Name = "team"

// state stores the aggregate for a particular team stream.
type state struct {
	id            string
	ownerUserID   string
	memberUserIDs map[string]struct{}
}

func RegisterWithCommandBus(bus domain.CommandBus, store domain.EventStore) {
	var (
		facHandler = NewFactoryHandler(store)
		// cmdHandler = NewCommandHandler(store)
	)

	bus.RegisterHandler(&FormTeam{}, facHandler)
}

// NewFactoryHandler handler team foactor commands.
func NewFactoryHandler(store domain.EventStore) domain.CommandHandler {
	return func(ctx context.Context, msg proto.Message) error {
		fac, ok := msg.(Factory)
		if !ok {
			return domain.ErrNoHandlerForCommand
		}

		log.Printf("[Team] Received factory command of type %T", fac)

		if err := fac.Validate(); err != nil {
			return fmt.Errorf("validating factory command for team: %w", err)
		}

		teamID := uuid.New().String()

		event := fac.InitialEvent(teamID)

		stream, err := store.GetStream(ctx, Name, teamID)
		if err != nil {
			return fmt.Errorf("getting stream for team %s: %w", teamID, err)
		}

		if err := stream.Publish(ctx, event); err != nil {
			return fmt.Errorf("publishing event for team %s: %w", teamID, err)
		}

		return nil
	}
}

// NewCommandHandler handles team commands.
func NewCommandHandler(store domain.EventStore) domain.CommandHandler {
	return func(ctx context.Context, msg proto.Message) error {
		cmd, ok := msg.(Command)
		if !ok {
			return domain.ErrNoHandlerForCommand
		}

		log.Printf("[Team] Received command of type %T", cmd)

		if err := cmd.Validate(); err != nil {
			return fmt.Errorf("validating command for team %s: %w", cmd.GetTeamId(), err)
		}

		stream, err := store.GetStream(ctx, Name, cmd.GetTeamId())
		if err != nil {
			return fmt.Errorf("getting stream for team %s: %w", cmd.GetTeamId(), err)
		}

		// rehydrate state
		team := RehydrateState(stream.Events())
		event, err := cmd.Execute(team)
		if err != nil {
			return fmt.Errorf("executing command for team %s: %w", cmd.GetTeamId(), err)
		}

		if !event.ProtoReflect().IsValid() {
			return nil
		}

		if err := stream.Publish(ctx, event); err != nil {
			return fmt.Errorf("publishing event for team %s: %w", cmd.GetTeamId(), err)
		}

		return nil
	}
}

// RehydrateState rehydrates an team's state.
func RehydrateState(events []domain.Event) *state {
	log.Printf("[Team] Rehydrating state with %d events", len(events))
	s := &state{}

	for _, e := range events {
		if ue, ok := e.Data().(Event); ok {
			ue.Apply(s)
		}
	}

	return s
}
