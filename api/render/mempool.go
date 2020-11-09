package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	gotezos "github.com/goat-systems/go-tezos/v2"
)

func MempoolOperation(acb gotezos.Operations) *genModels.MempoolOperation {
	return &genModels.MempoolOperation{
		ChainID:   acb.ChainID,
		Hash:      acb.Hash,
		Protocol:  acb.Protocol,
		Branch:    acb.Branch,
		Signature: acb.Signature,
		Contents:  acb.Contents,
	}
}

func MempoolOperationsList(ops []gotezos.Operations) []*genModels.MempoolOperation {
	mps := make([]*genModels.MempoolOperation, len(ops))
	for i := range ops {
		mps[i] = MempoolOperation(ops[i])
	}

	return mps
}
