package api

import (
	"github.com/bullblock-io/tezTracker/gen/restapi/operations"
	"github.com/bullblock-io/tezTracker/services/cmc"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// SetHandlers initializes the API handlers with underlying services.
func SetHandlers(serv *operations.TezTrackerAPI, db *gorm.DB) {
	serv.Logger = logrus.Infof
	serv.BlocksGetBlocksHeadHandler = &getHeadBlockHandler{db}
	serv.BlocksGetBlocksListHandler = &getBlockListHandler{db}
	serv.BlocksGetBlockHandler = &getBlockHandler{db}
	serv.OperationsListGetOperationsListHandler = &getOperationListHandler{db}
	serv.AppInfoGetInfoHandler = &getInfoHandler{&cmc.CoinGecko{}, db}
	serv.AccountsGetAccountsListHandler = &getAccountListHandler{db}
	serv.AccountsGetAccountHandler = &getAccountHandler{db}
	serv.BlocksGetBlockEndorsementsHandler = &getBlockEndorsementsHandler{db}
	serv.AccountsGetBakersListHandler = &getBakerListHandler{db}
	serv.AccountsGetAccountDelegatorsHandler = &getAccountDelegatorsHandler{db}
}
