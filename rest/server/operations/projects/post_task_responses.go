// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostTaskCreatedCode is the HTTP code returned for type PostTaskCreated
const PostTaskCreatedCode int = 201

/*PostTaskCreated Created

swagger:response postTaskCreated
*/
type PostTaskCreated struct {
}

// NewPostTaskCreated creates PostTaskCreated with default headers values
func NewPostTaskCreated() *PostTaskCreated {
	return &PostTaskCreated{}
}

// WriteResponse to the client
func (o *PostTaskCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
}

// PostTaskBadRequestCode is the HTTP code returned for type PostTaskBadRequest
const PostTaskBadRequestCode int = 400

/*PostTaskBadRequest Bad Request

swagger:response postTaskBadRequest
*/
type PostTaskBadRequest struct {
}

// NewPostTaskBadRequest creates PostTaskBadRequest with default headers values
func NewPostTaskBadRequest() *PostTaskBadRequest {
	return &PostTaskBadRequest{}
}

// WriteResponse to the client
func (o *PostTaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}
