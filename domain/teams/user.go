package teams

import (
	"google.golang.org/protobuf/proto"
)

const UserAggregate = "user"

// User is a user who is part of a team.
type User struct {
	userID  string
	teamID  string
	isOwner bool
}

// RehydrateUser rehydrates a User aggregate from an event
// stream
func RehydrateUser(events []proto.Message, userID string) *User {
	u := &User{
		userID: userID,
	}

	for _, e := range events {
		if ue, ok := e.(interface{ Apply(*User) }); ok {
			ue.Apply(u)
		}
	}

	return u
}

// OnLeaderChanged is called when a User receives a LeaderChanged event.
func (u *User) OnLeaderChanged(evt *OwnerChanged) *BecameOwner {
	if u.teamID == evt.GetTeamId() {
		return &BecameOwner{
			UserId: u.userID,
			TeamId: u.teamID,
		}
	}

	return nil
}
