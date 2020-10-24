package teams

import (
	"fmt"

	"github.com/gofrs/uuid"
	"google.golang.org/protobuf/proto"
)

// Execute the CreateTeam command
func (cmd *CreateTeam) Execute(m *Member) (proto.Message, error) {
	if m.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", m.userID, m.teamID)
	}

	return &TeamCreated{
		TeamId: uuid.Must(uuid.NewV4()).String(),
		UserId: cmd.GetUserId(),
	}, nil
}

// Execute the JoinTeam command
func (cmd *JoinTeam) Execute(m *Member) (proto.Message, error) {
	if m.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", m.userID, m.teamID)
	}

	return &TeamJoined{
		UserId: cmd.GetUserId(),
		TeamId: uuid.Must(uuid.NewV4()).String(),
	}, nil
}

// Execute the LeaveTeam command
func (cmd *LeaveTeam) Execute(m *Member) (proto.Message, error) {
	if m.teamID == "" {
		return nil, fmt.Errorf("user %s is not in a team", m.userID)
	}

	return &TeamLeft{
		UserId: cmd.GetUserId(),
		TeamId: uuid.Must(uuid.NewV4()).String(),
	}, nil
}

// Execute the DisbandTeam command
func (cmd *DisbandTeam) Execute(m *Member) (proto.Message, error) {
	if m.teamID == "" {
		return nil, fmt.Errorf("user %s is not in a team", m.userID)
	}
	if !m.isOwner {
		return nil, fmt.Errorf("user %s is not the team owner", m.userID)
	}

	return &TeamDisbanded{
		UserId: cmd.GetUserId(),
		TeamId: uuid.Must(uuid.NewV4()).String(),
	}, nil
}
