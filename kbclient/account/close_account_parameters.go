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

// NewCloseAccountParams creates a new CloseAccountParams object
// with the default values initialized.
func NewCloseAccountParams() *CloseAccountParams {
	var (
		cancelAllSubscriptionsDefault    = bool(false)
		itemAdjustUnpaidInvoicesDefault  = bool(false)
		removeFutureNotificationsDefault = bool(true)
		writeOffUnpaidInvoicesDefault    = bool(false)
	)
	return &CloseAccountParams{
		CancelAllSubscriptions:    &cancelAllSubscriptionsDefault,
		ItemAdjustUnpaidInvoices:  &itemAdjustUnpaidInvoicesDefault,
		RemoveFutureNotifications: &removeFutureNotificationsDefault,
		WriteOffUnpaidInvoices:    &writeOffUnpaidInvoicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCloseAccountParamsWithTimeout creates a new CloseAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloseAccountParamsWithTimeout(timeout time.Duration) *CloseAccountParams {
	var (
		cancelAllSubscriptionsDefault    = bool(false)
		itemAdjustUnpaidInvoicesDefault  = bool(false)
		removeFutureNotificationsDefault = bool(true)
		writeOffUnpaidInvoicesDefault    = bool(false)
	)
	return &CloseAccountParams{
		CancelAllSubscriptions:    &cancelAllSubscriptionsDefault,
		ItemAdjustUnpaidInvoices:  &itemAdjustUnpaidInvoicesDefault,
		RemoveFutureNotifications: &removeFutureNotificationsDefault,
		WriteOffUnpaidInvoices:    &writeOffUnpaidInvoicesDefault,

		timeout: timeout,
	}
}

// NewCloseAccountParamsWithContext creates a new CloseAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloseAccountParamsWithContext(ctx context.Context) *CloseAccountParams {
	var (
		cancelAllSubscriptionsDefault    = bool(false)
		itemAdjustUnpaidInvoicesDefault  = bool(false)
		removeFutureNotificationsDefault = bool(true)
		writeOffUnpaidInvoicesDefault    = bool(false)
	)
	return &CloseAccountParams{
		CancelAllSubscriptions:    &cancelAllSubscriptionsDefault,
		ItemAdjustUnpaidInvoices:  &itemAdjustUnpaidInvoicesDefault,
		RemoveFutureNotifications: &removeFutureNotificationsDefault,
		WriteOffUnpaidInvoices:    &writeOffUnpaidInvoicesDefault,

		Context: ctx,
	}
}

// NewCloseAccountParamsWithHTTPClient creates a new CloseAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloseAccountParamsWithHTTPClient(client *http.Client) *CloseAccountParams {
	var (
		cancelAllSubscriptionsDefault    = bool(false)
		itemAdjustUnpaidInvoicesDefault  = bool(false)
		removeFutureNotificationsDefault = bool(true)
		writeOffUnpaidInvoicesDefault    = bool(false)
	)
	return &CloseAccountParams{
		CancelAllSubscriptions:    &cancelAllSubscriptionsDefault,
		ItemAdjustUnpaidInvoices:  &itemAdjustUnpaidInvoicesDefault,
		RemoveFutureNotifications: &removeFutureNotificationsDefault,
		WriteOffUnpaidInvoices:    &writeOffUnpaidInvoicesDefault,
		HTTPClient:                client,
	}
}

/*CloseAccountParams contains all the parameters to send to the API endpoint
for the close account operation typically these are written to a http.Request
*/
type CloseAccountParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*CancelAllSubscriptions*/
	CancelAllSubscriptions *bool
	/*ItemAdjustUnpaidInvoices*/
	ItemAdjustUnpaidInvoices *bool
	/*RemoveFutureNotifications*/
	RemoveFutureNotifications *bool
	/*WriteOffUnpaidInvoices*/
	WriteOffUnpaidInvoices *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the close account params
func (o *CloseAccountParams) WithTimeout(timeout time.Duration) *CloseAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the close account params
func (o *CloseAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the close account params
func (o *CloseAccountParams) WithContext(ctx context.Context) *CloseAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the close account params
func (o *CloseAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the close account params
func (o *CloseAccountParams) WithHTTPClient(client *http.Client) *CloseAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the close account params
func (o *CloseAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the close account params
func (o *CloseAccountParams) WithXKillbillComment(xKillbillComment *string) *CloseAccountParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the close account params
func (o *CloseAccountParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the close account params
func (o *CloseAccountParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CloseAccountParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the close account params
func (o *CloseAccountParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the close account params
func (o *CloseAccountParams) WithXKillbillReason(xKillbillReason *string) *CloseAccountParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the close account params
func (o *CloseAccountParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the close account params
func (o *CloseAccountParams) WithAccountID(accountID strfmt.UUID) *CloseAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the close account params
func (o *CloseAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithCancelAllSubscriptions adds the cancelAllSubscriptions to the close account params
func (o *CloseAccountParams) WithCancelAllSubscriptions(cancelAllSubscriptions *bool) *CloseAccountParams {
	o.SetCancelAllSubscriptions(cancelAllSubscriptions)
	return o
}

// SetCancelAllSubscriptions adds the cancelAllSubscriptions to the close account params
func (o *CloseAccountParams) SetCancelAllSubscriptions(cancelAllSubscriptions *bool) {
	o.CancelAllSubscriptions = cancelAllSubscriptions
}

// WithItemAdjustUnpaidInvoices adds the itemAdjustUnpaidInvoices to the close account params
func (o *CloseAccountParams) WithItemAdjustUnpaidInvoices(itemAdjustUnpaidInvoices *bool) *CloseAccountParams {
	o.SetItemAdjustUnpaidInvoices(itemAdjustUnpaidInvoices)
	return o
}

// SetItemAdjustUnpaidInvoices adds the itemAdjustUnpaidInvoices to the close account params
func (o *CloseAccountParams) SetItemAdjustUnpaidInvoices(itemAdjustUnpaidInvoices *bool) {
	o.ItemAdjustUnpaidInvoices = itemAdjustUnpaidInvoices
}

// WithRemoveFutureNotifications adds the removeFutureNotifications to the close account params
func (o *CloseAccountParams) WithRemoveFutureNotifications(removeFutureNotifications *bool) *CloseAccountParams {
	o.SetRemoveFutureNotifications(removeFutureNotifications)
	return o
}

// SetRemoveFutureNotifications adds the removeFutureNotifications to the close account params
func (o *CloseAccountParams) SetRemoveFutureNotifications(removeFutureNotifications *bool) {
	o.RemoveFutureNotifications = removeFutureNotifications
}

// WithWriteOffUnpaidInvoices adds the writeOffUnpaidInvoices to the close account params
func (o *CloseAccountParams) WithWriteOffUnpaidInvoices(writeOffUnpaidInvoices *bool) *CloseAccountParams {
	o.SetWriteOffUnpaidInvoices(writeOffUnpaidInvoices)
	return o
}

// SetWriteOffUnpaidInvoices adds the writeOffUnpaidInvoices to the close account params
func (o *CloseAccountParams) SetWriteOffUnpaidInvoices(writeOffUnpaidInvoices *bool) {
	o.WriteOffUnpaidInvoices = writeOffUnpaidInvoices
}

// WriteToRequest writes these params to a swagger request
func (o *CloseAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CancelAllSubscriptions != nil {

		// query param cancelAllSubscriptions
		var qrCancelAllSubscriptions bool
		if o.CancelAllSubscriptions != nil {
			qrCancelAllSubscriptions = *o.CancelAllSubscriptions
		}
		qCancelAllSubscriptions := swag.FormatBool(qrCancelAllSubscriptions)
		if qCancelAllSubscriptions != "" {
			if err := r.SetQueryParam("cancelAllSubscriptions", qCancelAllSubscriptions); err != nil {
				return err
			}
		}

	}

	if o.ItemAdjustUnpaidInvoices != nil {

		// query param itemAdjustUnpaidInvoices
		var qrItemAdjustUnpaidInvoices bool
		if o.ItemAdjustUnpaidInvoices != nil {
			qrItemAdjustUnpaidInvoices = *o.ItemAdjustUnpaidInvoices
		}
		qItemAdjustUnpaidInvoices := swag.FormatBool(qrItemAdjustUnpaidInvoices)
		if qItemAdjustUnpaidInvoices != "" {
			if err := r.SetQueryParam("itemAdjustUnpaidInvoices", qItemAdjustUnpaidInvoices); err != nil {
				return err
			}
		}

	}

	if o.RemoveFutureNotifications != nil {

		// query param removeFutureNotifications
		var qrRemoveFutureNotifications bool
		if o.RemoveFutureNotifications != nil {
			qrRemoveFutureNotifications = *o.RemoveFutureNotifications
		}
		qRemoveFutureNotifications := swag.FormatBool(qrRemoveFutureNotifications)
		if qRemoveFutureNotifications != "" {
			if err := r.SetQueryParam("removeFutureNotifications", qRemoveFutureNotifications); err != nil {
				return err
			}
		}

	}

	if o.WriteOffUnpaidInvoices != nil {

		// query param writeOffUnpaidInvoices
		var qrWriteOffUnpaidInvoices bool
		if o.WriteOffUnpaidInvoices != nil {
			qrWriteOffUnpaidInvoices = *o.WriteOffUnpaidInvoices
		}
		qWriteOffUnpaidInvoices := swag.FormatBool(qrWriteOffUnpaidInvoices)
		if qWriteOffUnpaidInvoices != "" {
			if err := r.SetQueryParam("writeOffUnpaidInvoices", qWriteOffUnpaidInvoices); err != nil {
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
