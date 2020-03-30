// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountBakingListHandlerFunc turns a function with the right signature into a get account baking list handler
type GetAccountBakingListHandlerFunc func(GetAccountBakingListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountBakingListHandlerFunc) Handle(params GetAccountBakingListParams) middleware.Responder {
	return fn(params)
}

// GetAccountBakingListHandler interface for that can handle valid get account baking list params
type GetAccountBakingListHandler interface {
	Handle(GetAccountBakingListParams) middleware.Responder
}

// NewGetAccountBakingList creates a new http.Handler for the get account baking list operation
func NewGetAccountBakingList(ctx *middleware.Context, handler GetAccountBakingListHandler) *GetAccountBakingList {
	return &GetAccountBakingList{Context: ctx, Handler: handler}
}

/*GetAccountBakingList swagger:route GET /v2/data/{platform}/{network}/accounts/baking/{accountId} Accounts getAccountBakingList

GetAccountBakingList get account baking list API

*/
type GetAccountBakingList struct {
	Context *middleware.Context
	Handler GetAccountBakingListHandler
}

func (o *GetAccountBakingList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountBakingListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
