package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
	"github.com/go-openapi/strfmt"
)

func ProposalVoters(vp []models.ProposalVoter) []*genModels.ProposalVoter {
	votes := make([]*genModels.ProposalVoter, len(vp))
	for i := range vp {
		votes[i] = ProposalVoter(vp[i])
	}
	return votes
}

func ProposalVoter(v models.ProposalVoter) *genModels.ProposalVoter {
	return &genModels.ProposalVoter{
		BlockLevel: v.BlockLevel,
		Name:       v.Alias,
		Operation:  v.Operation,
		Pkh:        v.Pkh,
		Proposal:   v.Proposal,
		Rolls:      v.Rolls,
		Timestamp:  strfmt.DateTime(v.Timestamp),
	}
}

func BallotVoters(vp []models.ProposalVoter) []*genModels.BallotVoter {
	votes := make([]*genModels.BallotVoter, len(vp))
	for i := range vp {
		votes[i] = BallotVoter(vp[i])
	}
	return votes
}

func BallotVoter(v models.ProposalVoter) *genModels.BallotVoter {
	return &genModels.BallotVoter{
		BlockLevel: v.BlockLevel,
		Name:       v.Alias,
		Operation:  v.Operation,
		Pkh:        v.Pkh,
		Decision:   v.Ballot,
		Rolls:      v.Rolls,
		Timestamp:  strfmt.DateTime(v.Timestamp),
	}
}

func NonVoters(vp []models.Voter) []*genModels.NonVoter {
	votes := make([]*genModels.NonVoter, len(vp))
	for i := range vp {
		votes[i] = NonVoter(vp[i])
	}
	return votes
}

func NonVoter(v models.Voter) *genModels.NonVoter {
	return &genModels.NonVoter{
		Name:  v.Alias,
		Pkh:   v.Pkh,
		Rolls: v.Rolls,
	}
}
