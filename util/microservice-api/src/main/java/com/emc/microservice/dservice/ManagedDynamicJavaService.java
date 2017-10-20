/*
 * Copyright (c) 2015-2016 EMC Corporation All Rights Reserved
 */
package com.emc.microservice.dservice;

import com.emc.microservice.resource.ManagedResource;

/**
 * Created by liebea on 3/22/15.
 * Drink responsibly
 */
public interface ManagedDynamicJavaService extends
        ManagedResource<DynamicJavaServiceDescriptor, DynamicJavaServiceConfiguration> {

    <T> T getInstance();
}
