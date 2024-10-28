using System.Collections.Generic;
using System;

using Pulumi;
using Pulumi.Bitlaunch;

return await Deployment.RunAsync(() =>
{
    var token = Environment.GetEnvironmentVariable("BITLAUNCH_TOKEN");

    if(String.IsNullOrEmpty(token))
        throw new Exception("BITLAUNCH_TOKEN env. var not present");

    var password = Environment.GetEnvironmentVariable("SERVER_PASSWORD");

    if(String.IsNullOrEmpty(password))
        throw new Exception("SERVER_PASSWORD env. var not present");
    
    // If you set token in code, you need to use custom provider, otherwise will get an error:
    // Provider is missing a required configuration key, try `pulumi config set bitlaunch:token`: API Token
    var providerArgs = new ProviderArgs
    {
        Token = token
    };
    var provider = new Provider("BitlaunchProvider", providerArgs);
    var invokeOptions = new InvokeOptions { Provider = provider };

    var hostName = "BitLaunch";

    // Add your resources here
    // e.g. var resource = new Resource("name", new ResourceArgs { });        
    var region = Pulumi.Bitlaunch.GetRegion.Invoke(new()
        {
            Host = hostName,
            RegionName = "Bucharest",
        },
        invokeOptions);

    var image = Pulumi.Bitlaunch.GetImage.Invoke(new()
        {
            Host = hostName,
            DistroName = "Ubuntu",
            VersionName = "Ubuntu 24.04 LTS",
        },
        invokeOptions);

    var size = Pulumi.Bitlaunch.GetSize.Invoke(new ()
        {
            CpuCount = 1,
            Host = hostName,
            MemoryMb = 1024,
        },
        invokeOptions);

    var serverArgs = new ServerArgs
    {
        Host = hostName,
        ImageId = image.Apply(img => img.Id),
        RegionId = region.Apply(region => region.Id),
        SizeId = size.Apply(size => size.Id),
        WaitForIp = true,
        Password = password,
    };

    var server = new Pulumi.Bitlaunch.Server(
            "testServer", 
            serverArgs, 
            new CustomResourceOptions{ Provider=provider }
        );

    // Export outputs here
    return new Dictionary<string, object?>
    {
        ["ipAddress"] = server.Ipv4
    };
});
