// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetNFTContractOperationsChartURL generates an URL for the get n f t contract operations chart operation
type GetNFTContractOperationsChartURL struct {
	ContractID string
	Network    string

	From   int64
	Period string
	To     int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNFTContractOperationsChartURL) WithBasePath(bp string) *GetNFTContractOperationsChartURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetNFTContractOperationsChartURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetNFTContractOperationsChartURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/data/{network}/nft_contracts/{contract_id}/operations/chart"

	contractID := o.ContractID
	if contractID != "" {
		_path = strings.Replace(_path, "{contract_id}", contractID, -1)
	} else {
		return nil, errors.New("contractId is required on GetNFTContractOperationsChartURL")
	}

	network := o.Network
	if network != "" {
		_path = strings.Replace(_path, "{network}", network, -1)
	} else {
		return nil, errors.New("network is required on GetNFTContractOperationsChartURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	fromQ := swag.FormatInt64(o.From)
	if fromQ != "" {
		qs.Set("from", fromQ)
	}

	periodQ := o.Period
	if periodQ != "" {
		qs.Set("period", periodQ)
	}

	toQ := swag.FormatInt64(o.To)
	if toQ != "" {
		qs.Set("to", toQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetNFTContractOperationsChartURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetNFTContractOperationsChartURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetNFTContractOperationsChartURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetNFTContractOperationsChartURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetNFTContractOperationsChartURL")
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
func (o *GetNFTContractOperationsChartURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
