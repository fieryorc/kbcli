// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewUploadCatalogXMLParams creates a new UploadCatalogXMLParams object
// with the default values initialized.
func NewUploadCatalogXMLParams() *UploadCatalogXMLParams {
	var ()
	return &UploadCatalogXMLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadCatalogXMLParamsWithTimeout creates a new UploadCatalogXMLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadCatalogXMLParamsWithTimeout(timeout time.Duration) *UploadCatalogXMLParams {
	var ()
	return &UploadCatalogXMLParams{

		timeout: timeout,
	}
}

// NewUploadCatalogXMLParamsWithContext creates a new UploadCatalogXMLParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadCatalogXMLParamsWithContext(ctx context.Context) *UploadCatalogXMLParams {
	var ()
	return &UploadCatalogXMLParams{

		Context: ctx,
	}
}

// NewUploadCatalogXMLParamsWithHTTPClient creates a new UploadCatalogXMLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadCatalogXMLParamsWithHTTPClient(client *http.Client) *UploadCatalogXMLParams {
	var ()
	return &UploadCatalogXMLParams{
		HTTPClient: client,
	}
}

/*UploadCatalogXMLParams contains all the parameters to send to the API endpoint
for the upload catalog Xml operation typically these are written to a http.Request
*/
type UploadCatalogXMLParams struct {

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
	Body *string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithTimeout(timeout time.Duration) *UploadCatalogXMLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithContext(ctx context.Context) *UploadCatalogXMLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithHTTPClient(client *http.Client) *UploadCatalogXMLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithXKillbillAPIKey(xKillbillAPIKey string) *UploadCatalogXMLParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithXKillbillAPISecret(xKillbillAPISecret string) *UploadCatalogXMLParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithXKillbillComment(xKillbillComment *string) *UploadCatalogXMLParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadCatalogXMLParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithXKillbillReason(xKillbillReason *string) *UploadCatalogXMLParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload catalog Xml params
func (o *UploadCatalogXMLParams) WithBody(body *string) *UploadCatalogXMLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload catalog Xml params
func (o *UploadCatalogXMLParams) SetBody(body *string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UploadCatalogXMLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
