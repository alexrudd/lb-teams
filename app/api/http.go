package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alexrudd/lb-teams/domain/invite"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type InviteFactoryHandler func(context.Context, invite.Factory) error
type InviteCommandHandler func(context.Context, invite.Command) error

// NewHTTPInviteHandler returns an http.Handler for receiving invite commands
// and queries.
func NewHTTPInviteHandler(
	inviteFactoryHandler InviteFactoryHandler,
	inviteCommandHandler InviteCommandHandler,
	view *invite.PendingInvitesView,
) http.Handler {
	mux := http.NewServeMux()
	// SendTeamFormationInvite
	mux.Handle("/SendTeamFormationInvite", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "this endpoint only accepts POST requests", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		cmd := &invite.SendTeamFormationInvite{}

		err = protojson.Unmarshal(b, cmd)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = inviteFactoryHandler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	handler := func(msg func() proto.Message) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(rw, "this endpoint only accepts POST requests", http.StatusMethodNotAllowed)
				return
			}

			defer r.Body.Close()
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}

			m := msg()

			err = protojson.Unmarshal(b, m)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusBadRequest)
				return
			}

			if err = inviteCommandHandler(r.Context(), m.(invite.Command)); err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}
	// DeclineTeamFormationInvite
	mux.Handle("/DeclineTeamFormationInvite", handler(func() proto.Message { return &invite.DeclineTeamFormationInvite{} }))
	// CancelTeamFormationInvite
	mux.Handle("/CancelTeamFormationInvite", handler(func() proto.Message { return &invite.CancelTeamFormationInvite{} }))
	// AcceptTeamFormationInvite
	mux.Handle("/AcceptTeamFormationInvite", handler(func() proto.Message { return &invite.AcceptTeamFormationInvite{} }))

	// Query Inbox
	mux.Handle("/Inbox", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(rw, "this endpoint only accepts GET requests", http.StatusMethodNotAllowed)
			return
		}

		userID := r.URL.Query().Get("userid")
		if userID == "" {
			http.Error(rw, "must specify url query paramter userid", http.StatusBadRequest)
			return
		}

		inbox := transform(view.GetUserInbox(userID))

		b, err := json.Marshal(inbox)
		if err != nil {
			http.Error(rw, fmt.Sprintf("marshalling response: %s", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(rw, string(b))
	}))

	return mux
}

type serializableInbox struct {
	UserID  string
	Invites []serializableInvite
}

func transform(in invite.UserInbox) serializableInbox {
	out := serializableInbox{
		UserID: in.UserID(),
	}

	for _, i := range in.Invites() {
		out.Invites = append(out.Invites, serializableInvite{
			ID:            i.ID(),
			TeamID:        i.TeamID(),
			InviterUserID: i.InviterUserID(),
		})
	}

	return out
}

type serializableInvite struct {
	ID            string
	TeamID        string
	InviterUserID string
}
