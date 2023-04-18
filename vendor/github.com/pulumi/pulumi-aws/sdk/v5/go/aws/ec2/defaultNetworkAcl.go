// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a VPC's default network ACL. This resource can manage the default network ACL of the default or a non-default VPC.
//
// > **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `ec2.DefaultNetworkAcl` behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management.
//
// Every VPC has a default network ACL that can be managed but not destroyed. When the provider first adopts the Default Network ACL, it **immediately removes all rules in the ACL**. It then proceeds to create any rules specified in the configuration. This step is required so that only the rules specified in the configuration are created.
//
// This resource treats its inline rules as absolute; only the rules defined inline are created, and any additions/removals external to this resource will result in diffs being shown. For these reasons, this resource is incompatible with the `ec2.NetworkAclRule` resource.
//
// For more information about Network ACLs, see the AWS Documentation on [Network ACLs][aws-network-acls].
//
// ## Example Usage
// ### Basic Example
//
// The following config gives the Default Network ACL the same rules that AWS includes but pulls the resource under management by this provider. This means that any ACL rules added or changed will be detected as drift.
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
//			mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewDefaultNetworkAcl(ctx, "default", &ec2.DefaultNetworkAclArgs{
//				DefaultNetworkAclId: mainvpc.DefaultNetworkAclId,
//				Ingress: ec2.DefaultNetworkAclIngressArray{
//					&ec2.DefaultNetworkAclIngressArgs{
//						Protocol:  pulumi.String("-1"),
//						RuleNo:    pulumi.Int(100),
//						Action:    pulumi.String("allow"),
//						CidrBlock: pulumi.String("0.0.0.0/0"),
//						FromPort:  pulumi.Int(0),
//						ToPort:    pulumi.Int(0),
//					},
//				},
//				Egress: ec2.DefaultNetworkAclEgressArray{
//					&ec2.DefaultNetworkAclEgressArgs{
//						Protocol:  pulumi.String("-1"),
//						RuleNo:    pulumi.Int(100),
//						Action:    pulumi.String("allow"),
//						CidrBlock: pulumi.String("0.0.0.0/0"),
//						FromPort:  pulumi.Int(0),
//						ToPort:    pulumi.Int(0),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example: Deny All Egress Traffic, Allow Ingress
//
// The following denies all Egress traffic by omitting any `egress` rules, while including the default `ingress` rule to allow all traffic.
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
//			mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewDefaultNetworkAcl(ctx, "default", &ec2.DefaultNetworkAclArgs{
//				DefaultNetworkAclId: mainvpc.DefaultNetworkAclId,
//				Ingress: ec2.DefaultNetworkAclIngressArray{
//					&ec2.DefaultNetworkAclIngressArgs{
//						Protocol:  pulumi.String("-1"),
//						RuleNo:    pulumi.Int(100),
//						Action:    pulumi.String("allow"),
//						CidrBlock: pulumi.Any(aws_default_vpc.Mainvpc.Cidr_block),
//						FromPort:  pulumi.Int(0),
//						ToPort:    pulumi.Int(0),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example: Deny All Traffic To Any Subnet In The Default Network ACL
//
// This config denies all traffic in the Default ACL. This can be useful if you want to lock down the VPC to force all resources to assign a non-default ACL.
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
//			mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewDefaultNetworkAcl(ctx, "default", &ec2.DefaultNetworkAclArgs{
//				DefaultNetworkAclId: mainvpc.DefaultNetworkAclId,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Managing Subnets In A Default Network ACL
//
// Within a VPC, all Subnets must be associated with a Network ACL. In order to "delete" the association between a Subnet and a non-default Network ACL, the association is destroyed by replacing it with an association between the Subnet and the Default ACL instead.
//
// When managing the Default Network ACL, you cannot "remove" Subnets. Instead, they must be reassigned to another Network ACL, or the Subnet itself must be destroyed. Because of these requirements, removing the `subnetIds` attribute from the configuration of a `ec2.DefaultNetworkAcl` resource may result in a reoccurring plan, until the Subnets are reassigned to another Network ACL or are destroyed.
//
// Because Subnets are by default associated with the Default Network ACL, any non-explicit association will show up as a plan to remove the Subnet. For example: if you have a custom `ec2.NetworkAcl` with two subnets attached, and you remove the `ec2.NetworkAcl` resource, after successfully destroying this resource future plans will show a diff on the managed `ec2.DefaultNetworkAcl`, as those two Subnets have been orphaned by the now destroyed network acl and thus adopted by the Default Network ACL. In order to avoid a reoccurring plan, they will need to be reassigned, destroyed, or added to the `subnetIds` attribute of the `ec2.DefaultNetworkAcl` entry.
//
// As an alternative to the above, you can also specify the following lifecycle configuration in your `ec2.DefaultNetworkAcl` resource:
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
//			_, err := ec2.NewDefaultNetworkAcl(ctx, "default", nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Removing `ec2.DefaultNetworkAcl` From Your Configuration
//
// Each AWS VPC comes with a Default Network ACL that cannot be deleted. The `ec2.DefaultNetworkAcl` allows you to manage this Network ACL, but the provider cannot destroy it. Removing this resource from your configuration will remove it from your statefile and management, **but will not destroy the Network ACL.** All Subnets associations and ingress or egress rules will be left as they are at the time of removal. You can resume managing them via the AWS Console.
//
// ## Import
//
// Default Network ACLs can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/defaultNetworkAcl:DefaultNetworkAcl sample acl-7aaabd18
//
// ```
type DefaultNetworkAcl struct {
	pulumi.CustomResourceState

	// ARN of the Default Network ACL
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Network ACL ID to manage. This attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`
	// Configuration block for an egress rule. Detailed below.
	Egress DefaultNetworkAclEgressArrayOutput `pulumi:"egress"`
	// Configuration block for an ingress rule. Detailed below.
	Ingress DefaultNetworkAclIngressArrayOutput `pulumi:"ingress"`
	// ID of the AWS account that owns the Default Network ACL
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// ID of the associated VPC
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewDefaultNetworkAcl(ctx *pulumi.Context,
	name string, args *DefaultNetworkAclArgs, opts ...pulumi.ResourceOption) (*DefaultNetworkAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultNetworkAclId == nil {
		return nil, errors.New("invalid value for required argument 'DefaultNetworkAclId'")
	}
	var resource DefaultNetworkAcl
	err := ctx.RegisterResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultNetworkAcl gets an existing DefaultNetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultNetworkAclState, opts ...pulumi.ResourceOption) (*DefaultNetworkAcl, error) {
	var resource DefaultNetworkAcl
	err := ctx.ReadResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultNetworkAcl resources.
type defaultNetworkAclState struct {
	// ARN of the Default Network ACL
	Arn *string `pulumi:"arn"`
	// Network ACL ID to manage. This attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId *string `pulumi:"defaultNetworkAclId"`
	// Configuration block for an egress rule. Detailed below.
	Egress []DefaultNetworkAclEgress `pulumi:"egress"`
	// Configuration block for an ingress rule. Detailed below.
	Ingress []DefaultNetworkAclIngress `pulumi:"ingress"`
	// ID of the AWS account that owns the Default Network ACL
	OwnerId *string `pulumi:"ownerId"`
	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	SubnetIds []string `pulumi:"subnetIds"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// ID of the associated VPC
	VpcId *string `pulumi:"vpcId"`
}

type DefaultNetworkAclState struct {
	// ARN of the Default Network ACL
	Arn pulumi.StringPtrInput
	// Network ACL ID to manage. This attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringPtrInput
	// Configuration block for an egress rule. Detailed below.
	Egress DefaultNetworkAclEgressArrayInput
	// Configuration block for an ingress rule. Detailed below.
	Ingress DefaultNetworkAclIngressArrayInput
	// ID of the AWS account that owns the Default Network ACL
	OwnerId pulumi.StringPtrInput
	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// ID of the associated VPC
	VpcId pulumi.StringPtrInput
}

func (DefaultNetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNetworkAclState)(nil)).Elem()
}

type defaultNetworkAclArgs struct {
	// Network ACL ID to manage. This attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId string `pulumi:"defaultNetworkAclId"`
	// Configuration block for an egress rule. Detailed below.
	Egress []DefaultNetworkAclEgress `pulumi:"egress"`
	// Configuration block for an ingress rule. Detailed below.
	Ingress []DefaultNetworkAclIngress `pulumi:"ingress"`
	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	SubnetIds []string `pulumi:"subnetIds"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a DefaultNetworkAcl resource.
type DefaultNetworkAclArgs struct {
	// Network ACL ID to manage. This attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringInput
	// Configuration block for an egress rule. Detailed below.
	Egress DefaultNetworkAclEgressArrayInput
	// Configuration block for an ingress rule. Detailed below.
	Ingress DefaultNetworkAclIngressArrayInput
	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (DefaultNetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNetworkAclArgs)(nil)).Elem()
}

type DefaultNetworkAclInput interface {
	pulumi.Input

	ToDefaultNetworkAclOutput() DefaultNetworkAclOutput
	ToDefaultNetworkAclOutputWithContext(ctx context.Context) DefaultNetworkAclOutput
}

func (*DefaultNetworkAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNetworkAcl)(nil)).Elem()
}

func (i *DefaultNetworkAcl) ToDefaultNetworkAclOutput() DefaultNetworkAclOutput {
	return i.ToDefaultNetworkAclOutputWithContext(context.Background())
}

func (i *DefaultNetworkAcl) ToDefaultNetworkAclOutputWithContext(ctx context.Context) DefaultNetworkAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNetworkAclOutput)
}

// DefaultNetworkAclArrayInput is an input type that accepts DefaultNetworkAclArray and DefaultNetworkAclArrayOutput values.
// You can construct a concrete instance of `DefaultNetworkAclArrayInput` via:
//
//	DefaultNetworkAclArray{ DefaultNetworkAclArgs{...} }
type DefaultNetworkAclArrayInput interface {
	pulumi.Input

	ToDefaultNetworkAclArrayOutput() DefaultNetworkAclArrayOutput
	ToDefaultNetworkAclArrayOutputWithContext(context.Context) DefaultNetworkAclArrayOutput
}

type DefaultNetworkAclArray []DefaultNetworkAclInput

func (DefaultNetworkAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNetworkAcl)(nil)).Elem()
}

func (i DefaultNetworkAclArray) ToDefaultNetworkAclArrayOutput() DefaultNetworkAclArrayOutput {
	return i.ToDefaultNetworkAclArrayOutputWithContext(context.Background())
}

func (i DefaultNetworkAclArray) ToDefaultNetworkAclArrayOutputWithContext(ctx context.Context) DefaultNetworkAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNetworkAclArrayOutput)
}

// DefaultNetworkAclMapInput is an input type that accepts DefaultNetworkAclMap and DefaultNetworkAclMapOutput values.
// You can construct a concrete instance of `DefaultNetworkAclMapInput` via:
//
//	DefaultNetworkAclMap{ "key": DefaultNetworkAclArgs{...} }
type DefaultNetworkAclMapInput interface {
	pulumi.Input

	ToDefaultNetworkAclMapOutput() DefaultNetworkAclMapOutput
	ToDefaultNetworkAclMapOutputWithContext(context.Context) DefaultNetworkAclMapOutput
}

type DefaultNetworkAclMap map[string]DefaultNetworkAclInput

func (DefaultNetworkAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNetworkAcl)(nil)).Elem()
}

func (i DefaultNetworkAclMap) ToDefaultNetworkAclMapOutput() DefaultNetworkAclMapOutput {
	return i.ToDefaultNetworkAclMapOutputWithContext(context.Background())
}

func (i DefaultNetworkAclMap) ToDefaultNetworkAclMapOutputWithContext(ctx context.Context) DefaultNetworkAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNetworkAclMapOutput)
}

type DefaultNetworkAclOutput struct{ *pulumi.OutputState }

func (DefaultNetworkAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNetworkAcl)(nil)).Elem()
}

func (o DefaultNetworkAclOutput) ToDefaultNetworkAclOutput() DefaultNetworkAclOutput {
	return o
}

func (o DefaultNetworkAclOutput) ToDefaultNetworkAclOutputWithContext(ctx context.Context) DefaultNetworkAclOutput {
	return o
}

// ARN of the Default Network ACL
func (o DefaultNetworkAclOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Network ACL ID to manage. This attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
func (o DefaultNetworkAclOutput) DefaultNetworkAclId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringOutput { return v.DefaultNetworkAclId }).(pulumi.StringOutput)
}

// Configuration block for an egress rule. Detailed below.
func (o DefaultNetworkAclOutput) Egress() DefaultNetworkAclEgressArrayOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) DefaultNetworkAclEgressArrayOutput { return v.Egress }).(DefaultNetworkAclEgressArrayOutput)
}

// Configuration block for an ingress rule. Detailed below.
func (o DefaultNetworkAclOutput) Ingress() DefaultNetworkAclIngressArrayOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) DefaultNetworkAclIngressArrayOutput { return v.Ingress }).(DefaultNetworkAclIngressArrayOutput)
}

// ID of the AWS account that owns the Default Network ACL
func (o DefaultNetworkAclOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
func (o DefaultNetworkAclOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DefaultNetworkAclOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o DefaultNetworkAclOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// ID of the associated VPC
func (o DefaultNetworkAclOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNetworkAcl) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type DefaultNetworkAclArrayOutput struct{ *pulumi.OutputState }

func (DefaultNetworkAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNetworkAcl)(nil)).Elem()
}

func (o DefaultNetworkAclArrayOutput) ToDefaultNetworkAclArrayOutput() DefaultNetworkAclArrayOutput {
	return o
}

func (o DefaultNetworkAclArrayOutput) ToDefaultNetworkAclArrayOutputWithContext(ctx context.Context) DefaultNetworkAclArrayOutput {
	return o
}

func (o DefaultNetworkAclArrayOutput) Index(i pulumi.IntInput) DefaultNetworkAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultNetworkAcl {
		return vs[0].([]*DefaultNetworkAcl)[vs[1].(int)]
	}).(DefaultNetworkAclOutput)
}

type DefaultNetworkAclMapOutput struct{ *pulumi.OutputState }

func (DefaultNetworkAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNetworkAcl)(nil)).Elem()
}

func (o DefaultNetworkAclMapOutput) ToDefaultNetworkAclMapOutput() DefaultNetworkAclMapOutput {
	return o
}

func (o DefaultNetworkAclMapOutput) ToDefaultNetworkAclMapOutputWithContext(ctx context.Context) DefaultNetworkAclMapOutput {
	return o
}

func (o DefaultNetworkAclMapOutput) MapIndex(k pulumi.StringInput) DefaultNetworkAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultNetworkAcl {
		return vs[0].(map[string]*DefaultNetworkAcl)[vs[1].(string)]
	}).(DefaultNetworkAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNetworkAclInput)(nil)).Elem(), &DefaultNetworkAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNetworkAclArrayInput)(nil)).Elem(), DefaultNetworkAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNetworkAclMapInput)(nil)).Elem(), DefaultNetworkAclMap{})
	pulumi.RegisterOutputType(DefaultNetworkAclOutput{})
	pulumi.RegisterOutputType(DefaultNetworkAclArrayOutput{})
	pulumi.RegisterOutputType(DefaultNetworkAclMapOutput{})
}
