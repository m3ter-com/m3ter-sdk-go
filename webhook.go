// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
)

// WebhookService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// This endpoint creates a new webhook destination. A webhook destination is a URL
// where webhook payloads will be sent.
func (r *WebhookService) New(ctx context.Context, params WebhookNewParams, opts ...option.RequestOption) (res *WebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationdestinations/webhooks", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the webhook Destination for the UUID.
func (r *WebhookService) Get(ctx context.Context, id string, query WebhookGetParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.OrgID, precfg.OrgID)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationdestinations/webhooks/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a destination to be used for a webhook.
func (r *WebhookService) Update(ctx context.Context, id string, params WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationdestinations/webhooks/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of all Destinations created in the Organization.
func (r *WebhookService) List(ctx context.Context, params WebhookListParams, opts ...option.RequestOption) (res *pagination.Cursor[Webhook], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationdestinations/webhooks", params.OrgID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of all Destinations created in the Organization.
func (r *WebhookService) ListAutoPaging(ctx context.Context, params WebhookListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Webhook] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// This endpoint deletes a specific webhook destination identified by its UUID.
func (r *WebhookService) Delete(ctx context.Context, id string, body WebhookDeleteParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.OrgID, precfg.OrgID)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationdestinations/webhooks/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Set the `active` status on a webhook integration destination.
//
// Use this endpoint to activate or deactivate a webhook integration destination.
// It toggles the `active` status of the specific wehbook destination with the
// given ID.
func (r *WebhookService) SetActive(ctx context.Context, id string, params WebhookSetActiveParams, opts ...option.RequestOption) (res *WebhookSetActiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationdestinations/webhooks/%s/active", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type M3terSignedCredentialsRequest struct {
	// The API key provided by m3ter. This key is part of the credential set required
	// for signing requests and authenticating with m3ter services.
	APIKey string `json:"apiKey,required"`
	// The secret associated with the API key. This secret is used in conjunction with
	// the API key to generate a signature for secure authentication.
	Secret string `json:"secret,required"`
	// Specifies the authorization type. For this schema, it is exclusively set to
	// M3TER_SIGNED_REQUEST.
	Type M3terSignedCredentialsRequestType `json:"type,required"`
	// A flag to indicate whether the credentials are empty.
	//
	// - TRUE - empty credentials.
	// - FALSE - credential details required.
	Empty bool `json:"empty"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version int64                             `json:"version"`
	JSON    m3terSignedCredentialsRequestJSON `json:"-"`
}

// m3terSignedCredentialsRequestJSON contains the JSON metadata for the struct
// [M3terSignedCredentialsRequest]
type m3terSignedCredentialsRequestJSON struct {
	APIKey      apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	Empty       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *M3terSignedCredentialsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r m3terSignedCredentialsRequestJSON) RawJSON() string {
	return r.raw
}

// Specifies the authorization type. For this schema, it is exclusively set to
// M3TER_SIGNED_REQUEST.
type M3terSignedCredentialsRequestType string

const (
	M3terSignedCredentialsRequestTypeM3TerSignedRequest M3terSignedCredentialsRequestType = "M3TER_SIGNED_REQUEST"
)

func (r M3terSignedCredentialsRequestType) IsKnown() bool {
	switch r {
	case M3terSignedCredentialsRequestTypeM3TerSignedRequest:
		return true
	}
	return false
}

type M3terSignedCredentialsRequestParam struct {
	// The API key provided by m3ter. This key is part of the credential set required
	// for signing requests and authenticating with m3ter services.
	APIKey param.Field[string] `json:"apiKey,required"`
	// The secret associated with the API key. This secret is used in conjunction with
	// the API key to generate a signature for secure authentication.
	Secret param.Field[string] `json:"secret,required"`
	// Specifies the authorization type. For this schema, it is exclusively set to
	// M3TER_SIGNED_REQUEST.
	Type param.Field[M3terSignedCredentialsRequestType] `json:"type,required"`
	// A flag to indicate whether the credentials are empty.
	//
	// - TRUE - empty credentials.
	// - FALSE - credential details required.
	Empty param.Field[bool] `json:"empty"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r M3terSignedCredentialsRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type M3terSignedCredentialsResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// the system the integration is for
	Destination string `json:"destination,required"`
	// the type of credentials
	Type string `json:"type,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The API key provided by m3ter. This key is part of the credential set required
	// for signing requests and authenticating with m3ter services.
	APIKey string `json:"apiKey"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// the destinationId the integration is for
	DestinationID string `json:"destinationId"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// the name of the credentials
	Name string `json:"name"`
	// The secret associated with the API key. This secret is used in conjunction with
	// the API key to generate a signature for secure authentication.
	Secret string                             `json:"secret"`
	JSON   m3terSignedCredentialsResponseJSON `json:"-"`
}

// m3terSignedCredentialsResponseJSON contains the JSON metadata for the struct
// [M3terSignedCredentialsResponse]
type m3terSignedCredentialsResponseJSON struct {
	ID             apijson.Field
	Destination    apijson.Field
	Type           apijson.Field
	Version        apijson.Field
	APIKey         apijson.Field
	CreatedBy      apijson.Field
	DestinationID  apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	Secret         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *M3terSignedCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r m3terSignedCredentialsResponseJSON) RawJSON() string {
	return r.raw
}

type Webhook struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64  `json:"version,required"`
	Active  bool   `json:"active"`
	Code    string `json:"code"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// Response representing a set of credentials used for signing m3ter requests.
	Credentials M3terSignedCredentialsResponse `json:"credentials"`
	Description string                         `json:"description"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	Name           string `json:"name"`
	// The URL to which webhook requests are sent.
	URL  string      `json:"url"`
	JSON webhookJSON `json:"-"`
}

// webhookJSON contains the JSON metadata for the struct [Webhook]
type webhookJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	Active         apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	Credentials    apijson.Field
	Description    apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Webhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookJSON) RawJSON() string {
	return r.raw
}

type WebhookNewResponse struct {
	// This schema defines the credentials required for m3ter request signing.
	Credentials M3terSignedCredentialsRequest `json:"credentials,required"`
	Description string                        `json:"description,required"`
	Name        string                        `json:"name,required"`
	// The URL to which the webhook requests will be sent.
	URL    string `json:"url,required"`
	Active bool   `json:"active"`
	Code   string `json:"code"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version int64                  `json:"version"`
	JSON    webhookNewResponseJSON `json:"-"`
}

// webhookNewResponseJSON contains the JSON metadata for the struct
// [WebhookNewResponse]
type webhookNewResponseJSON struct {
	Credentials apijson.Field
	Description apijson.Field
	Name        apijson.Field
	URL         apijson.Field
	Active      apijson.Field
	Code        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookUpdateResponse struct {
	// This schema defines the credentials required for m3ter request signing.
	Credentials M3terSignedCredentialsRequest `json:"credentials,required"`
	Description string                        `json:"description,required"`
	Name        string                        `json:"name,required"`
	// The URL to which the webhook requests will be sent.
	URL    string `json:"url,required"`
	Active bool   `json:"active"`
	Code   string `json:"code"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version int64                     `json:"version"`
	JSON    webhookUpdateResponseJSON `json:"-"`
}

// webhookUpdateResponseJSON contains the JSON metadata for the struct
// [WebhookUpdateResponse]
type webhookUpdateResponseJSON struct {
	Credentials apijson.Field
	Description apijson.Field
	Name        apijson.Field
	URL         apijson.Field
	Active      apijson.Field
	Code        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookSetActiveResponse struct {
	// This schema defines the credentials required for m3ter request signing.
	Credentials M3terSignedCredentialsRequest `json:"credentials,required"`
	Description string                        `json:"description,required"`
	Name        string                        `json:"name,required"`
	// The URL to which the webhook requests will be sent.
	URL    string `json:"url,required"`
	Active bool   `json:"active"`
	Code   string `json:"code"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version int64                        `json:"version"`
	JSON    webhookSetActiveResponseJSON `json:"-"`
}

// webhookSetActiveResponseJSON contains the JSON metadata for the struct
// [WebhookSetActiveResponse]
type webhookSetActiveResponseJSON struct {
	Credentials apijson.Field
	Description apijson.Field
	Name        apijson.Field
	URL         apijson.Field
	Active      apijson.Field
	Code        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookSetActiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookSetActiveResponseJSON) RawJSON() string {
	return r.raw
}

type WebhookNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// This schema defines the credentials required for m3ter request signing.
	Credentials param.Field[M3terSignedCredentialsRequestParam] `json:"credentials,required"`
	Description param.Field[string]                             `json:"description,required"`
	Name        param.Field[string]                             `json:"name,required"`
	// The URL to which the webhook requests will be sent.
	URL    param.Field[string] `json:"url,required"`
	Active param.Field[bool]   `json:"active"`
	Code   param.Field[string] `json:"code"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type WebhookUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// This schema defines the credentials required for m3ter request signing.
	Credentials param.Field[M3terSignedCredentialsRequestParam] `json:"credentials,required"`
	Description param.Field[string]                             `json:"description,required"`
	Name        param.Field[string]                             `json:"name,required"`
	// The URL to which the webhook requests will be sent.
	URL    param.Field[string] `json:"url,required"`
	Active param.Field[bool]   `json:"active"`
	Code   param.Field[string] `json:"code"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string]   `path:"orgId,required"`
	IDs   param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of WebhookIntegrations to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WebhookDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type WebhookSetActiveParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// active status of the webhook
	Active param.Field[bool] `query:"active"`
}

// URLQuery serializes [WebhookSetActiveParams]'s query parameters as `url.Values`.
func (r WebhookSetActiveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
