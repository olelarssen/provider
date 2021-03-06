// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetCredentialsParams creates a new GetCredentialsParams object
//
// There are no default values defined in the spec.
func NewGetCredentialsParams() GetCredentialsParams {

	return GetCredentialsParams{}
}

// GetCredentialsParams contains all the bound params for the get credentials operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetCredentials
type GetCredentialsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*client id
	  Required: true
	  In: query
	*/
	ClientID string
	/*domain for app
	  In: query
	*/
	Domain *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetCredentialsParams() beforehand.
func (o *GetCredentialsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClientID, qhkClientID, _ := qs.GetOK("client_id")
	if err := o.bindClientID(qClientID, qhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	qDomain, qhkDomain, _ := qs.GetOK("domain")
	if err := o.bindDomain(qDomain, qhkDomain, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClientID binds and validates parameter ClientID from query.
func (o *GetCredentialsParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("client_id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("client_id", "query", raw); err != nil {
		return err
	}
	o.ClientID = raw

	return nil
}

// bindDomain binds and validates parameter Domain from query.
func (o *GetCredentialsParams) bindDomain(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Domain = &raw

	return nil
}
