// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetWhaleMoversHandlerFunc turns a function with the right signature into a get whale movers handler
type GetWhaleMoversHandlerFunc func(GetWhaleMoversParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWhaleMoversHandlerFunc) Handle(params GetWhaleMoversParams) middleware.Responder {
	return fn(params)
}

// GetWhaleMoversHandler interface for that can handle valid get whale movers params
type GetWhaleMoversHandler interface {
	Handle(GetWhaleMoversParams) middleware.Responder
}

// NewGetWhaleMovers creates a new http.Handler for the get whale movers operation
func NewGetWhaleMovers(ctx *middleware.Context, handler GetWhaleMoversHandler) *GetWhaleMovers {
	return &GetWhaleMovers{Context: ctx, Handler: handler}
}

/*GetWhaleMovers swagger:route GET /v2/data/{platform}/{network}/whale/movers Accounts getWhaleMovers

GetWhaleMovers get whale movers API

*/
type GetWhaleMovers struct {
	Context *middleware.Context
	Handler GetWhaleMoversHandler
}

func (o *GetWhaleMovers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWhaleMoversParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}