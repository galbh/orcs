/* $Id$
 * 
 * This computer code is copyright 2014 EMC Corporation
 * All rights reserved
 */

package com.emc.ocopea.services.rest;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path(MicroServiceStateAPI.BASE_URI)
public interface MicroServiceStateAPI {

    public static final String BASE_URI = "/state";

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    ServiceState getServiceState();
}
