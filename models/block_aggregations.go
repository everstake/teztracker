package models

type BlockAggregationView struct {
	Level                int64 `json:"level"`
	Volume               int64 `json:"volume"`
	Fees                 int64 `json:"fees"`
	Endorsements         int64 `json:"endorsements"`
	Proposals            int64 `json:"proposals"`
	SeedNonceRevelations int64 `json:"seed_nonce_revelations"`
	Delegations          int64 `json:"delegations"`
	Transactions         int64 `json:"transactions"`
	ActivateAccounts     int64 `json:"activate_accounts"`
	Ballots              int64 `json:"ballots"`
	Originations         int64 `json:"originations"`
	Reveals              int64 `json:"reveals"`
	DoubleBakingEvidence int64 `json:"double_baking_evidences"`
}

func (*BlockAggregationView) TableName() string {
	return "block_aggregation_view"
}
