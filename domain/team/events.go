package team

import "google.golang.org/protobuf/proto"

type Event interface {
	Apply(*state)
	proto.Message
}

func (e *TeamFormed) Apply(s *state) {
	s.id = e.GetMemberUserId()
	s.ownerUserID = e.GetOwnerUserId()
	s.memberUserIDs[e.GetMemberUserId()] = struct{}{}
}
