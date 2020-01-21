package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getSnapshotsHandler struct {
	provider DbProvider
}

// Handle serves the Get Block List request.
func (h *getSnapshotsHandler) Handle(params operations.GetSnapshotsParams) middleware.Responder {
	net, err := ToNetwork(params.Network)
	if err != nil {
		return operations.NewGetSnapshotsBadRequest()
	}
	db, err := h.provider.GetDb(net)
	if err != nil {
		return operations.NewGetSnapshotsNotFound()
	}
	service := services.New(repos.New(db), net)

	limiter := NewLimiter(params.Limit, params.Offset)
	count, bs, err := service.Snapshots(limiter)
	if err != nil {
		logrus.Errorf("failed to get snapshots: %s", err.Error())
		return operations.NewGetSnapshotsNotFound()
	}

	return operations.NewGetSnapshotsOK().WithPayload(render.Snapshots(bs)).WithXTotalCount(count)
}
