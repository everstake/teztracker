package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
	"github.com/go-openapi/strfmt"
)

// Snapshot renders an app level model to a gennerated OpenAPI model.
func Period(p models.PeriodInfo) *genModels.PeriodInfo {
	pi := &genModels.PeriodInfo{
		PeriodType: p.Type,
		Period: &genModels.Period{
			ID:         p.Period,
			StartLevel: p.StartBlock,
			EndLevel:   p.EndBlock,
			StartTime:  strfmt.DateTime(p.StartTime),
			EndTime:    strfmt.DateTime(p.EndTime),
		},
		VoteStats: &genModels.VoteStats{
			NumVoters:      p.Bakers,
			NumVotersTotal: p.TotalBakers,
			VotesAvailable: p.TotalRolls,
			VotesCast:      p.Rolls,
		},
	}

	if p.BallotsStat != nil {
		pi.Ballots = &genModels.Ballots{
			Yay:           p.BallotsStat.Yay,
			Nay:           p.BallotsStat.Nay,
			Pass:          p.BallotsStat.Pass,
			Quorum:        p.BallotsStat.Quorum,
			Supermajority: p.BallotsStat.Supermajority,
		}
	}

	return pi
}
