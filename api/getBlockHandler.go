package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/blocks"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type getBlockHandler struct {
	db *gorm.DB
}

// Handle serves the Get Block request.
func (h *getBlockHandler) Handle(params blocks.GetBlockParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	block, err := service.GetBlockWithOperationGroups(params.Hash)

	if err != nil {
		if err == services.ErrNotFound {
			return blocks.NewGetBlockNotFound()
		}
		logrus.Errorf("failed to get block: %s", err.Error())
		return blocks.NewGetBlockInternalServerError()
	}

	return blocks.NewGetBlockOK().WithPayload(render.BlockResult(block))
}
