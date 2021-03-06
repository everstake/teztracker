// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountsListParams creates a new GetAccountsListParams object
// with the default values initialized.
func NewGetAccountsListParams() GetAccountsListParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(20)

		offsetDefault = int64(0)
	)

	return GetAccountsListParams{
		Limit: &limitDefault,

		Offset: &offsetDefault,
	}
}

// GetAccountsListParams contains all the bound params for the get accounts list operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAccountsList
type GetAccountsListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Not used
	  In: query
	  Collection Format: multi
	*/
	AccountDelegate []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	AccountID []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	AccountManager []string
	/*
	  In: query
	*/
	AfterID *string
	/*
	  In: query
	  Collection Format: multi
	*/
	BlockID []string
	/*
	  In: query
	  Collection Format: multi
	*/
	BlockLevel []int64
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	BlockNetid []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	BlockProtocol []string
	/*favorites accounts
	  In: query
	  Collection Format: multi
	*/
	Favorites []string
	/*
	  Maximum: 500
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*Not used
	  Required: true
	  In: path
	*/
	Network string
	/*Offset
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Offset *int64
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	OperationDestination []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	OperationID []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	OperationKind []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	OperationParticipant []string
	/*Not used
	  In: query
	  Collection Format: multi
	*/
	OperationSource []string
	/*Not used
	  In: query
	*/
	Order *string
	/*Not used
	  Required: true
	  In: path
	*/
	Platform string
	/*Not used
	  In: query
	*/
	SortBy *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAccountsListParams() beforehand.
func (o *GetAccountsListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccountDelegate, qhkAccountDelegate, _ := qs.GetOK("account_delegate")
	if err := o.bindAccountDelegate(qAccountDelegate, qhkAccountDelegate, route.Formats); err != nil {
		res = append(res, err)
	}

	qAccountID, qhkAccountID, _ := qs.GetOK("account_id")
	if err := o.bindAccountID(qAccountID, qhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	qAccountManager, qhkAccountManager, _ := qs.GetOK("account_manager")
	if err := o.bindAccountManager(qAccountManager, qhkAccountManager, route.Formats); err != nil {
		res = append(res, err)
	}

	qAfterID, qhkAfterID, _ := qs.GetOK("after_id")
	if err := o.bindAfterID(qAfterID, qhkAfterID, route.Formats); err != nil {
		res = append(res, err)
	}

	qBlockID, qhkBlockID, _ := qs.GetOK("block_id")
	if err := o.bindBlockID(qBlockID, qhkBlockID, route.Formats); err != nil {
		res = append(res, err)
	}

	qBlockLevel, qhkBlockLevel, _ := qs.GetOK("block_level")
	if err := o.bindBlockLevel(qBlockLevel, qhkBlockLevel, route.Formats); err != nil {
		res = append(res, err)
	}

	qBlockNetid, qhkBlockNetid, _ := qs.GetOK("block_netid")
	if err := o.bindBlockNetid(qBlockNetid, qhkBlockNetid, route.Formats); err != nil {
		res = append(res, err)
	}

	qBlockProtocol, qhkBlockProtocol, _ := qs.GetOK("block_protocol")
	if err := o.bindBlockProtocol(qBlockProtocol, qhkBlockProtocol, route.Formats); err != nil {
		res = append(res, err)
	}

	qFavorites, qhkFavorites, _ := qs.GetOK("favorites")
	if err := o.bindFavorites(qFavorites, qhkFavorites, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	rNetwork, rhkNetwork, _ := route.Params.GetOK("network")
	if err := o.bindNetwork(rNetwork, rhkNetwork, route.Formats); err != nil {
		res = append(res, err)
	}

	qOffset, qhkOffset, _ := qs.GetOK("offset")
	if err := o.bindOffset(qOffset, qhkOffset, route.Formats); err != nil {
		res = append(res, err)
	}

	qOperationDestination, qhkOperationDestination, _ := qs.GetOK("operation_destination")
	if err := o.bindOperationDestination(qOperationDestination, qhkOperationDestination, route.Formats); err != nil {
		res = append(res, err)
	}

	qOperationID, qhkOperationID, _ := qs.GetOK("operation_id")
	if err := o.bindOperationID(qOperationID, qhkOperationID, route.Formats); err != nil {
		res = append(res, err)
	}

	qOperationKind, qhkOperationKind, _ := qs.GetOK("operation_kind")
	if err := o.bindOperationKind(qOperationKind, qhkOperationKind, route.Formats); err != nil {
		res = append(res, err)
	}

	qOperationParticipant, qhkOperationParticipant, _ := qs.GetOK("operation_participant")
	if err := o.bindOperationParticipant(qOperationParticipant, qhkOperationParticipant, route.Formats); err != nil {
		res = append(res, err)
	}

	qOperationSource, qhkOperationSource, _ := qs.GetOK("operation_source")
	if err := o.bindOperationSource(qOperationSource, qhkOperationSource, route.Formats); err != nil {
		res = append(res, err)
	}

	qOrder, qhkOrder, _ := qs.GetOK("order")
	if err := o.bindOrder(qOrder, qhkOrder, route.Formats); err != nil {
		res = append(res, err)
	}

	rPlatform, rhkPlatform, _ := route.Params.GetOK("platform")
	if err := o.bindPlatform(rPlatform, rhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortBy, qhkSortBy, _ := qs.GetOK("sort_by")
	if err := o.bindSortBy(qSortBy, qhkSortBy, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAccountDelegate binds and validates array parameter AccountDelegate from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindAccountDelegate(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	accountDelegateIC := rawData

	if len(accountDelegateIC) == 0 {
		return nil
	}

	var accountDelegateIR []string
	for _, accountDelegateIV := range accountDelegateIC {
		accountDelegateI := accountDelegateIV

		accountDelegateIR = append(accountDelegateIR, accountDelegateI)
	}

	o.AccountDelegate = accountDelegateIR

	return nil
}

// bindAccountID binds and validates array parameter AccountID from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	accountIDIC := rawData

	if len(accountIDIC) == 0 {
		return nil
	}

	var accountIDIR []string
	for _, accountIDIV := range accountIDIC {
		accountIDI := accountIDIV

		accountIDIR = append(accountIDIR, accountIDI)
	}

	o.AccountID = accountIDIR

	return nil
}

// bindAccountManager binds and validates array parameter AccountManager from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindAccountManager(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	accountManagerIC := rawData

	if len(accountManagerIC) == 0 {
		return nil
	}

	var accountManagerIR []string
	for _, accountManagerIV := range accountManagerIC {
		accountManagerI := accountManagerIV

		accountManagerIR = append(accountManagerIR, accountManagerI)
	}

	o.AccountManager = accountManagerIR

	return nil
}

// bindAfterID binds and validates parameter AfterID from query.
func (o *GetAccountsListParams) bindAfterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AfterID = &raw

	return nil
}

// bindBlockID binds and validates array parameter BlockID from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindBlockID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	blockIDIC := rawData

	if len(blockIDIC) == 0 {
		return nil
	}

	var blockIDIR []string
	for _, blockIDIV := range blockIDIC {
		blockIDI := blockIDIV

		blockIDIR = append(blockIDIR, blockIDI)
	}

	o.BlockID = blockIDIR

	return nil
}

// bindBlockLevel binds and validates array parameter BlockLevel from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindBlockLevel(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	blockLevelIC := rawData

	if len(blockLevelIC) == 0 {
		return nil
	}

	var blockLevelIR []int64
	for i, blockLevelIV := range blockLevelIC {
		// items.Format: "int64"
		blockLevelI, err := swag.ConvertInt64(blockLevelIV)
		if err != nil {
			return errors.InvalidType(fmt.Sprintf("%s.%v", "block_level", i), "query", "int64", blockLevelI)
		}

		blockLevelIR = append(blockLevelIR, blockLevelI)
	}

	o.BlockLevel = blockLevelIR

	return nil
}

// bindBlockNetid binds and validates array parameter BlockNetid from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindBlockNetid(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	blockNetidIC := rawData

	if len(blockNetidIC) == 0 {
		return nil
	}

	var blockNetidIR []string
	for _, blockNetidIV := range blockNetidIC {
		blockNetidI := blockNetidIV

		blockNetidIR = append(blockNetidIR, blockNetidI)
	}

	o.BlockNetid = blockNetidIR

	return nil
}

// bindBlockProtocol binds and validates array parameter BlockProtocol from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindBlockProtocol(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	blockProtocolIC := rawData

	if len(blockProtocolIC) == 0 {
		return nil
	}

	var blockProtocolIR []string
	for _, blockProtocolIV := range blockProtocolIC {
		blockProtocolI := blockProtocolIV

		blockProtocolIR = append(blockProtocolIR, blockProtocolI)
	}

	o.BlockProtocol = blockProtocolIR

	return nil
}

// bindFavorites binds and validates array parameter Favorites from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindFavorites(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	favoritesIC := rawData

	if len(favoritesIC) == 0 {
		return nil
	}

	var favoritesIR []string
	for _, favoritesIV := range favoritesIC {
		favoritesI := favoritesIV

		favoritesIR = append(favoritesIR, favoritesI)
	}

	o.Favorites = favoritesIR

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetAccountsListParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAccountsListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *GetAccountsListParams) validateLimit(formats strfmt.Registry) error {

	if err := validate.MinimumInt("limit", "query", int64(*o.Limit), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("limit", "query", int64(*o.Limit), 500, false); err != nil {
		return err
	}

	return nil
}

// bindNetwork binds and validates parameter Network from path.
func (o *GetAccountsListParams) bindNetwork(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Network = raw

	return nil
}

// bindOffset binds and validates parameter Offset from query.
func (o *GetAccountsListParams) bindOffset(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAccountsListParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("offset", "query", "int64", raw)
	}
	o.Offset = &value

	if err := o.validateOffset(formats); err != nil {
		return err
	}

	return nil
}

// validateOffset carries on validations for parameter Offset
func (o *GetAccountsListParams) validateOffset(formats strfmt.Registry) error {

	if err := validate.MinimumInt("offset", "query", int64(*o.Offset), 0, false); err != nil {
		return err
	}

	return nil
}

// bindOperationDestination binds and validates array parameter OperationDestination from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindOperationDestination(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	operationDestinationIC := rawData

	if len(operationDestinationIC) == 0 {
		return nil
	}

	var operationDestinationIR []string
	for _, operationDestinationIV := range operationDestinationIC {
		operationDestinationI := operationDestinationIV

		operationDestinationIR = append(operationDestinationIR, operationDestinationI)
	}

	o.OperationDestination = operationDestinationIR

	return nil
}

// bindOperationID binds and validates array parameter OperationID from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindOperationID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	operationIDIC := rawData

	if len(operationIDIC) == 0 {
		return nil
	}

	var operationIDIR []string
	for _, operationIDIV := range operationIDIC {
		operationIDI := operationIDIV

		operationIDIR = append(operationIDIR, operationIDI)
	}

	o.OperationID = operationIDIR

	return nil
}

// bindOperationKind binds and validates array parameter OperationKind from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindOperationKind(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	operationKindIC := rawData

	if len(operationKindIC) == 0 {
		return nil
	}

	var operationKindIR []string
	for _, operationKindIV := range operationKindIC {
		operationKindI := operationKindIV

		operationKindIR = append(operationKindIR, operationKindI)
	}

	o.OperationKind = operationKindIR

	return nil
}

// bindOperationParticipant binds and validates array parameter OperationParticipant from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindOperationParticipant(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	operationParticipantIC := rawData

	if len(operationParticipantIC) == 0 {
		return nil
	}

	var operationParticipantIR []string
	for _, operationParticipantIV := range operationParticipantIC {
		operationParticipantI := operationParticipantIV

		operationParticipantIR = append(operationParticipantIR, operationParticipantI)
	}

	o.OperationParticipant = operationParticipantIR

	return nil
}

// bindOperationSource binds and validates array parameter OperationSource from query.
//
// Arrays are parsed according to CollectionFormat: "multi" (defaults to "csv" when empty).
func (o *GetAccountsListParams) bindOperationSource(rawData []string, hasKey bool, formats strfmt.Registry) error {

	// CollectionFormat: multi
	operationSourceIC := rawData

	if len(operationSourceIC) == 0 {
		return nil
	}

	var operationSourceIR []string
	for _, operationSourceIV := range operationSourceIC {
		operationSourceI := operationSourceIV

		operationSourceIR = append(operationSourceIR, operationSourceI)
	}

	o.OperationSource = operationSourceIR

	return nil
}

// bindOrder binds and validates parameter Order from query.
func (o *GetAccountsListParams) bindOrder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Order = &raw

	return nil
}

// bindPlatform binds and validates parameter Platform from path.
func (o *GetAccountsListParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Platform = raw

	return nil
}

// bindSortBy binds and validates parameter SortBy from query.
func (o *GetAccountsListParams) bindSortBy(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SortBy = &raw

	return nil
}
