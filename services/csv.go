package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/jszwec/csvutil"
	"log"
)

const (
	limit     = 250000
	frontHost = "https://teztracker.everstake.one/en/mainnet/tx/%s"
)

var header = []string{"block level", "timestamp", "operation", "coin", "in", "out", "from", "to", "fee", "reward", "loss", "status", "link"}

func (t *TezTracker) GetAccountReport(accountID string, from, to int64, operations []string) (resp []byte, err error) {
	var buf bytes.Buffer

	//Check that account is baker
	isBaker, _, err := t.repoProvider.GetBaker().Find(accountID)
	if err != nil {
		return nil, err
	}

	report, err := t.repoProvider.GetAccount().GetReport(accountID, models.AccountReportFilter{
		From:       from,
		To:         to,
		Operations: operations,
		IsBaker:    isBaker,
		Limit:      limit,
	})
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(&buf)
	writer.UseCRLF = true
	writer.Comma = ';'

	writer.Write(header)

	encoder := csvutil.NewEncoder(writer)

	for i := range report {
		//TODO get front host from env
		if report[i].OperationGroupHash.Valid {
			report[i].Link = fmt.Sprintf(frontHost, report[i].OperationGroupHash.String)
		}

		if report[i].Source == accountID {
			report[i].Out = report[i].Amount
		} else {
			report[i].In = report[i].Amount
		}

		err = encoder.Encode(report[i])
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()

	return buf.Bytes(), nil
}
