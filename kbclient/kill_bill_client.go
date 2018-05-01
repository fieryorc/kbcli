// Code generated by go-swagger; DO NOT EDIT.

package kbclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbclient/admin"
	"github.com/killbill/kbcli/kbclient/bundle"
	"github.com/killbill/kbcli/kbclient/catalog"
	"github.com/killbill/kbcli/kbclient/credit"
	"github.com/killbill/kbcli/kbclient/custom_field"
	"github.com/killbill/kbcli/kbclient/export"
	"github.com/killbill/kbcli/kbclient/invoice"
	"github.com/killbill/kbcli/kbclient/invoice_item"
	"github.com/killbill/kbcli/kbclient/invoice_payment"
	"github.com/killbill/kbcli/kbclient/nodes_info"
	"github.com/killbill/kbcli/kbclient/overdue"
	"github.com/killbill/kbcli/kbclient/payment"
	"github.com/killbill/kbcli/kbclient/payment_gateway"
	"github.com/killbill/kbcli/kbclient/payment_method"
	"github.com/killbill/kbcli/kbclient/payment_transaction"
	"github.com/killbill/kbcli/kbclient/plugin_info"
	"github.com/killbill/kbcli/kbclient/security"
	"github.com/killbill/kbcli/kbclient/subscription"
	"github.com/killbill/kbcli/kbclient/tag"
	"github.com/killbill/kbcli/kbclient/tag_definition"
	"github.com/killbill/kbcli/kbclient/tenant"
	"github.com/killbill/kbcli/kbclient/usage"
)

// New creates a new kill bill client
// The following snippet provides creating killbill client with basic auth.
//
//
// 	   trp := httptransport.New("127.0.0.1:8080" /*host*/, "" /*basePath*/, nil /*schemes*/)
// 	   // Add missing handler. OpenAPI runtime doesn't have this by default
// 	   trp.Producers["text/xml"] = runtime.TextProducer()
// 	   // Set tro true to print http/debug logs
// 	   trp.Debug = enableDebug
// 	   // Setup basic auth
// 	   authWriter := httptransport.BasicAuth("admin", "password")
// 	   client := kbclient.New(trp, strfmt.Default)
//     // Use client
//     client.Accounts.GetAccount(...)
//
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *KillBill {

	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := &KillBill{
		Transport: transport,
		defaults:  defaultsImpl{KillbillDefaults: defaults},
	}

	cli.Account = account.New(transport, formats, authInfo, &cli.defaults)

	cli.Admin = admin.New(transport, formats, authInfo, &cli.defaults)

	cli.Bundle = bundle.New(transport, formats, authInfo, &cli.defaults)

	cli.Catalog = catalog.New(transport, formats, authInfo, &cli.defaults)

	cli.Credit = credit.New(transport, formats, authInfo, &cli.defaults)

	cli.CustomField = custom_field.New(transport, formats, authInfo, &cli.defaults)

	cli.Export = export.New(transport, formats, authInfo, &cli.defaults)

	cli.Invoice = invoice.New(transport, formats, authInfo, &cli.defaults)

	cli.InvoiceItem = invoice_item.New(transport, formats, authInfo, &cli.defaults)

	cli.InvoicePayment = invoice_payment.New(transport, formats, authInfo, &cli.defaults)

	cli.NodesInfo = nodes_info.New(transport, formats, authInfo, &cli.defaults)

	cli.Overdue = overdue.New(transport, formats, authInfo, &cli.defaults)

	cli.Payment = payment.New(transport, formats, authInfo, &cli.defaults)

	cli.PaymentGateway = payment_gateway.New(transport, formats, authInfo, &cli.defaults)

	cli.PaymentMethod = payment_method.New(transport, formats, authInfo, &cli.defaults)

	cli.PaymentTransaction = payment_transaction.New(transport, formats, authInfo, &cli.defaults)

	cli.PluginInfo = plugin_info.New(transport, formats, authInfo, &cli.defaults)

	cli.Security = security.New(transport, formats, authInfo, &cli.defaults)

	cli.Subscription = subscription.New(transport, formats, authInfo, &cli.defaults)

	cli.Tag = tag.New(transport, formats, authInfo, &cli.defaults)

	cli.TagDefinition = tag_definition.New(transport, formats, authInfo, &cli.defaults)

	cli.Tenant = tenant.New(transport, formats, authInfo, &cli.defaults)

	cli.Usage = usage.New(transport, formats, authInfo, &cli.defaults)

	return cli
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// KillBill is a client for kill bill
type KillBill struct {
	Account *account.Client

	Admin *admin.Client

	Bundle *bundle.Client

	Catalog *catalog.Client

	Credit *credit.Client

	CustomField *custom_field.Client

	Export *export.Client

	Invoice *invoice.Client

	InvoiceItem *invoice_item.Client

	InvoicePayment *invoice_payment.Client

	NodesInfo *nodes_info.Client

	Overdue *overdue.Client

	Payment *payment.Client

	PaymentGateway *payment_gateway.Client

	PaymentMethod *payment_method.Client

	PaymentTransaction *payment_transaction.Client

	PluginInfo *plugin_info.Client

	Security *security.Client

	Subscription *subscription.Client

	Tag *tag.Client

	TagDefinition *tag_definition.Client

	Tenant *tenant.Client

	Usage *usage.Client

	Transport runtime.ClientTransport
	defaults  defaultsImpl
}

// SetTransport changes the transport on the client and all its subresources
func (c *KillBill) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Account.SetTransport(transport)

	c.Admin.SetTransport(transport)

	c.Bundle.SetTransport(transport)

	c.Catalog.SetTransport(transport)

	c.Credit.SetTransport(transport)

	c.CustomField.SetTransport(transport)

	c.Export.SetTransport(transport)

	c.Invoice.SetTransport(transport)

	c.InvoiceItem.SetTransport(transport)

	c.InvoicePayment.SetTransport(transport)

	c.NodesInfo.SetTransport(transport)

	c.Overdue.SetTransport(transport)

	c.Payment.SetTransport(transport)

	c.PaymentGateway.SetTransport(transport)

	c.PaymentMethod.SetTransport(transport)

	c.PaymentTransaction.SetTransport(transport)

	c.PluginInfo.SetTransport(transport)

	c.Security.SetTransport(transport)

	c.Subscription.SetTransport(transport)

	c.Tag.SetTransport(transport)

	c.TagDefinition.SetTransport(transport)

	c.Tenant.SetTransport(transport)

	c.Usage.SetTransport(transport)

}

// Defaults returns killbill defaults
func (c *KillBill) Defaults() KillbillDefaults {
	return c.defaults.KillbillDefaults
}

// SetDefaults sets killbill defaults
func (c *KillBill) SetDefaults(defaults KillbillDefaults) {
	c.defaults = defaultsImpl{KillbillDefaults: defaults}
}

// Implements killbill default values.
type KillbillDefaults struct {
	// XKillbillAPIKey property
	APIKey *string
	// XKillbillAPISecret property
	APISecret *string
	// XKillbillCreatedBy property
	CreatedBy *string
	// XKillbillComment property
	Comment *string
	// XKillbillReason property
	Reason *string
	// withStackTrace property
	WithStackTrace *bool
}

type defaultsImpl struct {
	KillbillDefaults
}

// Default API Key. If not set explicitly in params, this will be used.
func (d *defaultsImpl) XKillbillAPIKey() *string {
	return d.APIKey
}

// Default API Secret. If not set explicitly in params, this will be used.
func (d *defaultsImpl) XKillbillAPISecret() *string {
	return d.APISecret
}

// Default CreatedBy. If not set explicitly in params, this will be used.
func (d *defaultsImpl) XKillbillCreatedBy() *string {
	return d.CreatedBy
}

// Default Comment. If not set explicitly in params, this will be used.
func (d *defaultsImpl) XKillbillComment() *string {
	return d.Comment
}

// Default Reason. If not set explicitly in params, this will be used.
func (d *defaultsImpl) XKillbillReason() *string {
	return d.Reason
}

// Default WithStackTrace. If not set explicitly in params, this will be used.
func (d *defaultsImpl) KillbillWithStackTrace() *bool {
	return d.WithStackTrace
}
