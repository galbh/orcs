/*
 * Copyright (c) 2015-2016 EMC Corporation All Rights Reserved
 */
package com.emc.microservice.dservice;

import com.emc.microservice.resource.ResourceConfiguration;

import java.util.ArrayList;

/**
 * Created by liebea on 3/22/15.
 * Drink responsibly
 */
public class DynamicJavaServiceConfiguration extends ResourceConfiguration {
    public DynamicJavaServiceConfiguration() {
        super("DynamicJavaService", new ArrayList<>());
    }
}
