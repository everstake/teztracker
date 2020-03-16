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
		ProposalInfo(proposal string) (models.ProposalInfo, error)
		ProposalsList(id *int64, limit uint) ([]models.VotingProposal, error)
		StatsByKind(periodType string) ([]models.PeriodStats, error)
		VotersList(id int64, kind string, limit uint, offset uint) (periodProposals []models.ProposalVoter, err error)
		VotersCount(id int64, kind string) (count int64, err error)
		PeriodNonVotersList(id, blockLevel int64, limit uint, offset uint) (periodProposals []models.Voter, err error)
		PeriodNonVotersCount(id, blockLevel int64) (count int64, err error)
		ProtocolsList(limit uint, offset uint) ([]models.Protocol, error)
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

func (r *Repository) ProposalsList(id *int64, limit uint) (periodProposals []models.VotingProposal, err error) {
	db := r.db.Select("*, address as pkh").Table("tezos.proposal_stat_view").
		Joins("left join tezos.voting_proposal on proposal = hash").
		Joins("left join tezos.baker_alias on proposer = address").
		Where("kind = 'proposals'")

	if id != nil {
		db = db.Where("period = ?", &id)
	}

	err = db.Limit(limit).Scan(&periodProposals).Error
	if err != nil {
		return periodProposals, err
	}

	return periodProposals, nil
}

func (r *Repository) VotersList(id int64, kind string, limit uint, offset uint) (periodProposals []models.ProposalVoter, err error) {
	err = r.db.Select("v.*, v.source as pkh, operation_group_hash as operation, timestamp, baker_name as alias").Table("tezos.voting_view as v").
		Joins("left join tezos.operations as o on o.block_level = v.block_level and v.source = o.source and v.kind = o.kind").
		Joins("left join tezos.public_bakers as pb on pb.delegate = v.source").
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

func (r *Repository) VotersCount(id int64, kind string) (count int64, err error) {
	err = r.db.Table("tezos.voting_view as v").
		Where("v.period = ? and v.kind = ?", id, kind).
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) ProposalNonVotersList(id, blockLevel int64, limit uint, offset uint) (periodProposals []models.Voter, err error) {
	err = r.db.Select("pkh,r.rolls, baker_name as alias").
		Table("tezos.rolls as r").
		Joins("left join tezos.voting_view as vv on (vv.source = r.pkh and period = ?)", id).
		Joins("left join tezos.public_bakers on pkh = delegate").
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

func (r *Repository) PeriodNonVotersCount(id, blockLevel int64) (count int64, err error) {
	err = r.db.
		Table("tezos.rolls as r").
		Joins("left join tezos.voting_view as vv on (vv.source = r.pkh and period = ?)", id).
		Where("r.block_level = ? and period is null", blockLevel).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) ProposalInfo(proposal string) (proposalInfo models.ProposalInfo, err error) {
	err = r.db.Select("*, address as pkh").
		Table("tezos.voting_proposal as vp").
		Joins("left join tezos.baker_alias on proposer = address").
		Where("hash = ? ", proposal).Find(&proposalInfo).Error
	if err != nil {
		return proposalInfo, err
	}

	return proposalInfo, nil
}

func (r *Repository) ProtocolsList(limit uint, offset uint) (protocolsList []models.Protocol, err error) {
	err = r.db.Select("protocol as hash, min(level) as start_block, max(level) as end_block").
		Table("tezos.blocks").
		Group("protocol").
		Order("start_block asc").
		Limit(limit).
		Offset(offset).
		Scan(&protocolsList).Error
	if err != nil {
		return protocolsList, err
	}

	return protocolsList, nil
}
