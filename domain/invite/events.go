package invite

import "google.golang.org/protobuf/proto"

type Event interface {
	Apply(*state)
	proto.Message
}

func (e *TeamFormationInviteSent) Apply(s *state) {
	s.id = e.GetInviteId()
	s.inviterUserID = e.GetInviterUserId()
	s.inviteeUserID = e.GetInviteeUserId()
}
