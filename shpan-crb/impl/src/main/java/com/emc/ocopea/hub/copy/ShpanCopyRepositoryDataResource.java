package com.emc.ocopea.hub.copy;

import com.emc.microservice.Context;
import com.emc.microservice.MicroServiceApplication;
import com.emc.microservice.blobstore.BlobStoreAPI;
import com.emc.microservice.serialization.SerializationManager;
import com.emc.ocopea.crb.CopyMetaData;
import com.emc.ocopea.crb.CrbWebDataApi;
import com.emc.ocopea.devtools.checkstyle.NoJavadoc;
import com.emc.ocopea.util.MapBuilder;
import org.apache.commons.io.input.CountingInputStream;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.ws.rs.ClientErrorException;
import javax.ws.rs.InternalServerErrorException;
import javax.ws.rs.NotFoundException;
import javax.ws.rs.PathParam;
import javax.ws.rs.core.Application;
import javax.ws.rs.core.Response;
import java.io.IOException;
import java.io.InputStream;
import java.io.PipedInputStream;
import java.io.PipedOutputStream;
import java.util.Collections;

import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.DATA_NAMESPACE;
import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.META_NAMESPACE;
import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.REPOSITORY_NAMESPACE;
import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.SHPAN_PROTOCOL;
import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.SHPAN_VERSION;
import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.SIZE_INFO_KEY;
import static com.emc.ocopea.hub.copy.ShpanCopyRepositoryResource.ShpanRepositoryInfo;

public class ShpanCopyRepositoryDataResource implements CrbWebDataApi {
    private static final Logger log = LoggerFactory.getLogger(ShpanCopyRepositoryDataResource.class);

    private BlobStoreAPI copyStore;

    @NoJavadoc
    @javax.ws.rs.core.Context
    public void setApplication(Application application) {
        Context context = ((MicroServiceApplication) application).getMicroServiceContext();
        copyStore = context.getBlobStoreManager().getManagedResourceByName("copy-store").getBlobStoreAPI();
        SerializationManager serializationManager = context.getSerializationManager();
        if (!serializationManager.getSupportedWriters().contains(ShpanRepositoryInfo.class)) {
            serializationManager.registerJackson(ShpanRepositoryInfo.class);
        }
        if (!serializationManager.getSupportedWriters().contains(CopyMetaData.class)) {
            serializationManager.registerJackson(CopyMetaData.class);
        }
    }

    @Override
    public String createCopyInRepo(
            @PathParam("repoId") String repoId, @PathParam("copyId") String copyId, InputStream copyData) {
        if (copyStore.isExists(REPOSITORY_NAMESPACE, repoId)) {
            if (!copyStore.isExists(DATA_NAMESPACE, copyId)) {
                ShpanRepositoryInfo repositoryInfo =
                        copyStore.readBlob(REPOSITORY_NAMESPACE, repoId, ShpanRepositoryInfo.class);
                repositoryInfo.copies.add(copyId);
                CountingInputStream cis = new CountingInputStream(copyData);
                copyStore.create(DATA_NAMESPACE, copyId, Collections.emptyMap(), cis);
                repositoryInfo.stats.setNumberOfCopiesStored(repositoryInfo.stats.getNumberOfCopiesStored() + 1);
                repositoryInfo.addByteCount(cis.getByteCount());

                final CopyMetaData copyMetaData = copyStore.readBlob(META_NAMESPACE, copyId, CopyMetaData.class);
                copyStore.update(
                        META_NAMESPACE,
                        copyId,
                        Collections.emptyMap(),
                        CopyMetaData.class,
                        new CopyMetaData(
                                copyMetaData.getCopyId(),
                                copyMetaData.getRepoId(),
                                copyMetaData.getCopyTimeStamp(),
                                SHPAN_PROTOCOL,
                                SHPAN_VERSION,
                                copyMetaData.getFacility(),
                                repositoryInfo.instance.getRepoURL(),
                                MapBuilder.<String, String>newHashMap()
                                        .with(SIZE_INFO_KEY, Long.toString(cis.getByteCount()))
                                        .build()));
                copyStore.update(
                        REPOSITORY_NAMESPACE,
                        repoId,
                        Collections.emptyMap(),
                        ShpanRepositoryInfo.class,
                        repositoryInfo);
                return repositoryInfo.instance.getRepoURL();
            } else {
                throw new ClientErrorException("copy " + copyId + " already exists", Response.Status.CONFLICT);
            }
        } else {
            throw new NotFoundException("repository " + repoId + " does not exist");
        }

    }

    @Override
    public Response retrieveCopy(@PathParam("copyId") String copyId) {
        if (copyStore.isExists(DATA_NAMESPACE, copyId)) {
            try {
                PipedOutputStream output = new PipedOutputStream();
                PipedInputStream in = new PipedInputStream(output);
                new Thread(() -> {
                    copyStore.readBlob(DATA_NAMESPACE, copyId, output);
                    try {
                        output.close();
                    } catch (IOException e) {
                        log.error("error on closing connection to data store on retrieving", e);
                    }
                }).start();
                return Response.ok(in).build();
            } catch (IOException e) {
                throw new InternalServerErrorException(e);
            }
        } else {
            throw new NotFoundException("copy " + copyId + " does not exist");
        }
    }

}
