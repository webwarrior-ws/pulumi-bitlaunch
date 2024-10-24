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
    'GetRegionResult',
    'AwaitableGetRegionResult',
    'get_region',
    'get_region_output',
]

@pulumi.output_type
class GetRegionResult:
    """
    A collection of values returned by getRegion.
    """
    def __init__(__self__, host=None, id=None, iso=None, region_name=None, slug=None, unavailable_sizes=None):
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iso and not isinstance(iso, str):
            raise TypeError("Expected argument 'iso' to be a str")
        pulumi.set(__self__, "iso", iso)
        if region_name and not isinstance(region_name, str):
            raise TypeError("Expected argument 'region_name' to be a str")
        pulumi.set(__self__, "region_name", region_name)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)
        if unavailable_sizes and not isinstance(unavailable_sizes, list):
            raise TypeError("Expected argument 'unavailable_sizes' to be a list")
        pulumi.set(__self__, "unavailable_sizes", unavailable_sizes)

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
    @pulumi.getter
    def iso(self) -> str:
        """
        The ISO code for the region.
        """
        return pulumi.get(self, "iso")

    @property
    @pulumi.getter(name="regionName")
    def region_name(self) -> Optional[str]:
        """
        The name of the Region.
        """
        return pulumi.get(self, "region_name")

    @property
    @pulumi.getter
    def slug(self) -> Optional[str]:
        """
        The Specific Subregion slug.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter(name="unavailableSizes")
    def unavailable_sizes(self) -> Sequence[str]:
        """
        A list of the unavailable sizes for this subregion.
        """
        return pulumi.get(self, "unavailable_sizes")


class AwaitableGetRegionResult(GetRegionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionResult(
            host=self.host,
            id=self.id,
            iso=self.iso,
            region_name=self.region_name,
            slug=self.slug,
            unavailable_sizes=self.unavailable_sizes)


def get_region(host: Optional[str] = None,
               region_name: Optional[str] = None,
               slug: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionResult:
    """
    Holds available region configurations for a server. Matches https://developers.bitlaunch.io/reference/host-region-object

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_bitlaunch as bitlaunch

    config = pulumi.Config()
    token = config.require_object("token")
    example = bitlaunch.get_region(host="DigitalOcean",
        region_name="New York",
        slug="nyc1")
    ```
    <!--End PulumiCodeChooser -->


    :param str host: Host Provider (DigitalOcean, Vultr, etc.)
    :param str region_name: The name of the Region.
    :param str slug: The Specific Subregion slug.
    """
    __args__ = dict()
    __args__['host'] = host
    __args__['regionName'] = region_name
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('bitlaunch:index/getRegion:getRegion', __args__, opts=opts, typ=GetRegionResult).value

    return AwaitableGetRegionResult(
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        iso=pulumi.get(__ret__, 'iso'),
        region_name=pulumi.get(__ret__, 'region_name'),
        slug=pulumi.get(__ret__, 'slug'),
        unavailable_sizes=pulumi.get(__ret__, 'unavailable_sizes'))


@_utilities.lift_output_func(get_region)
def get_region_output(host: Optional[pulumi.Input[str]] = None,
                      region_name: Optional[pulumi.Input[Optional[str]]] = None,
                      slug: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionResult]:
    """
    Holds available region configurations for a server. Matches https://developers.bitlaunch.io/reference/host-region-object

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_bitlaunch as bitlaunch

    config = pulumi.Config()
    token = config.require_object("token")
    example = bitlaunch.get_region(host="DigitalOcean",
        region_name="New York",
        slug="nyc1")
    ```
    <!--End PulumiCodeChooser -->


    :param str host: Host Provider (DigitalOcean, Vultr, etc.)
    :param str region_name: The name of the Region.
    :param str slug: The Specific Subregion slug.
    """
    ...
