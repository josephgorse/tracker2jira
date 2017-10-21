// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostProjectHandlerFunc turns a function with the right signature into a post project handler
type PostProjectHandlerFunc func(PostProjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProjectHandlerFunc) Handle(params PostProjectParams) middleware.Responder {
	return fn(params)
}

// PostProjectHandler interface for that can handle valid post project params
type PostProjectHandler interface {
	Handle(PostProjectParams) middleware.Responder
}

// NewPostProject creates a new http.Handler for the post project operation
func NewPostProject(ctx *middleware.Context, handler PostProjectHandler) *PostProject {
	return &PostProject{Context: ctx, Handler: handler}
}

/*PostProject swagger:route POST /projects projects postProject

Adds a project configuration

Post a new project config

*/
type PostProject struct {
	Context *middleware.Context
	Handler PostProjectHandler
}

func (o *PostProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostProjectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
