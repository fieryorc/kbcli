// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewGetPerTenantConfigurationParams creates a new GetPerTenantConfigurationParams object
// with the default values initialized.
func NewGetPerTenantConfigurationParams() *GetPerTenantConfigurationParams {
	var ()
	return &GetPerTenantConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPerTenantConfigurationParamsWithTimeout creates a new GetPerTenantConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPerTenantConfigurationParamsWithTimeout(timeout time.Duration) *GetPerTenantConfigurationParams {
	var ()
	return &GetPerTenantConfigurationParams{

		timeout: timeout,
	}
}

// NewGetPerTenantConfigurationParamsWithContext creates a new GetPerTenantConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPerTenantConfigurationParamsWithContext(ctx context.Context) *GetPerTenantConfigurationParams {
	var ()
	return &GetPerTenantConfigurationParams{

		Context: ctx,
	}
}

// NewGetPerTenantConfigurationParamsWithHTTPClient creates a new GetPerTenantConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPerTenantConfigurationParamsWithHTTPClient(client *http.Client) *GetPerTenantConfigurationParams {
	var ()
	return &GetPerTenantConfigurationParams{
		HTTPClient: client,
	}
}

/*GetPerTenantConfigurationParams contains all the parameters to send to the API endpoint
for the get per tenant configuration operation typically these are written to a http.Request
*/
type GetPerTenantConfigurationParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) WithTimeout(timeout time.Duration) *GetPerTenantConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) WithContext(ctx context.Context) *GetPerTenantConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) WithHTTPClient(client *http.Client) *GetPerTenantConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetPerTenantConfigurationParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetPerTenantConfigurationParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get per tenant configuration params
func (o *GetPerTenantConfigurationParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WriteToRequest writes these params to a swagger request
func (o *GetPerTenantConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
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
