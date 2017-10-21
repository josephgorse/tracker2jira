// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostTaskHandlerFunc turns a function with the right signature into a post task handler
type PostTaskHandlerFunc func(PostTaskParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTaskHandlerFunc) Handle(params PostTaskParams) middleware.Responder {
	return fn(params)
}

// PostTaskHandler interface for that can handle valid post task params
type PostTaskHandler interface {
	Handle(PostTaskParams) middleware.Responder
}

// NewPostTask creates a new http.Handler for the post task operation
func NewPostTask(ctx *middleware.Context, handler PostTaskHandler) *PostTask {
	return &PostTask{Context: ctx, Handler: handler}
}

/*PostTask swagger:route POST /tasks projects postTask

Adds a task configuration

Post a new task config

*/
type PostTask struct {
	Context *middleware.Context
	Handler PostTaskHandler
}

func (o *PostTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTaskParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
