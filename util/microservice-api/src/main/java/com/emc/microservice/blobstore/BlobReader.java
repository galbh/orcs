/*
 * Copyright (c) 2014-2016 EMC Corporation All Rights Reserved
 */
package com.emc.microservice.blobstore;

import java.io.InputStream;

public interface BlobReader {
    void read(InputStream in);
}
