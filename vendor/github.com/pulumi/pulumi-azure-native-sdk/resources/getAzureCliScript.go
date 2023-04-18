// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a deployment script with a given name.
// API Version: 2020-10-01.
func LookupAzureCliScript(ctx *pulumi.Context, args *LookupAzureCliScriptArgs, opts ...pulumi.InvokeOption) (*LookupAzureCliScriptResult, error) {
	var rv LookupAzureCliScriptResult
	err := ctx.Invoke("azure-native:resources:getAzureCliScript", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAzureCliScriptArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment script.
	ScriptName string `pulumi:"scriptName"`
}

// Object model for the Azure CLI script.
type LookupAzureCliScriptResult struct {
	// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
	Arguments *string `pulumi:"arguments"`
	// Azure CLI module version to be used.
	AzCliVersion string `pulumi:"azCliVersion"`
	// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
	CleanupPreference *string `pulumi:"cleanupPreference"`
	// Container settings.
	ContainerSettings *ContainerConfigurationResponse `pulumi:"containerSettings"`
	// The environment variables to pass over to the script.
	EnvironmentVariables []EnvironmentVariableResponse `pulumi:"environmentVariables"`
	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
	ForceUpdateTag *string `pulumi:"forceUpdateTag"`
	// String Id used to locate any resource on Azure.
	Id string `pulumi:"id"`
	// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// Type of the script.
	// Expected value is 'AzureCLI'.
	Kind string `pulumi:"kind"`
	// The location of the ACI and the storage account for the deployment script.
	Location string `pulumi:"location"`
	// Name of this resource.
	Name string `pulumi:"name"`
	// List of script outputs.
	Outputs map[string]interface{} `pulumi:"outputs"`
	// Uri for the script. This is the entry point for the external script.
	PrimaryScriptUri *string `pulumi:"primaryScriptUri"`
	// State of the script execution. This only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
	RetentionInterval string `pulumi:"retentionInterval"`
	// Script body.
	ScriptContent *string `pulumi:"scriptContent"`
	// Contains the results of script execution.
	Status ScriptStatusResponse `pulumi:"status"`
	// Storage Account settings.
	StorageAccountSettings *StorageAccountConfigurationResponse `pulumi:"storageAccountSettings"`
	// Supporting files for the external script.
	SupportingScriptUris []string `pulumi:"supportingScriptUris"`
	// The system metadata related to this resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
	Timeout *string `pulumi:"timeout"`
	// Type of this resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAzureCliScriptResult
func (val *LookupAzureCliScriptResult) Defaults() *LookupAzureCliScriptResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CleanupPreference) {
		cleanupPreference_ := "Always"
		tmp.CleanupPreference = &cleanupPreference_
	}
	if isZero(tmp.Timeout) {
		timeout_ := "P1D"
		tmp.Timeout = &timeout_
	}
	return &tmp
}

func LookupAzureCliScriptOutput(ctx *pulumi.Context, args LookupAzureCliScriptOutputArgs, opts ...pulumi.InvokeOption) LookupAzureCliScriptResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureCliScriptResult, error) {
			args := v.(LookupAzureCliScriptArgs)
			r, err := LookupAzureCliScript(ctx, &args, opts...)
			var s LookupAzureCliScriptResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureCliScriptResultOutput)
}

type LookupAzureCliScriptOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the deployment script.
	ScriptName pulumi.StringInput `pulumi:"scriptName"`
}

func (LookupAzureCliScriptOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureCliScriptArgs)(nil)).Elem()
}

// Object model for the Azure CLI script.
type LookupAzureCliScriptResultOutput struct{ *pulumi.OutputState }

func (LookupAzureCliScriptResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureCliScriptResult)(nil)).Elem()
}

func (o LookupAzureCliScriptResultOutput) ToLookupAzureCliScriptResultOutput() LookupAzureCliScriptResultOutput {
	return o
}

func (o LookupAzureCliScriptResultOutput) ToLookupAzureCliScriptResultOutputWithContext(ctx context.Context) LookupAzureCliScriptResultOutput {
	return o
}

// Command line arguments to pass to the script. Arguments are separated by spaces. ex: -Name blue* -Location 'West US 2'
func (o LookupAzureCliScriptResultOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

// Azure CLI module version to be used.
func (o LookupAzureCliScriptResultOutput) AzCliVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.AzCliVersion }).(pulumi.StringOutput)
}

// The clean up preference when the script execution gets in a terminal state. Default setting is 'Always'.
func (o LookupAzureCliScriptResultOutput) CleanupPreference() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.CleanupPreference }).(pulumi.StringPtrOutput)
}

// Container settings.
func (o LookupAzureCliScriptResultOutput) ContainerSettings() ContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *ContainerConfigurationResponse { return v.ContainerSettings }).(ContainerConfigurationResponsePtrOutput)
}

// The environment variables to pass over to the script.
func (o LookupAzureCliScriptResultOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) []EnvironmentVariableResponse { return v.EnvironmentVariables }).(EnvironmentVariableResponseArrayOutput)
}

// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID.
func (o LookupAzureCliScriptResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

// String Id used to locate any resource on Azure.
func (o LookupAzureCliScriptResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional property. Managed identity to be used for this deployment script. Currently, only user-assigned MSI is supported.
func (o LookupAzureCliScriptResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// Type of the script.
// Expected value is 'AzureCLI'.
func (o LookupAzureCliScriptResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The location of the ACI and the storage account for the deployment script.
func (o LookupAzureCliScriptResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of this resource.
func (o LookupAzureCliScriptResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of script outputs.
func (o LookupAzureCliScriptResultOutput) Outputs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) map[string]interface{} { return v.Outputs }).(pulumi.MapOutput)
}

// Uri for the script. This is the entry point for the external script.
func (o LookupAzureCliScriptResultOutput) PrimaryScriptUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.PrimaryScriptUri }).(pulumi.StringPtrOutput)
}

// State of the script execution. This only appears in the response.
func (o LookupAzureCliScriptResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. Duration is based on ISO 8601 pattern (for example P1D means one day).
func (o LookupAzureCliScriptResultOutput) RetentionInterval() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.RetentionInterval }).(pulumi.StringOutput)
}

// Script body.
func (o LookupAzureCliScriptResultOutput) ScriptContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.ScriptContent }).(pulumi.StringPtrOutput)
}

// Contains the results of script execution.
func (o LookupAzureCliScriptResultOutput) Status() ScriptStatusResponseOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) ScriptStatusResponse { return v.Status }).(ScriptStatusResponseOutput)
}

// Storage Account settings.
func (o LookupAzureCliScriptResultOutput) StorageAccountSettings() StorageAccountConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *StorageAccountConfigurationResponse {
		return v.StorageAccountSettings
	}).(StorageAccountConfigurationResponsePtrOutput)
}

// Supporting files for the external script.
func (o LookupAzureCliScriptResultOutput) SupportingScriptUris() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) []string { return v.SupportingScriptUris }).(pulumi.StringArrayOutput)
}

// The system metadata related to this resource.
func (o LookupAzureCliScriptResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupAzureCliScriptResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Maximum allowed script execution time specified in ISO 8601 format. Default value is P1D
func (o LookupAzureCliScriptResultOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

// Type of this resource.
func (o LookupAzureCliScriptResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureCliScriptResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureCliScriptResultOutput{})
}
