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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/ziti/controller/rest_model"
)

// DetailServiceReader is a Reader for the DetailService structure.
type DetailServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDetailServiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDetailServiceOK creates a DetailServiceOK with default headers values
func NewDetailServiceOK() *DetailServiceOK {
	return &DetailServiceOK{}
}

/* DetailServiceOK describes a response with status code 200, with default header values.

A single service
*/
type DetailServiceOK struct {
	Payload *rest_model.DetailServiceEnvelope
}

func (o *DetailServiceOK) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] detailServiceOK  %+v", 200, o.Payload)
}
func (o *DetailServiceOK) GetPayload() *rest_model.DetailServiceEnvelope {
	return o.Payload
}

func (o *DetailServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailServiceEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailServiceUnauthorized creates a DetailServiceUnauthorized with default headers values
func NewDetailServiceUnauthorized() *DetailServiceUnauthorized {
	return &DetailServiceUnauthorized{}
}

/* DetailServiceUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailServiceUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] detailServiceUnauthorized  %+v", 401, o.Payload)
}
func (o *DetailServiceUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailServiceNotFound creates a DetailServiceNotFound with default headers values
func NewDetailServiceNotFound() *DetailServiceNotFound {
	return &DetailServiceNotFound{}
}

/* DetailServiceNotFound describes a response with status code 404, with default header values.

The requested resource does not exist
*/
type DetailServiceNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailServiceNotFound) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] detailServiceNotFound  %+v", 404, o.Payload)
}
func (o *DetailServiceNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailServiceTooManyRequests creates a DetailServiceTooManyRequests with default headers values
func NewDetailServiceTooManyRequests() *DetailServiceTooManyRequests {
	return &DetailServiceTooManyRequests{}
}

/* DetailServiceTooManyRequests describes a response with status code 429, with default header values.

The resource requested is rate limited and the rate limit has been exceeded
*/
type DetailServiceTooManyRequests struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailServiceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /services/{id}][%d] detailServiceTooManyRequests  %+v", 429, o.Payload)
}
func (o *DetailServiceTooManyRequests) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailServiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
