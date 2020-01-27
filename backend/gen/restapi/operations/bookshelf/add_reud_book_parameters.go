// Code generated by go-swagger; DO NOT EDIT.

package bookshelf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "backend/gen/models"
)

// NewAddReudBookParams creates a new AddReudBookParams object
// no default values defined in spec.
func NewAddReudBookParams() AddReudBookParams {

	return AddReudBookParams{}
}

// AddReudBookParams contains all the bound params for the add reud book operation
// typically these are obtained from a http.Request
//
// swagger:parameters addReudBook
type AddReudBookParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*本の情報は本棚に登録するために必要である。
	  Required: true
	  In: body
	*/
	Body *models.Book
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddReudBookParams() beforehand.
func (o *AddReudBookParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Book
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
