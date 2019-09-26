package api

import (
	"github.com/bullblock-io/tezTracker/api/render"
	"github.com/bullblock-io/tezTracker/gen/restapi/operations/accounts"
	"github.com/bullblock-io/tezTracker/repos"
	"github.com/bullblock-io/tezTracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type getAccountListHandler struct {
	db *gorm.DB
}

// Handle serves the Get Account List request.
func (h *getAccountListHandler) Handle(params accounts.GetAccountsListParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	limiter := NewLimiter(params.Limit, params.Offset)
	before := ""
	if params.AfterID != nil {
		before = *params.AfterID
	}
	accs, count, err := service.AccountList(before, limiter)
	if err != nil {
		logrus.Errorf("failed to get accounts: %s", err.Error())
		return accounts.NewGetAccountsListNotFound()
	}
	return accounts.NewGetAccountsListOK().WithPayload(render.Accounts(accs)).WithXTotalCount(count)
}
