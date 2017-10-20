package com.emc.microservice.discovery;

import javax.ws.rs.client.WebTarget;

/**
 * Created by liebea on 5/2/16.
 * Drink responsibly
 */
public interface WebAPIConnection {
    <T> T resolve(Class<T> resourceClass);

    WebTarget getWebTarget();
}
