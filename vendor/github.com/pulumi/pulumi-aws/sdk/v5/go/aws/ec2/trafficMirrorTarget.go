// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Traffic mirror target.\
// Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring
//
// ## Example Usage
//
// # To create a basic traffic mirror session
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
//			_, err := ec2.NewTrafficMirrorTarget(ctx, "nlb", &ec2.TrafficMirrorTargetArgs{
//				Description:            pulumi.String("NLB target"),
//				NetworkLoadBalancerArn: pulumi.Any(aws_lb.Lb.Arn),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewTrafficMirrorTarget(ctx, "eni", &ec2.TrafficMirrorTargetArgs{
//				Description:        pulumi.String("ENI target"),
//				NetworkInterfaceId: pulumi.Any(aws_instance.Test.Primary_network_interface_id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewTrafficMirrorTarget(ctx, "gwlb", &ec2.TrafficMirrorTargetArgs{
//				Description:                   pulumi.String("GWLB target"),
//				GatewayLoadBalancerEndpointId: pulumi.Any(aws_vpc_endpoint.Example.Id),
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
// Traffic mirror targets can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/trafficMirrorTarget:TrafficMirrorTarget target tmt-0c13a005422b86606
//
// ```
type TrafficMirrorTarget struct {
	pulumi.CustomResourceState

	// The ARN of the traffic mirror target.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the traffic mirror session.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
	GatewayLoadBalancerEndpointId pulumi.StringPtrOutput `pulumi:"gatewayLoadBalancerEndpointId"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId pulumi.StringPtrOutput `pulumi:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn pulumi.StringPtrOutput `pulumi:"networkLoadBalancerArn"`
	// The ID of the AWS account that owns the traffic mirror target.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTrafficMirrorTarget registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorTarget(ctx *pulumi.Context,
	name string, args *TrafficMirrorTargetArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorTarget, error) {
	if args == nil {
		args = &TrafficMirrorTargetArgs{}
	}

	var resource TrafficMirrorTarget
	err := ctx.RegisterResource("aws:ec2/trafficMirrorTarget:TrafficMirrorTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorTarget gets an existing TrafficMirrorTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorTargetState, opts ...pulumi.ResourceOption) (*TrafficMirrorTarget, error) {
	var resource TrafficMirrorTarget
	err := ctx.ReadResource("aws:ec2/trafficMirrorTarget:TrafficMirrorTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorTarget resources.
type trafficMirrorTargetState struct {
	// The ARN of the traffic mirror target.
	Arn *string `pulumi:"arn"`
	// A description of the traffic mirror session.
	Description *string `pulumi:"description"`
	// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
	GatewayLoadBalancerEndpointId *string `pulumi:"gatewayLoadBalancerEndpointId"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn *string `pulumi:"networkLoadBalancerArn"`
	// The ID of the AWS account that owns the traffic mirror target.
	OwnerId *string `pulumi:"ownerId"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TrafficMirrorTargetState struct {
	// The ARN of the traffic mirror target.
	Arn pulumi.StringPtrInput
	// A description of the traffic mirror session.
	Description pulumi.StringPtrInput
	// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
	GatewayLoadBalancerEndpointId pulumi.StringPtrInput
	// The network interface ID that is associated with the target.
	NetworkInterfaceId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn pulumi.StringPtrInput
	// The ID of the AWS account that owns the traffic mirror target.
	OwnerId pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (TrafficMirrorTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorTargetState)(nil)).Elem()
}

type trafficMirrorTargetArgs struct {
	// A description of the traffic mirror session.
	Description *string `pulumi:"description"`
	// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
	GatewayLoadBalancerEndpointId *string `pulumi:"gatewayLoadBalancerEndpointId"`
	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn *string `pulumi:"networkLoadBalancerArn"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

// The set of arguments for constructing a TrafficMirrorTarget resource.
type TrafficMirrorTargetArgs struct {
	// A description of the traffic mirror session.
	Description pulumi.StringPtrInput
	// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
	GatewayLoadBalancerEndpointId pulumi.StringPtrInput
	// The network interface ID that is associated with the target.
	NetworkInterfaceId pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
	NetworkLoadBalancerArn pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (TrafficMirrorTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorTargetArgs)(nil)).Elem()
}

type TrafficMirrorTargetInput interface {
	pulumi.Input

	ToTrafficMirrorTargetOutput() TrafficMirrorTargetOutput
	ToTrafficMirrorTargetOutputWithContext(ctx context.Context) TrafficMirrorTargetOutput
}

func (*TrafficMirrorTarget) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorTarget)(nil)).Elem()
}

func (i *TrafficMirrorTarget) ToTrafficMirrorTargetOutput() TrafficMirrorTargetOutput {
	return i.ToTrafficMirrorTargetOutputWithContext(context.Background())
}

func (i *TrafficMirrorTarget) ToTrafficMirrorTargetOutputWithContext(ctx context.Context) TrafficMirrorTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorTargetOutput)
}

// TrafficMirrorTargetArrayInput is an input type that accepts TrafficMirrorTargetArray and TrafficMirrorTargetArrayOutput values.
// You can construct a concrete instance of `TrafficMirrorTargetArrayInput` via:
//
//	TrafficMirrorTargetArray{ TrafficMirrorTargetArgs{...} }
type TrafficMirrorTargetArrayInput interface {
	pulumi.Input

	ToTrafficMirrorTargetArrayOutput() TrafficMirrorTargetArrayOutput
	ToTrafficMirrorTargetArrayOutputWithContext(context.Context) TrafficMirrorTargetArrayOutput
}

type TrafficMirrorTargetArray []TrafficMirrorTargetInput

func (TrafficMirrorTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorTarget)(nil)).Elem()
}

func (i TrafficMirrorTargetArray) ToTrafficMirrorTargetArrayOutput() TrafficMirrorTargetArrayOutput {
	return i.ToTrafficMirrorTargetArrayOutputWithContext(context.Background())
}

func (i TrafficMirrorTargetArray) ToTrafficMirrorTargetArrayOutputWithContext(ctx context.Context) TrafficMirrorTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorTargetArrayOutput)
}

// TrafficMirrorTargetMapInput is an input type that accepts TrafficMirrorTargetMap and TrafficMirrorTargetMapOutput values.
// You can construct a concrete instance of `TrafficMirrorTargetMapInput` via:
//
//	TrafficMirrorTargetMap{ "key": TrafficMirrorTargetArgs{...} }
type TrafficMirrorTargetMapInput interface {
	pulumi.Input

	ToTrafficMirrorTargetMapOutput() TrafficMirrorTargetMapOutput
	ToTrafficMirrorTargetMapOutputWithContext(context.Context) TrafficMirrorTargetMapOutput
}

type TrafficMirrorTargetMap map[string]TrafficMirrorTargetInput

func (TrafficMirrorTargetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorTarget)(nil)).Elem()
}

func (i TrafficMirrorTargetMap) ToTrafficMirrorTargetMapOutput() TrafficMirrorTargetMapOutput {
	return i.ToTrafficMirrorTargetMapOutputWithContext(context.Background())
}

func (i TrafficMirrorTargetMap) ToTrafficMirrorTargetMapOutputWithContext(ctx context.Context) TrafficMirrorTargetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorTargetMapOutput)
}

type TrafficMirrorTargetOutput struct{ *pulumi.OutputState }

func (TrafficMirrorTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorTarget)(nil)).Elem()
}

func (o TrafficMirrorTargetOutput) ToTrafficMirrorTargetOutput() TrafficMirrorTargetOutput {
	return o
}

func (o TrafficMirrorTargetOutput) ToTrafficMirrorTargetOutputWithContext(ctx context.Context) TrafficMirrorTargetOutput {
	return o
}

// The ARN of the traffic mirror target.
func (o TrafficMirrorTargetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description of the traffic mirror session.
func (o TrafficMirrorTargetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The VPC Endpoint Id of the Gateway Load Balancer that is associated with the target.
func (o TrafficMirrorTargetOutput) GatewayLoadBalancerEndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringPtrOutput { return v.GatewayLoadBalancerEndpointId }).(pulumi.StringPtrOutput)
}

// The network interface ID that is associated with the target.
func (o TrafficMirrorTargetOutput) NetworkInterfaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringPtrOutput { return v.NetworkInterfaceId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
func (o TrafficMirrorTargetOutput) NetworkLoadBalancerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringPtrOutput { return v.NetworkLoadBalancerArn }).(pulumi.StringPtrOutput)
}

// The ID of the AWS account that owns the traffic mirror target.
func (o TrafficMirrorTargetOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TrafficMirrorTargetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o TrafficMirrorTargetOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TrafficMirrorTarget) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type TrafficMirrorTargetArrayOutput struct{ *pulumi.OutputState }

func (TrafficMirrorTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorTarget)(nil)).Elem()
}

func (o TrafficMirrorTargetArrayOutput) ToTrafficMirrorTargetArrayOutput() TrafficMirrorTargetArrayOutput {
	return o
}

func (o TrafficMirrorTargetArrayOutput) ToTrafficMirrorTargetArrayOutputWithContext(ctx context.Context) TrafficMirrorTargetArrayOutput {
	return o
}

func (o TrafficMirrorTargetArrayOutput) Index(i pulumi.IntInput) TrafficMirrorTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMirrorTarget {
		return vs[0].([]*TrafficMirrorTarget)[vs[1].(int)]
	}).(TrafficMirrorTargetOutput)
}

type TrafficMirrorTargetMapOutput struct{ *pulumi.OutputState }

func (TrafficMirrorTargetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorTarget)(nil)).Elem()
}

func (o TrafficMirrorTargetMapOutput) ToTrafficMirrorTargetMapOutput() TrafficMirrorTargetMapOutput {
	return o
}

func (o TrafficMirrorTargetMapOutput) ToTrafficMirrorTargetMapOutputWithContext(ctx context.Context) TrafficMirrorTargetMapOutput {
	return o
}

func (o TrafficMirrorTargetMapOutput) MapIndex(k pulumi.StringInput) TrafficMirrorTargetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMirrorTarget {
		return vs[0].(map[string]*TrafficMirrorTarget)[vs[1].(string)]
	}).(TrafficMirrorTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorTargetInput)(nil)).Elem(), &TrafficMirrorTarget{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorTargetArrayInput)(nil)).Elem(), TrafficMirrorTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorTargetMapInput)(nil)).Elem(), TrafficMirrorTargetMap{})
	pulumi.RegisterOutputType(TrafficMirrorTargetOutput{})
	pulumi.RegisterOutputType(TrafficMirrorTargetArrayOutput{})
	pulumi.RegisterOutputType(TrafficMirrorTargetMapOutput{})
}
