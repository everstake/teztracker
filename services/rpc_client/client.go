package rpc_client

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tzblock "github.com/bullblock-io/go-tezos/v2/block"
	tzc "github.com/bullblock-io/go-tezos/v2/client"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/rpc_client/client"
	"github.com/everstake/teztracker/services/rpc_client/client/baking_rights"
	"github.com/everstake/teztracker/services/rpc_client/client/snapshots"
	genmodels "github.com/everstake/teztracker/services/rpc_client/models"
)

const headBlock = "head"
const BlocksInCycle = 4096

type Tezos struct {
	client        *client.Tezosrpc
	network       string
	isTestNetwork bool //we have to use a separate flag due to stupid nodes configs...
	tzcClient     *tzblock.BlockService
}

func (t *Tezos) BlocksInCycle() int64 {
	if t.isTestNetwork {
		return BlocksInCycle / 2
	}
	return BlocksInCycle
}

func New(cfg client.TransportConfig, network string, isTestNetwork bool) *Tezos {
	cli := client.NewHTTPClientWithConfig(nil, &cfg)
	return &Tezos{
		client:        cli,
		network:       network,
		tzcClient:     tzblock.NewBlockService(tzc.NewClient(cfg.Host)),
		isTestNetwork: isTestNetwork,
	}
}
func (t *Tezos) RightsFor(ctx context.Context, blockFrom, blockTo, currentHead int64) ([]models.FutureBakingRight, error) {
	all := true
	blockToUse := headBlock
	if currentHead >= blockFrom {
		blockToUse = strconv.FormatInt(blockFrom, 10)
	}

	params := baking_rights.NewGetBakingRightsParamsWithContext(ctx).
		WithNetwork(t.network).
		WithBlock(blockToUse).
		WithAll(&all)

	levels := []string{}
	for b := blockFrom; b <= blockTo; b++ {
		levels = append(levels, strconv.FormatInt(b, 10))
	}
	params.SetLevel(levels)
	resp, err := t.client.BakingRights.GetBakingRights(params)
	if err != nil {
		return nil, err
	}
	rights := make([]models.FutureBakingRight, len(resp.Payload))
	for i := range resp.Payload {
		if resp.Payload[i] != nil {
			rights[i] = genRightToModel(*resp.Payload[i])
		}
	}
	return rights, nil
}

func genRightToModel(m genmodels.BakingRight) models.FutureBakingRight {
	return models.FutureBakingRight{
		Level:         m.Level,
		Priority:      int(m.Priority),
		Delegate:      m.Delegate,
		EstimatedTime: time.Time(m.EstimatedTime),
	}
}

func (t *Tezos) SnapshotForCycle(ctx context.Context, cycle int64, useHead bool) (snap models.Snapshot, err error) {
	blockToUse := headBlock
	if !useHead {
		level := cycle*t.BlocksInCycle() + 1
		blockToUse = strconv.FormatInt(level, 10)
	}
	params := snapshots.NewGetRollSnapshotParamsWithContext(ctx).
		WithCycle(cycle).
		WithNetwork(t.network).
		WithBlock(blockToUse)
	resp, err := t.client.Snapshots.GetRollSnapshot(params)
	if err != nil {
		return snap, err
	}
	snapshot := resp.Payload
	snap.Cycle = cycle
	snap.BlockLevel = ((cycle-7)*t.BlocksInCycle() + 1) + (snapshot+1)*256 - 1
	rollParams := snapshots.NewGetRollsParamsWithContext(ctx).
		WithCycle(cycle).
		WithNetwork(t.network).
		WithSnap(snapshot).
		WithBlock(blockToUse)
	rollsResp, err := t.client.Snapshots.GetRolls(rollParams)
	if err != nil {
		return snap, err
	}
	if rollsResp == nil {
		return snap, fmt.Errorf("nil resp")
	}

	snap.Rolls = int64(len(rollsResp.Payload))
	return snap, nil
}

func (t *Tezos) DoubleBakingEvidence(ctx context.Context, blockLevel int, operationHash string) (dee models.DoubleBakingEvidence, err error) {
	block, err := t.tzcClient.Get(blockLevel)
	if err != nil {
		return dee, err
	}
	for i := range block.Operations {
		for _, op := range block.Operations[i] {
			if strings.EqualFold(op.Hash, operationHash) {
				dee, err := ToDoubleBakingEvidence(op)
				if err != nil {
					return dee, err
				}
				dee.BlockLevel = int64(block.Header.Level)
				dee.BlockHash = block.Hash
				return dee, err
			}
		}
	}
	return dee, fmt.Errorf("not found")
}

func ToDoubleBakingEvidence(op tzblock.Operations) (dee models.DoubleBakingEvidence, err error) {
	for i := range op.Contents {
		if op.Contents[i].Bh1 != nil {
			dee.DenouncedLevel = int64(op.Contents[i].Bh1.Level)
			dee.Priority = op.Contents[i].Bh1.Priority
			if meta := op.Contents[i].Metadata; meta != nil {
				for _, bu := range meta.BalanceUpdates {
					if strings.EqualFold(bu.Kind, "freezer") {
						switch strings.ToLower(bu.Category) {
						case "deposits":
							dee.Offender = bu.Delegate
							change, err := strconv.ParseInt(bu.Change, 10, 64)
							if err != nil {
								return dee, err
							}
							dee.LostDeposits = -change
						case "rewards":
							change, err := strconv.ParseInt(bu.Change, 10, 64)
							if err != nil {
								return dee, err
							}
							if change < 0 {
								dee.LostRewards = -change
							} else {
								dee.BakerReward = change
								dee.EvidenceBaker = bu.Delegate
							}

						}
					}
				}
			}
			return dee, nil
		}
	}
	return dee, fmt.Errorf("not a double baking evidence")

}

func (t *Tezos) DoubleEndorsementEvidenceLevel(ctx context.Context, blockLevel int, operationHash string) (int64, error) {
	block, err := t.tzcClient.Get(blockLevel)
	if err != nil {
		return 0, err
	}
	for i := range block.Operations {
		for _, op := range block.Operations[i] {
			if strings.EqualFold(op.Hash, operationHash) {
				return GetDoubleEndorsementEvidenceLevel(op)
			}
		}
	}
	return 0, fmt.Errorf("not found")
}

func GetDoubleEndorsementEvidenceLevel(op tzblock.Operations) (int64, error) {
	for i := range op.Contents {
		if op.Contents[i].Op1 != nil {
			return int64(op.Contents[i].Op1.Operations.Level), nil
		}
	}
	return 0, fmt.Errorf("not a double endorsement evidence")
}
