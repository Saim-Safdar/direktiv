// Code generated by go-swagger; DO NOT EDIT.

package directory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/direktiv/direktiv/pkg/api/models"
)

// CreateDirectoryReader is a Reader for the CreateDirectory structure.
type CreateDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDirectoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDirectoryOK creates a CreateDirectoryOK with default headers values
func NewCreateDirectoryOK() *CreateDirectoryOK {
	return &CreateDirectoryOK{}
}

/* CreateDirectoryOK describes a response with status code 200, with default header values.

directory has been created
*/
type CreateDirectoryOK struct {
	Payload models.OkBody
}

func (o *CreateDirectoryOK) Error() string {
	return fmt.Sprintf("[PUT /api/namespaces/{namespace}/tree/{directory}?op=create-directory][%d] createDirectoryOK  %+v", 200, o.Payload)
}
func (o *CreateDirectoryOK) GetPayload() models.OkBody {
	return o.Payload
}

func (o *CreateDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDirectoryDefault creates a CreateDirectoryDefault with default headers values
func NewCreateDirectoryDefault(code int) *CreateDirectoryDefault {
	return &CreateDirectoryDefault{
		_statusCode: code,
	}
}

/* CreateDirectoryDefault describes a response with status code -1, with default header values.

an error has occurred
*/
type CreateDirectoryDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create directory default response
func (o *CreateDirectoryDefault) Code() int {
	return o._statusCode
}

func (o *CreateDirectoryDefault) Error() string {
	return fmt.Sprintf("[PUT /api/namespaces/{namespace}/tree/{directory}?op=create-directory][%d] createDirectory default  %+v", o._statusCode, o.Payload)
}
func (o *CreateDirectoryDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDirectoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
