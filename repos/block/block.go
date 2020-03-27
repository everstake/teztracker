package block

import (
	"fmt"

	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

//go:generate mockgen -source ./block.go -destination ./mock_block/main.go Repo
type (
	// Repository is the block repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		Last() (block models.Block, err error)
		List(limit, offset uint, since uint64) (blocks []models.Block, err error)
		Filter(filter models.BlockFilter) (blocks []models.Block, err error)
		Find(filter models.Block) (found bool, block models.Block, err error)
		FindExtended(filter models.Block) (found bool, block models.Block, err error)
		ListExtended(limit, offset uint, since uint64) (blocks []models.Block, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) getDb() *gorm.DB {
	db := r.db.Select("blocks.*, pb.baker_name, bu.change as reward").
		Model(&models.Block{}).
		Joins("left join tezos.public_bakers as pb on delegate=baker").
		Joins("left join tezos.balance_updates as bu on (source_hash=hash and category='rewards' and source='block')")

	return db
}

// Last returns the last added block.
func (r *Repository) Last() (block models.Block, err error) {
	db := r.getDb()
	err = db.Order("blocks.level desc").First(&block).Error
	return block, err
}

// List returns a list of blocks from the newest to oldest.
// limit defines the limit for the maximum number of blocks returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the blocks with level < since will be returned.
func (r *Repository) List(limit, offset uint, since uint64) (blocks []models.Block, err error) {
	db := r.getDb()
	if since > 0 {
		db = db.Where("level < ?", since)
	}
	err = db.Order("blocks.level desc").
		Limit(limit).
		Offset(offset).
		Find(&blocks).Error
	return blocks, err
}

// Find looks up for a block with filter and extends it with aggregated info.
func (r *Repository) FindExtended(filter models.Block) (found bool, block models.Block, err error) {
	found, block, err = r.Find(filter)
	if err != nil || !found {
		return found, block, err
	}
	blocks, err := r.ExtendBlocks([]models.Block{block})
	if err != nil {
		return false, block, err
	}
	if len(blocks) == 0 {
		return false, block, fmt.Errorf("failed to extend block: empty slice")
	}
	return true, blocks[0], nil
}

// Find looks up for blocks with filter.
func (r *Repository) Find(filter models.Block) (found bool, block models.Block, err error) {
	db := r.getDb()
	if res := db.Where(&filter).Find(&block); res.Error != nil {
		if res.RecordNotFound() {
			return false, block, nil
		}
		return false, block, res.Error
	}
	return true, block, nil
}

// ListBlockAggregation returns a list of block aggreagation data for blocks in levels slice.
func (r *Repository) ListBlockAggregation(levels []int64) (blocks []models.BlockAggregationView, err error) {
	db := r.db.Model(&models.BlockAggregationView{})
	db = db.Where("level IN (?)", levels)
	err = db.Order("level desc").
		Find(&blocks).Error
	return blocks, err
}

// ListExtended returns a list of blocks with populated aggregation data from the newest to oldest.
// limit defines the limit for the maximum number of blocks returned.
// since is used to paginate results based on the level. As the result is ordered descendingly the blocks with level < since will be returned.
func (r *Repository) ListExtended(limit, offset uint, since uint64) (blocks []models.Block, err error) {
	blocks, err = r.List(limit, offset, since)
	if err != nil || len(blocks) == 0 {
		return blocks, err
	}
	return r.ExtendBlocks(blocks)
}

func (r *Repository) ExtendBlocks(blocks []models.Block) (extended []models.Block, err error) {
	count := len(blocks)
	ids := make([]int64, count)
	m := make(map[int64]*models.Block, count)
	for i := range blocks {
		ids[i] = blocks[i].Level.Int64
		m[blocks[i].Level.Int64] = &blocks[i]
	}
	aggInfo, err := r.ListBlockAggregation(ids)
	if err != nil {
		return blocks, err
	}
	for i := range aggInfo {
		level := aggInfo[i].Level
		if b, ok := m[level]; ok {
			b.BlockAggregation = &aggInfo[i]
		}
	}
	return blocks, err
}

func (r *Repository) Filter(filter models.BlockFilter) (blocks []models.Block, err error) {
	db := r.getDb()
	db = db.Or("blocks.level in (?)", filter.BlockLevels).Or("hash in (?)", filter.BlockHashes)
	err = db.Find(&blocks).Error
	return blocks, err
}
