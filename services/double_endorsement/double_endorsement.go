package double_endorsement

import (
	"context"

	"github.com/bullblock-io/tezTracker/models"
	"github.com/bullblock-io/tezTracker/repos/operation"
)

type BakesProvider interface {
	DoubleEndrsementEvidenceLevel(ctx context.Context, blockLevel int, operationHash string) (int64, error)
}

type UnitOfWork interface {
	GetOperation() operation.Repo
}
type LevelUpdater interface {
	UpdateLevel(operation models.Operation) error
}

const limit = 100

func SaveUnprocessedDoubleEndorsementEvidences(ctx context.Context, unit UnitOfWork, provider BakesProvider) (err error) {
	repo := unit.GetOperation()
	endorsements, err := repo.ListDoubleEndorsementsWithoutLevel(limit, 0)
	if err != nil {
		return err
	}
	for i := range endorsements {
		err = SaveDoubleEndorsementEvidenceLevelFor(ctx, endorsements[i], unit.GetOperation(), provider)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveDoubleEndorsementEvidenceLevelFor(ctx context.Context, op models.Operation, repo LevelUpdater, provider BakesProvider) error {
	level, err := provider.DoubleEndrsementEvidenceLevel(ctx, int(op.BlockLevel.Int64), op.OperationGroupHash.String)
	if err != nil {
		return err
	}
	op.Level = level

	return repo.UpdateLevel(op)
}
