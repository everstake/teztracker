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

	var bakingReq, endorsingReq, assetsReq bool

	for i := range operations {

		if operations[i] == "assets" {
			assetsReq = true
		}

		//For baker check that baking\endorsing\assets operations required
		if isBaker {
			switch operations[i] {
			case "baking":
				bakingReq = true
			case "endorsement":
				endorsingReq = true
			}

		}
	}

	report, err := t.repoProvider.GetAccount().GetReport(accountID, models.ReportFilter{
		From:         from,
		To:           to,
		Operations:   operations,
		EndorsingReq: endorsingReq,
		AssetsReq:    assetsReq,
		Limit:        limit,
	})
	if err != nil {
		return nil, err
	}

	var bakingReport []models.ExtendReport
	if bakingReq {
		bakingReport, err = t.repoProvider.GetAccount().GetBakingReport(accountID, models.ReportFilter{
			From:       from,
			To:         to,
			Operations: operations,
			Limit:      limit,
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
				record = report[i].OperationReport
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

func (t *TezTracker) GetAssetReport(assetID string, from, to int64, operations []string) (resp []byte, err error) {
	repo := t.repoProvider.GetAssets()
	token, err := repo.GetTokenInfo(assetID)
	if err != nil {
		return nil, err
	}

	report, err := repo.GetAssetReport(token.ID, models.ReportFilter{
		From:  from,
		To:    to,
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	writer := csv.NewWriter(&buf)
	writer.UseCRLF = true
	writer.Comma = ';'

	encoder := csvutil.NewEncoder(writer)

	for i := range report {
		if report[i].OperationGroupHash.Valid {
			report[i].Link = fmt.Sprintf(frontHost, report[i].OperationGroupHash.String)
		}

		err = encoder.Encode(report[i])
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()

	return buf.Bytes(), nil
}
