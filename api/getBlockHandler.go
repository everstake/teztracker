package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBlockHandler struct {
	provider DbProvider
}

// Handle serves the Get Block request.
func (h *getBlockHandler) Handle(params blocks.GetBlockParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return blocks.NewGetBlockBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return blocks.NewGetBlockInternalServerError()
	}
	service := services.New(repos.New(db), net)
	var block models.Block
	if params.Hash == "head" {
		block, err = service.HeadBlock()
	} else {
		block, err = service.GetBlockWithOperationGroups(params.Hash)
	}

	if err != nil {
		if err == services.ErrNotFound {
			return blocks.NewGetBlockNotFound()
		}
		logrus.Errorf("failed to get block: %s", err.Error())
		return blocks.NewGetBlockInternalServerError()
	}

	return blocks.NewGetBlockOK().WithPayload(render.BlockResult(block))
}
