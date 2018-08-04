// Code generated by go-swagger; DO NOT EDIT.

package overdue

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

// NewGetOverdueConfigXMLParams creates a new GetOverdueConfigXMLParams object
// with the default values initialized.
func NewGetOverdueConfigXMLParams() *GetOverdueConfigXMLParams {

	return &GetOverdueConfigXMLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOverdueConfigXMLParamsWithTimeout creates a new GetOverdueConfigXMLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOverdueConfigXMLParamsWithTimeout(timeout time.Duration) *GetOverdueConfigXMLParams {

	return &GetOverdueConfigXMLParams{

		timeout: timeout,
	}
}

// NewGetOverdueConfigXMLParamsWithContext creates a new GetOverdueConfigXMLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOverdueConfigXMLParamsWithContext(ctx context.Context) *GetOverdueConfigXMLParams {

	return &GetOverdueConfigXMLParams{

		Context: ctx,
	}
}

// NewGetOverdueConfigXMLParamsWithHTTPClient creates a new GetOverdueConfigXMLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOverdueConfigXMLParamsWithHTTPClient(client *http.Client) *GetOverdueConfigXMLParams {

	return &GetOverdueConfigXMLParams{
		HTTPClient: client,
	}
}

/*GetOverdueConfigXMLParams contains all the parameters to send to the API endpoint
for the get overdue config Xml operation typically these are written to a http.Request
*/
type GetOverdueConfigXMLParams struct {
	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get overdue config Xml params
func (o *GetOverdueConfigXMLParams) WithTimeout(timeout time.Duration) *GetOverdueConfigXMLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get overdue config Xml params
func (o *GetOverdueConfigXMLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get overdue config Xml params
func (o *GetOverdueConfigXMLParams) WithContext(ctx context.Context) *GetOverdueConfigXMLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get overdue config Xml params
func (o *GetOverdueConfigXMLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get overdue config Xml params
func (o *GetOverdueConfigXMLParams) WithHTTPClient(client *http.Client) *GetOverdueConfigXMLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get overdue config Xml params
func (o *GetOverdueConfigXMLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOverdueConfigXMLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
