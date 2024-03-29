// Code generated by go-swagger; DO NOT EDIT.

package n_f_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetNFTContractTokensListHandlerFunc turns a function with the right signature into a get n f t contract tokens list handler
type GetNFTContractTokensListHandlerFunc func(GetNFTContractTokensListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetNFTContractTokensListHandlerFunc) Handle(params GetNFTContractTokensListParams) middleware.Responder {
	return fn(params)
}

// GetNFTContractTokensListHandler interface for that can handle valid get n f t contract tokens list params
type GetNFTContractTokensListHandler interface {
	Handle(GetNFTContractTokensListParams) middleware.Responder
}

// NewGetNFTContractTokensList creates a new http.Handler for the get n f t contract tokens list operation
func NewGetNFTContractTokensList(ctx *middleware.Context, handler GetNFTContractTokensListHandler) *GetNFTContractTokensList {
	return &GetNFTContractTokensList{Context: ctx, Handler: handler}
}

/*GetNFTContractTokensList swagger:route GET /v2/data/{network}/nft_contracts/{contract_id}/tokens NFT getNFTContractTokensList

GetNFTContractTokensList get n f t contract tokens list API

*/
type GetNFTContractTokensList struct {
	Context *middleware.Context
	Handler GetNFTContractTokensListHandler
}

func (o *GetNFTContractTokensList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetNFTContractTokensListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
