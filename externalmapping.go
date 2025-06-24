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

// ExternalMappingService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalMappingService] method instead.
type ExternalMappingService struct {
	Options []option.RequestOption
}

// NewExternalMappingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExternalMappingService(opts ...option.RequestOption) (r *ExternalMappingService) {
	r = &ExternalMappingService{}
	r.Options = opts
	return
}

// Creates a new External Mapping.
//
// This endpoint enables you to create a new External Mapping for the specified
// Organization. You need to supply a request body with the details of the new
// External Mapping.
func (r *ExternalMappingService) New(ctx context.Context, params ExternalMappingNewParams, opts ...option.RequestOption) (res *ExternalMappingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/externalmappings", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve an External Mapping with the given UUID.
//
// This endpoint enables you to retrieve the External Mapping with the specified
// UUID for a specific Organization.
func (r *ExternalMappingService) Get(ctx context.Context, id string, query ExternalMappingGetParams, opts ...option.RequestOption) (res *ExternalMappingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/externalmappings/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an External Mapping with the given UUID.
//
// This endpoint enables you to update an existing External Mapping entity,
// identified by its UUID. You must supply a request body with the new details for
// the External Mapping.
func (r *ExternalMappingService) Update(ctx context.Context, id string, params ExternalMappingUpdateParams, opts ...option.RequestOption) (res *ExternalMappingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/externalmappings/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of all External Mapping entities.
//
// This endpoint retrieves a list of all External Mapping entities for a specific
// Organization. The list can be paginated for better management, and supports
// filtering using the external system.
func (r *ExternalMappingService) List(ctx context.Context, params ExternalMappingListParams, opts ...option.RequestOption) (res *pagination.Cursor[ExternalMappingResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/externalmappings", params.OrgID)
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

// Retrieve a list of all External Mapping entities.
//
// This endpoint retrieves a list of all External Mapping entities for a specific
// Organization. The list can be paginated for better management, and supports
// filtering using the external system.
func (r *ExternalMappingService) ListAutoPaging(ctx context.Context, params ExternalMappingListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ExternalMappingResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete an External Mapping with the given UUID.
func (r *ExternalMappingService) Delete(ctx context.Context, id string, body ExternalMappingDeleteParams, opts ...option.RequestOption) (res *ExternalMappingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/externalmappings/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve a list of External Mapping entities for a specified external system
// entity.
//
// Use this endpoint to retrieve a list of External Mapping entities associated
// with a specific external system entity. The list can be paginated for easier
// management.
func (r *ExternalMappingService) ListByExternalEntity(ctx context.Context, system string, externalTable string, externalID string, params ExternalMappingListByExternalEntityParams, opts ...option.RequestOption) (res *pagination.Cursor[ExternalMappingResponse], err error) {
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
	if system == "" {
		err = errors.New("missing required system parameter")
		return
	}
	if externalTable == "" {
		err = errors.New("missing required externalTable parameter")
		return
	}
	if externalID == "" {
		err = errors.New("missing required externalId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/externalmappings/externalid/%s/%s/%s", params.OrgID, system, externalTable, externalID)
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

// Retrieve a list of External Mapping entities for a specified external system
// entity.
//
// Use this endpoint to retrieve a list of External Mapping entities associated
// with a specific external system entity. The list can be paginated for easier
// management.
func (r *ExternalMappingService) ListByExternalEntityAutoPaging(ctx context.Context, system string, externalTable string, externalID string, params ExternalMappingListByExternalEntityParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ExternalMappingResponse] {
	return pagination.NewCursorAutoPager(r.ListByExternalEntity(ctx, system, externalTable, externalID, params, opts...))
}

// Retrieve a list of External Mapping entities for a specified m3ter entity.
//
// Use this endpoint to retrieve a list of External Mapping entities associated
// with a specific m3ter entity. The list can be paginated for easier management.
func (r *ExternalMappingService) ListByM3terEntity(ctx context.Context, entity string, m3terID string, params ExternalMappingListByM3terEntityParams, opts ...option.RequestOption) (res *pagination.Cursor[ExternalMappingResponse], err error) {
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
	if entity == "" {
		err = errors.New("missing required entity parameter")
		return
	}
	if m3terID == "" {
		err = errors.New("missing required m3terId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/externalmappings/external/%s/%s", params.OrgID, entity, m3terID)
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

// Retrieve a list of External Mapping entities for a specified m3ter entity.
//
// Use this endpoint to retrieve a list of External Mapping entities associated
// with a specific m3ter entity. The list can be paginated for easier management.
func (r *ExternalMappingService) ListByM3terEntityAutoPaging(ctx context.Context, entity string, m3terID string, params ExternalMappingListByM3terEntityParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ExternalMappingResponse] {
	return pagination.NewCursorAutoPager(r.ListByM3terEntity(ctx, entity, m3terID, params, opts...))
}

type ExternalMappingResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The unique identifier (UUID) of the entity in the external system.
	ExternalID string `json:"externalId,required"`
	// The name of the external system where the entity you are mapping resides.
	ExternalSystem string `json:"externalSystem,required"`
	// The name of the table in the external system where the entity resides.
	ExternalTable string `json:"externalTable,required"`
	// The name of the m3ter entity that is part of the External Mapping. For example,
	// this could be "Account".
	M3terEntity string `json:"m3terEntity,required"`
	// The unique identifier (UUID) of the m3ter entity.
	M3terID string `json:"m3terId,required"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// UUID of the configuration this mapping is for
	IntegrationConfigID string `json:"integrationConfigId"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                       `json:"version"`
	JSON    externalMappingResponseJSON `json:"-"`
}

// externalMappingResponseJSON contains the JSON metadata for the struct
// [ExternalMappingResponse]
type externalMappingResponseJSON struct {
	ID                  apijson.Field
	ExternalID          apijson.Field
	ExternalSystem      apijson.Field
	ExternalTable       apijson.Field
	M3terEntity         apijson.Field
	M3terID             apijson.Field
	CreatedBy           apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	IntegrationConfigID apijson.Field
	LastModifiedBy      apijson.Field
	Version             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ExternalMappingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r externalMappingResponseJSON) RawJSON() string {
	return r.raw
}

type ExternalMappingNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) of the entity in the external system. This UUID
	// should already exist in the external system.
	ExternalID param.Field[string] `json:"externalId,required"`
	// The name of the external system where the entity you are mapping resides.
	ExternalSystem param.Field[string] `json:"externalSystem,required"`
	// The name of the table in ther external system where the entity resides.
	ExternalTable param.Field[string] `json:"externalTable,required"`
	// The name of the m3ter entity that you are creating or modifying an External
	// Mapping for. As an example, this could be an "Account".
	M3terEntity param.Field[string] `json:"m3terEntity,required"`
	// The unique identifier (UUID) of the m3ter entity.
	M3terID param.Field[string] `json:"m3terId,required"`
	// UUID of the integration config to link this mapping to
	IntegrationConfigID param.Field[string] `json:"integrationConfigId"`
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

func (r ExternalMappingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalMappingGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type ExternalMappingUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) of the entity in the external system. This UUID
	// should already exist in the external system.
	ExternalID param.Field[string] `json:"externalId,required"`
	// The name of the external system where the entity you are mapping resides.
	ExternalSystem param.Field[string] `json:"externalSystem,required"`
	// The name of the table in ther external system where the entity resides.
	ExternalTable param.Field[string] `json:"externalTable,required"`
	// The name of the m3ter entity that you are creating or modifying an External
	// Mapping for. As an example, this could be an "Account".
	M3terEntity param.Field[string] `json:"m3terEntity,required"`
	// The unique identifier (UUID) of the m3ter entity.
	M3terID param.Field[string] `json:"m3terId,required"`
	// UUID of the integration config to link this mapping to
	IntegrationConfigID param.Field[string] `json:"integrationConfigId"`
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

func (r ExternalMappingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalMappingListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The name of the external system to use as a filter.
	//
	// For example, if you want to list only those external mappings created for your
	// Organization for the Salesforce external system, use:
	//
	// `?externalSystemId=Salesforce`
	ExternalSystemID param.Field[string] `query:"externalSystemId"`
	// ID of the integration config
	IntegrationConfigID param.Field[string] `query:"integrationConfigId"`
	// IDs for m3ter entities
	M3terIDs param.Field[[]string] `query:"m3terIds"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// External Mappings in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of External Mappings to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ExternalMappingListParams]'s query parameters as
// `url.Values`.
func (r ExternalMappingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalMappingDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type ExternalMappingListByExternalEntityParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// External Mappings in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of External Mappings to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ExternalMappingListByExternalEntityParams]'s query
// parameters as `url.Values`.
func (r ExternalMappingListByExternalEntityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExternalMappingListByM3terEntityParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// External Mappings in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of External Mappings to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ExternalMappingListByM3terEntityParams]'s query parameters
// as `url.Values`.
func (r ExternalMappingListByM3terEntityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
