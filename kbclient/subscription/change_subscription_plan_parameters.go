// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewChangeSubscriptionPlanParams creates a new ChangeSubscriptionPlanParams object
// with the default values initialized.
func NewChangeSubscriptionPlanParams() *ChangeSubscriptionPlanParams {
	var (
		callCompletionDefault = bool(false)
		callTimeoutSecDefault = int64(3)
	)
	return &ChangeSubscriptionPlanParams{
		CallCompletion: &callCompletionDefault,
		CallTimeoutSec: &callTimeoutSecDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeSubscriptionPlanParamsWithTimeout creates a new ChangeSubscriptionPlanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeSubscriptionPlanParamsWithTimeout(timeout time.Duration) *ChangeSubscriptionPlanParams {
	var (
		callCompletionDefault = bool(false)
		callTimeoutSecDefault = int64(3)
	)
	return &ChangeSubscriptionPlanParams{
		CallCompletion: &callCompletionDefault,
		CallTimeoutSec: &callTimeoutSecDefault,

		timeout: timeout,
	}
}

// NewChangeSubscriptionPlanParamsWithContext creates a new ChangeSubscriptionPlanParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeSubscriptionPlanParamsWithContext(ctx context.Context) *ChangeSubscriptionPlanParams {
	var (
		callCompletionDefault = bool(false)
		callTimeoutSecDefault = int64(3)
	)
	return &ChangeSubscriptionPlanParams{
		CallCompletion: &callCompletionDefault,
		CallTimeoutSec: &callTimeoutSecDefault,

		Context: ctx,
	}
}

// NewChangeSubscriptionPlanParamsWithHTTPClient creates a new ChangeSubscriptionPlanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeSubscriptionPlanParamsWithHTTPClient(client *http.Client) *ChangeSubscriptionPlanParams {
	var (
		callCompletionDefault = bool(false)
		callTimeoutSecDefault = int64(3)
	)
	return &ChangeSubscriptionPlanParams{
		CallCompletion: &callCompletionDefault,
		CallTimeoutSec: &callTimeoutSecDefault,
		HTTPClient:     client,
	}
}

/*ChangeSubscriptionPlanParams contains all the parameters to send to the API endpoint
for the change subscription plan operation typically these are written to a http.Request
*/
type ChangeSubscriptionPlanParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*BillingPolicy*/
	BillingPolicy *string
	/*Body*/
	Body *kbmodel.Subscription
	/*CallCompletion*/
	CallCompletion *bool
	/*CallTimeoutSec*/
	CallTimeoutSec *int64
	/*PluginProperty*/
	PluginProperty []string
	/*RequestedDate*/
	RequestedDate *strfmt.Date
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithTimeout(timeout time.Duration) *ChangeSubscriptionPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithContext(ctx context.Context) *ChangeSubscriptionPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithHTTPClient(client *http.Client) *ChangeSubscriptionPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithXKillbillComment(xKillbillComment *string) *ChangeSubscriptionPlanParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ChangeSubscriptionPlanParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithXKillbillReason(xKillbillReason *string) *ChangeSubscriptionPlanParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBillingPolicy adds the billingPolicy to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithBillingPolicy(billingPolicy *string) *ChangeSubscriptionPlanParams {
	o.SetBillingPolicy(billingPolicy)
	return o
}

// SetBillingPolicy adds the billingPolicy to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetBillingPolicy(billingPolicy *string) {
	o.BillingPolicy = billingPolicy
}

// WithBody adds the body to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithBody(body *kbmodel.Subscription) *ChangeSubscriptionPlanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetBody(body *kbmodel.Subscription) {
	o.Body = body
}

// WithCallCompletion adds the callCompletion to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithCallCompletion(callCompletion *bool) *ChangeSubscriptionPlanParams {
	o.SetCallCompletion(callCompletion)
	return o
}

// SetCallCompletion adds the callCompletion to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetCallCompletion(callCompletion *bool) {
	o.CallCompletion = callCompletion
}

// WithCallTimeoutSec adds the callTimeoutSec to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithCallTimeoutSec(callTimeoutSec *int64) *ChangeSubscriptionPlanParams {
	o.SetCallTimeoutSec(callTimeoutSec)
	return o
}

// SetCallTimeoutSec adds the callTimeoutSec to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetCallTimeoutSec(callTimeoutSec *int64) {
	o.CallTimeoutSec = callTimeoutSec
}

// WithPluginProperty adds the pluginProperty to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithPluginProperty(pluginProperty []string) *ChangeSubscriptionPlanParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithRequestedDate(requestedDate *strfmt.Date) *ChangeSubscriptionPlanParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WithSubscriptionID adds the subscriptionID to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) WithSubscriptionID(subscriptionID strfmt.UUID) *ChangeSubscriptionPlanParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the change subscription plan params
func (o *ChangeSubscriptionPlanParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeSubscriptionPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BillingPolicy != nil {

		// query param billingPolicy
		var qrBillingPolicy string
		if o.BillingPolicy != nil {
			qrBillingPolicy = *o.BillingPolicy
		}
		qBillingPolicy := qrBillingPolicy
		if qBillingPolicy != "" {
			if err := r.SetQueryParam("billingPolicy", qBillingPolicy); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CallCompletion != nil {

		// query param callCompletion
		var qrCallCompletion bool
		if o.CallCompletion != nil {
			qrCallCompletion = *o.CallCompletion
		}
		qCallCompletion := swag.FormatBool(qrCallCompletion)
		if qCallCompletion != "" {
			if err := r.SetQueryParam("callCompletion", qCallCompletion); err != nil {
				return err
			}
		}

	}

	if o.CallTimeoutSec != nil {

		// query param callTimeoutSec
		var qrCallTimeoutSec int64
		if o.CallTimeoutSec != nil {
			qrCallTimeoutSec = *o.CallTimeoutSec
		}
		qCallTimeoutSec := swag.FormatInt64(qrCallTimeoutSec)
		if qCallTimeoutSec != "" {
			if err := r.SetQueryParam("callTimeoutSec", qCallTimeoutSec); err != nil {
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

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date
		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {
			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
		}

	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
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
