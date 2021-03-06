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

// NewSetDefaultPaymentMethodParams creates a new SetDefaultPaymentMethodParams object
// with the default values initialized.
func NewSetDefaultPaymentMethodParams() *SetDefaultPaymentMethodParams {
	var (
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &SetDefaultPaymentMethodParams{
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDefaultPaymentMethodParamsWithTimeout creates a new SetDefaultPaymentMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDefaultPaymentMethodParamsWithTimeout(timeout time.Duration) *SetDefaultPaymentMethodParams {
	var (
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &SetDefaultPaymentMethodParams{
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,

		timeout: timeout,
	}
}

// NewSetDefaultPaymentMethodParamsWithContext creates a new SetDefaultPaymentMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDefaultPaymentMethodParamsWithContext(ctx context.Context) *SetDefaultPaymentMethodParams {
	var (
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &SetDefaultPaymentMethodParams{
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,

		Context: ctx,
	}
}

// NewSetDefaultPaymentMethodParamsWithHTTPClient creates a new SetDefaultPaymentMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDefaultPaymentMethodParamsWithHTTPClient(client *http.Client) *SetDefaultPaymentMethodParams {
	var (
		payAllUnpaidInvoicesDefault = bool(false)
	)
	return &SetDefaultPaymentMethodParams{
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,
		HTTPClient:           client,
	}
}

/*SetDefaultPaymentMethodParams contains all the parameters to send to the API endpoint
for the set default payment method operation typically these are written to a http.Request
*/
type SetDefaultPaymentMethodParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*PayAllUnpaidInvoices*/
	PayAllUnpaidInvoices *bool
	/*PaymentMethodID*/
	PaymentMethodID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithTimeout(timeout time.Duration) *SetDefaultPaymentMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithContext(ctx context.Context) *SetDefaultPaymentMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithHTTPClient(client *http.Client) *SetDefaultPaymentMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithXKillbillComment(xKillbillComment *string) *SetDefaultPaymentMethodParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *SetDefaultPaymentMethodParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithXKillbillReason(xKillbillReason *string) *SetDefaultPaymentMethodParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithAccountID(accountID strfmt.UUID) *SetDefaultPaymentMethodParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithPayAllUnpaidInvoices adds the payAllUnpaidInvoices to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithPayAllUnpaidInvoices(payAllUnpaidInvoices *bool) *SetDefaultPaymentMethodParams {
	o.SetPayAllUnpaidInvoices(payAllUnpaidInvoices)
	return o
}

// SetPayAllUnpaidInvoices adds the payAllUnpaidInvoices to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetPayAllUnpaidInvoices(payAllUnpaidInvoices *bool) {
	o.PayAllUnpaidInvoices = payAllUnpaidInvoices
}

// WithPaymentMethodID adds the paymentMethodID to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *SetDefaultPaymentMethodParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithPluginProperty(pluginProperty []string) *SetDefaultPaymentMethodParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *SetDefaultPaymentMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.PayAllUnpaidInvoices != nil {

		// query param payAllUnpaidInvoices
		var qrPayAllUnpaidInvoices bool
		if o.PayAllUnpaidInvoices != nil {
			qrPayAllUnpaidInvoices = *o.PayAllUnpaidInvoices
		}
		qPayAllUnpaidInvoices := swag.FormatBool(qrPayAllUnpaidInvoices)
		if qPayAllUnpaidInvoices != "" {
			if err := r.SetQueryParam("payAllUnpaidInvoices", qPayAllUnpaidInvoices); err != nil {
				return err
			}
		}

	}

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
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
