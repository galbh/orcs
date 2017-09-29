package crb_web_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"../../work/nazgul/copyrepo/copyrepo-web-api/src/main/resources/models"
)

// CreateCopyInRepoCreatedCode is the HTTP code returned for type CreateCopyInRepoCreated
const CreateCopyInRepoCreatedCode int = 201

/*CreateCopyInRepoCreated Successfully created a copy of the data and an entry in the Metadata-DB

swagger:response createCopyInRepoCreated
*/
type CreateCopyInRepoCreated struct {

	/*Copy URL - http://<crb-endpoint>/copies/<copyId>
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateCopyInRepoCreated creates CreateCopyInRepoCreated with default headers values
func NewCreateCopyInRepoCreated() *CreateCopyInRepoCreated {
	return &CreateCopyInRepoCreated{}
}

// WithPayload adds the payload to the create copy in repo created response
func (o *CreateCopyInRepoCreated) WithPayload(payload string) *CreateCopyInRepoCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy in repo created response
func (o *CreateCopyInRepoCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyInRepoCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// CreateCopyInRepoNotFoundCode is the HTTP code returned for type CreateCopyInRepoNotFound
const CreateCopyInRepoNotFoundCode int = 404

/*CreateCopyInRepoNotFound Repository repoId instance is not configured

swagger:response createCopyInRepoNotFound
*/
type CreateCopyInRepoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyInRepoNotFound creates CreateCopyInRepoNotFound with default headers values
func NewCreateCopyInRepoNotFound() *CreateCopyInRepoNotFound {
	return &CreateCopyInRepoNotFound{}
}

// WithPayload adds the payload to the create copy in repo not found response
func (o *CreateCopyInRepoNotFound) WithPayload(payload *models.Error) *CreateCopyInRepoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy in repo not found response
func (o *CreateCopyInRepoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyInRepoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCopyInRepoConflictCode is the HTTP code returned for type CreateCopyInRepoConflict
const CreateCopyInRepoConflictCode int = 409

/*CreateCopyInRepoConflict A copy instance with the same copyId exists. Either use different copyId or delete existing copyId instance before retrying

swagger:response createCopyInRepoConflict
*/
type CreateCopyInRepoConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyInRepoConflict creates CreateCopyInRepoConflict with default headers values
func NewCreateCopyInRepoConflict() *CreateCopyInRepoConflict {
	return &CreateCopyInRepoConflict{}
}

// WithPayload adds the payload to the create copy in repo conflict response
func (o *CreateCopyInRepoConflict) WithPayload(payload *models.Error) *CreateCopyInRepoConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy in repo conflict response
func (o *CreateCopyInRepoConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyInRepoConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCopyInRepoInternalServerErrorCode is the HTTP code returned for type CreateCopyInRepoInternalServerError
const CreateCopyInRepoInternalServerErrorCode int = 500

/*CreateCopyInRepoInternalServerError Internal error. Data could not be copied to copy-repo or copy instance metadata could not be created. If copy instance metadata can't be created, copy instance will be deleted

swagger:response createCopyInRepoInternalServerError
*/
type CreateCopyInRepoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyInRepoInternalServerError creates CreateCopyInRepoInternalServerError with default headers values
func NewCreateCopyInRepoInternalServerError() *CreateCopyInRepoInternalServerError {
	return &CreateCopyInRepoInternalServerError{}
}

// WithPayload adds the payload to the create copy in repo internal server error response
func (o *CreateCopyInRepoInternalServerError) WithPayload(payload *models.Error) *CreateCopyInRepoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy in repo internal server error response
func (o *CreateCopyInRepoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyInRepoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateCopyInRepoDefault Unexpected error

swagger:response createCopyInRepoDefault
*/
type CreateCopyInRepoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateCopyInRepoDefault creates CreateCopyInRepoDefault with default headers values
func NewCreateCopyInRepoDefault(code int) *CreateCopyInRepoDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateCopyInRepoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create copy in repo default response
func (o *CreateCopyInRepoDefault) WithStatusCode(code int) *CreateCopyInRepoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create copy in repo default response
func (o *CreateCopyInRepoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create copy in repo default response
func (o *CreateCopyInRepoDefault) WithPayload(payload *models.Error) *CreateCopyInRepoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create copy in repo default response
func (o *CreateCopyInRepoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCopyInRepoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
