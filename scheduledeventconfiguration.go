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

// ScheduledEventConfigurationService contains methods and other services that help
// with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduledEventConfigurationService] method instead.
type ScheduledEventConfigurationService struct {
	Options []option.RequestOption
}

// NewScheduledEventConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewScheduledEventConfigurationService(opts ...option.RequestOption) (r *ScheduledEventConfigurationService) {
	r = &ScheduledEventConfigurationService{}
	r.Options = opts
	return
}

// Create a new ScheduledEventConfiguration.
func (r *ScheduledEventConfigurationService) New(ctx context.Context, params ScheduledEventConfigurationNewParams, opts ...option.RequestOption) (res *ScheduledEventConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/scheduledevents/configurations", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a ScheduledEventConfiguration for the given UUID.
func (r *ScheduledEventConfigurationService) Get(ctx context.Context, id string, query ScheduledEventConfigurationGetParams, opts ...option.RequestOption) (res *ScheduledEventConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/scheduledevents/configurations/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a ScheduledEventConfiguration for the given UUID.
func (r *ScheduledEventConfigurationService) Update(ctx context.Context, id string, params ScheduledEventConfigurationUpdateParams, opts ...option.RequestOption) (res *ScheduledEventConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/scheduledevents/configurations/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of ScheduledEventConfiguration entities
func (r *ScheduledEventConfigurationService) List(ctx context.Context, params ScheduledEventConfigurationListParams, opts ...option.RequestOption) (res *pagination.Cursor[ScheduledEventConfigurationResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/scheduledevents/configurations", params.OrgID)
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

// Retrieve a list of ScheduledEventConfiguration entities
func (r *ScheduledEventConfigurationService) ListAutoPaging(ctx context.Context, params ScheduledEventConfigurationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ScheduledEventConfigurationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the ScheduledEventConfiguration for the given UUID.
func (r *ScheduledEventConfigurationService) Delete(ctx context.Context, id string, body ScheduledEventConfigurationDeleteParams, opts ...option.RequestOption) (res *ScheduledEventConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/scheduledevents/configurations/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ScheduledEventConfigurationResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The referenced configuration or billing entity for which the desired scheduled
	// Event will trigger.
	Entity string `json:"entity,required"`
	// A DateTime field for which the desired scheduled Event will trigger - this must
	// be a DateTime field on the referenced billing or configuration entity.
	Field string `json:"field,required"`
	// The name of the custom Scheduled Event Configuration.
	Name string `json:"name,required"`
	// The offset in days from the specified DateTime field on the referenced entity
	// when the scheduled Event will trigger.
	Offset int64 `json:"offset,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this item.
	LastModifiedBy string                                  `json:"lastModifiedBy"`
	JSON           scheduledEventConfigurationResponseJSON `json:"-"`
}

// scheduledEventConfigurationResponseJSON contains the JSON metadata for the
// struct [ScheduledEventConfigurationResponse]
type scheduledEventConfigurationResponseJSON struct {
	ID             apijson.Field
	Entity         apijson.Field
	Field          apijson.Field
	Name           apijson.Field
	Offset         apijson.Field
	Version        apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScheduledEventConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledEventConfigurationResponseJSON) RawJSON() string {
	return r.raw
}

type ScheduledEventConfigurationNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The referenced configuration or billing entity for which the desired scheduled
	// Event will trigger.
	Entity param.Field[string] `json:"entity,required"`
	// A DateTime field for which the desired scheduled Event will trigger - this must
	// be a DateTime field on the referenced billing or configuration entity.
	Field param.Field[string] `json:"field,required"`
	// The name of the custom Scheduled Event Configuration.
	//
	// This must be in the format:
	//
	// - scheduled._name of entity_._custom event name_
	//
	// For example:
	//
	// - `scheduled.bill.endDateEvent`
	Name param.Field[string] `json:"name,required"`
	// The offset in days from the specified DateTime field on the referenced entity
	// when the scheduled Event will trigger.
	Offset param.Field[int64] `json:"offset,required"`
	// The version number of the scheduled event configuration:
	//
	//   - **Create entity**: Not valid for initial insertion - do not use for Create. On
	//     initial Create, version is set at 1 and listed in the response.
	//   - **Update Entity**: On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r ScheduledEventConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduledEventConfigurationGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type ScheduledEventConfigurationUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The referenced configuration or billing entity for which the desired scheduled
	// Event will trigger.
	Entity param.Field[string] `json:"entity,required"`
	// A DateTime field for which the desired scheduled Event will trigger - this must
	// be a DateTime field on the referenced billing or configuration entity.
	Field param.Field[string] `json:"field,required"`
	// The name of the custom Scheduled Event Configuration.
	//
	// This must be in the format:
	//
	// - scheduled._name of entity_._custom event name_
	//
	// For example:
	//
	// - `scheduled.bill.endDateEvent`
	Name param.Field[string] `json:"name,required"`
	// The offset in days from the specified DateTime field on the referenced entity
	// when the scheduled Event will trigger.
	Offset param.Field[int64] `json:"offset,required"`
	// The version number of the scheduled event configuration:
	//
	//   - **Create entity**: Not valid for initial insertion - do not use for Create. On
	//     initial Create, version is set at 1 and listed in the response.
	//   - **Update Entity**: On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r ScheduledEventConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScheduledEventConfigurationListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// list of UUIDs to retrieve
	IDs param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of ScheduledEventConfigurations to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ScheduledEventConfigurationListParams]'s query parameters
// as `url.Values`.
func (r ScheduledEventConfigurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ScheduledEventConfigurationDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
