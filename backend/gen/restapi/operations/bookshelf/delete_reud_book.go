// Code generated by go-swagger; DO NOT EDIT.

package bookshelf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteReudBookHandlerFunc turns a function with the right signature into a delete reud book handler
type DeleteReudBookHandlerFunc func(DeleteReudBookParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteReudBookHandlerFunc) Handle(params DeleteReudBookParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteReudBookHandler interface for that can handle valid delete reud book params
type DeleteReudBookHandler interface {
	Handle(DeleteReudBookParams, interface{}) middleware.Responder
}

// NewDeleteReudBook creates a new http.Handler for the delete reud book operation
func NewDeleteReudBook(ctx *middleware.Context, handler DeleteReudBookHandler) *DeleteReudBook {
	return &DeleteReudBook{Context: ctx, Handler: handler}
}

/*DeleteReudBook swagger:route DELETE /bookshelf/{bookId} bookshelf deleteReudBook

delete book by id

amazonの投稿した本を削除する機能

*/
type DeleteReudBook struct {
	Context *middleware.Context
	Handler DeleteReudBookHandler
}

func (o *DeleteReudBook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteReudBookParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
