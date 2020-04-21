package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Operation renders an app level model to a gennerated OpenAPI model.
func Operation(b models.Operation, dbe *models.DoubleOperationEvidenceExtended) *genModels.OperationsRow {
	ts := b.Timestamp.Unix()

	row := genModels.OperationsRow{
		OperationID:         b.OperationID.Ptr(),
		OperationGroupHash:  b.OperationGroupHash.Ptr(),
		Kind:                b.Kind.Ptr(),
		Level:               b.Level,
		Delegate:            b.Delegate,
		DelegateName:        b.DelegateName,
		Slots:               b.Slots,
		Nonce:               b.Nonce,
		Pkh:                 b.Pkh,
		Secret:              b.Secret,
		Source:              b.Source,
		SourceName:          b.SourceName,
		Fee:                 b.Fee,
		Counter:             b.Counter,
		Reward:              b.Reward,
		GasLimit:            b.GasLimit,
		StorageLimit:        b.StorageLimit,
		PublicKey:           b.PublicKey,
		Amount:              b.Amount,
		Destination:         b.Destination,
		DestinationName:     b.DestinationName,
		Parameters:          b.Parameters,
		ManagerPubkey:       b.ManagerPubkey,
		Balance:             b.Balance,
		Spendable:           b.Spendable,
		Delegatable:         b.Delegatable,
		DelegationAmount:    b.DelegationAmount,
		Script:              b.Script,
		Storage:             b.Storage,
		Status:              b.Status,
		ConsumedGas:         b.ConsumedGas,
		StorageSize:         b.StorageSize,
		PaidStorageSizeDiff: b.PaidStorageSizeDiff,
		OriginatedContracts: b.OriginatedContracts,
		BlockHash:           b.BlockHash.Ptr(),
		BlockLevel:          b.BlockLevel.Ptr(),
		Ballot:              b.Ballot,
		Proposal:            b.Proposal,
		Cycle:               b.Cycle,
		Confirmations:       &b.Confirmations,
		EndorsementReward:   b.EndorsementReward,
		EndorsementDeposit:  b.EndorsementDeposit,
		ClaimedAmount:       b.ClaimedAmount,
		Timestamp:           &ts,
	}
	if dbe != nil {
		row.DoubleOperationDetails = &genModels.DoubleOperationDetails{
			BakerReward:       dbe.BakerReward,
			DenouncedLevel:    dbe.DenouncedLevel,
			EvidenceBaker:     dbe.EvidenceBaker,
			EvidenceBakerName: dbe.EvidenceBakerName,
			LostDeposits:      dbe.LostDeposits,
			LostFees:          dbe.LostFees,
			LostRewards:       dbe.LostRewards,
			Offender:          dbe.Offender,
			OffenderName:      dbe.OffenderName,
			Priority:          int64(dbe.Priority),
		}
	}
	return &row
}

// Operations renders a slice of app level Operations into a slice of OpenAPI models.
func Operations(bs []models.Operation) []*genModels.OperationsRow {
	operations := make([]*genModels.OperationsRow, len(bs))
	for i := range bs {
		operations[i] = Operation(bs[i], bs[i].DoubleOperationEvidenceExtended)
	}
	return operations
}

func DoubleOperations(bs []models.DoubleOperationEvidenceExtended) []*genModels.OperationsRow {
	operations := make([]*genModels.OperationsRow, len(bs))
	for i := range bs {
		operations[i] = Operation(bs[i].Operation, &bs[i])
	}
	return operations
}
