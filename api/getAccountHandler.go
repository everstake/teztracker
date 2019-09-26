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

type getAccountHandler struct {
	db *gorm.DB
}

// Handle serves the Get Account request.
func (h *getAccountHandler) Handle(params accounts.GetAccountParams) middleware.Responder {
	service := services.New(repos.New(h.db))

	acc, err := service.GetAccount(params.AccountID)

	if err != nil {
		if err == services.ErrNotFound {
			return accounts.NewGetAccountNotFound()
		}
		logrus.Errorf("failed to get acc: %s", err.Error())
		return accounts.NewGetAccountInternalServerError()
	}

	return accounts.NewGetAccountOK().WithPayload(render.Account(acc))
}
