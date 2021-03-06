// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountSecurityDepositListHandlerFunc turns a function with the right signature into a get account security deposit list handler
type GetAccountSecurityDepositListHandlerFunc func(GetAccountSecurityDepositListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountSecurityDepositListHandlerFunc) Handle(params GetAccountSecurityDepositListParams) middleware.Responder {
	return fn(params)
}

// GetAccountSecurityDepositListHandler interface for that can handle valid get account security deposit list params
type GetAccountSecurityDepositListHandler interface {
	Handle(GetAccountSecurityDepositListParams) middleware.Responder
}

// NewGetAccountSecurityDepositList creates a new http.Handler for the get account security deposit list operation
func NewGetAccountSecurityDepositList(ctx *middleware.Context, handler GetAccountSecurityDepositListHandler) *GetAccountSecurityDepositList {
	return &GetAccountSecurityDepositList{Context: ctx, Handler: handler}
}

/*GetAccountSecurityDepositList swagger:route GET /v2/data/{platform}/{network}/accounts/security_deposit/{accountId}/future Accounts getAccountSecurityDepositList

GetAccountSecurityDepositList get account security deposit list API

*/
type GetAccountSecurityDepositList struct {
	Context *middleware.Context
	Handler GetAccountSecurityDepositListHandler
}

func (o *GetAccountSecurityDepositList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountSecurityDepositListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
