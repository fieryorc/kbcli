// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyTransactionCustomFieldsReader is a Reader for the ModifyTransactionCustomFields structure.
type ModifyTransactionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyTransactionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyTransactionCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyTransactionCustomFieldsBadRequest()
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

// NewModifyTransactionCustomFieldsNoContent creates a ModifyTransactionCustomFieldsNoContent with default headers values
func NewModifyTransactionCustomFieldsNoContent() *ModifyTransactionCustomFieldsNoContent {
	return &ModifyTransactionCustomFieldsNoContent{}
}

/*ModifyTransactionCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyTransactionCustomFieldsNoContent struct {
}

func (o *ModifyTransactionCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/paymentTransactions/{transactionId}/customFields][%d] modifyTransactionCustomFieldsNoContent ", 204)
}

func (o *ModifyTransactionCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyTransactionCustomFieldsBadRequest creates a ModifyTransactionCustomFieldsBadRequest with default headers values
func NewModifyTransactionCustomFieldsBadRequest() *ModifyTransactionCustomFieldsBadRequest {
	return &ModifyTransactionCustomFieldsBadRequest{}
}

/*ModifyTransactionCustomFieldsBadRequest handles this case with default header values.

Invalid transaction id supplied
*/
type ModifyTransactionCustomFieldsBadRequest struct {
}

func (o *ModifyTransactionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/paymentTransactions/{transactionId}/customFields][%d] modifyTransactionCustomFieldsBadRequest ", 400)
}

func (o *ModifyTransactionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
