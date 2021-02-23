package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
)

type Operation struct {
	repos services.Provider
}

func NewOperationTask(repos services.Provider) Operation {

	return Operation{repos: repos}
}

func (o Operation) GetEventData(data interface{}) (interface{}, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	operation := models.Operation{}

	err = json.Unmarshal(bt, &operation)
	if err != nil {
		return nil, err
	}

	op, err := o.repos.GetOperation().List(nil, nil, nil, nil, 1, 0, 0, []int64{operation.OperationID.Int64})
	if err != nil {
		return nil, err
	}

	if len(op) != 1 {
		return nil, fmt.Errorf("Wrong resp len: %d", len(op))
	}

	return op[0], nil
}
