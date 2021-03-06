// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAssetReportHandlerFunc turns a function with the right signature into a get asset report handler
type GetAssetReportHandlerFunc func(GetAssetReportParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAssetReportHandlerFunc) Handle(params GetAssetReportParams) middleware.Responder {
	return fn(params)
}

// GetAssetReportHandler interface for that can handle valid get asset report params
type GetAssetReportHandler interface {
	Handle(GetAssetReportParams) middleware.Responder
}

// NewGetAssetReport creates a new http.Handler for the get asset report operation
func NewGetAssetReport(ctx *middleware.Context, handler GetAssetReportHandler) *GetAssetReport {
	return &GetAssetReport{Context: ctx, Handler: handler}
}

/*GetAssetReport swagger:route GET /v2/data/{platform}/{network}/assets/{assetId}/report Assets getAssetReport

GetAssetReport get asset report API

*/
type GetAssetReport struct {
	Context *middleware.Context
	Handler GetAssetReportHandler
}

func (o *GetAssetReport) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAssetReportParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
