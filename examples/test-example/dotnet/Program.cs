using System.Collections.Generic;
using System;

using Pulumi;
using Pulumi.Bitlaunch;

return await Deployment.RunAsync(() =>
{
    var token = Environment.GetEnvironmentVariable("BITLAUNCH_TOKEN");

    if(String.IsNullOrEmpty(token))
        throw new Exception("BITLAUNCH_TOKEN env. var not present");
    
    Pulumi.Bitlaunch.Config.Token = token;
    
    var providerArgs = new ProviderArgs
    {
        Token = token
    };
    var provider = new Provider("BitlaunchProvider", providerArgs);

    // Add your resources here
    // e.g. var resource = new Resource("name", new ResourceArgs { });        
    var region = Pulumi.Bitlaunch.GetRegion.Invoke(new()
        {
            Host = "BitLaunch",
            RegionName = "Bucharest",
        });

    var image = Pulumi.Bitlaunch.GetImage.Invoke(new()
        {
            Host = "DigitalOcean",
            DistroName = "Ubuntu",
            VersionName = "24.04 (LTS) x64",
        });

    var size = Pulumi.Bitlaunch.GetSize.Invoke(new ()
        {
            CpuCount = 1,
            Host = "BitLaunch",
            MemoryMb = 1024,
        });

    var serverArgs = new ServerArgs();
    serverArgs.Host = "BitLaunch";
    serverArgs.ImageId = image.Apply(img => img.Id);
    serverArgs.RegionId = region.Apply(region => region.Id);
    serverArgs.SizeId = size.Apply(size => size.Id);
    serverArgs.WaitForIp = true;
    
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
