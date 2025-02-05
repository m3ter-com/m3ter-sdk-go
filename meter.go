// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
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
func (r *MeterService) New(ctx context.Context, orgID string, body MeterNewParams, opts ...option.RequestOption) (res *Meter, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/meters", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve the Meter with the given UUID.
func (r *MeterService) Get(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Meter, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/meters/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Meter with the given UUID.
//
// **Note:** If you have created Custom Fields for a Meter, when you use this
// endpoint to update the Meter use the `customFields` parameter to preserve those
// Custom Fields. If you omit them from the update request, they will be lost.
func (r *MeterService) Update(ctx context.Context, orgID string, id string, body MeterUpdateParams, opts ...option.RequestOption) (res *Meter, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/meters/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of Meter entities
func (r *MeterService) List(ctx context.Context, orgID string, query MeterListParams, opts ...option.RequestOption) (res *pagination.Cursor[Meter], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/meters", orgID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieve a list of Meter entities
func (r *MeterService) ListAutoPaging(ctx context.Context, orgID string, query MeterListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Meter] {
	return pagination.NewCursorAutoPager(r.List(ctx, orgID, query, opts...))
}

type Meter struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
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
	CustomFields map[string]MeterCustomFieldsUnion `json:"customFields"`
	// Used to submit categorized raw usage data values for ingest into the platform -
	// either numeric quantitative values or non-numeric data values. At least one
	// required per Meter; maximum 15 per Meter.
	DataFields []MeterDataField `json:"dataFields"`
	// Used to submit usage data values for ingest into the platform that are the
	// result of a calculation performed on `dataFields`, `customFields`, or system
	// `Timestamp` fields. Raw usage data is not submitted using `derivedFields`.
	// Maximum 15 per Meter. _(Optional)_.
	DerivedFields []MeterDerivedField `json:"derivedFields"`
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
	ProductID string    `json:"productId"`
	JSON      meterJSON `json:"-"`
}

// meterJSON contains the JSON metadata for the struct [Meter]
type meterJSON struct {
	ID             apijson.Field
	Version        apijson.Field
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
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Meter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type MeterCustomFieldsUnion interface {
	ImplementsMeterCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MeterCustomFieldsUnion)(nil)).Elem(),
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

type MeterDataField struct {
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category MeterDataFieldsCategory `json:"category,required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code string `json:"code,required"`
	// Descriptive name of the field.
	Name string `json:"name,required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit string             `json:"unit"`
	JSON meterDataFieldJSON `json:"-"`
}

// meterDataFieldJSON contains the JSON metadata for the struct [MeterDataField]
type meterDataFieldJSON struct {
	Category    apijson.Field
	Code        apijson.Field
	Name        apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterDataField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterDataFieldJSON) RawJSON() string {
	return r.raw
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type MeterDataFieldsCategory string

const (
	MeterDataFieldsCategoryWho      MeterDataFieldsCategory = "WHO"
	MeterDataFieldsCategoryWhere    MeterDataFieldsCategory = "WHERE"
	MeterDataFieldsCategoryWhat     MeterDataFieldsCategory = "WHAT"
	MeterDataFieldsCategoryOther    MeterDataFieldsCategory = "OTHER"
	MeterDataFieldsCategoryMetadata MeterDataFieldsCategory = "METADATA"
	MeterDataFieldsCategoryMeasure  MeterDataFieldsCategory = "MEASURE"
	MeterDataFieldsCategoryIncome   MeterDataFieldsCategory = "INCOME"
	MeterDataFieldsCategoryCost     MeterDataFieldsCategory = "COST"
)

func (r MeterDataFieldsCategory) IsKnown() bool {
	switch r {
	case MeterDataFieldsCategoryWho, MeterDataFieldsCategoryWhere, MeterDataFieldsCategoryWhat, MeterDataFieldsCategoryOther, MeterDataFieldsCategoryMetadata, MeterDataFieldsCategoryMeasure, MeterDataFieldsCategoryIncome, MeterDataFieldsCategoryCost:
		return true
	}
	return false
}

type MeterDerivedField struct {
	// The calculation used to transform the value of submitted `dataFields` in usage
	// data. Calculation can reference `dataFields`, `customFields`, or system
	// `Timestamp` fields. _(Example: datafieldms datafieldgb)_
	Calculation string `json:"calculation,required"`
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category MeterDerivedFieldsCategory `json:"category,required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code string `json:"code,required"`
	// Descriptive name of the field.
	Name string `json:"name,required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit string                `json:"unit"`
	JSON meterDerivedFieldJSON `json:"-"`
}

// meterDerivedFieldJSON contains the JSON metadata for the struct
// [MeterDerivedField]
type meterDerivedFieldJSON struct {
	Calculation apijson.Field
	Category    apijson.Field
	Code        apijson.Field
	Name        apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeterDerivedField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meterDerivedFieldJSON) RawJSON() string {
	return r.raw
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type MeterDerivedFieldsCategory string

const (
	MeterDerivedFieldsCategoryWho      MeterDerivedFieldsCategory = "WHO"
	MeterDerivedFieldsCategoryWhere    MeterDerivedFieldsCategory = "WHERE"
	MeterDerivedFieldsCategoryWhat     MeterDerivedFieldsCategory = "WHAT"
	MeterDerivedFieldsCategoryOther    MeterDerivedFieldsCategory = "OTHER"
	MeterDerivedFieldsCategoryMetadata MeterDerivedFieldsCategory = "METADATA"
	MeterDerivedFieldsCategoryMeasure  MeterDerivedFieldsCategory = "MEASURE"
	MeterDerivedFieldsCategoryIncome   MeterDerivedFieldsCategory = "INCOME"
	MeterDerivedFieldsCategoryCost     MeterDerivedFieldsCategory = "COST"
)

func (r MeterDerivedFieldsCategory) IsKnown() bool {
	switch r {
	case MeterDerivedFieldsCategoryWho, MeterDerivedFieldsCategoryWhere, MeterDerivedFieldsCategoryWhat, MeterDerivedFieldsCategoryOther, MeterDerivedFieldsCategoryMetadata, MeterDerivedFieldsCategoryMeasure, MeterDerivedFieldsCategoryIncome, MeterDerivedFieldsCategoryCost:
		return true
	}
	return false
}

type MeterNewParams struct {
	// Code of the Meter - unique short code used to identify the Meter.
	//
	// **NOTE:** Code has a maximum length of 80 characters and must not contain
	// non-printable or whitespace characters (except space), and cannot start/end with
	// whitespace.
	Code param.Field[string] `json:"code,required"`
	// Used to submit categorized raw usage data values for ingest into the platform -
	// either numeric quantitative values or non-numeric data values. At least one
	// required per Meter; maximum 15 per Meter.
	DataFields param.Field[[]MeterNewParamsDataField] `json:"dataFields,required"`
	// Used to submit usage data values for ingest into the platform that are the
	// result of a calculation performed on `dataFields`, `customFields`, or system
	// `Timestamp` fields. Raw usage data is not submitted using `derivedFields`.
	// Maximum 15 per Meter. _(Optional)_.
	//
	// **Note:** Required parameter. If you want to create a Meter without Derived
	// Fields, use an empty array `[]`. If you use a `null`, you'll receive an error.
	DerivedFields param.Field[[]MeterNewParamsDerivedField] `json:"derivedFields,required"`
	// Descriptive name for the Meter.
	Name param.Field[string] `json:"name,required"`
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

type MeterNewParamsDataField struct {
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category param.Field[MeterNewParamsDataFieldsCategory] `json:"category,required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name of the field.
	Name param.Field[string] `json:"name,required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit param.Field[string] `json:"unit"`
}

func (r MeterNewParamsDataField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type MeterNewParamsDataFieldsCategory string

const (
	MeterNewParamsDataFieldsCategoryWho      MeterNewParamsDataFieldsCategory = "WHO"
	MeterNewParamsDataFieldsCategoryWhere    MeterNewParamsDataFieldsCategory = "WHERE"
	MeterNewParamsDataFieldsCategoryWhat     MeterNewParamsDataFieldsCategory = "WHAT"
	MeterNewParamsDataFieldsCategoryOther    MeterNewParamsDataFieldsCategory = "OTHER"
	MeterNewParamsDataFieldsCategoryMetadata MeterNewParamsDataFieldsCategory = "METADATA"
	MeterNewParamsDataFieldsCategoryMeasure  MeterNewParamsDataFieldsCategory = "MEASURE"
	MeterNewParamsDataFieldsCategoryIncome   MeterNewParamsDataFieldsCategory = "INCOME"
	MeterNewParamsDataFieldsCategoryCost     MeterNewParamsDataFieldsCategory = "COST"
)

func (r MeterNewParamsDataFieldsCategory) IsKnown() bool {
	switch r {
	case MeterNewParamsDataFieldsCategoryWho, MeterNewParamsDataFieldsCategoryWhere, MeterNewParamsDataFieldsCategoryWhat, MeterNewParamsDataFieldsCategoryOther, MeterNewParamsDataFieldsCategoryMetadata, MeterNewParamsDataFieldsCategoryMeasure, MeterNewParamsDataFieldsCategoryIncome, MeterNewParamsDataFieldsCategoryCost:
		return true
	}
	return false
}

type MeterNewParamsDerivedField struct {
	// The calculation used to transform the value of submitted `dataFields` in usage
	// data. Calculation can reference `dataFields`, `customFields`, or system
	// `Timestamp` fields. _(Example: datafieldms datafieldgb)_
	Calculation param.Field[string] `json:"calculation,required"`
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category param.Field[MeterNewParamsDerivedFieldsCategory] `json:"category,required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name of the field.
	Name param.Field[string] `json:"name,required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit param.Field[string] `json:"unit"`
}

func (r MeterNewParamsDerivedField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type MeterNewParamsDerivedFieldsCategory string

const (
	MeterNewParamsDerivedFieldsCategoryWho      MeterNewParamsDerivedFieldsCategory = "WHO"
	MeterNewParamsDerivedFieldsCategoryWhere    MeterNewParamsDerivedFieldsCategory = "WHERE"
	MeterNewParamsDerivedFieldsCategoryWhat     MeterNewParamsDerivedFieldsCategory = "WHAT"
	MeterNewParamsDerivedFieldsCategoryOther    MeterNewParamsDerivedFieldsCategory = "OTHER"
	MeterNewParamsDerivedFieldsCategoryMetadata MeterNewParamsDerivedFieldsCategory = "METADATA"
	MeterNewParamsDerivedFieldsCategoryMeasure  MeterNewParamsDerivedFieldsCategory = "MEASURE"
	MeterNewParamsDerivedFieldsCategoryIncome   MeterNewParamsDerivedFieldsCategory = "INCOME"
	MeterNewParamsDerivedFieldsCategoryCost     MeterNewParamsDerivedFieldsCategory = "COST"
)

func (r MeterNewParamsDerivedFieldsCategory) IsKnown() bool {
	switch r {
	case MeterNewParamsDerivedFieldsCategoryWho, MeterNewParamsDerivedFieldsCategoryWhere, MeterNewParamsDerivedFieldsCategoryWhat, MeterNewParamsDerivedFieldsCategoryOther, MeterNewParamsDerivedFieldsCategoryMetadata, MeterNewParamsDerivedFieldsCategoryMeasure, MeterNewParamsDerivedFieldsCategoryIncome, MeterNewParamsDerivedFieldsCategoryCost:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type MeterNewParamsCustomFieldsUnion interface {
	ImplementsMeterNewParamsCustomFieldsUnion()
}

type MeterUpdateParams struct {
	// Code of the Meter - unique short code used to identify the Meter.
	//
	// **NOTE:** Code has a maximum length of 80 characters and must not contain
	// non-printable or whitespace characters (except space), and cannot start/end with
	// whitespace.
	Code param.Field[string] `json:"code,required"`
	// Used to submit categorized raw usage data values for ingest into the platform -
	// either numeric quantitative values or non-numeric data values. At least one
	// required per Meter; maximum 15 per Meter.
	DataFields param.Field[[]MeterUpdateParamsDataField] `json:"dataFields,required"`
	// Used to submit usage data values for ingest into the platform that are the
	// result of a calculation performed on `dataFields`, `customFields`, or system
	// `Timestamp` fields. Raw usage data is not submitted using `derivedFields`.
	// Maximum 15 per Meter. _(Optional)_.
	//
	// **Note:** Required parameter. If you want to create a Meter without Derived
	// Fields, use an empty array `[]`. If you use a `null`, you'll receive an error.
	DerivedFields param.Field[[]MeterUpdateParamsDerivedField] `json:"derivedFields,required"`
	// Descriptive name for the Meter.
	Name param.Field[string] `json:"name,required"`
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

type MeterUpdateParamsDataField struct {
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category param.Field[MeterUpdateParamsDataFieldsCategory] `json:"category,required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name of the field.
	Name param.Field[string] `json:"name,required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit param.Field[string] `json:"unit"`
}

func (r MeterUpdateParamsDataField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type MeterUpdateParamsDataFieldsCategory string

const (
	MeterUpdateParamsDataFieldsCategoryWho      MeterUpdateParamsDataFieldsCategory = "WHO"
	MeterUpdateParamsDataFieldsCategoryWhere    MeterUpdateParamsDataFieldsCategory = "WHERE"
	MeterUpdateParamsDataFieldsCategoryWhat     MeterUpdateParamsDataFieldsCategory = "WHAT"
	MeterUpdateParamsDataFieldsCategoryOther    MeterUpdateParamsDataFieldsCategory = "OTHER"
	MeterUpdateParamsDataFieldsCategoryMetadata MeterUpdateParamsDataFieldsCategory = "METADATA"
	MeterUpdateParamsDataFieldsCategoryMeasure  MeterUpdateParamsDataFieldsCategory = "MEASURE"
	MeterUpdateParamsDataFieldsCategoryIncome   MeterUpdateParamsDataFieldsCategory = "INCOME"
	MeterUpdateParamsDataFieldsCategoryCost     MeterUpdateParamsDataFieldsCategory = "COST"
)

func (r MeterUpdateParamsDataFieldsCategory) IsKnown() bool {
	switch r {
	case MeterUpdateParamsDataFieldsCategoryWho, MeterUpdateParamsDataFieldsCategoryWhere, MeterUpdateParamsDataFieldsCategoryWhat, MeterUpdateParamsDataFieldsCategoryOther, MeterUpdateParamsDataFieldsCategoryMetadata, MeterUpdateParamsDataFieldsCategoryMeasure, MeterUpdateParamsDataFieldsCategoryIncome, MeterUpdateParamsDataFieldsCategoryCost:
		return true
	}
	return false
}

type MeterUpdateParamsDerivedField struct {
	// The calculation used to transform the value of submitted `dataFields` in usage
	// data. Calculation can reference `dataFields`, `customFields`, or system
	// `Timestamp` fields. _(Example: datafieldms datafieldgb)_
	Calculation param.Field[string] `json:"calculation,required"`
	// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
	Category param.Field[MeterUpdateParamsDerivedFieldsCategory] `json:"category,required"`
	// Short code to identify the field
	//
	// **NOTE:** Code has a maximum length of 80 characters and can only contain
	// letters, numbers, underscore, and the dollar character, and must not start with
	// a number.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name of the field.
	Name param.Field[string] `json:"name,required"`
	// The units to measure the data with. Should conform to _Unified Code for Units of
	// Measure_ (UCUM). Required only for numeric field categories.
	Unit param.Field[string] `json:"unit"`
}

func (r MeterUpdateParamsDerivedField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of field (WHO, WHAT, WHERE, MEASURE, METADATA, INCOME, COST, OTHER).
type MeterUpdateParamsDerivedFieldsCategory string

const (
	MeterUpdateParamsDerivedFieldsCategoryWho      MeterUpdateParamsDerivedFieldsCategory = "WHO"
	MeterUpdateParamsDerivedFieldsCategoryWhere    MeterUpdateParamsDerivedFieldsCategory = "WHERE"
	MeterUpdateParamsDerivedFieldsCategoryWhat     MeterUpdateParamsDerivedFieldsCategory = "WHAT"
	MeterUpdateParamsDerivedFieldsCategoryOther    MeterUpdateParamsDerivedFieldsCategory = "OTHER"
	MeterUpdateParamsDerivedFieldsCategoryMetadata MeterUpdateParamsDerivedFieldsCategory = "METADATA"
	MeterUpdateParamsDerivedFieldsCategoryMeasure  MeterUpdateParamsDerivedFieldsCategory = "MEASURE"
	MeterUpdateParamsDerivedFieldsCategoryIncome   MeterUpdateParamsDerivedFieldsCategory = "INCOME"
	MeterUpdateParamsDerivedFieldsCategoryCost     MeterUpdateParamsDerivedFieldsCategory = "COST"
)

func (r MeterUpdateParamsDerivedFieldsCategory) IsKnown() bool {
	switch r {
	case MeterUpdateParamsDerivedFieldsCategoryWho, MeterUpdateParamsDerivedFieldsCategoryWhere, MeterUpdateParamsDerivedFieldsCategoryWhat, MeterUpdateParamsDerivedFieldsCategoryOther, MeterUpdateParamsDerivedFieldsCategoryMetadata, MeterUpdateParamsDerivedFieldsCategoryMeasure, MeterUpdateParamsDerivedFieldsCategoryIncome, MeterUpdateParamsDerivedFieldsCategoryCost:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type MeterUpdateParamsCustomFieldsUnion interface {
	ImplementsMeterUpdateParamsCustomFieldsUnion()
}

type MeterListParams struct {
	// list of codes to retrieve
	Codes param.Field[[]string] `query:"codes"`
	// list of ids to retrieve
	IDs param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Meters to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// The UUIDs of the products to retrieve meters for
	ProductID param.Field[[]string] `query:"productId"`
}

// URLQuery serializes [MeterListParams]'s query parameters as `url.Values`.
func (r MeterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
