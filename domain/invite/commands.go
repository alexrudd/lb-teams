package invite

import (
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Factory interface {
	Validate() error
	InitialEvent(id string) Event
}

type Command interface {
	// Authorize(identity domain.Identity) error
	Validate() error
	GetInviteId() string
	Execute(*state) (Event, error)
}

// Form team
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

// Decline formation
func (c *DeclineTeamFormationInvite) Validate() error {
	if c.GetInviteId() == "" {
		return errors.New("command must specify an invite ID")
	}

	return nil
}

func (c *DeclineTeamFormationInvite) Execute(s *state) (Event, error) {
	if s.status != pending {
		return nil, errors.New("invite must be in pending state to be declined")
	}

	return &TeamFormationInviteDeclined{
		InviteId:      c.GetInviteId(),
		InviterUserId: s.inviterUserID,
		InviteeUserId: s.inviteeUserID,
		Timestamp:     timestamppb.Now(),
	}, nil
}

// Cancel formation
func (c *CancelTeamFormationInvite) Validate() error {
	if c.GetInviteId() == "" {
		return errors.New("command must specify an invite ID")
	}

	return nil
}

func (c *CancelTeamFormationInvite) Execute(s *state) (Event, error) {
	if s.status != pending {
		return nil, errors.New("invite must be in pending state to be cancelled")
	}

	return &TeamFormationInviteCancelled{
		InviteId:      c.GetInviteId(),
		InviterUserId: s.inviterUserID,
		InviteeUserId: s.inviteeUserID,
		Timestamp:     timestamppb.Now(),
	}, nil
}

// Formation Expired
func (c *ExpireTeamFormationInvite) Validate() error {
	if c.GetInviteId() == "" {
		return errors.New("command must specify an invite ID")
	}

	return nil
}

func (c *ExpireTeamFormationInvite) Execute(s *state) (Event, error) {
	if s.status != pending {
		return nil, errors.New("invite must be in pending state to be expired")
	}

	return &TeamFormationInviteExpired{
		InviteId:      c.GetInviteId(),
		InviterUserId: s.inviterUserID,
		InviteeUserId: s.inviteeUserID,
		Timestamp:     timestamppb.Now(),
	}, nil
}

// Formation Accepted
func (c *AcceptTeamFormationInvite) Validate() error {
	if c.GetInviteId() == "" {
		return errors.New("command must specify an invite ID")
	}

	return nil
}

func (c *AcceptTeamFormationInvite) Execute(s *state) (Event, error) {
	if s.status != pending {
		return nil, errors.New("invite must be in pending state to be accepted")
	}

	return &TeamFormationInviteAccepted{
		InviteId:      c.GetInviteId(),
		InviterUserId: s.inviterUserID,
		InviteeUserId: s.inviteeUserID,
		Timestamp:     timestamppb.Now(),
	}, nil
}
