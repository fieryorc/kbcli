// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewGetInvoicePaymentsParams creates a new GetInvoicePaymentsParams object
// with the default values initialized.
func NewGetInvoicePaymentsParams() *GetInvoicePaymentsParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentsParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoicePaymentsParamsWithTimeout creates a new GetInvoicePaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoicePaymentsParamsWithTimeout(timeout time.Duration) *GetInvoicePaymentsParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentsParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetInvoicePaymentsParamsWithContext creates a new GetInvoicePaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoicePaymentsParamsWithContext(ctx context.Context) *GetInvoicePaymentsParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentsParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetInvoicePaymentsParamsWithHTTPClient creates a new GetInvoicePaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoicePaymentsParamsWithHTTPClient(client *http.Client) *GetInvoicePaymentsParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentsParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*GetInvoicePaymentsParams contains all the parameters to send to the API endpoint
for the get invoice payments operation typically these are written to a http.Request
*/
type GetInvoicePaymentsParams struct {

	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string
	/*PluginProperty*/
	PluginProperty []string
	/*WithAttempts*/
	WithAttempts *bool
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithTimeout(timeout time.Duration) *GetInvoicePaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithContext(ctx context.Context) *GetInvoicePaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithHTTPClient(client *http.Client) *GetInvoicePaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithAccountID(accountID strfmt.UUID) *GetInvoicePaymentsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithAudit(audit *string) *GetInvoicePaymentsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPluginProperty adds the pluginProperty to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithPluginProperty(pluginProperty []string) *GetInvoicePaymentsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithAttempts adds the withAttempts to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithWithAttempts(withAttempts *bool) *GetInvoicePaymentsParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get invoice payments params
func (o *GetInvoicePaymentsParams) WithWithPluginInfo(withPluginInfo *bool) *GetInvoicePaymentsParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get invoice payments params
func (o *GetInvoicePaymentsParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoicePaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.WithAttempts != nil {

		// query param withAttempts
		var qrWithAttempts bool
		if o.WithAttempts != nil {
			qrWithAttempts = *o.WithAttempts
		}
		qWithAttempts := swag.FormatBool(qrWithAttempts)
		if qWithAttempts != "" {
			if err := r.SetQueryParam("withAttempts", qWithAttempts); err != nil {
				return err
			}
		}

	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool
		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {
			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
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
