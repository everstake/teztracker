package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
	wsmodels "github.com/everstake/teztracker/ws/models"
)

type Operation struct {
	repos services.Provider
}

func NewOperationTask(repos services.Provider) Operation {

	return Operation{repos: repos}
}

func (o Operation) GetEventData(data interface{}) ([]wsmodels.EventType, interface{}, error) {
	bt, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	operation := models.Operation{}

	err = json.Unmarshal(bt, &operation)
	if err != nil {
		return nil, nil, err
	}

	op, err := o.repos.GetOperation().List(nil, nil, nil, nil, 1, 0, 0, []int64{operation.OperationID.Int64})
	if err != nil {
		return nil, nil, err
	}

	if len(op) != 1 {
		return nil, nil, fmt.Errorf("Wrong resp len: %d", len(op))
	}

	apiOperation := render.Operation(op[0], nil)

	return []wsmodels.EventType{wsmodels.EventTypeOperation, wsmodels.EventType(fmt.Sprint(op[0].Kind.String, "s"))}, apiOperation, nil
}
