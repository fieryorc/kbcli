// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// NewDeleteBundleTagsParams creates a new DeleteBundleTagsParams object
// with the default values initialized.
func NewDeleteBundleTagsParams() *DeleteBundleTagsParams {
	var ()
	return &DeleteBundleTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBundleTagsParamsWithTimeout creates a new DeleteBundleTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBundleTagsParamsWithTimeout(timeout time.Duration) *DeleteBundleTagsParams {
	var ()
	return &DeleteBundleTagsParams{

		timeout: timeout,
	}
}

// NewDeleteBundleTagsParamsWithContext creates a new DeleteBundleTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBundleTagsParamsWithContext(ctx context.Context) *DeleteBundleTagsParams {
	var ()
	return &DeleteBundleTagsParams{

		Context: ctx,
	}
}

// NewDeleteBundleTagsParamsWithHTTPClient creates a new DeleteBundleTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBundleTagsParamsWithHTTPClient(client *http.Client) *DeleteBundleTagsParams {
	var ()
	return &DeleteBundleTagsParams{
		HTTPClient: client,
	}
}

/*DeleteBundleTagsParams contains all the parameters to send to the API endpoint
for the delete bundle tags operation typically these are written to a http.Request
*/
type DeleteBundleTagsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*BundleID*/
	BundleID strfmt.UUID
	/*TagDef*/
	TagDef []strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithTimeout(timeout time.Duration) *DeleteBundleTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithContext(ctx context.Context) *DeleteBundleTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithHTTPClient(client *http.Client) *DeleteBundleTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithXKillbillComment(xKillbillComment *string) *DeleteBundleTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteBundleTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithXKillbillReason(xKillbillReason *string) *DeleteBundleTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBundleID adds the bundleID to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithBundleID(bundleID strfmt.UUID) *DeleteBundleTagsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithTagDef adds the tagDef to the delete bundle tags params
func (o *DeleteBundleTagsParams) WithTagDef(tagDef []strfmt.UUID) *DeleteBundleTagsParams {
	o.SetTagDef(tagDef)
	return o
}

// SetTagDef adds the tagDef to the delete bundle tags params
func (o *DeleteBundleTagsParams) SetTagDef(tagDef []strfmt.UUID) {
	o.TagDef = tagDef
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBundleTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	var valuesTagDef []string
	for _, v := range o.TagDef {
		valuesTagDef = append(valuesTagDef, v.String())
	}

	joinedTagDef := swag.JoinByFormat(valuesTagDef, "multi")
	// query array param tagDef
	if err := r.SetQueryParam("tagDef", joinedTagDef...); err != nil {
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
