package services

import (
	"context"
	"github.com/everstake/teztracker/services/public_baker"
	"sync/atomic"
	"time"

	"github.com/everstake/teztracker/config"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services/counter"
	"github.com/everstake/teztracker/services/double_baking"
	"github.com/everstake/teztracker/services/double_endorsement"
	"github.com/everstake/teztracker/services/future_rights"
	"github.com/everstake/teztracker/services/rpc_client"
	"github.com/everstake/teztracker/services/rpc_client/client"
	"github.com/everstake/teztracker/services/snapshots"
	"github.com/jinzhu/gorm"
	"github.com/roylee0704/gron"
	log "github.com/sirupsen/logrus"
)

func AddToCron(cron *gron.Cron, cfg config.Config, db *gorm.DB, rpcConfig client.TransportConfig, network models.Network, isTestNetwork bool) {

	if cfg.CounterIntervalHours > 0 {
		dur := time.Duration(cfg.CounterIntervalHours) * time.Hour
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
					log.Errorf("EndorsementRights saver failed: %s", err.Error())
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

		dur := 30 * time.Second
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
}
