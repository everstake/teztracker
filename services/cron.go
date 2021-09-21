package services

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/services/assets"
	"github.com/everstake/teztracker/services/counter"
	"github.com/everstake/teztracker/services/double_baking"
	"github.com/everstake/teztracker/services/double_endorsement"
	"github.com/everstake/teztracker/services/ipfs"
	"github.com/everstake/teztracker/services/nft"
	"github.com/everstake/teztracker/services/public_baker"
	"github.com/everstake/teztracker/services/rolls"
	"github.com/everstake/teztracker/services/snapshots"
	"github.com/everstake/teztracker/services/thirdparty_bakers"

	"github.com/everstake/teztracker/services/future_rights"

	"github.com/everstake/teztracker/services/cmc"
	"github.com/everstake/teztracker/services/mailer"
	"github.com/everstake/teztracker/ws"

	"github.com/everstake/teztracker/config"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services/rpc_client"
	"github.com/everstake/teztracker/services/rpc_client/client"
	wsmodels "github.com/everstake/teztracker/ws/models"
	"github.com/jinzhu/gorm"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

func AddToCron(cron *gron.Cron, cfg config.Config, db *gorm.DB, ws *ws.Hub, mail mailer.Mail, marketDataProvider *cmc.CoinGecko, rpcConfig client.TransportConfig, network models.Network, isTestNetwork bool) {

	if cfg.CounterIntervalSeconds > 0 {
		dur := time.Duration(cfg.CounterIntervalSeconds) * time.Second
		log.Infof("Sheduling counter saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			unitOfWork := repos.New(db)

			err := counter.SaveCounters(unitOfWork.GetOperation(), unitOfWork.GetOperationCounter())
			if err != nil {
				log.Errorf("counter saver failed: %s", err.Error())
				return
			}

		})
	} else {
		log.Infof("no sheduling counter due to missing CounterIntervalHours in config")
	}
	if cfg.SnapshotCheckIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.SnapshotCheckIntervalMinutes) * time.Minute
		log.Infof("Sheduling snapshots parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				count, err := snapshots.SaveNewSnapshots(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("Snapshots saver failed: %s", err.Error())
					return
				}
				log.Tracef("Snapshots saved %d rights", count)
			} else {
				log.Tracef("skipping Snapshots saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling snapshots parser due to missing FutureRightsIntervalMinutes in config")
	}

	if cfg.FutureRightsIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.FutureRightsIntervalMinutes) * time.Minute
		log.Infof("Sheduling future rights parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				count, err := future_rights.SaveNewBakingRights(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("BakingRights saver failed: %s", err.Error())
					return
				}
				log.Tracef("BakingRights saved %d rights", count)
			} else {
				log.Tracef("skipping BakingRights saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling future rights parser due to missing FutureRightsIntervalMinutes in config")
	}

	if cfg.FutureRightsIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.FutureRightsIntervalMinutes) * time.Minute

		log.Infof("Sheduling future rights parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				count, err := future_rights.SaveNewEndorsementRights(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("BakingRights saver failed: %s", err.Error())
					return
				}
				log.Tracef("EndorsementRights saved %d rights", count)
			} else {
				log.Tracef("skipping EndorsementRights saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling future rights parser due to missing FutureRightsIntervalMinutes in config")
	}

	if cfg.VotingRollsIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.VotingRollsIntervalMinutes) * time.Minute
		log.Infof("Sheduling rolls parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				count, err := rolls.SaveRolls(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("Rolls saver failed: %s", err.Error())
					return
				}
				log.Tracef("Rolls saved %d count", count)
			} else {
				log.Tracef("skipping Rolls saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling rolls parser due to missing in config")
	}

	if cfg.DoubleBakingCheckIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.DoubleBakingCheckIntervalMinutes) * time.Minute
		log.Infof("Sheduling double baking parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				err := double_baking.SaveUnprocessedDoubleBakingEvidences(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("double baking saver failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping double baking saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling double baking parser due to missing DoubleBakingCheckIntervalMinutes in config")
	}
	if cfg.DoubleEndorsementCheckIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.DoubleEndorsementCheckIntervalMinutes) * time.Minute
		log.Infof("Sheduling double endorsement parser saver every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				err := double_endorsement.SaveUnprocessedDoubleEndorsementEvidences(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("double endorsement saver failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping double endorsement saver as the previous job is still running")
			}
		})
	} else {
		log.Infof("no sheduling double endorsement parser due to missing DoubleEndorsementCheckIntervalMinutes in config")
	}

	func() {
		var jobIsRunning uint32

		//Todo refactor
		dur := 1 * time.Minute
		log.Infof("Sheduling baker materialized view update every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)
				err := unitOfWork.GetBaker().RefreshView()
				if err != nil {
					log.Errorf("materialized view update failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping materialized view update as the previous job is still running")
			}
		})
	}()

	func() {
		var jobIsRunning uint32

		dur := 1 * time.Hour
		log.Infof("Sheduling insert whale accounts every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)
				unitOfWork := repos.New(db)

				year, months, day := time.Now().Date()
				startOfDay := time.Date(year, months, day, 0, 0, 0, 0, time.Local)

				err := unitOfWork.GetChart().InsertWhaleAccounts(startOfDay.Unix())
				if err != nil {
					log.Errorf("insert whale accounts failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping materialized view update as the previous job is still running")
			}
		})
	}()

	if cfg.BakerRegistryCheckIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.BakerRegistryCheckIntervalMinutes) * time.Minute
		log.Infof("Sheduling materialized view update every %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)
				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				err := public_baker.MonitorPublicBakers(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("public bakers update failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping public bakers update as the previous job is still running")
			}
		})
	}

	if cfg.AssetsParseIntervalMinutes > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.AssetsParseIntervalMinutes) * time.Minute
		log.Infof("Sheduling parse assets operations %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)
				rpc := rpc_client.New(rpcConfig, string(network), isTestNetwork)
				err := assets.ProcessAssetOperations(context.TODO(), unitOfWork, rpc)
				if err != nil {
					log.Errorf("assets operations parse failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping assets operations parse as the previous job is still running")
			}
		})
	}

	if cfg.NFTTokensParseIntervalSeconds > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.NFTTokensParseIntervalSeconds) * time.Second

		log.Infof("Sheduling parse nft tokens %s", dur)

		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)
				ipfsClient, err := ipfs.NewIPFSClient(cfg.IPFSClient)
				if err != nil {
					log.Fatalf("Wrong IPFS client url: %s")
				}

				err = nft.ProcessNFTMintOperations(context.TODO(), unitOfWork, ipfsClient)
				if err != nil {
					log.Errorf("nft tokens failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping nft tokens parse as the previous job is still running")
			}
		})

	}

	if cfg.NFTTokensParseIntervalSeconds > 0 {
		var jobIsRunning uint32

		dur := time.Duration(cfg.NFTTokensParseIntervalSeconds) * time.Second

		log.Infof("Sheduling update nft tokens %s", dur)

		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)

				err := nft.ProcessNFTOperations(context.TODO(), unitOfWork)
				if err != nil {
					log.Errorf("nft tokens update failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping nft tokens update as the previous job is still running")
			}
		})

	}

	if !isTestNetwork { // dispatch user verifications every minutes
		var jobIsRunning uint32

		log.Infof("Sheduling check email verifications")
		cron.AddFunc(gron.Every(time.Minute), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				service := New(repos.New(db), models.NetworkMain)
				err := service.SendNewVerifications(mail)
				if err != nil {
					log.Errorf("dispatch mail verification failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("skipping check email verifications as the previous job is still running")
			}
		})
	}

	if cfg.ThirdPartyBakersIntervalMinutes > 0 && !isTestNetwork {
		var jobIsRunning uint32

		dur := time.Duration(cfg.ThirdPartyBakersIntervalMinutes) * time.Minute
		log.Infof("Sheduling update third party bakers %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)
				err := thirdparty_bakers.UpdateBakers(context.TODO(), unitOfWork)
				if err != nil {
					log.Errorf("third party bakers failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("updating third party bakers as the previous job is still running")
			}
		})
	}

	if cfg.BakersSocialMediaHours > 0 && !isTestNetwork {
		var jobIsRunning uint32

		dur := time.Duration(cfg.BakersSocialMediaHours) * time.Hour
		log.Infof("Sheduling update bakers social media %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				unitOfWork := repos.New(db)
				err := New(unitOfWork, network).UpdateBakersSocialMedia()
				if err != nil {
					log.Errorf("UpdateBakersSocialMedia failed: %s", err.Error())
					return
				}
			} else {
				log.Tracef("updating bakers social media as the previous job is still running")
			}
		})
	}

	//Info cron
	func() {
		var jobIsRunning uint32

		dur := 1 * time.Minute
		log.Infof("Sheduling info ws publish %s", dur)
		cron.AddFunc(gron.Every(dur), func() {
			// Ensure jobs are not stacking up. If the previous job is still running - skip this run.
			if atomic.CompareAndSwapUint32(&jobIsRunning, 0, 1) {
				defer atomic.StoreUint32(&jobIsRunning, 0)

				//Outside network always main for RPC
				serviceNetwork := network
				if isTestNetwork {
					serviceNetwork = models.NetworkFlorence
				}

				service := New(repos.New(db), serviceNetwork)
				ratio, err := service.GetStakingRatio()
				if err != nil {
					log.Errorf("failed to get staking ratio: %s", err.Error())
				}

				for curr := range cmc.AvailableCurrencies {
					md, err := marketDataProvider.GetTezosMarketData(curr)
					if err != nil {
						log.Errorf("GetTezosMarketData failed: %s", err.Error())
						continue
					}

					apiInfo := render.Info(curr, md, ratio, service.BlocksInCycle())
					err = ws.Broadcast(wsmodels.BasicMessage{
						Event: wsmodels.EventType(fmt.Sprintf("%s_%s", string(wsmodels.EventTypeInfo), curr)),
						Data:  apiInfo,
					})
					if err != nil {
						log.Errorf("Broadcast info message failed: %s", err.Error())
						continue
					}
				}
			} else {
				log.Tracef("skipping info job is still running")
			}
		})
	}()

}
