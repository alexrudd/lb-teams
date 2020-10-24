package main

import (
	"flag"
	"net/http"

	"github.com/alexrudd/lb-teams/app/api"
	"github.com/alexrudd/lb-teams/domain/teams"
	"github.com/alexrudd/lb-teams/infra/lb"
)

func main() {
	var (
		liftbridgeAddr = flag.String("liftbridge-addr", "liftbridge:9292", "the address of a liftbridge cluster")
		httpAddr       = flag.String("http-addr", ":8080", "the address to listen for commands on")
	)

	// infra
	eventStore, err := lb.NewLiftBridgeEventStore([]string{*liftbridgeAddr})
	if err != nil {
		panic(err)
	}

	// domain
	teamsHandler := teams.NewMemberCommandHandler(eventStore)

	// app
	httpHandler := api.NewHTTPTeamsHandler(teamsHandler)

	if err := http.ListenAndServe(*httpAddr, httpHandler); err != nil {
		panic(err)
	}
}
