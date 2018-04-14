// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// AddBundleBlockingStateReader is a Reader for the AddBundleBlockingState structure.
type AddBundleBlockingStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBundleBlockingStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddBundleBlockingStateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddBundleBlockingStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddBundleBlockingStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddBundleBlockingStateCreated creates a AddBundleBlockingStateCreated with default headers values
func NewAddBundleBlockingStateCreated() *AddBundleBlockingStateCreated {
	return &AddBundleBlockingStateCreated{}
}

/*AddBundleBlockingStateCreated handles this case with default header values.

Blocking state created successfully
*/
type AddBundleBlockingStateCreated struct {
	Payload []*kbmodel.BlockingState
}

func (o *AddBundleBlockingStateCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/block][%d] addBundleBlockingStateCreated  %+v", 201, o.Payload)
}

func (o *AddBundleBlockingStateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBundleBlockingStateBadRequest creates a AddBundleBlockingStateBadRequest with default headers values
func NewAddBundleBlockingStateBadRequest() *AddBundleBlockingStateBadRequest {
	return &AddBundleBlockingStateBadRequest{}
}

/*AddBundleBlockingStateBadRequest handles this case with default header values.

Invalid bundle id supplied
*/
type AddBundleBlockingStateBadRequest struct {
}

func (o *AddBundleBlockingStateBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/block][%d] addBundleBlockingStateBadRequest ", 400)
}

func (o *AddBundleBlockingStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddBundleBlockingStateNotFound creates a AddBundleBlockingStateNotFound with default headers values
func NewAddBundleBlockingStateNotFound() *AddBundleBlockingStateNotFound {
	return &AddBundleBlockingStateNotFound{}
}

/*AddBundleBlockingStateNotFound handles this case with default header values.

Bundle not found
*/
type AddBundleBlockingStateNotFound struct {
}

func (o *AddBundleBlockingStateNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/block][%d] addBundleBlockingStateNotFound ", 404)
}

func (o *AddBundleBlockingStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}