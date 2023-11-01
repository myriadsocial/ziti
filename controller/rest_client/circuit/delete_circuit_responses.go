// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package circuit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/ziti/controller/rest_model"
)

// DeleteCircuitReader is a Reader for the DeleteCircuit structure.
type DeleteCircuitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCircuitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCircuitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCircuitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCircuitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteCircuitConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteCircuitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCircuitOK creates a DeleteCircuitOK with default headers values
func NewDeleteCircuitOK() *DeleteCircuitOK {
	return &DeleteCircuitOK{}
}

/* DeleteCircuitOK describes a response with status code 200, with default header values.

The delete request was successful and the resource has been removed
*/
type DeleteCircuitOK struct {
	Payload *rest_model.Empty
}

func (o *DeleteCircuitOK) Error() string {
	return fmt.Sprintf("[DELETE /circuits/{id}][%d] deleteCircuitOK  %+v", 200, o.Payload)
}
func (o *DeleteCircuitOK) GetPayload() *rest_model.Empty {
	return o.Payload
}

func (o *DeleteCircuitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.Empty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCircuitBadRequest creates a DeleteCircuitBadRequest with default headers values
func NewDeleteCircuitBadRequest() *DeleteCircuitBadRequest {
	return &DeleteCircuitBadRequest{}
}

/* DeleteCircuitBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type DeleteCircuitBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCircuitBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /circuits/{id}][%d] deleteCircuitBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCircuitBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCircuitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCircuitUnauthorized creates a DeleteCircuitUnauthorized with default headers values
func NewDeleteCircuitUnauthorized() *DeleteCircuitUnauthorized {
	return &DeleteCircuitUnauthorized{}
}

/* DeleteCircuitUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DeleteCircuitUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCircuitUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /circuits/{id}][%d] deleteCircuitUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCircuitUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCircuitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCircuitConflict creates a DeleteCircuitConflict with default headers values
func NewDeleteCircuitConflict() *DeleteCircuitConflict {
	return &DeleteCircuitConflict{}
}

/* DeleteCircuitConflict describes a response with status code 409, with default header values.

The resource requested to be removed/altered cannot be as it is referenced by another object.
*/
type DeleteCircuitConflict struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCircuitConflict) Error() string {
	return fmt.Sprintf("[DELETE /circuits/{id}][%d] deleteCircuitConflict  %+v", 409, o.Payload)
}
func (o *DeleteCircuitConflict) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCircuitConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCircuitTooManyRequests creates a DeleteCircuitTooManyRequests with default headers values
func NewDeleteCircuitTooManyRequests() *DeleteCircuitTooManyRequests {
	return &DeleteCircuitTooManyRequests{}
}

/* DeleteCircuitTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DeleteCircuitTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DeleteCircuitTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /circuits/{id}][%d] deleteCircuitTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteCircuitTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DeleteCircuitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
