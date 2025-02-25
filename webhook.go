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
	// This schema defines the credentials required for m3ter request signing.
	Credentials WebhookCredentials `json:"credentials"`
	Description string             `json:"description"`
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

// This schema defines the credentials required for m3ter request signing.
type WebhookCredentials struct {
	JSON webhookCredentialsJSON `json:"-"`
}

// webhookCredentialsJSON contains the JSON metadata for the struct
// [WebhookCredentials]
type webhookCredentialsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookCredentialsJSON) RawJSON() string {
	return r.raw
}

type WebhookNewResponse struct {
	// The credentials required for the webhook.
	Credentials WebhookNewResponseCredentials `json:"credentials,required"`
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

// The credentials required for the webhook.
type WebhookNewResponseCredentials struct {
	JSON webhookNewResponseCredentialsJSON `json:"-"`
}

// webhookNewResponseCredentialsJSON contains the JSON metadata for the struct
// [WebhookNewResponseCredentials]
type webhookNewResponseCredentialsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookNewResponseCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookNewResponseCredentialsJSON) RawJSON() string {
	return r.raw
}

type WebhookUpdateResponse struct {
	// The credentials required for the webhook.
	Credentials WebhookUpdateResponseCredentials `json:"credentials,required"`
	Description string                           `json:"description,required"`
	Name        string                           `json:"name,required"`
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

// The credentials required for the webhook.
type WebhookUpdateResponseCredentials struct {
	JSON webhookUpdateResponseCredentialsJSON `json:"-"`
}

// webhookUpdateResponseCredentialsJSON contains the JSON metadata for the struct
// [WebhookUpdateResponseCredentials]
type webhookUpdateResponseCredentialsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponseCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseCredentialsJSON) RawJSON() string {
	return r.raw
}

type WebhookSetActiveResponse struct {
	// The credentials required for the webhook.
	Credentials WebhookSetActiveResponseCredentials `json:"credentials,required"`
	Description string                              `json:"description,required"`
	Name        string                              `json:"name,required"`
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

// The credentials required for the webhook.
type WebhookSetActiveResponseCredentials struct {
	JSON webhookSetActiveResponseCredentialsJSON `json:"-"`
}

// webhookSetActiveResponseCredentialsJSON contains the JSON metadata for the
// struct [WebhookSetActiveResponseCredentials]
type webhookSetActiveResponseCredentialsJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookSetActiveResponseCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookSetActiveResponseCredentialsJSON) RawJSON() string {
	return r.raw
}

type WebhookNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The credentials required for the webhook.
	Credentials param.Field[WebhookNewParamsCredentials] `json:"credentials,required"`
	Description param.Field[string]                      `json:"description,required"`
	Name        param.Field[string]                      `json:"name,required"`
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

// The credentials required for the webhook.
type WebhookNewParamsCredentials struct {
}

func (r WebhookNewParamsCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type WebhookUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The credentials required for the webhook.
	Credentials param.Field[WebhookUpdateParamsCredentials] `json:"credentials,required"`
	Description param.Field[string]                         `json:"description,required"`
	Name        param.Field[string]                         `json:"name,required"`
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

// The credentials required for the webhook.
type WebhookUpdateParamsCredentials struct {
}

func (r WebhookUpdateParamsCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookListParams struct {
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
	OrgID param.Field[string] `path:"orgId,required"`
}

type WebhookSetActiveParams struct {
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
