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

// NewDeleteAccountCustomFieldsParams creates a new DeleteAccountCustomFieldsParams object
// with the default values initialized.
func NewDeleteAccountCustomFieldsParams() *DeleteAccountCustomFieldsParams {
	var ()
	return &DeleteAccountCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccountCustomFieldsParamsWithTimeout creates a new DeleteAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAccountCustomFieldsParamsWithTimeout(timeout time.Duration) *DeleteAccountCustomFieldsParams {
	var ()
	return &DeleteAccountCustomFieldsParams{

		timeout: timeout,
	}
}

// NewDeleteAccountCustomFieldsParamsWithContext creates a new DeleteAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAccountCustomFieldsParamsWithContext(ctx context.Context) *DeleteAccountCustomFieldsParams {
	var ()
	return &DeleteAccountCustomFieldsParams{

		Context: ctx,
	}
}

// NewDeleteAccountCustomFieldsParamsWithHTTPClient creates a new DeleteAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAccountCustomFieldsParamsWithHTTPClient(client *http.Client) *DeleteAccountCustomFieldsParams {
	var ()
	return &DeleteAccountCustomFieldsParams{
		HTTPClient: client,
	}
}

/*DeleteAccountCustomFieldsParams contains all the parameters to send to the API endpoint
for the delete account custom fields operation typically these are written to a http.Request
*/
type DeleteAccountCustomFieldsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*CustomField*/
	CustomField []strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithTimeout(timeout time.Duration) *DeleteAccountCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithContext(ctx context.Context) *DeleteAccountCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithHTTPClient(client *http.Client) *DeleteAccountCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeleteAccountCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteAccountCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeleteAccountCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithAccountID(accountID strfmt.UUID) *DeleteAccountCustomFieldsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithCustomField adds the customField to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeleteAccountCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete account custom fields params
func (o *DeleteAccountCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccountCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesCustomField []string
	for _, v := range o.CustomField {
		valuesCustomField = append(valuesCustomField, v.String())
	}

	joinedCustomField := swag.JoinByFormat(valuesCustomField, "multi")
	// query array param customField
	if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
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
