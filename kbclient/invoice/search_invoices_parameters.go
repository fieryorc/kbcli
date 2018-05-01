// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewSearchInvoicesParams creates a new SearchInvoicesParams object
// with the default values initialized.
func NewSearchInvoicesParams() *SearchInvoicesParams {
	var (
		auditDefault     = string("NONE")
		limitDefault     = int64(100)
		offsetDefault    = int64(0)
		withItemsDefault = bool(false)
	)
	return &SearchInvoicesParams{
		Audit:     &auditDefault,
		Limit:     &limitDefault,
		Offset:    &offsetDefault,
		WithItems: &withItemsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchInvoicesParamsWithTimeout creates a new SearchInvoicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchInvoicesParamsWithTimeout(timeout time.Duration) *SearchInvoicesParams {
	var (
		auditDefault     = string("NONE")
		limitDefault     = int64(100)
		offsetDefault    = int64(0)
		withItemsDefault = bool(false)
	)
	return &SearchInvoicesParams{
		Audit:     &auditDefault,
		Limit:     &limitDefault,
		Offset:    &offsetDefault,
		WithItems: &withItemsDefault,

		timeout: timeout,
	}
}

// NewSearchInvoicesParamsWithContext creates a new SearchInvoicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchInvoicesParamsWithContext(ctx context.Context) *SearchInvoicesParams {
	var (
		auditDefault     = string("NONE")
		limitDefault     = int64(100)
		offsetDefault    = int64(0)
		withItemsDefault = bool(false)
	)
	return &SearchInvoicesParams{
		Audit:     &auditDefault,
		Limit:     &limitDefault,
		Offset:    &offsetDefault,
		WithItems: &withItemsDefault,

		Context: ctx,
	}
}

// NewSearchInvoicesParamsWithHTTPClient creates a new SearchInvoicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchInvoicesParamsWithHTTPClient(client *http.Client) *SearchInvoicesParams {
	var (
		auditDefault     = string("NONE")
		limitDefault     = int64(100)
		offsetDefault    = int64(0)
		withItemsDefault = bool(false)
	)
	return &SearchInvoicesParams{
		Audit:      &auditDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		WithItems:  &withItemsDefault,
		HTTPClient: client,
	}
}

/*SearchInvoicesParams contains all the parameters to send to the API endpoint
for the search invoices operation typically these are written to a http.Request
*/
type SearchInvoicesParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*SearchKey*/
	SearchKey string
	/*WithItems*/
	WithItems *bool

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the search invoices params
func (o *SearchInvoicesParams) WithTimeout(timeout time.Duration) *SearchInvoicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search invoices params
func (o *SearchInvoicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search invoices params
func (o *SearchInvoicesParams) WithContext(ctx context.Context) *SearchInvoicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search invoices params
func (o *SearchInvoicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search invoices params
func (o *SearchInvoicesParams) WithHTTPClient(client *http.Client) *SearchInvoicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search invoices params
func (o *SearchInvoicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the search invoices params
func (o *SearchInvoicesParams) WithXKillbillAPIKey(xKillbillAPIKey string) *SearchInvoicesParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the search invoices params
func (o *SearchInvoicesParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the search invoices params
func (o *SearchInvoicesParams) WithXKillbillAPISecret(xKillbillAPISecret string) *SearchInvoicesParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the search invoices params
func (o *SearchInvoicesParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the search invoices params
func (o *SearchInvoicesParams) WithAudit(audit *string) *SearchInvoicesParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search invoices params
func (o *SearchInvoicesParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search invoices params
func (o *SearchInvoicesParams) WithLimit(limit *int64) *SearchInvoicesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search invoices params
func (o *SearchInvoicesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search invoices params
func (o *SearchInvoicesParams) WithOffset(offset *int64) *SearchInvoicesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search invoices params
func (o *SearchInvoicesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearchKey adds the searchKey to the search invoices params
func (o *SearchInvoicesParams) WithSearchKey(searchKey string) *SearchInvoicesParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search invoices params
func (o *SearchInvoicesParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WithWithItems adds the withItems to the search invoices params
func (o *SearchInvoicesParams) WithWithItems(withItems *bool) *SearchInvoicesParams {
	o.SetWithItems(withItems)
	return o
}

// SetWithItems adds the withItems to the search invoices params
func (o *SearchInvoicesParams) SetWithItems(withItems *bool) {
	o.WithItems = withItems
}

// WriteToRequest writes these params to a swagger request
func (o *SearchInvoicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.WithItems != nil {

		// query param withItems
		var qrWithItems bool
		if o.WithItems != nil {
			qrWithItems = *o.WithItems
		}
		qWithItems := swag.FormatBool(qrWithItems)
		if qWithItems != "" {
			if err := r.SetQueryParam("withItems", qWithItems); err != nil {
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
