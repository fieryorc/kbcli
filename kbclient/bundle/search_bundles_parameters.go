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

// NewSearchBundlesParams creates a new SearchBundlesParams object
// with the default values initialized.
func NewSearchBundlesParams() *SearchBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchBundlesParamsWithTimeout creates a new SearchBundlesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchBundlesParamsWithTimeout(timeout time.Duration) *SearchBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewSearchBundlesParamsWithContext creates a new SearchBundlesParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchBundlesParamsWithContext(ctx context.Context) *SearchBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewSearchBundlesParamsWithHTTPClient creates a new SearchBundlesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchBundlesParamsWithHTTPClient(client *http.Client) *SearchBundlesParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &SearchBundlesParams{
		Audit:      &auditDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*SearchBundlesParams contains all the parameters to send to the API endpoint
for the search bundles operation typically these are written to a http.Request
*/
type SearchBundlesParams struct {

	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*SearchKey*/
	SearchKey string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the search bundles params
func (o *SearchBundlesParams) WithTimeout(timeout time.Duration) *SearchBundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search bundles params
func (o *SearchBundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search bundles params
func (o *SearchBundlesParams) WithContext(ctx context.Context) *SearchBundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search bundles params
func (o *SearchBundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search bundles params
func (o *SearchBundlesParams) WithHTTPClient(client *http.Client) *SearchBundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search bundles params
func (o *SearchBundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the search bundles params
func (o *SearchBundlesParams) WithAudit(audit *string) *SearchBundlesParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search bundles params
func (o *SearchBundlesParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search bundles params
func (o *SearchBundlesParams) WithLimit(limit *int64) *SearchBundlesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search bundles params
func (o *SearchBundlesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search bundles params
func (o *SearchBundlesParams) WithOffset(offset *int64) *SearchBundlesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search bundles params
func (o *SearchBundlesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearchKey adds the searchKey to the search bundles params
func (o *SearchBundlesParams) WithSearchKey(searchKey string) *SearchBundlesParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search bundles params
func (o *SearchBundlesParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WriteToRequest writes these params to a swagger request
func (o *SearchBundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchKey
	if err := r.SetPathParam("searchKey", o.SearchKey); err != nil {
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
