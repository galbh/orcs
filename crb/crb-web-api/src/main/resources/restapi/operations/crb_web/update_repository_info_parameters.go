package crb_web

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"../../work/nazgul/copyrepo/copyrepo-web-api/src/main/resources/models"
)

// NewUpdateRepositoryInfoParams creates a new UpdateRepositoryInfoParams object
// with the default values initialized.
func NewUpdateRepositoryInfoParams() UpdateRepositoryInfoParams {
	var ()
	return UpdateRepositoryInfoParams{}
}

// UpdateRepositoryInfoParams contains all the bound params for the update repository info operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateRepositoryInfo
type UpdateRepositoryInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The id of the copy repository instance
	  Required: true
	  In: path
	*/
	RepoID string
	/*repository access info
	  Required: true
	  In: body
	*/
	RepoInfo *models.RepositoryInfo
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateRepositoryInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rRepoID, rhkRepoID, _ := route.Params.GetOK("repoId")
	if err := o.bindRepoID(rRepoID, rhkRepoID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.RepositoryInfo
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("repoInfo", "body"))
			} else {
				res = append(res, errors.NewParseError("repoInfo", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.RepoInfo = &body
			}
		}

	} else {
		res = append(res, errors.Required("repoInfo", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRepositoryInfoParams) bindRepoID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.RepoID = raw

	return nil
}
