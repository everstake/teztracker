package mailer

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strings"
)

const (
	from     = "teztracker@everstake.one"
	fromName = "TezTracker"
	subject  = "Notification from TezTracker"

	ValidatorDelegationMsg = "validator_delegation"
	DelegatorDelegationMsg = "delegator_delegation"
	TransferMsg            = "transfer"
)

var templatesPath = map[string]string{
	ValidatorDelegationMsg: "./services/mailer/templates/validator_delegation.txt",
	DelegatorDelegationMsg: "./services/mailer/templates/delegator_delegation.txt",
	TransferMsg:            "./services/mailer/templates/transfer.txt",
}

type Mailer struct {
	templates map[string]string
	dialer    *gomail.Dialer
}

type Mail interface {
	Send(email string, msgType string, values map[string]string) error
}

func New(host string, port int, user, password string) *Mailer {
	d := gomail.NewDialer(host, port, user, password)
	return &Mailer{
		templates: loadTemplates(),
		dialer:    d,
	}
}

func (m Mailer) makeContent(msgType string, values map[string]string) (string, error) {
	content, ok := m.templates[msgType]
	if !ok {
		return "", fmt.Errorf("can`t find template for `%s` type", msgType)
	}
	for key, value := range values {
		item := fmt.Sprintf("{%s}", key)
		content = strings.ReplaceAll(content, item, value)
	}
	return content, nil
}

func (m Mailer) Send(email string, msgType string, values map[string]string) error {
	content, err := m.makeContent(msgType, values)
	if err != nil {
		return fmt.Errorf("makeContent: %s", err.Error())
	}
	err = m.send(email, content)
	if err != nil {
		return fmt.Errorf("send: %s", err.Error())
	}
	return nil
}

func (m Mailer) send(email string, data string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("Subject", subject)
	msg.SetAddressHeader("From", from, fromName)
	msg.SetHeader("To", email)
	msg.SetBody("text/plain", data)
	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("dialer.DialAndSend: %s", err.Error())
	}
	return nil
}

func loadTemplates() (mp map[string]string) {
	mp = make(map[string]string)
	for msg, path := range templatesPath {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Error(fmt.Sprintf("Mailer: loadTemplates: can`t read file(%s): %s", path, err.Error()))
			continue
		}
		mp[msg] = string(content)
	}

	return mp
}
