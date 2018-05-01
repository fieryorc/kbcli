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

// NewGetPaymentsForAccountParams creates a new GetPaymentsForAccountParams object
// with the default values initialized.
func NewGetPaymentsForAccountParams() *GetPaymentsForAccountParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentsForAccountParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsForAccountParamsWithTimeout creates a new GetPaymentsForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsForAccountParamsWithTimeout(timeout time.Duration) *GetPaymentsForAccountParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentsForAccountParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetPaymentsForAccountParamsWithContext creates a new GetPaymentsForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentsForAccountParamsWithContext(ctx context.Context) *GetPaymentsForAccountParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentsForAccountParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetPaymentsForAccountParamsWithHTTPClient creates a new GetPaymentsForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentsForAccountParamsWithHTTPClient(client *http.Client) *GetPaymentsForAccountParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentsForAccountParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*GetPaymentsForAccountParams contains all the parameters to send to the API endpoint
for the get payments for account operation typically these are written to a http.Request
*/
type GetPaymentsForAccountParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
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

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get payments for account params
func (o *GetPaymentsForAccountParams) WithTimeout(timeout time.Duration) *GetPaymentsForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payments for account params
func (o *GetPaymentsForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payments for account params
func (o *GetPaymentsForAccountParams) WithContext(ctx context.Context) *GetPaymentsForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payments for account params
func (o *GetPaymentsForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payments for account params
func (o *GetPaymentsForAccountParams) WithHTTPClient(client *http.Client) *GetPaymentsForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payments for account params
func (o *GetPaymentsForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get payments for account params
func (o *GetPaymentsForAccountParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetPaymentsForAccountParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get payments for account params
func (o *GetPaymentsForAccountParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get payments for account params
func (o *GetPaymentsForAccountParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetPaymentsForAccountParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get payments for account params
func (o *GetPaymentsForAccountParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get payments for account params
func (o *GetPaymentsForAccountParams) WithAccountID(accountID strfmt.UUID) *GetPaymentsForAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get payments for account params
func (o *GetPaymentsForAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get payments for account params
func (o *GetPaymentsForAccountParams) WithAudit(audit *string) *GetPaymentsForAccountParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payments for account params
func (o *GetPaymentsForAccountParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPluginProperty adds the pluginProperty to the get payments for account params
func (o *GetPaymentsForAccountParams) WithPluginProperty(pluginProperty []string) *GetPaymentsForAccountParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payments for account params
func (o *GetPaymentsForAccountParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithAttempts adds the withAttempts to the get payments for account params
func (o *GetPaymentsForAccountParams) WithWithAttempts(withAttempts *bool) *GetPaymentsForAccountParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get payments for account params
func (o *GetPaymentsForAccountParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get payments for account params
func (o *GetPaymentsForAccountParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentsForAccountParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payments for account params
func (o *GetPaymentsForAccountParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
