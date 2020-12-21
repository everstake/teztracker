package user_profile

import (
	"github.com/everstake/teztracker/models"
	"github.com/jinzhu/gorm"
)

type (
	// Repository is the user profile repo implementation.
	Repository struct {
		db *gorm.DB
	}

	Repo interface {
		FindUserByAccount(accountID string) (user models.User, found bool, err error)
		CreateUser(user models.User) error
		UpdateUser(user models.User) error
		GetUsersAndAddresses(addresses []string) (users []models.UserAndAddress, err error)

		GetUserAddresses(accountID string) (addresses []models.UserAddress, err error)
		GetUserAddress(accountID string, address string) (model models.UserAddress, found bool, err error)
		CreateUserAddress(address models.UserAddress) error
		DeleteUserAddress(accountID string, address string) error
		UpdateUserAddress(address models.UserAddress) error
		GetUserAddressesCount(accountID string) (count uint64, err error)

		UserNotesList(accountID string) (notes []models.UserNote, err error)
		FindUserNote(accountID string, text string) (note models.UserNote, found bool, err error)
		CreateUserNote(models.UserNote) error
		DeleteUserNote(accountID string, text string) error
		UpdateUserNote(models.UserNote) error
		GetUserNotesCount(accountID string) (count uint64, err error)
	}
)

// New creates an instance of repository using the provided db.
func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Model(&models.User{}),
	}
}

func (r *Repository) FindUserByAccount(accountID string) (user models.User, found bool, err error) {
	if res := r.db.Where("account_id = ?", accountID).First(&user); res.Error != nil {
		if res.RecordNotFound() {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

func (r *Repository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *Repository) UpdateUser(user models.User) error {
	return r.db.Where("account_id = ?", user.AccountID).Update(&user).Error
}

func (r *Repository) GetUsersAndAddresses(addresses []string) (items []models.UserAndAddress, err error) {
	err = r.db.Select("user_addresses.*, users.email").
		Where("user_addresses.address in (?)", addresses).
		Table("tezos.user_addresses").
		Joins("left join tezos.users ON (user_addresses.account_id = users.account_id)").
		Find(&items).
		Error
	return items, err
}

func (r *Repository) GetUserAddresses(accountID string) (addresses []models.UserAddress, err error) {
	err = r.db.Where("account_id = ?", accountID).Find(&addresses).Error
	return addresses, err
}

func (r *Repository) GetUserAddress(accountID string, address string) (model models.UserAddress, found bool, err error) {
	res := r.db.Model(&models.UserAddress{}).
		Where("account_id = ?", accountID).
		Where("address = ?", address).
		First(&address)
	if res.RecordNotFound() {
		return model, false, nil
	}
	if res.Error != nil {
		return model, false, err
	}
	return model, true, nil
}

func (r *Repository) CreateUserAddress(address models.UserAddress) error {
	return r.db.Model(&models.UserAddress{}).Create(&address).Error
}

func (r *Repository) DeleteUserAddress(accountID string, address string) error {
	return r.db.Delete(&models.UserAddress{}).
		Where("account_id = ?", accountID).
		Where("address = ?", address).Error
}

func (r *Repository) UpdateUserAddress(address models.UserAddress) error {
	return r.db.Model(&models.UserAddress{}).
		Where("account_id = ?", address.AccountID).
		Where("address = ?", address).Update(&address).Error
}

func (r *Repository) GetUserAddressesCount(accountID string) (count uint64, err error) {
	 err = r.db.Model(&models.UserAddress{}).
	 	Select("count(*)").
		Where("account_id = ?", accountID).First(&count).Error
	 return count, err
}

func (r *Repository) UserNotesList(accountID string) (notes []models.UserNote, err error) {
	err = r.db.Model(&models.UserAddress{}).Where("account_id = ?", accountID).Find(&notes).Error
	return notes, err
}

func (r *Repository) FindUserNote(accountID string, text string) (note models.UserNote, found bool, err error) {
	if res := r.db.Model(&models.UserNote{}).Where("account_id = ?", accountID).Where("text = ?", text).First(&note); res.Error != nil {
		if res.RecordNotFound() {
			return note, false, nil
		}
		return note, false, err
	}
	return note, true, nil
}

func (r *Repository) CreateUserNote(note models.UserNote) error {
	return r.db.Model(&models.UserNote{}).Create(&note).Error
}

func (r *Repository) UpdateUserNote(note models.UserNote) error {
	return r.db.Where("account_id = ?", note.AccountID).Where("text = ?", note.Text).Update(&note).Error
}

func (r *Repository) DeleteUserNote(accountID string, text string) error {
	return r.db.Delete(&models.UserNote{}).
		Where("account_id = ?", accountID).
		Where("text = ?", text).Error
}

func (r *Repository) GetUserNotesCount(accountID string) (count uint64, err error) {
	err = r.db.Model(&models.UserNote{}).
		Select("count(*)").
		Where("account_id = ?", accountID).First(&count).Error
	return count, err
}

