// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// SearchPaymentsReader is a Reader for the SearchPayments structure.
type SearchPaymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchPaymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchPaymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewSearchPaymentsOK creates a SearchPaymentsOK with default headers values
func NewSearchPaymentsOK() *SearchPaymentsOK {
	return &SearchPaymentsOK{}
}

/*SearchPaymentsOK handles this case with default header values.

successful operation
*/
type SearchPaymentsOK struct {
	Payload []*kbmodel.Payment
}

func (o *SearchPaymentsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/search/{searchKey}][%d] searchPaymentsOK  %+v", 200, o.Payload)
}

func (o *SearchPaymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
