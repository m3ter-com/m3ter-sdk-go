// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
)

// This section provides Endpoints for operations that allow you to retrieve
// detailed information about individual Events, list all Events or specific Event
// Types, and explore dynamic fields available for each Event Type.
//
// Events encompass specific instances of state changes within the system, such as
// the creation of a new Prepayment/Commitment for an Account. Each Event is
// classified under an Event Type framework, providing context about what kind of
// change occurred to generate the Event.
//
// **Events for Configuration and Billing Entities**
//
// Many Event Types cover common configuration and billing objects, where the Event
// is generated for a state change of one of these objects - for when the
// configuration or billing object is **created**, **deleted**, or **updated**.
//
// For example:
//
// - configuration.commitment.created
// - configuration.commitment.deleted
// - configuration.commitment.updated
// - configuration.account.created
// - configuration.account.deleted
// - configuration.account.updated
// - billing.bill.created
// - billing.bill.deleted
// - billing.bill.created
//
// **Events for Errors or Failures**
//
// There are also Event Types for certain kinds of error that can occur:
//
// - For an Integration:
//
//   - validation
//   - authentication
//   - perform
//   - missing account mapping
//   - disabled
//
// - For a Usage Data Ingest Submission:
//
//   - validation failure
//
// - For Data Export Jobs:
//   - data export job failure
//
// **Scheduled Events**
//
// In addition to system-generated Events that occur when a configuration entity
// undergoes a state change at creation, update, or deletion of the entity, you can
// use API calls to create and configure _Scheduled Event Configurations_.
// Scheduled Events are custom Event types, which you can set up by referencing
// Date/Time fields on configuration and billing entities. See the
// [ScheduledEventConfigurations](https://www.m3ter.com/docs/api#tag/ScheduledEventConfigurations)
// section of this API Reference for more details.
//
// **Notifications for Events**
//
// You can create Notification rules based on Events and these rules can reference
// and apply calculations to the Event's fields. This allows you to set up
// customized alerts to be sent out via webhooks when the Event occurs and any
// conditions you've built into the Notification rule's calculation are satisfied.
//
// See the [Notifications](https://www.m3ter.com/docs/api#tag/Notifications)
// section for more details.
//
// **Other Events**
//
// When Events occur, they can cause other Events, such as when a Notification is
// triggered by the Event it is based on. For these Events there are currently two
// categories:
//
// - Notification
// - IntegrationEvent
//
// Also see
// [Utilizing Events and Notifications](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications)
// and
// [Object Definitions and API Calls](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/object-definitions-and-api-calls)
// in the m3ter documentation for more guidance.
//
// EventService contains methods and other services that help with interacting with
// the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r *EventService) {
	r = &EventService{}
	r.Options = opts
	return
}

// Retrieve a specific Event.
//
// Retrieves detailed information about the specific Event with the given UUID. An
// Event corresponds to a unique instance of a state change within the system,
// classified under a specific Event Type.
func (r *EventService) Get(ctx context.Context, id string, query EventGetParams, opts ...option.RequestOption) (res *EventResponse, err error) {
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("organizations/%s/events/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all Events.
//
// Retrieve a list of all Events, with options to filter the returned list based on
// various criteria. Each Event represents a unique instance of a state change
// within the system, classified under a specific kind of Event.
//
// **NOTES:** You can:
//
//   - Use `eventName` as a valid Query parameter to filter the list of Events
//     returned. For example:
//     `.../organizations/{orgId}/events?eventName=configuration.commitment.created`
//   - Use the
//     [List Notification Events](https://www.m3ter.com/docs/api#tag/Events/operation/ListEventTypes)
//     endpoint in this section. The response lists the valid Query parameters.
func (r *EventService) List(ctx context.Context, params EventListParams, opts ...option.RequestOption) (res *pagination.Cursor[EventResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("organizations/%s/events", params.OrgID)
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

// List all Events.
//
// Retrieve a list of all Events, with options to filter the returned list based on
// various criteria. Each Event represents a unique instance of a state change
// within the system, classified under a specific kind of Event.
//
// **NOTES:** You can:
//
//   - Use `eventName` as a valid Query parameter to filter the list of Events
//     returned. For example:
//     `.../organizations/{orgId}/events?eventName=configuration.commitment.created`
//   - Use the
//     [List Notification Events](https://www.m3ter.com/docs/api#tag/Events/operation/ListEventTypes)
//     endpoint in this section. The response lists the valid Query parameters.
func (r *EventService) ListAutoPaging(ctx context.Context, params EventListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[EventResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// List Event Fields.
//
// Retrieves a list of Fields for a specific Event Type. These Fields are dynamic
// and forward compatibile, enabling calculation operations on the Event schema.
//
// **Notes:**
//
//   - In many of the Response schema for this call, such as when you retrieve the
//     Fields for a `configuration.commitment.created` Event Type, `new` represents
//     the attributes the newly created object has. The Response to a call to
//     retrieve the Fields for a `configuration.commitment.updated` Event Type will
//     contain Field values for both the `old` and `new` objects. The Response to a
//     call to retrieve the Fields for a `configuration.commitment.deleted` Event
//     Type will only contain `old` Fields, for values at point of deletion. Having
//     access to reference both `new` and `old` Field values for an object can be
//     very useful if you want to base a Notification rule on an Event and include a
//     calculation in the rule that, for example, compares `new` values with `old` -
//     for example, trigger a Notification only when a Commitment has been updated
//     and the `new` value for the `amount` is at least $1,000 greater than the `old`
//     value.
//   - Some Event types will show `customFields` even though the specific billing or
//     configuration object the Event is for does not yet have the custom fields
//     functionality implemented. For these Events, their `customFields` values will
//     not be populated until such time as the custom fields functionality is
//     implemented for them
func (r *EventService) GetFields(ctx context.Context, params EventGetFieldsParams, opts ...option.RequestOption) (res *EventGetFieldsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/events/fields", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Retrieve a list of Notification Event Types.
//
// This endpoint retrieves a list of Event Types that can have Notification rules
// configured.
func (r *EventService) GetTypes(ctx context.Context, query EventGetTypesParams, opts ...option.RequestOption) (res *EventGetTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.OrgID, precfg.OrgID)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/events/types", query.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing an Event entity.
type EventResponse struct {
	// The uniqie identifier (UUID) of the Event.
	ID string `json:"id" api:"required"`
	// When an Event was actioned. It follows the ISO 8601 date and time format.
	//
	// You can action an Event to indicate that it has been followed up and resolved -
	// this is useful when dealing with integration error Events or ingest failure
	// Events.
	DtActioned time.Time `json:"dtActioned" api:"required" format:"date-time"`
	// The name of the Event as it is registered in the system. This name is used to
	// categorize and trigger associated actions.
	EventName string `json:"eventName" api:"required"`
	// The time when the Event was triggered, using the ISO 8601 date and time format.
	EventTime time.Time `json:"eventTime" api:"required" format:"date-time"`
	// The Data Transfer Object (DTO) containing the details of the Event.
	M3terEvent interface{}       `json:"m3terEvent" api:"required"`
	JSON       eventResponseJSON `json:"-"`
}

// eventResponseJSON contains the JSON metadata for the struct [EventResponse]
type eventResponseJSON struct {
	ID          apijson.Field
	DtActioned  apijson.Field
	EventName   apijson.Field
	EventTime   apijson.Field
	M3terEvent  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing the list of Fields for an Event Type.
type EventGetFieldsResponse struct {
	// An object containing the list of Fields for the queried Event Type.
	//
	// See the 200 Response sample where we have queried to get the Fields for the
	// `configuration.commitment.created` Event Type.
	//
	// **Note:** `new` represents the attributes the newly created object has.
	Events map[string]map[string]string `json:"events"`
	JSON   eventGetFieldsResponseJSON   `json:"-"`
}

// eventGetFieldsResponseJSON contains the JSON metadata for the struct
// [EventGetFieldsResponse]
type eventGetFieldsResponseJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventGetFieldsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventGetFieldsResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing list of Event Types that can have Notification rules
// configured.
type EventGetTypesResponse struct {
	// An array containing a list of all Event Types for which Notification rules can
	// be configured. Each Event Type is represented by a string.
	Events []string                  `json:"events"`
	JSON   eventGetTypesResponseJSON `json:"-"`
}

// eventGetTypesResponseJSON contains the JSON metadata for the struct
// [EventGetTypesResponse]
type eventGetTypesResponseJSON struct {
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventGetTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventGetTypesResponseJSON) RawJSON() string {
	return r.raw
}

type EventGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type EventListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The Account ID associated with the Event to filter the results. Returns the
	// Events that have been generated for the Account.
	AccountID param.Field[string] `query:"accountId"`
	EventName param.Field[string] `query:"eventName"`
	// The category of Events to filter the results by. Options:
	//
	// - Notification
	// - IntegrationEvent
	// - IngestValidationFailure
	// - DataExportJobFailure
	EventType param.Field[string] `query:"eventType"`
	// List of Event UUIDs to filter the results.
	//
	// **NOTE:** cannot be used with other filters.
	IDs param.Field[[]string] `query:"ids"`
	// A Boolean flag indicating whether to return Events that have been actioned.
	//
	// - **TRUE** - include actioned Events.
	// - **FALSE** - exclude actioned Events.
	IncludeActioned param.Field[bool] `query:"includeActioned"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Events in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Short code of the Notification to filter the results. Returns the Events that
	// have triggered the Notification.
	NotificationCode param.Field[string] `query:"notificationCode"`
	// Notification UUID to filter the results. Returns the Events that have triggered
	// the Notification.
	NotificationID param.Field[string] `query:"notificationId"`
	// The maximum number of Events to retrieve per page.
	PageSize   param.Field[int64]  `query:"pageSize"`
	ResourceID param.Field[string] `query:"resourceId"`
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventGetFieldsParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The name of the specific Event Type to use as a list filter, for example
	// `configuration.commitment.created`.
	EventName param.Field[string] `query:"eventName"`
}

// URLQuery serializes [EventGetFieldsParams]'s query parameters as `url.Values`.
func (r EventGetFieldsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EventGetTypesParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
