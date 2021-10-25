// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNFTContractOwnershipHandlerFunc turns a function with the right signature into a get n f t contract ownership handler
type GetNFTContractOwnershipHandlerFunc func(GetNFTContractOwnershipParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNFTContractOwnershipHandlerFunc) Handle(params GetNFTContractOwnershipParams) middleware.Responder {
	return fn(params)
}

// GetNFTContractOwnershipHandler interface for that can handle valid get n f t contract ownership params
type GetNFTContractOwnershipHandler interface {
	Handle(GetNFTContractOwnershipParams) middleware.Responder
}

// NewGetNFTContractOwnership creates a new http.Handler for the get n f t contract ownership operation
func NewGetNFTContractOwnership(ctx *middleware.Context, handler GetNFTContractOwnershipHandler) *GetNFTContractOwnership {
	return &GetNFTContractOwnership{Context: ctx, Handler: handler}
}

/*GetNFTContractOwnership swagger:route GET /v2/data/{network}/nft_contracts/{contract_id}/ownership NFT getNFTContractOwnership

GetNFTContractOwnership get n f t contract ownership API

*/
type GetNFTContractOwnership struct {
	Context *middleware.Context
	Handler GetNFTContractOwnershipHandler
}

func (o *GetNFTContractOwnership) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNFTContractOwnershipParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
