package teams

// Apply the TeamCreated event.
func (evt *TeamCreated) Apply(u *User) {
	u.teamID = evt.GetTeamId()
	u.isOwner = true
}

// Apply the TeamJoined event.
func (evt *TeamJoined) Apply(u *User) {
	u.teamID = evt.GetTeamId()
	u.isOwner = false
}

// Apply the TeamLeft event.
func (evt *TeamLeft) Apply(u *User) {
	u.teamID = ""
	u.isOwner = false
}

// Apply the LeaderChanged event.
func (evt *OwnerChanged) Apply(u *User) {
	u.teamID = evt.GetTeamId()
	u.isOwner = false
}

// Apply the LeaderChanged event.
func (evt *BecameOwner) Apply(u *User) {
	u.teamID = evt.GetTeamId()
	u.isOwner = true
}

// Apply the TeamDisbanded event.
func (evt *TeamDisbanded) Apply(u *User) {
	u.teamID = ""
	u.isOwner = false
}
