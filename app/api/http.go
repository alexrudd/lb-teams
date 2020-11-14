package api

import (
	"io/ioutil"
	"net/http"

	"github.com/alexrudd/lb-teams/domain/user"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// NewHTTPUserCommandHandler returns an http.Handler for receiving user
// commands.
func NewHTTPUserCommandHandler(handler user.CommandHandler) http.Handler {
	mux := http.NewServeMux()

	handleRequest := func(rw http.ResponseWriter, r *http.Request, msg proto.Message) {
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		err = protojson.Unmarshal(b, msg)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		cmd, ok := msg.(user.Command)
		if !ok {
			http.Error(rw, "request does not implement user.Command", http.StatusBadRequest)
			return
		}

		if err = handler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// InviteUserToTeam
	mux.Handle("/InviteUserToTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &user.InviteUserToTeam{})
	}))

	// AcceptInvite
	mux.Handle("/AcceptInvite", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &user.AcceptInvite{})
	}))

	// LeaveTeam
	mux.Handle("/LeaveTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &user.LeaveTeam{})
	}))

	return mux
}
