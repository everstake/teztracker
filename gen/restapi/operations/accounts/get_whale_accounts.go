// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetWhaleAccountsHandlerFunc turns a function with the right signature into a get whale accounts handler
type GetWhaleAccountsHandlerFunc func(GetWhaleAccountsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWhaleAccountsHandlerFunc) Handle(params GetWhaleAccountsParams) middleware.Responder {
	return fn(params)
}

// GetWhaleAccountsHandler interface for that can handle valid get whale accounts params
type GetWhaleAccountsHandler interface {
	Handle(GetWhaleAccountsParams) middleware.Responder
}

// NewGetWhaleAccounts creates a new http.Handler for the get whale accounts operation
func NewGetWhaleAccounts(ctx *middleware.Context, handler GetWhaleAccountsHandler) *GetWhaleAccounts {
	return &GetWhaleAccounts{Context: ctx, Handler: handler}
}

/*GetWhaleAccounts swagger:route GET /v2/data/{platform}/{network}/whale/accounts Accounts getWhaleAccounts

GetWhaleAccounts get whale accounts API

*/
type GetWhaleAccounts struct {
	Context *middleware.Context
	Handler GetWhaleAccountsHandler
}

func (o *GetWhaleAccounts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWhaleAccountsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
