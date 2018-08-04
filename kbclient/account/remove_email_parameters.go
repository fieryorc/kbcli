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

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveEmailParams creates a new RemoveEmailParams object
// with the default values initialized.
func NewRemoveEmailParams() *RemoveEmailParams {
	var ()
	return &RemoveEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveEmailParamsWithTimeout creates a new RemoveEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveEmailParamsWithTimeout(timeout time.Duration) *RemoveEmailParams {
	var ()
	return &RemoveEmailParams{

		timeout: timeout,
	}
}

// NewRemoveEmailParamsWithContext creates a new RemoveEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveEmailParamsWithContext(ctx context.Context) *RemoveEmailParams {
	var ()
	return &RemoveEmailParams{

		Context: ctx,
	}
}

// NewRemoveEmailParamsWithHTTPClient creates a new RemoveEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveEmailParamsWithHTTPClient(client *http.Client) *RemoveEmailParams {
	var ()
	return &RemoveEmailParams{
		HTTPClient: client,
	}
}

/*RemoveEmailParams contains all the parameters to send to the API endpoint
for the remove email operation typically these are written to a http.Request
*/
type RemoveEmailParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Email*/
	Email string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the remove email params
func (o *RemoveEmailParams) WithTimeout(timeout time.Duration) *RemoveEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove email params
func (o *RemoveEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove email params
func (o *RemoveEmailParams) WithContext(ctx context.Context) *RemoveEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove email params
func (o *RemoveEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove email params
func (o *RemoveEmailParams) WithHTTPClient(client *http.Client) *RemoveEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove email params
func (o *RemoveEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the remove email params
func (o *RemoveEmailParams) WithXKillbillComment(xKillbillComment *string) *RemoveEmailParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the remove email params
func (o *RemoveEmailParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the remove email params
func (o *RemoveEmailParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RemoveEmailParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the remove email params
func (o *RemoveEmailParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the remove email params
func (o *RemoveEmailParams) WithXKillbillReason(xKillbillReason *string) *RemoveEmailParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the remove email params
func (o *RemoveEmailParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the remove email params
func (o *RemoveEmailParams) WithAccountID(accountID strfmt.UUID) *RemoveEmailParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the remove email params
func (o *RemoveEmailParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithEmail adds the email to the remove email params
func (o *RemoveEmailParams) WithEmail(email string) *RemoveEmailParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the remove email params
func (o *RemoveEmailParams) SetEmail(email string) {
	o.Email = email
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param email
	if err := r.SetPathParam("email", o.Email); err != nil {
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
