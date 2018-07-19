// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default API Key. If not set explicitly in params, this will be used.
	XKillbillAPIKey() *string
	// Default API Secret. If not set explicitly in params, this will be used.
	XKillbillAPISecret() *string
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// ICatalog - interface for Catalog client.
type ICatalog interface {
	/*
		AddSimplePlan adds a simple plan entry in the current version of the catalog
	*/
	AddSimplePlan(ctx context.Context, params *AddSimplePlanParams) (*AddSimplePlanCreated, error)

	/*
		DeleteCatalog deletes all versions for a per tenant catalog
	*/
	DeleteCatalog(ctx context.Context, params *DeleteCatalogParams) (*DeleteCatalogNoContent, error)

	/*
		GetAvailableAddons retrieves available add ons for a given product
	*/
	GetAvailableAddons(ctx context.Context, params *GetAvailableAddonsParams) (*GetAvailableAddonsOK, error)

	/*
		GetAvailableBasePlans retrieves available base plans
	*/
	GetAvailableBasePlans(ctx context.Context, params *GetAvailableBasePlansParams) (*GetAvailableBasePlansOK, error)

	/*
		GetCatalogJSON retrieves the catalog as JSON
	*/
	GetCatalogJSON(ctx context.Context, params *GetCatalogJSONParams) (*GetCatalogJSONOK, error)

	/*
		GetCatalogVersions retrieves a list of catalog versions
	*/
	GetCatalogVersions(ctx context.Context, params *GetCatalogVersionsParams) (*GetCatalogVersionsOK, error)

	/*
		GetCatalogXML retrieves the full catalog as XML
	*/
	GetCatalogXML(ctx context.Context, params *GetCatalogXMLParams) (*GetCatalogXMLOK, error)

	/*
		GetPhaseForSubscriptionAndDate retrieves phase for a given subscription and date
	*/
	GetPhaseForSubscriptionAndDate(ctx context.Context, params *GetPhaseForSubscriptionAndDateParams) (*GetPhaseForSubscriptionAndDateOK, error)

	/*
		GetPlanForSubscriptionAndDate retrieves plan for a given subscription and date
	*/
	GetPlanForSubscriptionAndDate(ctx context.Context, params *GetPlanForSubscriptionAndDateParams) (*GetPlanForSubscriptionAndDateOK, error)

	/*
		GetPriceListForSubscriptionAndDate retrieves price list for a given subscription and date
	*/
	GetPriceListForSubscriptionAndDate(ctx context.Context, params *GetPriceListForSubscriptionAndDateParams) (*GetPriceListForSubscriptionAndDateOK, error)

	/*
		GetProductForSubscriptionAndDate retrieves product for a given subscription and date
	*/
	GetProductForSubscriptionAndDate(ctx context.Context, params *GetProductForSubscriptionAndDateParams) (*GetProductForSubscriptionAndDateOK, error)

	/*
		UploadCatalogXML uploads the full catalog as XML
	*/
	UploadCatalogXML(ctx context.Context, params *UploadCatalogXMLParams) (*UploadCatalogXMLCreated, error)
}

/*
AddSimplePlan adds a simple plan entry in the current version of the catalog
*/
func (a *Client) AddSimplePlan(ctx context.Context, params *AddSimplePlanParams) (*AddSimplePlanCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSimplePlanParams()
	}
	getParams := NewAddSimplePlanParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSimplePlan",
		Method:             "POST",
		PathPattern:        "/1.0/kb/catalog/simplePlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddSimplePlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*AddSimplePlanCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addSimplePlan",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &AddSimplePlanReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*AddSimplePlanCreated), nil

}

/*
DeleteCatalog deletes all versions for a per tenant catalog
*/
func (a *Client) DeleteCatalog(ctx context.Context, params *DeleteCatalogParams) (*DeleteCatalogNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCatalogParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCatalog",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/catalog",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCatalogReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCatalogNoContent), nil

}

/*
GetAvailableAddons retrieves available add ons for a given product
*/
func (a *Client) GetAvailableAddons(ctx context.Context, params *GetAvailableAddonsParams) (*GetAvailableAddonsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAvailableAddonsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAvailableAddons",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/availableAddons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAvailableAddonsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAvailableAddonsOK), nil

}

/*
GetAvailableBasePlans retrieves available base plans
*/
func (a *Client) GetAvailableBasePlans(ctx context.Context, params *GetAvailableBasePlansParams) (*GetAvailableBasePlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAvailableBasePlansParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAvailableBasePlans",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/availableBasePlans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAvailableBasePlansReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAvailableBasePlansOK), nil

}

/*
GetCatalogJSON retrieves the catalog as JSON
*/
func (a *Client) GetCatalogJSON(ctx context.Context, params *GetCatalogJSONParams) (*GetCatalogJSONOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogJSONParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogJson",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogJSONReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogJSONOK), nil

}

/*
GetCatalogVersions retrieves a list of catalog versions
*/
func (a *Client) GetCatalogVersions(ctx context.Context, params *GetCatalogVersionsParams) (*GetCatalogVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogVersionsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogVersions",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogVersionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogVersionsOK), nil

}

/*
GetCatalogXML retrieves the full catalog as XML
*/
func (a *Client) GetCatalogXML(ctx context.Context, params *GetCatalogXMLParams) (*GetCatalogXMLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogXMLParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogXml",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/xml",
		ProducesMediaTypes: []string{"text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCatalogXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogXMLOK), nil

}

/*
GetPhaseForSubscriptionAndDate retrieves phase for a given subscription and date
*/
func (a *Client) GetPhaseForSubscriptionAndDate(ctx context.Context, params *GetPhaseForSubscriptionAndDateParams) (*GetPhaseForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPhaseForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPhaseForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/phase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPhaseForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPhaseForSubscriptionAndDateOK), nil

}

/*
GetPlanForSubscriptionAndDate retrieves plan for a given subscription and date
*/
func (a *Client) GetPlanForSubscriptionAndDate(ctx context.Context, params *GetPlanForSubscriptionAndDateParams) (*GetPlanForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlanForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPlanForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/plan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPlanForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPlanForSubscriptionAndDateOK), nil

}

/*
GetPriceListForSubscriptionAndDate retrieves price list for a given subscription and date
*/
func (a *Client) GetPriceListForSubscriptionAndDate(ctx context.Context, params *GetPriceListForSubscriptionAndDateParams) (*GetPriceListForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPriceListForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPriceListForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/priceList",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPriceListForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPriceListForSubscriptionAndDateOK), nil

}

/*
GetProductForSubscriptionAndDate retrieves product for a given subscription and date
*/
func (a *Client) GetProductForSubscriptionAndDate(ctx context.Context, params *GetProductForSubscriptionAndDateParams) (*GetProductForSubscriptionAndDateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProductForSubscriptionAndDateParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProductForSubscriptionAndDate",
		Method:             "GET",
		PathPattern:        "/1.0/kb/catalog/product",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProductForSubscriptionAndDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProductForSubscriptionAndDateOK), nil

}

/*
UploadCatalogXML uploads the full catalog as XML
*/
func (a *Client) UploadCatalogXML(ctx context.Context, params *UploadCatalogXMLParams) (*UploadCatalogXMLCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadCatalogXMLParams()
	}
	getParams := NewUploadCatalogXMLParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}
	getParams.XKillbillAPIKey = params.XKillbillAPIKey

	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}
	getParams.XKillbillAPISecret = params.XKillbillAPISecret

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadCatalogXml",
		Method:             "POST",
		PathPattern:        "/1.0/kb/catalog/xml",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadCatalogXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*UploadCatalogXMLCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadCatalogXml",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"text/xml"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &UploadCatalogXMLReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*UploadCatalogXMLCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
