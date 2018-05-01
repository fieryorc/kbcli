// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// GetPriceListForSubscriptionAndDateReader is a Reader for the GetPriceListForSubscriptionAndDate structure.
type GetPriceListForSubscriptionAndDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPriceListForSubscriptionAndDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPriceListForSubscriptionAndDateOK()
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

// NewGetPriceListForSubscriptionAndDateOK creates a GetPriceListForSubscriptionAndDateOK with default headers values
func NewGetPriceListForSubscriptionAndDateOK() *GetPriceListForSubscriptionAndDateOK {
	return &GetPriceListForSubscriptionAndDateOK{}
}

/*GetPriceListForSubscriptionAndDateOK handles this case with default header values.

successful operation
*/
type GetPriceListForSubscriptionAndDateOK struct {
	Payload *kbmodel.PriceList
}

func (o *GetPriceListForSubscriptionAndDateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/priceList][%d] getPriceListForSubscriptionAndDateOK  %+v", 200, o.Payload)
}

func (o *GetPriceListForSubscriptionAndDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.PriceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
