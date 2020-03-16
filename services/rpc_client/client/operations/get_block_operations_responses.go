// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBlockOperationsReader is a Reader for the GetBlockOperations structure.
type GetBlockOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlockOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetBlockOperationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlockOperationsOK creates a GetBlockOperationsOK with default headers values
func NewGetBlockOperationsOK() *GetBlockOperationsOK {
	return &GetBlockOperationsOK{}
}

/*GetBlockOperationsOK handles this case with default header values.

Endpoint for contract script
*/
type GetBlockOperationsOK struct {
	Payload [][]interface{}
}

func (o *GetBlockOperationsOK) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/{block}/operations][%d] getBlockOperationsOK  %+v", 200, o.Payload)
}

func (o *GetBlockOperationsOK) GetPayload() [][]interface{} {
	return o.Payload
}

func (o *GetBlockOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockOperationsInternalServerError creates a GetBlockOperationsInternalServerError with default headers values
func NewGetBlockOperationsInternalServerError() *GetBlockOperationsInternalServerError {
	return &GetBlockOperationsInternalServerError{}
}

/*GetBlockOperationsInternalServerError handles this case with default header values.

Internal error
*/
type GetBlockOperationsInternalServerError struct {
}

func (o *GetBlockOperationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/{block}/operations][%d] getBlockOperationsInternalServerError ", 500)
}

func (o *GetBlockOperationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
