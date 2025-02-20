package m3ter

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// NewClientWithServiceUserAuth generates a new client with the default option read
// from the environment (M3TER_API_KEY, M3TER_API_SECRET, M3TER_API_TOKEN). The
// option passed in as arguments are applied after these default arguments, and all
// option will be passed down to the services and requests that this client makes.
//
// Additionally, NewClientWithServiceUserAuth makes an API call to the oauth/token
// endpoint to exchange the provided APIKey and APISecret for an access token.
func NewClientWithServiceUserAuth(ctx context.Context, opts ...option.RequestOption) (*Client, error) {
	var client *Client
	token := new(string)
	expiry := new(time.Time)

	authMiddleware := option.Middleware(func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		//only perform token refresh if we don't already have a valid token
		if expiry.Before(time.Now()) {
			tokenReq := AuthenticationGetBearerTokenParams{
				GrantType: F(AuthenticationGetBearerTokenParamsGrantTypeClientCredentials),
			}

			basicAuthOption := func(r *requestconfig.RequestConfig) error {
				r.Request.SetBasicAuth(r.APIKey, r.APISecret)
				return nil
			}

			reqTime := time.Now()

			tokenRes, tokenErr := (client).Authentication.GetBearerToken(ctx, tokenReq, basicAuthOption)
			if tokenErr != nil {
				return nil, fmt.Errorf("Failed to retrieve a token using the specified API key and API secret: %w", tokenErr)
			}

			*expiry = reqTime.Add(time.Duration(tokenRes.ExpiresIn) * time.Second)
			*token = tokenRes.AccessToken
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
		return next(req)
	})

	optsWithAuth := append(opts, option.WithMiddleware(authMiddleware))
	client = NewClient(optsWithAuth...)
	//don't use the token refresh middleware on the token refresh endpoint
	client.Authentication.Options = opts

	return client, nil
}
