// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPriceListForSubscriptionAndDateParams creates a new GetPriceListForSubscriptionAndDateParams object
// with the default values initialized.
func NewGetPriceListForSubscriptionAndDateParams() *GetPriceListForSubscriptionAndDateParams {
	var ()
	return &GetPriceListForSubscriptionAndDateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPriceListForSubscriptionAndDateParamsWithTimeout creates a new GetPriceListForSubscriptionAndDateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPriceListForSubscriptionAndDateParamsWithTimeout(timeout time.Duration) *GetPriceListForSubscriptionAndDateParams {
	var ()
	return &GetPriceListForSubscriptionAndDateParams{

		timeout: timeout,
	}
}

// NewGetPriceListForSubscriptionAndDateParamsWithContext creates a new GetPriceListForSubscriptionAndDateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPriceListForSubscriptionAndDateParamsWithContext(ctx context.Context) *GetPriceListForSubscriptionAndDateParams {
	var ()
	return &GetPriceListForSubscriptionAndDateParams{

		Context: ctx,
	}
}

// NewGetPriceListForSubscriptionAndDateParamsWithHTTPClient creates a new GetPriceListForSubscriptionAndDateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPriceListForSubscriptionAndDateParamsWithHTTPClient(client *http.Client) *GetPriceListForSubscriptionAndDateParams {
	var ()
	return &GetPriceListForSubscriptionAndDateParams{
		HTTPClient: client,
	}
}

/*GetPriceListForSubscriptionAndDateParams contains all the parameters to send to the API endpoint
for the get price list for subscription and date operation typically these are written to a http.Request
*/
type GetPriceListForSubscriptionAndDateParams struct {

	/*RequestedDate*/
	RequestedDate *strfmt.Date
	/*SubscriptionID*/
	SubscriptionID *strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) WithTimeout(timeout time.Duration) *GetPriceListForSubscriptionAndDateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) WithContext(ctx context.Context) *GetPriceListForSubscriptionAndDateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) WithHTTPClient(client *http.Client) *GetPriceListForSubscriptionAndDateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestedDate adds the requestedDate to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) WithRequestedDate(requestedDate *strfmt.Date) *GetPriceListForSubscriptionAndDateParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WithSubscriptionID adds the subscriptionID to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) WithSubscriptionID(subscriptionID *strfmt.UUID) *GetPriceListForSubscriptionAndDateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the get price list for subscription and date params
func (o *GetPriceListForSubscriptionAndDateParams) SetSubscriptionID(subscriptionID *strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPriceListForSubscriptionAndDateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date
		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {
			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
		}

	}

	if o.SubscriptionID != nil {

		// query param subscriptionId
		var qrSubscriptionID strfmt.UUID
		if o.SubscriptionID != nil {
			qrSubscriptionID = *o.SubscriptionID
		}
		qSubscriptionID := qrSubscriptionID.String()
		if qSubscriptionID != "" {
			if err := r.SetQueryParam("subscriptionId", qSubscriptionID); err != nil {
				return err
			}
		}

	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
