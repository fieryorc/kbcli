// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewNotifyStateChangedParams creates a new NotifyStateChangedParams object
// with the default values initialized.
func NewNotifyStateChangedParams() *NotifyStateChangedParams {
	var ()
	return &NotifyStateChangedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNotifyStateChangedParamsWithTimeout creates a new NotifyStateChangedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNotifyStateChangedParamsWithTimeout(timeout time.Duration) *NotifyStateChangedParams {
	var ()
	return &NotifyStateChangedParams{

		timeout: timeout,
	}
}

// NewNotifyStateChangedParamsWithContext creates a new NotifyStateChangedParams object
// with the default values initialized, and the ability to set a context for a request
func NewNotifyStateChangedParamsWithContext(ctx context.Context) *NotifyStateChangedParams {
	var ()
	return &NotifyStateChangedParams{

		Context: ctx,
	}
}

// NewNotifyStateChangedParamsWithHTTPClient creates a new NotifyStateChangedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNotifyStateChangedParamsWithHTTPClient(client *http.Client) *NotifyStateChangedParams {
	var ()
	return &NotifyStateChangedParams{
		HTTPClient: client,
	}
}

/*NotifyStateChangedParams contains all the parameters to send to the API endpoint
for the notify state changed operation typically these are written to a http.Request
*/
type NotifyStateChangedParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.PaymentTransaction
	/*ControlPluginName*/
	ControlPluginName []string
	/*TransactionID*/
	TransactionID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the notify state changed params
func (o *NotifyStateChangedParams) WithTimeout(timeout time.Duration) *NotifyStateChangedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notify state changed params
func (o *NotifyStateChangedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notify state changed params
func (o *NotifyStateChangedParams) WithContext(ctx context.Context) *NotifyStateChangedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notify state changed params
func (o *NotifyStateChangedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notify state changed params
func (o *NotifyStateChangedParams) WithHTTPClient(client *http.Client) *NotifyStateChangedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notify state changed params
func (o *NotifyStateChangedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the notify state changed params
func (o *NotifyStateChangedParams) WithXKillbillAPIKey(xKillbillAPIKey string) *NotifyStateChangedParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the notify state changed params
func (o *NotifyStateChangedParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the notify state changed params
func (o *NotifyStateChangedParams) WithXKillbillAPISecret(xKillbillAPISecret string) *NotifyStateChangedParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the notify state changed params
func (o *NotifyStateChangedParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the notify state changed params
func (o *NotifyStateChangedParams) WithXKillbillComment(xKillbillComment *string) *NotifyStateChangedParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the notify state changed params
func (o *NotifyStateChangedParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the notify state changed params
func (o *NotifyStateChangedParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *NotifyStateChangedParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the notify state changed params
func (o *NotifyStateChangedParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the notify state changed params
func (o *NotifyStateChangedParams) WithXKillbillReason(xKillbillReason *string) *NotifyStateChangedParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the notify state changed params
func (o *NotifyStateChangedParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the notify state changed params
func (o *NotifyStateChangedParams) WithBody(body *kbmodel.PaymentTransaction) *NotifyStateChangedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the notify state changed params
func (o *NotifyStateChangedParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the notify state changed params
func (o *NotifyStateChangedParams) WithControlPluginName(controlPluginName []string) *NotifyStateChangedParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the notify state changed params
func (o *NotifyStateChangedParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithTransactionID adds the transactionID to the notify state changed params
func (o *NotifyStateChangedParams) WithTransactionID(transactionID strfmt.UUID) *NotifyStateChangedParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the notify state changed params
func (o *NotifyStateChangedParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *NotifyStateChangedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesControlPluginName := o.ControlPluginName

	joinedControlPluginName := swag.JoinByFormat(valuesControlPluginName, "multi")
	// query array param controlPluginName
	if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
		return err
	}

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID.String()); err != nil {
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
