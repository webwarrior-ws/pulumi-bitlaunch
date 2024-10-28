open System
open Pulumi
open Pulumi.Bitlaunch
open Pulumi.FSharp

let Deploy () = 
    let token = Environment.GetEnvironmentVariable("BITLAUNCH_TOKEN")

    if String.IsNullOrEmpty(token) then
        raise (Exception("BITLAUNCH_TOKEN env. var not present"))

    let password = Environment.GetEnvironmentVariable("SERVER_PASSWORD")

    if String.IsNullOrEmpty(password) then
        raise (Exception("SERVER_PASSWORD env. var not present"))

    let providerArgs = ProviderArgs(Token = token)
    let provider = Provider("BitlaunchProvider", providerArgs)
    let invokeOptions = InvokeOptions(Provider = provider)

    let hostName = "BitLaunch"

    let region = 
        Pulumi.Bitlaunch.GetRegion.Invoke(
            GetRegionInvokeArgs(Host = hostName, RegionName = "Bucharest"),
            invokeOptions
        )

    let image = 
        Pulumi.Bitlaunch.GetImage.Invoke(
            GetImageInvokeArgs(Host = hostName, DistroName = "Ubuntu", VersionName = "Ubuntu 24.04 LTS"),
            invokeOptions
        )

    let size = 
        Pulumi.Bitlaunch.GetSize.Invoke(
            GetSizeInvokeArgs(CpuCount = 1, Host = hostName, MemoryMb = 1024),
            invokeOptions
        )

    let serverArgs = ServerArgs(
        Host = hostName,
        ImageId = io (image.Apply(fun img -> img.Id)),
        RegionId = io (region.Apply(fun region -> region.Id)),
        SizeId = io (size.Apply(fun size -> size.Id)),
        WaitForIp = true,
        Password = password
    )

    let server = 
        Pulumi.Bitlaunch.Server(
            "testServer", 
            serverArgs, 
            CustomResourceOptions(Provider = provider)
        )

    // Export outputs here
    dict [ "ipAddress", box server.Ipv4 ]

[<EntryPoint>]
let Main _ = Deployment.run Deploy
