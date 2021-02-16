package pusher

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services"
	"github.com/everstake/teztracker/services/mailer"
	wsmodels "github.com/everstake/teztracker/ws/models"
	log "github.com/sirupsen/logrus"
)

type EmailPusher struct {
	mail    mailer.Mail
	service services.Provider
}

func NewEmailPusher(mail mailer.Mail, service services.Provider) *EmailPusher {
	return &EmailPusher{
		mail:    mail,
		service: service,
	}
}

func (p EmailPusher) Push(event wsmodels.EventType, data interface{}) (err error) {
	if p.mail == nil {
		return nil
	}
	switch event {
	case wsmodels.EventTypeOperation:
		err = p.sendOperation(data)
	case wsmodels.EventTypeAssetOperation:
		err = p.sendAssetOperation(data)
	default:
		return nil
	}
	if err != nil {
		return fmt.Errorf("send[%s]: %s", event, err.Error())
	}
	return nil
}

func (p EmailPusher) sendOperation(data interface{}) error {
	operation, ok := data.(models.Operation)
	if !ok {
		return fmt.Errorf("wrong data")
	}
	if !operation.Kind.Valid {
		return nil
	}
	var addresses []string
	var msgType string
	msgValues := make(map[string]string)
	if operation.Kind.String == "delegation" {
		addresses = append(addresses, operation.Delegate, operation.Source)
		msgValues["delegator"] = operation.Source
		msgValues["validator"] = operation.Delegate
	}
	if operation.Kind.String == "transaction" {
		addresses = append(addresses, operation.Source, operation.Destination)
		msgValues["from"] = operation.Source
		msgValues["to"] = operation.Destination
		msgValues["amount"] = fmt.Sprintf("%f", float64(operation.Amount/1e6))
		msgValues["operation"] = operation.OperationGroupHash.String
		msgValues["token"] = "xtz"
	}
	if len(addresses) == 0 {
		return nil
	}
	userProfileRepo := p.service.GetUserProfile()
	users, err := userProfileRepo.GetVerifiedUsersAndAddresses(addresses)
	if err != nil {
		return fmt.Errorf("userProfileRepo.GetVerifiedUsersAndAddresses: %s", err.Error())
	}
	for _, user := range users {
		if user.Email == "" {
			continue
		}
		switch operation.Kind.String {
		case "delegation":
			if !user.DelegationsEnabled {
				continue
			}
			if msgValues["delegator"] == user.Address {
				msgType = mailer.DelegatorDelegationMsg
			} else {
				msgType = mailer.ValidatorDelegationMsg
			}
		case "transaction":
			if !user.InTransfersEnabled && msgValues["to"] == user.Address {
				continue
			}
			if !user.OutTransfersEnabled && msgValues["from"] == user.Address {
				continue
			}
			msgType = mailer.InTransferMsg
			if msgValues["from"] == user.Address {
				msgType = mailer.OutTransferMsg
			}
		}
		err = p.mail.Send(user.Email, msgType, msgValues)
		if err != nil {
			log.Errorf("Watcher: mail: cant send to %s: %s", user.Email, err.Error())
		}
	}
	return nil
}

func (p EmailPusher) sendAssetOperation(data interface{}) error {
	operation, ok := data.(models.AssetOperation)
	if !ok {
		return fmt.Errorf("wrong data")
	}
	if operation.Type != "transfer" {
		return nil
	}
	addresses := []string{operation.Receiver, operation.Sender}
	userProfileRepo := p.service.GetUserProfile()
	users, err := userProfileRepo.GetVerifiedUsersAndAddresses(addresses)
	if err != nil {
		return fmt.Errorf("userProfileRepo.GetVerifiedUsersAndAddresses: %s", err.Error())
	}
	msgValues := make(map[string]string)
	msgValues["from"] = operation.Sender
	msgValues["to"] = operation.Receiver
	msgValues["amount"] = fmt.Sprintf("%f", float64(operation.Amount))
	msgValues["operation"] = operation.OperationGroupHash
	registeredToken, _ := p.service.GetAssets().GetRegisteredToken(operation.TokenId)
	msgValues["token"] = registeredToken.Name
	for _, user := range users {
		if user.Email == "" {
			continue
		}
		if !user.InTransfersEnabled && msgValues["to"] == user.Address {
			continue
		}
		if !user.OutTransfersEnabled && msgValues["from"] == user.Address {
			continue
		}
		msgType := mailer.InTransferMsg
		if msgValues["from"] == user.Address {
			msgType = mailer.OutTransferMsg
		}
		err = p.mail.Send(user.Email, msgType, msgValues)
		if err != nil {
			log.Errorf("Watcher: mail: cant send to %s: %s", user.Email, err.Error())
		}
	}
	return nil
}
