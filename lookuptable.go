// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// LookupTableService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupTableService] method instead.
type LookupTableService struct {
	Options                 []option.RequestOption
	LookupTableRevisions    *LookupTableLookupTableRevisionService
	LookupTableRevisionData *LookupTableLookupTableRevisionDataService
}

// NewLookupTableService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLookupTableService(opts ...option.RequestOption) (r *LookupTableService) {
	r = &LookupTableService{}
	r.Options = opts
	r.LookupTableRevisions = NewLookupTableLookupTableRevisionService(opts...)
	r.LookupTableRevisionData = NewLookupTableLookupTableRevisionDataService(opts...)
	return
}

// Create a new Lookup Table.
func (r *LookupTableService) New(ctx context.Context, params LookupTableNewParams, opts ...option.RequestOption) (res *LookupTableResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Lookup Table by UUID.
func (r *LookupTableService) Get(ctx context.Context, id string, params LookupTableGetParams, opts ...option.RequestOption) (res *LookupTableResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update the Lookup Table with the given UUID.
func (r *LookupTableService) Update(ctx context.Context, id string, params LookupTableUpdateParams, opts ...option.RequestOption) (res *LookupTableResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list Lookup Tables created for the Organization:
//
//   - Returned list can be filtered by Lookup Table `code` query parameter.
//   - If you want to include any non-default fields for the returned Lookup Tables,
//     use the additional query parameter to specify which you want to include in the
//     response.
func (r *LookupTableService) List(ctx context.Context, params LookupTableListParams, opts ...option.RequestOption) (res *pagination.Cursor[LookupTableResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables", params.OrgID)
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

// Retrieve a list Lookup Tables created for the Organization:
//
//   - Returned list can be filtered by Lookup Table `code` query parameter.
//   - If you want to include any non-default fields for the returned Lookup Tables,
//     use the additional query parameter to specify which you want to include in the
//     response.
func (r *LookupTableService) ListAutoPaging(ctx context.Context, params LookupTableListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[LookupTableResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Lookup Table with the given UUID.
func (r *LookupTableService) Delete(ctx context.Context, id string, body LookupTableDeleteParams, opts ...option.RequestOption) (res *LookupTableResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LookupTableRequestParam struct {
	// Code of the Lookup Table - unique short code used to identify the Lookup Table.
	//
	// **NOTE:** Code has a maximum length of 80 characters and must not contain
	// non-printable or whitespace characters (except space), and cannot start/end with
	// whitespace.
	Code param.Field[string] `json:"code" api:"required"`
	// Descriptive name for the Lookup Table.
	Name param.Field[string] `json:"name" api:"required"`
	// User defined fields enabling you to attach custom data. The value for a custom
	// field can be either a string or a number.
	//
	// If `customFields` can also be defined for this entity at the Organizational
	// level, `customField` values defined at individual level override values of
	// `customFields` with the same name defined at Organization level.
	//
	// See
	// [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields)
	// in the m3ter documentation for more information.
	CustomFields param.Field[map[string]LookupTableRequestCustomFieldsUnionParam] `json:"customFields"`
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

func (r LookupTableRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type LookupTableRequestCustomFieldsUnionParam interface {
	ImplementsLookupTableRequestCustomFieldsUnionParam()
}

type LookupTableResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// Response containing a LookupTableRevision entity
	ActiveRevision LookupTableRevisionResponse `json:"activeRevision"`
	// The code of the Lookup Table
	Code string `json:"code"`
	// The id of the user who created this Lookup Table.
	CreatedBy string `json:"createdBy"`
	// User defined fields enabling you to attach custom data. The value for a custom
	// field can be either a string or a number.
	//
	// If `customFields` can also be defined for this entity at the Organizational
	// level,`customField` values defined at individual level override values of
	// `customFields` with the same name defined at Organization level.
	//
	// See
	// [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields)
	// in the m3ter documentation for more information.
	CustomFields map[string]LookupTableResponseCustomFieldsUnion `json:"customFields"`
	// The DateTime when the Lookup Table was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this Lookup Table.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Lookup Table
	Name string `json:"name"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                   `json:"version"`
	JSON    lookupTableResponseJSON `json:"-"`
}

// lookupTableResponseJSON contains the JSON metadata for the struct
// [LookupTableResponse]
type lookupTableResponseJSON struct {
	ID             apijson.Field
	ActiveRevision apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	CustomFields   apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type LookupTableResponseCustomFieldsUnion interface {
	ImplementsLookupTableResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LookupTableResponseCustomFieldsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type LookupTableNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID              param.Field[string]     `path:"orgId" api:"required"`
	LookupTableRequest LookupTableRequestParam `json:"lookup_table_request" api:"required"`
}

func (r LookupTableNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LookupTableRequest)
}

type LookupTableGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Comma separated list of additional non-default fields to be included in the
	// response. For example,if you want to include the active Revision for the Lookup
	// Tables returned, set `additional=activeRevision` in the query.
	Additional param.Field[[]string] `query:"additional"`
}

// URLQuery serializes [LookupTableGetParams]'s query parameters as `url.Values`.
func (r LookupTableGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupTableUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID              param.Field[string]     `path:"orgId" api:"required"`
	LookupTableRequest LookupTableRequestParam `json:"lookup_table_request" api:"required"`
}

func (r LookupTableUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LookupTableRequest)
}

type LookupTableListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Comma separated list of additional non-default fields to be included in the
	// response. For example,if you want to include the active Revision for each of the
	// Lookup Tables returned, set `additional=activeRevision` in the query.
	Additional param.Field[[]string] `query:"additional"`
	// List of Lookup Table codes to retrieve.
	Codes param.Field[[]string] `query:"codes"`
	// Token to supply for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Lookup Tables to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [LookupTableListParams]'s query parameters as `url.Values`.
func (r LookupTableListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupTableDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
