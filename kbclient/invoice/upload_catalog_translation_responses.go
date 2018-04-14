// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UploadCatalogTranslationReader is a Reader for the UploadCatalogTranslation structure.
type UploadCatalogTranslationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadCatalogTranslationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUploadCatalogTranslationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadCatalogTranslationCreated creates a UploadCatalogTranslationCreated with default headers values
func NewUploadCatalogTranslationCreated() *UploadCatalogTranslationCreated {
	return &UploadCatalogTranslationCreated{}
}

/*UploadCatalogTranslationCreated handles this case with default header values.

Uploaded catalog translation Successfully
*/
type UploadCatalogTranslationCreated struct {
	Payload string
}

func (o *UploadCatalogTranslationCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/catalogTranslation/{locale}][%d] uploadCatalogTranslationCreated  %+v", 201, o.Payload)
}

func (o *UploadCatalogTranslationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}