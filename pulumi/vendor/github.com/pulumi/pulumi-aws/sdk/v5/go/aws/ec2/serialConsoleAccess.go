// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage whether serial console access is enabled for your AWS account in the current AWS region.
//
// > **NOTE:** Removing this resource disables serial console access.
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
//			_, err := ec2.NewSerialConsoleAccess(ctx, "example", &ec2.SerialConsoleAccessArgs{
//				Enabled: pulumi.Bool(true),
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
// Serial console access state can be imported, e.g.,
//
// ```sh
//
//	$ pulumi import aws:ec2/serialConsoleAccess:SerialConsoleAccess example default
//
// ```
type SerialConsoleAccess struct {
	pulumi.CustomResourceState

	// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
}

// NewSerialConsoleAccess registers a new resource with the given unique name, arguments, and options.
func NewSerialConsoleAccess(ctx *pulumi.Context,
	name string, args *SerialConsoleAccessArgs, opts ...pulumi.ResourceOption) (*SerialConsoleAccess, error) {
	if args == nil {
		args = &SerialConsoleAccessArgs{}
	}

	var resource SerialConsoleAccess
	err := ctx.RegisterResource("aws:ec2/serialConsoleAccess:SerialConsoleAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSerialConsoleAccess gets an existing SerialConsoleAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSerialConsoleAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SerialConsoleAccessState, opts ...pulumi.ResourceOption) (*SerialConsoleAccess, error) {
	var resource SerialConsoleAccess
	err := ctx.ReadResource("aws:ec2/serialConsoleAccess:SerialConsoleAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SerialConsoleAccess resources.
type serialConsoleAccessState struct {
	// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

type SerialConsoleAccessState struct {
	// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (SerialConsoleAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*serialConsoleAccessState)(nil)).Elem()
}

type serialConsoleAccessArgs struct {
	// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
}

// The set of arguments for constructing a SerialConsoleAccess resource.
type SerialConsoleAccessArgs struct {
	// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
}

func (SerialConsoleAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serialConsoleAccessArgs)(nil)).Elem()
}

type SerialConsoleAccessInput interface {
	pulumi.Input

	ToSerialConsoleAccessOutput() SerialConsoleAccessOutput
	ToSerialConsoleAccessOutputWithContext(ctx context.Context) SerialConsoleAccessOutput
}

func (*SerialConsoleAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**SerialConsoleAccess)(nil)).Elem()
}

func (i *SerialConsoleAccess) ToSerialConsoleAccessOutput() SerialConsoleAccessOutput {
	return i.ToSerialConsoleAccessOutputWithContext(context.Background())
}

func (i *SerialConsoleAccess) ToSerialConsoleAccessOutputWithContext(ctx context.Context) SerialConsoleAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SerialConsoleAccessOutput)
}

// SerialConsoleAccessArrayInput is an input type that accepts SerialConsoleAccessArray and SerialConsoleAccessArrayOutput values.
// You can construct a concrete instance of `SerialConsoleAccessArrayInput` via:
//
//	SerialConsoleAccessArray{ SerialConsoleAccessArgs{...} }
type SerialConsoleAccessArrayInput interface {
	pulumi.Input

	ToSerialConsoleAccessArrayOutput() SerialConsoleAccessArrayOutput
	ToSerialConsoleAccessArrayOutputWithContext(context.Context) SerialConsoleAccessArrayOutput
}

type SerialConsoleAccessArray []SerialConsoleAccessInput

func (SerialConsoleAccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SerialConsoleAccess)(nil)).Elem()
}

func (i SerialConsoleAccessArray) ToSerialConsoleAccessArrayOutput() SerialConsoleAccessArrayOutput {
	return i.ToSerialConsoleAccessArrayOutputWithContext(context.Background())
}

func (i SerialConsoleAccessArray) ToSerialConsoleAccessArrayOutputWithContext(ctx context.Context) SerialConsoleAccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SerialConsoleAccessArrayOutput)
}

// SerialConsoleAccessMapInput is an input type that accepts SerialConsoleAccessMap and SerialConsoleAccessMapOutput values.
// You can construct a concrete instance of `SerialConsoleAccessMapInput` via:
//
//	SerialConsoleAccessMap{ "key": SerialConsoleAccessArgs{...} }
type SerialConsoleAccessMapInput interface {
	pulumi.Input

	ToSerialConsoleAccessMapOutput() SerialConsoleAccessMapOutput
	ToSerialConsoleAccessMapOutputWithContext(context.Context) SerialConsoleAccessMapOutput
}

type SerialConsoleAccessMap map[string]SerialConsoleAccessInput

func (SerialConsoleAccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SerialConsoleAccess)(nil)).Elem()
}

func (i SerialConsoleAccessMap) ToSerialConsoleAccessMapOutput() SerialConsoleAccessMapOutput {
	return i.ToSerialConsoleAccessMapOutputWithContext(context.Background())
}

func (i SerialConsoleAccessMap) ToSerialConsoleAccessMapOutputWithContext(ctx context.Context) SerialConsoleAccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SerialConsoleAccessMapOutput)
}

type SerialConsoleAccessOutput struct{ *pulumi.OutputState }

func (SerialConsoleAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SerialConsoleAccess)(nil)).Elem()
}

func (o SerialConsoleAccessOutput) ToSerialConsoleAccessOutput() SerialConsoleAccessOutput {
	return o
}

func (o SerialConsoleAccessOutput) ToSerialConsoleAccessOutputWithContext(ctx context.Context) SerialConsoleAccessOutput {
	return o
}

// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
func (o SerialConsoleAccessOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SerialConsoleAccess) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type SerialConsoleAccessArrayOutput struct{ *pulumi.OutputState }

func (SerialConsoleAccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SerialConsoleAccess)(nil)).Elem()
}

func (o SerialConsoleAccessArrayOutput) ToSerialConsoleAccessArrayOutput() SerialConsoleAccessArrayOutput {
	return o
}

func (o SerialConsoleAccessArrayOutput) ToSerialConsoleAccessArrayOutputWithContext(ctx context.Context) SerialConsoleAccessArrayOutput {
	return o
}

func (o SerialConsoleAccessArrayOutput) Index(i pulumi.IntInput) SerialConsoleAccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SerialConsoleAccess {
		return vs[0].([]*SerialConsoleAccess)[vs[1].(int)]
	}).(SerialConsoleAccessOutput)
}

type SerialConsoleAccessMapOutput struct{ *pulumi.OutputState }

func (SerialConsoleAccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SerialConsoleAccess)(nil)).Elem()
}

func (o SerialConsoleAccessMapOutput) ToSerialConsoleAccessMapOutput() SerialConsoleAccessMapOutput {
	return o
}

func (o SerialConsoleAccessMapOutput) ToSerialConsoleAccessMapOutputWithContext(ctx context.Context) SerialConsoleAccessMapOutput {
	return o
}

func (o SerialConsoleAccessMapOutput) MapIndex(k pulumi.StringInput) SerialConsoleAccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SerialConsoleAccess {
		return vs[0].(map[string]*SerialConsoleAccess)[vs[1].(string)]
	}).(SerialConsoleAccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SerialConsoleAccessInput)(nil)).Elem(), &SerialConsoleAccess{})
	pulumi.RegisterInputType(reflect.TypeOf((*SerialConsoleAccessArrayInput)(nil)).Elem(), SerialConsoleAccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SerialConsoleAccessMapInput)(nil)).Elem(), SerialConsoleAccessMap{})
	pulumi.RegisterOutputType(SerialConsoleAccessOutput{})
	pulumi.RegisterOutputType(SerialConsoleAccessArrayOutput{})
	pulumi.RegisterOutputType(SerialConsoleAccessMapOutput{})
}
