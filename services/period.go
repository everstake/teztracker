package services

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
	"math"
)

const (
	cyclesOnVotingPeriod      = 8
	firstPeriodStartBlock     = 1
	periodKindProposal        = "proposals"
	periodKindBallot          = "ballot"
	quorumFormulaChangePeriod = 19
	supermajority             = 80
	genesisQuorum             = 0.8
	minQuorum                 = 0.2
	maxQuorum                 = 0.7
)

func (t *TezTracker) VotingPeriodsList() ([]models.PeriodInfo, error) {
	periods, err := t.repoProvider.GetVotingPeriod().List()
	if err != nil {
		return periods, err
	}

	return periods, nil
}

func (t *TezTracker) VotingPeriodStats(id *int64) (info models.PeriodStats, err error) {
	repo := t.repoProvider.GetVotingPeriod()
	if id == nil {
		info.Period, err = repo.GetCurrentPeriodId()
		if err != nil {
			return
		}
	} else {
		info.Period = *id
	}

	info, err = repo.Info(info.Period)
	if err != nil {
		return info, err
	}

	var proposalHash string
	if info.Kind == periodKindBallot {

		ballots, err := repo.BallotsList(info.Period)
		if err != nil {
			return info, err
		}

		var b models.BallotsStat
		for _, value := range ballots {
			switch value.Ballot {
			case "yay":
				b.Yay = value.Rolls
			case "nay":
				b.Nay = value.Rolls
			case "pass":
				b.Pass = value.Rolls
			}

			proposalHash = value.Proposal
		}

		b.Supermajority = supermajority

		b.Quorum, err = t.calcQuorumForPeriod(info.Period)

		info.BallotsStat = &b
	}

	_, block, err := t.repoProvider.GetBlock().Find(models.Block{
		Level: null.IntFrom(info.StartBlock),
	})
	if err != nil {
		return info, err
	}

	//If block have active proposal replace
	//for old proposals
	if block.ActiveProposal != "" {
		proposalHash = block.ActiveProposal
	}

	//Get additional information if proposalHash presented
	if proposalHash != "" {
		proposalInfo, err := repo.ProposalInfo(proposalHash)
		if err != nil {
			return info, err
		}

		info.Proposal = &proposalInfo
	}

	return info, nil
}

func (t *TezTracker) calcQuorumForPeriod(id int64) (quorum float64, err error) {
	stats, err := t.repoProvider.GetVotingPeriod().StatsByKind(periodKindBallot)
	if err != nil {
		return quorum, err
	}

	calcFormula := QuorumOldFormula
	quorum = genesisQuorum
	//Used on new formula
	avP := genesisQuorum

	if len(stats) == 0 {
		return quorum, nil
	}

	if stats[0].Period == id {
		return quorum, nil
	}

	for i := 1; i < len(stats); i++ {

		if stats[i].Period > quorumFormulaChangePeriod {
			calcFormula = QuorumNewFormula
		}

		quorum, avP = calcFormula(avP, float64(stats[i-1].Rolls)/float64(stats[i-1].TotalRolls))
		//Truncate
		avP = math.Trunc(quorum*10000) / 10000
		quorum = math.Trunc(quorum*10000) / 10000

		if stats[i].Period == id {
			return quorum, nil
		}
	}

	return 0, fmt.Errorf("Level not found")
}

type QuorumFormula func(prevAverPart, actualPart float64) (quorum float64, averPart float64)

func QuorumOldFormula(prevAverPart, actualPart float64) (q float64, avP float64) {
	avP = 0.8*prevAverPart + 0.2*actualPart
	return avP, avP
}

func QuorumNewFormula(prevAverPart, actualPart float64) (q float64, avP float64) {
	_, avP = QuorumOldFormula(prevAverPart, actualPart)
	return minQuorum + avP*(maxQuorum-minQuorum), avP
}

func (t *TezTracker) ProposalsByPeriodID(id *int64, limits Limiter) (proposals []models.VotingProposal, count int64, err error) {
	proposals, err = t.repoProvider.GetVotingPeriod().ProposalsList(id, limits.Limit())
	if err != nil {
		return proposals, 0, err
	}

	return proposals, 0, nil
}

func (t *TezTracker) GetProposalVoters(id int64, limits Limiter) (votes []models.ProposalVoter, count int64, err error) {
	periodRepo := t.repoProvider.GetVotingPeriod()
	count, err = periodRepo.VotersCount(id, periodKindProposal)
	if err != nil {
		return votes, 0, err
	}

	votes, err = periodRepo.VotersList(id, periodKindProposal, limits.Limit(), limits.Offset())
	if err != nil {
		return votes, 0, err
	}

	return votes, count, nil
}

func (t *TezTracker) GetBallotVoters(id int64, limits Limiter) (votes []models.ProposalVoter, count int64, err error) {
	periodRepo := t.repoProvider.GetVotingPeriod()
	count, err = periodRepo.VotersCount(id, periodKindBallot)
	if err != nil {
		return votes, 0, err
	}

	votes, err = t.repoProvider.GetVotingPeriod().VotersList(id, periodKindBallot, limits.Limit(), limits.Offset())
	if err != nil {
		return votes, 0, err
	}

	return votes, count, nil
}

func (t *TezTracker) GetProtocolsList(limits Limiter) (protocols []models.Protocol, count int64, err error) {
	protocols, err = t.repoProvider.GetVotingPeriod().ProtocolsList(limits.Limit(), limits.Offset())
	if err != nil {
		return protocols, 0, err
	}

	return protocols, count, nil
}

func (t *TezTracker) GetPeriodNonVoters(id int64, limits Limiter) (proposals []models.Voter, count int64, err error) {
	periodRepo := t.repoProvider.GetVotingPeriod()

	count, err = periodRepo.PeriodNonVotersCount(id)
	if err != nil {
		return proposals, count, err
	}

	proposals, err = periodRepo.PeriodNonVotersList(id, limits.Limit(), limits.Offset())
	if err != nil {
		return proposals, 0, err
	}

	return proposals, count, err
}

func (t *TezTracker) calcVotingPeriod(id int64) (startBlock, endBlock int64) {
	blocksInPeriod := t.BlocksInCycle() * cyclesOnVotingPeriod
	startBlock = blocksInPeriod*id + firstPeriodStartBlock
	return startBlock, startBlock + blocksInPeriod - 1
}
