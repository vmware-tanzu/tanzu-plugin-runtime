// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVSphereNetworksHandlerFunc turns a function with the right signature into a get v sphere networks handler
type GetVSphereNetworksHandlerFunc func(GetVSphereNetworksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVSphereNetworksHandlerFunc) Handle(params GetVSphereNetworksParams) middleware.Responder {
	return fn(params)
}

// GetVSphereNetworksHandler interface for that can handle valid get v sphere networks params
type GetVSphereNetworksHandler interface {
	Handle(GetVSphereNetworksParams) middleware.Responder
}

// NewGetVSphereNetworks creates a new http.Handler for the get v sphere networks operation
func NewGetVSphereNetworks(ctx *middleware.Context, handler GetVSphereNetworksHandler) *GetVSphereNetworks {
	return &GetVSphereNetworks{Context: ctx, Handler: handler}
}

/*
GetVSphereNetworks swagger:route GET /api/providers/vsphere/networks vsphere getVSphereNetworks

Retrieve networks associated with the datacenter in vSphere
*/
type GetVSphereNetworks struct {
	Context *middleware.Context
	Handler GetVSphereNetworksHandler
}

func (o *GetVSphereNetworks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVSphereNetworksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
