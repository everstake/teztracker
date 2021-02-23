package api

import (
	"github.com/everstake/teztracker/api/render"
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/gen/restapi/operations/profile"
	"github.com/everstake/teztracker/gen/restapi/operations/voting"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos"
	"github.com/everstake/teztracker/services"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

type getUserProfileHandler struct {
	provider DbProvider
}

func (h *getUserProfileHandler) Handle(params profile.GetUserProfileParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return voting.NewGetPeriodBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserProfileBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserProfileInternalServerError()
	}
	return profile.NewGetUserProfileOK().WithPayload(&genModels.UserProfile{
		Email:    &user.Email,
		Username: &user.Username,
		Verified: &user.Verified,
	})
}

type updateUserProfileHandler struct {
	provider DbProvider
}

func (h *updateUserProfileHandler) Handle(params profile.UpdateProfileParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewUpdateProfileBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewUpdateProfileBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewUpdateProfileInternalServerError()
	}

	user.Email = params.Data.Email
	user.Username = params.Data.Username
	err = service.UpdateUser(user)
	if err != nil {
		logrus.Errorf("failed update user profile: %s", err.Error())
		return profile.NewUpdateProfileInternalServerError()
	}
	return profile.NewUpdateProfileOK()
}

type getUserAddressesHandler struct {
	provider DbProvider
}

func (h *getUserAddressesHandler) Handle(params profile.GetUserAddressesParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}
	addresses, err := service.GetUserAddresses(user.AccountID)
	if err != nil {
		logrus.Errorf("failed get user addresses: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}
	return profile.NewGetUserAddressesOK().WithPayload(render.UserAddresses(addresses))
}

type createOrUpdateUserAddressHandler struct {
	provider DbProvider
}

func (h *createOrUpdateUserAddressHandler) Handle(params profile.CreateOrUpdateUserAddressParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}

	err = service.CreateOrUpdateUserAddress(models.UserAddress{
		AccountID:           user.AccountID,
		Address:             *params.Data.Address,
		DelegationsEnabled:  *params.Data.DelegationsEnabled,
		InTransfersEnabled:  *params.Data.InTransfersEnabled,
		OutTransfersEnabled: *params.Data.OutTransfersEnabled,
	})
	if err == models.UserLimitReachedErr || err == models.AccountNotFoundErr {
		return profile.NewCreateOrUpdateUserAddressBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed create or update user address: %s", err.Error())
		return profile.NewCreateOrUpdateUserAddressInternalServerError()
	}
	return profile.NewCreateOrUpdateUserAddressOK()
}

type deleteUserAddressHandler struct {
	provider DbProvider
}

func (h *deleteUserAddressHandler) Handle(params profile.DeleteUserAddressParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}

	err = service.DeleteUserAddress(user.AccountID, params.Data.Address)
	if err != nil {
		logrus.Errorf("failed delete user address: %s", err.Error())
		return profile.NewDeleteUserAddressInternalServerError()
	}
	return profile.NewDeleteUserAddressOK()
}

type getUserNotesHandler struct {
	provider DbProvider
}

func (h *getUserNotesHandler) Handle(params profile.GetUserNotesParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}

	notes, err := service.GetUserNotes(user.AccountID)
	if err != nil {
		logrus.Errorf("failed get user notes: %s", err.Error())
		return profile.NewGetUserNotesInternalServerError()
	}
	return profile.NewGetUserNotesOK().WithPayload(render.UserNotesWithBalance(notes))
}

type createOrUpdateUserNoteHandler struct {
	provider DbProvider
}

func (h *createOrUpdateUserNoteHandler) Handle(params profile.CreateOrUpdateNoteParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}

	err = service.CreateOrUpdateUserNote(models.UserNote{
		AccountID:   user.AccountID,
		Address:     params.Data.Address,
		Alias:       params.Data.Alias,
		Tag:         params.Data.Tag,
		Description: params.Data.Description,
	})
	if err == models.UserLimitReachedErr || err == models.AccountNotFoundErr {
		return profile.NewCreateOrUpdateNoteBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user notes: %s", err.Error())
		return profile.NewCreateOrUpdateNoteInternalServerError()
	}
	return profile.NewCreateOrUpdateNoteOK()
}

type deleteUserNoteHandler struct {
	provider DbProvider
}

func (h *deleteUserNoteHandler) Handle(params profile.DeleteUserNoteParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}

	err = service.DeleteUserNote(user.AccountID, params.Data.Address)
	if err != nil {
		logrus.Errorf("failed get user notes: %s", err.Error())
		return profile.NewDeleteUserNoteInternalServerError()
	}
	return profile.NewDeleteUserNoteOK()
}

type verifyEmailHandler struct {
	provider DbProvider
}

func (h *verifyEmailHandler) Handle(params profile.VerifyEmailParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	user, err := service.GetOrCreateUser(params.Address)
	if err == models.AccountNotFoundErr {
		logrus.Warnf("account not found")
		return profile.NewGetUserAddressesBadRequest()
	}
	if err != nil {
		logrus.Errorf("failed get user profile: %s", err.Error())
		return profile.NewGetUserAddressesInternalServerError()
	}

	err = service.EmailVerification(user.AccountID)
	if err != nil {
		logrus.Errorf("failed get user notes: %s", err.Error())
		return profile.NewVerifyEmailInternalServerError()
	}
	return profile.NewVerifyEmailOK()
}

type verifyEmailTokenHandler struct {
	provider DbProvider
}

func (h *verifyEmailTokenHandler) Handle(params profile.VerifyEmailTokenParams) middleware.Responder {
	db, err := h.provider.GetDb(models.NetworkMain)
	if err != nil {
		return profile.NewGetUserAddressesBadRequest()
	}
	service := services.New(repos.New(db), models.NetworkMain)
	err = service.EmailTokenVerification(params.Data.Token)
	if err != nil {
		logrus.Errorf("failed get user notes: %s", err.Error())
		return profile.NewVerifyEmailTokenInternalServerError()
	}
	return profile.NewVerifyEmailTokenOK()
}
