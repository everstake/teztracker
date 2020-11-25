package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
	wsmodels "github.com/everstake/teztracker/ws/models"
)

type Block struct {
	repos services.Provider
}

func NewBlockTask(repos services.Provider) Block {

	return Block{repos: repos}
}

func (b Block) GetEventData(data interface{}) ([]wsmodels.EventType, interface{}, error) {

	bt, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	block := models.Block{}

	err = json.Unmarshal(bt, &block)
	if err != nil {
		return nil, nil, err
	}

	//Extend block
	found, extBlock, err := b.repos.GetBlock().FindExtended(block)
	if err != nil {
		return nil, nil, err
	}

	if !found {
		return nil, nil, fmt.Errorf("Block %d not found", block.Level)
	}

	//TODO move render as smodels render
	apiBlock := render.Block(extBlock)

	return []wsmodels.EventType{wsmodels.EventTypeBlock}, apiBlock, nil
}
