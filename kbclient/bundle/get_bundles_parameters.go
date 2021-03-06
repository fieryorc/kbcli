// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBundlesParams creates a new GetBundlesParams object
// with the default values initialized.
func NewGetBundlesParams() *GetBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundlesParamsWithTimeout creates a new GetBundlesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBundlesParamsWithTimeout(timeout time.Duration) *GetBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetBundlesParamsWithContext creates a new GetBundlesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBundlesParamsWithContext(ctx context.Context) *GetBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetBundlesParamsWithHTTPClient creates a new GetBundlesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBundlesParamsWithHTTPClient(client *http.Client) *GetBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetBundlesParams{
		Audit:      &auditDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetBundlesParams contains all the parameters to send to the API endpoint
for the get bundles operation typically these are written to a http.Request
*/
type GetBundlesParams struct {

	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get bundles params
func (o *GetBundlesParams) WithTimeout(timeout time.Duration) *GetBundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundles params
func (o *GetBundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundles params
func (o *GetBundlesParams) WithContext(ctx context.Context) *GetBundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundles params
func (o *GetBundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundles params
func (o *GetBundlesParams) WithHTTPClient(client *http.Client) *GetBundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundles params
func (o *GetBundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get bundles params
func (o *GetBundlesParams) WithAudit(audit *string) *GetBundlesParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundles params
func (o *GetBundlesParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the get bundles params
func (o *GetBundlesParams) WithLimit(limit *int64) *GetBundlesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get bundles params
func (o *GetBundlesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get bundles params
func (o *GetBundlesParams) WithOffset(offset *int64) *GetBundlesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get bundles params
func (o *GetBundlesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
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
