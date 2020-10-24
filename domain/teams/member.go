package teams

import (
	"context"
	"fmt"

	"github.com/alexrudd/lb-teams/domain"

	"google.golang.org/protobuf/proto"
)

// MemberCommand is a command that can be executed on a Member.
type MemberCommand interface {
	GetUserId() string
	Execute(*Member) (proto.Message, error)
}

// MemberCommandHandler is a function that handles Member commands
type MemberCommandHandler func(context.Context, MemberCommand) error

// Member is a user who is part of a team.
type Member struct {
	userID  string
	teamID  string
	isOwner bool
}

// RehydrateMember rehydrates a Member aggregate from an event
// stream
func RehydrateMember(messages []proto.Message) *Member {
	if len(messages) == 0 {
		return nil
	}

	m := &Member{}

	for _, msg := range messages {
		switch event := msg.(type) {
		case *TeamCreated:
			m.userID = event.GetUserId()
			m.teamID = event.GetTeamId()
			m.isOwner = true
		case *TeamJoined:
			m.userID = event.GetUserId()
			m.teamID = event.GetTeamId()
			m.isOwner = false
		case *TeamLeft:
			m.userID = event.GetUserId()
			m.teamID = ""
			m.isOwner = false
		case *TeamDisbanded:
			m.userID = event.GetUserId()
			m.teamID = ""
			m.isOwner = false
		}
	}

	return m
}

// NewMemberCommandHandler handler member commands.
func NewMemberCommandHandler(store domain.EventStore) MemberCommandHandler {
	return func(ctx context.Context, cmd MemberCommand) error {
		stream, err := store.GetStream(ctx, cmd.GetUserId())
		if err != nil {
			return fmt.Errorf("getting member stream for user %s: %w", cmd.GetUserId(), err)
		}

		// rehydrate or create member
		member := RehydrateMember(stream.Events())
		if member == nil {
			member = &Member{userID: cmd.GetUserId()}
		}

		event, err := cmd.Execute(member)
		if err != nil {
			return fmt.Errorf("executing member command for user %s: %w", cmd.GetUserId(), err)
		}

		if err := stream.Publish(ctx, event); err != nil {
			return fmt.Errorf("publishing member event for user %s: %w", cmd.GetUserId(), err)
		}

		return nil
	}
}
