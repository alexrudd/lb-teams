package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/alexrudd/lb-teams/app/api"
	"github.com/alexrudd/lb-teams/domain/invite"
	"github.com/alexrudd/lb-teams/domain/team"
	"github.com/alexrudd/lb-teams/domain/views"
	"github.com/alexrudd/lb-teams/infra/im"
	"github.com/alexrudd/lb-teams/infra/lb"

	lift "github.com/liftbridge-io/go-liftbridge/v2"
)

func main() {
	var (
		liftbridgeAddr = flag.String("liftbridge-addr", "liftbridge:9292", "the address of a liftbridge cluster")
		httpAddr       = flag.String("http-addr", ":8080", "the address to listen for commands on")
	)

	// infra
	lbc, err := lift.Connect([]string{*liftbridgeAddr})
	if err != nil {
		panic(err)
	}

	eventStore := lb.NewLiftBridgeEventStore(lbc)
	defer eventStore.Close()

	cmdBus := im.NewBus()

	// domain
	invite.RegisterWithCommandBus(cmdBus, eventStore)
	team.RegisterWithCommandBus(cmdBus, eventStore)

	// views
	pendingInvitesView, err := views.InitialisePendingInvitesView(eventStore)
	if err != nil {
		panic(err)
	}
	teamFormationProcessor, err := views.NewTeamsToFormProcessor(eventStore)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			teamFormationProcessor.Run(cmdBus)
			time.Sleep(time.Second)
		}
	}()

	// app
	httpHandler := api.NewHTTPInviteHandler(cmdBus, pendingInvitesView)

	if err := http.ListenAndServe(*httpAddr, httpHandler); err != nil {
		panic(err)
	}
}
