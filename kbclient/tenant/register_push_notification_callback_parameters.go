// Code generated by go-swagger; DO NOT EDIT.

package tenant

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
)

// NewRegisterPushNotificationCallbackParams creates a new RegisterPushNotificationCallbackParams object
// with the default values initialized.
func NewRegisterPushNotificationCallbackParams() *RegisterPushNotificationCallbackParams {
	var ()
	return &RegisterPushNotificationCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterPushNotificationCallbackParamsWithTimeout creates a new RegisterPushNotificationCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterPushNotificationCallbackParamsWithTimeout(timeout time.Duration) *RegisterPushNotificationCallbackParams {
	var ()
	return &RegisterPushNotificationCallbackParams{

		timeout: timeout,
	}
}

// NewRegisterPushNotificationCallbackParamsWithContext creates a new RegisterPushNotificationCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterPushNotificationCallbackParamsWithContext(ctx context.Context) *RegisterPushNotificationCallbackParams {
	var ()
	return &RegisterPushNotificationCallbackParams{

		Context: ctx,
	}
}

// NewRegisterPushNotificationCallbackParamsWithHTTPClient creates a new RegisterPushNotificationCallbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterPushNotificationCallbackParamsWithHTTPClient(client *http.Client) *RegisterPushNotificationCallbackParams {
	var ()
	return &RegisterPushNotificationCallbackParams{
		HTTPClient: client,
	}
}

/*RegisterPushNotificationCallbackParams contains all the parameters to send to the API endpoint
for the register push notification callback operation typically these are written to a http.Request
*/
type RegisterPushNotificationCallbackParams struct {

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
	/*Cb*/
	Cb *string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithTimeout(timeout time.Duration) *RegisterPushNotificationCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithContext(ctx context.Context) *RegisterPushNotificationCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithHTTPClient(client *http.Client) *RegisterPushNotificationCallbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillAPIKey(xKillbillAPIKey string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillAPISecret(xKillbillAPISecret string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillComment(xKillbillComment *string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithXKillbillReason(xKillbillReason *string) *RegisterPushNotificationCallbackParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithCb adds the cb to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) WithCb(cb *string) *RegisterPushNotificationCallbackParams {
	o.SetCb(cb)
	return o
}

// SetCb adds the cb to the register push notification callback params
func (o *RegisterPushNotificationCallbackParams) SetCb(cb *string) {
	o.Cb = cb
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterPushNotificationCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Cb != nil {

		// query param cb
		var qrCb string
		if o.Cb != nil {
			qrCb = *o.Cb
		}
		qCb := qrCb
		if qCb != "" {
			if err := r.SetQueryParam("cb", qCb); err != nil {
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
