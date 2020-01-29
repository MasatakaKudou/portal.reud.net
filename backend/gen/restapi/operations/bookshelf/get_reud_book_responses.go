// Code generated by go-swagger; DO NOT EDIT.

package bookshelf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "backend/gen/models"
)

// GetReudBookOKCode is the HTTP code returned for type GetReudBookOK
const GetReudBookOKCode int = 200

/*GetReudBookOK 本の情報、IDはreudが本を同サイトから消す際に使用される。

swagger:response getReudBookOK
*/
type GetReudBookOK struct {

	/*
	  In: Body
	*/
	Payload []*models.StoredBook `json:"body,omitempty"`
}

// NewGetReudBookOK creates GetReudBookOK with default headers values
func NewGetReudBookOK() *GetReudBookOK {

	return &GetReudBookOK{}
}

// WithPayload adds the payload to the get reud book o k response
func (o *GetReudBookOK) WithPayload(payload []*models.StoredBook) *GetReudBookOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get reud book o k response
func (o *GetReudBookOK) SetPayload(payload []*models.StoredBook) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReudBookOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.StoredBook, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetReudBookInternalServerErrorCode is the HTTP code returned for type GetReudBookInternalServerError
const GetReudBookInternalServerErrorCode int = 500

/*GetReudBookInternalServerError Internal Server Error

swagger:response getReudBookInternalServerError
*/
type GetReudBookInternalServerError struct {
}

// NewGetReudBookInternalServerError creates GetReudBookInternalServerError with default headers values
func NewGetReudBookInternalServerError() *GetReudBookInternalServerError {

	return &GetReudBookInternalServerError{}
}

// WriteResponse to the client
func (o *GetReudBookInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
