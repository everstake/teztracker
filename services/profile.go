package services

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/guregu/null"
)

const (
	userAddressesLimit = 50
	userNotesLimit     = 50
)

func (t *TezTracker) GetOrCreateUser(account string) (user models.User, err error) {
	profileRepo := t.repoProvider.GetUserProfile()
	user, found, err := profileRepo.FindUserByAccount(account)
	if err != nil {
		return user, fmt.Errorf("repo.GetUserProfile: %s", err.Error())
	}
	if found {
		return user, nil
	}
	accountRepo := t.repoProvider.GetAccount()
	accFound, _, err := accountRepo.Find(models.Account{
		AccountID: null.NewString(account, true),
	})
	if err != nil {
		return user, fmt.Errorf("accountRepo.Find: %s", err.Error())
	}
	if !accFound {
		return user, models.AccountNotFoundErr
	}
	user = models.User{
		AccountID: account,
	}
	err = profileRepo.CreateUser(user)
	if err != nil {
		return user, fmt.Errorf("profileRepo.CreateUser: %s", err.Error())
	}
	return user, nil
}

func (t *TezTracker) UpdateUser(user models.User) error {
	return t.repoProvider.GetUserProfile().UpdateUser(user)
}

func (t *TezTracker) GetUserAddresses(accountID string) (addresses []models.UserAddress, err error) {
	return t.repoProvider.GetUserProfile().GetUserAddresses(accountID)
}

func (t *TezTracker) CreateOrUpdateUserAddress(userAddress models.UserAddress) error {
	userProfileRepo := t.repoProvider.GetUserProfile()
	user, found, err := userProfileRepo.GetUserAddress(userAddress.AccountID, userAddress.Address)
	if err != nil {
		return fmt.Errorf("GetUserAddress: %s", err.Error())
	}
	if found {
		err = userProfileRepo.UpdateUserAddress(userAddress)
		if err != nil {
			return fmt.Errorf("UpdateUserAddress: %s", err.Error())
		}
		return nil
	}
	accountRepo := t.repoProvider.GetAccount()
	accFound, _, err := accountRepo.Find(models.Account{
		AccountID: null.NewString(userAddress.Address, true),
	})
	if !accFound {
		return models.AccountNotFoundErr
	}
	count, err := userProfileRepo.GetUserAddressesCount(user.AccountID)
	if err != nil {
		return fmt.Errorf("userProfileRepo.GetUserAddressesCount: %s", err.Error())
	}
	if count == userAddressesLimit {
		return models.UserLimitReachedErr
	}
	err = userProfileRepo.CreateUserAddress(userAddress)
	if err != nil {
		return fmt.Errorf("CreateUserAddress: %s", err.Error())
	}
	return nil
}

func (t *TezTracker) DeleteUserAddress(accountID string, address string) error {
	return t.repoProvider.GetUserProfile().DeleteUserAddress(accountID, address)
}

func (t *TezTracker) GetUserNotes(accountID string) (notes []models.UserNote, err error) {
	return t.repoProvider.GetUserProfile().UserNotesList(accountID)
}

func (t *TezTracker) CreateOrUpdateUserNote(note models.UserNote) error {
	userProfileRepo := t.repoProvider.GetUserProfile()
	user, found, err := userProfileRepo.FindUserNote(note.AccountID, note.Text)
	if err != nil {
		return fmt.Errorf("FindUserNote: %s", err.Error())
	}
	if found {
		err = userProfileRepo.UpdateUserNote(note)
		if err != nil {
			return fmt.Errorf("UpdateUserNote: %s", err.Error())
		}
		return nil
	}
	count, err := userProfileRepo.GetUserNotesCount(user.AccountID)
	if err != nil {
		return fmt.Errorf("userProfileRepo.GetUserNotesCount: %s", err.Error())
	}
	if count == userNotesLimit {
		return models.UserLimitReachedErr
	}
	err = userProfileRepo.CreateUserNote(note)
	if err != nil {
		return fmt.Errorf("CreateUserNote: %s", err.Error())
	}
	return nil
}

func (t *TezTracker) DeleteUserNote(accountID string, text string) error {
	return t.repoProvider.GetUserProfile().DeleteUserNote(accountID, text)
}
