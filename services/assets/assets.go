package assets

import (
	"blockwatch.cc/tzindex/chain"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/assets"
	"github.com/everstake/teztracker/repos/operation"
	"github.com/everstake/teztracker/services/michelson"
	"github.com/everstake/teztracker/services/rpc_client"
	"strconv"
	"strings"
)

type AssetRepo interface {
	GetUnprocessedAssetTxs(tokenID string) ([]models.Operation, error)
}

type AssetProvider interface {
	Operation(ctx context.Context, blockHash, transactionHash string) (op rpc_client.Operation, err error)
	Script(ctx context.Context, contractHash string) (bm michelson.BigMap, err error)
}

type UnitOfWork interface {
	GetAssets() assets.Repo
	GetOperation() operation.Repo
}

const limit = 100

type Transfer struct {
	Transfer interface{} `json:"transfer"`
	From     string      `json:"from"`
	To       string      `json:"to"`
	Value    string      `json:"value"`

	//	Mint
	Mint []string `json:"mint"`
}

func ProcessAssetOperations(ctx context.Context, unit UnitOfWork, provider AssetProvider) (err error) {

	repo := unit.GetAssets()

	_, tokens, err := repo.GetTokensList()
	if err != nil {
		return err
	}

	for tI := range tokens {

		ops, err := repo.GetUnprocessedAssetTxs(tokens[tI].AccountId)
		if err != nil {
			return err
		}

		if len(ops) == 0 {
			continue
		}

		script, err := provider.Script(ctx, tokens[tI].AccountId)
		if err != nil {
			return err
		}

		container := michelson.NewBigMapContainer()

		container.InitPath(script.Code.Args[0].Args[0])

		var groupHashOperations []models.Operation
		for opsI := range ops {
			container.FlushValues()

			groupHashOperations, err = unit.GetOperation().List([]string{ops[opsI].OperationGroupHash.String}, []string{"transaction"}, nil, nil, 10, 0, 0, nil)
			if err != nil {
				return err
			}

			if len(groupHashOperations) == 0 {
				return fmt.Errorf("Some error")
			}

			op, err := provider.Operation(ctx, groupHashOperations[0].BlockHash.String, groupHashOperations[0].OperationGroupHash.String)
			if err != nil {
				return err
			}

			if len(op.Contents) == 0 {
				return fmt.Errorf("Some error Contents")
			}

			var value *rpc_client.Contents
			//With reveal
			for i := range op.Contents {
				if op.Contents[i].Kind != "transaction" {
					continue
				}

				value = op.Contents[i]
			}

			var transfer Transfer

			switch len(groupHashOperations) {
			//Direct call
			case 1:
				container.ParseValues(value.Parameters.Entrypoint, value.Parameters.Value)
			//Internal SC call
			case 2:
				if len(value.Metadata.InternalOperationResults) > 0 {
					container.ParseValues(value.Metadata.InternalOperationResults[0].Parameters.Entrypoint, value.Metadata.InternalOperationResults[0].Parameters.Value)
				} else {
					container.ParseValues(value.Parameters.Entrypoint, value.Parameters.Value)
				}

			}

			c, err := container.MarshalJSON()
			if err != nil {
				return err
			}

			err = json.Unmarshal(c, &transfer)
			if err != nil {
				return err
			}

			var opType, to string
			var amount int64

			switch {
			//Transfer call
			case transfer.Transfer != nil:

				opType = "transfer"
				switch transfer.Transfer.(type) {
				case []interface{}:
					fields := transfer.Transfer.([]interface{})

					if len(fields) != 2 {
						continue
					}

					split := strings.Split(fields[1].(string), "#")

					to = split[0]
					if !chain.HasAddressPrefix(split[0]) {
						to, err = convertAddress(split[0])
						if err != nil {
							return err
						}
					}

					amount, err = strconv.ParseInt(split[1], 10, 64)
					if err != nil {
						return err
					}

				case string:

					to = transfer.To
					if !chain.HasAddressPrefix(transfer.To) {
						to, err = convertAddress(transfer.To)
						if err != nil {
							return err
						}
					}

					amount, err = strconv.ParseInt(transfer.Value, 10, 64)
					if err != nil {
						return err
					}
				}
			//Mint call
			case transfer.Mint != nil:
				//from = groupHashOperations[0].Source
				if len(transfer.Mint) != 2 {
					continue
				}

				to = transfer.Mint[0]
				amount, err = strconv.ParseInt(transfer.Mint[1], 10, 64)
				if err != nil {
					return err
				}
				opType = "mint"
			//	Some other call
			default:
				opType = value.Parameters.Entrypoint
				if opType == "default" {
					call := map[string]interface{}{}

					err = json.Unmarshal(c, &call)
					if err != nil {
						return err
					}

					for key := range call {
						opType = key
					}
				}

				to = tokens[tI].AccountId
			}

			err = repo.CreateAssetOperations(models.AssetOperation{
				TokenId:            tokens[tI].ID,
				OperationId:        groupHashOperations[0].OperationID.Int64,
				OperationGroupHash: op.Hash,
				Sender:             value.Source,
				Receiver:           to,
				Amount:             amount,
				Type:               opType,
				Timestamp:          groupHashOperations[0].Timestamp,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func convertAddress(hexAddr string) (string, error) {
	bt, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", err
	}

	ad := chain.Address{}
	err = ad.UnmarshalBinary(bt)
	if err != nil {
		return "", err
	}

	return ad.String(), nil
}
