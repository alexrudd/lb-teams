package invite

import "google.golang.org/protobuf/proto"

type Event interface {
	Apply(*state)
	proto.Message
}

func (e *TeamFormationInviteSent) Apply(s *state) {
	s.id = e.GetInviteId()
	s.status = pending
	s.inviterUserID = e.GetInviterUserId()
	s.inviteeUserID = e.GetInviteeUserId()
}

func (e *TeamInviteSent) Apply(s *state) {
	s.id = e.GetInviteId()
	s.status = pending
	s.teamID = e.GetTeamId()
	s.inviterUserID = e.GetInviterUserId()
	s.inviteeUserID = e.GetInviteeUserId()
}

func (e *TeamFormationInviteDeclined) Apply(s *state) {
	s.status = declined
}

func (e *TeamFormationInviteCancelled) Apply(s *state) {
	s.status = cancelled
}

func (e *TeamFormationInviteExpired) Apply(s *state) {
	s.status = expired
}

func (e *TeamFormationInviteAccepted) Apply(s *state) {
	s.status = accepted
}

func (e *TeamInviteDeclined) Apply(s *state) {
	s.status = declined
}

func (e *TeamInviteCancelled) Apply(s *state) {
	s.status = cancelled
}

func (e *TeamInviteExpired) Apply(s *state) {
	s.status = expired
}

func (e *TeamInviteAccepted) Apply(s *state) {
	s.status = accepted
}
