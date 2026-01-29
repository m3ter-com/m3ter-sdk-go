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

// ResourceGroupService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceGroupService] method instead.
type ResourceGroupService struct {
	Options []option.RequestOption
}

// NewResourceGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceGroupService(opts ...option.RequestOption) (r *ResourceGroupService) {
	r = &ResourceGroupService{}
	r.Options = opts
	return
}

// Create a ResourceGroup for the UUID
func (r *ResourceGroupService) New(ctx context.Context, type_ string, params ResourceGroupNewParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s", params.OrgID, type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the ResourceGroup for the UUID
func (r *ResourceGroupService) Get(ctx context.Context, type_ string, id string, query ResourceGroupGetParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s", query.OrgID, type_, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the ResourceGroup for the UUID
func (r *ResourceGroupService) Update(ctx context.Context, type_ string, id string, params ResourceGroupUpdateParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s", params.OrgID, type_, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of ResourceGroup entities
func (r *ResourceGroupService) List(ctx context.Context, type_ string, params ResourceGroupListParams, opts ...option.RequestOption) (res *pagination.Cursor[ResourceGroupResponse], err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s", params.OrgID, type_)
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

// Retrieve a list of ResourceGroup entities
func (r *ResourceGroupService) ListAutoPaging(ctx context.Context, type_ string, params ResourceGroupListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ResourceGroupResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, type_, params, opts...))
}

// Delete a ResourceGroup for the UUID
func (r *ResourceGroupService) Delete(ctx context.Context, type_ string, id string, body ResourceGroupDeleteParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.OrgID, precfg.OrgID)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s", body.OrgID, type_, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add an item to a ResourceGroup.
func (r *ResourceGroupService) AddResource(ctx context.Context, type_ string, resourceGroupID string, params ResourceGroupAddResourceParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resourceGroupId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s/addresource", params.OrgID, type_, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a list of items for a ResourceGroup
func (r *ResourceGroupService) ListContents(ctx context.Context, type_ string, resourceGroupID string, params ResourceGroupListContentsParams, opts ...option.RequestOption) (res *pagination.Cursor[ResourceGroupListContentsResponse], err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resourceGroupId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s/contents", params.OrgID, type_, resourceGroupID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Retrieve a list of items for a ResourceGroup
func (r *ResourceGroupService) ListContentsAutoPaging(ctx context.Context, type_ string, resourceGroupID string, params ResourceGroupListContentsParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ResourceGroupListContentsResponse] {
	return pagination.NewCursorAutoPager(r.ListContents(ctx, type_, resourceGroupID, params, opts...))
}

// Retrieve a list of permission policies for a ResourceGroup
func (r *ResourceGroupService) ListPermissions(ctx context.Context, type_ string, resourceGroupID string, params ResourceGroupListPermissionsParams, opts ...option.RequestOption) (res *pagination.Cursor[PermissionPolicyResponse], err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resourceGroupId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s/permissions", params.OrgID, type_, resourceGroupID)
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

// Retrieve a list of permission policies for a ResourceGroup
func (r *ResourceGroupService) ListPermissionsAutoPaging(ctx context.Context, type_ string, resourceGroupID string, params ResourceGroupListPermissionsParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PermissionPolicyResponse] {
	return pagination.NewCursorAutoPager(r.ListPermissions(ctx, type_, resourceGroupID, params, opts...))
}

// Remove an item from a ResourceGroup.
func (r *ResourceGroupService) RemoveResource(ctx context.Context, type_ string, resourceGroupID string, params ResourceGroupRemoveResourceParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
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
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resourceGroupId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/resourcegroups/%s/%s/removeresource", params.OrgID, type_, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ResourceGroupResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The unique identifier (UUID) of the user who created this Resource Group.
	CreatedBy string `json:"createdBy"`
	// The date and time _(in ISO-8601 format)_ when the Resource Group was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the Resource Group was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this Resource Group.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Resource Group.
	Name string `json:"name"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                     `json:"version"`
	JSON    resourceGroupResponseJSON `json:"-"`
}

// resourceGroupResponseJSON contains the JSON metadata for the struct
// [ResourceGroupResponse]
type resourceGroupResponseJSON struct {
	ID             apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupListContentsResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The id of the user who created this item for the resource group.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the item was created for the resource group.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the resource group item was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this item for the resource group.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The UUID of the item.
	TargetID   string                                      `json:"targetId"`
	TargetType ResourceGroupListContentsResponseTargetType `json:"targetType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                 `json:"version"`
	JSON    resourceGroupListContentsResponseJSON `json:"-"`
}

// resourceGroupListContentsResponseJSON contains the JSON metadata for the struct
// [ResourceGroupListContentsResponse]
type resourceGroupListContentsResponseJSON struct {
	ID             apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	TargetID       apijson.Field
	TargetType     apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ResourceGroupListContentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceGroupListContentsResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceGroupListContentsResponseTargetType string

const (
	ResourceGroupListContentsResponseTargetTypeItem  ResourceGroupListContentsResponseTargetType = "ITEM"
	ResourceGroupListContentsResponseTargetTypeGroup ResourceGroupListContentsResponseTargetType = "GROUP"
)

func (r ResourceGroupListContentsResponseTargetType) IsKnown() bool {
	switch r {
	case ResourceGroupListContentsResponseTargetTypeItem, ResourceGroupListContentsResponseTargetTypeGroup:
		return true
	}
	return false
}

type ResourceGroupNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID   param.Field[string] `path:"orgId,required"`
	Name    param.Field[string] `json:"name,required"`
	Version param.Field[int64]  `json:"version"`
}

func (r ResourceGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceGroupGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type ResourceGroupUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID   param.Field[string] `path:"orgId,required"`
	Name    param.Field[string] `json:"name,required"`
	Version param.Field[int64]  `json:"version"`
}

func (r ResourceGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResourceGroupListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of ResourceGroups to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ResourceGroupListParams]'s query parameters as
// `url.Values`.
func (r ResourceGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ResourceGroupDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type ResourceGroupAddResourceParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The id of the item or group you want to:
	//
	// - _Add Item_ call: add to a Resource Group.
	// - _Remove Item_ call: remove from the Resource Group.
	TargetID param.Field[string] `json:"targetId,required"`
	// When adding to or removing from a Resource Group, specify whether a single item
	// or group:
	//
	// - `item`
	//   - _Add Item_ call: use to add a single meter to a Resource Group
	//   - _Remove Item_ call: use to remove a single from a Resource Group.
	//
	// - `group`
	//   - _Add Item_ call: use to add a Resource Group to another Resource Group and
	//     form a nested Resource Group
	//   - _Remove Item_ call: use remove a nested Resource Group from a Resource
	//     Group.
	TargetType param.Field[ResourceGroupAddResourceParamsTargetType] `json:"targetType,required"`
	// The version number of the group.
	Version param.Field[int64] `json:"version"`
}

func (r ResourceGroupAddResourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When adding to or removing from a Resource Group, specify whether a single item
// or group:
//
// - `item`
//   - _Add Item_ call: use to add a single meter to a Resource Group
//   - _Remove Item_ call: use to remove a single from a Resource Group.
//
// - `group`
//   - _Add Item_ call: use to add a Resource Group to another Resource Group and
//     form a nested Resource Group
//   - _Remove Item_ call: use remove a nested Resource Group from a Resource
//     Group.
type ResourceGroupAddResourceParamsTargetType string

const (
	ResourceGroupAddResourceParamsTargetTypeItem  ResourceGroupAddResourceParamsTargetType = "ITEM"
	ResourceGroupAddResourceParamsTargetTypeGroup ResourceGroupAddResourceParamsTargetType = "GROUP"
)

func (r ResourceGroupAddResourceParamsTargetType) IsKnown() bool {
	switch r {
	case ResourceGroupAddResourceParamsTargetTypeItem, ResourceGroupAddResourceParamsTargetTypeGroup:
		return true
	}
	return false
}

type ResourceGroupListContentsParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of ResourceGroupItems to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ResourceGroupListContentsParams]'s query parameters as
// `url.Values`.
func (r ResourceGroupListContentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ResourceGroupListPermissionsParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of PermissionPolicy entities to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ResourceGroupListPermissionsParams]'s query parameters as
// `url.Values`.
func (r ResourceGroupListPermissionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ResourceGroupRemoveResourceParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The id of the item or group you want to:
	//
	// - _Add Item_ call: add to a Resource Group.
	// - _Remove Item_ call: remove from the Resource Group.
	TargetID param.Field[string] `json:"targetId,required"`
	// When adding to or removing from a Resource Group, specify whether a single item
	// or group:
	//
	// - `item`
	//   - _Add Item_ call: use to add a single meter to a Resource Group
	//   - _Remove Item_ call: use to remove a single from a Resource Group.
	//
	// - `group`
	//   - _Add Item_ call: use to add a Resource Group to another Resource Group and
	//     form a nested Resource Group
	//   - _Remove Item_ call: use remove a nested Resource Group from a Resource
	//     Group.
	TargetType param.Field[ResourceGroupRemoveResourceParamsTargetType] `json:"targetType,required"`
	// The version number of the group.
	Version param.Field[int64] `json:"version"`
}

func (r ResourceGroupRemoveResourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When adding to or removing from a Resource Group, specify whether a single item
// or group:
//
// - `item`
//   - _Add Item_ call: use to add a single meter to a Resource Group
//   - _Remove Item_ call: use to remove a single from a Resource Group.
//
// - `group`
//   - _Add Item_ call: use to add a Resource Group to another Resource Group and
//     form a nested Resource Group
//   - _Remove Item_ call: use remove a nested Resource Group from a Resource
//     Group.
type ResourceGroupRemoveResourceParamsTargetType string

const (
	ResourceGroupRemoveResourceParamsTargetTypeItem  ResourceGroupRemoveResourceParamsTargetType = "ITEM"
	ResourceGroupRemoveResourceParamsTargetTypeGroup ResourceGroupRemoveResourceParamsTargetType = "GROUP"
)

func (r ResourceGroupRemoveResourceParamsTargetType) IsKnown() bool {
	switch r {
	case ResourceGroupRemoveResourceParamsTargetTypeItem, ResourceGroupRemoveResourceParamsTargetTypeGroup:
		return true
	}
	return false
}
