package views

import (
	"context"
	"fmt"
	"log"

	"github.com/alexrudd/lb-teams/domain"
	"github.com/alexrudd/lb-teams/domain/invite"
)

type PendingInvitesView struct {
	userInboxes map[string]*UserInbox
}

func InitialisePendingInvitesView(store domain.EventStore) (*PendingInvitesView, error) {
	stream, err := store.GetAggregateStream(context.Background(), invite.Name)
	if err != nil {
		return nil, fmt.Errorf("getting aggregate stream for %s: %w", invite.Name, err)
	}

	piv := &PendingInvitesView{
		userInboxes: map[string]*UserInbox{},
	}

	if err := stream.Bind(piv.eventHandler); err != nil {
		return nil, fmt.Errorf("binding to aggregate stream for %s: %w", invite.Name, err)
	}

	return piv, nil
}

func (piv *PendingInvitesView) GetUserInbox(id string) UserInbox {
	if i, ok := piv.userInboxes[id]; ok {
		return *i
	}

	return UserInbox{
		userID: id,
	}
}

func (piv *PendingInvitesView) eventHandler(ctx context.Context, event domain.Event) {
	switch e := event.Data().(type) {
	case *invite.TeamFormationInviteSent:
		if _, ok := piv.userInboxes[e.GetInviteeUserId()]; !ok {
			piv.userInboxes[e.GetInviteeUserId()] = &UserInbox{
				userID:  e.GetInviteeUserId(),
				invites: map[string]*Invite{},
			}
		}

		piv.userInboxes[e.GetInviteeUserId()].invites[e.GetInviteId()] = &Invite{
			id:            e.GetInviteId(),
			inviterUserID: e.GetInviterUserId(),
			inviteeUserID: e.GetInviteeUserId(),
		}
	case *invite.TeamFormationInviteDeclined:
		if _, ok := piv.userInboxes[e.GetInviteeUserId()]; ok {
			delete(piv.userInboxes[e.GetInviteeUserId()].invites, e.GetInviteId())
		}
	case *invite.TeamFormationInviteCancelled:
		if _, ok := piv.userInboxes[e.GetInviteeUserId()]; ok {
			delete(piv.userInboxes[e.GetInviteeUserId()].invites, e.GetInviteId())
		}
	case *invite.TeamFormationInviteExpired:
		if _, ok := piv.userInboxes[e.GetInviteeUserId()]; ok {
			delete(piv.userInboxes[e.GetInviteeUserId()].invites, e.GetInviteId())
		}
	case *invite.TeamFormationInviteAccepted:
		if _, ok := piv.userInboxes[e.GetInviteeUserId()]; ok {
			delete(piv.userInboxes[e.GetInviteeUserId()].invites, e.GetInviteId())
		}
	default:
		log.Printf("[PendingInvitesView] ignoring event of type %T", event.Data())
	}
}

type UserInbox struct {
	userID  string
	invites map[string]*Invite
}

func (ui *UserInbox) UserID() string {
	return ui.userID
}

func (ui *UserInbox) Invites() []Invite {
	var out []Invite

	for _, i := range ui.invites {
		out = append(out, *i)
	}

	return out
}

type Invite struct {
	id            string
	teamID        string
	inviterUserID string
	inviteeUserID string
}

func (i *Invite) ID() string {
	return i.id
}

func (i *Invite) TeamID() string {
	return i.teamID
}

func (i *Invite) InviterUserID() string {
	return i.inviterUserID
}
