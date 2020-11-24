package views

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alexrudd/lb-teams/domain"
	"github.com/alexrudd/lb-teams/domain/invite"
	"github.com/alexrudd/lb-teams/domain/team"
)

type TeamsFormationProcessor struct {
	todo map[string]*teamFormationJob
}

type teamFormationJob struct {
	ownerUserID   string
	memberUserID  string
	lastProcessed time.Time
}

func NewTeamsToFormProcessor(store domain.EventStore) (*TeamsFormationProcessor, error) {
	tfp := &TeamsFormationProcessor{
		todo: map[string]*teamFormationJob{},
	}

	stream, err := store.GetAggregateStream(context.Background(), invite.Name)
	if err != nil {
		return nil, fmt.Errorf("getting aggregate stream for %s: %w", invite.Name, err)
	}

	if err := stream.Bind(tfp.eventHandler); err != nil {
		return nil, fmt.Errorf("binding to aggregate stream for %s: %w", invite.Name, err)
	}

	stream, err = store.GetAggregateStream(context.Background(), team.Name)
	if err != nil {
		return nil, fmt.Errorf("getting aggregate stream for %s: %w", team.Name, err)
	}

	if err := stream.Bind(tfp.eventHandler); err != nil {
		return nil, fmt.Errorf("binding to aggregate stream for %s: %w", team.Name, err)
	}

	return tfp, nil
}

func (tfp *TeamsFormationProcessor) Run(bus domain.CommandBus) {
	for _, job := range tfp.todo {
		if job.lastProcessed.Before(time.Now().Add(-100 * time.Millisecond)) {
			log.Printf("Sending FormSquad command for %s and %s", job.ownerUserID, job.memberUserID)

			if err := bus.SubmitCommand(context.Background(), &team.FormTeam{
				OwnerUserId:  job.ownerUserID,
				MemberUserId: job.memberUserID,
			}); err != nil {
				log.Printf("submitting FormSquad command: %s", err)
			}

			job.lastProcessed = time.Now()
		}
	}
}

func (tfp *TeamsFormationProcessor) eventHandler(ctx context.Context, event domain.Event) {
	switch e := event.Data().(type) {
	case *invite.TeamFormationInviteAccepted:
		tfp.todo[e.GetInviterUserId()] = &teamFormationJob{
			ownerUserID:  e.GetInviterUserId(),
			memberUserID: e.GetInviteeUserId(),
		}
	case *team.TeamFormed:
		delete(tfp.todo, e.GetOwnerUserId())
	default:
		log.Printf("[TeamsFormationProcess] ignoring event of type %T", event.Data())
	}
}
