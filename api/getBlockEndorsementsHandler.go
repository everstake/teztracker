package api

import (
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/gen/restapi/operations/blocks"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getBlockEndorsementsHandler struct {
	provider DbProvider
}

// Handle serves the Get Block request.
func (h *getBlockEndorsementsHandler) Handle(params blocks.GetBlockEndorsementsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return blocks.NewGetBlockEndorsementsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return blocks.NewGetBlockEndorsementsInternalServerError()
	}
	service := services.New(repos.New(db), net)
	operations, count, err := service.GetBlockEndorsements(params.Hash)

	if err != nil {
		if err == services.ErrNotFound {
			return blocks.NewGetBlockEndorsementsNotFound()
		}
		logrus.Errorf("failed to get block: %s", err.Error())
		return blocks.NewGetBlockEndorsementsInternalServerError()
	}

	return blocks.NewGetBlockEndorsementsOK().WithPayload(render.Operations(operations)).WithXTotalCount(count)
}
