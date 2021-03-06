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

// NewGetAuthorizeParams creates a new GetAuthorizeParams object
//
// There are no default values defined in the spec.
func NewGetAuthorizeParams() GetAuthorizeParams {

	return GetAuthorizeParams{}
}

// GetAuthorizeParams contains all the bound params for the get authorize operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAuthorize
type GetAuthorizeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*access type
	  In: query
	*/
	AccessType *string
	/*client id
	  Required: true
	  In: query
	*/
	ClientID string
	/*code challenge
	  In: query
	*/
	CodeChallenge *string
	/*code challenge method
	  In: query
	*/
	CodeChallengeMethod *string
	/*redirect uri
	  Required: true
	  In: query
	*/
	RedirectURI string
	/*response type
	  Required: true
	  In: query
	*/
	ResponseType string
	/*scope
	  Required: true
	  In: query
	*/
	Scope string
	/*state
	  In: query
	*/
	State *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAuthorizeParams() beforehand.
func (o *GetAuthorizeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAccessType, qhkAccessType, _ := qs.GetOK("access_type")
	if err := o.bindAccessType(qAccessType, qhkAccessType, route.Formats); err != nil {
		res = append(res, err)
	}

	qClientID, qhkClientID, _ := qs.GetOK("client_id")
	if err := o.bindClientID(qClientID, qhkClientID, route.Formats); err != nil {
		res = append(res, err)
	}

	qCodeChallenge, qhkCodeChallenge, _ := qs.GetOK("code_challenge")
	if err := o.bindCodeChallenge(qCodeChallenge, qhkCodeChallenge, route.Formats); err != nil {
		res = append(res, err)
	}

	qCodeChallengeMethod, qhkCodeChallengeMethod, _ := qs.GetOK("code_challenge_method")
	if err := o.bindCodeChallengeMethod(qCodeChallengeMethod, qhkCodeChallengeMethod, route.Formats); err != nil {
		res = append(res, err)
	}

	qRedirectURI, qhkRedirectURI, _ := qs.GetOK("redirect_uri")
	if err := o.bindRedirectURI(qRedirectURI, qhkRedirectURI, route.Formats); err != nil {
		res = append(res, err)
	}

	qResponseType, qhkResponseType, _ := qs.GetOK("response_type")
	if err := o.bindResponseType(qResponseType, qhkResponseType, route.Formats); err != nil {
		res = append(res, err)
	}

	qScope, qhkScope, _ := qs.GetOK("scope")
	if err := o.bindScope(qScope, qhkScope, route.Formats); err != nil {
		res = append(res, err)
	}

	qState, qhkState, _ := qs.GetOK("state")
	if err := o.bindState(qState, qhkState, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAccessType binds and validates parameter AccessType from query.
func (o *GetAuthorizeParams) bindAccessType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.AccessType = &raw

	return nil
}

// bindClientID binds and validates parameter ClientID from query.
func (o *GetAuthorizeParams) bindClientID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindCodeChallenge binds and validates parameter CodeChallenge from query.
func (o *GetAuthorizeParams) bindCodeChallenge(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CodeChallenge = &raw

	return nil
}

// bindCodeChallengeMethod binds and validates parameter CodeChallengeMethod from query.
func (o *GetAuthorizeParams) bindCodeChallengeMethod(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.CodeChallengeMethod = &raw

	return nil
}

// bindRedirectURI binds and validates parameter RedirectURI from query.
func (o *GetAuthorizeParams) bindRedirectURI(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("redirect_uri", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("redirect_uri", "query", raw); err != nil {
		return err
	}
	o.RedirectURI = raw

	return nil
}

// bindResponseType binds and validates parameter ResponseType from query.
func (o *GetAuthorizeParams) bindResponseType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("response_type", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("response_type", "query", raw); err != nil {
		return err
	}
	o.ResponseType = raw

	if err := o.validateResponseType(formats); err != nil {
		return err
	}

	return nil
}

// validateResponseType carries on validations for parameter ResponseType
func (o *GetAuthorizeParams) validateResponseType(formats strfmt.Registry) error {

	if err := validate.EnumCase("response_type", "query", o.ResponseType, []interface{}{"code", "token"}, true); err != nil {
		return err
	}

	return nil
}

// bindScope binds and validates parameter Scope from query.
func (o *GetAuthorizeParams) bindScope(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("scope", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("scope", "query", raw); err != nil {
		return err
	}
	o.Scope = raw

	return nil
}

// bindState binds and validates parameter State from query.
func (o *GetAuthorizeParams) bindState(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.State = &raw

	return nil
}
