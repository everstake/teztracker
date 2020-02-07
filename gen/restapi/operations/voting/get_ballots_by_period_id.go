// Code generated by go-swagger; DO NOT EDIT.

package voting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBallotsByPeriodIDHandlerFunc turns a function with the right signature into a get ballots by period ID handler
type GetBallotsByPeriodIDHandlerFunc func(GetBallotsByPeriodIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBallotsByPeriodIDHandlerFunc) Handle(params GetBallotsByPeriodIDParams) middleware.Responder {
	return fn(params)
}

// GetBallotsByPeriodIDHandler interface for that can handle valid get ballots by period ID params
type GetBallotsByPeriodIDHandler interface {
	Handle(GetBallotsByPeriodIDParams) middleware.Responder
}

// NewGetBallotsByPeriodID creates a new http.Handler for the get ballots by period ID operation
func NewGetBallotsByPeriodID(ctx *middleware.Context, handler GetBallotsByPeriodIDHandler) *GetBallotsByPeriodID {
	return &GetBallotsByPeriodID{Context: ctx, Handler: handler}
}

/*GetBallotsByPeriodID swagger:route GET /v2/ballots/{id} Voting getBallotsByPeriodId

GetBallotsByPeriodID get ballots by period ID API

*/
type GetBallotsByPeriodID struct {
	Context *middleware.Context
	Handler GetBallotsByPeriodIDHandler
}

func (o *GetBallotsByPeriodID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBallotsByPeriodIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
