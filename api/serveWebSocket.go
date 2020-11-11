package api

import (
	"github.com/go-openapi/runtime"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/everstake/teztracker/gen/restapi/operations/w_s"
	"github.com/go-openapi/runtime/middleware"
)

type serveWS struct {
	provider WSProvider
}

func (h *serveWS) Handle(params w_s.ConnectToWSParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return w_s.NewConnectToWSBadRequest()
	}

	hub, err := h.provider.GetWS(net)
	if err != nil {

	}

	return middleware.ResponderFunc(func(rw http.ResponseWriter, _ runtime.Producer) {
		conn, err := hub.GetUpgrader().Upgrade(rw, params.HTTPRequest, nil)
		if err != nil {
			log.Println(err)
			return
		}

		hub.RegisterClient(conn)
	})
}
