package double_baking

import (
	"context"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos/double_baking"
	"github.com/bullblock-io/tezTracker/repos/operation"
)

type EvidenceRepo interface {
	Last() (found bool, evidence models.DoubleBakingEvidence, err error)
	Create(evidence models.DoubleBakingEvidence) error
}

type BakesProvider interface {
	DoubleBakingEvidence(ctx context.Context, blockLevel int, operationHash string) (dee models.DoubleBakingEvidence, err error)
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
	evidence, err := provider.DoubleBakingEvidence(ctx, int(op.BlockLevel.Int64), op.OperationGroupHash.String)
	if err != nil {
		return err
	}
	evidence.OperationID = op.OperationID.Int64

	return repo.Create(evidence)
}
