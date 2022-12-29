// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigParamsOKCode is the HTTP code returned for type GetConfigParamsOK
const GetConfigParamsOKCode int = 200

/*
GetConfigParamsOK OK

swagger:response getConfigParamsOK
*/
type GetConfigParamsOK struct {

	/*
	  In: Body
	*/
	Payload *models.OperParams `json:"body,omitempty"`
}

// NewGetConfigParamsOK creates GetConfigParamsOK with default headers values
func NewGetConfigParamsOK() *GetConfigParamsOK {

	return &GetConfigParamsOK{}
}

// WithPayload adds the payload to the get config params o k response
func (o *GetConfigParamsOK) WithPayload(payload *models.OperParams) *GetConfigParamsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params o k response
func (o *GetConfigParamsOK) SetPayload(payload *models.OperParams) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsNoContentCode is the HTTP code returned for type GetConfigParamsNoContent
const GetConfigParamsNoContentCode int = 204

/*
GetConfigParamsNoContent OK

swagger:response getConfigParamsNoContent
*/
type GetConfigParamsNoContent struct {
}

// NewGetConfigParamsNoContent creates GetConfigParamsNoContent with default headers values
func NewGetConfigParamsNoContent() *GetConfigParamsNoContent {

	return &GetConfigParamsNoContent{}
}

// WriteResponse to the client
func (o *GetConfigParamsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// GetConfigParamsBadRequestCode is the HTTP code returned for type GetConfigParamsBadRequest
const GetConfigParamsBadRequestCode int = 400

/*
GetConfigParamsBadRequest Malformed arguments for API call

swagger:response getConfigParamsBadRequest
*/
type GetConfigParamsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsBadRequest creates GetConfigParamsBadRequest with default headers values
func NewGetConfigParamsBadRequest() *GetConfigParamsBadRequest {

	return &GetConfigParamsBadRequest{}
}

// WithPayload adds the payload to the get config params bad request response
func (o *GetConfigParamsBadRequest) WithPayload(payload *models.Error) *GetConfigParamsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params bad request response
func (o *GetConfigParamsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsUnauthorizedCode is the HTTP code returned for type GetConfigParamsUnauthorized
const GetConfigParamsUnauthorizedCode int = 401

/*
GetConfigParamsUnauthorized Invalid authentication credentials

swagger:response getConfigParamsUnauthorized
*/
type GetConfigParamsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsUnauthorized creates GetConfigParamsUnauthorized with default headers values
func NewGetConfigParamsUnauthorized() *GetConfigParamsUnauthorized {

	return &GetConfigParamsUnauthorized{}
}

// WithPayload adds the payload to the get config params unauthorized response
func (o *GetConfigParamsUnauthorized) WithPayload(payload *models.Error) *GetConfigParamsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params unauthorized response
func (o *GetConfigParamsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsForbiddenCode is the HTTP code returned for type GetConfigParamsForbidden
const GetConfigParamsForbiddenCode int = 403

/*
GetConfigParamsForbidden Capacity insufficient

swagger:response getConfigParamsForbidden
*/
type GetConfigParamsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsForbidden creates GetConfigParamsForbidden with default headers values
func NewGetConfigParamsForbidden() *GetConfigParamsForbidden {

	return &GetConfigParamsForbidden{}
}

// WithPayload adds the payload to the get config params forbidden response
func (o *GetConfigParamsForbidden) WithPayload(payload *models.Error) *GetConfigParamsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params forbidden response
func (o *GetConfigParamsForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsNotFoundCode is the HTTP code returned for type GetConfigParamsNotFound
const GetConfigParamsNotFoundCode int = 404

/*
GetConfigParamsNotFound Resource not found

swagger:response getConfigParamsNotFound
*/
type GetConfigParamsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsNotFound creates GetConfigParamsNotFound with default headers values
func NewGetConfigParamsNotFound() *GetConfigParamsNotFound {

	return &GetConfigParamsNotFound{}
}

// WithPayload adds the payload to the get config params not found response
func (o *GetConfigParamsNotFound) WithPayload(payload *models.Error) *GetConfigParamsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params not found response
func (o *GetConfigParamsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsConflictCode is the HTTP code returned for type GetConfigParamsConflict
const GetConfigParamsConflictCode int = 409

/*
GetConfigParamsConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response getConfigParamsConflict
*/
type GetConfigParamsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsConflict creates GetConfigParamsConflict with default headers values
func NewGetConfigParamsConflict() *GetConfigParamsConflict {

	return &GetConfigParamsConflict{}
}

// WithPayload adds the payload to the get config params conflict response
func (o *GetConfigParamsConflict) WithPayload(payload *models.Error) *GetConfigParamsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params conflict response
func (o *GetConfigParamsConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsInternalServerErrorCode is the HTTP code returned for type GetConfigParamsInternalServerError
const GetConfigParamsInternalServerErrorCode int = 500

/*
GetConfigParamsInternalServerError Internal service error

swagger:response getConfigParamsInternalServerError
*/
type GetConfigParamsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsInternalServerError creates GetConfigParamsInternalServerError with default headers values
func NewGetConfigParamsInternalServerError() *GetConfigParamsInternalServerError {

	return &GetConfigParamsInternalServerError{}
}

// WithPayload adds the payload to the get config params internal server error response
func (o *GetConfigParamsInternalServerError) WithPayload(payload *models.Error) *GetConfigParamsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params internal server error response
func (o *GetConfigParamsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigParamsServiceUnavailableCode is the HTTP code returned for type GetConfigParamsServiceUnavailable
const GetConfigParamsServiceUnavailableCode int = 503

/*
GetConfigParamsServiceUnavailable Maintanence mode

swagger:response getConfigParamsServiceUnavailable
*/
type GetConfigParamsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigParamsServiceUnavailable creates GetConfigParamsServiceUnavailable with default headers values
func NewGetConfigParamsServiceUnavailable() *GetConfigParamsServiceUnavailable {

	return &GetConfigParamsServiceUnavailable{}
}

// WithPayload adds the payload to the get config params service unavailable response
func (o *GetConfigParamsServiceUnavailable) WithPayload(payload *models.Error) *GetConfigParamsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config params service unavailable response
func (o *GetConfigParamsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigParamsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
