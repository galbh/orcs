// Copyright (c) [2017] Dell Inc. or its subsidiaries. All Rights Reserved.
package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// StoreCopyMetaDataHandlerFunc turns a function with the right signature into a store copy meta data handler
type StoreCopyMetaDataHandlerFunc func(StoreCopyMetaDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn StoreCopyMetaDataHandlerFunc) Handle(params StoreCopyMetaDataParams) middleware.Responder {
	return fn(params)
}

// StoreCopyMetaDataHandler interface for that can handle valid store copy meta data params
type StoreCopyMetaDataHandler interface {
	Handle(StoreCopyMetaDataParams) middleware.Responder
}

// NewStoreCopyMetaData creates a new http.Handler for the store copy meta data operation
func NewStoreCopyMetaData(ctx *middleware.Context, handler StoreCopyMetaDataHandler) *StoreCopyMetaData {
	return &StoreCopyMetaData{Context: ctx, Handler: handler}
}

/*StoreCopyMetaData swagger:route POST /copies/{copyId} crb-web storeCopyMetaData

Store a copy instance meta data

Store a copy instance attributes

*/
type StoreCopyMetaData struct {
	Context *middleware.Context
	Handler StoreCopyMetaDataHandler
}

func (o *StoreCopyMetaData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewStoreCopyMetaDataParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
