package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getHeadBlockHandler struct {
	provider DbProvider
}

// Handle serves the Get Head Block request.
func (h *getHeadBlockHandler) Handle(params blocks.GetBlocksHeadParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return blocks.NewGetBlocksHeadBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return blocks.NewGetBlocksHeadInternalServerError()
	}

	service := services.New(repos.New(db), net)
	block, err := service.HeadBlock()
	if err != nil {
		logrus.Errorf("failed to get Head block: %s", err.Error())
		return blocks.NewGetBlocksHeadInternalServerError()
	}
	return blocks.NewGetBlocksHeadOK().WithPayload(render.Block(block))
}
