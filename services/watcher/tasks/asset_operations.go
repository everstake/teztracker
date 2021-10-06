package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
)

type AssetOperation struct {
	repos services.Provider
}

func NewAssetOperation(repos services.Provider) AssetOperation {

	return AssetOperation{repos: repos}
}

func (ao AssetOperation) GetEventData(data interface{}) (interface{}, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	op := models.AssetOperation{}

	err = json.Unmarshal(bt, &op)
	if err != nil {
		return nil, err
	}

	operations, err := ao.repos.GetAssets().FindOperations([]int64{op.OperationId}, 1)
	if err != nil {
		return nil, err
	}
	if len(operations) == 0 {
		return nil, fmt.Errorf("not found")
	}

	return operations[0], nil
}
