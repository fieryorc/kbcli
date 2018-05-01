// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewInvalidatesCacheParams creates a new InvalidatesCacheParams object
// with the default values initialized.
func NewInvalidatesCacheParams() *InvalidatesCacheParams {
	var ()
	return &InvalidatesCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidatesCacheParamsWithTimeout creates a new InvalidatesCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvalidatesCacheParamsWithTimeout(timeout time.Duration) *InvalidatesCacheParams {
	var ()
	return &InvalidatesCacheParams{

		timeout: timeout,
	}
}

// NewInvalidatesCacheParamsWithContext creates a new InvalidatesCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvalidatesCacheParamsWithContext(ctx context.Context) *InvalidatesCacheParams {
	var ()
	return &InvalidatesCacheParams{

		Context: ctx,
	}
}

// NewInvalidatesCacheParamsWithHTTPClient creates a new InvalidatesCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvalidatesCacheParamsWithHTTPClient(client *http.Client) *InvalidatesCacheParams {
	var ()
	return &InvalidatesCacheParams{
		HTTPClient: client,
	}
}

/*InvalidatesCacheParams contains all the parameters to send to the API endpoint
for the invalidates cache operation typically these are written to a http.Request
*/
type InvalidatesCacheParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*CacheName*/
	CacheName *string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the invalidates cache params
func (o *InvalidatesCacheParams) WithTimeout(timeout time.Duration) *InvalidatesCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidates cache params
func (o *InvalidatesCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidates cache params
func (o *InvalidatesCacheParams) WithContext(ctx context.Context) *InvalidatesCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidates cache params
func (o *InvalidatesCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidates cache params
func (o *InvalidatesCacheParams) WithHTTPClient(client *http.Client) *InvalidatesCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidates cache params
func (o *InvalidatesCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the invalidates cache params
func (o *InvalidatesCacheParams) WithXKillbillAPIKey(xKillbillAPIKey string) *InvalidatesCacheParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the invalidates cache params
func (o *InvalidatesCacheParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the invalidates cache params
func (o *InvalidatesCacheParams) WithXKillbillAPISecret(xKillbillAPISecret string) *InvalidatesCacheParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the invalidates cache params
func (o *InvalidatesCacheParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithCacheName adds the cacheName to the invalidates cache params
func (o *InvalidatesCacheParams) WithCacheName(cacheName *string) *InvalidatesCacheParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the invalidates cache params
func (o *InvalidatesCacheParams) SetCacheName(cacheName *string) {
	o.CacheName = cacheName
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidatesCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CacheName != nil {

		// query param cacheName
		var qrCacheName string
		if o.CacheName != nil {
			qrCacheName = *o.CacheName
		}
		qCacheName := qrCacheName
		if qCacheName != "" {
			if err := r.SetQueryParam("cacheName", qCacheName); err != nil {
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
