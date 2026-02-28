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

// Endpoints for listing, creating, updating, retrieving, or deleting Aggregations.
//
// An Aggregation links to a Meter and targets a Data Field or Derived Field on the
// Meter. You define the method of aggregation used to convert the usage data
// collected by the targeted Meter field into a numerical unit of measurement.
//
// You can then use the unit of measurement an Aggregation yields as a metric for
// pricing Product Plans and apply usage-based pricing to your products and
// services. You might also want to aggregate raw data measures for other purposes,
// such as to feed into analytical or business performance tools.
//
// **Notes:**
//
//   - **Contrast with Compound Aggregations**. Standard or simple Aggregations of
//     this type, which apply an aggregation method directly to Meter usage data
//     fields, are contrasted with
//     [Compound Aggregations](https://www.m3ter.com/docs/api#tag/CompoundAggregation).
//     A Compound Aggregation typically references one or more simple Aggregations
//     and applies a calculation to them to derive pricing metrics needed to serve
//     more complex usage-based pricing scenarios.
//   - **Segmented Aggregations**. Segmented Aggregations allow you to segment the
//     usage data collected by a single Meter. This capability is very useful for
//     implementing some pricing and billing use cases. See
//     [Segmented Aggregations](https://www.m3ter.com/docs/guides/usage-data-aggregations/segmented-aggregations)
//     in our main documentation for more details.
//
// AggregationService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAggregationService] method instead.
type AggregationService struct {
	Options []option.RequestOption
}

// NewAggregationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAggregationService(opts ...option.RequestOption) (r *AggregationService) {
	r = &AggregationService{}
	r.Options = opts
	return
}

// Create a new Aggregation.
func (r *AggregationService) New(ctx context.Context, params AggregationNewParams, opts ...option.RequestOption) (res *AggregationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/aggregations", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Aggregation with the given UUID.
func (r *AggregationService) Get(ctx context.Context, id string, query AggregationGetParams, opts ...option.RequestOption) (res *AggregationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/aggregations/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Aggregation with the given UUID.
//
// **Note:** If you have created Custom Fields for an Aggregation, when you use
// this endpoint to update the Aggregation use the `customFields` parameter to
// preserve those Custom Fields. If you omit them from the update request, they
// will be lost.
func (r *AggregationService) Update(ctx context.Context, id string, params AggregationUpdateParams, opts ...option.RequestOption) (res *AggregationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/aggregations/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Aggregations that can be filtered by Product, Aggregation ID,
// or Code.
func (r *AggregationService) List(ctx context.Context, params AggregationListParams, opts ...option.RequestOption) (res *pagination.Cursor[AggregationResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/aggregations", params.OrgID)
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

// Retrieve a list of Aggregations that can be filtered by Product, Aggregation ID,
// or Code.
func (r *AggregationService) ListAutoPaging(ctx context.Context, params AggregationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AggregationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Aggregation with the given UUID.
func (r *AggregationService) Delete(ctx context.Context, id string, body AggregationDeleteParams, opts ...option.RequestOption) (res *AggregationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/aggregations/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AggregationResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// Optional Product ID this Aggregation should be attributed to for accounting
	// purposes.
	AccountingProductID string `json:"accountingProductId"`
	// Specifies the computation method applied to usage data collected in
	// `targetField`. Aggregation unit value depends on the **Category** configured for
	// the selected targetField.
	//
	// Enum:
	//
	//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
	//     **Cost** `targetField`.
	//
	//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
	//     or **Cost** `targetField`.
	//
	//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
	//     or **Cost** `targetField`.
	//
	//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
	//     **Income**, or **Cost** `targetField`.
	//
	//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
	//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`)
	//     value of usage data measurement submissions. If using this method, please
	//     ensure _distinct_ `ts` values are used for usage data measurment submissions.
	//
	//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
	//     **Measure**, **Income**, or **Cost** `targetField`.
	//
	//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
	//     values. Can be applied to a **Metadata** `targetField`.
	//
	//   - **CUSTOM_SQL**. Uses an SQL query expression. The `customSQL` parameter is
	//     used for the SQL query.
	Aggregation AggregationResponseAggregation `json:"aggregation"`
	// Code of the Aggregation. A unique short code to identify the Aggregation.
	Code string `json:"code"`
	// The id of the user who created this aggregation.
	CreatedBy    string                                          `json:"createdBy"`
	CustomFields map[string]AggregationResponseCustomFieldsUnion `json:"customFields"`
	// The SQL query expression to be used for a Custom SQL Aggregation.
	CustomSql string `json:"customSql"`
	// Aggregation value used when no usage data is available to be aggregated.
	// _(Optional)_.
	//
	// **Note:** Set to 0, if you expect to reference the Aggregation in a Compound
	// Aggregation. This ensures that any null values are passed in correctly to the
	// Compound Aggregation calculation with a value = 0.
	DefaultValue float64 `json:"defaultValue"`
	// The DateTime when the aggregation was created _(in ISO 8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the aggregation was last modified _(in ISO 8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this aggregation.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The UUID of the Meter used as the source of usage data for the Aggregation.
	//
	// Each Aggregation is a child of a Meter, so the Meter must be selected.
	MeterID string `json:"meterId"`
	// Descriptive name for the Aggregation.
	Name string `json:"name"`
	// Defines how much of a quantity equates to 1 unit. Used when setting the price
	// per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at
	// rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at
	// $0.25 per unit.
	//
	// If `quantityPerUnit` is set to a value other than one, rounding is typically set
	// to UP.
	QuantityPerUnit float64 `json:"quantityPerUnit"`
	// Specifies how you want to deal with non-integer, fractional number Aggregation
	// values.
	//
	// **NOTES:**
	//
	//   - **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is
	//     rounded to 4.
	//   - Also used in combination with `quantityPerUnit`. Rounds the number of units
	//     after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value
	//     other than one, you would typically set Rounding to **UP**. For example,
	//     suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` =
	//     500, and set charge rate at $0.25 per unit used. If your customer used 48,900
	//     KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up
	//     to 98 \* 0.25 = $2.45.
	//
	// Enum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???
	Rounding AggregationResponseRounding `json:"rounding"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segments`.
	//
	// The `Codes` of the fields in the target Meter to use for segmentation purposes.
	//
	// String `dataFields` on the target Meter can be segmented. Any string
	// `derivedFields` on the target Meter, such as one that concatenates two string
	// `dataFields`, can also be segmented.
	SegmentedFields []string `json:"segmentedFields"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segmentedFields`.
	//
	// Contains the values that are to be used as the segments, read from the fields in
	// the meter pointed at by `segmentedFields`.
	Segments []map[string]string `json:"segments"`
	// `Code` of the target `dataField` or `derivedField` on the Meter used as the
	// basis for the Aggregation.
	TargetField string `json:"targetField"`
	// User defined or following the _Unified Code for Units of Measure_ (UCUM).
	//
	// Used as the label for billing, indicating to your customers what they are being
	// charged for.
	Unit string `json:"unit"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                   `json:"version"`
	JSON    aggregationResponseJSON `json:"-"`
}

// aggregationResponseJSON contains the JSON metadata for the struct
// [AggregationResponse]
type aggregationResponseJSON struct {
	ID                  apijson.Field
	AccountingProductID apijson.Field
	Aggregation         apijson.Field
	Code                apijson.Field
	CreatedBy           apijson.Field
	CustomFields        apijson.Field
	CustomSql           apijson.Field
	DefaultValue        apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	LastModifiedBy      apijson.Field
	MeterID             apijson.Field
	Name                apijson.Field
	QuantityPerUnit     apijson.Field
	Rounding            apijson.Field
	SegmentedFields     apijson.Field
	Segments            apijson.Field
	TargetField         apijson.Field
	Unit                apijson.Field
	Version             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AggregationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregationResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the computation method applied to usage data collected in
// `targetField`. Aggregation unit value depends on the **Category** configured for
// the selected targetField.
//
// Enum:
//
//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
//     **Cost** `targetField`.
//
//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`.
//
//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`)
//     value of usage data measurement submissions. If using this method, please
//     ensure _distinct_ `ts` values are used for usage data measurment submissions.
//
//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
//     **Measure**, **Income**, or **Cost** `targetField`.
//
//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
//     values. Can be applied to a **Metadata** `targetField`.
//
//   - **CUSTOM_SQL**. Uses an SQL query expression. The `customSQL` parameter is
//     used for the SQL query.
type AggregationResponseAggregation string

const (
	AggregationResponseAggregationSum       AggregationResponseAggregation = "SUM"
	AggregationResponseAggregationMin       AggregationResponseAggregation = "MIN"
	AggregationResponseAggregationMax       AggregationResponseAggregation = "MAX"
	AggregationResponseAggregationCount     AggregationResponseAggregation = "COUNT"
	AggregationResponseAggregationLatest    AggregationResponseAggregation = "LATEST"
	AggregationResponseAggregationMean      AggregationResponseAggregation = "MEAN"
	AggregationResponseAggregationUnique    AggregationResponseAggregation = "UNIQUE"
	AggregationResponseAggregationCustomSql AggregationResponseAggregation = "CUSTOM_SQL"
)

func (r AggregationResponseAggregation) IsKnown() bool {
	switch r {
	case AggregationResponseAggregationSum, AggregationResponseAggregationMin, AggregationResponseAggregationMax, AggregationResponseAggregationCount, AggregationResponseAggregationLatest, AggregationResponseAggregationMean, AggregationResponseAggregationUnique, AggregationResponseAggregationCustomSql:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type AggregationResponseCustomFieldsUnion interface {
	ImplementsAggregationResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AggregationResponseCustomFieldsUnion)(nil)).Elem(),
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

// Specifies how you want to deal with non-integer, fractional number Aggregation
// values.
//
// **NOTES:**
//
//   - **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is
//     rounded to 4.
//   - Also used in combination with `quantityPerUnit`. Rounds the number of units
//     after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value
//     other than one, you would typically set Rounding to **UP**. For example,
//     suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` =
//     500, and set charge rate at $0.25 per unit used. If your customer used 48,900
//     KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up
//     to 98 \* 0.25 = $2.45.
//
// Enum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???
type AggregationResponseRounding string

const (
	AggregationResponseRoundingUp      AggregationResponseRounding = "UP"
	AggregationResponseRoundingDown    AggregationResponseRounding = "DOWN"
	AggregationResponseRoundingNearest AggregationResponseRounding = "NEAREST"
	AggregationResponseRoundingNone    AggregationResponseRounding = "NONE"
)

func (r AggregationResponseRounding) IsKnown() bool {
	switch r {
	case AggregationResponseRoundingUp, AggregationResponseRoundingDown, AggregationResponseRoundingNearest, AggregationResponseRoundingNone:
		return true
	}
	return false
}

type AggregationNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Specifies the computation method applied to usage data collected in
	// `targetField`. Aggregation unit value depends on the **Category** configured for
	// the selected `targetField`.
	//
	// Enum:
	//
	//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
	//     **Cost** `targetField`.
	//
	//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
	//     or **Cost** `targetField`.
	//
	//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
	//     or **Cost** `targetField`.
	//
	//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
	//     **Income**, or **Cost** `targetField`.
	//
	//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
	//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`)
	//     value of usage data measurement submissions. If using this method, please
	//     ensure _distinct_ `ts` values are used for usage data measurment submissions.
	//
	//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
	//     **Measure**, **Income**, or **Cost** `targetField`.
	//
	//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
	//     values. Can be applied to a **Metadata** `targetField`.
	//
	//   - **CUSTOM_SQL**. Uses an SQL query expression. If you select this Aggregation
	//     type, use the `customSQL` request parameter to enter an SQL query.
	Aggregation param.Field[AggregationNewParamsAggregation] `json:"aggregation" api:"required"`
	// The UUID of the Meter used as the source of usage data for the Aggregation.
	//
	// Each Aggregation is a child of a Meter, so the Meter must be selected.
	MeterID param.Field[string] `json:"meterId" api:"required"`
	// Descriptive name for the Aggregation.
	Name param.Field[string] `json:"name" api:"required"`
	// Defines how much of a quantity equates to 1 unit. Used when setting the price
	// per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at
	// rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at
	// $0.25 per unit.
	//
	// **Note:** If `quantityPerUnit` is set to a value other than one, `rounding` is
	// typically set to `"UP"`.
	QuantityPerUnit param.Field[float64] `json:"quantityPerUnit" api:"required"`
	// Specifies how you want to deal with non-integer, fractional number Aggregation
	// values.
	//
	// **NOTES:**
	//
	//   - **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is
	//     rounded to 4.
	//   - Also used in combination with `quantityPerUnit`. Rounds the number of units
	//     after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value
	//     other than one, you would typically set Rounding to **UP**. For example,
	//     suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` =
	//     500, and set charge rate at $0.25 per unit used. If your customer used 48,900
	//     KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up
	//     to 98 \* 0.25 = $2.45.
	//
	// Enum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???
	Rounding param.Field[AggregationNewParamsRounding] `json:"rounding" api:"required"`
	// `Code` of the target `dataField` or `derivedField` on the Meter used as the
	// basis for the Aggregation.
	TargetField param.Field[string] `json:"targetField" api:"required"`
	// User defined label for units shown for Bill line items, indicating to your
	// customers what they are being charged for.
	Unit param.Field[string] `json:"unit" api:"required"`
	// Optional Product ID this Aggregation should be attributed to for accounting
	// purposes.
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Code of the new Aggregation. A unique short code to identify the Aggregation.
	Code         param.Field[string]                                           `json:"code"`
	CustomFields param.Field[map[string]AggregationNewParamsCustomFieldsUnion] `json:"customFields"`
	// Enter the SQL query expression to be used for a Custom SQL Aggregation. Custom
	// SQL queries should be run against the Measurements table - for more details see
	// [Custom SQL Aggregations](https://www.m3ter.com/docs/guides/usage-data-aggregations/custom-sql-aggregations)
	// in your main User documentation.
	//
	// **NOTE:** The `customSql` Aggregation type is currently available in Preview
	// release. If you are interested in using this feature, please get in touch with
	// m3ter Support or your m3ter contact.
	CustomSql param.Field[string] `json:"customSql"`
	// Aggregation value used when no usage data is available to be aggregated.
	// _(Optional)_.
	//
	// **Note:** Set to 0, if you expect to reference the Aggregation in a Compound
	// Aggregation. This ensures that any null values are passed in correctly to the
	// Compound Aggregation calculation with a value = 0.
	DefaultValue param.Field[float64] `json:"defaultValue"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segments`.
	//
	// Enter the `Codes` of the fields in the target Meter to use for segmentation
	// purposes.
	//
	// String `dataFields` on the target Meter can be segmented. Any string
	// `derivedFields` on the target Meter, such as one that concatenates two string
	// `dataFields`, can also be segmented.
	SegmentedFields param.Field[[]string] `json:"segmentedFields"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segmentedFields`.
	//
	// Enter the values that are to be used as the segments, read from the fields in
	// the meter pointed at by `segmentedFields`.
	//
	// Note that you can use _wildcards_ or _defaults_ when setting up segment values.
	// For more details on how to do this with an example, see
	// [Using Wildcards - API Calls](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/segmented-aggregations#using-wildcards---api-calls)
	// in our main User Docs.
	Segments param.Field[[]map[string]string] `json:"segments"`
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

func (r AggregationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the computation method applied to usage data collected in
// `targetField`. Aggregation unit value depends on the **Category** configured for
// the selected `targetField`.
//
// Enum:
//
//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
//     **Cost** `targetField`.
//
//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`.
//
//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`)
//     value of usage data measurement submissions. If using this method, please
//     ensure _distinct_ `ts` values are used for usage data measurment submissions.
//
//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
//     **Measure**, **Income**, or **Cost** `targetField`.
//
//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
//     values. Can be applied to a **Metadata** `targetField`.
//
//   - **CUSTOM_SQL**. Uses an SQL query expression. If you select this Aggregation
//     type, use the `customSQL` request parameter to enter an SQL query.
type AggregationNewParamsAggregation string

const (
	AggregationNewParamsAggregationSum       AggregationNewParamsAggregation = "SUM"
	AggregationNewParamsAggregationMin       AggregationNewParamsAggregation = "MIN"
	AggregationNewParamsAggregationMax       AggregationNewParamsAggregation = "MAX"
	AggregationNewParamsAggregationCount     AggregationNewParamsAggregation = "COUNT"
	AggregationNewParamsAggregationLatest    AggregationNewParamsAggregation = "LATEST"
	AggregationNewParamsAggregationMean      AggregationNewParamsAggregation = "MEAN"
	AggregationNewParamsAggregationUnique    AggregationNewParamsAggregation = "UNIQUE"
	AggregationNewParamsAggregationCustomSql AggregationNewParamsAggregation = "CUSTOM_SQL"
)

func (r AggregationNewParamsAggregation) IsKnown() bool {
	switch r {
	case AggregationNewParamsAggregationSum, AggregationNewParamsAggregationMin, AggregationNewParamsAggregationMax, AggregationNewParamsAggregationCount, AggregationNewParamsAggregationLatest, AggregationNewParamsAggregationMean, AggregationNewParamsAggregationUnique, AggregationNewParamsAggregationCustomSql:
		return true
	}
	return false
}

// Specifies how you want to deal with non-integer, fractional number Aggregation
// values.
//
// **NOTES:**
//
//   - **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is
//     rounded to 4.
//   - Also used in combination with `quantityPerUnit`. Rounds the number of units
//     after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value
//     other than one, you would typically set Rounding to **UP**. For example,
//     suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` =
//     500, and set charge rate at $0.25 per unit used. If your customer used 48,900
//     KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up
//     to 98 \* 0.25 = $2.45.
//
// Enum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???
type AggregationNewParamsRounding string

const (
	AggregationNewParamsRoundingUp      AggregationNewParamsRounding = "UP"
	AggregationNewParamsRoundingDown    AggregationNewParamsRounding = "DOWN"
	AggregationNewParamsRoundingNearest AggregationNewParamsRounding = "NEAREST"
	AggregationNewParamsRoundingNone    AggregationNewParamsRounding = "NONE"
)

func (r AggregationNewParamsRounding) IsKnown() bool {
	switch r {
	case AggregationNewParamsRoundingUp, AggregationNewParamsRoundingDown, AggregationNewParamsRoundingNearest, AggregationNewParamsRoundingNone:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type AggregationNewParamsCustomFieldsUnion interface {
	ImplementsAggregationNewParamsCustomFieldsUnion()
}

type AggregationGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type AggregationUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Specifies the computation method applied to usage data collected in
	// `targetField`. Aggregation unit value depends on the **Category** configured for
	// the selected `targetField`.
	//
	// Enum:
	//
	//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
	//     **Cost** `targetField`.
	//
	//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
	//     or **Cost** `targetField`.
	//
	//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
	//     or **Cost** `targetField`.
	//
	//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
	//     **Income**, or **Cost** `targetField`.
	//
	//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
	//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`)
	//     value of usage data measurement submissions. If using this method, please
	//     ensure _distinct_ `ts` values are used for usage data measurment submissions.
	//
	//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
	//     **Measure**, **Income**, or **Cost** `targetField`.
	//
	//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
	//     values. Can be applied to a **Metadata** `targetField`.
	//
	//   - **CUSTOM_SQL**. Uses an SQL query expression. If you select this Aggregation
	//     type, use the `customSQL` request parameter to enter an SQL query.
	Aggregation param.Field[AggregationUpdateParamsAggregation] `json:"aggregation" api:"required"`
	// The UUID of the Meter used as the source of usage data for the Aggregation.
	//
	// Each Aggregation is a child of a Meter, so the Meter must be selected.
	MeterID param.Field[string] `json:"meterId" api:"required"`
	// Descriptive name for the Aggregation.
	Name param.Field[string] `json:"name" api:"required"`
	// Defines how much of a quantity equates to 1 unit. Used when setting the price
	// per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at
	// rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at
	// $0.25 per unit.
	//
	// **Note:** If `quantityPerUnit` is set to a value other than one, `rounding` is
	// typically set to `"UP"`.
	QuantityPerUnit param.Field[float64] `json:"quantityPerUnit" api:"required"`
	// Specifies how you want to deal with non-integer, fractional number Aggregation
	// values.
	//
	// **NOTES:**
	//
	//   - **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is
	//     rounded to 4.
	//   - Also used in combination with `quantityPerUnit`. Rounds the number of units
	//     after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value
	//     other than one, you would typically set Rounding to **UP**. For example,
	//     suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` =
	//     500, and set charge rate at $0.25 per unit used. If your customer used 48,900
	//     KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up
	//     to 98 \* 0.25 = $2.45.
	//
	// Enum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???
	Rounding param.Field[AggregationUpdateParamsRounding] `json:"rounding" api:"required"`
	// `Code` of the target `dataField` or `derivedField` on the Meter used as the
	// basis for the Aggregation.
	TargetField param.Field[string] `json:"targetField" api:"required"`
	// User defined label for units shown for Bill line items, indicating to your
	// customers what they are being charged for.
	Unit param.Field[string] `json:"unit" api:"required"`
	// Optional Product ID this Aggregation should be attributed to for accounting
	// purposes.
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Code of the new Aggregation. A unique short code to identify the Aggregation.
	Code         param.Field[string]                                              `json:"code"`
	CustomFields param.Field[map[string]AggregationUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// Enter the SQL query expression to be used for a Custom SQL Aggregation. Custom
	// SQL queries should be run against the Measurements table - for more details see
	// [Custom SQL Aggregations](https://www.m3ter.com/docs/guides/usage-data-aggregations/custom-sql-aggregations)
	// in your main User documentation.
	//
	// **NOTE:** The `customSql` Aggregation type is currently available in Preview
	// release. If you are interested in using this feature, please get in touch with
	// m3ter Support or your m3ter contact.
	CustomSql param.Field[string] `json:"customSql"`
	// Aggregation value used when no usage data is available to be aggregated.
	// _(Optional)_.
	//
	// **Note:** Set to 0, if you expect to reference the Aggregation in a Compound
	// Aggregation. This ensures that any null values are passed in correctly to the
	// Compound Aggregation calculation with a value = 0.
	DefaultValue param.Field[float64] `json:"defaultValue"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segments`.
	//
	// Enter the `Codes` of the fields in the target Meter to use for segmentation
	// purposes.
	//
	// String `dataFields` on the target Meter can be segmented. Any string
	// `derivedFields` on the target Meter, such as one that concatenates two string
	// `dataFields`, can also be segmented.
	SegmentedFields param.Field[[]string] `json:"segmentedFields"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segmentedFields`.
	//
	// Enter the values that are to be used as the segments, read from the fields in
	// the meter pointed at by `segmentedFields`.
	//
	// Note that you can use _wildcards_ or _defaults_ when setting up segment values.
	// For more details on how to do this with an example, see
	// [Using Wildcards - API Calls](https://www.m3ter.com/docs/guides/setting-up-usage-data-meters-and-aggregations/segmented-aggregations#using-wildcards---api-calls)
	// in our main User Docs.
	Segments param.Field[[]map[string]string] `json:"segments"`
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

func (r AggregationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the computation method applied to usage data collected in
// `targetField`. Aggregation unit value depends on the **Category** configured for
// the selected `targetField`.
//
// Enum:
//
//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
//     **Cost** `targetField`.
//
//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`.
//
//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp (`ts`)
//     value of usage data measurement submissions. If using this method, please
//     ensure _distinct_ `ts` values are used for usage data measurment submissions.
//
//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
//     **Measure**, **Income**, or **Cost** `targetField`.
//
//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
//     values. Can be applied to a **Metadata** `targetField`.
//
//   - **CUSTOM_SQL**. Uses an SQL query expression. If you select this Aggregation
//     type, use the `customSQL` request parameter to enter an SQL query.
type AggregationUpdateParamsAggregation string

const (
	AggregationUpdateParamsAggregationSum       AggregationUpdateParamsAggregation = "SUM"
	AggregationUpdateParamsAggregationMin       AggregationUpdateParamsAggregation = "MIN"
	AggregationUpdateParamsAggregationMax       AggregationUpdateParamsAggregation = "MAX"
	AggregationUpdateParamsAggregationCount     AggregationUpdateParamsAggregation = "COUNT"
	AggregationUpdateParamsAggregationLatest    AggregationUpdateParamsAggregation = "LATEST"
	AggregationUpdateParamsAggregationMean      AggregationUpdateParamsAggregation = "MEAN"
	AggregationUpdateParamsAggregationUnique    AggregationUpdateParamsAggregation = "UNIQUE"
	AggregationUpdateParamsAggregationCustomSql AggregationUpdateParamsAggregation = "CUSTOM_SQL"
)

func (r AggregationUpdateParamsAggregation) IsKnown() bool {
	switch r {
	case AggregationUpdateParamsAggregationSum, AggregationUpdateParamsAggregationMin, AggregationUpdateParamsAggregationMax, AggregationUpdateParamsAggregationCount, AggregationUpdateParamsAggregationLatest, AggregationUpdateParamsAggregationMean, AggregationUpdateParamsAggregationUnique, AggregationUpdateParamsAggregationCustomSql:
		return true
	}
	return false
}

// Specifies how you want to deal with non-integer, fractional number Aggregation
// values.
//
// **NOTES:**
//
//   - **NEAREST** rounds to the nearest half: 5.1 is rounded to 5, and 3.5 is
//     rounded to 4.
//   - Also used in combination with `quantityPerUnit`. Rounds the number of units
//     after `quantityPerUnit` is applied. If you set `quantityPerUnit` to a value
//     other than one, you would typically set Rounding to **UP**. For example,
//     suppose you charge by kilobytes per second (KiBy/s), set `quantityPerUnit` =
//     500, and set charge rate at $0.25 per unit used. If your customer used 48,900
//     KiBy/s in a billing period, the charge would be 48,900 / 500 = 97.8 rounded up
//     to 98 \* 0.25 = $2.45.
//
// Enum: ???UP??? ???DOWN??? ???NEAREST??? ???NONE???
type AggregationUpdateParamsRounding string

const (
	AggregationUpdateParamsRoundingUp      AggregationUpdateParamsRounding = "UP"
	AggregationUpdateParamsRoundingDown    AggregationUpdateParamsRounding = "DOWN"
	AggregationUpdateParamsRoundingNearest AggregationUpdateParamsRounding = "NEAREST"
	AggregationUpdateParamsRoundingNone    AggregationUpdateParamsRounding = "NONE"
)

func (r AggregationUpdateParamsRounding) IsKnown() bool {
	switch r {
	case AggregationUpdateParamsRoundingUp, AggregationUpdateParamsRoundingDown, AggregationUpdateParamsRoundingNearest, AggregationUpdateParamsRoundingNone:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type AggregationUpdateParamsCustomFieldsUnion interface {
	ImplementsAggregationUpdateParamsCustomFieldsUnion()
}

type AggregationListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// List of Aggregation codes to retrieve. These are unique short codes to identify
	// each Aggregation.
	Codes param.Field[[]string] `query:"codes"`
	// List of Aggregation IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi-page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Aggregations to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// The UUIDs of the Products to retrieve Aggregations for.
	ProductID param.Field[[]string] `query:"productId"`
}

// URLQuery serializes [AggregationListParams]'s query parameters as `url.Values`.
func (r AggregationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AggregationDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
