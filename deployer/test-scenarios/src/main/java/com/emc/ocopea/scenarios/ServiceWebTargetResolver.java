package com.emc.ocopea.scenarios;

import javax.ws.rs.client.WebTarget;

/**
 * Created by liebea on 6/20/16.
 * Drink responsibly
 */
public interface ServiceWebTargetResolver {
    public WebTarget resolveWebTarget(String serviceURN, String path);
}
