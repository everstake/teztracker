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

type getBlockEndorsementsHandler struct {
	db *gorm.DB
}

// Handle serves the Get Block request.
func (h *getBlockEndorsementsHandler) Handle(params blocks.GetBlockEndorsementsParams) middleware.Responder {
	service := services.New(repos.New(h.db))
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
