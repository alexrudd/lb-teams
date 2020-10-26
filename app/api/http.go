package api

import (
	"io/ioutil"
	"net/http"

	"github.com/alexrudd/lb-teams/domain/teams"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewHTTPTeamsHandler returns an http.Handler for receiving teams commands.
func NewHTTPTeamsHandler(handler teams.UserCommandHandler) http.Handler {
	mux := http.NewServeMux()

	handleRequest := func(rw http.ResponseWriter, r *http.Request, cmd teams.UserCommand) {
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		err = protojson.Unmarshal(b, cmd)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = handler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// CreateTeam
	mux.Handle("/CreateTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &teams.CreateTeam{})
	}))

	// JoinTeam
	mux.Handle("/JoinTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &teams.JoinTeam{})
	}))

	// LeaveTeam
	mux.Handle("/LeaveTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &teams.LeaveTeam{})
	}))

	// ChangeOwner
	mux.Handle("/ChangeOwner", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &teams.ChangeOwner{})
	}))

	// DisbandTeam
	mux.Handle("/DisbandTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		handleRequest(rw, r, &teams.DisbandTeam{})
	}))

	return mux
}
