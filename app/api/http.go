package api

import (
	"net/http"

	"github.com/alexrudd/lb-teams/domain/teams"
	"github.com/golang/protobuf/jsonpb"
)

// NewHTTPTeamsHandler returns an http.Handler for receiving teams commands.
func NewHTTPTeamsHandler(handler teams.MemberCommandHandler) http.Handler {
	mux := http.NewServeMux()

	// CreateTeam
	mux.Handle("/CreateTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cmd := &teams.CreateTeam{}

		err := jsonpb.Unmarshal(r.Body, cmd)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = handler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	// JoinTeam
	mux.Handle("/JoinTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cmd := &teams.JoinTeam{}

		err := jsonpb.Unmarshal(r.Body, cmd)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = handler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	// LeaveTeam
	mux.Handle("/LeaveTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cmd := &teams.LeaveTeam{}

		err := jsonpb.Unmarshal(r.Body, cmd)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = handler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	// DisbandTeam
	mux.Handle("/DisbandTeam", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		cmd := &teams.DisbandTeam{}

		err := jsonpb.Unmarshal(r.Body, cmd)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err = handler(r.Context(), cmd); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	return mux
}
