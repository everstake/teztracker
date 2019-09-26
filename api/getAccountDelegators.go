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

type getAccountDelegatorsHandler struct {
	db *gorm.DB
}

// Handle serves the Get Account Delegators request.
func (h *getAccountDelegatorsHandler) Handle(params accounts.GetAccountDelegatorsParams) middleware.Responder {
	service := services.New(repos.New(h.db))
	limiter := NewLimiter(params.Limit, params.Offset)
	accs, count, err := service.AccountDelegatorsList(params.AccountID, limiter)
	if err != nil {
		logrus.Errorf("failed to get account's delegators: %s", err.Error())
		return accounts.NewGetAccountInternalServerError()
	}
	return accounts.NewGetAccountDelegatorsOK().WithPayload(render.Accounts(accs)).WithXTotalCount(count)
}
