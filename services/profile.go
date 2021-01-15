package services

import (
	"fmt"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/services/mailer"
	"github.com/guregu/null"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"regexp"
)

const (
	userAddressesLimit = 50
	userNotesLimit     = 50
	emailTokenLen      = 30
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

func (t *TezTracker) UpdateUser(newUser models.User) error {
	profileRepo := t.repoProvider.GetUserProfile()
	user, found, err := profileRepo.FindUserByAccount(newUser.AccountID)
	if err != nil {
		return fmt.Errorf("profileRepo.GetUserProfile: %s", err.Error())
	}
	if !found {
		return fmt.Errorf("profileRepo.GetUserProfile: not found")
	}
	if newUser.Email != user.Email && newUser.Email != "" {
		if !isEmailValid(newUser.Email) {
			return fmt.Errorf("profileRepo.GetUserProfile: invalid address")
		}
		newUser.Verified = false
	}
	return profileRepo.UpdateUser(newUser)
}

func (t *TezTracker) EmailVerification(accountID string) error {
	profileRepo := t.repoProvider.GetUserProfile()
	user, found, err := profileRepo.FindUserByAccount(accountID)
	if err != nil {
		return fmt.Errorf("profileRepo.GetUserProfile: %s", err.Error())
	}
	if !found {
		return fmt.Errorf("profileRepo.GetUserProfile: not found")
	}
	if user.Email == "" {
		return fmt.Errorf("empty user email")
	}
	_, found, err = profileRepo.GetEmailVerification(accountID, user.Email)
	if err != nil {
		return fmt.Errorf("repoProvider.GetEmailVerification: %s", err.Error())
	}
	if found {
		return nil
	}
	token := randomStr(emailTokenLen)
	err = profileRepo.CreateEmailToken(models.EmailVerification{
		AccountID: accountID,
		Email:     user.Email,
		Token:     token,
	})
	if err != nil {
		return fmt.Errorf("repoProvider.CreateEmailToken: %s", err.Error())
	}
	return nil
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
	userNote, found, err := userProfileRepo.FindUserNote(note.AccountID, note.Text)
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
	count, err := userProfileRepo.GetUserNotesCount(userNote.AccountID)
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

func (t *TezTracker) SendNewVerifications(pusher mailer.Mail) error {
	userRepo := t.repoProvider.GetUserProfile()
	verifications, err := userRepo.GetEmailVerifications(false, nil)
	if err != nil {
		return fmt.Errorf("userRepo.GetEmailVerifications: %s", err.Error())
	}
	for _, verification := range verifications {
		err = pusher.Send(verification.Email, mailer.VerificationMsg, map[string]string{
			"token": verification.Token,
		})
		if err != nil {
			log.Error("SendNewVerifications: pusher.Send: %s", err.Error())
			continue
		}
		verification.Sent = true
		err = userRepo.UpdateEmailVerifications(verification)
		if err != nil {
			log.Error("SendNewVerifications: userRepo.UpdateEmailVerifications: %s", err.Error())
			continue
		}
	}
	return nil
}

func (t *TezTracker) EmailTokenVerification(token string) error {
	profileRepo := t.repoProvider.GetUserProfile()
	verifications, err := profileRepo.GetEmailVerifications(true, []string{token})
	if err != nil {
		return fmt.Errorf("userRepo.GetEmailVerifications: %s", err.Error())
	}
	if len(verifications) == 0 {
		return fmt.Errorf("verification not found")
	}
	verification := verifications[0]
	user, found, err := profileRepo.FindUserByAccount(verification.AccountID)
	if err != nil {
		return fmt.Errorf("profileRepo.GetUserProfile: %s", err.Error())
	}
	if !found {
		return fmt.Errorf("profileRepo.GetUserProfile: not found")
	}
	user.Email = verification.Email
	user.Verified = true
	err = profileRepo.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("profileRepo.UpdateUser: %s", err.Error())
	}
	verification.Verified = true
	err = profileRepo.UpdateEmailVerifications(verification)
	if err != nil {
		return fmt.Errorf("profileRepo.UpdateEmailVerifications: %s", err.Error())
	}
	return nil
}

func randomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
