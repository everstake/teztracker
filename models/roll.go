package models

type Roll struct {
	Pkh          string `json:"pkh"`
	Rolls        int64  `json:"rolls"`
	BlockLevel   int64  `json:"block_level"`
	VotingPeriod int64  `json:"voting_period"`
	Cycle        int64  `json:"cycle"`
}
