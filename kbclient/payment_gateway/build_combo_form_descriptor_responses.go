// Code generated by go-swagger; DO NOT EDIT.

package payment_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// BuildComboFormDescriptorReader is a Reader for the BuildComboFormDescriptor structure.
type BuildComboFormDescriptorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildComboFormDescriptorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildComboFormDescriptorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBuildComboFormDescriptorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewBuildComboFormDescriptorOK creates a BuildComboFormDescriptorOK with default headers values
func NewBuildComboFormDescriptorOK() *BuildComboFormDescriptorOK {
	return &BuildComboFormDescriptorOK{}
}

/*BuildComboFormDescriptorOK handles this case with default header values.

successful operation
*/
type BuildComboFormDescriptorOK struct {
	Payload *kbmodel.HostedPaymentPageFormDescriptor
}

func (o *BuildComboFormDescriptorOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentGateways/hosted/form][%d] buildComboFormDescriptorOK  %+v", 200, o.Payload)
}

func (o *BuildComboFormDescriptorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.HostedPaymentPageFormDescriptor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildComboFormDescriptorBadRequest creates a BuildComboFormDescriptorBadRequest with default headers values
func NewBuildComboFormDescriptorBadRequest() *BuildComboFormDescriptorBadRequest {
	return &BuildComboFormDescriptorBadRequest{}
}

/*BuildComboFormDescriptorBadRequest handles this case with default header values.

Invalid data for Account or PaymentMethod
*/
type BuildComboFormDescriptorBadRequest struct {
}

func (o *BuildComboFormDescriptorBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentGateways/hosted/form][%d] buildComboFormDescriptorBadRequest ", 400)
}

func (o *BuildComboFormDescriptorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
