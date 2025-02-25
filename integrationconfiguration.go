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

// IntegrationConfigurationService contains methods and other services that help
// with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationConfigurationService] method instead.
type IntegrationConfigurationService struct {
	Options []option.RequestOption
}

// NewIntegrationConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIntegrationConfigurationService(opts ...option.RequestOption) (r *IntegrationConfigurationService) {
	r = &IntegrationConfigurationService{}
	r.Options = opts
	return
}

// Set the integration configuration for the entity.
func (r *IntegrationConfigurationService) New(ctx context.Context, params IntegrationConfigurationNewParams, opts ...option.RequestOption) (res *IntegrationConfigurationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the integration configuration for the given UUID.
//
// This endpoint retrieves the configuration details of a specific integration
// within an organization. It is useful for obtaining the settings and parameters
// of an integration.
func (r *IntegrationConfigurationService) Get(ctx context.Context, id string, query IntegrationConfigurationGetParams, opts ...option.RequestOption) (res *IntegrationConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the integration configuration for the given UUID.
//
// This endpoint allows you to update the configuration of a specific integration
// within your organization. It is used to modify settings or parameters of an
// existing integration.
func (r *IntegrationConfigurationService) Update(ctx context.Context, id string, params IntegrationConfigurationUpdateParams, opts ...option.RequestOption) (res *IntegrationConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List all integration configurations.
//
// This endpoint retrieves a list of all integration configurations for the
// specified Organization. The list can be paginated for easier management.
func (r *IntegrationConfigurationService) List(ctx context.Context, params IntegrationConfigurationListParams, opts ...option.RequestOption) (res *pagination.Cursor[IntegrationConfigurationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs", params.OrgID)
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

// List all integration configurations.
//
// This endpoint retrieves a list of all integration configurations for the
// specified Organization. The list can be paginated for easier management.
func (r *IntegrationConfigurationService) ListAutoPaging(ctx context.Context, params IntegrationConfigurationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[IntegrationConfigurationListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the integration configuration for the given UUID.
//
// Use this endpoint to delete the configuration of a specific integration within
// your organization. It is intended for removing integration settings that are no
// longer needed.
func (r *IntegrationConfigurationService) Delete(ctx context.Context, id string, body IntegrationConfigurationDeleteParams, opts ...option.RequestOption) (res *IntegrationConfigurationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Enables a previously disabled integration configuration, allowing it to be
// operational again.
func (r *IntegrationConfigurationService) Enable(ctx context.Context, id string, body IntegrationConfigurationEnableParams, opts ...option.RequestOption) (res *IntegrationConfigurationEnableResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs/%s/enable", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve the integration configuration for the entity
func (r *IntegrationConfigurationService) GetByEntity(ctx context.Context, entityType string, params IntegrationConfigurationGetByEntityParams, opts ...option.RequestOption) (res *IntegrationConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if entityType == "" {
		err = errors.New("missing required entityType parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/integrationconfigs/entity/%s", params.OrgID, entityType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type IntegrationConfiguration struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The destination system for the integration.
	Destination string `json:"destination,required"`
	// The unique identifier (UUID) of the entity the integration is for.
	EntityID string `json:"entityId,required"`
	// The type of entity the integration is for _(e.g. Bill)_.
	EntityType string                         `json:"entityType,required"`
	Status     IntegrationConfigurationStatus `json:"status,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The date and time the integration was completed _(in ISO-8601 format)_.
	DtCompleted time.Time `json:"dtCompleted" format:"date-time"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The date and time the integration was started _(in ISO-8601 format)_.
	DtStarted time.Time `json:"dtStarted" format:"date-time"`
	// Describes any errors encountered during the integration run.
	Error string `json:"error"`
	// The external ID in the destination system if available.
	ExternalID string `json:"externalId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The URL of the entity in the destination system if available.
	URL  string                       `json:"url"`
	JSON integrationConfigurationJSON `json:"-"`
}

// integrationConfigurationJSON contains the JSON metadata for the struct
// [IntegrationConfiguration]
type integrationConfigurationJSON struct {
	ID             apijson.Field
	Destination    apijson.Field
	EntityID       apijson.Field
	EntityType     apijson.Field
	Status         apijson.Field
	Version        apijson.Field
	CreatedBy      apijson.Field
	DtCompleted    apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	DtStarted      apijson.Field
	Error          apijson.Field
	ExternalID     apijson.Field
	LastModifiedBy apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IntegrationConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigurationJSON) RawJSON() string {
	return r.raw
}

type IntegrationConfigurationStatus string

const (
	IntegrationConfigurationStatusWaiting                IntegrationConfigurationStatus = "WAITING"
	IntegrationConfigurationStatusStarted                IntegrationConfigurationStatus = "STARTED"
	IntegrationConfigurationStatusComplete               IntegrationConfigurationStatus = "COMPLETE"
	IntegrationConfigurationStatusError                  IntegrationConfigurationStatus = "ERROR"
	IntegrationConfigurationStatusAwaitingRetry          IntegrationConfigurationStatus = "AWAITING_RETRY"
	IntegrationConfigurationStatusAuthFailed             IntegrationConfigurationStatus = "AUTH_FAILED"
	IntegrationConfigurationStatusAccountingPeriodClosed IntegrationConfigurationStatus = "ACCOUNTING_PERIOD_CLOSED"
	IntegrationConfigurationStatusInvoiceAlreadyPaid     IntegrationConfigurationStatus = "INVOICE_ALREADY_PAID"
	IntegrationConfigurationStatusDisabled               IntegrationConfigurationStatus = "DISABLED"
)

func (r IntegrationConfigurationStatus) IsKnown() bool {
	switch r {
	case IntegrationConfigurationStatusWaiting, IntegrationConfigurationStatusStarted, IntegrationConfigurationStatusComplete, IntegrationConfigurationStatusError, IntegrationConfigurationStatusAwaitingRetry, IntegrationConfigurationStatusAuthFailed, IntegrationConfigurationStatusAccountingPeriodClosed, IntegrationConfigurationStatusInvoiceAlreadyPaid, IntegrationConfigurationStatusDisabled:
		return true
	}
	return false
}

type IntegrationConfigurationNewResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The type of destination _(e.g. Netsuite, webhooks)_.
	Destination string `json:"destination,required"`
	// The type of entity the integration is for _(e.g. Bill)_.
	EntityType string `json:"entityType,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A flag indicating whether the integration configuration is authorized.
	//
	// - TRUE - authorized.
	// - FALSE - not authorized.
	Authorized bool `json:"authorized"`
	// Configuration data for the integration
	ConfigData map[string]interface{} `json:"configData"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The unique identifier (UUID) of the entity the integration is for.
	DestinationID string `json:"destinationId"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// A flag indicating whether the integration configuration is currently enabled or
	// disabled.
	//
	// - TRUE - enabled.
	// - FALSE - disabled.
	Enabled bool `json:"enabled"`
	// The unique identifier (UUID) of the entity this integration is for _(e.g. the ID
	// of a notification configuration. Optional.)_
	EntityID string `json:"entityId"`
	// UUID of the credentials to use for this integration
	IntegrationCredentialsID string `json:"integrationCredentialsId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the configuration
	Name string `json:"name"`
	// Specifies the type of trigger for the integration.
	//
	// Deprecated: deprecated
	TriggerType IntegrationConfigurationNewResponseTriggerType `json:"triggerType"`
	JSON        integrationConfigurationNewResponseJSON        `json:"-"`
}

// integrationConfigurationNewResponseJSON contains the JSON metadata for the
// struct [IntegrationConfigurationNewResponse]
type integrationConfigurationNewResponseJSON struct {
	ID                       apijson.Field
	Destination              apijson.Field
	EntityType               apijson.Field
	Version                  apijson.Field
	Authorized               apijson.Field
	ConfigData               apijson.Field
	CreatedBy                apijson.Field
	DestinationID            apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	Enabled                  apijson.Field
	EntityID                 apijson.Field
	IntegrationCredentialsID apijson.Field
	LastModifiedBy           apijson.Field
	Name                     apijson.Field
	TriggerType              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *IntegrationConfigurationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigurationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of trigger for the integration.
type IntegrationConfigurationNewResponseTriggerType string

const (
	IntegrationConfigurationNewResponseTriggerTypeEvent    IntegrationConfigurationNewResponseTriggerType = "EVENT"
	IntegrationConfigurationNewResponseTriggerTypeSchedule IntegrationConfigurationNewResponseTriggerType = "SCHEDULE"
)

func (r IntegrationConfigurationNewResponseTriggerType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationNewResponseTriggerTypeEvent, IntegrationConfigurationNewResponseTriggerTypeSchedule:
		return true
	}
	return false
}

type IntegrationConfigurationUpdateResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The type of destination _(e.g. Netsuite, webhooks)_.
	Destination string `json:"destination,required"`
	// The type of entity the integration is for _(e.g. Bill)_.
	EntityType string `json:"entityType,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A flag indicating whether the integration configuration is authorized.
	//
	// - TRUE - authorized.
	// - FALSE - not authorized.
	Authorized bool `json:"authorized"`
	// Configuration data for the integration
	ConfigData map[string]interface{} `json:"configData"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The unique identifier (UUID) of the entity the integration is for.
	DestinationID string `json:"destinationId"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// A flag indicating whether the integration configuration is currently enabled or
	// disabled.
	//
	// - TRUE - enabled.
	// - FALSE - disabled.
	Enabled bool `json:"enabled"`
	// The unique identifier (UUID) of the entity this integration is for _(e.g. the ID
	// of a notification configuration. Optional.)_
	EntityID string `json:"entityId"`
	// UUID of the credentials to use for this integration
	IntegrationCredentialsID string `json:"integrationCredentialsId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the configuration
	Name string `json:"name"`
	// Specifies the type of trigger for the integration.
	//
	// Deprecated: deprecated
	TriggerType IntegrationConfigurationUpdateResponseTriggerType `json:"triggerType"`
	JSON        integrationConfigurationUpdateResponseJSON        `json:"-"`
}

// integrationConfigurationUpdateResponseJSON contains the JSON metadata for the
// struct [IntegrationConfigurationUpdateResponse]
type integrationConfigurationUpdateResponseJSON struct {
	ID                       apijson.Field
	Destination              apijson.Field
	EntityType               apijson.Field
	Version                  apijson.Field
	Authorized               apijson.Field
	ConfigData               apijson.Field
	CreatedBy                apijson.Field
	DestinationID            apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	Enabled                  apijson.Field
	EntityID                 apijson.Field
	IntegrationCredentialsID apijson.Field
	LastModifiedBy           apijson.Field
	Name                     apijson.Field
	TriggerType              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *IntegrationConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of trigger for the integration.
type IntegrationConfigurationUpdateResponseTriggerType string

const (
	IntegrationConfigurationUpdateResponseTriggerTypeEvent    IntegrationConfigurationUpdateResponseTriggerType = "EVENT"
	IntegrationConfigurationUpdateResponseTriggerTypeSchedule IntegrationConfigurationUpdateResponseTriggerType = "SCHEDULE"
)

func (r IntegrationConfigurationUpdateResponseTriggerType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationUpdateResponseTriggerTypeEvent, IntegrationConfigurationUpdateResponseTriggerTypeSchedule:
		return true
	}
	return false
}

type IntegrationConfigurationListResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The type of destination _(e.g. Netsuite, webhooks)_.
	Destination string `json:"destination,required"`
	// The type of entity the integration is for _(e.g. Bill)_.
	EntityType string `json:"entityType,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A flag indicating whether the integration configuration is authorized.
	//
	// - TRUE - authorized.
	// - FALSE - not authorized.
	Authorized bool `json:"authorized"`
	// Configuration data for the integration
	ConfigData map[string]interface{} `json:"configData"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The unique identifier (UUID) of the entity the integration is for.
	DestinationID string `json:"destinationId"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// A flag indicating whether the integration configuration is currently enabled or
	// disabled.
	//
	// - TRUE - enabled.
	// - FALSE - disabled.
	Enabled bool `json:"enabled"`
	// The unique identifier (UUID) of the entity this integration is for _(e.g. the ID
	// of a notification configuration. Optional.)_
	EntityID string `json:"entityId"`
	// UUID of the credentials to use for this integration
	IntegrationCredentialsID string `json:"integrationCredentialsId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the configuration
	Name string `json:"name"`
	// Specifies the type of trigger for the integration.
	//
	// Deprecated: deprecated
	TriggerType IntegrationConfigurationListResponseTriggerType `json:"triggerType"`
	JSON        integrationConfigurationListResponseJSON        `json:"-"`
}

// integrationConfigurationListResponseJSON contains the JSON metadata for the
// struct [IntegrationConfigurationListResponse]
type integrationConfigurationListResponseJSON struct {
	ID                       apijson.Field
	Destination              apijson.Field
	EntityType               apijson.Field
	Version                  apijson.Field
	Authorized               apijson.Field
	ConfigData               apijson.Field
	CreatedBy                apijson.Field
	DestinationID            apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	Enabled                  apijson.Field
	EntityID                 apijson.Field
	IntegrationCredentialsID apijson.Field
	LastModifiedBy           apijson.Field
	Name                     apijson.Field
	TriggerType              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *IntegrationConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigurationListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of trigger for the integration.
type IntegrationConfigurationListResponseTriggerType string

const (
	IntegrationConfigurationListResponseTriggerTypeEvent    IntegrationConfigurationListResponseTriggerType = "EVENT"
	IntegrationConfigurationListResponseTriggerTypeSchedule IntegrationConfigurationListResponseTriggerType = "SCHEDULE"
)

func (r IntegrationConfigurationListResponseTriggerType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationListResponseTriggerTypeEvent, IntegrationConfigurationListResponseTriggerTypeSchedule:
		return true
	}
	return false
}

type IntegrationConfigurationDeleteResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The type of destination _(e.g. Netsuite, webhooks)_.
	Destination string `json:"destination,required"`
	// The type of entity the integration is for _(e.g. Bill)_.
	EntityType string `json:"entityType,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A flag indicating whether the integration configuration is authorized.
	//
	// - TRUE - authorized.
	// - FALSE - not authorized.
	Authorized bool `json:"authorized"`
	// Configuration data for the integration
	ConfigData map[string]interface{} `json:"configData"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The unique identifier (UUID) of the entity the integration is for.
	DestinationID string `json:"destinationId"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// A flag indicating whether the integration configuration is currently enabled or
	// disabled.
	//
	// - TRUE - enabled.
	// - FALSE - disabled.
	Enabled bool `json:"enabled"`
	// The unique identifier (UUID) of the entity this integration is for _(e.g. the ID
	// of a notification configuration. Optional.)_
	EntityID string `json:"entityId"`
	// UUID of the credentials to use for this integration
	IntegrationCredentialsID string `json:"integrationCredentialsId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the configuration
	Name string `json:"name"`
	// Specifies the type of trigger for the integration.
	//
	// Deprecated: deprecated
	TriggerType IntegrationConfigurationDeleteResponseTriggerType `json:"triggerType"`
	JSON        integrationConfigurationDeleteResponseJSON        `json:"-"`
}

// integrationConfigurationDeleteResponseJSON contains the JSON metadata for the
// struct [IntegrationConfigurationDeleteResponse]
type integrationConfigurationDeleteResponseJSON struct {
	ID                       apijson.Field
	Destination              apijson.Field
	EntityType               apijson.Field
	Version                  apijson.Field
	Authorized               apijson.Field
	ConfigData               apijson.Field
	CreatedBy                apijson.Field
	DestinationID            apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	Enabled                  apijson.Field
	EntityID                 apijson.Field
	IntegrationCredentialsID apijson.Field
	LastModifiedBy           apijson.Field
	Name                     apijson.Field
	TriggerType              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *IntegrationConfigurationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigurationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of trigger for the integration.
type IntegrationConfigurationDeleteResponseTriggerType string

const (
	IntegrationConfigurationDeleteResponseTriggerTypeEvent    IntegrationConfigurationDeleteResponseTriggerType = "EVENT"
	IntegrationConfigurationDeleteResponseTriggerTypeSchedule IntegrationConfigurationDeleteResponseTriggerType = "SCHEDULE"
)

func (r IntegrationConfigurationDeleteResponseTriggerType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationDeleteResponseTriggerTypeEvent, IntegrationConfigurationDeleteResponseTriggerTypeSchedule:
		return true
	}
	return false
}

type IntegrationConfigurationEnableResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The type of destination _(e.g. Netsuite, webhooks)_.
	Destination string `json:"destination,required"`
	// The type of entity the integration is for _(e.g. Bill)_.
	EntityType string `json:"entityType,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A flag indicating whether the integration configuration is authorized.
	//
	// - TRUE - authorized.
	// - FALSE - not authorized.
	Authorized bool `json:"authorized"`
	// Configuration data for the integration
	ConfigData map[string]interface{} `json:"configData"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The unique identifier (UUID) of the entity the integration is for.
	DestinationID string `json:"destinationId"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// A flag indicating whether the integration configuration is currently enabled or
	// disabled.
	//
	// - TRUE - enabled.
	// - FALSE - disabled.
	Enabled bool `json:"enabled"`
	// The unique identifier (UUID) of the entity this integration is for _(e.g. the ID
	// of a notification configuration. Optional.)_
	EntityID string `json:"entityId"`
	// UUID of the credentials to use for this integration
	IntegrationCredentialsID string `json:"integrationCredentialsId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the configuration
	Name string `json:"name"`
	// Specifies the type of trigger for the integration.
	//
	// Deprecated: deprecated
	TriggerType IntegrationConfigurationEnableResponseTriggerType `json:"triggerType"`
	JSON        integrationConfigurationEnableResponseJSON        `json:"-"`
}

// integrationConfigurationEnableResponseJSON contains the JSON metadata for the
// struct [IntegrationConfigurationEnableResponse]
type integrationConfigurationEnableResponseJSON struct {
	ID                       apijson.Field
	Destination              apijson.Field
	EntityType               apijson.Field
	Version                  apijson.Field
	Authorized               apijson.Field
	ConfigData               apijson.Field
	CreatedBy                apijson.Field
	DestinationID            apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	Enabled                  apijson.Field
	EntityID                 apijson.Field
	IntegrationCredentialsID apijson.Field
	LastModifiedBy           apijson.Field
	Name                     apijson.Field
	TriggerType              apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *IntegrationConfigurationEnableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r integrationConfigurationEnableResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of trigger for the integration.
type IntegrationConfigurationEnableResponseTriggerType string

const (
	IntegrationConfigurationEnableResponseTriggerTypeEvent    IntegrationConfigurationEnableResponseTriggerType = "EVENT"
	IntegrationConfigurationEnableResponseTriggerTypeSchedule IntegrationConfigurationEnableResponseTriggerType = "SCHEDULE"
)

func (r IntegrationConfigurationEnableResponseTriggerType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationEnableResponseTriggerTypeEvent, IntegrationConfigurationEnableResponseTriggerTypeSchedule:
		return true
	}
	return false
}

type IntegrationConfigurationNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// A flexible object to include any additional configuration data specific to the
	// integration.
	ConfigData param.Field[map[string]interface{}] `json:"configData,required"`
	// Base model for defining integration credentials across different types of
	// integrations.
	Credentials param.Field[IntegrationConfigurationNewParamsCredentials] `json:"credentials,required"`
	// Denotes the integration destination. This field identifies the target platform
	// or service for the integration.
	Destination param.Field[string] `json:"destination,required"`
	// The unique identifier (UUID) for the integration destination.
	DestinationID param.Field[string] `json:"destinationId,required"`
	// The unique identifier (UUID) of the entity. This field is used to specify which
	// entity's integration configuration you're updating.
	EntityID param.Field[string] `json:"entityId,required"`
	// Specifies the type of entity for which the integration configuration is being
	// updated. Must be a valid alphanumeric string.
	EntityType               param.Field[string] `json:"entityType,required"`
	IntegrationCredentialsID param.Field[string] `json:"integrationCredentialsId,required"`
	Name                     param.Field[string] `json:"name,required"`
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

func (r IntegrationConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Base model for defining integration credentials across different types of
// integrations.
type IntegrationConfigurationNewParamsCredentials struct {
	// Specifies the type of authorization required for the integration.
	Type        param.Field[IntegrationConfigurationNewParamsCredentialsType]        `json:"type,required"`
	Destination param.Field[IntegrationConfigurationNewParamsCredentialsDestination] `json:"destination"`
	// A flag to indicate whether the credentials are empty.
	//
	// - TRUE - empty credentials.
	// - FALSE - credential details required.
	Empty param.Field[bool] `json:"empty"`
	// The name of the credentials
	Name param.Field[string] `json:"name"`
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

func (r IntegrationConfigurationNewParamsCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of authorization required for the integration.
type IntegrationConfigurationNewParamsCredentialsType string

const (
	IntegrationConfigurationNewParamsCredentialsTypeHTTPBasic              IntegrationConfigurationNewParamsCredentialsType = "HTTP_BASIC"
	IntegrationConfigurationNewParamsCredentialsTypeOAuthClientCredentials IntegrationConfigurationNewParamsCredentialsType = "OAUTH_CLIENT_CREDENTIALS"
	IntegrationConfigurationNewParamsCredentialsTypeM3TerSignedRequest     IntegrationConfigurationNewParamsCredentialsType = "M3TER_SIGNED_REQUEST"
	IntegrationConfigurationNewParamsCredentialsTypeAwsIntegration         IntegrationConfigurationNewParamsCredentialsType = "AWS_INTEGRATION"
	IntegrationConfigurationNewParamsCredentialsTypePaddleAuth             IntegrationConfigurationNewParamsCredentialsType = "PADDLE_AUTH"
	IntegrationConfigurationNewParamsCredentialsTypeNetsuiteAuth           IntegrationConfigurationNewParamsCredentialsType = "NETSUITE_AUTH"
	IntegrationConfigurationNewParamsCredentialsTypeChargebeeAuth          IntegrationConfigurationNewParamsCredentialsType = "CHARGEBEE_AUTH"
	IntegrationConfigurationNewParamsCredentialsTypeM3TerServiceUser       IntegrationConfigurationNewParamsCredentialsType = "M3TER_SERVICE_USER"
	IntegrationConfigurationNewParamsCredentialsTypeStripeSignedRequest    IntegrationConfigurationNewParamsCredentialsType = "STRIPE_SIGNED_REQUEST"
)

func (r IntegrationConfigurationNewParamsCredentialsType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationNewParamsCredentialsTypeHTTPBasic, IntegrationConfigurationNewParamsCredentialsTypeOAuthClientCredentials, IntegrationConfigurationNewParamsCredentialsTypeM3TerSignedRequest, IntegrationConfigurationNewParamsCredentialsTypeAwsIntegration, IntegrationConfigurationNewParamsCredentialsTypePaddleAuth, IntegrationConfigurationNewParamsCredentialsTypeNetsuiteAuth, IntegrationConfigurationNewParamsCredentialsTypeChargebeeAuth, IntegrationConfigurationNewParamsCredentialsTypeM3TerServiceUser, IntegrationConfigurationNewParamsCredentialsTypeStripeSignedRequest:
		return true
	}
	return false
}

type IntegrationConfigurationNewParamsCredentialsDestination string

const (
	IntegrationConfigurationNewParamsCredentialsDestinationWebhook           IntegrationConfigurationNewParamsCredentialsDestination = "WEBHOOK"
	IntegrationConfigurationNewParamsCredentialsDestinationNetsuite          IntegrationConfigurationNewParamsCredentialsDestination = "NETSUITE"
	IntegrationConfigurationNewParamsCredentialsDestinationStripe            IntegrationConfigurationNewParamsCredentialsDestination = "STRIPE"
	IntegrationConfigurationNewParamsCredentialsDestinationStripeTest        IntegrationConfigurationNewParamsCredentialsDestination = "STRIPE_TEST"
	IntegrationConfigurationNewParamsCredentialsDestinationAws               IntegrationConfigurationNewParamsCredentialsDestination = "AWS"
	IntegrationConfigurationNewParamsCredentialsDestinationPaddle            IntegrationConfigurationNewParamsCredentialsDestination = "PADDLE"
	IntegrationConfigurationNewParamsCredentialsDestinationPaddleSandbox     IntegrationConfigurationNewParamsCredentialsDestination = "PADDLE_SANDBOX"
	IntegrationConfigurationNewParamsCredentialsDestinationSalesforce        IntegrationConfigurationNewParamsCredentialsDestination = "SALESFORCE"
	IntegrationConfigurationNewParamsCredentialsDestinationXero              IntegrationConfigurationNewParamsCredentialsDestination = "XERO"
	IntegrationConfigurationNewParamsCredentialsDestinationChargebee         IntegrationConfigurationNewParamsCredentialsDestination = "CHARGEBEE"
	IntegrationConfigurationNewParamsCredentialsDestinationQuickbooks        IntegrationConfigurationNewParamsCredentialsDestination = "QUICKBOOKS"
	IntegrationConfigurationNewParamsCredentialsDestinationQuickbooksSandbox IntegrationConfigurationNewParamsCredentialsDestination = "QUICKBOOKS_SANDBOX"
	IntegrationConfigurationNewParamsCredentialsDestinationM3Ter             IntegrationConfigurationNewParamsCredentialsDestination = "M3TER"
)

func (r IntegrationConfigurationNewParamsCredentialsDestination) IsKnown() bool {
	switch r {
	case IntegrationConfigurationNewParamsCredentialsDestinationWebhook, IntegrationConfigurationNewParamsCredentialsDestinationNetsuite, IntegrationConfigurationNewParamsCredentialsDestinationStripe, IntegrationConfigurationNewParamsCredentialsDestinationStripeTest, IntegrationConfigurationNewParamsCredentialsDestinationAws, IntegrationConfigurationNewParamsCredentialsDestinationPaddle, IntegrationConfigurationNewParamsCredentialsDestinationPaddleSandbox, IntegrationConfigurationNewParamsCredentialsDestinationSalesforce, IntegrationConfigurationNewParamsCredentialsDestinationXero, IntegrationConfigurationNewParamsCredentialsDestinationChargebee, IntegrationConfigurationNewParamsCredentialsDestinationQuickbooks, IntegrationConfigurationNewParamsCredentialsDestinationQuickbooksSandbox, IntegrationConfigurationNewParamsCredentialsDestinationM3Ter:
		return true
	}
	return false
}

type IntegrationConfigurationGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type IntegrationConfigurationUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// A flexible object to include any additional configuration data specific to the
	// integration.
	ConfigData param.Field[map[string]interface{}] `json:"configData,required"`
	// Base model for defining integration credentials across different types of
	// integrations.
	Credentials param.Field[IntegrationConfigurationUpdateParamsCredentials] `json:"credentials,required"`
	// Denotes the integration destination. This field identifies the target platform
	// or service for the integration.
	Destination param.Field[string] `json:"destination,required"`
	// The unique identifier (UUID) for the integration destination.
	DestinationID param.Field[string] `json:"destinationId,required"`
	// The unique identifier (UUID) of the entity. This field is used to specify which
	// entity's integration configuration you're updating.
	EntityID param.Field[string] `json:"entityId,required"`
	// Specifies the type of entity for which the integration configuration is being
	// updated. Must be a valid alphanumeric string.
	EntityType               param.Field[string] `json:"entityType,required"`
	IntegrationCredentialsID param.Field[string] `json:"integrationCredentialsId,required"`
	Name                     param.Field[string] `json:"name,required"`
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

func (r IntegrationConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Base model for defining integration credentials across different types of
// integrations.
type IntegrationConfigurationUpdateParamsCredentials struct {
	// Specifies the type of authorization required for the integration.
	Type        param.Field[IntegrationConfigurationUpdateParamsCredentialsType]        `json:"type,required"`
	Destination param.Field[IntegrationConfigurationUpdateParamsCredentialsDestination] `json:"destination"`
	// A flag to indicate whether the credentials are empty.
	//
	// - TRUE - empty credentials.
	// - FALSE - credential details required.
	Empty param.Field[bool] `json:"empty"`
	// The name of the credentials
	Name param.Field[string] `json:"name"`
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

func (r IntegrationConfigurationUpdateParamsCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of authorization required for the integration.
type IntegrationConfigurationUpdateParamsCredentialsType string

const (
	IntegrationConfigurationUpdateParamsCredentialsTypeHTTPBasic              IntegrationConfigurationUpdateParamsCredentialsType = "HTTP_BASIC"
	IntegrationConfigurationUpdateParamsCredentialsTypeOAuthClientCredentials IntegrationConfigurationUpdateParamsCredentialsType = "OAUTH_CLIENT_CREDENTIALS"
	IntegrationConfigurationUpdateParamsCredentialsTypeM3TerSignedRequest     IntegrationConfigurationUpdateParamsCredentialsType = "M3TER_SIGNED_REQUEST"
	IntegrationConfigurationUpdateParamsCredentialsTypeAwsIntegration         IntegrationConfigurationUpdateParamsCredentialsType = "AWS_INTEGRATION"
	IntegrationConfigurationUpdateParamsCredentialsTypePaddleAuth             IntegrationConfigurationUpdateParamsCredentialsType = "PADDLE_AUTH"
	IntegrationConfigurationUpdateParamsCredentialsTypeNetsuiteAuth           IntegrationConfigurationUpdateParamsCredentialsType = "NETSUITE_AUTH"
	IntegrationConfigurationUpdateParamsCredentialsTypeChargebeeAuth          IntegrationConfigurationUpdateParamsCredentialsType = "CHARGEBEE_AUTH"
	IntegrationConfigurationUpdateParamsCredentialsTypeM3TerServiceUser       IntegrationConfigurationUpdateParamsCredentialsType = "M3TER_SERVICE_USER"
	IntegrationConfigurationUpdateParamsCredentialsTypeStripeSignedRequest    IntegrationConfigurationUpdateParamsCredentialsType = "STRIPE_SIGNED_REQUEST"
)

func (r IntegrationConfigurationUpdateParamsCredentialsType) IsKnown() bool {
	switch r {
	case IntegrationConfigurationUpdateParamsCredentialsTypeHTTPBasic, IntegrationConfigurationUpdateParamsCredentialsTypeOAuthClientCredentials, IntegrationConfigurationUpdateParamsCredentialsTypeM3TerSignedRequest, IntegrationConfigurationUpdateParamsCredentialsTypeAwsIntegration, IntegrationConfigurationUpdateParamsCredentialsTypePaddleAuth, IntegrationConfigurationUpdateParamsCredentialsTypeNetsuiteAuth, IntegrationConfigurationUpdateParamsCredentialsTypeChargebeeAuth, IntegrationConfigurationUpdateParamsCredentialsTypeM3TerServiceUser, IntegrationConfigurationUpdateParamsCredentialsTypeStripeSignedRequest:
		return true
	}
	return false
}

type IntegrationConfigurationUpdateParamsCredentialsDestination string

const (
	IntegrationConfigurationUpdateParamsCredentialsDestinationWebhook           IntegrationConfigurationUpdateParamsCredentialsDestination = "WEBHOOK"
	IntegrationConfigurationUpdateParamsCredentialsDestinationNetsuite          IntegrationConfigurationUpdateParamsCredentialsDestination = "NETSUITE"
	IntegrationConfigurationUpdateParamsCredentialsDestinationStripe            IntegrationConfigurationUpdateParamsCredentialsDestination = "STRIPE"
	IntegrationConfigurationUpdateParamsCredentialsDestinationStripeTest        IntegrationConfigurationUpdateParamsCredentialsDestination = "STRIPE_TEST"
	IntegrationConfigurationUpdateParamsCredentialsDestinationAws               IntegrationConfigurationUpdateParamsCredentialsDestination = "AWS"
	IntegrationConfigurationUpdateParamsCredentialsDestinationPaddle            IntegrationConfigurationUpdateParamsCredentialsDestination = "PADDLE"
	IntegrationConfigurationUpdateParamsCredentialsDestinationPaddleSandbox     IntegrationConfigurationUpdateParamsCredentialsDestination = "PADDLE_SANDBOX"
	IntegrationConfigurationUpdateParamsCredentialsDestinationSalesforce        IntegrationConfigurationUpdateParamsCredentialsDestination = "SALESFORCE"
	IntegrationConfigurationUpdateParamsCredentialsDestinationXero              IntegrationConfigurationUpdateParamsCredentialsDestination = "XERO"
	IntegrationConfigurationUpdateParamsCredentialsDestinationChargebee         IntegrationConfigurationUpdateParamsCredentialsDestination = "CHARGEBEE"
	IntegrationConfigurationUpdateParamsCredentialsDestinationQuickbooks        IntegrationConfigurationUpdateParamsCredentialsDestination = "QUICKBOOKS"
	IntegrationConfigurationUpdateParamsCredentialsDestinationQuickbooksSandbox IntegrationConfigurationUpdateParamsCredentialsDestination = "QUICKBOOKS_SANDBOX"
	IntegrationConfigurationUpdateParamsCredentialsDestinationM3Ter             IntegrationConfigurationUpdateParamsCredentialsDestination = "M3TER"
)

func (r IntegrationConfigurationUpdateParamsCredentialsDestination) IsKnown() bool {
	switch r {
	case IntegrationConfigurationUpdateParamsCredentialsDestinationWebhook, IntegrationConfigurationUpdateParamsCredentialsDestinationNetsuite, IntegrationConfigurationUpdateParamsCredentialsDestinationStripe, IntegrationConfigurationUpdateParamsCredentialsDestinationStripeTest, IntegrationConfigurationUpdateParamsCredentialsDestinationAws, IntegrationConfigurationUpdateParamsCredentialsDestinationPaddle, IntegrationConfigurationUpdateParamsCredentialsDestinationPaddleSandbox, IntegrationConfigurationUpdateParamsCredentialsDestinationSalesforce, IntegrationConfigurationUpdateParamsCredentialsDestinationXero, IntegrationConfigurationUpdateParamsCredentialsDestinationChargebee, IntegrationConfigurationUpdateParamsCredentialsDestinationQuickbooks, IntegrationConfigurationUpdateParamsCredentialsDestinationQuickbooksSandbox, IntegrationConfigurationUpdateParamsCredentialsDestinationM3Ter:
		return true
	}
	return false
}

type IntegrationConfigurationListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// integration configurations in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of integration configurations to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [IntegrationConfigurationListParams]'s query parameters as
// `url.Values`.
func (r IntegrationConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntegrationConfigurationDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type IntegrationConfigurationEnableParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type IntegrationConfigurationGetByEntityParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// UUID of the entity to retrieve IntegrationConfigs for
	EntityID param.Field[string] `query:"entityId"`
}

// URLQuery serializes [IntegrationConfigurationGetByEntityParams]'s query
// parameters as `url.Values`.
func (r IntegrationConfigurationGetByEntityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
