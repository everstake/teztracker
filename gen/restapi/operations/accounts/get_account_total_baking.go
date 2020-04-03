// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountTotalBakingHandlerFunc turns a function with the right signature into a get account total baking handler
type GetAccountTotalBakingHandlerFunc func(GetAccountTotalBakingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountTotalBakingHandlerFunc) Handle(params GetAccountTotalBakingParams) middleware.Responder {
	return fn(params)
}

// GetAccountTotalBakingHandler interface for that can handle valid get account total baking params
type GetAccountTotalBakingHandler interface {
	Handle(GetAccountTotalBakingParams) middleware.Responder
}

// NewGetAccountTotalBaking creates a new http.Handler for the get account total baking operation
func NewGetAccountTotalBaking(ctx *middleware.Context, handler GetAccountTotalBakingHandler) *GetAccountTotalBaking {
	return &GetAccountTotalBaking{Context: ctx, Handler: handler}
}

/*GetAccountTotalBaking swagger:route GET /v2/data/{platform}/{network}/accounts/baking/{accountId}/total Accounts getAccountTotalBaking

GetAccountTotalBaking get account total baking API

*/
type GetAccountTotalBaking struct {
	Context *middleware.Context
	Handler GetAccountTotalBakingHandler
}

func (o *GetAccountTotalBaking) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountTotalBakingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}