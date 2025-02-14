// Code generated by go-swagger; DO NOT EDIT.

package global_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetGlobalServiceReader is a Reader for the GetGlobalService structure.
type GetGlobalServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalServiceOK creates a GetGlobalServiceOK with default headers values
func NewGetGlobalServiceOK() *GetGlobalServiceOK {
	return &GetGlobalServiceOK{}
}

/* GetGlobalServiceOK describes a response with status code 200, with default header values.

successfully got service details
*/
type GetGlobalServiceOK struct {
}

func (o *GetGlobalServiceOK) Error() string {
	return fmt.Sprintf("[GET /api/functions/{serviceName}][%d] getGlobalServiceOK ", 200)
}

func (o *GetGlobalServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
