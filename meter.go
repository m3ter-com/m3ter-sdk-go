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

// MeterService contains methods and other services that help with interacting with
// the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeterService] method instead.
type MeterService struct {
	Options []option.RequestOption
}

// NewMeterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeterService(opts ...option.RequestOption) (r *MeterService) {
	r = &MeterService{}
	r.Options = opts
	return
}

// Create a new Meter.
//
// When you create a Meter, you can define two types of field for usage data
// collection and ingest into the platform:
//
//   - `dataFields` to collect raw usage data measures - numeric quantitative data
//     values or non-numeric point data values.
//   - `derivedFields` to derive usage data measures that are the result of applying
//     a calculation to `dataFields`, `customFields`, or system `Timestamp` fields.
//
// You can also:
//
//   - Create `customFields` for a Meter, which allows you to attach custom data to
//     the Meter as name/value pairs.
//   - Create Global Meters, which are not tied to a specific Product and allow you
//     collect to usage data that will form the basis of usage-based pricing across
//     more than one of your Products.
//
// **IMPORTANT! - use of PII:** The use of any of your end-customers' Personally
// Identifiable Information (PII) in m3ter is restricted to a few fields on the
// **Account** entity. Please ensure that any fields you configure for Meters, such
// as Data Fields or Derived Fields, do not contain any end-customer PII data. See
// the [Introduction section](https://www.m3ter.com/docs/api#section/Introduction)
// above for more details.
//
// See also:
//
// - [Reviewing Meter Options](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/reviewing-meter-options).
func (r *MeterService) New(ctx context.Context, params MeterNewParams, opts ...option.RequestOption) (res *MeterResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/meters", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Meter with the given UUID.
func (r *MeterService) Get(ctx context.Context, id string, query MeterGetParams, opts ...option.RequestOption) (res *MeterResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/meters/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Meter with the given UUID.
//
// **Note:** If you have created Custom Fields for a Meter, when you use this
// endpoint to update the Meter use the `customFields` parameter to preserve those
// Custom Fields. If you omit them from the update request, they will be lost.
func (r *MeterService) Update(ctx context.Context, id string, params MeterUpdateParams, opts ...option.RequestOption) (res *MeterResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/meters/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Meters that can be filtered by Product, Meter ID, or Meter
// short code.
func (r *MeterService) List(ctx context.Context, params MeterListParams, opts ...option.RequestOption) (res *pagination.Cursor[MeterResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/meters", params.OrgID)
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

// Retrieve a list of Meters that can be filtered by Product, Meter ID, or Meter
// short code.
func (r *MeterService) ListAutoPaging(ctx context.Context, params MeterListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[MeterResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Meter with the given UUID.
func (r *MeterService) Delete(ctx context.Context, id string, body MeterDeleteParams, opts ...option.RequestOption) (res *MeterResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/meters/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DataField struct {
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category DataFieldCategory `json:"category" api:"required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code string `json:"code" api:"required"`
	// Descriptive name of the field.
	Name string `json:"name" api:"required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit string        `json:"unit"`
	JSON dataFieldJSON `json:"-"`
}

// dataFieldJSON contains the JSON metadata for the struct [DataField]
type dataFieldJSON struct {
	Category    apijson.Field
	Code        apijson.Field
	Name        apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataFieldJSON) RawJSON() string {
	return r.raw
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type DataFieldCategory string

const (
	DataFieldCategoryWho      DataFieldCategory = "WHO"
	DataFieldCategoryWhere    DataFieldCategory = "WHERE"
	DataFieldCategoryWhat     DataFieldCategory = "WHAT"
	DataFieldCategoryOther    DataFieldCategory = "OTHER"
	DataFieldCategoryMetadata DataFieldCategory = "METADATA"
	DataFieldCategoryMeasure  DataFieldCategory = "MEASURE"
	DataFieldCategoryIncome   DataFieldCategory = "INCOME"
	DataFieldCategoryCost     DataFieldCategory = "COST"
)

func (r DataFieldCategory) IsKnown() bool {
	switch r {
	case DataFieldCategoryWho, DataFieldCategoryWhere, DataFieldCategoryWhat, DataFieldCategoryOther, DataFieldCategoryMetadata, DataFieldCategoryMeasure, DataFieldCategoryIncome, DataFieldCategoryCost:
		return true
	}
	return false
}

type DataFieldParam struct {
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category param.Field[DataFieldCategory] `json:"category" api:"required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code param.Field[string] `json:"code" api:"required"`
	// Descriptive name of the field.
	Name param.Field[string] `json:"name" api:"required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit param.Field[string] `json:"unit"`
}

func (r DataFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DerivedField struct {
	// The calculation used to transform the value of submitted `dataFields` in usage
	// data. Calculation can reference `dataFields`, `customFields`, or system
	// `Timestamp` fields. _(Example: datafieldms datafieldgb)_
	Calculation string           `json:"calculation" api:"required"`
	JSON        derivedFieldJSON `json:"-"`
	DataField
}

// derivedFieldJSON contains the JSON metadata for the struct [DerivedField]
type derivedFieldJSON struct {
	Calculation apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DerivedField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r derivedFieldJSON) RawJSON() string {
	return r.raw
}

type DerivedFieldParam struct {
	// The calculation used to transform the value of submitted `dataFields` in usage
	// data. Calculation can reference `dataFields`, `customFields`, or system
	// `Timestamp` fields. _(Example: datafieldms datafieldgb)_
	Calculation param.Field[string] `json:"calculation" api:"required"`
	DataFieldParam
}

func (r DerivedFieldParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MeterResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// Code of the Meter - unique short code used to identify the Meter.
	Code string `json:"code"`
	// The id of the user who created this meter.
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
	CustomFields map[string]MeterResponseCustomFieldsUnion `json:"customFields"`
	// Used to submit categorized raw usage data values for ingest into the platform -
	// either numeric quantitative values or non-numeric data values. At least one
	// required per Meter; maximum 15 per Meter.
	DataFields []DataField `json:"dataFields"`
	// Used to submit usage data values for ingest into the platform that are the
	// result of a calculation performed on `dataFields`, `customFields`, or system
	// `Timestamp` fields. Raw usage data is not submitted using `derivedFields`.
	// Maximum 15 per Meter. _(Optional)_.
	DerivedFields []DerivedField `json:"derivedFields"`
	// The DateTime when the meter was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the meter was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// UUID of the MeterGroup the Meter belongs to. _(Optional)_.
	GroupID string `json:"groupId"`
	// The id of the user who last modified this meter.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Descriptive name for the Meter.
	Name string `json:"name"`
	// UUID of the Product the Meter belongs to. _(Optional)_ - if blank, the Meter is
	// global.
	ProductID string `json:"productId"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64             `json:"version"`
	JSON    meterResponseJSON `json:"-"`
}

// meterResponseJSON contains the JSON metadata for the struct [MeterResponse]
type meterResponseJSON struct {
	ID             apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	CustomFields   apijson.Field
	DataFields     apijson.Field
	DerivedFields  apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	GroupID        apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	ProductID      apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MeterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type MeterResponseCustomFieldsUnion interface {
	ImplementsMeterResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterResponseCustomFieldsUnion)(nil)).Elem(),
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

type MeterNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Code of the Meter - unique short code used to identify the Meter.
	//
	// **NOTE:** Code has a maximum length of 80 characters and must not contain
	// non-printable or whitespace characters (except space), and cannot start/end with
	// whitespace.
	Code param.Field[string] `json:"code" api:"required"`
	// Used to submit categorized raw usage data values for ingest into the platform -
	// either numeric quantitative values or non-numeric data values. At least one
	// required per Meter; maximum 15 per Meter.
	DataFields param.Field[[]DataFieldParam] `json:"dataFields" api:"required"`
	// Used to submit usage data values for ingest into the platform that are the
	// result of a calculation performed on `dataFields`, `customFields`, or system
	// `Timestamp` fields. Raw usage data is not submitted using `derivedFields`.
	// Maximum 15 per Meter. _(Optional)_.
	//
	// **Note:** Required parameter. If you want to create a Meter without Derived
	// Fields, use an empty array `[]`. If you use a `null`, you'll receive an error.
	DerivedFields param.Field[[]DerivedFieldParam] `json:"derivedFields" api:"required"`
	// Descriptive name for the Meter.
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
	CustomFields param.Field[map[string]MeterNewParamsCustomFieldsUnion] `json:"customFields"`
	// UUID of the group the Meter belongs to. _(Optional)_.
	GroupID param.Field[string] `json:"groupId"`
	// UUID of the product the Meter belongs to. _(Optional)_ - if left blank, the
	// Meter is global.
	ProductID param.Field[string] `json:"productId"`
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

func (r MeterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type MeterNewParamsCustomFieldsUnion interface {
	ImplementsMeterNewParamsCustomFieldsUnion()
}

type MeterGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type MeterUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Code of the Meter - unique short code used to identify the Meter.
	//
	// **NOTE:** Code has a maximum length of 80 characters and must not contain
	// non-printable or whitespace characters (except space), and cannot start/end with
	// whitespace.
	Code param.Field[string] `json:"code" api:"required"`
	// Used to submit categorized raw usage data values for ingest into the platform -
	// either numeric quantitative values or non-numeric data values. At least one
	// required per Meter; maximum 15 per Meter.
	DataFields param.Field[[]DataFieldParam] `json:"dataFields" api:"required"`
	// Used to submit usage data values for ingest into the platform that are the
	// result of a calculation performed on `dataFields`, `customFields`, or system
	// `Timestamp` fields. Raw usage data is not submitted using `derivedFields`.
	// Maximum 15 per Meter. _(Optional)_.
	//
	// **Note:** Required parameter. If you want to create a Meter without Derived
	// Fields, use an empty array `[]`. If you use a `null`, you'll receive an error.
	DerivedFields param.Field[[]DerivedFieldParam] `json:"derivedFields" api:"required"`
	// Descriptive name for the Meter.
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
	CustomFields param.Field[map[string]MeterUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// UUID of the group the Meter belongs to. _(Optional)_.
	GroupID param.Field[string] `json:"groupId"`
	// UUID of the product the Meter belongs to. _(Optional)_ - if left blank, the
	// Meter is global.
	ProductID param.Field[string] `json:"productId"`
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

func (r MeterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type MeterUpdateParamsCustomFieldsUnion interface {
	ImplementsMeterUpdateParamsCustomFieldsUnion()
}

type MeterListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// List of Meter codes to retrieve. These are the unique short codes that identify
	// each Meter.
	Codes param.Field[[]string] `query:"codes"`
	// List of Meter IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Meters to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// The UUIDs of the Products to retrieve Meters for.
	ProductID param.Field[[]string] `query:"productId"`
}

// URLQuery serializes [MeterListParams]'s query parameters as `url.Values`.
func (r MeterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeterDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
