package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
)

type Block struct {
	repos services.Provider
}

func NewBlockTask(repos services.Provider) Block {

	return Block{repos: repos}
}

func (b Block) GetEventData(data interface{}) ( interface{}, error) {

	bt, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	block := models.Block{}

	err = json.Unmarshal(bt, &block)
	if err != nil {
		return nil, err
	}

	//Extend block
	found, extBlock, err := b.repos.GetBlock().FindExtended(block)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("Block %d not found", block.Level)
	}

	return extBlock, nil
}
