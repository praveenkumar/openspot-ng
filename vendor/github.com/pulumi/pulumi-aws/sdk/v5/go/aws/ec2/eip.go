// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Elastic IP resource.
//
// > **Note:** EIP may require IGW to exist prior to association. Use `dependsOn` to set an explicit dependency on the IGW.
//
// > **Note:** Do not use `networkInterface` to associate the EIP to `lb.LoadBalancer` or `ec2.NatGateway` resources. Instead use the `allocationId` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
//
// ## Example Usage
// ### Single EIP associated with an instance
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
//			_, err := ec2.NewEip(ctx, "lb", &ec2.EipArgs{
//				Instance: pulumi.Any(aws_instance.Web.Id),
//				Vpc:      pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Multiple EIPs associated with a single network interface
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
//			_, err := ec2.NewNetworkInterface(ctx, "multi-ip", &ec2.NetworkInterfaceArgs{
//				SubnetId: pulumi.Any(aws_subnet.Main.Id),
//				PrivateIps: pulumi.StringArray{
//					pulumi.String("10.0.0.10"),
//					pulumi.String("10.0.0.11"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewEip(ctx, "one", &ec2.EipArgs{
//				Vpc:                    pulumi.Bool(true),
//				NetworkInterface:       multi_ip.ID(),
//				AssociateWithPrivateIp: pulumi.String("10.0.0.10"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewEip(ctx, "two", &ec2.EipArgs{
//				Vpc:                    pulumi.Bool(true),
//				NetworkInterface:       multi_ip.ID(),
//				AssociateWithPrivateIp: pulumi.String("10.0.0.11"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Attaching an EIP to an Instance with a pre-assigned private ip (VPC Only)
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
//			_, err := ec2.NewVpc(ctx, "default", &ec2.VpcArgs{
//				CidrBlock:          pulumi.String("10.0.0.0/16"),
//				EnableDnsHostnames: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			gw, err := ec2.NewInternetGateway(ctx, "gw", &ec2.InternetGatewayArgs{
//				VpcId: _default.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			myTestSubnet, err := ec2.NewSubnet(ctx, "myTestSubnet", &ec2.SubnetArgs{
//				VpcId:               _default.ID(),
//				CidrBlock:           pulumi.String("10.0.0.0/24"),
//				MapPublicIpOnLaunch: pulumi.Bool(true),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				gw,
//			}))
//			if err != nil {
//				return err
//			}
//			foo, err := ec2.NewInstance(ctx, "foo", &ec2.InstanceArgs{
//				Ami:          pulumi.String("ami-5189a661"),
//				InstanceType: pulumi.String("t2.micro"),
//				PrivateIp:    pulumi.String("10.0.0.12"),
//				SubnetId:     myTestSubnet.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewEip(ctx, "bar", &ec2.EipArgs{
//				Vpc:                    pulumi.Bool(true),
//				Instance:               foo.ID(),
//				AssociateWithPrivateIp: pulumi.String("10.0.0.12"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				gw,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Allocating EIP from the BYOIP pool
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
//			_, err := ec2.NewEip(ctx, "byoip-ip", &ec2.EipArgs{
//				PublicIpv4Pool: pulumi.String("ipv4pool-ec2-012345"),
//				Vpc:            pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// EIPs in a VPC can be imported using their Allocation ID, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/eip:Eip bar eipalloc-00a10e96
//
// ```
//
//	EIPs in EC2-Classic can be imported using their Public IP, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/eip:Eip bar 52.0.0.0
//
// ```
type Eip struct {
	pulumi.CustomResourceState

	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIp pulumi.StringPtrOutput `pulumi:"associateWithPrivateIp"`
	// ID representing the association of the address with an instance in a VPC.
	AssociationId pulumi.StringOutput `pulumi:"associationId"`
	// Carrier IP address.
	CarrierIp pulumi.StringOutput `pulumi:"carrierIp"`
	// Customer owned IP.
	CustomerOwnedIp pulumi.StringOutput `pulumi:"customerOwnedIp"`
	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
	CustomerOwnedIpv4Pool pulumi.StringPtrOutput `pulumi:"customerOwnedIpv4Pool"`
	// Indicates if this EIP is for use in VPC (`vpc`) or EC2-Classic (`standard`).
	Domain pulumi.StringOutput `pulumi:"domain"`
	// EC2 instance ID.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup pulumi.StringOutput `pulumi:"networkBorderGroup"`
	// Network interface ID to associate with.
	NetworkInterface pulumi.StringOutput `pulumi:"networkInterface"`
	// The Private DNS associated with the Elastic IP address (if in VPC).
	PrivateDns pulumi.StringOutput `pulumi:"privateDns"`
	// Contains the private IP address (if in VPC).
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// Public DNS associated with the Elastic IP address.
	PublicDns pulumi.StringOutput `pulumi:"publicDns"`
	// Contains the public IP address.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// EC2 IPv4 address pool identifier or `amazon`.
	// This option is only available for VPC EIPs.
	PublicIpv4Pool pulumi.StringOutput `pulumi:"publicIpv4Pool"`
	// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Boolean if the EIP is in a VPC or not.
	// Defaults to `true` unless the region supports EC2-Classic.
	Vpc pulumi.BoolOutput `pulumi:"vpc"`
}

// NewEip registers a new resource with the given unique name, arguments, and options.
func NewEip(ctx *pulumi.Context,
	name string, args *EipArgs, opts ...pulumi.ResourceOption) (*Eip, error) {
	if args == nil {
		args = &EipArgs{}
	}

	var resource Eip
	err := ctx.RegisterResource("aws:ec2/eip:Eip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEip gets an existing Eip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipState, opts ...pulumi.ResourceOption) (*Eip, error) {
	var resource Eip
	err := ctx.ReadResource("aws:ec2/eip:Eip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Eip resources.
type eipState struct {
	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address *string `pulumi:"address"`
	// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
	AllocationId *string `pulumi:"allocationId"`
	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIp *string `pulumi:"associateWithPrivateIp"`
	// ID representing the association of the address with an instance in a VPC.
	AssociationId *string `pulumi:"associationId"`
	// Carrier IP address.
	CarrierIp *string `pulumi:"carrierIp"`
	// Customer owned IP.
	CustomerOwnedIp *string `pulumi:"customerOwnedIp"`
	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
	CustomerOwnedIpv4Pool *string `pulumi:"customerOwnedIpv4Pool"`
	// Indicates if this EIP is for use in VPC (`vpc`) or EC2-Classic (`standard`).
	Domain *string `pulumi:"domain"`
	// EC2 instance ID.
	Instance *string `pulumi:"instance"`
	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup *string `pulumi:"networkBorderGroup"`
	// Network interface ID to associate with.
	NetworkInterface *string `pulumi:"networkInterface"`
	// The Private DNS associated with the Elastic IP address (if in VPC).
	PrivateDns *string `pulumi:"privateDns"`
	// Contains the private IP address (if in VPC).
	PrivateIp *string `pulumi:"privateIp"`
	// Public DNS associated with the Elastic IP address.
	PublicDns *string `pulumi:"publicDns"`
	// Contains the public IP address.
	PublicIp *string `pulumi:"publicIp"`
	// EC2 IPv4 address pool identifier or `amazon`.
	// This option is only available for VPC EIPs.
	PublicIpv4Pool *string `pulumi:"publicIpv4Pool"`
	// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Boolean if the EIP is in a VPC or not.
	// Defaults to `true` unless the region supports EC2-Classic.
	Vpc *bool `pulumi:"vpc"`
}

type EipState struct {
	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address pulumi.StringPtrInput
	// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
	AllocationId pulumi.StringPtrInput
	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIp pulumi.StringPtrInput
	// ID representing the association of the address with an instance in a VPC.
	AssociationId pulumi.StringPtrInput
	// Carrier IP address.
	CarrierIp pulumi.StringPtrInput
	// Customer owned IP.
	CustomerOwnedIp pulumi.StringPtrInput
	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
	CustomerOwnedIpv4Pool pulumi.StringPtrInput
	// Indicates if this EIP is for use in VPC (`vpc`) or EC2-Classic (`standard`).
	Domain pulumi.StringPtrInput
	// EC2 instance ID.
	Instance pulumi.StringPtrInput
	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup pulumi.StringPtrInput
	// Network interface ID to associate with.
	NetworkInterface pulumi.StringPtrInput
	// The Private DNS associated with the Elastic IP address (if in VPC).
	PrivateDns pulumi.StringPtrInput
	// Contains the private IP address (if in VPC).
	PrivateIp pulumi.StringPtrInput
	// Public DNS associated with the Elastic IP address.
	PublicDns pulumi.StringPtrInput
	// Contains the public IP address.
	PublicIp pulumi.StringPtrInput
	// EC2 IPv4 address pool identifier or `amazon`.
	// This option is only available for VPC EIPs.
	PublicIpv4Pool pulumi.StringPtrInput
	// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Boolean if the EIP is in a VPC or not.
	// Defaults to `true` unless the region supports EC2-Classic.
	Vpc pulumi.BoolPtrInput
}

func (EipState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipState)(nil)).Elem()
}

type eipArgs struct {
	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address *string `pulumi:"address"`
	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIp *string `pulumi:"associateWithPrivateIp"`
	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
	CustomerOwnedIpv4Pool *string `pulumi:"customerOwnedIpv4Pool"`
	// EC2 instance ID.
	Instance *string `pulumi:"instance"`
	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup *string `pulumi:"networkBorderGroup"`
	// Network interface ID to associate with.
	NetworkInterface *string `pulumi:"networkInterface"`
	// EC2 IPv4 address pool identifier or `amazon`.
	// This option is only available for VPC EIPs.
	PublicIpv4Pool *string `pulumi:"publicIpv4Pool"`
	// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Boolean if the EIP is in a VPC or not.
	// Defaults to `true` unless the region supports EC2-Classic.
	Vpc *bool `pulumi:"vpc"`
}

// The set of arguments for constructing a Eip resource.
type EipArgs struct {
	// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
	Address pulumi.StringPtrInput
	// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	AssociateWithPrivateIp pulumi.StringPtrInput
	// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
	CustomerOwnedIpv4Pool pulumi.StringPtrInput
	// EC2 instance ID.
	Instance pulumi.StringPtrInput
	// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
	NetworkBorderGroup pulumi.StringPtrInput
	// Network interface ID to associate with.
	NetworkInterface pulumi.StringPtrInput
	// EC2 IPv4 address pool identifier or `amazon`.
	// This option is only available for VPC EIPs.
	PublicIpv4Pool pulumi.StringPtrInput
	// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Boolean if the EIP is in a VPC or not.
	// Defaults to `true` unless the region supports EC2-Classic.
	Vpc pulumi.BoolPtrInput
}

func (EipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipArgs)(nil)).Elem()
}

type EipInput interface {
	pulumi.Input

	ToEipOutput() EipOutput
	ToEipOutputWithContext(ctx context.Context) EipOutput
}

func (*Eip) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (i *Eip) ToEipOutput() EipOutput {
	return i.ToEipOutputWithContext(context.Background())
}

func (i *Eip) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipOutput)
}

// EipArrayInput is an input type that accepts EipArray and EipArrayOutput values.
// You can construct a concrete instance of `EipArrayInput` via:
//
//	EipArray{ EipArgs{...} }
type EipArrayInput interface {
	pulumi.Input

	ToEipArrayOutput() EipArrayOutput
	ToEipArrayOutputWithContext(context.Context) EipArrayOutput
}

type EipArray []EipInput

func (EipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (i EipArray) ToEipArrayOutput() EipArrayOutput {
	return i.ToEipArrayOutputWithContext(context.Background())
}

func (i EipArray) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipArrayOutput)
}

// EipMapInput is an input type that accepts EipMap and EipMapOutput values.
// You can construct a concrete instance of `EipMapInput` via:
//
//	EipMap{ "key": EipArgs{...} }
type EipMapInput interface {
	pulumi.Input

	ToEipMapOutput() EipMapOutput
	ToEipMapOutputWithContext(context.Context) EipMapOutput
}

type EipMap map[string]EipInput

func (EipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (i EipMap) ToEipMapOutput() EipMapOutput {
	return i.ToEipMapOutputWithContext(context.Background())
}

func (i EipMap) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipMapOutput)
}

type EipOutput struct{ *pulumi.OutputState }

func (EipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (o EipOutput) ToEipOutput() EipOutput {
	return o
}

func (o EipOutput) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return o
}

// IP address from an EC2 BYOIP pool. This option is only available for VPC EIPs.
func (o EipOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// ID that AWS assigns to represent the allocation of the Elastic IP address for use with instances in a VPC.
func (o EipOutput) AllocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.AllocationId }).(pulumi.StringOutput)
}

// User-specified primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
func (o EipOutput) AssociateWithPrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.AssociateWithPrivateIp }).(pulumi.StringPtrOutput)
}

// ID representing the association of the address with an instance in a VPC.
func (o EipOutput) AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.AssociationId }).(pulumi.StringOutput)
}

// Carrier IP address.
func (o EipOutput) CarrierIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.CarrierIp }).(pulumi.StringOutput)
}

// Customer owned IP.
func (o EipOutput) CustomerOwnedIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.CustomerOwnedIp }).(pulumi.StringOutput)
}

// ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing).
func (o EipOutput) CustomerOwnedIpv4Pool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringPtrOutput { return v.CustomerOwnedIpv4Pool }).(pulumi.StringPtrOutput)
}

// Indicates if this EIP is for use in VPC (`vpc`) or EC2-Classic (`standard`).
func (o EipOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// EC2 instance ID.
func (o EipOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// Location from which the IP address is advertised. Use this parameter to limit the address to this location.
func (o EipOutput) NetworkBorderGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.NetworkBorderGroup }).(pulumi.StringOutput)
}

// Network interface ID to associate with.
func (o EipOutput) NetworkInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.NetworkInterface }).(pulumi.StringOutput)
}

// The Private DNS associated with the Elastic IP address (if in VPC).
func (o EipOutput) PrivateDns() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PrivateDns }).(pulumi.StringOutput)
}

// Contains the private IP address (if in VPC).
func (o EipOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// Public DNS associated with the Elastic IP address.
func (o EipOutput) PublicDns() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PublicDns }).(pulumi.StringOutput)
}

// Contains the public IP address.
func (o EipOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

// EC2 IPv4 address pool identifier or `amazon`.
// This option is only available for VPC EIPs.
func (o EipOutput) PublicIpv4Pool() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.PublicIpv4Pool }).(pulumi.StringOutput)
}

// Map of tags to assign to the resource. Tags can only be applied to EIPs in a VPC. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EipOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o EipOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Boolean if the EIP is in a VPC or not.
// Defaults to `true` unless the region supports EC2-Classic.
func (o EipOutput) Vpc() pulumi.BoolOutput {
	return o.ApplyT(func(v *Eip) pulumi.BoolOutput { return v.Vpc }).(pulumi.BoolOutput)
}

type EipArrayOutput struct{ *pulumi.OutputState }

func (EipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (o EipArrayOutput) ToEipArrayOutput() EipArrayOutput {
	return o
}

func (o EipArrayOutput) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return o
}

func (o EipArrayOutput) Index(i pulumi.IntInput) EipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].([]*Eip)[vs[1].(int)]
	}).(EipOutput)
}

type EipMapOutput struct{ *pulumi.OutputState }

func (EipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (o EipMapOutput) ToEipMapOutput() EipMapOutput {
	return o
}

func (o EipMapOutput) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return o
}

func (o EipMapOutput) MapIndex(k pulumi.StringInput) EipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].(map[string]*Eip)[vs[1].(string)]
	}).(EipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipInput)(nil)).Elem(), &Eip{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipArrayInput)(nil)).Elem(), EipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipMapInput)(nil)).Elem(), EipMap{})
	pulumi.RegisterOutputType(EipOutput{})
	pulumi.RegisterOutputType(EipArrayOutput{})
	pulumi.RegisterOutputType(EipMapOutput{})
}