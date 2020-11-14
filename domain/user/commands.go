package user

import "fmt"

// Command is a command that can be executed on a User.
type Command interface {
	GetUserId() string
	Execute(*User) (Event, error)
}

// Executes an AssignTeamToUser command
func (c *AssignTeamToUser) Execute(u *User) (Event, error) {
	if u.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", u.userID, u.teamID)
	}

	return &UserAssignedTeam{
		UserId: c.GetUserId(),
		TeamId: c.GetTeamId(),
	}, nil
}

// Executes an InviteUserToTeam command
func (c *InviteUserToTeam) Execute(u *User) (Event, error) {
	if u.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", u.userID, u.teamID)
	}

	return &UserInvitedToTeam{
		UserId: c.GetUserId(),
		TeamId: c.GetTeamId(),
	}, nil
}

// Executes an AcceptInvite command
func (c *AcceptInvite) Execute(u *User) (Event, error) {
	if u.teamID != "" {
		return nil, fmt.Errorf("user %s is already in a team (%s)", u.userID, u.teamID)
	}

	if _, invited := u.invites[c.GetTeamId()]; !invited {
		return nil, fmt.Errorf("user %s does not have an invite to team %s", u.userID, c.GetTeamId())
	}

	return &UserAcceptedInvite{
		UserId: c.GetUserId(),
		TeamId: c.GetTeamId(),
	}, nil
}

// Executes a LeaveTeam command
func (c *LeaveTeam) Execute(u *User) (Event, error) {
	if u.teamID == "" {
		return nil, fmt.Errorf("user %s is not in a team", u.userID)
	}

	return &UserLeftTeam{
		UserId: c.GetUserId(),
		TeamId: u.teamID,
	}, nil
}
