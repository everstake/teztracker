package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func Proposals(vp []models.VotingProposal) []*genModels.Proposal {
	proposals := make([]*genModels.Proposal, len(vp))
	for i := range vp {
		proposals[i] = Proposal(vp[i])
	}
	return proposals
}

func Proposal(p models.VotingProposal) *genModels.Proposal {
	return &genModels.Proposal{
		Hash:         p.Proposal,
		Period:       p.Period,
		VotesCasted:  p.Rolls,
		VotesNum:     p.Bakers,
		ProposalFile: p.ProposalFile,
		Proposer: &genModels.ProposalProposer{
			Name: p.Name,
			Pkh:  p.Pkh,
		},
		ShortDescription: p.ShortDescription,
		Title:            p.Title,
	}
}
