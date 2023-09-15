// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a TargetHttpsProxy resource, which is used by one or more
// global forwarding rule to route incoming HTTPS requests to a URL map.
//
// To get more information about TargetHttpsProxy, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
//
// ## Example Usage
//
// ## Import
//
// # TargetHttpsProxy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/targetHttpsProxy:TargetHttpsProxy default projects/{{project}}/global/targetHttpsProxies/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetHttpsProxy:TargetHttpsProxy default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetHttpsProxy:TargetHttpsProxy default {{name}}
//
// ```
type TargetHttpsProxy struct {
	pulumi.CustomResourceState

	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	CertificateMap pulumi.StringPtrOutput `pulumi:"certificateMap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HttpKeepAliveTimeoutSec pulumi.IntPtrOutput `pulumi:"httpKeepAliveTimeoutSec"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// The unique identifier for the resource.
	ProxyId pulumi.IntOutput `pulumi:"proxyId"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is `NONE`.
	// Possible values are: `NONE`, `ENABLE`, `DISABLE`.
	QuicOverride pulumi.StringPtrOutput `pulumi:"quicOverride"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of SslCertificate resource URLs or Certificate Manager certificate URLs that are used to authenticate
	// connections between users and the load balancer. At least one resource must be specified.
	SslCertificates pulumi.StringArrayOutput `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrOutput `pulumi:"sslPolicy"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	//
	// ***
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewTargetHttpsProxy registers a new resource with the given unique name, arguments, and options.
func NewTargetHttpsProxy(ctx *pulumi.Context,
	name string, args *TargetHttpsProxyArgs, opts ...pulumi.ResourceOption) (*TargetHttpsProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UrlMap == nil {
		return nil, errors.New("invalid value for required argument 'UrlMap'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TargetHttpsProxy
	err := ctx.RegisterResource("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetHttpsProxy gets an existing TargetHttpsProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetHttpsProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetHttpsProxyState, opts ...pulumi.ResourceOption) (*TargetHttpsProxy, error) {
	var resource TargetHttpsProxy
	err := ctx.ReadResource("gcp:compute/targetHttpsProxy:TargetHttpsProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetHttpsProxy resources.
type targetHttpsProxyState struct {
	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	CertificateMap *string `pulumi:"certificateMap"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HttpKeepAliveTimeoutSec *int `pulumi:"httpKeepAliveTimeoutSec"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// The unique identifier for the resource.
	ProxyId *int `pulumi:"proxyId"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is `NONE`.
	// Possible values are: `NONE`, `ENABLE`, `DISABLE`.
	QuicOverride *string `pulumi:"quicOverride"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of SslCertificate resource URLs or Certificate Manager certificate URLs that are used to authenticate
	// connections between users and the load balancer. At least one resource must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	//
	// ***
	UrlMap *string `pulumi:"urlMap"`
}

type TargetHttpsProxyState struct {
	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	CertificateMap pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HttpKeepAliveTimeoutSec pulumi.IntPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.IntPtrInput
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is `NONE`.
	// Possible values are: `NONE`, `ENABLE`, `DISABLE`.
	QuicOverride pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A list of SslCertificate resource URLs or Certificate Manager certificate URLs that are used to authenticate
	// connections between users and the load balancer. At least one resource must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	//
	// ***
	UrlMap pulumi.StringPtrInput
}

func (TargetHttpsProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpsProxyState)(nil)).Elem()
}

type targetHttpsProxyArgs struct {
	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	CertificateMap *string `pulumi:"certificateMap"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HttpKeepAliveTimeoutSec *int `pulumi:"httpKeepAliveTimeoutSec"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is `NONE`.
	// Possible values are: `NONE`, `ENABLE`, `DISABLE`.
	QuicOverride *string `pulumi:"quicOverride"`
	// A list of SslCertificate resource URLs or Certificate Manager certificate URLs that are used to authenticate
	// connections between users and the load balancer. At least one resource must be specified.
	SslCertificates []string `pulumi:"sslCertificates"`
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy *string `pulumi:"sslPolicy"`
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	//
	// ***
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a TargetHttpsProxy resource.
type TargetHttpsProxyArgs struct {
	// A reference to the CertificateMap resource uri that identifies a certificate map
	// associated with the given target proxy. This field can only be set for global target proxies.
	// Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
	CertificateMap pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	HttpKeepAliveTimeoutSec pulumi.IntPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used.
	// Default value is `NONE`.
	// Possible values are: `NONE`, `ENABLE`, `DISABLE`.
	QuicOverride pulumi.StringPtrInput
	// A list of SslCertificate resource URLs or Certificate Manager certificate URLs that are used to authenticate
	// connections between users and the load balancer. At least one resource must be specified.
	SslCertificates pulumi.StringArrayInput
	// A reference to the SslPolicy resource that will be associated with
	// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
	// resource will not have any SSL policy configured.
	SslPolicy pulumi.StringPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL
	// to the BackendService.
	//
	// ***
	UrlMap pulumi.StringInput
}

func (TargetHttpsProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetHttpsProxyArgs)(nil)).Elem()
}

type TargetHttpsProxyInput interface {
	pulumi.Input

	ToTargetHttpsProxyOutput() TargetHttpsProxyOutput
	ToTargetHttpsProxyOutputWithContext(ctx context.Context) TargetHttpsProxyOutput
}

func (*TargetHttpsProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpsProxy)(nil)).Elem()
}

func (i *TargetHttpsProxy) ToTargetHttpsProxyOutput() TargetHttpsProxyOutput {
	return i.ToTargetHttpsProxyOutputWithContext(context.Background())
}

func (i *TargetHttpsProxy) ToTargetHttpsProxyOutputWithContext(ctx context.Context) TargetHttpsProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyOutput)
}

// TargetHttpsProxyArrayInput is an input type that accepts TargetHttpsProxyArray and TargetHttpsProxyArrayOutput values.
// You can construct a concrete instance of `TargetHttpsProxyArrayInput` via:
//
//	TargetHttpsProxyArray{ TargetHttpsProxyArgs{...} }
type TargetHttpsProxyArrayInput interface {
	pulumi.Input

	ToTargetHttpsProxyArrayOutput() TargetHttpsProxyArrayOutput
	ToTargetHttpsProxyArrayOutputWithContext(context.Context) TargetHttpsProxyArrayOutput
}

type TargetHttpsProxyArray []TargetHttpsProxyInput

func (TargetHttpsProxyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetHttpsProxy)(nil)).Elem()
}

func (i TargetHttpsProxyArray) ToTargetHttpsProxyArrayOutput() TargetHttpsProxyArrayOutput {
	return i.ToTargetHttpsProxyArrayOutputWithContext(context.Background())
}

func (i TargetHttpsProxyArray) ToTargetHttpsProxyArrayOutputWithContext(ctx context.Context) TargetHttpsProxyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyArrayOutput)
}

// TargetHttpsProxyMapInput is an input type that accepts TargetHttpsProxyMap and TargetHttpsProxyMapOutput values.
// You can construct a concrete instance of `TargetHttpsProxyMapInput` via:
//
//	TargetHttpsProxyMap{ "key": TargetHttpsProxyArgs{...} }
type TargetHttpsProxyMapInput interface {
	pulumi.Input

	ToTargetHttpsProxyMapOutput() TargetHttpsProxyMapOutput
	ToTargetHttpsProxyMapOutputWithContext(context.Context) TargetHttpsProxyMapOutput
}

type TargetHttpsProxyMap map[string]TargetHttpsProxyInput

func (TargetHttpsProxyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetHttpsProxy)(nil)).Elem()
}

func (i TargetHttpsProxyMap) ToTargetHttpsProxyMapOutput() TargetHttpsProxyMapOutput {
	return i.ToTargetHttpsProxyMapOutputWithContext(context.Background())
}

func (i TargetHttpsProxyMap) ToTargetHttpsProxyMapOutputWithContext(ctx context.Context) TargetHttpsProxyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetHttpsProxyMapOutput)
}

type TargetHttpsProxyOutput struct{ *pulumi.OutputState }

func (TargetHttpsProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetHttpsProxy)(nil)).Elem()
}

func (o TargetHttpsProxyOutput) ToTargetHttpsProxyOutput() TargetHttpsProxyOutput {
	return o
}

func (o TargetHttpsProxyOutput) ToTargetHttpsProxyOutputWithContext(ctx context.Context) TargetHttpsProxyOutput {
	return o
}

// A reference to the CertificateMap resource uri that identifies a certificate map
// associated with the given target proxy. This field can only be set for global target proxies.
// Accepted format is `//certificatemanager.googleapis.com/projects/{project}/locations/{location}/certificateMaps/{resourceName}`.
func (o TargetHttpsProxyOutput) CertificateMap() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringPtrOutput { return v.CertificateMap }).(pulumi.StringPtrOutput)
}

// Creation timestamp in RFC3339 text format.
func (o TargetHttpsProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o TargetHttpsProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies how long to keep a connection open, after completing a response,
// while there is no matching traffic (in seconds). If an HTTP keepalive is
// not specified, a default value (610 seconds) will be used. For Global
// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
// load balancer (classic), this option is not available publicly.
func (o TargetHttpsProxyOutput) HttpKeepAliveTimeoutSec() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.IntPtrOutput { return v.HttpKeepAliveTimeoutSec }).(pulumi.IntPtrOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
func (o TargetHttpsProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TargetHttpsProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references
// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
func (o TargetHttpsProxyOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.BoolOutput { return v.ProxyBind }).(pulumi.BoolOutput)
}

// The unique identifier for the resource.
func (o TargetHttpsProxyOutput) ProxyId() pulumi.IntOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.IntOutput { return v.ProxyId }).(pulumi.IntOutput)
}

// Specifies the QUIC override policy for this resource. This determines
// whether the load balancer will attempt to negotiate QUIC with clients
// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
// specified, Google manages whether QUIC is used.
// Default value is `NONE`.
// Possible values are: `NONE`, `ENABLE`, `DISABLE`.
func (o TargetHttpsProxyOutput) QuicOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringPtrOutput { return v.QuicOverride }).(pulumi.StringPtrOutput)
}

// The URI of the created resource.
func (o TargetHttpsProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// A list of SslCertificate resource URLs or Certificate Manager certificate URLs that are used to authenticate
// connections between users and the load balancer. At least one resource must be specified.
func (o TargetHttpsProxyOutput) SslCertificates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringArrayOutput { return v.SslCertificates }).(pulumi.StringArrayOutput)
}

// A reference to the SslPolicy resource that will be associated with
// the TargetHttpsProxy resource. If not set, the TargetHttpsProxy
// resource will not have any SSL policy configured.
func (o TargetHttpsProxyOutput) SslPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringPtrOutput { return v.SslPolicy }).(pulumi.StringPtrOutput)
}

// A reference to the UrlMap resource that defines the mapping from URL
// to the BackendService.
//
// ***
func (o TargetHttpsProxyOutput) UrlMap() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetHttpsProxy) pulumi.StringOutput { return v.UrlMap }).(pulumi.StringOutput)
}

type TargetHttpsProxyArrayOutput struct{ *pulumi.OutputState }

func (TargetHttpsProxyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetHttpsProxy)(nil)).Elem()
}

func (o TargetHttpsProxyArrayOutput) ToTargetHttpsProxyArrayOutput() TargetHttpsProxyArrayOutput {
	return o
}

func (o TargetHttpsProxyArrayOutput) ToTargetHttpsProxyArrayOutputWithContext(ctx context.Context) TargetHttpsProxyArrayOutput {
	return o
}

func (o TargetHttpsProxyArrayOutput) Index(i pulumi.IntInput) TargetHttpsProxyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TargetHttpsProxy {
		return vs[0].([]*TargetHttpsProxy)[vs[1].(int)]
	}).(TargetHttpsProxyOutput)
}

type TargetHttpsProxyMapOutput struct{ *pulumi.OutputState }

func (TargetHttpsProxyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetHttpsProxy)(nil)).Elem()
}

func (o TargetHttpsProxyMapOutput) ToTargetHttpsProxyMapOutput() TargetHttpsProxyMapOutput {
	return o
}

func (o TargetHttpsProxyMapOutput) ToTargetHttpsProxyMapOutputWithContext(ctx context.Context) TargetHttpsProxyMapOutput {
	return o
}

func (o TargetHttpsProxyMapOutput) MapIndex(k pulumi.StringInput) TargetHttpsProxyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TargetHttpsProxy {
		return vs[0].(map[string]*TargetHttpsProxy)[vs[1].(string)]
	}).(TargetHttpsProxyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpsProxyInput)(nil)).Elem(), &TargetHttpsProxy{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpsProxyArrayInput)(nil)).Elem(), TargetHttpsProxyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetHttpsProxyMapInput)(nil)).Elem(), TargetHttpsProxyMap{})
	pulumi.RegisterOutputType(TargetHttpsProxyOutput{})
	pulumi.RegisterOutputType(TargetHttpsProxyArrayOutput{})
	pulumi.RegisterOutputType(TargetHttpsProxyMapOutput{})
}