package public_baker

import (
	script "blockwatch.cc/tzindex/micheline"
	"context"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/baker"
	"github.com/everstake/teztracker/repos/operation"
	"github.com/everstake/teztracker/services/michelson"
	"github.com/everstake/teztracker/services/rpc_client"
)

type BakesProvider interface {
	Operation(ctx context.Context, blockHash, transactionHash string) (op rpc_client.Operation, err error)
	Script(ctx context.Context, contractHash string) (bm michelson.BigMap, err error)
}
type UnitOfWork interface {
	GetBaker() baker.Repo
	GetOperation() operation.Repo
}

const (
	BakerRegistryContract    = "KT1ChNsEFxwyCbJyWGSL3KdjeXE28AY1Kaog"
	operationKindTransaction = "transaction"
)

func MonitorPublicBakers(ctx context.Context, unit UnitOfWork, rpc BakesProvider) (err error) {

	bakerRepo := unit.GetBaker()
	publicBakers, err := bakerRepo.PublicBakersList()
	if err != nil {
		return err
	}

	operationRepo := unit.GetOperation()
	operations, err := operationRepo.List(nil, []string{operationKindTransaction}, nil, []string{BakerRegistryContract}, 0, 0, 0)
	if err != nil {
		return err
	}

	operationsByBaker := map[string]models.Operation{}

	//Group operations by source
	for key := range operations {
		if elem, ok := operationsByBaker[operations[key].Source]; ok {
			if elem.Level < operationsByBaker[operations[key].Source].Level {
				elem = operationsByBaker[operations[key].Source]
			}
			continue
		}
		operationsByBaker[operations[key].Source] = operations[key]
	}

	bakersMap := map[string]models.PublicBaker{}

	//Group bakers by pkh
	for key := range publicBakers {
		bakersMap[publicBakers[key].Delegate] = publicBakers[key]
	}

	contractContainer, err := InitContractScript(ctx, rpc, BakerRegistryContract)
	if err != nil {
		return err
	}

	for key, value := range operationsByBaker {
		//Baker have actual data
		if value.Level <= bakersMap[key].LastUpdateId {
			continue
		}

		publicBaker, isStorageOperation, err := GetPublicBakerInfo(ctx, rpc, contractContainer, value.BlockHash.String, value.OperationGroupHash.String)
		if err != nil {
			return err
		}

		//Not process storage operations
		if isStorageOperation {
			continue
		}

		publicBaker.LastUpdateId = value.OperationID.Int64

		err = bakerRepo.SavePublicBaker(publicBaker)
		if err != nil {
			return err
		}
	}

	return nil
}

func InitContractScript(ctx context.Context, rpc BakesProvider, contractHash string) (container michelson.BigMapContainer, err error) {
	contractScript, err := rpc.Script(ctx, contractHash)
	if err != nil {
		return container, err
	}

	//Insert params locate on L branch
	//R branch contains params for storage
	container.InitPath(contractScript.Code.Args[0].Args[0])

	return container, nil
}

func GetPublicBakerInfo(ctx context.Context, rpc BakesProvider, container michelson.BigMapContainer, blockHash, operationHash string) (publicBaker models.PublicBaker, isStorageOperation bool, err error) {
	op, err := rpc.Operation(ctx, blockHash, operationHash)
	if err != nil {
		return publicBaker, false, err
	}

	//Find required operation
	var index int
	for i := range op.Contents {
		if op.Contents[i].Kind != operationKindTransaction || op.Contents[i].Parameters.Value == nil {
			continue
		}
		index = i
	}

	v := op.Contents[index].Parameters.Value
	//Check unusual operations
	if op.Contents[index].Parameters.Value.OpCode != script.D_LEFT {
		//Not process updates of internal contract storage
		if op.Contents[index].Parameters.Value.OpCode == script.D_RIGHT {
			return publicBaker, true, nil
		}

		//For txs where missed highest level of tree
		v = &script.Prim{
			OpCode: script.D_LEFT,
			Args:   []*script.Prim{op.Contents[index].Parameters.Value},
		}
	}

	container.ParseValues(v)

	bt, err := container.MarshalJSON()
	if err != nil {
		return publicBaker, false, err
	}

	err = publicBaker.Unmarshal(bt)
	if err != nil {
		return publicBaker, false, err
	}

	return publicBaker, false, nil
}
