package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
	"github.com/go-openapi/strfmt"
)

// Block renders an app level model to a gennerated OpenAPI model.
func Block(b models.Block) *genModels.BlocksRow {
	ts := b.Timestamp.Unix()

	genBlock := genModels.BlocksRow{
		Level:                    b.Level.Ptr(),
		Proto:                    b.Proto.Ptr(),
		BlockTime:                b.BlockTime,
		Predecessor:              b.Predecessor.Ptr(),
		Timestamp:                &ts,
		ValidationPass:           b.ValidationPass.Ptr(),
		Fitness:                  b.Fitness.Ptr(),
		Context:                  b.Context,
		Signature:                b.Signature,
		Protocol:                 b.Protocol.Ptr(),
		Priority:                 b.Priority.Ptr(),
		ChainID:                  b.ChainID,
		Hash:                     b.Hash.Ptr(),
		Reward:                   &b.Reward,
		Deposit:                  b.Deposit,
		OperationsHash:           b.OperationsHash,
		PeriodKind:               b.PeriodKind,
		CurrentExpectedQuorum:    b.CurrentExpectedQuorum,
		ActiveProposal:           b.ActiveProposal,
		Baker:                    b.Baker,
		BakerName:                b.BakerName,
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
		genBlock.DoubleBakingEvidence = b.BlockAggregation.DoubleBakingEvidences
		genBlock.DoubleEndorsementEvidence = b.BlockAggregation.DoubleEndorsementEvidences
		genBlock.NumberOfOperations = b.BlockAggregation.NumberOfOperations
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

// Blocks renders a slice of app level Blocks into a slice of OpenAPI models.
func BlocksBakingRights(bs []models.Block) []*genModels.BakingRightsPerBlock {
	blocks := make([]*genModels.BakingRightsPerBlock, len(bs))
	for i := range bs {
		blocks[i] = BlockBakingRights(bs[i])
	}
	return blocks
}

// BlockBakingRights renders an app level block model into a OpenAPI model.
func BlockBakingRights(b models.Block) *genModels.BakingRightsPerBlock {
	br := genModels.BakingRightsPerBlock{Baker: b.Baker, Level: b.Level.Int64, BlockHash: b.Hash.String, BakerPriority: b.Priority.Ptr()}
	br.Rights = BakingRights(b.BakingRights)
	return &br
}

func BakingRights(br []models.FutureBakingRight) []*genModels.BakingRightsRow {
	rights := make([]*genModels.BakingRightsRow, len(br))
	for i, r := range br {
		rights[i] = BakingRight(r)
	}
	return rights
}

func BakingRight(r models.FutureBakingRight) *genModels.BakingRightsRow {
	priority := int64(r.Priority)
	return &genModels.BakingRightsRow{
		Level:         r.Level,
		Delegate:      r.Delegate,
		DelegateName:  r.DelegateName,
		Priority:      &priority,
		EstimatedTime: strfmt.DateTime(r.EstimatedTime),
		Reward:        r.Reward,
		Deposit:       r.Deposit,
	}
}

func BlocksFutureBakingRights(r []models.FutureBlockBakingRight) []*genModels.FutureBakingRightsPerBlock {
	blocks := make([]*genModels.FutureBakingRightsPerBlock, len(r))
	for i := range r {
		blocks[i] = FutureBlockBakingRights(r[i])
	}
	return blocks
}
func FutureBlockBakingRights(r models.FutureBlockBakingRight) *genModels.FutureBakingRightsPerBlock {
	resp := genModels.FutureBakingRightsPerBlock{Level: r.Level, Rights: make([]*genModels.BakingRightsRow, len(r.Rights))}
	for i := range r.Rights {
		priority := int64(r.Rights[i].Priority)
		resp.Rights[i] = &genModels.BakingRightsRow{
			Delegate:      r.Rights[i].Delegate,
			DelegateName:  r.Rights[i].DelegateName,
			Priority:      &priority,
			EstimatedTime: strfmt.DateTime(r.Rights[i].EstimatedTime),
		}
	}
	return &resp
}
