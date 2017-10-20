package com.emc.ocopea.site.app;

/**
 * Created by liebea on 10/18/16.
 * Drink responsibly
 */
public enum DeployedDataServiceState {
    pending,
    queued,
    creating,
    created,
    errorcreating,
    binding,
    errorbinding,
    bound,
    errorunbinding,
    unbinding,
    unbound,
    errorremoving,
    removing,
    removed

}
