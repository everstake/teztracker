package api

import (
	"github.com/everstake/teztracker/gen/restapi/operations"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/cmc"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type DbProvider interface {
	GetDb(models.Network) (*gorm.DB, error)
}

// SetHandlers initializes the API handlers with underlying services.
func SetHandlers(serv *operations.TezTrackerAPI, db DbProvider) {
	serv.Logger = logrus.Infof
	serv.BlocksGetBlocksHeadHandler = &getHeadBlockHandler{db}
	serv.BlocksGetBlocksListHandler = &getBlockListHandler{db}
	serv.BlocksGetBlockHandler = &getBlockHandler{db}
	serv.OperationsListGetOperationsListHandler = &getOperationListHandler{db}
	serv.AppInfoGetInfoHandler = &getInfoHandler{cmc.NewCoinGecko(), db}
	serv.AccountsGetAccountsListHandler = &getAccountListHandler{db}
	serv.AccountsGetAccountHandler = &getAccountHandler{db}
	serv.BlocksGetBlockEndorsementsHandler = &getBlockEndorsementsHandler{db}
	serv.AccountsGetBakersListHandler = &getBakerListHandler{db}
	serv.AccountsGetAccountDelegatorsHandler = &getAccountDelegatorsHandler{db}
	serv.AccountsGetContractsListHandler = &getContractListHandler{db}
	serv.BlocksGetBakingRightsHandler = &getBakingRightsHandler{db}
	serv.BlocksGetFutureBakingRightsHandler = &getFutureBakingRightsHandler{db}
	serv.GetSnapshotsHandler = &getSnapshotsHandler{db}
	serv.BlocksGetBlockBakingRightsHandler = &getBlockBakingRightsHandler{db}
	serv.OperationsListGetDoubleBakingsListHandler = &getDoubleBakingsListHandler{db}
	serv.VotingGetPeriodHandler = &getPeriodInfoHandler{db}
	serv.VotingGetPeriodsListHandler = &getPeriodListHandler{db}
	serv.VotingGetProposalsByPeriodIDHandler = &getProposalListHandler{db}
	serv.VotingGetProposalVotesListHandler = &getProposalVotesHandler{db}
	serv.VotingGetNonVotersByPeriodIDHandler = &getNonVotersHandler{db}
	serv.VotingGetBallotsByPeriodIDHandler = &getBallotsHandler{db}
}
