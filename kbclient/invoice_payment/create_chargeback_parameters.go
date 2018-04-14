// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateChargebackParams creates a new CreateChargebackParams object
// with the default values initialized.
func NewCreateChargebackParams() *CreateChargebackParams {
	var ()
	return &CreateChargebackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChargebackParamsWithTimeout creates a new CreateChargebackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateChargebackParamsWithTimeout(timeout time.Duration) *CreateChargebackParams {
	var ()
	return &CreateChargebackParams{

		timeout: timeout,
	}
}

// NewCreateChargebackParamsWithContext creates a new CreateChargebackParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateChargebackParamsWithContext(ctx context.Context) *CreateChargebackParams {
	var ()
	return &CreateChargebackParams{

		Context: ctx,
	}
}

// NewCreateChargebackParamsWithHTTPClient creates a new CreateChargebackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateChargebackParamsWithHTTPClient(client *http.Client) *CreateChargebackParams {
	var ()
	return &CreateChargebackParams{
		HTTPClient: client,
	}
}

/*CreateChargebackParams contains all the parameters to send to the API endpoint
for the create chargeback operation typically these are written to a http.Request
*/
type CreateChargebackParams struct {

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
	Body *kbmodel.InvoicePaymentTransaction
	/*PaymentID*/
	PaymentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create chargeback params
func (o *CreateChargebackParams) WithTimeout(timeout time.Duration) *CreateChargebackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create chargeback params
func (o *CreateChargebackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create chargeback params
func (o *CreateChargebackParams) WithContext(ctx context.Context) *CreateChargebackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create chargeback params
func (o *CreateChargebackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create chargeback params
func (o *CreateChargebackParams) WithHTTPClient(client *http.Client) *CreateChargebackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create chargeback params
func (o *CreateChargebackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create chargeback params
func (o *CreateChargebackParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateChargebackParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create chargeback params
func (o *CreateChargebackParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create chargeback params
func (o *CreateChargebackParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateChargebackParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create chargeback params
func (o *CreateChargebackParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create chargeback params
func (o *CreateChargebackParams) WithXKillbillComment(xKillbillComment *string) *CreateChargebackParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create chargeback params
func (o *CreateChargebackParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create chargeback params
func (o *CreateChargebackParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateChargebackParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create chargeback params
func (o *CreateChargebackParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create chargeback params
func (o *CreateChargebackParams) WithXKillbillReason(xKillbillReason *string) *CreateChargebackParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create chargeback params
func (o *CreateChargebackParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create chargeback params
func (o *CreateChargebackParams) WithBody(body *kbmodel.InvoicePaymentTransaction) *CreateChargebackParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create chargeback params
func (o *CreateChargebackParams) SetBody(body *kbmodel.InvoicePaymentTransaction) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the create chargeback params
func (o *CreateChargebackParams) WithPaymentID(paymentID strfmt.UUID) *CreateChargebackParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create chargeback params
func (o *CreateChargebackParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChargebackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}