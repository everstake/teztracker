package pusher

import (
	"fmt"
	"github.com/everstake/teztracker/api/render"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/ws"
	wsmodels "github.com/everstake/teztracker/ws/models"
	"strings"
)

type WSPusher struct {
	ws *ws.Hub
}

func NewWSPusher(ws *ws.Hub) *WSPusher {
	return &WSPusher{
		ws: ws,
	}
}

func (p WSPusher) Push(event wsmodels.EventType, data interface{}) (err error) {
	switch event {
	case wsmodels.EventTypeBlock:
		err = p.sendBlock(data)
	case wsmodels.EventTypeOperation:
		err = p.sendOperation(data)
	case wsmodels.EventTypeAccountCreated:
		err = p.sendAccount(data)
	default:
		return nil
	}
	if err != nil {
		return fmt.Errorf("send[%s]: %s", event, err.Error())
	}
	return nil
}

func (p WSPusher) sendBlock(data interface{}) error {
	block, ok := data.(models.Block)
	if !ok {
		return fmt.Errorf("wrong data")
	}
	apiBlock := render.Block(block)
	p.ws.Broadcast(wsmodels.BasicMessage{Event: wsmodels.EventTypeBlock, Data: apiBlock})
	return nil
}

func (p WSPusher) sendOperation(data interface{}) error {
	op, ok := data.(models.Operation)
	if !ok {
		return fmt.Errorf("wrong data")
	}
	apiOperation := render.Operation(op, nil)
	p.ws.Broadcast(wsmodels.BasicMessage{Event: wsmodels.EventTypeOperation, Data: apiOperation})
	p.ws.Broadcast(wsmodels.BasicMessage{Event: wsmodels.EventType(fmt.Sprint(op.Kind.String, "s")), Data: apiOperation})
	return nil
}

func (p WSPusher) sendAccount(data interface{}) error {
	acc, ok := data.(models.AccountListView)
	if !ok {
		return fmt.Errorf("wrong data")
	}
	apiAccount := render.Account(acc)
	accountType := "accounts"
	if strings.Contains(acc.AccountID.String, models.ContractAccountPrefix) {
		accountType = "contracts"
	}
	p.ws.Broadcast(wsmodels.BasicMessage{Event: wsmodels.EventTypeAccountCreated, Data: apiAccount})
	p.ws.Broadcast(wsmodels.BasicMessage{Event: wsmodels.EventType(accountType), Data: apiAccount})
	return nil
}
