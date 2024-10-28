// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bitlaunch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/webwarrior-ws/pulumi-bitlaunch/sdk/go/bitlaunch/internal"
)

// Virtual Machine Resource. Matches https://developers.bitlaunch.io/reference/server-object
type Server struct {
	pulumi.CustomResourceState

	// The creation date of the server.
	Created pulumi.StringOutput `pulumi:"created"`
	// The host for the server to reside on.
	Host pulumi.StringOutput `pulumi:"host"`
	// The description of the image installed on the server.
	ImageDescription pulumi.StringOutput `pulumi:"imageDescription"`
	// The image ID to use on the server.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// A script to run on first boot of the server. Only hosts with initScript enabled can use this feature.
	Initscript pulumi.StringPtrOutput `pulumi:"initscript"`
	// The name of the key.
	Ipv4 pulumi.StringOutput `pulumi:"ipv4"`
	// The name of the server.
	Name pulumi.StringOutput `pulumi:"name"`
	// The root user password to set on the server. Must be used if no SSH keys designated.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The hourly rate of the server that will be deducted from your account balance every hour.
	Rate pulumi.IntOutput `pulumi:"rate"`
	// The region ID of the location that the server will reside at.
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// The size ID of the server to be provisioned to.
	SizeId pulumi.StringOutput `pulumi:"sizeId"`
	// An array of SSH key IDs to place on the server for authentication. Must be used if no password is designated of if the selected image does not support passwords.
	SshKeys pulumi.StringArrayOutput `pulumi:"sshKeys"`
	// The name of the key.
	Status pulumi.StringOutput `pulumi:"status"`
	// Wait to get IP Address
	WaitForIp pulumi.BoolPtrOutput `pulumi:"waitForIp"`
}

// NewServer registers a new resource with the given unique name, arguments, and options.
func NewServer(ctx *pulumi.Context,
	name string, args *ServerArgs, opts ...pulumi.ResourceOption) (*Server, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.RegionId == nil {
		return nil, errors.New("invalid value for required argument 'RegionId'")
	}
	if args.SizeId == nil {
		return nil, errors.New("invalid value for required argument 'SizeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Server
	err := ctx.RegisterResource("bitlaunch:index/server:Server", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServer gets an existing Server resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerState, opts ...pulumi.ResourceOption) (*Server, error) {
	var resource Server
	err := ctx.ReadResource("bitlaunch:index/server:Server", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Server resources.
type serverState struct {
	// The creation date of the server.
	Created *string `pulumi:"created"`
	// The host for the server to reside on.
	Host *string `pulumi:"host"`
	// The description of the image installed on the server.
	ImageDescription *string `pulumi:"imageDescription"`
	// The image ID to use on the server.
	ImageId *string `pulumi:"imageId"`
	// A script to run on first boot of the server. Only hosts with initScript enabled can use this feature.
	Initscript *string `pulumi:"initscript"`
	// The name of the key.
	Ipv4 *string `pulumi:"ipv4"`
	// The name of the server.
	Name *string `pulumi:"name"`
	// The root user password to set on the server. Must be used if no SSH keys designated.
	Password *string `pulumi:"password"`
	// The hourly rate of the server that will be deducted from your account balance every hour.
	Rate *int `pulumi:"rate"`
	// The region ID of the location that the server will reside at.
	RegionId *string `pulumi:"regionId"`
	// The size ID of the server to be provisioned to.
	SizeId *string `pulumi:"sizeId"`
	// An array of SSH key IDs to place on the server for authentication. Must be used if no password is designated of if the selected image does not support passwords.
	SshKeys []string `pulumi:"sshKeys"`
	// The name of the key.
	Status *string `pulumi:"status"`
	// Wait to get IP Address
	WaitForIp *bool `pulumi:"waitForIp"`
}

type ServerState struct {
	// The creation date of the server.
	Created pulumi.StringPtrInput
	// The host for the server to reside on.
	Host pulumi.StringPtrInput
	// The description of the image installed on the server.
	ImageDescription pulumi.StringPtrInput
	// The image ID to use on the server.
	ImageId pulumi.StringPtrInput
	// A script to run on first boot of the server. Only hosts with initScript enabled can use this feature.
	Initscript pulumi.StringPtrInput
	// The name of the key.
	Ipv4 pulumi.StringPtrInput
	// The name of the server.
	Name pulumi.StringPtrInput
	// The root user password to set on the server. Must be used if no SSH keys designated.
	Password pulumi.StringPtrInput
	// The hourly rate of the server that will be deducted from your account balance every hour.
	Rate pulumi.IntPtrInput
	// The region ID of the location that the server will reside at.
	RegionId pulumi.StringPtrInput
	// The size ID of the server to be provisioned to.
	SizeId pulumi.StringPtrInput
	// An array of SSH key IDs to place on the server for authentication. Must be used if no password is designated of if the selected image does not support passwords.
	SshKeys pulumi.StringArrayInput
	// The name of the key.
	Status pulumi.StringPtrInput
	// Wait to get IP Address
	WaitForIp pulumi.BoolPtrInput
}

func (ServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverState)(nil)).Elem()
}

type serverArgs struct {
	// The host for the server to reside on.
	Host string `pulumi:"host"`
	// The image ID to use on the server.
	ImageId string `pulumi:"imageId"`
	// A script to run on first boot of the server. Only hosts with initScript enabled can use this feature.
	Initscript *string `pulumi:"initscript"`
	// The name of the server.
	Name *string `pulumi:"name"`
	// The root user password to set on the server. Must be used if no SSH keys designated.
	Password *string `pulumi:"password"`
	// The region ID of the location that the server will reside at.
	RegionId string `pulumi:"regionId"`
	// The size ID of the server to be provisioned to.
	SizeId string `pulumi:"sizeId"`
	// An array of SSH key IDs to place on the server for authentication. Must be used if no password is designated of if the selected image does not support passwords.
	SshKeys []string `pulumi:"sshKeys"`
	// Wait to get IP Address
	WaitForIp *bool `pulumi:"waitForIp"`
}

// The set of arguments for constructing a Server resource.
type ServerArgs struct {
	// The host for the server to reside on.
	Host pulumi.StringInput
	// The image ID to use on the server.
	ImageId pulumi.StringInput
	// A script to run on first boot of the server. Only hosts with initScript enabled can use this feature.
	Initscript pulumi.StringPtrInput
	// The name of the server.
	Name pulumi.StringPtrInput
	// The root user password to set on the server. Must be used if no SSH keys designated.
	Password pulumi.StringPtrInput
	// The region ID of the location that the server will reside at.
	RegionId pulumi.StringInput
	// The size ID of the server to be provisioned to.
	SizeId pulumi.StringInput
	// An array of SSH key IDs to place on the server for authentication. Must be used if no password is designated of if the selected image does not support passwords.
	SshKeys pulumi.StringArrayInput
	// Wait to get IP Address
	WaitForIp pulumi.BoolPtrInput
}

func (ServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverArgs)(nil)).Elem()
}

type ServerInput interface {
	pulumi.Input

	ToServerOutput() ServerOutput
	ToServerOutputWithContext(ctx context.Context) ServerOutput
}

func (*Server) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (i *Server) ToServerOutput() ServerOutput {
	return i.ToServerOutputWithContext(context.Background())
}

func (i *Server) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerOutput)
}

// ServerArrayInput is an input type that accepts ServerArray and ServerArrayOutput values.
// You can construct a concrete instance of `ServerArrayInput` via:
//
//	ServerArray{ ServerArgs{...} }
type ServerArrayInput interface {
	pulumi.Input

	ToServerArrayOutput() ServerArrayOutput
	ToServerArrayOutputWithContext(context.Context) ServerArrayOutput
}

type ServerArray []ServerInput

func (ServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (i ServerArray) ToServerArrayOutput() ServerArrayOutput {
	return i.ToServerArrayOutputWithContext(context.Background())
}

func (i ServerArray) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerArrayOutput)
}

// ServerMapInput is an input type that accepts ServerMap and ServerMapOutput values.
// You can construct a concrete instance of `ServerMapInput` via:
//
//	ServerMap{ "key": ServerArgs{...} }
type ServerMapInput interface {
	pulumi.Input

	ToServerMapOutput() ServerMapOutput
	ToServerMapOutputWithContext(context.Context) ServerMapOutput
}

type ServerMap map[string]ServerInput

func (ServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (i ServerMap) ToServerMapOutput() ServerMapOutput {
	return i.ToServerMapOutputWithContext(context.Background())
}

func (i ServerMap) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerMapOutput)
}

type ServerOutput struct{ *pulumi.OutputState }

func (ServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Server)(nil)).Elem()
}

func (o ServerOutput) ToServerOutput() ServerOutput {
	return o
}

func (o ServerOutput) ToServerOutputWithContext(ctx context.Context) ServerOutput {
	return o
}

// The creation date of the server.
func (o ServerOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// The host for the server to reside on.
func (o ServerOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The description of the image installed on the server.
func (o ServerOutput) ImageDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ImageDescription }).(pulumi.StringOutput)
}

// The image ID to use on the server.
func (o ServerOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// A script to run on first boot of the server. Only hosts with initScript enabled can use this feature.
func (o ServerOutput) Initscript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Initscript }).(pulumi.StringPtrOutput)
}

// The name of the key.
func (o ServerOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Ipv4 }).(pulumi.StringOutput)
}

// The name of the server.
func (o ServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The root user password to set on the server. Must be used if no SSH keys designated.
func (o ServerOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The hourly rate of the server that will be deducted from your account balance every hour.
func (o ServerOutput) Rate() pulumi.IntOutput {
	return o.ApplyT(func(v *Server) pulumi.IntOutput { return v.Rate }).(pulumi.IntOutput)
}

// The region ID of the location that the server will reside at.
func (o ServerOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// The size ID of the server to be provisioned to.
func (o ServerOutput) SizeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.SizeId }).(pulumi.StringOutput)
}

// An array of SSH key IDs to place on the server for authentication. Must be used if no password is designated of if the selected image does not support passwords.
func (o ServerOutput) SshKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Server) pulumi.StringArrayOutput { return v.SshKeys }).(pulumi.StringArrayOutput)
}

// The name of the key.
func (o ServerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Server) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Wait to get IP Address
func (o ServerOutput) WaitForIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Server) pulumi.BoolPtrOutput { return v.WaitForIp }).(pulumi.BoolPtrOutput)
}

type ServerArrayOutput struct{ *pulumi.OutputState }

func (ServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Server)(nil)).Elem()
}

func (o ServerArrayOutput) ToServerArrayOutput() ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) ToServerArrayOutputWithContext(ctx context.Context) ServerArrayOutput {
	return o
}

func (o ServerArrayOutput) Index(i pulumi.IntInput) ServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Server {
		return vs[0].([]*Server)[vs[1].(int)]
	}).(ServerOutput)
}

type ServerMapOutput struct{ *pulumi.OutputState }

func (ServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Server)(nil)).Elem()
}

func (o ServerMapOutput) ToServerMapOutput() ServerMapOutput {
	return o
}

func (o ServerMapOutput) ToServerMapOutputWithContext(ctx context.Context) ServerMapOutput {
	return o
}

func (o ServerMapOutput) MapIndex(k pulumi.StringInput) ServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Server {
		return vs[0].(map[string]*Server)[vs[1].(string)]
	}).(ServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerInput)(nil)).Elem(), &Server{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerArrayInput)(nil)).Elem(), ServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerMapInput)(nil)).Elem(), ServerMap{})
	pulumi.RegisterOutputType(ServerOutput{})
	pulumi.RegisterOutputType(ServerArrayOutput{})
	pulumi.RegisterOutputType(ServerMapOutput{})
}