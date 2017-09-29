// Copyright (c) [2017] Dell Inc. or its subsidiaries. All Rights Reserved.
package com.emc.ocopea.site.app;

/**
 * Created by liebea on 10/18/16.
 * Drink responsibly
 */
public enum DeployedAppServiceState {
    pending,
    queued,
    deploying,
    deployed,
    error,
    stopping,
    errorstopping,
    stopped
}
