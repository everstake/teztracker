package voting_periods

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the snapshots repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		List() ([]models.PeriodInfo, error)
		Info(id int64) (models.PeriodStats, error)
		GetCurrentPeriodId() (int64, error)
		BallotsList(id int64) ([]models.PeriodBallot, error)
		ProposalsList(id int64, limit uint) ([]models.VotingProposal, error)
		StatsByKind(periodType string) ([]models.PeriodStats, error)
		VotersList(id int64, kind string, limit uint, offset uint) (periodProposals []models.ProposalVoter, err error)
		ProposalNonVotersList(id, blockLevel int64, limit uint, offset uint) (periodProposals []models.Voter, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() (periods []models.PeriodInfo, err error) {
	err = r.db.Select("*").
		Table("tezos.voting_period").
		Order("id asc").
		Scan(&periods).Error
	if err != nil {
		return periods, err
	}

	return periods, nil
}

func (r *Repository) GetCurrentPeriodId() (id int64, err error) {
	period := struct {
		ID int64
	}{}

	err = r.db.Select("max(id) as id").Table("tezos.voting_period").
		Find(&period).Error
	if err != nil {
		return 0, err
	}

	return period.ID, nil
}

func (r *Repository) Info(id int64) (periodInfo models.PeriodStats, err error) {
	err = r.db.Select("vp.*, psw.*, ptsv.total_rolls, ptsv.total_bakers").Table("tezos.voting_period as vp").
		Joins("left join tezos.period_stat_view as psw on id = psw.period").
		Joins("left join tezos.period_total_stat_view as ptsv on id = ptsv.period").
		Where("id = ?", id).
		Find(&periodInfo).Error
	if err != nil {
		return periodInfo, err
	}

	return periodInfo, nil
}

func (r *Repository) StatsByKind(periodKind string) (periods []models.PeriodStats, err error) {
	err = r.db.Select("psv.*, ptsv.total_rolls, ptsv.total_bakers").
		Table("tezos.period_stat_view as psv").
		Joins("left join tezos.period_total_stat_view as ptsv on psv.period = ptsv.period").
		Where("kind = ?", periodKind).
		Order("psv.period asc").Scan(&periods).Error
	if err != nil {
		return periods, err
	}

	return periods, nil
}

func (r *Repository) BallotsList(id int64) (periodBallots []models.PeriodBallot, err error) {
	err = r.db.Select("*").Table("tezos.proposal_stat_view").
		Where("period = ?", id).Scan(&periodBallots).Error
	if err != nil {
		return periodBallots, err
	}

	return periodBallots, nil
}

func (r *Repository) ProposalsList(id int64, limit uint) (periodProposals []models.VotingProposal, err error) {
	err = r.db.Select("*").Table("tezos.proposal_stat_view").
		Where("period = ? and kind = 'proposals'", id).
		Limit(limit).Scan(&periodProposals).Error
	if err != nil {
		return periodProposals, err
	}

	return periodProposals, nil
}

func (r *Repository) VotersList(id int64, kind string, limit uint, offset uint) (periodProposals []models.ProposalVoter, err error) {
	err = r.db.Select("v.*, v.source as pkh, operation_group_hash as operation, timestamp, name as alias").Table("tezos.voting_view as v").
		Joins("left join tezos.operations as o on o.block_level = v.block_level and v.source = o.source and v.kind = o.kind").
		Joins("left join tezos.baker_alias as ba on ba.address = v.source").
		Where("v.period = ? and v.kind = ?", id, kind).
		Order(" v.block_level  desc").
		Limit(limit).
		Offset(offset).
		Scan(&periodProposals).Error
	if err != nil {
		return periodProposals, err
	}

	return periodProposals, nil
}

func (r *Repository) ProposalNonVotersList(id, blockLevel int64, limit uint, offset uint) (periodProposals []models.Voter, err error) {
	err = r.db.Select("pkh,r.rolls,name as alias").
		Table("tezos.rolls as r").
		Joins("left join tezos.voting_view as vv on (vv.source = r.pkh and period = ?)", id).
		Joins("left join tezos.baker_alias on pkh = address").
		Where("r.block_level = ? and period is null", blockLevel).
		Order("r.rolls desc").
		Limit(limit).
		Offset(offset).
		Scan(&periodProposals).Error
	if err != nil {
		return periodProposals, err
	}

	return periodProposals, nil
}
