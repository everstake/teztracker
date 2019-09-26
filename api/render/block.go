package render

import (
	genModels "github.com/bullblock-io/tezTracker/gen/models"
	"github.com/bullblock-io/tezTracker/models"
)

// Block renders an app level model to a gennerated OpenAPI model.
func Block(b models.Block) *genModels.BlocksRow {
	ts := b.Timestamp.Unix()

	genBlock := genModels.BlocksRow{
		Level:                    b.Level.Ptr(),
		Proto:                    b.Proto.Ptr(),
		Predecessor:              b.Predecessor.Ptr(),
		Timestamp:                &ts,
		ValidationPass:           b.ValidationPass.Ptr(),
		Fitness:                  b.Fitness.Ptr(),
		Context:                  b.Context,
		Signature:                b.Signature,
		Protocol:                 b.Protocol.Ptr(),
		ChainID:                  b.ChainID,
		Hash:                     b.Hash.Ptr(),
		OperationsHash:           b.OperationsHash,
		PeriodKind:               b.PeriodKind,
		CurrentExpectedQuorum:    b.CurrentExpectedQuorum,
		ActiveProposal:           b.ActiveProposal,
		Baker:                    b.Baker,
		NonceHash:                b.NonceHash,
		ConsumedGas:              b.ConsumedGas,
		MetaLevel:                b.MetaLevel,
		MetaLevelPosition:        b.MetaLevelPosition,
		MetaCycle:                b.MetaCycle,
		MetaCyclePosition:        b.MetaCyclePosition,
		MetaVotingPeriod:         b.MetaVotingPeriod,
		MetaVotingPeriodPosition: b.MetaVotingPeriodPosition,
		ExpectedCommitment:       b.ExpectedCommitment,
	}

	if b.BlockAggregation != nil {
		genBlock.Volume = b.BlockAggregation.Volume
		genBlock.Fees = b.BlockAggregation.Fees
		genBlock.Endorsements = b.BlockAggregation.Endorsements
		genBlock.Proposals = b.BlockAggregation.Proposals
		genBlock.SeedNonceRevelations = b.BlockAggregation.SeedNonceRevelations
		genBlock.Delegations = b.BlockAggregation.Delegations
		genBlock.Transactions = b.BlockAggregation.Transactions
		genBlock.ActivateAccounts = b.BlockAggregation.ActivateAccounts
		genBlock.Ballots = b.BlockAggregation.Ballots
		genBlock.Originations = b.BlockAggregation.Originations
		genBlock.Reveals = b.BlockAggregation.Reveals
		genBlock.DoubleBakingEvidence = b.BlockAggregation.DoubleBakingEvidence
	}

	return &genBlock
}

// Blocks renders a slice of app level Blocks into a slice of OpenAPI models.
func Blocks(bs []models.Block) []*genModels.BlocksRow {
	blocks := make([]*genModels.BlocksRow, len(bs))
	for i := range bs {
		blocks[i] = Block(bs[i])
	}
	return blocks
}

// BlockResult renders an app level block model into a OpenAPI model with operation groups.
func BlockResult(b models.Block) *genModels.BlockResult {
	groups := make([]*genModels.OperationGroupsRow, len(b.OperationGroups))
	for i, og := range b.OperationGroups {
		if og == nil {
			continue
		}
		groups[i] = OperationGroup(*og)
	}

	br := genModels.BlockResult{Block: Block(b), OperationGroups: groups}
	return &br
}
