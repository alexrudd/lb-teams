package team

import (
	"errors"
)

type Factory interface {
	Validate() error
	InitialEvent(id string) Event
}

type Command interface {
	// Authorize(identity domain.Identity) error
	Validate() error
	GetTeamId() string
	Execute(*state) (Event, error)
}

// Form team
func (c *FormTeam) Validate() error {
	if c.GetOwnerUserId() == "" {
		return errors.New("command must specify an owner")
	} else if c.GetMemberUserId() == "" {
		return errors.New("command must specify a member user ID")
	}

	return nil
}

func (c *FormTeam) InitialEvent(id string) Event {
	return &TeamFormed{
		OwnerUserId:  c.GetOwnerUserId(),
		MemberUserId: c.GetMemberUserId(),
	}
}
