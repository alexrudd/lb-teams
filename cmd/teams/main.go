package main

import (
	"flag"
	"net/http"

	"github.com/alexrudd/lb-teams/app/api"
	"github.com/alexrudd/lb-teams/domain/teams"
	"github.com/alexrudd/lb-teams/infra/lb"
	"github.com/nats-io/nats.go"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
)

func main() {
	var (
		natsAddr       = flag.String("nats-addr", "nats://liftbridge:4222", "the address of a nats cluster")
		liftbridgeAddr = flag.String("liftbridge-addr", "liftbridge:9292", "the address of a liftbridge cluster")
		httpAddr       = flag.String("http-addr", ":8080", "the address to listen for commands on")
	)

	// infra
	nc, err := nats.Connect(*natsAddr)
	if err != nil {
		panic(err)
	}

	lbc, err := lift.Connect([]string{*liftbridgeAddr})
	if err != nil {
		panic(err)
	}

	eventStore := lb.NewLiftBridgeEventStore(nc, lbc)

	// domain
	teamsHandler := teams.NewUserCommandHandler(eventStore)
	if err := teams.SetupEventHandlers(eventStore); err != nil {
		panic(err)
	}

	// app
	httpHandler := api.NewHTTPTeamsHandler(teamsHandler)

	if err := http.ListenAndServe(*httpAddr, httpHandler); err != nil {
		panic(err)
	}
}
