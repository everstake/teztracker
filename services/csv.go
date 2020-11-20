package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/jszwec/csvutil"
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

	var bakingReq, endorsingReq bool
	//For baker check that baking\endorsing operations required
	if isBaker {
		for i := range operations {
			switch operations[i] {
			case "baking":
				bakingReq = true
			case "endorsement":
				endorsingReq = true
			}
		}
	}

	report, err := t.repoProvider.GetAccount().GetReport(accountID, models.AccountReportFilter{
		From:         from,
		To:           to,
		Operations:   operations,
		EndorsingReq: endorsingReq,
		Limit:        limit,
	})
	if err != nil {
		return nil, err
	}

	var bakingReport []models.BakerReport
	if bakingReq {
		bakingReport, err = t.repoProvider.GetAccount().GetBakingReport(accountID, models.AccountReportFilter{
			From:         from,
			To:           to,
			Operations:   operations,
			EndorsingReq: isBaker,
			Limit:        limit,
		})
		if err != nil {
			return nil, err
		}
	}

	writer := csv.NewWriter(&buf)
	writer.UseCRLF = true
	writer.Comma = ';'

	encoder := csvutil.NewEncoder(writer)

	var j int
	var record interface{}
	for i := 0; i < len(report); {

		//Merge sort
		if j < len(bakingReport) && report[i].BlockLevel <= bakingReport[j].BlockLevel {
			record = bakingReport[j]
			j++
		} else {
			//TODO get front host from env
			if report[i].OperationGroupHash.Valid {
				report[i].Link = fmt.Sprintf(frontHost, report[i].OperationGroupHash.String)
			}

			report[i].In = report[i].Amount
			if report[i].Source == accountID {
				report[i].Out = report[i].Amount
			}

			record = report[i]
			if !isBaker {
				record = report[i].AccountReport
			}
			i++
		}

		err = encoder.Encode(record)
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()

	return buf.Bytes(), nil
}
