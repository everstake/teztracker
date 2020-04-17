package rpc_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/services/michelson"
	"github.com/everstake/teztracker/services/rpc_client/client/contracts"
	"github.com/everstake/teztracker/services/rpc_client/client/operations"
	"strconv"
	"strings"
	"time"

	script "blockwatch.cc/tzindex/micheline"
	tzblock "github.com/bullblock-io/go-tezos/v2/block"
	tzc "github.com/bullblock-io/go-tezos/v2/client"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/rpc_client/client"
	"github.com/everstake/teztracker/services/rpc_client/client/baking_rights"
	"github.com/everstake/teztracker/services/rpc_client/client/endorsing_rights"
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

	//If scheme provided in config add it for library
	url := cfg.Host
	if len(cfg.Schemes) != 0 {
		url = fmt.Sprintf("%s://%s", cfg.Schemes[0], cfg.Host)
	}

	return &Tezos{
		client:        cli,
		network:       network,
		tzcClient:     tzblock.NewBlockService(tzc.NewClient(url)),
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

func (t *Tezos) EndorsementRightsFor(ctx context.Context, blockFrom, blockTo, currentHead int64) ([]models.FutureEndorsementRight, error) {
	blockToUse := headBlock
	if currentHead >= blockFrom {
		blockToUse = strconv.FormatInt(blockFrom, 10)
	}

	params := endorsing_rights.NewGetEndorsingRightsParamsWithContext(ctx).
		WithNetwork(t.network).
		WithBlock(blockToUse)

	levels := []string{}
	for b := blockFrom; b <= blockTo; b++ {
		levels = append(levels, strconv.FormatInt(b, 10))
	}
	params.SetLevel(levels)
	resp, err := t.client.EndorsingRights.GetEndorsingRights(params)
	if err != nil {
		return nil, err
	}
	rights := make([]models.FutureEndorsementRight, len(resp.Payload))
	for i := range resp.Payload {
		if resp.Payload[i] != nil {
			rights[i] = genEndorsementRightToModel(*resp.Payload[i])
		}
	}
	return rights, nil
}

func genEndorsementRightToModel(m genmodels.EndorsementRight) models.FutureEndorsementRight {
	return models.FutureEndorsementRight{
		Level:         m.Level,
		Slots:         m.Slots,
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

func ToDoubleOperationEvidence(op tzblock.Operations) (dee models.DoubleOperationEvidence, err error) {
	for i := range op.Contents {
		if op.Contents[i].Op1 != nil {
			dee.DenouncedLevel = int64(op.Contents[i].Op1.Operations.Level)
			err = parseDoubleOperationMetaData(&dee, op.Contents[i].Metadata)
			if err != nil {
				return dee, err
			}
			return dee, nil
		}
	}
	return dee, fmt.Errorf("not a double endorsement evidence")
}

func parseDoubleOperationMetaData(dee *models.DoubleOperationEvidence, meta *tzblock.ContentsMetadata) (err error) {
	if dee == nil {
		return fmt.Errorf("Empty double operation model")
	}

	if meta == nil {
		return nil
	}

	for _, bu := range meta.BalanceUpdates {
		if strings.EqualFold(bu.Kind, "freezer") {
			switch strings.ToLower(bu.Category) {
			case "deposits":
				dee.Offender = bu.Delegate
				change, err := strconv.ParseInt(bu.Change, 10, 64)
				if err != nil {
					return err
				}
				dee.LostDeposits = -change
			case "rewards":
				change, err := strconv.ParseInt(bu.Change, 10, 64)
				if err != nil {
					return err
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

	return nil
}

func (t *Tezos) DoubleOperationEvidence(ctx context.Context, blockLevel int, operationHash string) (dee models.DoubleOperationEvidence, err error) {
	block, err := t.tzcClient.Get(blockLevel)
	if err != nil {
		return dee, err
	}
	for i := range block.Operations {
		for _, op := range block.Operations[i] {
			if strings.EqualFold(op.Hash, operationHash) {
				dee, err := ToDoubleOperationEvidence(op)
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

//Todo move models somewhere
//Block parse
type Parameters struct {
	Entrypoint string       `json:"entrypoint"`
	Value      *script.Prim `json:"value"`
}

type Contents struct {
	Parameters Parameters `json:"parameters`
	Kind       string     `json:"kind"`
}

type Operation struct {
	Hash     string      `json:"hash"`
	Contents []*Contents `json:"contents"`
}

var sc [][]Operation

func (t *Tezos) Operation(ctx context.Context, blockHash, transactionHash string) (op Operation, err error) {

	params := operations.NewGetBlockOperationsParamsWithContext(ctx).WithBlock(blockHash)
	operations, err := t.client.Operations.GetBlockOperations(params)
	if err != nil {
		return op, err
	}

	for _, cont := range operations.Payload {
		for _, genOp := range cont {
			bt, err := json.Marshal(genOp)
			if err != nil {
				return op, err
			}

			err = json.Unmarshal(bt, &op)
			if err != nil {
				return op, err
			}

			if op.Hash == transactionHash {
				return op, nil
			}

		}
	}
	return op, fmt.Errorf("Operation not found")
}

func (t *Tezos) Script(ctx context.Context, contractHash string) (bm michelson.BigMap, err error) {
	params := contracts.NewGetContractScriptParamsWithContext(ctx).WithContract(contractHash)
	resp, err := t.client.Contracts.GetContractScript(params)
	if err != nil {
		return bm, err
	}

	bytes, err := json.Marshal(resp.Payload)
	if err != nil {
		return bm, err
	}

	err = json.Unmarshal(bytes, &bm)
	if err != nil {
		return bm, err
	}

	return bm, nil
}
