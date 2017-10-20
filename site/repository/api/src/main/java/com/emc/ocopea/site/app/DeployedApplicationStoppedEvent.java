package com.emc.ocopea.site.app;

import java.util.Date;
import java.util.UUID;

/**
 * Created by liebea on 10/18/16.
 * Drink responsibly
 */
public class DeployedApplicationStoppedEvent extends DeployedApplicationEvent {
    private DeployedApplicationStoppedEvent() {
        super(null, 0L, null, null);
    }

    protected DeployedApplicationStoppedEvent(UUID id, long version, Date timestamp, String message) {
        super(id, version, timestamp, message);
    }
}

