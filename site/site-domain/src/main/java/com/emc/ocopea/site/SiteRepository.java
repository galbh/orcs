package com.emc.ocopea.site;

/**
 * Created by liebea on 10/31/16.
 * Drink responsibly
 */
public interface SiteRepository {
    Site load();

    void persist(Site site);
}
