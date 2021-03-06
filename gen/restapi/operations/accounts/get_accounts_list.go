// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountsListHandlerFunc turns a function with the right signature into a get accounts list handler
type GetAccountsListHandlerFunc func(GetAccountsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountsListHandlerFunc) Handle(params GetAccountsListParams) middleware.Responder {
	return fn(params)
}

// GetAccountsListHandler interface for that can handle valid get accounts list params
type GetAccountsListHandler interface {
	Handle(GetAccountsListParams) middleware.Responder
}

// NewGetAccountsList creates a new http.Handler for the get accounts list operation
func NewGetAccountsList(ctx *middleware.Context, handler GetAccountsListHandler) *GetAccountsList {
	return &GetAccountsList{Context: ctx, Handler: handler}
}

/*GetAccountsList swagger:route GET /v2/data/{platform}/{network}/accounts Accounts getAccountsList

GetAccountsList get accounts list API

*/
type GetAccountsList struct {
	Context *middleware.Context
	Handler GetAccountsListHandler
}

func (o *GetAccountsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountsListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
