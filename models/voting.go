package models

import "time"

type PeriodType string

type PeriodStats struct {
	Rolls       int64
	Bakers      int64
	BlockLevel  int64
	Period      int64
	Kind        string
	TotalBakers int64
	TotalRolls  int64
	BallotsStat *BallotsStat
	Proposal    *VotingProposal
	PeriodInfo
}

type PeriodInfo struct {
	ID         int64
	Type       string
	StartBlock int64
	EndBlock   int64
	StartTime  time.Time
	EndTime    time.Time
}

type BallotsStat struct {
	Yay           int64
	Nay           int64
	Pass          int64
	Quorum        float64
	Supermajority float64
}

type PeriodBallot struct {
	Rolls    int64
	Ballot   string
	Proposal string
}

type ProposalVoter struct {
	Proposal   string
	BlockLevel int64
	Operation  string
	Timestamp  time.Time
	Ballot     string
	Voter
}

type Voter struct {
	Pkh   string
	Rolls int64
	Alias string
}

type VotingProposal struct {
	PeriodBallot
	Bakers     int64
	BlockLevel int64
	Period     int64
	Kind       string
}
