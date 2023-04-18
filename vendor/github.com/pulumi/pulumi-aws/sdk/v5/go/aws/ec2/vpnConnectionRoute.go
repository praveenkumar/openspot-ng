// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a static route between a VPN connection and a customer gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			vpnGateway, err := ec2.NewVpnGateway(ctx, "vpnGateway", &ec2.VpnGatewayArgs{
//				VpcId: vpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			customerGateway, err := ec2.NewCustomerGateway(ctx, "customerGateway", &ec2.CustomerGatewayArgs{
//				BgpAsn:    pulumi.String("65000"),
//				IpAddress: pulumi.String("172.0.0.1"),
//				Type:      pulumi.String("ipsec.1"),
//			})
//			if err != nil {
//				return err
//			}
//			main, err := ec2.NewVpnConnection(ctx, "main", &ec2.VpnConnectionArgs{
//				VpnGatewayId:      vpnGateway.ID(),
//				CustomerGatewayId: customerGateway.ID(),
//				Type:              pulumi.String("ipsec.1"),
//				StaticRoutesOnly:  pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpnConnectionRoute(ctx, "office", &ec2.VpnConnectionRouteArgs{
//				DestinationCidrBlock: pulumi.String("192.168.10.0/24"),
//				VpnConnectionId:      main.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VpnConnectionRoute struct {
	pulumi.CustomResourceState

	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId pulumi.StringOutput `pulumi:"vpnConnectionId"`
}

// NewVpnConnectionRoute registers a new resource with the given unique name, arguments, and options.
func NewVpnConnectionRoute(ctx *pulumi.Context,
	name string, args *VpnConnectionRouteArgs, opts ...pulumi.ResourceOption) (*VpnConnectionRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.VpnConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'VpnConnectionId'")
	}
	var resource VpnConnectionRoute
	err := ctx.RegisterResource("aws:ec2/vpnConnectionRoute:VpnConnectionRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnConnectionRoute gets an existing VpnConnectionRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnConnectionRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnConnectionRouteState, opts ...pulumi.ResourceOption) (*VpnConnectionRoute, error) {
	var resource VpnConnectionRoute
	err := ctx.ReadResource("aws:ec2/vpnConnectionRoute:VpnConnectionRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnConnectionRoute resources.
type vpnConnectionRouteState struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
}

type VpnConnectionRouteState struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock pulumi.StringPtrInput
	// The ID of the VPN connection.
	VpnConnectionId pulumi.StringPtrInput
}

func (VpnConnectionRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionRouteState)(nil)).Elem()
}

type vpnConnectionRouteArgs struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId string `pulumi:"vpnConnectionId"`
}

// The set of arguments for constructing a VpnConnectionRoute resource.
type VpnConnectionRouteArgs struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock pulumi.StringInput
	// The ID of the VPN connection.
	VpnConnectionId pulumi.StringInput
}

func (VpnConnectionRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionRouteArgs)(nil)).Elem()
}

type VpnConnectionRouteInput interface {
	pulumi.Input

	ToVpnConnectionRouteOutput() VpnConnectionRouteOutput
	ToVpnConnectionRouteOutputWithContext(ctx context.Context) VpnConnectionRouteOutput
}

func (*VpnConnectionRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnConnectionRoute)(nil)).Elem()
}

func (i *VpnConnectionRoute) ToVpnConnectionRouteOutput() VpnConnectionRouteOutput {
	return i.ToVpnConnectionRouteOutputWithContext(context.Background())
}

func (i *VpnConnectionRoute) ToVpnConnectionRouteOutputWithContext(ctx context.Context) VpnConnectionRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionRouteOutput)
}

// VpnConnectionRouteArrayInput is an input type that accepts VpnConnectionRouteArray and VpnConnectionRouteArrayOutput values.
// You can construct a concrete instance of `VpnConnectionRouteArrayInput` via:
//
//	VpnConnectionRouteArray{ VpnConnectionRouteArgs{...} }
type VpnConnectionRouteArrayInput interface {
	pulumi.Input

	ToVpnConnectionRouteArrayOutput() VpnConnectionRouteArrayOutput
	ToVpnConnectionRouteArrayOutputWithContext(context.Context) VpnConnectionRouteArrayOutput
}

type VpnConnectionRouteArray []VpnConnectionRouteInput

func (VpnConnectionRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnConnectionRoute)(nil)).Elem()
}

func (i VpnConnectionRouteArray) ToVpnConnectionRouteArrayOutput() VpnConnectionRouteArrayOutput {
	return i.ToVpnConnectionRouteArrayOutputWithContext(context.Background())
}

func (i VpnConnectionRouteArray) ToVpnConnectionRouteArrayOutputWithContext(ctx context.Context) VpnConnectionRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionRouteArrayOutput)
}

// VpnConnectionRouteMapInput is an input type that accepts VpnConnectionRouteMap and VpnConnectionRouteMapOutput values.
// You can construct a concrete instance of `VpnConnectionRouteMapInput` via:
//
//	VpnConnectionRouteMap{ "key": VpnConnectionRouteArgs{...} }
type VpnConnectionRouteMapInput interface {
	pulumi.Input

	ToVpnConnectionRouteMapOutput() VpnConnectionRouteMapOutput
	ToVpnConnectionRouteMapOutputWithContext(context.Context) VpnConnectionRouteMapOutput
}

type VpnConnectionRouteMap map[string]VpnConnectionRouteInput

func (VpnConnectionRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnConnectionRoute)(nil)).Elem()
}

func (i VpnConnectionRouteMap) ToVpnConnectionRouteMapOutput() VpnConnectionRouteMapOutput {
	return i.ToVpnConnectionRouteMapOutputWithContext(context.Background())
}

func (i VpnConnectionRouteMap) ToVpnConnectionRouteMapOutputWithContext(ctx context.Context) VpnConnectionRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionRouteMapOutput)
}

type VpnConnectionRouteOutput struct{ *pulumi.OutputState }

func (VpnConnectionRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnConnectionRoute)(nil)).Elem()
}

func (o VpnConnectionRouteOutput) ToVpnConnectionRouteOutput() VpnConnectionRouteOutput {
	return o
}

func (o VpnConnectionRouteOutput) ToVpnConnectionRouteOutputWithContext(ctx context.Context) VpnConnectionRouteOutput {
	return o
}

// The CIDR block associated with the local subnet of the customer network.
func (o VpnConnectionRouteOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnectionRoute) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// The ID of the VPN connection.
func (o VpnConnectionRouteOutput) VpnConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnConnectionRoute) pulumi.StringOutput { return v.VpnConnectionId }).(pulumi.StringOutput)
}

type VpnConnectionRouteArrayOutput struct{ *pulumi.OutputState }

func (VpnConnectionRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnConnectionRoute)(nil)).Elem()
}

func (o VpnConnectionRouteArrayOutput) ToVpnConnectionRouteArrayOutput() VpnConnectionRouteArrayOutput {
	return o
}

func (o VpnConnectionRouteArrayOutput) ToVpnConnectionRouteArrayOutputWithContext(ctx context.Context) VpnConnectionRouteArrayOutput {
	return o
}

func (o VpnConnectionRouteArrayOutput) Index(i pulumi.IntInput) VpnConnectionRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnConnectionRoute {
		return vs[0].([]*VpnConnectionRoute)[vs[1].(int)]
	}).(VpnConnectionRouteOutput)
}

type VpnConnectionRouteMapOutput struct{ *pulumi.OutputState }

func (VpnConnectionRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnConnectionRoute)(nil)).Elem()
}

func (o VpnConnectionRouteMapOutput) ToVpnConnectionRouteMapOutput() VpnConnectionRouteMapOutput {
	return o
}

func (o VpnConnectionRouteMapOutput) ToVpnConnectionRouteMapOutputWithContext(ctx context.Context) VpnConnectionRouteMapOutput {
	return o
}

func (o VpnConnectionRouteMapOutput) MapIndex(k pulumi.StringInput) VpnConnectionRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnConnectionRoute {
		return vs[0].(map[string]*VpnConnectionRoute)[vs[1].(string)]
	}).(VpnConnectionRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnConnectionRouteInput)(nil)).Elem(), &VpnConnectionRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnConnectionRouteArrayInput)(nil)).Elem(), VpnConnectionRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnConnectionRouteMapInput)(nil)).Elem(), VpnConnectionRouteMap{})
	pulumi.RegisterOutputType(VpnConnectionRouteOutput{})
	pulumi.RegisterOutputType(VpnConnectionRouteArrayOutput{})
	pulumi.RegisterOutputType(VpnConnectionRouteMapOutput{})
}
