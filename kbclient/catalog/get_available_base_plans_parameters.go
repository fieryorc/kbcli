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

// NewGetAvailableBasePlansParams creates a new GetAvailableBasePlansParams object
// with the default values initialized.
func NewGetAvailableBasePlansParams() *GetAvailableBasePlansParams {
	var ()
	return &GetAvailableBasePlansParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvailableBasePlansParamsWithTimeout creates a new GetAvailableBasePlansParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAvailableBasePlansParamsWithTimeout(timeout time.Duration) *GetAvailableBasePlansParams {
	var ()
	return &GetAvailableBasePlansParams{

		timeout: timeout,
	}
}

// NewGetAvailableBasePlansParamsWithContext creates a new GetAvailableBasePlansParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAvailableBasePlansParamsWithContext(ctx context.Context) *GetAvailableBasePlansParams {
	var ()
	return &GetAvailableBasePlansParams{

		Context: ctx,
	}
}

// NewGetAvailableBasePlansParamsWithHTTPClient creates a new GetAvailableBasePlansParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAvailableBasePlansParamsWithHTTPClient(client *http.Client) *GetAvailableBasePlansParams {
	var ()
	return &GetAvailableBasePlansParams{
		HTTPClient: client,
	}
}

/*GetAvailableBasePlansParams contains all the parameters to send to the API endpoint
for the get available base plans operation typically these are written to a http.Request
*/
type GetAvailableBasePlansParams struct {

	/*AccountID*/
	AccountID *strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get available base plans params
func (o *GetAvailableBasePlansParams) WithTimeout(timeout time.Duration) *GetAvailableBasePlansParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get available base plans params
func (o *GetAvailableBasePlansParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get available base plans params
func (o *GetAvailableBasePlansParams) WithContext(ctx context.Context) *GetAvailableBasePlansParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get available base plans params
func (o *GetAvailableBasePlansParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get available base plans params
func (o *GetAvailableBasePlansParams) WithHTTPClient(client *http.Client) *GetAvailableBasePlansParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get available base plans params
func (o *GetAvailableBasePlansParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get available base plans params
func (o *GetAvailableBasePlansParams) WithAccountID(accountID *strfmt.UUID) *GetAvailableBasePlansParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get available base plans params
func (o *GetAvailableBasePlansParams) SetAccountID(accountID *strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvailableBasePlansParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID strfmt.UUID
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID.String()
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
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
