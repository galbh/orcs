// Copyright (c) [2017] Dell Inc. or its subsidiaries. All Rights Reserved.
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"../../work/nazgul/copyrepo/copyrepo-web-api/src/main/resources/restapi/operations/crb_web"
	"../../work/nazgul/copyrepo/copyrepo-web-api/src/main/resources/restapi/operations/crb_web_data"
)

// NewCrbAPI creates a new Crb instance
func NewCrbAPI(spec *loads.Document) *CrbAPI {
	return &CrbAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		BinConsumer:     runtime.ByteStreamConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		BinProducer:     runtime.ByteStreamProducer(),
		CrbWebDataCreateCopyInRepoHandler: crb_web_data.CreateCopyInRepoHandlerFunc(func(params crb_web_data.CreateCopyInRepoParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebDataCreateCopyInRepo has not yet been implemented")
		}),
		CrbWebDeleteCopyHandler: crb_web.DeleteCopyHandlerFunc(func(params crb_web.DeleteCopyParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebDeleteCopy has not yet been implemented")
		}),
		CrbWebDeleteRepositoryHandler: crb_web.DeleteRepositoryHandlerFunc(func(params crb_web.DeleteRepositoryParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebDeleteRepository has not yet been implemented")
		}),
		CrbWebGetCopyInstancesHandler: crb_web.GetCopyInstancesHandlerFunc(func(params crb_web.GetCopyInstancesParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebGetCopyInstances has not yet been implemented")
		}),
		CrbWebGetCopyMetaDataHandler: crb_web.GetCopyMetaDataHandlerFunc(func(params crb_web.GetCopyMetaDataParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebGetCopyMetaData has not yet been implemented")
		}),
		CrbWebGetInfoHandler: crb_web.GetInfoHandlerFunc(func(params crb_web.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebGetInfo has not yet been implemented")
		}),
		CrbWebGetRepositoryInfoHandler: crb_web.GetRepositoryInfoHandlerFunc(func(params crb_web.GetRepositoryInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebGetRepositoryInfo has not yet been implemented")
		}),
		CrbWebGetRepositoryStatsHandler: crb_web.GetRepositoryStatsHandlerFunc(func(params crb_web.GetRepositoryStatsParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebGetRepositoryStats has not yet been implemented")
		}),
		CrbWebListRepositoryInstancesHandler: crb_web.ListRepositoryInstancesHandlerFunc(func(params crb_web.ListRepositoryInstancesParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebListRepositoryInstances has not yet been implemented")
		}),
		CrbWebDataRetrieveCopyHandler: crb_web_data.RetrieveCopyHandlerFunc(func(params crb_web_data.RetrieveCopyParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebDataRetrieveCopy has not yet been implemented")
		}),
		CrbWebStoreCopyMetaDataHandler: crb_web.StoreCopyMetaDataHandlerFunc(func(params crb_web.StoreCopyMetaDataParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebStoreCopyMetaData has not yet been implemented")
		}),
		CrbWebStoreRepositoryInfoHandler: crb_web.StoreRepositoryInfoHandlerFunc(func(params crb_web.StoreRepositoryInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebStoreRepositoryInfo has not yet been implemented")
		}),
		CrbWebUpdateRepositoryInfoHandler: crb_web.UpdateRepositoryInfoHandlerFunc(func(params crb_web.UpdateRepositoryInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation CrbWebUpdateRepositoryInfo has not yet been implemented")
		}),
	}
}

/*CrbAPI Copy Service Broker provides a way to copy(store) data-stream associated with a copyId and retrieve data-stream for a given copyId */
type CrbAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// BinConsumer registers a consumer for a "application/octet-stream" mime type
	BinConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// BinProducer registers a producer for a "application/octet-stream" mime type
	BinProducer runtime.Producer

	// CrbWebDataCreateCopyInRepoHandler sets the operation handler for the create copy in repo operation
	CrbWebDataCreateCopyInRepoHandler crb_web_data.CreateCopyInRepoHandler
	// CrbWebDeleteCopyHandler sets the operation handler for the delete copy operation
	CrbWebDeleteCopyHandler crb_web.DeleteCopyHandler
	// CrbWebDeleteRepositoryHandler sets the operation handler for the delete repository operation
	CrbWebDeleteRepositoryHandler crb_web.DeleteRepositoryHandler
	// CrbWebGetCopyInstancesHandler sets the operation handler for the get copy instances operation
	CrbWebGetCopyInstancesHandler crb_web.GetCopyInstancesHandler
	// CrbWebGetCopyMetaDataHandler sets the operation handler for the get copy meta data operation
	CrbWebGetCopyMetaDataHandler crb_web.GetCopyMetaDataHandler
	// CrbWebGetInfoHandler sets the operation handler for the get info operation
	CrbWebGetInfoHandler crb_web.GetInfoHandler
	// CrbWebGetRepositoryInfoHandler sets the operation handler for the get repository info operation
	CrbWebGetRepositoryInfoHandler crb_web.GetRepositoryInfoHandler
	// CrbWebGetRepositoryStatsHandler sets the operation handler for the get repository stats operation
	CrbWebGetRepositoryStatsHandler crb_web.GetRepositoryStatsHandler
	// CrbWebListRepositoryInstancesHandler sets the operation handler for the list repository instances operation
	CrbWebListRepositoryInstancesHandler crb_web.ListRepositoryInstancesHandler
	// CrbWebDataRetrieveCopyHandler sets the operation handler for the retrieve copy operation
	CrbWebDataRetrieveCopyHandler crb_web_data.RetrieveCopyHandler
	// CrbWebStoreCopyMetaDataHandler sets the operation handler for the store copy meta data operation
	CrbWebStoreCopyMetaDataHandler crb_web.StoreCopyMetaDataHandler
	// CrbWebStoreRepositoryInfoHandler sets the operation handler for the store repository info operation
	CrbWebStoreRepositoryInfoHandler crb_web.StoreRepositoryInfoHandler
	// CrbWebUpdateRepositoryInfoHandler sets the operation handler for the update repository info operation
	CrbWebUpdateRepositoryInfoHandler crb_web.UpdateRepositoryInfoHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *CrbAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CrbAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CrbAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CrbAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CrbAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CrbAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CrbAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CrbAPI
func (o *CrbAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinConsumer == nil {
		unregistered = append(unregistered, "BinConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}

	if o.CrbWebDataCreateCopyInRepoHandler == nil {
		unregistered = append(unregistered, "crb_web_data.CreateCopyInRepoHandler")
	}

	if o.CrbWebDeleteCopyHandler == nil {
		unregistered = append(unregistered, "crb_web.DeleteCopyHandler")
	}

	if o.CrbWebDeleteRepositoryHandler == nil {
		unregistered = append(unregistered, "crb_web.DeleteRepositoryHandler")
	}

	if o.CrbWebGetCopyInstancesHandler == nil {
		unregistered = append(unregistered, "crb_web.GetCopyInstancesHandler")
	}

	if o.CrbWebGetCopyMetaDataHandler == nil {
		unregistered = append(unregistered, "crb_web.GetCopyMetaDataHandler")
	}

	if o.CrbWebGetInfoHandler == nil {
		unregistered = append(unregistered, "crb_web.GetInfoHandler")
	}

	if o.CrbWebGetRepositoryInfoHandler == nil {
		unregistered = append(unregistered, "crb_web.GetRepositoryInfoHandler")
	}

	if o.CrbWebGetRepositoryStatsHandler == nil {
		unregistered = append(unregistered, "crb_web.GetRepositoryStatsHandler")
	}

	if o.CrbWebListRepositoryInstancesHandler == nil {
		unregistered = append(unregistered, "crb_web.ListRepositoryInstancesHandler")
	}

	if o.CrbWebDataRetrieveCopyHandler == nil {
		unregistered = append(unregistered, "crb_web_data.RetrieveCopyHandler")
	}

	if o.CrbWebStoreCopyMetaDataHandler == nil {
		unregistered = append(unregistered, "crb_web.StoreCopyMetaDataHandler")
	}

	if o.CrbWebStoreRepositoryInfoHandler == nil {
		unregistered = append(unregistered, "crb_web.StoreRepositoryInfoHandler")
	}

	if o.CrbWebUpdateRepositoryInfoHandler == nil {
		unregistered = append(unregistered, "crb_web.UpdateRepositoryInfoHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CrbAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CrbAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *CrbAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "application/octet-stream":
			result["application/octet-stream"] = o.BinConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *CrbAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CrbAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the crb API
func (o *CrbAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CrbAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repoId}/copies/{copyId}/data"] = crb_web_data.NewCreateCopyInRepo(o.context, o.CrbWebDataCreateCopyInRepoHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/copies/{copyId}"] = crb_web.NewDeleteCopy(o.context, o.CrbWebDeleteCopyHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/repositories/{repoId}"] = crb_web.NewDeleteRepository(o.context, o.CrbWebDeleteRepositoryHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/copies"] = crb_web.NewGetCopyInstances(o.context, o.CrbWebGetCopyInstancesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/copies/{copyId}"] = crb_web.NewGetCopyMetaData(o.context, o.CrbWebGetCopyMetaDataHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/info"] = crb_web.NewGetInfo(o.context, o.CrbWebGetInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repoId}"] = crb_web.NewGetRepositoryInfo(o.context, o.CrbWebGetRepositoryInfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories/{repoId}/stats"] = crb_web.NewGetRepositoryStats(o.context, o.CrbWebGetRepositoryStatsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories"] = crb_web.NewListRepositoryInstances(o.context, o.CrbWebListRepositoryInstancesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/copies/{copyId}/data"] = crb_web_data.NewRetrieveCopy(o.context, o.CrbWebDataRetrieveCopyHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/copies/{copyId}"] = crb_web.NewStoreCopyMetaData(o.context, o.CrbWebStoreCopyMetaDataHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/repositories/{repoId}"] = crb_web.NewStoreRepositoryInfo(o.context, o.CrbWebStoreRepositoryInfoHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/repositories/{repoId}"] = crb_web.NewUpdateRepositoryInfo(o.context, o.CrbWebUpdateRepositoryInfoHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CrbAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *CrbAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
