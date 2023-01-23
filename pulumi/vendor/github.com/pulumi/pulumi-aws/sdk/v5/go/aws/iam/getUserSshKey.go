// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a SSH public key associated with the specified IAM user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.GetUserSshKey(ctx, &iam.GetUserSshKeyArgs{
//				Encoding:       "SSH",
//				SshPublicKeyId: "APKARUZ32GUTKIGARLXE",
//				Username:       "test-user",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetUserSshKey(ctx *pulumi.Context, args *GetUserSshKeyArgs, opts ...pulumi.InvokeOption) (*GetUserSshKeyResult, error) {
	var rv GetUserSshKeyResult
	err := ctx.Invoke("aws:iam/getUserSshKey:getUserSshKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserSshKey.
type GetUserSshKeyArgs struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding string `pulumi:"encoding"`
	// Unique identifier for the SSH public key.
	SshPublicKeyId string `pulumi:"sshPublicKeyId"`
	// Name of the IAM user associated with the SSH public key.
	Username string `pulumi:"username"`
}

// A collection of values returned by getUserSshKey.
type GetUserSshKeyResult struct {
	Encoding string `pulumi:"encoding"`
	// MD5 message digest of the SSH public key.
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// SSH public key.
	PublicKey      string `pulumi:"publicKey"`
	SshPublicKeyId string `pulumi:"sshPublicKeyId"`
	// Status of the SSH public key. Active means that the key can be used for authentication with an CodeCommit repository. Inactive means that the key cannot be used.
	Status   string `pulumi:"status"`
	Username string `pulumi:"username"`
}

func GetUserSshKeyOutput(ctx *pulumi.Context, args GetUserSshKeyOutputArgs, opts ...pulumi.InvokeOption) GetUserSshKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserSshKeyResult, error) {
			args := v.(GetUserSshKeyArgs)
			r, err := GetUserSshKey(ctx, &args, opts...)
			var s GetUserSshKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserSshKeyResultOutput)
}

// A collection of arguments for invoking getUserSshKey.
type GetUserSshKeyOutputArgs struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding pulumi.StringInput `pulumi:"encoding"`
	// Unique identifier for the SSH public key.
	SshPublicKeyId pulumi.StringInput `pulumi:"sshPublicKeyId"`
	// Name of the IAM user associated with the SSH public key.
	Username pulumi.StringInput `pulumi:"username"`
}

func (GetUserSshKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSshKeyArgs)(nil)).Elem()
}

// A collection of values returned by getUserSshKey.
type GetUserSshKeyResultOutput struct{ *pulumi.OutputState }

func (GetUserSshKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserSshKeyResult)(nil)).Elem()
}

func (o GetUserSshKeyResultOutput) ToGetUserSshKeyResultOutput() GetUserSshKeyResultOutput {
	return o
}

func (o GetUserSshKeyResultOutput) ToGetUserSshKeyResultOutputWithContext(ctx context.Context) GetUserSshKeyResultOutput {
	return o
}

func (o GetUserSshKeyResultOutput) Encoding() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.Encoding }).(pulumi.StringOutput)
}

// MD5 message digest of the SSH public key.
func (o GetUserSshKeyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserSshKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

// SSH public key.
func (o GetUserSshKeyResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func (o GetUserSshKeyResultOutput) SshPublicKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.SshPublicKeyId }).(pulumi.StringOutput)
}

// Status of the SSH public key. Active means that the key can be used for authentication with an CodeCommit repository. Inactive means that the key cannot be used.
func (o GetUserSshKeyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o GetUserSshKeyResultOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserSshKeyResult) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserSshKeyResultOutput{})
}
