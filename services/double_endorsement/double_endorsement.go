package double_endorsement

import (
	"context"
	"github.com/everstake/teztracker/repos/double_endorsement"

	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/operation"
)

type EvidenceRepo interface {
	Last() (found bool, evidence models.DoubleOperationEvidenceExtended, err error)
	Create(evidence models.DoubleOperationEvidence) error
}

type BakesProvider interface {
	DoubleOperationEvidence(ctx context.Context, blockLevel int, operationHash string) (dee models.DoubleOperationEvidence, err error)
}

type UnitOfWork interface {
	GetOperation() operation.Repo
	GetDoubleEndorsement() double_endorsement.Repo
}

const limit = 100

func SaveUnprocessedDoubleEndorsementEvidences(ctx context.Context, unit UnitOfWork, provider BakesProvider) (err error) {
	repo := unit.GetDoubleEndorsement()
	found, lastEvidence, err := repo.Last()
	if err != nil {
		return err
	}
	lastKnownOperationID := int64(0)
	if found {
		lastKnownOperationID = lastEvidence.OperationID
	}

	newDoubleBakes, err := unit.GetOperation().ListAsc([]string{"double_endorsement_evidence"}, limit, 0, lastKnownOperationID)
	if err != nil {
		return err
	}

	for i := range newDoubleBakes {
		err = SaveDoubleEndorsementEvidenceFor(ctx, newDoubleBakes[i], unit.GetDoubleEndorsement(), provider)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveDoubleEndorsementEvidenceFor(ctx context.Context, op models.Operation, repo EvidenceRepo, provider BakesProvider) error {
	evidence, err := provider.DoubleOperationEvidence(ctx, int(op.BlockLevel.Int64), op.OperationGroupHash.String)
	if err != nil {
		return err
	}
	evidence.OperationID = op.OperationID.Int64
	evidence.Type = models.DoubleOperationTypeEndorsement

	return repo.Create(evidence)
}
