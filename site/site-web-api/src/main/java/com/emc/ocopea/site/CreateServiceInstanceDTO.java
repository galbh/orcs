package com.emc.ocopea.site;

/**
 * Created by liebea on 7/21/15.
 * Drink responsibly
 */
public class CreateServiceInstanceDTO {
    private final String name;
    private final String dsb;
    private final String basedOnCopyId;

    private CreateServiceInstanceDTO() {
        this(null, null, null, null);
    }

    public CreateServiceInstanceDTO(String name, String dsb, String type, String basedOnCopyId) {
        this.name = name;
        this.dsb = dsb;
        this.basedOnCopyId = basedOnCopyId;
    }

    public String getDsb() {
        return dsb;
    }

    public String getName() {
        return name;
    }

    public String getBasedOnCopyId() {
        return basedOnCopyId;
    }
}
