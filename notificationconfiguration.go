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

// NotificationConfigurationService contains methods and other services that help
// with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNotificationConfigurationService] method instead.
type NotificationConfigurationService struct {
	Options []option.RequestOption
}

// NewNotificationConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNotificationConfigurationService(opts ...option.RequestOption) (r *NotificationConfigurationService) {
	r = &NotificationConfigurationService{}
	r.Options = opts
	return
}

// Create a new Notification for an Event.
//
// This endpoint enables you to create a new Event Notification for the specified
// Organization. You need to supply a request body with the details of the new
// Notification.
func (r *NotificationConfigurationService) New(ctx context.Context, params NotificationConfigurationNewParams, opts ...option.RequestOption) (res *NotificationConfigurationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/notifications/configurations", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the details of a specific Notification using its UUID. Includes the
// Event the Notification is based on, and any calculation referencing the Event's
// field and which defines further conditions that must be met to trigger the
// Notification when the Event occurs.
func (r *NotificationConfigurationService) Get(ctx context.Context, id string, query NotificationConfigurationGetParams, opts ...option.RequestOption) (res *NotificationConfigurationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/notifications/configurations/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Notification with the given UUID.
//
// This endpoint modifies the configuration details of an existing Notification.
// You can change the Event that triggers the Notification and/or update the
// conditions for sending the Notification.
func (r *NotificationConfigurationService) Update(ctx context.Context, id string, params NotificationConfigurationUpdateParams, opts ...option.RequestOption) (res *NotificationConfigurationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/notifications/configurations/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Event Notifications for the specified Organization.
//
// This endpoint retrieves a list of all Event Notifications for the Organization
// identified by its UUID. The list can be paginated for easier management. The
// list also supports filtering by parameters such as Notification UUID.
func (r *NotificationConfigurationService) List(ctx context.Context, params NotificationConfigurationListParams, opts ...option.RequestOption) (res *pagination.Cursor[NotificationConfigurationResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/notifications/configurations", params.OrgID)
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

// Retrieve a list of Event Notifications for the specified Organization.
//
// This endpoint retrieves a list of all Event Notifications for the Organization
// identified by its UUID. The list can be paginated for easier management. The
// list also supports filtering by parameters such as Notification UUID.
func (r *NotificationConfigurationService) ListAutoPaging(ctx context.Context, params NotificationConfigurationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[NotificationConfigurationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Notification with the given UUID.
//
// This endpoint permanently removes a specified Notification and its
// configuration. This action cannot be undone.
func (r *NotificationConfigurationService) Delete(ctx context.Context, id string, body NotificationConfigurationDeleteParams, opts ...option.RequestOption) (res *NotificationConfigurationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/notifications/configurations/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type NotificationConfigurationResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The short code for the Notification.
	Code string `json:"code,required"`
	// The description for the Notification providing a brief overview of its purpose
	// and functionality.
	Description string `json:"description,required"`
	// The name of the Notification.
	Name string `json:"name,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A Boolean flag indicating whether or not the Notification is active.
	//
	// - **TRUE** - active Notification.
	// - **FALSE** - inactive Notification.
	Active bool `json:"active"`
	// A Boolean flag indicating whether the Notification is always triggered,
	// regardless of other conditions and omitting reference to any calculation. This
	// means the Notification will be triggered simply by the Event it is based on
	// occurring and with no further conditions having to be met.
	//
	//   - **TRUE** - the Notification is always triggered and omits any reference to the
	//     calculation to check for other conditions being true before triggering the
	//     Notification.
	//   - **FALSE** - the Notification is only triggered when the Event it is based on
	//     occurs and any calculation is checked and all conditions defined by the
	//     calculation are met.
	AlwaysFireEvent bool `json:"alwaysFireEvent"`
	// A logical expression that that is evaluated to a Boolean. If it evaluates as
	// `True`, a Notification for the Event is created and sent to the configured
	// destination. Calculations can reference numeric, string, and boolean Event
	// fields.
	//
	// See
	// [Creating Calculations](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/key-concepts-and-relationships#creating-calculations)
	// in the m3ter documentation for more information.
	Calculation string `json:"calculation"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The name of the Event that the Notification is based on. When this Event occurs
	// and the calculation evaluates to `True`, the Notification is triggered.
	//
	// **Note:** If the Notification is set to always fire, then the Notification will
	// always be sent when the Event it is based on occurs, and without any other
	// conditions defined by a calculation having to be met.
	EventName string `json:"eventName"`
	// The ID of the user who last modified this item.
	LastModifiedBy string                                `json:"lastModifiedBy"`
	JSON           notificationConfigurationResponseJSON `json:"-"`
}

// notificationConfigurationResponseJSON contains the JSON metadata for the struct
// [NotificationConfigurationResponse]
type notificationConfigurationResponseJSON struct {
	ID              apijson.Field
	Code            apijson.Field
	Description     apijson.Field
	Name            apijson.Field
	Version         apijson.Field
	Active          apijson.Field
	AlwaysFireEvent apijson.Field
	Calculation     apijson.Field
	CreatedBy       apijson.Field
	DtCreated       apijson.Field
	DtLastModified  apijson.Field
	EventName       apijson.Field
	LastModifiedBy  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NotificationConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationConfigurationResponseJSON) RawJSON() string {
	return r.raw
}

type NotificationConfigurationNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The short code for the Notification.
	Code param.Field[string] `json:"code,required"`
	// The description for the Notification providing a brief overview of its purpose
	// and functionality.
	Description param.Field[string] `json:"description,required"`
	// The name of the _Event type_ that the Notification is based on. When an Event of
	// this type occurs and any calculation built into the Notification evaluates to
	// `True`, the Notification is triggered.
	//
	// **Note:** If the Notification is set to always fire, then the Notification will
	// always be sent when the Event of the type it is based on occurs, and without any
	// other conditions defined by a calculation having to be met.
	EventName param.Field[string] `json:"eventName,required"`
	// The name of the Notification.
	Name param.Field[string] `json:"name,required"`
	// Boolean flag that sets the Notification as active or inactive. Only active
	// Notifications are sent when triggered by the Event they are based on:
	//
	// - **TRUE** - set Notification as active.
	// - **FALSE** - set Notification as inactive.
	Active param.Field[bool] `json:"active"`
	// A Boolean flag indicating whether the Notification is always triggered,
	// regardless of other conditions and omitting reference to any calculation. This
	// means the Notification will be triggered simply by the Event it is based on
	// occurring and with no further conditions having to be met.
	//
	//   - **TRUE** - the Notification is always triggered and omits any reference to the
	//     calculation to check for other conditions being true before triggering the
	//     Notification.
	//   - **FALSE** - the Notification is only triggered when the Event it is based on
	//     occurs and any calculation is checked and all conditions defined by the
	//     calculation are met.
	AlwaysFireEvent param.Field[bool] `json:"alwaysFireEvent"`
	// A logical expression that that is evaluated to a Boolean. If it evaluates as
	// `True`, a Notification for the Event is created and sent to the configured
	// destination. Calculations can reference numeric, string, and boolean Event
	// fields.
	//
	// See
	// [Creating Calculations](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/key-concepts-and-relationships#creating-calculations)
	// in the m3ter documentation for more information.
	Calculation param.Field[string] `json:"calculation"`
	// The version number for the Notification:
	//
	//   - **Create:** Not valid for initial insertion of new entity - _do not use for
	//     Create_. On initial Create, version is set at 1 and listed in the response.
	//   - **Update:** On Update, version is required and must match the existing version
	//     because a check is performed to ensure sequential versioning is preserved.
	//     Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r NotificationConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NotificationConfigurationGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type NotificationConfigurationUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The short code for the Notification.
	Code param.Field[string] `json:"code,required"`
	// The description for the Notification providing a brief overview of its purpose
	// and functionality.
	Description param.Field[string] `json:"description,required"`
	// The name of the _Event type_ that the Notification is based on. When an Event of
	// this type occurs and any calculation built into the Notification evaluates to
	// `True`, the Notification is triggered.
	//
	// **Note:** If the Notification is set to always fire, then the Notification will
	// always be sent when the Event of the type it is based on occurs, and without any
	// other conditions defined by a calculation having to be met.
	EventName param.Field[string] `json:"eventName,required"`
	// The name of the Notification.
	Name param.Field[string] `json:"name,required"`
	// Boolean flag that sets the Notification as active or inactive. Only active
	// Notifications are sent when triggered by the Event they are based on:
	//
	// - **TRUE** - set Notification as active.
	// - **FALSE** - set Notification as inactive.
	Active param.Field[bool] `json:"active"`
	// A Boolean flag indicating whether the Notification is always triggered,
	// regardless of other conditions and omitting reference to any calculation. This
	// means the Notification will be triggered simply by the Event it is based on
	// occurring and with no further conditions having to be met.
	//
	//   - **TRUE** - the Notification is always triggered and omits any reference to the
	//     calculation to check for other conditions being true before triggering the
	//     Notification.
	//   - **FALSE** - the Notification is only triggered when the Event it is based on
	//     occurs and any calculation is checked and all conditions defined by the
	//     calculation are met.
	AlwaysFireEvent param.Field[bool] `json:"alwaysFireEvent"`
	// A logical expression that that is evaluated to a Boolean. If it evaluates as
	// `True`, a Notification for the Event is created and sent to the configured
	// destination. Calculations can reference numeric, string, and boolean Event
	// fields.
	//
	// See
	// [Creating Calculations](https://www.m3ter.com/docs/guides/utilizing-events-and-notifications/key-concepts-and-relationships#creating-calculations)
	// in the m3ter documentation for more information.
	Calculation param.Field[string] `json:"calculation"`
	// The version number for the Notification:
	//
	//   - **Create:** Not valid for initial insertion of new entity - _do not use for
	//     Create_. On initial Create, version is set at 1 and listed in the response.
	//   - **Update:** On Update, version is required and must match the existing version
	//     because a check is performed to ensure sequential versioning is preserved.
	//     Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r NotificationConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NotificationConfigurationListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// A Boolean flag indicating whether to retrieve only active or only inactive
	// Notifications.
	//
	// - **TRUE** - only active Notifications are returned.
	// - **FALSE** - only inactive Notifications are returned.
	Active param.Field[bool] `query:"active"`
	// Use this to filter the Notifications returned - only those Notifications that
	// are based on the _Event type_ specified by `eventName` are returned.
	EventName param.Field[string] `query:"eventName"`
	// A list of specific Notification UUIDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Notifications in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Notifications to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [NotificationConfigurationListParams]'s query parameters as
// `url.Values`.
func (r NotificationConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NotificationConfigurationDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
