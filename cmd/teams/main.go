package main

import (
	"flag"
	"net/http"

	"github.com/alexrudd/lb-teams/app/api"
	"github.com/alexrudd/lb-teams/domain/invite"
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

	// domain
	inviteFacHandler := invite.NewFactoryHandler(eventStore)
	inviteCmdHandler := invite.NewCommandHandler(eventStore)
	pendingInvitesView, err := invite.InitialisePendingInvitesView(eventStore)
	if err != nil {
		panic(err)
	}

	// app
	httpHandler := api.NewHTTPInviteHandler(inviteFacHandler, inviteCmdHandler, pendingInvitesView)

	if err := http.ListenAndServe(*httpAddr, httpHandler); err != nil {
		panic(err)
	}
}
