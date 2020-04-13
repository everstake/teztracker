// Code generated by go-swagger; DO NOT EDIT.

package operations_list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetDoubleEndorsementsListURL generates an URL for the get double endorsements list operation
type GetDoubleEndorsementsListURL struct {
	Network  string
	Platform string

	AccountID   []string
	BlockID     []string
	BlockLevel  []int64
	Limit       *int64
	Offset      *int64
	OperationID []string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDoubleEndorsementsListURL) WithBasePath(bp string) *GetDoubleEndorsementsListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDoubleEndorsementsListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetDoubleEndorsementsListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/data/{platform}/{network}/double_endorsement"

	network := o.Network
	if network != "" {
		_path = strings.Replace(_path, "{network}", network, -1)
	} else {
		return nil, errors.New("network is required on GetDoubleEndorsementsListURL")
	}

	platform := o.Platform
	if platform != "" {
		_path = strings.Replace(_path, "{platform}", platform, -1)
	} else {
		return nil, errors.New("platform is required on GetDoubleEndorsementsListURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var accountIDIR []string
	for _, accountIDI := range o.AccountID {
		accountIDIS := accountIDI
		if accountIDIS != "" {
			accountIDIR = append(accountIDIR, accountIDIS)
		}
	}

	accountID := swag.JoinByFormat(accountIDIR, "multi")

	for _, qsv := range accountID {
		qs.Add("account_id", qsv)
	}

	var blockIDIR []string
	for _, blockIDI := range o.BlockID {
		blockIDIS := blockIDI
		if blockIDIS != "" {
			blockIDIR = append(blockIDIR, blockIDIS)
		}
	}

	blockID := swag.JoinByFormat(blockIDIR, "multi")

	for _, qsv := range blockID {
		qs.Add("block_id", qsv)
	}

	var blockLevelIR []string
	for _, blockLevelI := range o.BlockLevel {
		blockLevelIS := swag.FormatInt64(blockLevelI)
		if blockLevelIS != "" {
			blockLevelIR = append(blockLevelIR, blockLevelIS)
		}
	}

	blockLevel := swag.JoinByFormat(blockLevelIR, "multi")

	for _, qsv := range blockLevel {
		qs.Add("block_level", qsv)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt64(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	var operationIDIR []string
	for _, operationIDI := range o.OperationID {
		operationIDIS := operationIDI
		if operationIDIS != "" {
			operationIDIR = append(operationIDIR, operationIDIS)
		}
	}

	operationID := swag.JoinByFormat(operationIDIR, "multi")

	for _, qsv := range operationID {
		qs.Add("operation_id", qsv)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetDoubleEndorsementsListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetDoubleEndorsementsListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetDoubleEndorsementsListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetDoubleEndorsementsListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetDoubleEndorsementsListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetDoubleEndorsementsListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
