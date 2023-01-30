// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetNetworkInterfaces(ctx *pulumi.Context, args *GetNetworkInterfacesArgs, opts ...pulumi.InvokeOption) (*GetNetworkInterfacesResult, error) {
	var rv GetNetworkInterfacesResult
	err := ctx.Invoke("aws:ec2/getNetworkInterfaces:getNetworkInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInterfaces.
type GetNetworkInterfacesArgs struct {
	// Custom filter block as described below.
	Filters []GetNetworkInterfacesFilter `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired network interfaces.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getNetworkInterfaces.
type GetNetworkInterfacesResult struct {
	Filters []GetNetworkInterfacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of all the network interface ids found.
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetNetworkInterfacesOutput(ctx *pulumi.Context, args GetNetworkInterfacesOutputArgs, opts ...pulumi.InvokeOption) GetNetworkInterfacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetNetworkInterfacesResult, error) {
			args := v.(GetNetworkInterfacesArgs)
			r, err := GetNetworkInterfaces(ctx, &args, opts...)
			var s GetNetworkInterfacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetNetworkInterfacesResultOutput)
}

// A collection of arguments for invoking getNetworkInterfaces.
type GetNetworkInterfacesOutputArgs struct {
	// Custom filter block as described below.
	Filters GetNetworkInterfacesFilterArrayInput `pulumi:"filters"`
	// Map of tags, each pair of which must exactly match
	// a pair on the desired network interfaces.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetNetworkInterfacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkInterfacesArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkInterfaces.
type GetNetworkInterfacesResultOutput struct{ *pulumi.OutputState }

func (GetNetworkInterfacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNetworkInterfacesResult)(nil)).Elem()
}

func (o GetNetworkInterfacesResultOutput) ToGetNetworkInterfacesResultOutput() GetNetworkInterfacesResultOutput {
	return o
}

func (o GetNetworkInterfacesResultOutput) ToGetNetworkInterfacesResultOutputWithContext(ctx context.Context) GetNetworkInterfacesResultOutput {
	return o
}

func (o GetNetworkInterfacesResultOutput) Filters() GetNetworkInterfacesFilterArrayOutput {
	return o.ApplyT(func(v GetNetworkInterfacesResult) []GetNetworkInterfacesFilter { return v.Filters }).(GetNetworkInterfacesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNetworkInterfacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNetworkInterfacesResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of all the network interface ids found.
func (o GetNetworkInterfacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNetworkInterfacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetNetworkInterfacesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetNetworkInterfacesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNetworkInterfacesResultOutput{})
}