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
)

// NewGetSubscriptionTagsParams creates a new GetSubscriptionTagsParams object
// with the default values initialized.
func NewGetSubscriptionTagsParams() *GetSubscriptionTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetSubscriptionTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubscriptionTagsParamsWithTimeout creates a new GetSubscriptionTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubscriptionTagsParamsWithTimeout(timeout time.Duration) *GetSubscriptionTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetSubscriptionTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetSubscriptionTagsParamsWithContext creates a new GetSubscriptionTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubscriptionTagsParamsWithContext(ctx context.Context) *GetSubscriptionTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetSubscriptionTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetSubscriptionTagsParamsWithHTTPClient creates a new GetSubscriptionTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubscriptionTagsParamsWithHTTPClient(client *http.Client) *GetSubscriptionTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetSubscriptionTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetSubscriptionTagsParams contains all the parameters to send to the API endpoint
for the get subscription tags operation typically these are written to a http.Request
*/
type GetSubscriptionTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*IncludedDeleted*/
	IncludedDeleted *bool
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithTimeout(timeout time.Duration) *GetSubscriptionTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithContext(ctx context.Context) *GetSubscriptionTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithHTTPClient(client *http.Client) *GetSubscriptionTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetSubscriptionTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetSubscriptionTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithAudit(audit *string) *GetSubscriptionTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetSubscriptionTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithSubscriptionID adds the subscriptionID to the get subscription tags params
func (o *GetSubscriptionTagsParams) WithSubscriptionID(subscriptionID strfmt.UUID) *GetSubscriptionTagsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the get subscription tags params
func (o *GetSubscriptionTagsParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubscriptionTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
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
