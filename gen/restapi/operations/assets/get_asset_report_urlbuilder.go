// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetAssetReportURL generates an URL for the get asset report operation
type GetAssetReportURL struct {
	AssetID  string
	Network  string
	Platform string

	From          int64
	OperationType []string
	To            int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAssetReportURL) WithBasePath(bp string) *GetAssetReportURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAssetReportURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAssetReportURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/data/{platform}/{network}/assets/{assetId}/report"

	assetID := o.AssetID
	if assetID != "" {
		_path = strings.Replace(_path, "{assetId}", assetID, -1)
	} else {
		return nil, errors.New("assetId is required on GetAssetReportURL")
	}

	network := o.Network
	if network != "" {
		_path = strings.Replace(_path, "{network}", network, -1)
	} else {
		return nil, errors.New("network is required on GetAssetReportURL")
	}

	platform := o.Platform
	if platform != "" {
		_path = strings.Replace(_path, "{platform}", platform, -1)
	} else {
		return nil, errors.New("platform is required on GetAssetReportURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	fromQ := swag.FormatInt64(o.From)
	if fromQ != "" {
		qs.Set("from", fromQ)
	}

	var operationTypeIR []string
	for _, operationTypeI := range o.OperationType {
		operationTypeIS := operationTypeI
		if operationTypeIS != "" {
			operationTypeIR = append(operationTypeIR, operationTypeIS)
		}
	}

	operationType := swag.JoinByFormat(operationTypeIR, "multi")

	for _, qsv := range operationType {
		qs.Add("operation_type", qsv)
	}

	toQ := swag.FormatInt64(o.To)
	if toQ != "" {
		qs.Set("to", toQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAssetReportURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAssetReportURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAssetReportURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAssetReportURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAssetReportURL")
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
func (o *GetAssetReportURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
