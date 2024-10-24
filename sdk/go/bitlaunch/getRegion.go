// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bitlaunch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/webwarrior-ws/pulumi-bitlaunch/sdk/go/bitlaunch/internal"
)

// Holds available region configurations for a server. Matches https://developers.bitlaunch.io/reference/host-region-object
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/webwarrior-ws/pulumi-bitlaunch/sdk/go/bitlaunch"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			token := cfg.RequireObject("token")
//			_, err := bitlaunch.GetRegion(ctx, &bitlaunch.GetRegionArgs{
//				Host:       "DigitalOcean",
//				RegionName: pulumi.StringRef("New York"),
//				Slug:       pulumi.StringRef("nyc1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetRegion(ctx *pulumi.Context, args *GetRegionArgs, opts ...pulumi.InvokeOption) (*GetRegionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRegionResult
	err := ctx.Invoke("bitlaunch:index/getRegion:getRegion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegion.
type GetRegionArgs struct {
	// Host Provider (DigitalOcean, Vultr, etc.)
	Host string `pulumi:"host"`
	// The name of the Region.
	RegionName *string `pulumi:"regionName"`
	// The Specific Subregion slug.
	Slug *string `pulumi:"slug"`
}

// A collection of values returned by getRegion.
type GetRegionResult struct {
	// Host Provider (DigitalOcean, Vultr, etc.)
	Host string `pulumi:"host"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ISO code for the region.
	Iso string `pulumi:"iso"`
	// The name of the Region.
	RegionName *string `pulumi:"regionName"`
	// The Specific Subregion slug.
	Slug *string `pulumi:"slug"`
	// A list of the unavailable sizes for this subregion.
	UnavailableSizes []string `pulumi:"unavailableSizes"`
}

func GetRegionOutput(ctx *pulumi.Context, args GetRegionOutputArgs, opts ...pulumi.InvokeOption) GetRegionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRegionResult, error) {
			args := v.(GetRegionArgs)
			r, err := GetRegion(ctx, &args, opts...)
			var s GetRegionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRegionResultOutput)
}

// A collection of arguments for invoking getRegion.
type GetRegionOutputArgs struct {
	// Host Provider (DigitalOcean, Vultr, etc.)
	Host pulumi.StringInput `pulumi:"host"`
	// The name of the Region.
	RegionName pulumi.StringPtrInput `pulumi:"regionName"`
	// The Specific Subregion slug.
	Slug pulumi.StringPtrInput `pulumi:"slug"`
}

func (GetRegionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionArgs)(nil)).Elem()
}

// A collection of values returned by getRegion.
type GetRegionResultOutput struct{ *pulumi.OutputState }

func (GetRegionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRegionResult)(nil)).Elem()
}

func (o GetRegionResultOutput) ToGetRegionResultOutput() GetRegionResultOutput {
	return o
}

func (o GetRegionResultOutput) ToGetRegionResultOutputWithContext(ctx context.Context) GetRegionResultOutput {
	return o
}

// Host Provider (DigitalOcean, Vultr, etc.)
func (o GetRegionResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionResult) string { return v.Host }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRegionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The ISO code for the region.
func (o GetRegionResultOutput) Iso() pulumi.StringOutput {
	return o.ApplyT(func(v GetRegionResult) string { return v.Iso }).(pulumi.StringOutput)
}

// The name of the Region.
func (o GetRegionResultOutput) RegionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegionResult) *string { return v.RegionName }).(pulumi.StringPtrOutput)
}

// The Specific Subregion slug.
func (o GetRegionResultOutput) Slug() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRegionResult) *string { return v.Slug }).(pulumi.StringPtrOutput)
}

// A list of the unavailable sizes for this subregion.
func (o GetRegionResultOutput) UnavailableSizes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRegionResult) []string { return v.UnavailableSizes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRegionResultOutput{})
}
