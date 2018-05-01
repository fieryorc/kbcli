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

// NewGetChildrenAccountsParams creates a new GetChildrenAccountsParams object
// with the default values initialized.
func NewGetChildrenAccountsParams() *GetChildrenAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetChildrenAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChildrenAccountsParamsWithTimeout creates a new GetChildrenAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChildrenAccountsParamsWithTimeout(timeout time.Duration) *GetChildrenAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetChildrenAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetChildrenAccountsParamsWithContext creates a new GetChildrenAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChildrenAccountsParamsWithContext(ctx context.Context) *GetChildrenAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetChildrenAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetChildrenAccountsParamsWithHTTPClient creates a new GetChildrenAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChildrenAccountsParamsWithHTTPClient(client *http.Client) *GetChildrenAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
	)
	return &GetChildrenAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetChildrenAccountsParams contains all the parameters to send to the API endpoint
for the get children accounts operation typically these are written to a http.Request
*/
type GetChildrenAccountsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID
	/*AccountWithBalance*/
	AccountWithBalance *bool
	/*AccountWithBalanceAndCBA*/
	AccountWithBalanceAndCBA *bool
	/*Audit*/
	Audit *string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get children accounts params
func (o *GetChildrenAccountsParams) WithTimeout(timeout time.Duration) *GetChildrenAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get children accounts params
func (o *GetChildrenAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get children accounts params
func (o *GetChildrenAccountsParams) WithContext(ctx context.Context) *GetChildrenAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get children accounts params
func (o *GetChildrenAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get children accounts params
func (o *GetChildrenAccountsParams) WithHTTPClient(client *http.Client) *GetChildrenAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get children accounts params
func (o *GetChildrenAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get children accounts params
func (o *GetChildrenAccountsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetChildrenAccountsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get children accounts params
func (o *GetChildrenAccountsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get children accounts params
func (o *GetChildrenAccountsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetChildrenAccountsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get children accounts params
func (o *GetChildrenAccountsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get children accounts params
func (o *GetChildrenAccountsParams) WithAccountID(accountID strfmt.UUID) *GetChildrenAccountsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get children accounts params
func (o *GetChildrenAccountsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAccountWithBalance adds the accountWithBalance to the get children accounts params
func (o *GetChildrenAccountsParams) WithAccountWithBalance(accountWithBalance *bool) *GetChildrenAccountsParams {
	o.SetAccountWithBalance(accountWithBalance)
	return o
}

// SetAccountWithBalance adds the accountWithBalance to the get children accounts params
func (o *GetChildrenAccountsParams) SetAccountWithBalance(accountWithBalance *bool) {
	o.AccountWithBalance = accountWithBalance
}

// WithAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the get children accounts params
func (o *GetChildrenAccountsParams) WithAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) *GetChildrenAccountsParams {
	o.SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA)
	return o
}

// SetAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the get children accounts params
func (o *GetChildrenAccountsParams) SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) {
	o.AccountWithBalanceAndCBA = accountWithBalanceAndCBA
}

// WithAudit adds the audit to the get children accounts params
func (o *GetChildrenAccountsParams) WithAudit(audit *string) *GetChildrenAccountsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get children accounts params
func (o *GetChildrenAccountsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WriteToRequest writes these params to a swagger request
func (o *GetChildrenAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.AccountWithBalance != nil {

		// query param accountWithBalance
		var qrAccountWithBalance bool
		if o.AccountWithBalance != nil {
			qrAccountWithBalance = *o.AccountWithBalance
		}
		qAccountWithBalance := swag.FormatBool(qrAccountWithBalance)
		if qAccountWithBalance != "" {
			if err := r.SetQueryParam("accountWithBalance", qAccountWithBalance); err != nil {
				return err
			}
		}

	}

	if o.AccountWithBalanceAndCBA != nil {

		// query param accountWithBalanceAndCBA
		var qrAccountWithBalanceAndCBA bool
		if o.AccountWithBalanceAndCBA != nil {
			qrAccountWithBalanceAndCBA = *o.AccountWithBalanceAndCBA
		}
		qAccountWithBalanceAndCBA := swag.FormatBool(qrAccountWithBalanceAndCBA)
		if qAccountWithBalanceAndCBA != "" {
			if err := r.SetQueryParam("accountWithBalanceAndCBA", qAccountWithBalanceAndCBA); err != nil {
				return err
			}
		}

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
