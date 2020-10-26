package teams

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/proto"
)

// Execute the CreateTeam command
func (cmd *CreateTeam) Execute(u *User) (proto.Message, error) {
	if u.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", u.userID, u.teamID)
	}

	return &TeamCreated{
		TeamId: cmd.GetTeamId(),
		UserId: cmd.GetUserId(),
	}, nil
}

// Execute the JoinTeam command
func (cmd *JoinTeam) Execute(u *User) (proto.Message, error) {
	if u.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", u.userID, u.teamID)
	}

	return &TeamJoined{
		UserId: cmd.GetUserId(),
		TeamId: cmd.GetTeamId(),
	}, nil
}

// Execute the LeaveTeam command
func (cmd *LeaveTeam) Execute(u *User) (proto.Message, error) {
	if u.teamID == "" {
		return nil, fmt.Errorf("user %s is not in a team", u.userID)
	}
	if u.isOwner {
		return nil, errors.New("team owners cannot leave the team")
	}

	return &TeamLeft{
		UserId: cmd.GetUserId(),
		TeamId: u.teamID,
	}, nil
}

// Execute the ChangeOwner command
func (cmd *ChangeOwner) Execute(u *User) (proto.Message, error) {
	if u.teamID == "" {
		return nil, fmt.Errorf("user %s is not in a team", u.userID)
	}
	if !u.isOwner {
		return nil, fmt.Errorf("user %s is not the team owner", u.userID)
	}

	return &OwnerChanged{
		UserId:         u.userID,
		TeamId:         u.teamID,
		NewOwnerUserId: cmd.GetNewOwnerUserId(),
	}, nil
}

// Execute the DisbandTeam command
func (cmd *DisbandTeam) Execute(u *User) (proto.Message, error) {
	if u.teamID == "" {
		return nil, fmt.Errorf("user %s is not in a team", u.userID)
	}
	if !u.isOwner {
		return nil, fmt.Errorf("user %s is not the team owner", u.userID)
	}

	return &TeamDisbanded{
		UserId: cmd.GetUserId(),
		TeamId: u.teamID,
	}, nil
}
