package double_baking

import (
	"context"

	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/double_baking"
	"github.com/everstake/teztracker/repos/operation"
)

type EvidenceRepo interface {
	Last() (found bool, evidence models.DoubleOperationEvidence, err error)
	Create(evidence models.DoubleOperationEvidence) error
}

type BakesProvider interface {
	DoubleOperationEvidence(ctx context.Context, blockLevel int, operationHash string) (dee models.DoubleOperationEvidence, err error)
}

type UnitOfWork interface {
	GetDoubleBaking() double_baking.Repo
	GetOperation() operation.Repo
}

const limit = 100

func SaveUnprocessedDoubleBakingEvidences(ctx context.Context, unit UnitOfWork, provider BakesProvider) (err error) {
	repo := unit.GetDoubleBaking()
	found, lastEvidence, err := repo.Last()
	if err != nil {
		return err
	}
	lastKnownOperationID := int64(0)
	if found {
		lastKnownOperationID = lastEvidence.OperationID
	}

	newDoubleBakes, err := unit.GetOperation().ListAsc([]string{"double_baking_evidence"}, limit, 0, lastKnownOperationID)
	if err != nil {
		return err
	}

	for i := range newDoubleBakes {
		err = SaveDoubleBakingEvidenceFor(ctx, newDoubleBakes[i], unit.GetDoubleBaking(), provider)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveDoubleBakingEvidenceFor(ctx context.Context, op models.Operation, repo EvidenceRepo, provider BakesProvider) error {
	evidence, err := provider.DoubleOperationEvidence(ctx, int(op.BlockLevel.Int64), op.OperationGroupHash.String)
	if err != nil {
		return err
	}
	evidence.OperationID = op.OperationID.Int64
	evidence.Type = models.DoubleOperationTypeBaking

	return repo.Create(evidence)
}
