// Code generated by go-swagger; DO NOT EDIT.

package payment_gateway

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

// NewBuildFormDescriptorParams creates a new BuildFormDescriptorParams object
// with the default values initialized.
func NewBuildFormDescriptorParams() *BuildFormDescriptorParams {
	var ()
	return &BuildFormDescriptorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildFormDescriptorParamsWithTimeout creates a new BuildFormDescriptorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildFormDescriptorParamsWithTimeout(timeout time.Duration) *BuildFormDescriptorParams {
	var ()
	return &BuildFormDescriptorParams{

		timeout: timeout,
	}
}

// NewBuildFormDescriptorParamsWithContext creates a new BuildFormDescriptorParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildFormDescriptorParamsWithContext(ctx context.Context) *BuildFormDescriptorParams {
	var ()
	return &BuildFormDescriptorParams{

		Context: ctx,
	}
}

// NewBuildFormDescriptorParamsWithHTTPClient creates a new BuildFormDescriptorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildFormDescriptorParamsWithHTTPClient(client *http.Client) *BuildFormDescriptorParams {
	var ()
	return &BuildFormDescriptorParams{
		HTTPClient: client,
	}
}

/*BuildFormDescriptorParams contains all the parameters to send to the API endpoint
for the build form descriptor operation typically these are written to a http.Request
*/
type BuildFormDescriptorParams struct {

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
	/*AccountID*/
	AccountID strfmt.UUID
	/*Body*/
	Body *kbmodel.HostedPaymentPageFields
	/*ControlPluginName*/
	ControlPluginName []string
	/*PaymentMethodID*/
	PaymentMethodID *strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the build form descriptor params
func (o *BuildFormDescriptorParams) WithTimeout(timeout time.Duration) *BuildFormDescriptorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build form descriptor params
func (o *BuildFormDescriptorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build form descriptor params
func (o *BuildFormDescriptorParams) WithContext(ctx context.Context) *BuildFormDescriptorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build form descriptor params
func (o *BuildFormDescriptorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build form descriptor params
func (o *BuildFormDescriptorParams) WithHTTPClient(client *http.Client) *BuildFormDescriptorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build form descriptor params
func (o *BuildFormDescriptorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the build form descriptor params
func (o *BuildFormDescriptorParams) WithXKillbillAPIKey(xKillbillAPIKey string) *BuildFormDescriptorParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the build form descriptor params
func (o *BuildFormDescriptorParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the build form descriptor params
func (o *BuildFormDescriptorParams) WithXKillbillAPISecret(xKillbillAPISecret string) *BuildFormDescriptorParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the build form descriptor params
func (o *BuildFormDescriptorParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the build form descriptor params
func (o *BuildFormDescriptorParams) WithXKillbillComment(xKillbillComment *string) *BuildFormDescriptorParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the build form descriptor params
func (o *BuildFormDescriptorParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the build form descriptor params
func (o *BuildFormDescriptorParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *BuildFormDescriptorParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the build form descriptor params
func (o *BuildFormDescriptorParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the build form descriptor params
func (o *BuildFormDescriptorParams) WithXKillbillReason(xKillbillReason *string) *BuildFormDescriptorParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the build form descriptor params
func (o *BuildFormDescriptorParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the build form descriptor params
func (o *BuildFormDescriptorParams) WithAccountID(accountID strfmt.UUID) *BuildFormDescriptorParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the build form descriptor params
func (o *BuildFormDescriptorParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the build form descriptor params
func (o *BuildFormDescriptorParams) WithBody(body *kbmodel.HostedPaymentPageFields) *BuildFormDescriptorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the build form descriptor params
func (o *BuildFormDescriptorParams) SetBody(body *kbmodel.HostedPaymentPageFields) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the build form descriptor params
func (o *BuildFormDescriptorParams) WithControlPluginName(controlPluginName []string) *BuildFormDescriptorParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the build form descriptor params
func (o *BuildFormDescriptorParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentMethodID adds the paymentMethodID to the build form descriptor params
func (o *BuildFormDescriptorParams) WithPaymentMethodID(paymentMethodID *strfmt.UUID) *BuildFormDescriptorParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the build form descriptor params
func (o *BuildFormDescriptorParams) SetPaymentMethodID(paymentMethodID *strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the build form descriptor params
func (o *BuildFormDescriptorParams) WithPluginProperty(pluginProperty []string) *BuildFormDescriptorParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the build form descriptor params
func (o *BuildFormDescriptorParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *BuildFormDescriptorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
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

	if o.PaymentMethodID != nil {

		// query param paymentMethodId
		var qrPaymentMethodID strfmt.UUID
		if o.PaymentMethodID != nil {
			qrPaymentMethodID = *o.PaymentMethodID
		}
		qPaymentMethodID := qrPaymentMethodID.String()
		if qPaymentMethodID != "" {
			if err := r.SetQueryParam("paymentMethodId", qPaymentMethodID); err != nil {
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
