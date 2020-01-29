// Code generated by go-swagger; DO NOT EDIT.

package bookshelf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddReudBookHandlerFunc turns a function with the right signature into a add reud book handler
type AddReudBookHandlerFunc func(AddReudBookParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddReudBookHandlerFunc) Handle(params AddReudBookParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddReudBookHandler interface for that can handle valid add reud book params
type AddReudBookHandler interface {
	Handle(AddReudBookParams, interface{}) middleware.Responder
}

// NewAddReudBook creates a new http.Handler for the add reud book operation
func NewAddReudBook(ctx *middleware.Context, handler AddReudBookHandler) *AddReudBook {
	return &AddReudBook{Context: ctx, Handler: handler}
}

/*AddReudBook swagger:route POST /bookshelf bookshelf addReudBook

Add a new book to bookshelf

amazonの本を投稿できる機能(amazonにあるもの)

*/
type AddReudBook struct {
	Context *middleware.Context
	Handler AddReudBookHandler
}

func (o *AddReudBook) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddReudBookParams()

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
