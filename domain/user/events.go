package user

import "google.golang.org/protobuf/proto"

type Event interface {
	Apply(*User)
	proto.Message
}

func (e *UserAssignedTeam) Apply(u *User) {
	u.teamID = e.GetTeamId()
}

func (e *UserInvitedToTeam) Apply(u *User) {
	u.invites[e.GetTeamId()] = struct{}{}
}

func (e *UserAcceptedInvite) Apply(u *User) {
	delete(u.invites, e.GetTeamId())
	u.teamID = e.GetTeamId()
}

func (e *UserLeftTeam) Apply(u *User) {
	u.teamID = ""
}
