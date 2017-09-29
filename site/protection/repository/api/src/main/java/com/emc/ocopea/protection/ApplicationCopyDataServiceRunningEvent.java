package com.emc.ocopea.protection;

import java.util.Date;
import java.util.UUID;

/**
 * Created by liebea on 10/26/16.
 * Drink responsibly
 */
public class ApplicationCopyDataServiceRunningEvent extends ApplicationCopyEvent {
    private final String dsbURN;
    private final String serviceId;

    private ApplicationCopyDataServiceRunningEvent() {
        this(null, null, 0, null, null, null, null);
    }

    public ApplicationCopyDataServiceRunningEvent(UUID id, UUID appInstanceId, long version, Date timeStamp,
                                                  String message, String dsbURN, String serviceId) {
        super(id, appInstanceId, version, timeStamp, message);
        this.dsbURN = dsbURN;
        this.serviceId = serviceId;
    }

    public String getDsbURN() {
        return dsbURN;
    }

    public String getServiceId() {
        return serviceId;
    }

    @Override
    public String toString() {
        return "ApplicationCopyDataServiceRunningEvent{" +
                "dsbURN='" + dsbURN + '\'' +
                ", serviceId='" + serviceId + '\'' +
                '}';
    }
}

