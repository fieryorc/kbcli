// Code generated by go-swagger; DO NOT EDIT.

package usage

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

// GetUsageReader is a Reader for the GetUsage structure.
type GetUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUsageBadRequest()
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

// NewGetUsageOK creates a GetUsageOK with default headers values
func NewGetUsageOK() *GetUsageOK {
	return &GetUsageOK{}
}

/*GetUsageOK handles this case with default header values.

successful operation
*/
type GetUsageOK struct {
	Payload *kbmodel.RolledUpUsage
}

func (o *GetUsageOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/usages/{subscriptionId}/{unitType}][%d] getUsageOK  %+v", 200, o.Payload)
}

func (o *GetUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.RolledUpUsage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsageBadRequest creates a GetUsageBadRequest with default headers values
func NewGetUsageBadRequest() *GetUsageBadRequest {
	return &GetUsageBadRequest{}
}

/*GetUsageBadRequest handles this case with default header values.

Missing start date or end date
*/
type GetUsageBadRequest struct {
}

func (o *GetUsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/usages/{subscriptionId}/{unitType}][%d] getUsageBadRequest ", 400)
}

func (o *GetUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
