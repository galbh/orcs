package com.emc.ocopea.site.app;

import com.emc.ocopea.protection.ProtectionWebAPI;

/**
 * Created by liebea on 10/19/16.
 * Drink responsibly
 */
public interface PolicyEngineConnector {
    ProtectionWebAPI get(String policyType);
}
