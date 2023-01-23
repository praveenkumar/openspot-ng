// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource attaches a security group to an Elastic Network Interface (ENI).
// It can be used to attach a security group to any existing ENI, be it a
// secondary ENI or one attached as the primary interface on an instance.
//
// > **NOTE on instances, interfaces, and security groups:** This provider currently
// provides the capability to assign security groups via the [`ec2.Instance`][1]
// and the [`ec2.NetworkInterface`][2] resources. Using this resource in
// conjunction with security groups provided in-line in those resources will cause
// conflicts, and will lead to spurious diffs and undefined behavior - please use
// one or the other.
//
// ## Example Usage
//
// The following provides a very basic example of setting up an instance (provided
// by `instance`) in the default security group, creating a security group
// (provided by `sg`) and then attaching the security group to the instance's
// primary network interface via the `ec2.NetworkInterfaceSecurityGroupAttachment` resource,
// named `sgAttachment`:
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
//			ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
//				MostRecent: pulumi.BoolRef(true),
//				Filters: []ec2.GetAmiFilter{
//					{
//						Name: "name",
//						Values: []string{
//							"amzn-ami-hvm-*",
//						},
//					},
//				},
//				Owners: []string{
//					"amazon",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			instance, err := ec2.NewInstance(ctx, "instance", &ec2.InstanceArgs{
//				InstanceType: pulumi.String("t2.micro"),
//				Ami:          *pulumi.String(ami.Id),
//				Tags: pulumi.StringMap{
//					"type": pulumi.String("test-instance"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
//				Tags: pulumi.StringMap{
//					"type": pulumi.String("test-security-group"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewNetworkInterfaceSecurityGroupAttachment(ctx, "sgAttachment", &ec2.NetworkInterfaceSecurityGroupAttachmentArgs{
//				SecurityGroupId:    sg.ID(),
//				NetworkInterfaceId: instance.PrimaryNetworkInterfaceId,
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
// In this example, `instance` is provided by the `ec2.Instance` data source,
// fetching an external instance, possibly not managed by this provider.
// `sgAttachment` then attaches to the output instance's `networkInterfaceId`:
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
//			instance, err := ec2.LookupInstance(ctx, &ec2.LookupInstanceArgs{
//				InstanceId: pulumi.StringRef("i-1234567890abcdef0"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			sg, err := ec2.NewSecurityGroup(ctx, "sg", &ec2.SecurityGroupArgs{
//				Tags: pulumi.StringMap{
//					"type": pulumi.String("test-security-group"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewNetworkInterfaceSecurityGroupAttachment(ctx, "sgAttachment", &ec2.NetworkInterfaceSecurityGroupAttachmentArgs{
//				SecurityGroupId:    sg.ID(),
//				NetworkInterfaceId: *pulumi.String(instance.NetworkInterfaceId),
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
// Network Interface Security Group attachments can be imported using the associated network interface ID and security group ID, separated by an underscore (`_`). For example
//
// ```sh
//
//	$ pulumi import aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment sg_attachment eni-1234567890abcdef0_sg-1234567890abcdef0
//
// ```
type NetworkInterfaceSecurityGroupAttachment struct {
	pulumi.CustomResourceState

	// The ID of the network interface to attach to.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The ID of the security group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
}

// NewNetworkInterfaceSecurityGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceSecurityGroupAttachment(ctx *pulumi.Context,
	name string, args *NetworkInterfaceSecurityGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	var resource NetworkInterfaceSecurityGroupAttachment
	err := ctx.RegisterResource("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceSecurityGroupAttachment gets an existing NetworkInterfaceSecurityGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceSecurityGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceSecurityGroupAttachmentState, opts ...pulumi.ResourceOption) (*NetworkInterfaceSecurityGroupAttachment, error) {
	var resource NetworkInterfaceSecurityGroupAttachment
	err := ctx.ReadResource("aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAttachment resources.
type networkInterfaceSecurityGroupAttachmentState struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The ID of the security group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

type NetworkInterfaceSecurityGroupAttachmentState struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId pulumi.StringPtrInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringPtrInput
}

func (NetworkInterfaceSecurityGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAttachmentState)(nil)).Elem()
}

type networkInterfaceSecurityGroupAttachmentArgs struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the security group.
	SecurityGroupId string `pulumi:"securityGroupId"`
}

// The set of arguments for constructing a NetworkInterfaceSecurityGroupAttachment resource.
type NetworkInterfaceSecurityGroupAttachmentArgs struct {
	// The ID of the network interface to attach to.
	NetworkInterfaceId pulumi.StringInput
	// The ID of the security group.
	SecurityGroupId pulumi.StringInput
}

func (NetworkInterfaceSecurityGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceSecurityGroupAttachmentArgs)(nil)).Elem()
}

type NetworkInterfaceSecurityGroupAttachmentInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAttachmentOutput() NetworkInterfaceSecurityGroupAttachmentOutput
	ToNetworkInterfaceSecurityGroupAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentOutput
}

func (*NetworkInterfaceSecurityGroupAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceSecurityGroupAttachment)(nil)).Elem()
}

func (i *NetworkInterfaceSecurityGroupAttachment) ToNetworkInterfaceSecurityGroupAttachmentOutput() NetworkInterfaceSecurityGroupAttachmentOutput {
	return i.ToNetworkInterfaceSecurityGroupAttachmentOutputWithContext(context.Background())
}

func (i *NetworkInterfaceSecurityGroupAttachment) ToNetworkInterfaceSecurityGroupAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAttachmentOutput)
}

// NetworkInterfaceSecurityGroupAttachmentArrayInput is an input type that accepts NetworkInterfaceSecurityGroupAttachmentArray and NetworkInterfaceSecurityGroupAttachmentArrayOutput values.
// You can construct a concrete instance of `NetworkInterfaceSecurityGroupAttachmentArrayInput` via:
//
//	NetworkInterfaceSecurityGroupAttachmentArray{ NetworkInterfaceSecurityGroupAttachmentArgs{...} }
type NetworkInterfaceSecurityGroupAttachmentArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAttachmentArrayOutput() NetworkInterfaceSecurityGroupAttachmentArrayOutput
	ToNetworkInterfaceSecurityGroupAttachmentArrayOutputWithContext(context.Context) NetworkInterfaceSecurityGroupAttachmentArrayOutput
}

type NetworkInterfaceSecurityGroupAttachmentArray []NetworkInterfaceSecurityGroupAttachmentInput

func (NetworkInterfaceSecurityGroupAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterfaceSecurityGroupAttachment)(nil)).Elem()
}

func (i NetworkInterfaceSecurityGroupAttachmentArray) ToNetworkInterfaceSecurityGroupAttachmentArrayOutput() NetworkInterfaceSecurityGroupAttachmentArrayOutput {
	return i.ToNetworkInterfaceSecurityGroupAttachmentArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceSecurityGroupAttachmentArray) ToNetworkInterfaceSecurityGroupAttachmentArrayOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAttachmentArrayOutput)
}

// NetworkInterfaceSecurityGroupAttachmentMapInput is an input type that accepts NetworkInterfaceSecurityGroupAttachmentMap and NetworkInterfaceSecurityGroupAttachmentMapOutput values.
// You can construct a concrete instance of `NetworkInterfaceSecurityGroupAttachmentMapInput` via:
//
//	NetworkInterfaceSecurityGroupAttachmentMap{ "key": NetworkInterfaceSecurityGroupAttachmentArgs{...} }
type NetworkInterfaceSecurityGroupAttachmentMapInput interface {
	pulumi.Input

	ToNetworkInterfaceSecurityGroupAttachmentMapOutput() NetworkInterfaceSecurityGroupAttachmentMapOutput
	ToNetworkInterfaceSecurityGroupAttachmentMapOutputWithContext(context.Context) NetworkInterfaceSecurityGroupAttachmentMapOutput
}

type NetworkInterfaceSecurityGroupAttachmentMap map[string]NetworkInterfaceSecurityGroupAttachmentInput

func (NetworkInterfaceSecurityGroupAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterfaceSecurityGroupAttachment)(nil)).Elem()
}

func (i NetworkInterfaceSecurityGroupAttachmentMap) ToNetworkInterfaceSecurityGroupAttachmentMapOutput() NetworkInterfaceSecurityGroupAttachmentMapOutput {
	return i.ToNetworkInterfaceSecurityGroupAttachmentMapOutputWithContext(context.Background())
}

func (i NetworkInterfaceSecurityGroupAttachmentMap) ToNetworkInterfaceSecurityGroupAttachmentMapOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceSecurityGroupAttachmentMapOutput)
}

type NetworkInterfaceSecurityGroupAttachmentOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkInterfaceSecurityGroupAttachment)(nil)).Elem()
}

func (o NetworkInterfaceSecurityGroupAttachmentOutput) ToNetworkInterfaceSecurityGroupAttachmentOutput() NetworkInterfaceSecurityGroupAttachmentOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAttachmentOutput) ToNetworkInterfaceSecurityGroupAttachmentOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentOutput {
	return o
}

// The ID of the network interface to attach to.
func (o NetworkInterfaceSecurityGroupAttachmentOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterfaceSecurityGroupAttachment) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// The ID of the security group.
func (o NetworkInterfaceSecurityGroupAttachmentOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkInterfaceSecurityGroupAttachment) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

type NetworkInterfaceSecurityGroupAttachmentArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkInterfaceSecurityGroupAttachment)(nil)).Elem()
}

func (o NetworkInterfaceSecurityGroupAttachmentArrayOutput) ToNetworkInterfaceSecurityGroupAttachmentArrayOutput() NetworkInterfaceSecurityGroupAttachmentArrayOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAttachmentArrayOutput) ToNetworkInterfaceSecurityGroupAttachmentArrayOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentArrayOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAttachmentArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceSecurityGroupAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkInterfaceSecurityGroupAttachment {
		return vs[0].([]*NetworkInterfaceSecurityGroupAttachment)[vs[1].(int)]
	}).(NetworkInterfaceSecurityGroupAttachmentOutput)
}

type NetworkInterfaceSecurityGroupAttachmentMapOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceSecurityGroupAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkInterfaceSecurityGroupAttachment)(nil)).Elem()
}

func (o NetworkInterfaceSecurityGroupAttachmentMapOutput) ToNetworkInterfaceSecurityGroupAttachmentMapOutput() NetworkInterfaceSecurityGroupAttachmentMapOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAttachmentMapOutput) ToNetworkInterfaceSecurityGroupAttachmentMapOutputWithContext(ctx context.Context) NetworkInterfaceSecurityGroupAttachmentMapOutput {
	return o
}

func (o NetworkInterfaceSecurityGroupAttachmentMapOutput) MapIndex(k pulumi.StringInput) NetworkInterfaceSecurityGroupAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkInterfaceSecurityGroupAttachment {
		return vs[0].(map[string]*NetworkInterfaceSecurityGroupAttachment)[vs[1].(string)]
	}).(NetworkInterfaceSecurityGroupAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAttachmentInput)(nil)).Elem(), &NetworkInterfaceSecurityGroupAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAttachmentArrayInput)(nil)).Elem(), NetworkInterfaceSecurityGroupAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceSecurityGroupAttachmentMapInput)(nil)).Elem(), NetworkInterfaceSecurityGroupAttachmentMap{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAttachmentOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAttachmentArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceSecurityGroupAttachmentMapOutput{})
}
