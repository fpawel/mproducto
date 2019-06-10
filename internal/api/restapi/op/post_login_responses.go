// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/fpawel/mproducto/internal/api/model"
)

// PostLoginOKCode is the HTTP code returned for type PostLoginOK
const PostLoginOKCode int = 200

/*PostLoginOK The api token of the user

swagger:response postLoginOK
*/
type PostLoginOK struct {

	/*
	  In: Body
	*/
	Payload *model.APIKey `json:"body,omitempty"`
}

// NewPostLoginOK creates PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {

	return &PostLoginOK{}
}

// WithPayload adds the payload to the post login o k response
func (o *PostLoginOK) WithPayload(payload *model.APIKey) *PostLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login o k response
func (o *PostLoginOK) SetPayload(payload *model.APIKey) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostLoginDefault Error

swagger:response postLoginDefault
*/
type PostLoginDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewPostLoginDefault creates PostLoginDefault with default headers values
func NewPostLoginDefault(code int) *PostLoginDefault {
	if code <= 0 {
		code = 500
	}

	return &PostLoginDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post login default response
func (o *PostLoginDefault) WithStatusCode(code int) *PostLoginDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post login default response
func (o *PostLoginDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post login default response
func (o *PostLoginDefault) WithPayload(payload *model.Error) *PostLoginDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post login default response
func (o *PostLoginDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostLoginDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
