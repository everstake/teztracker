package models

import "time"

type PeriodType string

type PeriodInfo struct {
	Rolls       int64
	Bakers      int64
	BlockLevel  int64
	Period      int64
	Type        string
	Kind        string
	StartBlock  int64
	EndBlock    int64
	Cycle       int8
	StartTime   time.Time
	EndTime     time.Time
	TotalBakers int64
	TotalRolls  int64
	BallotsStat *BallotsStat
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