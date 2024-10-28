# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetImageResult',
    'AwaitableGetImageResult',
    'get_image',
    'get_image_output',
]

@pulumi.output_type
class GetImageResult:
    """
    A collection of values returned by getImage.
    """
    def __init__(__self__, distro_name=None, extra_cost_per_month=None, host=None, id=None, is_windows=None, min_disk_size=None, password_unsupported=None, type=None, unavailable_regions=None, version_name=None):
        if distro_name and not isinstance(distro_name, str):
            raise TypeError("Expected argument 'distro_name' to be a str")
        pulumi.set(__self__, "distro_name", distro_name)
        if extra_cost_per_month and not isinstance(extra_cost_per_month, int):
            raise TypeError("Expected argument 'extra_cost_per_month' to be a int")
        pulumi.set(__self__, "extra_cost_per_month", extra_cost_per_month)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_windows and not isinstance(is_windows, bool):
            raise TypeError("Expected argument 'is_windows' to be a bool")
        pulumi.set(__self__, "is_windows", is_windows)
        if min_disk_size and not isinstance(min_disk_size, int):
            raise TypeError("Expected argument 'min_disk_size' to be a int")
        pulumi.set(__self__, "min_disk_size", min_disk_size)
        if password_unsupported and not isinstance(password_unsupported, bool):
            raise TypeError("Expected argument 'password_unsupported' to be a bool")
        pulumi.set(__self__, "password_unsupported", password_unsupported)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unavailable_regions and not isinstance(unavailable_regions, list):
            raise TypeError("Expected argument 'unavailable_regions' to be a list")
        pulumi.set(__self__, "unavailable_regions", unavailable_regions)
        if version_name and not isinstance(version_name, str):
            raise TypeError("Expected argument 'version_name' to be a str")
        pulumi.set(__self__, "version_name", version_name)

    @property
    @pulumi.getter(name="distroName")
    def distro_name(self) -> Optional[str]:
        """
        The name of the Linux Distibution or one-click app.
        """
        return pulumi.get(self, "distro_name")

    @property
    @pulumi.getter(name="extraCostPerMonth")
    def extra_cost_per_month(self) -> int:
        """
        Extra monthly cost.
        """
        return pulumi.get(self, "extra_cost_per_month")

    @property
    @pulumi.getter
    def host(self) -> str:
        """
        Host Provider (DigitalOcean, Vultr, etc.)
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isWindows")
    def is_windows(self) -> bool:
        """
        Flag to determine if the image is Windows-based.
        """
        return pulumi.get(self, "is_windows")

    @property
    @pulumi.getter(name="minDiskSize")
    def min_disk_size(self) -> int:
        """
        The minimum disk size available in GB.
        """
        return pulumi.get(self, "min_disk_size")

    @property
    @pulumi.getter(name="passwordUnsupported")
    def password_unsupported(self) -> bool:
        """
        If setting a password is supported.
        """
        return pulumi.get(self, "password_unsupported")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the image: image or app.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="unavailableRegions")
    def unavailable_regions(self) -> Sequence[str]:
        """
        A list of unavailable subregion IDs.
        """
        return pulumi.get(self, "unavailable_regions")

    @property
    @pulumi.getter(name="versionName")
    def version_name(self) -> Optional[str]:
        """
        The Specific Image Version
        """
        return pulumi.get(self, "version_name")


class AwaitableGetImageResult(GetImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageResult(
            distro_name=self.distro_name,
            extra_cost_per_month=self.extra_cost_per_month,
            host=self.host,
            id=self.id,
            is_windows=self.is_windows,
            min_disk_size=self.min_disk_size,
            password_unsupported=self.password_unsupported,
            type=self.type,
            unavailable_regions=self.unavailable_regions,
            version_name=self.version_name)


def get_image(distro_name: Optional[str] = None,
              host: Optional[str] = None,
              version_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageResult:
    """
    Holds details on Images and apps available when configuring a server. Matches https://developers.bitlaunch.io/reference/host-image-object

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_bitlaunch as bitlaunch

    config = pulumi.Config()
    token = config.require_object("token")
    example = bitlaunch.get_image(host="DigitalOcean",
        distro_name="Ubuntu")
    ```
    <!--End PulumiCodeChooser -->


    :param str distro_name: The name of the Linux Distibution or one-click app.
    :param str host: Host Provider (DigitalOcean, Vultr, etc.)
    :param str version_name: The Specific Image Version
    """
    __args__ = dict()
    __args__['distroName'] = distro_name
    __args__['host'] = host
    __args__['versionName'] = version_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('bitlaunch:index/getImage:getImage', __args__, opts=opts, typ=GetImageResult).value

    return AwaitableGetImageResult(
        distro_name=pulumi.get(__ret__, 'distro_name'),
        extra_cost_per_month=pulumi.get(__ret__, 'extra_cost_per_month'),
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        is_windows=pulumi.get(__ret__, 'is_windows'),
        min_disk_size=pulumi.get(__ret__, 'min_disk_size'),
        password_unsupported=pulumi.get(__ret__, 'password_unsupported'),
        type=pulumi.get(__ret__, 'type'),
        unavailable_regions=pulumi.get(__ret__, 'unavailable_regions'),
        version_name=pulumi.get(__ret__, 'version_name'))


@_utilities.lift_output_func(get_image)
def get_image_output(distro_name: Optional[pulumi.Input[Optional[str]]] = None,
                     host: Optional[pulumi.Input[str]] = None,
                     version_name: Optional[pulumi.Input[Optional[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetImageResult]:
    """
    Holds details on Images and apps available when configuring a server. Matches https://developers.bitlaunch.io/reference/host-image-object

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_bitlaunch as bitlaunch

    config = pulumi.Config()
    token = config.require_object("token")
    example = bitlaunch.get_image(host="DigitalOcean",
        distro_name="Ubuntu")
    ```
    <!--End PulumiCodeChooser -->


    :param str distro_name: The name of the Linux Distibution or one-click app.
    :param str host: Host Provider (DigitalOcean, Vultr, etc.)
    :param str version_name: The Specific Image Version
    """
    ...