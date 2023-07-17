// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Traffic mirror filter rule.\
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
//			filter, err := ec2.NewTrafficMirrorFilter(ctx, "filter", &ec2.TrafficMirrorFilterArgs{
//				Description: pulumi.String("traffic mirror filter - example"),
//				NetworkServices: pulumi.StringArray{
//					pulumi.String("amazon-dns"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewTrafficMirrorFilterRule(ctx, "ruleout", &ec2.TrafficMirrorFilterRuleArgs{
//				Description:           pulumi.String("test rule"),
//				TrafficMirrorFilterId: filter.ID(),
//				DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
//				SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
//				RuleNumber:            pulumi.Int(1),
//				RuleAction:            pulumi.String("accept"),
//				TrafficDirection:      pulumi.String("egress"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewTrafficMirrorFilterRule(ctx, "rulein", &ec2.TrafficMirrorFilterRuleArgs{
//				Description:           pulumi.String("test rule"),
//				TrafficMirrorFilterId: filter.ID(),
//				DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
//				SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
//				RuleNumber:            pulumi.Int(1),
//				RuleAction:            pulumi.String("accept"),
//				TrafficDirection:      pulumi.String("ingress"),
//				Protocol:              pulumi.Int(6),
//				DestinationPortRange: &ec2.TrafficMirrorFilterRuleDestinationPortRangeArgs{
//					FromPort: pulumi.Int(22),
//					ToPort:   pulumi.Int(53),
//				},
//				SourcePortRange: &ec2.TrafficMirrorFilterRuleSourcePortRangeArgs{
//					FromPort: pulumi.Int(0),
//					ToPort:   pulumi.Int(10),
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
//
// ## Import
//
// Traffic mirror rules can be imported using the `traffic_mirror_filter_id` and `id` separated by `:` e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule rule tmf-0fbb93ddf38198f64:tmfr-05a458f06445d0aee
//
// ```
type TrafficMirrorFilterRule struct {
	pulumi.CustomResourceState

	// ARN of the traffic mirror filter rule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the traffic mirror filter rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange TrafficMirrorFilterRuleDestinationPortRangePtrOutput `pulumi:"destinationPortRange"`
	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol pulumi.IntPtrOutput `pulumi:"protocol"`
	// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber pulumi.IntOutput `pulumi:"ruleNumber"`
	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock pulumi.StringOutput `pulumi:"sourceCidrBlock"`
	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange TrafficMirrorFilterRuleSourcePortRangePtrOutput `pulumi:"sourcePortRange"`
	// Direction of traffic to be captured. Valid values are `ingress` and `egress`
	//
	// Traffic mirror port range support following attributes:
	TrafficDirection pulumi.StringOutput `pulumi:"trafficDirection"`
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId pulumi.StringOutput `pulumi:"trafficMirrorFilterId"`
}

// NewTrafficMirrorFilterRule registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorFilterRule(ctx *pulumi.Context,
	name string, args *TrafficMirrorFilterRuleArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.RuleAction == nil {
		return nil, errors.New("invalid value for required argument 'RuleAction'")
	}
	if args.RuleNumber == nil {
		return nil, errors.New("invalid value for required argument 'RuleNumber'")
	}
	if args.SourceCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'SourceCidrBlock'")
	}
	if args.TrafficDirection == nil {
		return nil, errors.New("invalid value for required argument 'TrafficDirection'")
	}
	if args.TrafficMirrorFilterId == nil {
		return nil, errors.New("invalid value for required argument 'TrafficMirrorFilterId'")
	}
	var resource TrafficMirrorFilterRule
	err := ctx.RegisterResource("aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorFilterRule gets an existing TrafficMirrorFilterRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorFilterRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorFilterRuleState, opts ...pulumi.ResourceOption) (*TrafficMirrorFilterRule, error) {
	var resource TrafficMirrorFilterRule
	err := ctx.ReadResource("aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorFilterRule resources.
type trafficMirrorFilterRuleState struct {
	// ARN of the traffic mirror filter rule.
	Arn *string `pulumi:"arn"`
	// Description of the traffic mirror filter rule.
	Description *string `pulumi:"description"`
	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange *TrafficMirrorFilterRuleDestinationPortRange `pulumi:"destinationPortRange"`
	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int `pulumi:"protocol"`
	// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction *string `pulumi:"ruleAction"`
	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber *int `pulumi:"ruleNumber"`
	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string `pulumi:"sourceCidrBlock"`
	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange *TrafficMirrorFilterRuleSourcePortRange `pulumi:"sourcePortRange"`
	// Direction of traffic to be captured. Valid values are `ingress` and `egress`
	//
	// Traffic mirror port range support following attributes:
	TrafficDirection *string `pulumi:"trafficDirection"`
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId *string `pulumi:"trafficMirrorFilterId"`
}

type TrafficMirrorFilterRuleState struct {
	// ARN of the traffic mirror filter rule.
	Arn pulumi.StringPtrInput
	// Description of the traffic mirror filter rule.
	Description pulumi.StringPtrInput
	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock pulumi.StringPtrInput
	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange TrafficMirrorFilterRuleDestinationPortRangePtrInput
	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol pulumi.IntPtrInput
	// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction pulumi.StringPtrInput
	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber pulumi.IntPtrInput
	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock pulumi.StringPtrInput
	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange TrafficMirrorFilterRuleSourcePortRangePtrInput
	// Direction of traffic to be captured. Valid values are `ingress` and `egress`
	//
	// Traffic mirror port range support following attributes:
	TrafficDirection pulumi.StringPtrInput
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId pulumi.StringPtrInput
}

func (TrafficMirrorFilterRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterRuleState)(nil)).Elem()
}

type trafficMirrorFilterRuleArgs struct {
	// Description of the traffic mirror filter rule.
	Description *string `pulumi:"description"`
	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange *TrafficMirrorFilterRuleDestinationPortRange `pulumi:"destinationPortRange"`
	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *int `pulumi:"protocol"`
	// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction string `pulumi:"ruleAction"`
	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber int `pulumi:"ruleNumber"`
	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock string `pulumi:"sourceCidrBlock"`
	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange *TrafficMirrorFilterRuleSourcePortRange `pulumi:"sourcePortRange"`
	// Direction of traffic to be captured. Valid values are `ingress` and `egress`
	//
	// Traffic mirror port range support following attributes:
	TrafficDirection string `pulumi:"trafficDirection"`
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId string `pulumi:"trafficMirrorFilterId"`
}

// The set of arguments for constructing a TrafficMirrorFilterRule resource.
type TrafficMirrorFilterRuleArgs struct {
	// Description of the traffic mirror filter rule.
	Description pulumi.StringPtrInput
	// Destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock pulumi.StringInput
	// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	DestinationPortRange TrafficMirrorFilterRuleDestinationPortRangePtrInput
	// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol pulumi.IntPtrInput
	// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
	RuleAction pulumi.StringInput
	// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber pulumi.IntInput
	// Source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock pulumi.StringInput
	// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
	SourcePortRange TrafficMirrorFilterRuleSourcePortRangePtrInput
	// Direction of traffic to be captured. Valid values are `ingress` and `egress`
	//
	// Traffic mirror port range support following attributes:
	TrafficDirection pulumi.StringInput
	// ID of the traffic mirror filter to which this rule should be added
	TrafficMirrorFilterId pulumi.StringInput
}

func (TrafficMirrorFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorFilterRuleArgs)(nil)).Elem()
}

type TrafficMirrorFilterRuleInput interface {
	pulumi.Input

	ToTrafficMirrorFilterRuleOutput() TrafficMirrorFilterRuleOutput
	ToTrafficMirrorFilterRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleOutput
}

func (*TrafficMirrorFilterRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilterRule)(nil)).Elem()
}

func (i *TrafficMirrorFilterRule) ToTrafficMirrorFilterRuleOutput() TrafficMirrorFilterRuleOutput {
	return i.ToTrafficMirrorFilterRuleOutputWithContext(context.Background())
}

func (i *TrafficMirrorFilterRule) ToTrafficMirrorFilterRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterRuleOutput)
}

// TrafficMirrorFilterRuleArrayInput is an input type that accepts TrafficMirrorFilterRuleArray and TrafficMirrorFilterRuleArrayOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterRuleArrayInput` via:
//
//	TrafficMirrorFilterRuleArray{ TrafficMirrorFilterRuleArgs{...} }
type TrafficMirrorFilterRuleArrayInput interface {
	pulumi.Input

	ToTrafficMirrorFilterRuleArrayOutput() TrafficMirrorFilterRuleArrayOutput
	ToTrafficMirrorFilterRuleArrayOutputWithContext(context.Context) TrafficMirrorFilterRuleArrayOutput
}

type TrafficMirrorFilterRuleArray []TrafficMirrorFilterRuleInput

func (TrafficMirrorFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilterRule)(nil)).Elem()
}

func (i TrafficMirrorFilterRuleArray) ToTrafficMirrorFilterRuleArrayOutput() TrafficMirrorFilterRuleArrayOutput {
	return i.ToTrafficMirrorFilterRuleArrayOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterRuleArray) ToTrafficMirrorFilterRuleArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterRuleArrayOutput)
}

// TrafficMirrorFilterRuleMapInput is an input type that accepts TrafficMirrorFilterRuleMap and TrafficMirrorFilterRuleMapOutput values.
// You can construct a concrete instance of `TrafficMirrorFilterRuleMapInput` via:
//
//	TrafficMirrorFilterRuleMap{ "key": TrafficMirrorFilterRuleArgs{...} }
type TrafficMirrorFilterRuleMapInput interface {
	pulumi.Input

	ToTrafficMirrorFilterRuleMapOutput() TrafficMirrorFilterRuleMapOutput
	ToTrafficMirrorFilterRuleMapOutputWithContext(context.Context) TrafficMirrorFilterRuleMapOutput
}

type TrafficMirrorFilterRuleMap map[string]TrafficMirrorFilterRuleInput

func (TrafficMirrorFilterRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilterRule)(nil)).Elem()
}

func (i TrafficMirrorFilterRuleMap) ToTrafficMirrorFilterRuleMapOutput() TrafficMirrorFilterRuleMapOutput {
	return i.ToTrafficMirrorFilterRuleMapOutputWithContext(context.Background())
}

func (i TrafficMirrorFilterRuleMap) ToTrafficMirrorFilterRuleMapOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorFilterRuleMapOutput)
}

type TrafficMirrorFilterRuleOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrafficMirrorFilterRule)(nil)).Elem()
}

func (o TrafficMirrorFilterRuleOutput) ToTrafficMirrorFilterRuleOutput() TrafficMirrorFilterRuleOutput {
	return o
}

func (o TrafficMirrorFilterRuleOutput) ToTrafficMirrorFilterRuleOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleOutput {
	return o
}

// ARN of the traffic mirror filter rule.
func (o TrafficMirrorFilterRuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the traffic mirror filter rule.
func (o TrafficMirrorFilterRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Destination CIDR block to assign to the Traffic Mirror rule.
func (o TrafficMirrorFilterRuleOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// Destination port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
func (o TrafficMirrorFilterRuleOutput) DestinationPortRange() TrafficMirrorFilterRuleDestinationPortRangePtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) TrafficMirrorFilterRuleDestinationPortRangePtrOutput {
		return v.DestinationPortRange
	}).(TrafficMirrorFilterRuleDestinationPortRangePtrOutput)
}

// Protocol number, for example 17 (UDP), to assign to the Traffic Mirror rule. For information about the protocol value, see [Protocol Numbers](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
func (o TrafficMirrorFilterRuleOutput) Protocol() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.IntPtrOutput { return v.Protocol }).(pulumi.IntPtrOutput)
}

// Action to take (accept | reject) on the filtered traffic. Valid values are `accept` and `reject`
func (o TrafficMirrorFilterRuleOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringOutput { return v.RuleAction }).(pulumi.StringOutput)
}

// Number of the Traffic Mirror rule. This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
func (o TrafficMirrorFilterRuleOutput) RuleNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.IntOutput { return v.RuleNumber }).(pulumi.IntOutput)
}

// Source CIDR block to assign to the Traffic Mirror rule.
func (o TrafficMirrorFilterRuleOutput) SourceCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringOutput { return v.SourceCidrBlock }).(pulumi.StringOutput)
}

// Source port range. Supported only when the protocol is set to TCP(6) or UDP(17). See Traffic mirror port range documented below
func (o TrafficMirrorFilterRuleOutput) SourcePortRange() TrafficMirrorFilterRuleSourcePortRangePtrOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) TrafficMirrorFilterRuleSourcePortRangePtrOutput {
		return v.SourcePortRange
	}).(TrafficMirrorFilterRuleSourcePortRangePtrOutput)
}

// Direction of traffic to be captured. Valid values are `ingress` and `egress`
//
// Traffic mirror port range support following attributes:
func (o TrafficMirrorFilterRuleOutput) TrafficDirection() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringOutput { return v.TrafficDirection }).(pulumi.StringOutput)
}

// ID of the traffic mirror filter to which this rule should be added
func (o TrafficMirrorFilterRuleOutput) TrafficMirrorFilterId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrafficMirrorFilterRule) pulumi.StringOutput { return v.TrafficMirrorFilterId }).(pulumi.StringOutput)
}

type TrafficMirrorFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TrafficMirrorFilterRule)(nil)).Elem()
}

func (o TrafficMirrorFilterRuleArrayOutput) ToTrafficMirrorFilterRuleArrayOutput() TrafficMirrorFilterRuleArrayOutput {
	return o
}

func (o TrafficMirrorFilterRuleArrayOutput) ToTrafficMirrorFilterRuleArrayOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleArrayOutput {
	return o
}

func (o TrafficMirrorFilterRuleArrayOutput) Index(i pulumi.IntInput) TrafficMirrorFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TrafficMirrorFilterRule {
		return vs[0].([]*TrafficMirrorFilterRule)[vs[1].(int)]
	}).(TrafficMirrorFilterRuleOutput)
}

type TrafficMirrorFilterRuleMapOutput struct{ *pulumi.OutputState }

func (TrafficMirrorFilterRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TrafficMirrorFilterRule)(nil)).Elem()
}

func (o TrafficMirrorFilterRuleMapOutput) ToTrafficMirrorFilterRuleMapOutput() TrafficMirrorFilterRuleMapOutput {
	return o
}

func (o TrafficMirrorFilterRuleMapOutput) ToTrafficMirrorFilterRuleMapOutputWithContext(ctx context.Context) TrafficMirrorFilterRuleMapOutput {
	return o
}

func (o TrafficMirrorFilterRuleMapOutput) MapIndex(k pulumi.StringInput) TrafficMirrorFilterRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TrafficMirrorFilterRule {
		return vs[0].(map[string]*TrafficMirrorFilterRule)[vs[1].(string)]
	}).(TrafficMirrorFilterRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterRuleInput)(nil)).Elem(), &TrafficMirrorFilterRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterRuleArrayInput)(nil)).Elem(), TrafficMirrorFilterRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrafficMirrorFilterRuleMapInput)(nil)).Elem(), TrafficMirrorFilterRuleMap{})
	pulumi.RegisterOutputType(TrafficMirrorFilterRuleOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(TrafficMirrorFilterRuleMapOutput{})
}
