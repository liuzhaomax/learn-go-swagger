// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetPetReader is a Reader for the GetPet structure.
type GetPetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPetOK creates a GetPetOK with default headers values
func NewGetPetOK() *GetPetOK {
	return &GetPetOK{}
}

/*
GetPetOK describes a response with status code 200, with default header values.

A random pet
*/
type GetPetOK struct {
}

// IsSuccess returns true when this get pet o k response has a 2xx status code
func (o *GetPetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pet o k response has a 3xx status code
func (o *GetPetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pet o k response has a 4xx status code
func (o *GetPetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pet o k response has a 5xx status code
func (o *GetPetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pet o k response a status code equal to that given
func (o *GetPetOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPetOK) Error() string {
	return fmt.Sprintf("[GET /pet][%d] getPetOK ", 200)
}

func (o *GetPetOK) String() string {
	return fmt.Sprintf("[GET /pet][%d] getPetOK ", 200)
}

func (o *GetPetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPetDefault creates a GetPetDefault with default headers values
func NewGetPetDefault(code int) *GetPetDefault {
	return &GetPetDefault{
		_statusCode: code,
	}
}

/*
GetPetDefault describes a response with status code -1, with default header values.

unexpected error
*/
type GetPetDefault struct {
	_statusCode int
}

// Code gets the status code for the get pet default response
func (o *GetPetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get pet default response has a 2xx status code
func (o *GetPetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get pet default response has a 3xx status code
func (o *GetPetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get pet default response has a 4xx status code
func (o *GetPetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get pet default response has a 5xx status code
func (o *GetPetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get pet default response a status code equal to that given
func (o *GetPetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetPetDefault) Error() string {
	return fmt.Sprintf("[GET /pet][%d] getPet default ", o._statusCode)
}

func (o *GetPetDefault) String() string {
	return fmt.Sprintf("[GET /pet][%d] getPet default ", o._statusCode)
}

func (o *GetPetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
