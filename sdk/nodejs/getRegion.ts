// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Holds available region configurations for a server. Matches https://developers.bitlaunch.io/reference/host-region-object
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as bitlaunch from "@pulumi/bitlaunch";
 *
 * const config = new pulumi.Config();
 * const token = config.requireObject("token");
 * const example = bitlaunch.getRegion({
 *     host: "DigitalOcean",
 *     regionName: "New York",
 *     slug: "nyc1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRegion(args: GetRegionArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("bitlaunch:index/getRegion:getRegion", {
        "host": args.host,
        "regionName": args.regionName,
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionArgs {
    /**
     * Host Provider (DigitalOcean, Vultr, etc.)
     */
    host: string;
    /**
     * The name of the Region.
     */
    regionName?: string;
    /**
     * The Specific Subregion slug.
     */
    slug?: string;
}

/**
 * A collection of values returned by getRegion.
 */
export interface GetRegionResult {
    /**
     * Host Provider (DigitalOcean, Vultr, etc.)
     */
    readonly host: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ISO code for the region.
     */
    readonly iso: string;
    /**
     * The name of the Region.
     */
    readonly regionName?: string;
    /**
     * The Specific Subregion slug.
     */
    readonly slug?: string;
    /**
     * A list of the unavailable sizes for this subregion.
     */
    readonly unavailableSizes: string[];
}
/**
 * Holds available region configurations for a server. Matches https://developers.bitlaunch.io/reference/host-region-object
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as bitlaunch from "@pulumi/bitlaunch";
 *
 * const config = new pulumi.Config();
 * const token = config.requireObject("token");
 * const example = bitlaunch.getRegion({
 *     host: "DigitalOcean",
 *     regionName: "New York",
 *     slug: "nyc1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRegionOutput(args: GetRegionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionResult> {
    return pulumi.output(args).apply((a: any) => getRegion(a, opts))
}

/**
 * A collection of arguments for invoking getRegion.
 */
export interface GetRegionOutputArgs {
    /**
     * Host Provider (DigitalOcean, Vultr, etc.)
     */
    host: pulumi.Input<string>;
    /**
     * The name of the Region.
     */
    regionName?: pulumi.Input<string>;
    /**
     * The Specific Subregion slug.
     */
    slug?: pulumi.Input<string>;
}
