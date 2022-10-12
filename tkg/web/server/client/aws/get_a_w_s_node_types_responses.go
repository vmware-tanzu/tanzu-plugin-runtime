// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetAWSNodeTypesReader is a Reader for the GetAWSNodeTypes structure.
type GetAWSNodeTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAWSNodeTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAWSNodeTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAWSNodeTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAWSNodeTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAWSNodeTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAWSNodeTypesOK creates a GetAWSNodeTypesOK with default headers values
func NewGetAWSNodeTypesOK() *GetAWSNodeTypesOK {
	return &GetAWSNodeTypesOK{}
}

/*
GetAWSNodeTypesOK handles this case with default header values.

Successful retrieval of AWS node types
*/
type GetAWSNodeTypesOK struct {
	Payload []string
}

func (o *GetAWSNodeTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/nodetypes][%d] getAWSNodeTypesOK  %+v", 200, o.Payload)
}

func (o *GetAWSNodeTypesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAWSNodeTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSNodeTypesBadRequest creates a GetAWSNodeTypesBadRequest with default headers values
func NewGetAWSNodeTypesBadRequest() *GetAWSNodeTypesBadRequest {
	return &GetAWSNodeTypesBadRequest{}
}

/*
GetAWSNodeTypesBadRequest handles this case with default header values.

Bad request
*/
type GetAWSNodeTypesBadRequest struct {
	Payload *models.Error
}

func (o *GetAWSNodeTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/nodetypes][%d] getAWSNodeTypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAWSNodeTypesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSNodeTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSNodeTypesUnauthorized creates a GetAWSNodeTypesUnauthorized with default headers values
func NewGetAWSNodeTypesUnauthorized() *GetAWSNodeTypesUnauthorized {
	return &GetAWSNodeTypesUnauthorized{}
}

/*
GetAWSNodeTypesUnauthorized handles this case with default header values.

Incorrect credentials
*/
type GetAWSNodeTypesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAWSNodeTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/nodetypes][%d] getAWSNodeTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAWSNodeTypesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSNodeTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAWSNodeTypesInternalServerError creates a GetAWSNodeTypesInternalServerError with default headers values
func NewGetAWSNodeTypesInternalServerError() *GetAWSNodeTypesInternalServerError {
	return &GetAWSNodeTypesInternalServerError{}
}

/*
GetAWSNodeTypesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAWSNodeTypesInternalServerError struct {
	Payload *models.Error
}

func (o *GetAWSNodeTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/providers/aws/nodetypes][%d] getAWSNodeTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAWSNodeTypesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAWSNodeTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
