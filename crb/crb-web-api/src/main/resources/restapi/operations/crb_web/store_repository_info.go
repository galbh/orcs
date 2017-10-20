// Copyright (c) [2017] Dell Inc. or its subsidiaries. All Rights Reserved.
package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StoreRepositoryInfoHandlerFunc turns a function with the right signature into a store repository info handler
type StoreRepositoryInfoHandlerFunc func(StoreRepositoryInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StoreRepositoryInfoHandlerFunc) Handle(params StoreRepositoryInfoParams) middleware.Responder {
	return fn(params)
}

// StoreRepositoryInfoHandler interface for that can handle valid store repository info params
type StoreRepositoryInfoHandler interface {
	Handle(StoreRepositoryInfoParams) middleware.Responder
}

// NewStoreRepositoryInfo creates a new http.Handler for the store repository info operation
func NewStoreRepositoryInfo(ctx *middleware.Context, handler StoreRepositoryInfoHandler) *StoreRepositoryInfo {
	return &StoreRepositoryInfo{Context: ctx, Handler: handler}
}

/*StoreRepositoryInfo swagger:route POST /repositories/{repoId} crb-web storeRepositoryInfo

Store repository info of specified repoId.

Store repository configuration info.

*/
type StoreRepositoryInfo struct {
	Context *middleware.Context
	Handler StoreRepositoryInfoHandler
}

func (o *StoreRepositoryInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewStoreRepositoryInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
