// Copyright (c) [2017] Dell Inc. or its subsidiaries. All Rights Reserved.
package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"../../work/nazgul/copyrepo/copyrepo-web-api/src/main/resources/models"
)

// ListRepositoryInstancesOKCode is the HTTP code returned for type ListRepositoryInstancesOK
const ListRepositoryInstancesOKCode int = 200

/*ListRepositoryInstancesOK Successfully listed Repo instances managed by crb

swagger:response listRepositoryInstancesOK
*/
type ListRepositoryInstancesOK struct {

	/*repoId is currently used to access RepositoryInstance.
	  In: Body
	*/
	Payload []*models.RepositoryInstance `json:"body,omitempty"`
}

// NewListRepositoryInstancesOK creates ListRepositoryInstancesOK with default headers values
func NewListRepositoryInstancesOK() *ListRepositoryInstancesOK {
	return &ListRepositoryInstancesOK{}
}

// WithPayload adds the payload to the list repository instances o k response
func (o *ListRepositoryInstancesOK) WithPayload(payload []*models.RepositoryInstance) *ListRepositoryInstancesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list repository instances o k response
func (o *ListRepositoryInstancesOK) SetPayload(payload []*models.RepositoryInstance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRepositoryInstancesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.RepositoryInstance, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ListRepositoryInstancesNoContentCode is the HTTP code returned for type ListRepositoryInstancesNoContent
const ListRepositoryInstancesNoContentCode int = 204

/*ListRepositoryInstancesNoContent No repositories content in Metadata-DB

swagger:response listRepositoryInstancesNoContent
*/
type ListRepositoryInstancesNoContent struct {
}

// NewListRepositoryInstancesNoContent creates ListRepositoryInstancesNoContent with default headers values
func NewListRepositoryInstancesNoContent() *ListRepositoryInstancesNoContent {
	return &ListRepositoryInstancesNoContent{}
}

// WriteResponse to the client
func (o *ListRepositoryInstancesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// ListRepositoryInstancesInternalServerErrorCode is the HTTP code returned for type ListRepositoryInstancesInternalServerError
const ListRepositoryInstancesInternalServerErrorCode int = 500

/*ListRepositoryInstancesInternalServerError Internal error -- most likely Metadata-DB not accessible

swagger:response listRepositoryInstancesInternalServerError
*/
type ListRepositoryInstancesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListRepositoryInstancesInternalServerError creates ListRepositoryInstancesInternalServerError with default headers values
func NewListRepositoryInstancesInternalServerError() *ListRepositoryInstancesInternalServerError {
	return &ListRepositoryInstancesInternalServerError{}
}

// WithPayload adds the payload to the list repository instances internal server error response
func (o *ListRepositoryInstancesInternalServerError) WithPayload(payload *models.Error) *ListRepositoryInstancesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list repository instances internal server error response
func (o *ListRepositoryInstancesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRepositoryInstancesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListRepositoryInstancesDefault Unexpected error

swagger:response listRepositoryInstancesDefault
*/
type ListRepositoryInstancesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListRepositoryInstancesDefault creates ListRepositoryInstancesDefault with default headers values
func NewListRepositoryInstancesDefault(code int) *ListRepositoryInstancesDefault {
	if code <= 0 {
		code = 500
	}

	return &ListRepositoryInstancesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list repository instances default response
func (o *ListRepositoryInstancesDefault) WithStatusCode(code int) *ListRepositoryInstancesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list repository instances default response
func (o *ListRepositoryInstancesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list repository instances default response
func (o *ListRepositoryInstancesDefault) WithPayload(payload *models.Error) *ListRepositoryInstancesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list repository instances default response
func (o *ListRepositoryInstancesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListRepositoryInstancesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
