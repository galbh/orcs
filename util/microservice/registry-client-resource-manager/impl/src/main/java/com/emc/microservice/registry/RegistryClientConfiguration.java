package com.emc.microservice.registry;

import com.emc.microservice.resource.ResourceConfiguration;

import java.util.Collections;
import java.util.HashMap;

public class RegistryClientConfiguration extends ResourceConfiguration {

    public RegistryClientConfiguration() {
        super("Registry Client", Collections.emptyList());
        setPropertyValues(new HashMap<>());
    }
}
