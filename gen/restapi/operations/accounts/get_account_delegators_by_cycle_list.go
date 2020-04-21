// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountDelegatorsByCycleListHandlerFunc turns a function with the right signature into a get account delegators by cycle list handler
type GetAccountDelegatorsByCycleListHandlerFunc func(GetAccountDelegatorsByCycleListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountDelegatorsByCycleListHandlerFunc) Handle(params GetAccountDelegatorsByCycleListParams) middleware.Responder {
	return fn(params)
}

// GetAccountDelegatorsByCycleListHandler interface for that can handle valid get account delegators by cycle list params
type GetAccountDelegatorsByCycleListHandler interface {
	Handle(GetAccountDelegatorsByCycleListParams) middleware.Responder
}

// NewGetAccountDelegatorsByCycleList creates a new http.Handler for the get account delegators by cycle list operation
func NewGetAccountDelegatorsByCycleList(ctx *middleware.Context, handler GetAccountDelegatorsByCycleListHandler) *GetAccountDelegatorsByCycleList {
	return &GetAccountDelegatorsByCycleList{Context: ctx, Handler: handler}
}

/*GetAccountDelegatorsByCycleList swagger:route GET /v2/data/{platform}/{network}/accounts/{accountId}/delegators/{cycleId} Accounts getAccountDelegatorsByCycleList

GetAccountDelegatorsByCycleList get account delegators by cycle list API

*/
type GetAccountDelegatorsByCycleList struct {
	Context *middleware.Context
	Handler GetAccountDelegatorsByCycleListHandler
}

func (o *GetAccountDelegatorsByCycleList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountDelegatorsByCycleListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
