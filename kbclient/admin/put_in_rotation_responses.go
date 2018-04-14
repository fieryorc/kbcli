// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutInRotationReader is a Reader for the PutInRotation structure.
type PutInRotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutInRotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutInRotationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutInRotationNoContent creates a PutInRotationNoContent with default headers values
func NewPutInRotationNoContent() *PutInRotationNoContent {
	return &PutInRotationNoContent{}
}

/*PutInRotationNoContent handles this case with default header values.

Successful operation
*/
type PutInRotationNoContent struct {
}

func (o *PutInRotationNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/admin/healthcheck][%d] putInRotationNoContent ", 204)
}

func (o *PutInRotationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}