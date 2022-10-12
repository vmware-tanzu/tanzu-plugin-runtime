// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateDockerRegionalClusterHandlerFunc turns a function with the right signature into a create docker regional cluster handler
type CreateDockerRegionalClusterHandlerFunc func(CreateDockerRegionalClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateDockerRegionalClusterHandlerFunc) Handle(params CreateDockerRegionalClusterParams) middleware.Responder {
	return fn(params)
}

// CreateDockerRegionalClusterHandler interface for that can handle valid create docker regional cluster params
type CreateDockerRegionalClusterHandler interface {
	Handle(CreateDockerRegionalClusterParams) middleware.Responder
}

// NewCreateDockerRegionalCluster creates a new http.Handler for the create docker regional cluster operation
func NewCreateDockerRegionalCluster(ctx *middleware.Context, handler CreateDockerRegionalClusterHandler) *CreateDockerRegionalCluster {
	return &CreateDockerRegionalCluster{Context: ctx, Handler: handler}
}

/*
CreateDockerRegionalCluster swagger:route POST /api/providers/docker/create docker createDockerRegionalCluster

Create Docker regional cluster
*/
type CreateDockerRegionalCluster struct {
	Context *middleware.Context
	Handler CreateDockerRegionalClusterHandler
}

func (o *CreateDockerRegionalCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateDockerRegionalClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
