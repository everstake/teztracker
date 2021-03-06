// Code generated by go-swagger; DO NOT EDIT.

package voting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVotingRollsParams creates a new GetVotingRollsParams object
// with the default values initialized.
func NewGetVotingRollsParams() *GetVotingRollsParams {
	var ()
	return &GetVotingRollsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVotingRollsParamsWithTimeout creates a new GetVotingRollsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVotingRollsParamsWithTimeout(timeout time.Duration) *GetVotingRollsParams {
	var ()
	return &GetVotingRollsParams{

		timeout: timeout,
	}
}

// NewGetVotingRollsParamsWithContext creates a new GetVotingRollsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVotingRollsParamsWithContext(ctx context.Context) *GetVotingRollsParams {
	var ()
	return &GetVotingRollsParams{

		Context: ctx,
	}
}

// NewGetVotingRollsParamsWithHTTPClient creates a new GetVotingRollsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVotingRollsParamsWithHTTPClient(client *http.Client) *GetVotingRollsParams {
	var ()
	return &GetVotingRollsParams{
		HTTPClient: client,
	}
}

/*GetVotingRollsParams contains all the parameters to send to the API endpoint
for the get voting rolls operation typically these are written to a http.Request
*/
type GetVotingRollsParams struct {

	/*Block*/
	Block string
	/*Network*/
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voting rolls params
func (o *GetVotingRollsParams) WithTimeout(timeout time.Duration) *GetVotingRollsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voting rolls params
func (o *GetVotingRollsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voting rolls params
func (o *GetVotingRollsParams) WithContext(ctx context.Context) *GetVotingRollsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voting rolls params
func (o *GetVotingRollsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voting rolls params
func (o *GetVotingRollsParams) WithHTTPClient(client *http.Client) *GetVotingRollsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voting rolls params
func (o *GetVotingRollsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlock adds the block to the get voting rolls params
func (o *GetVotingRollsParams) WithBlock(block string) *GetVotingRollsParams {
	o.SetBlock(block)
	return o
}

// SetBlock adds the block to the get voting rolls params
func (o *GetVotingRollsParams) SetBlock(block string) {
	o.Block = block
}

// WithNetwork adds the network to the get voting rolls params
func (o *GetVotingRollsParams) WithNetwork(network string) *GetVotingRollsParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the get voting rolls params
func (o *GetVotingRollsParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *GetVotingRollsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param block
	if err := r.SetPathParam("block", o.Block); err != nil {
		return err
	}

	// path param network
	if err := r.SetPathParam("network", o.Network); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
