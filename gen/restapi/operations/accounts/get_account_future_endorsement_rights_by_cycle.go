// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountFutureEndorsementRightsByCycleHandlerFunc turns a function with the right signature into a get account future endorsement rights by cycle handler
type GetAccountFutureEndorsementRightsByCycleHandlerFunc func(GetAccountFutureEndorsementRightsByCycleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountFutureEndorsementRightsByCycleHandlerFunc) Handle(params GetAccountFutureEndorsementRightsByCycleParams) middleware.Responder {
	return fn(params)
}

// GetAccountFutureEndorsementRightsByCycleHandler interface for that can handle valid get account future endorsement rights by cycle params
type GetAccountFutureEndorsementRightsByCycleHandler interface {
	Handle(GetAccountFutureEndorsementRightsByCycleParams) middleware.Responder
}

// NewGetAccountFutureEndorsementRightsByCycle creates a new http.Handler for the get account future endorsement rights by cycle operation
func NewGetAccountFutureEndorsementRightsByCycle(ctx *middleware.Context, handler GetAccountFutureEndorsementRightsByCycleHandler) *GetAccountFutureEndorsementRightsByCycle {
	return &GetAccountFutureEndorsementRightsByCycle{Context: ctx, Handler: handler}
}

/*GetAccountFutureEndorsementRightsByCycle swagger:route GET /v2/data/{platform}/{network}/accounts/endorsing/{accountId}/future_endorsement_rights/{cycleId} Accounts getAccountFutureEndorsementRightsByCycle

GetAccountFutureEndorsementRightsByCycle get account future endorsement rights by cycle API

*/
type GetAccountFutureEndorsementRightsByCycle struct {
	Context *middleware.Context
	Handler GetAccountFutureEndorsementRightsByCycleHandler
}

func (o *GetAccountFutureEndorsementRightsByCycle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountFutureEndorsementRightsByCycleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
