// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"net/http"
	"slices"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// Endpoint for retrieving a JSON Web Token (JWT) bearer token for a ServiceUser
// using the Client Credentials Grant flow.
//
// A ServiceUser represents the automated process you want to grant access to your
// Organization - that is, as an API user.
//
// AuthenticationService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthenticationService] method instead.
type AuthenticationService struct {
	Options []option.RequestOption
}

// NewAuthenticationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthenticationService(opts ...option.RequestOption) (r *AuthenticationService) {
	r = &AuthenticationService{}
	r.Options = opts
	return
}

// Get authentication token
func (r *AuthenticationService) GetBearerToken(ctx context.Context, body AuthenticationGetBearerTokenParams, opts ...option.RequestOption) (res *AuthenticationGetBearerTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "oauth/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthenticationGetBearerTokenResponse struct {
	// The access token.
	AccessToken string `json:"access_token" api:"required"`
	// Token expiry time in seconds.
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// Not used.
	Scope string `json:"scope"`
	// The token type, which in this case is "bearer".
	TokenType string                                   `json:"token_type"`
	JSON      authenticationGetBearerTokenResponseJSON `json:"-"`
}

// authenticationGetBearerTokenResponseJSON contains the JSON metadata for the
// struct [AuthenticationGetBearerTokenResponse]
type authenticationGetBearerTokenResponseJSON struct {
	AccessToken apijson.Field
	ExpiresIn   apijson.Field
	Scope       apijson.Field
	TokenType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuthenticationGetBearerTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticationGetBearerTokenResponseJSON) RawJSON() string {
	return r.raw
}

type AuthenticationGetBearerTokenParams struct {
	// The grant type.
	GrantType param.Field[AuthenticationGetBearerTokenParamsGrantType] `json:"grant_type" api:"required"`
	// Not used. The JWT scope.
	Scope param.Field[string] `json:"scope"`
}

func (r AuthenticationGetBearerTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The grant type.
type AuthenticationGetBearerTokenParamsGrantType string

const (
	AuthenticationGetBearerTokenParamsGrantTypeClientCredentials AuthenticationGetBearerTokenParamsGrantType = "client_credentials"
)

func (r AuthenticationGetBearerTokenParamsGrantType) IsKnown() bool {
	switch r {
	case AuthenticationGetBearerTokenParamsGrantTypeClientCredentials:
		return true
	}
	return false
}
