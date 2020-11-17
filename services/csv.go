package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/everstake/teztracker/models"
	"log"
)

const limit = 257000
const frontHost = "https://teztracker.everstake.one/en/mainnet/tx/%s"

var header = []string{"block level", "timestamp", "operation", "coin", "in", "out", "from", "to", "fee", "reward", "loss", "status", "link"}

func (t *TezTracker) GetAccountReport(accountID string, from, to int64, operations []string) (resp []byte, err error) {
	var buf bytes.Buffer

	//Check that account is baker
	//isBaker, _, err := t.repoProvider.GetBaker().Find(accountID)
	//if err != nil {
	//	return nil, err
	//}
	isBaker := false

	report, err := t.repoProvider.GetAccount().GetReport(accountID, models.AccountReportFilter{
		From:       from,
		To:         to,
		Operations: operations,
		IsBaker:    isBaker,
	})
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(&buf)
	writer.UseCRLF = true
	writer.Comma = ';'

	writer.Write(header)

	for i := range report {
		//TODO get front host from env
		if report[i].OperationGroupHash.Valid {
			report[i].Link = fmt.Sprintf(frontHost, report[i].OperationGroupHash.String)
		}

		//TODO Marshall to csv form
		writer.Write([]string{report[i].Link})
	}

	//Remove
	data := [][]string{
		{"opaxMZALov5bo6TE47jfeN3gpyMxnjEHn1UKcZjqXP7pKPn91cR",
			"tz1UAxKyit5AmzUauHNvV5TMwCQhQZuJPZc6",
			"2020-07-15 02:14:53"},
		{"opauvrecoE6BN5YHsUN4ivmZroEXEhqLSdhTyFF2sVpyQRqiwKS",
			"tz1dqs6dtzzTfhVvQYUPDXHXZKY3SR8GozTX",
			"2020-07-11 20:47:25"},
		{"opao7MFwkWF1gsZuQVTJkX6W7SsPcWUQQNcpWNhk247PxNH7akr", "tz1WjhcpYaxeV6VAQk9XpTnieBEdgk6eafkq", "2020-06-10 19:57:40"},
	}

	for i := range data {
		err = writer.Write(data[i])
		if err != nil {
			log.Print(err)
		}
	}

	writer.Flush()

	return buf.Bytes(), nil
}
