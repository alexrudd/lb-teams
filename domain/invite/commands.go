package invite

import "errors"

type Factory interface {
	Validate() error
	InitialEvent(id string) Event
}

type Command interface {
	Validate() error
	GetInviteId() string
	Execute(*state) (Event, error)
}

func (c *SendTeamFormationInvite) Validate() error {
	if c.GetInviterUserId() == "" {
		return errors.New("command must specify and inviter")
	} else if c.GetInviteeUserId() == "" {
		return errors.New("command must specify and invitee")
	}

	return nil
}

func (c *SendTeamFormationInvite) InitialEvent(id string) Event {
	return &TeamFormationInviteSent{
		InviteId:      id,
		InviterUserId: c.GetInviterUserId(),
		InviteeUserId: c.GetInviteeUserId(),
	}
}
