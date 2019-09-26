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

type getBlockListHandler struct {
	db *gorm.DB
}

// Handle serves the Get Block List request.
func (h *getBlockListHandler) Handle(params blocks.GetBlocksListParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	limiter := NewLimiter(params.Limit, params.Offset)
	before := uint64(0)
	if params.BeforeLevel != nil {
		before = uint64(*params.BeforeLevel)
	}
	bs, count, err := service.BlockList(before, limiter)
	if err != nil {
		logrus.Errorf("failed to get blocks: %s", err.Error())
		return blocks.NewGetBlocksListNotFound()
	}

	return blocks.NewGetBlocksListOK().WithPayload(render.Blocks(bs)).WithXTotalCount(count)
}
