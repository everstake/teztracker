// Code generated by go-swagger; DO NOT EDIT.

package app_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetBakerChartInfoHandlerFunc turns a function with the right signature into a get baker chart info handler
type GetBakerChartInfoHandlerFunc func(GetBakerChartInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBakerChartInfoHandlerFunc) Handle(params GetBakerChartInfoParams) middleware.Responder {
	return fn(params)
}

// GetBakerChartInfoHandler interface for that can handle valid get baker chart info params
type GetBakerChartInfoHandler interface {
	Handle(GetBakerChartInfoParams) middleware.Responder
}

// NewGetBakerChartInfo creates a new http.Handler for the get baker chart info operation
func NewGetBakerChartInfo(ctx *middleware.Context, handler GetBakerChartInfoHandler) *GetBakerChartInfo {
	return &GetBakerChartInfo{Context: ctx, Handler: handler}
}

/*GetBakerChartInfo swagger:route GET /v2/data/{platform}/{network}/charts/bakers App Info getBakerChartInfo

GetBakerChartInfo get baker chart info API

*/
type GetBakerChartInfo struct {
	Context *middleware.Context
	Handler GetBakerChartInfoHandler
}

func (o *GetBakerChartInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetBakerChartInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
