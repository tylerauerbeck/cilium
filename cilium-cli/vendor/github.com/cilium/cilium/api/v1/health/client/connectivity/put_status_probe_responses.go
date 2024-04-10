// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package connectivity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/health/models"
)

// PutStatusProbeReader is a Reader for the PutStatusProbe structure.
type PutStatusProbeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutStatusProbeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutStatusProbeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPutStatusProbeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutStatusProbeFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /status/probe] PutStatusProbe", response, response.Code())
	}
}

// NewPutStatusProbeOK creates a PutStatusProbeOK with default headers values
func NewPutStatusProbeOK() *PutStatusProbeOK {
	return &PutStatusProbeOK{}
}

/*
PutStatusProbeOK describes a response with status code 200, with default header values.

Success
*/
type PutStatusProbeOK struct {
	Payload *models.HealthStatusResponse
}

// IsSuccess returns true when this put status probe o k response has a 2xx status code
func (o *PutStatusProbeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put status probe o k response has a 3xx status code
func (o *PutStatusProbeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put status probe o k response has a 4xx status code
func (o *PutStatusProbeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put status probe o k response has a 5xx status code
func (o *PutStatusProbeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put status probe o k response a status code equal to that given
func (o *PutStatusProbeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put status probe o k response
func (o *PutStatusProbeOK) Code() int {
	return 200
}

func (o *PutStatusProbeOK) Error() string {
	return fmt.Sprintf("[PUT /status/probe][%d] putStatusProbeOK  %+v", 200, o.Payload)
}

func (o *PutStatusProbeOK) String() string {
	return fmt.Sprintf("[PUT /status/probe][%d] putStatusProbeOK  %+v", 200, o.Payload)
}

func (o *PutStatusProbeOK) GetPayload() *models.HealthStatusResponse {
	return o.Payload
}

func (o *PutStatusProbeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutStatusProbeForbidden creates a PutStatusProbeForbidden with default headers values
func NewPutStatusProbeForbidden() *PutStatusProbeForbidden {
	return &PutStatusProbeForbidden{}
}

/*
PutStatusProbeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutStatusProbeForbidden struct {
}

// IsSuccess returns true when this put status probe forbidden response has a 2xx status code
func (o *PutStatusProbeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put status probe forbidden response has a 3xx status code
func (o *PutStatusProbeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put status probe forbidden response has a 4xx status code
func (o *PutStatusProbeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put status probe forbidden response has a 5xx status code
func (o *PutStatusProbeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put status probe forbidden response a status code equal to that given
func (o *PutStatusProbeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put status probe forbidden response
func (o *PutStatusProbeForbidden) Code() int {
	return 403
}

func (o *PutStatusProbeForbidden) Error() string {
	return fmt.Sprintf("[PUT /status/probe][%d] putStatusProbeForbidden ", 403)
}

func (o *PutStatusProbeForbidden) String() string {
	return fmt.Sprintf("[PUT /status/probe][%d] putStatusProbeForbidden ", 403)
}

func (o *PutStatusProbeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutStatusProbeFailed creates a PutStatusProbeFailed with default headers values
func NewPutStatusProbeFailed() *PutStatusProbeFailed {
	return &PutStatusProbeFailed{}
}

/*
PutStatusProbeFailed describes a response with status code 500, with default header values.

Internal error occurred while conducting connectivity probe
*/
type PutStatusProbeFailed struct {
	Payload models.Error
}

// IsSuccess returns true when this put status probe failed response has a 2xx status code
func (o *PutStatusProbeFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put status probe failed response has a 3xx status code
func (o *PutStatusProbeFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put status probe failed response has a 4xx status code
func (o *PutStatusProbeFailed) IsClientError() bool {
	return false
}

// IsServerError returns true when this put status probe failed response has a 5xx status code
func (o *PutStatusProbeFailed) IsServerError() bool {
	return true
}

// IsCode returns true when this put status probe failed response a status code equal to that given
func (o *PutStatusProbeFailed) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put status probe failed response
func (o *PutStatusProbeFailed) Code() int {
	return 500
}

func (o *PutStatusProbeFailed) Error() string {
	return fmt.Sprintf("[PUT /status/probe][%d] putStatusProbeFailed  %+v", 500, o.Payload)
}

func (o *PutStatusProbeFailed) String() string {
	return fmt.Sprintf("[PUT /status/probe][%d] putStatusProbeFailed  %+v", 500, o.Payload)
}

func (o *PutStatusProbeFailed) GetPayload() models.Error {
	return o.Payload
}

func (o *PutStatusProbeFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
