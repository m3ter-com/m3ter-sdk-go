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

// Endpoints for creating/updating/deleting Lookup Table Revisions.
//
// Lookup Tables utilize a "Table and Revision" model, which lets you update data
// cleanly, and extend the schema without modifying existing calculations:
//
//   - Create Revisions for a Lookup Table, which you can use to define data schema
//     and lookup keys.
//   - Populate draft Revisions with data values. You can create and edit multiple
//     draft Revisions, but only one can be published at any given time.
//   - Publish a Revision to activate it. When you use Lookup functions in your
//     calculations that reference the Lookup Table, the data values defined for the
//     published Revision are used.
//   - When you want different, updated data values to be used, publish the draft
//     Revision containing the required new values. The currently published Revision
//     is archived automatically.
//
// **Beta Version!** The Lookup Table feature is currently available in Beta
// release version. See
// [Feature Release Stages](https://www.m3ter.com/docs/guides/getting-started/feature-release-stages)
// for Beta release definition. Lookup Table Revision endpoints will only be
// available if Lookup Tables have been enabled for your Organization. For more
// details see
// [Lookup Tables (Beta)](https://www.m3ter.com/docs/guides/lookup-tables) in our
// main User documentation.
//
// LookupTableLookupTableRevisionService contains methods and other services that
// help with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupTableLookupTableRevisionService] method instead.
type LookupTableLookupTableRevisionService struct {
	Options []option.RequestOption
}

// NewLookupTableLookupTableRevisionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLookupTableLookupTableRevisionService(opts ...option.RequestOption) (r *LookupTableLookupTableRevisionService) {
	r = &LookupTableLookupTableRevisionService{}
	r.Options = opts
	return
}

// Create a new Revision for a Lookup Table.
//
// Fields and Keys for Revision schema: Use the `fields` parameter to define a
// Revision schema containing up to 10 number or string fields. Use the `keys`
// parameter to specify which are the key fields:
//
//   - At least one field must be a non-key field and at least one a key field.
//   - Up to 5 key fields can be defined.
//   - Using multiple key fields: ensure that the order in which they are defined
//     matches the order in which you want to use them in any Lookup functions that
//     reference the Revision's Lookup Table, because this is the order in which they
//     will be passed into the function. The order of non-key fields is not
//     constrained in this way.
//
// Revision status: when you first create a Lookup Table Revision it has DRAFT
// status. Use the
// [Update LookupTableRevision Status](www.m3ter.com/docs/api#tag/LookupTableRevision/operation/UpdateLookupTableRevisionStatus)
// call to change a Revision's status.
func (r *LookupTableLookupTableRevisionService) New(ctx context.Context, lookupTableID string, params LookupTableLookupTableRevisionNewParams, opts ...option.RequestOption) (res *LookupTableRevisionResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions", params.OrgID, lookupTableID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Lookup Table Revision for the given UUID.
func (r *LookupTableLookupTableRevisionService) Get(ctx context.Context, lookupTableID string, id string, query LookupTableLookupTableRevisionGetParams, opts ...option.RequestOption) (res *LookupTableRevisionResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s", query.OrgID, lookupTableID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Lookup Table Revision for the given UUID.
//
// **NOTE:** If you've already added data to a Lookup Table Revision - see the
// following
// [Lookup Table Revision Data](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData)
// section - then you won't be able to change the Revision's field schema and
// you'll receive an error if you try do this. Create a new Revision instead, or
// delete the data items first.
func (r *LookupTableLookupTableRevisionService) Update(ctx context.Context, lookupTableID string, id string, params LookupTableLookupTableRevisionUpdateParams, opts ...option.RequestOption) (res *LookupTableRevisionResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s", params.OrgID, lookupTableID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List LookupTableRevision entities
func (r *LookupTableLookupTableRevisionService) List(ctx context.Context, lookupTableID string, params LookupTableLookupTableRevisionListParams, opts ...option.RequestOption) (res *pagination.Cursor[LookupTableRevisionResponse], err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions", params.OrgID, lookupTableID)
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

// List LookupTableRevision entities
func (r *LookupTableLookupTableRevisionService) ListAutoPaging(ctx context.Context, lookupTableID string, params LookupTableLookupTableRevisionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[LookupTableRevisionResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, lookupTableID, params, opts...))
}

// Delete the Lookup Table Revision for the given UUID.
func (r *LookupTableLookupTableRevisionService) Delete(ctx context.Context, lookupTableID string, id string, body LookupTableLookupTableRevisionDeleteParams, opts ...option.RequestOption) (res *LookupTableRevisionResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s", body.OrgID, lookupTableID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update the status of a Lookup Table Revision for the given UUID.
func (r *LookupTableLookupTableRevisionService) UpdateStatus(ctx context.Context, lookupTableID string, id string, params LookupTableLookupTableRevisionUpdateStatusParams, opts ...option.RequestOption) (res *LookupTableRevisionResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/status", params.OrgID, lookupTableID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type LookupTableRevisionRequestParam struct {
	// The list of fields of the Lookup Table Revision.
	Fields param.Field[[]LookupTableRevisionRequestFieldParam] `json:"fields" api:"required"`
	// The ordered keys of the Lookup Table Revision.
	Keys param.Field[[]string] `json:"keys" api:"required"`
	// Descriptive name for the Lookup Table Revision.
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
	CustomFields param.Field[map[string]LookupTableRevisionRequestCustomFieldsUnionParam] `json:"customFields"`
	// The start date of the Lookup Table Revision.
	StartDate param.Field[time.Time] `json:"startDate" format:"date-time"`
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

func (r LookupTableRevisionRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field of a Lookup Table Revision
type LookupTableRevisionRequestFieldParam struct {
	// Type of a Lookup Table Revision Field
	Type param.Field[LookupTableRevisionRequestFieldsType] `json:"type" api:"required"`
	// The name of the field
	Name param.Field[string] `json:"name"`
}

func (r LookupTableRevisionRequestFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of a Lookup Table Revision Field
type LookupTableRevisionRequestFieldsType string

const (
	LookupTableRevisionRequestFieldsTypeString LookupTableRevisionRequestFieldsType = "STRING"
	LookupTableRevisionRequestFieldsTypeNumber LookupTableRevisionRequestFieldsType = "NUMBER"
)

func (r LookupTableRevisionRequestFieldsType) IsKnown() bool {
	switch r {
	case LookupTableRevisionRequestFieldsTypeString, LookupTableRevisionRequestFieldsTypeNumber:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type LookupTableRevisionRequestCustomFieldsUnionParam interface {
	ImplementsLookupTableRevisionRequestCustomFieldsUnionParam()
}

type LookupTableRevisionResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// The id of the user who created the Lookup Table Revision.
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
	CustomFields map[string]LookupTableRevisionResponseCustomFieldsUnion `json:"customFields"`
	// The DateTime when the Lookup Table Revision was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table Revision was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The list of fields of the Lookup Table Revision.
	Fields    []LookupTableRevisionResponseField `json:"fields"`
	ItemCount int64                              `json:"itemCount"`
	// The ordered keys of the Lookup Table Revision
	Keys []string `json:"keys"`
	// The id of the user who last modified the Lookup Table Revision.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Lookup Table Revision.
	Name string `json:"name"`
	// The start date of the Lookup Table Revision
	StartDate time.Time `json:"startDate" format:"date-time"`
	// Status of a Lookup Table Revision
	Status LookupTableRevisionResponseStatus `json:"status"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                           `json:"version"`
	JSON    lookupTableRevisionResponseJSON `json:"-"`
}

// lookupTableRevisionResponseJSON contains the JSON metadata for the struct
// [LookupTableRevisionResponse]
type lookupTableRevisionResponseJSON struct {
	ID             apijson.Field
	CreatedBy      apijson.Field
	CustomFields   apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	Fields         apijson.Field
	ItemCount      apijson.Field
	Keys           apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	StartDate      apijson.Field
	Status         apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableRevisionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableRevisionResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type LookupTableRevisionResponseCustomFieldsUnion interface {
	ImplementsLookupTableRevisionResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LookupTableRevisionResponseCustomFieldsUnion)(nil)).Elem(),
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

// Field of a Lookup Table Revision
type LookupTableRevisionResponseField struct {
	// Type of a Lookup Table Revision Field
	Type LookupTableRevisionResponseFieldsType `json:"type" api:"required"`
	// The name of the field
	Name string                               `json:"name"`
	JSON lookupTableRevisionResponseFieldJSON `json:"-"`
}

// lookupTableRevisionResponseFieldJSON contains the JSON metadata for the struct
// [LookupTableRevisionResponseField]
type lookupTableRevisionResponseFieldJSON struct {
	Type        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupTableRevisionResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableRevisionResponseFieldJSON) RawJSON() string {
	return r.raw
}

// Type of a Lookup Table Revision Field
type LookupTableRevisionResponseFieldsType string

const (
	LookupTableRevisionResponseFieldsTypeString LookupTableRevisionResponseFieldsType = "STRING"
	LookupTableRevisionResponseFieldsTypeNumber LookupTableRevisionResponseFieldsType = "NUMBER"
)

func (r LookupTableRevisionResponseFieldsType) IsKnown() bool {
	switch r {
	case LookupTableRevisionResponseFieldsTypeString, LookupTableRevisionResponseFieldsTypeNumber:
		return true
	}
	return false
}

// Status of a Lookup Table Revision
type LookupTableRevisionResponseStatus string

const (
	LookupTableRevisionResponseStatusDraft     LookupTableRevisionResponseStatus = "DRAFT"
	LookupTableRevisionResponseStatusPublished LookupTableRevisionResponseStatus = "PUBLISHED"
	LookupTableRevisionResponseStatusArchived  LookupTableRevisionResponseStatus = "ARCHIVED"
)

func (r LookupTableRevisionResponseStatus) IsKnown() bool {
	switch r {
	case LookupTableRevisionResponseStatusDraft, LookupTableRevisionResponseStatusPublished, LookupTableRevisionResponseStatusArchived:
		return true
	}
	return false
}

// Request containing the status details for a LookupTableRevision entity
type LookupTableRevisionStatusRequestParam struct {
	// Status of a Lookup Table Revision
	Status param.Field[LookupTableRevisionStatusRequestStatus] `json:"status"`
	// The version of the LookupTableRevision.
	Version param.Field[int64] `json:"version"`
}

func (r LookupTableRevisionStatusRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of a Lookup Table Revision
type LookupTableRevisionStatusRequestStatus string

const (
	LookupTableRevisionStatusRequestStatusDraft     LookupTableRevisionStatusRequestStatus = "DRAFT"
	LookupTableRevisionStatusRequestStatusPublished LookupTableRevisionStatusRequestStatus = "PUBLISHED"
	LookupTableRevisionStatusRequestStatusArchived  LookupTableRevisionStatusRequestStatus = "ARCHIVED"
)

func (r LookupTableRevisionStatusRequestStatus) IsKnown() bool {
	switch r {
	case LookupTableRevisionStatusRequestStatusDraft, LookupTableRevisionStatusRequestStatusPublished, LookupTableRevisionStatusRequestStatusArchived:
		return true
	}
	return false
}

type LookupTableLookupTableRevisionNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Request containing a LookupTableRevision entity
	LookupTableRevisionRequest LookupTableRevisionRequestParam `json:"lookup_table_revision_request" api:"required"`
}

func (r LookupTableLookupTableRevisionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LookupTableRevisionRequest)
}

type LookupTableLookupTableRevisionGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type LookupTableLookupTableRevisionUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Request containing a LookupTableRevision entity
	LookupTableRevisionRequest LookupTableRevisionRequestParam `json:"lookup_table_revision_request" api:"required"`
}

func (r LookupTableLookupTableRevisionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LookupTableRevisionRequest)
}

type LookupTableLookupTableRevisionListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// List of Lookup Table Revision IDs to retrieve
	IDs param.Field[[]string] `query:"ids"`
	// Token to supply for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of LookupTable to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [LookupTableLookupTableRevisionListParams]'s query
// parameters as `url.Values`.
func (r LookupTableLookupTableRevisionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupTableLookupTableRevisionDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type LookupTableLookupTableRevisionUpdateStatusParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Request containing the status details for a LookupTableRevision entity
	LookupTableRevisionStatusRequest LookupTableRevisionStatusRequestParam `json:"lookup_table_revision_status_request" api:"required"`
}

func (r LookupTableLookupTableRevisionUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.LookupTableRevisionStatusRequest)
}
