// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetAuthorizeURL generates an URL for the get authorize operation
type GetAuthorizeURL struct {
	AccessType          *string
	ClientID            string
	CodeChallenge       *string
	CodeChallengeMethod *string
	RedirectURI         string
	ResponseType        string
	Scope               string
	State               *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAuthorizeURL) WithBasePath(bp string) *GetAuthorizeURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAuthorizeURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAuthorizeURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/authorize"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var accessTypeQ string
	if o.AccessType != nil {
		accessTypeQ = *o.AccessType
	}
	if accessTypeQ != "" {
		qs.Set("access_type", accessTypeQ)
	}

	clientIDQ := o.ClientID
	if clientIDQ != "" {
		qs.Set("client_id", clientIDQ)
	}

	var codeChallengeQ string
	if o.CodeChallenge != nil {
		codeChallengeQ = *o.CodeChallenge
	}
	if codeChallengeQ != "" {
		qs.Set("code_challenge", codeChallengeQ)
	}

	var codeChallengeMethodQ string
	if o.CodeChallengeMethod != nil {
		codeChallengeMethodQ = *o.CodeChallengeMethod
	}
	if codeChallengeMethodQ != "" {
		qs.Set("code_challenge_method", codeChallengeMethodQ)
	}

	redirectURIQ := o.RedirectURI
	if redirectURIQ != "" {
		qs.Set("redirect_uri", redirectURIQ)
	}

	responseTypeQ := o.ResponseType
	if responseTypeQ != "" {
		qs.Set("response_type", responseTypeQ)
	}

	scopeQ := o.Scope
	if scopeQ != "" {
		qs.Set("scope", scopeQ)
	}

	var stateQ string
	if o.State != nil {
		stateQ = *o.State
	}
	if stateQ != "" {
		qs.Set("state", stateQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAuthorizeURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAuthorizeURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAuthorizeURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAuthorizeURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAuthorizeURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetAuthorizeURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
