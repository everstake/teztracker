package api

import (
	"github.com/everstake/teztracker/gen/restapi/operations"
	"github.com/everstake/teztracker/infrustructure"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/cmc"
	"github.com/everstake/teztracker/services/mempool"
	"github.com/everstake/teztracker/ws"
	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type DbProvider interface {
	GetDb(models.Network) (*gorm.DB, error)
}

type MempoolProvider interface {
	GetMempool(net models.Network) (*mempool.Mempool, error)
}

type WSProvider interface {
	GetWS(models.Network) (*ws.Hub, error)
}

// SetHandlers initializes the API handlers with underlying services.
func SetHandlers(serv *operations.TezTrackerAPI, db *infrustructure.Provider) {
	serv.Logger = logrus.Infof
	serv.BlocksGetBlocksHeadHandler = &getHeadBlockHandler{db}
	serv.BlocksGetBlocksListHandler = &getBlockListHandler{db}
	serv.BlocksGetBlockEndorsementsHandler = &getBlockEndorsementsHandler{db}
	serv.BlocksGetBlockHandler = &getBlockHandler{db}
	serv.OperationsListGetOperationsListHandler = &getOperationListHandler{db}
	serv.AppInfoGetInfoHandler = &getInfoHandler{cmc.NewCoinGecko(), db, cache.New(cacheTTL, cacheTTL)}
	serv.AppInfoGetChartsInfoHandler = &getChartsInfoHandler{db}
	serv.AppInfoGetHealthCheckInfoHandler = &getHealthHandler{db}
	serv.AppInfoGetBakerChartInfoHandler = &getBakerChartHandler{db}
	serv.AppInfoGetBlocksPriorityChartInfoHandler = &getBlocksPriorityHandler{db}
	//Account
	serv.AccountsGetAccountsListHandler = &getAccountListHandler{db}
	serv.AccountsGetAccountsTopBalanceListHandler = &getAccountTopBalanceListHandler{db}
	serv.AccountsGetAccountHandler = &getAccountHandler{db}
	serv.AccountsGetAccountBalanceListHandler = &getAccountBalanceListHandler{db}
	serv.AccountsGetAccountBakingListHandler = &getAccountBakingListHandler{db}
	serv.AccountsGetBakersListHandler = &getBakerListHandler{db}
	serv.AccountsGetPublicBakersListHandler = &getPublicBakerListHandler{db}
	serv.AccountsGetPublicBakersListForSearchHandler = &getPublicBakerSearchListHandler{db}
	serv.AccountsGetAccountDelegatorsHandler = &getAccountDelegatorsHandler{db}
	serv.AccountsGetContractsListHandler = &getContractListHandler{db}
	serv.AccountsGetAccountBakedBlocksListHandler = &getAccountBakedBlocksListHandler{db}
	serv.AccountsGetAccountTotalBakingHandler = &getAccountTotalBakingHandler{db}
	serv.AccountsGetAccountFutureBakingHandler = &getAccountFutureBakingHandler{db}
	serv.AccountsGetAccountFutureBakingRightsByCycleHandler = &getAccountFutureBakingRightsHandler{db}
	serv.AccountsGetAccountEndorsingListHandler = &getAccountEndorsingListHandler{db}
	serv.AccountsGetAccountTotalEndorsingHandler = &getAccountTotalEndorsingHandler{db}
	serv.AccountsGetAccountEndorsementsByCycleListHandler = &getAccountEndorsementsHandler{db}
	serv.AccountsGetAccountRewardsListHandler = &getAccountRewardsListHandler{db}
	serv.AccountsGetAccountDelegatorsByCycleListHandler = &getAccountDelegatorsByCycleListHandler{db}
	serv.AccountsGetAccountSecurityDepositListHandler = &getBakerSecurityDepositFutureListHandler{db}
	serv.BlocksGetBakingRightsHandler = &getBakingRightsHandler{db}
	serv.BlocksGetFutureBakingRightsHandler = &getFutureBakingRightsHandler{db}
	serv.AccountsGetAccountFutureEndorsingHandler = &getAccountFutureEndorsingHandler{db}
	serv.AccountsGetAccountFutureEndorsementRightsByCycleHandler = &getAccountFutureEndorsingRightsHandler{db}
	serv.GetSnapshotsHandler = &getSnapshotsHandler{db}
	serv.BlocksGetBlockBakingRightsHandler = &getBlockBakingRightsHandler{db}
	serv.OperationsListGetDoubleBakingsListHandler = &getDoubleBakingsListHandler{db}
	serv.OperationsListGetDoubleEndorsementsListHandler = &getDoubleEndorsementListHandler{db}
	serv.VotingGetPeriodHandler = &getPeriodInfoHandler{db}
	serv.VotingGetPeriodsListHandler = &getPeriodListHandler{db}
	serv.VotingGetProposalsByPeriodIDHandler = &getProposalListHandler{db}
	serv.VotingGetProposalVotesListHandler = &getProposalVotesHandler{db}
	serv.VotingGetProtocolsListHandler = &getProtocolListHandler{db}
	serv.VotingGetNonVotersByPeriodIDHandler = &getNonVotersHandler{db}
	serv.VotingGetBallotsByPeriodIDHandler = &getBallotsHandler{db}
	//	Assets
	serv.AssetsGetAssetTokenHoldersListHandler = &getAssetHoldersHandler{db}
	serv.AssetsGetAssetTokenInfoHandler = &getAssetInfoHandler{db}
	serv.AssetsGetAssetsListHandler = &getAssetsListHandler{db}
	serv.AssetsGetAssetOperationsListHandler = &getAssetOperationListHandler{db}
	//	Mempool
	serv.MempoolGetMempoolOperationsHandler = &getMempoolHandler{db}

	//	WS
	serv.WsConnectToWSHandler = &serveWS{provider: db}
}
