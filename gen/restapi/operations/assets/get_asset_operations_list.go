// Code generated by go-swagger; DO NOT EDIT.

package assets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAssetOperationsListHandlerFunc turns a function with the right signature into a get asset operations list handler
type GetAssetOperationsListHandlerFunc func(GetAssetOperationsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAssetOperationsListHandlerFunc) Handle(params GetAssetOperationsListParams) middleware.Responder {
	return fn(params)
}

// GetAssetOperationsListHandler interface for that can handle valid get asset operations list params
type GetAssetOperationsListHandler interface {
	Handle(GetAssetOperationsListParams) middleware.Responder
}

// NewGetAssetOperationsList creates a new http.Handler for the get asset operations list operation
func NewGetAssetOperationsList(ctx *middleware.Context, handler GetAssetOperationsListHandler) *GetAssetOperationsList {
	return &GetAssetOperationsList{Context: ctx, Handler: handler}
}

/*GetAssetOperationsList swagger:route GET /v2/data/{network}/assets/operations Assets getAssetOperationsList

GetAssetOperationsList get asset operations list API

*/
type GetAssetOperationsList struct {
	Context *middleware.Context
	Handler GetAssetOperationsListHandler
}

func (o *GetAssetOperationsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAssetOperationsListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
