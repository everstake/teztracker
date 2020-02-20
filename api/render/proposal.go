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
		Hash:        p.Proposal,
		MinQuorum:   0,
		Period:      p.Period,
		VotesCasted: p.Rolls,
		VotesNum:    p.Bakers,
	}
}
