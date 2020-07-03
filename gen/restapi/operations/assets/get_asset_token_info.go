// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAssetTokenInfoHandlerFunc turns a function with the right signature into a get asset token info handler
type GetAssetTokenInfoHandlerFunc func(GetAssetTokenInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAssetTokenInfoHandlerFunc) Handle(params GetAssetTokenInfoParams) middleware.Responder {
	return fn(params)
}

// GetAssetTokenInfoHandler interface for that can handle valid get asset token info params
type GetAssetTokenInfoHandler interface {
	Handle(GetAssetTokenInfoParams) middleware.Responder
}

// NewGetAssetTokenInfo creates a new http.Handler for the get asset token info operation
func NewGetAssetTokenInfo(ctx *middleware.Context, handler GetAssetTokenInfoHandler) *GetAssetTokenInfo {
	return &GetAssetTokenInfo{Context: ctx, Handler: handler}
}

/*GetAssetTokenInfo swagger:route GET /v2/data/{network}/assets/{asset_id} Assets getAssetTokenInfo

GetAssetTokenInfo get asset token info API

*/
type GetAssetTokenInfo struct {
	Context *middleware.Context
	Handler GetAssetTokenInfoHandler
}

func (o *GetAssetTokenInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAssetTokenInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
