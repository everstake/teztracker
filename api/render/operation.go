package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)

// Operation renders an app level model to a gennerated OpenAPI model.
func Operation(b models.Operation) *genModels.OperationsRow {
	ts := b.Timestamp.Unix()

	return &genModels.OperationsRow{
		OperationID:         b.OperationID.Ptr(),
		OperationGroupHash:  b.OperationGroupHash.Ptr(),
		Kind:                b.Kind.Ptr(),
		Level:               b.Level,
		Delegate:            b.Delegate,
		Slots:               b.Slots,
		Nonce:               b.Nonce,
		Pkh:                 b.Pkh,
		Secret:              b.Secret,
		Source:              b.Source,
		Fee:                 b.Fee,
		Counter:             b.Counter,
		GasLimit:            b.GasLimit,
		StorageLimit:        b.StorageLimit,
		PublicKey:           b.PublicKey,
		Amount:              b.Amount,
		Destination:         b.Destination,
		Parameters:          b.Parameters,
		ManagerPubkey:       b.ManagerPubkey,
		Balance:             b.Balance,
		Spendable:           b.Spendable,
		Delegatable:         b.Delegatable,
		Script:              b.Script,
		Storage:             b.Storage,
		Status:              b.Status,
		ConsumedGas:         b.ConsumedGas,
		StorageSize:         b.StorageSize,
		PaidStorageSizeDiff: b.PaidStorageSizeDiff,
		OriginatedContracts: b.OriginatedContracts,
		BlockHash:           b.BlockHash.Ptr(),
		BlockLevel:          b.BlockLevel.Ptr(),
		Timestamp:           &ts,
	}
}

// Operations renders a slice of app level Operations into a slice of OpenAPI models.
func Operations(bs []models.Operation) []*genModels.OperationsRow {
	operations := make([]*genModels.OperationsRow, len(bs))
	for i := range bs {
		operations[i] = Operation(bs[i])
	}
	return operations
}
