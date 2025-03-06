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

// CompoundAggregationService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompoundAggregationService] method instead.
type CompoundAggregationService struct {
	Options []option.RequestOption
}

// NewCompoundAggregationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCompoundAggregationService(opts ...option.RequestOption) (r *CompoundAggregationService) {
	r = &CompoundAggregationService{}
	r.Options = opts
	return
}

// Create a new CompoundAggregation.
//
// This endpoint allows you to create a new CompoundAggregation for a specific
// Organization. The request body must include all the necessary details such as
// the Calculation formula.
func (r *CompoundAggregationService) New(ctx context.Context, params CompoundAggregationNewParams, opts ...option.RequestOption) (res *AggregationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/compoundaggregations", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a CompoundAggregation using the given UUID.
//
// This endpoint returns a specific CompoundAggregation associated with an
// Organization. It provides detailed information about the CompoundAggregation.
func (r *CompoundAggregationService) Get(ctx context.Context, id string, query CompoundAggregationGetParams, opts ...option.RequestOption) (res *CompoundAggregationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/compoundaggregations/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the CompoundAggregation with the given UUID.
//
// This endpoint allows you to update the details of a specific CompoundAggregation
// associated with an Organization. Use it to modify details of an existing
// CompoundAggregation such as the Calculation formula.
//
// **Note:** If you have created Custom Fields for a Compound Aggregation, when you
// use this endpoint to update the Compound Aggregation use the `customFields`
// parameter to preserve those Custom Fields. If you omit them from the update
// request, they will be lost.
func (r *CompoundAggregationService) Update(ctx context.Context, id string, params CompoundAggregationUpdateParams, opts ...option.RequestOption) (res *AggregationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/compoundaggregations/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of all CompoundAggregations.
//
// This endpoint retrieves a list of CompoundAggregations associated with a
// specific organization. CompoundAggregations enable you to define numerical
// measures based on simple Aggregations of usage data. It supports pagination, and
// includes various query parameters to filter the CompoundAggregations based on
// Product, CompoundAggregation IDs or short codes.
func (r *CompoundAggregationService) List(ctx context.Context, params CompoundAggregationListParams, opts ...option.RequestOption) (res *pagination.Cursor[CompoundAggregationResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/compoundaggregations", params.OrgID)
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

// Retrieve a list of all CompoundAggregations.
//
// This endpoint retrieves a list of CompoundAggregations associated with a
// specific organization. CompoundAggregations enable you to define numerical
// measures based on simple Aggregations of usage data. It supports pagination, and
// includes various query parameters to filter the CompoundAggregations based on
// Product, CompoundAggregation IDs or short codes.
func (r *CompoundAggregationService) ListAutoPaging(ctx context.Context, params CompoundAggregationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CompoundAggregationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a CompoundAggregation with the given UUID.
//
// This endpoint enables deletion of a specific CompoundAggregation associated with
// a specific Organization. Useful when you need to remove an existing
// CompoundAggregation that is no longer required, such as when changing pricing or
// planning models.
func (r *CompoundAggregationService) Delete(ctx context.Context, id string, body CompoundAggregationDeleteParams, opts ...option.RequestOption) (res *CompoundAggregationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/compoundaggregations/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CompoundAggregationResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version             int64  `json:"version,required"`
	AccountingProductID string `json:"accountingProductId"`
	// This field is a string that represents the formula for the calculation. This
	// formula determines how the CompoundAggregation is calculated from the underlying
	// usage data.
	Calculation string `json:"calculation"`
	// Code of the Aggregation. A unique short code to identify the Aggregation.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this CompoundAggregation.
	CreatedBy    string                                                  `json:"createdBy"`
	CustomFields map[string]CompoundAggregationResponseCustomFieldsUnion `json:"customFields"`
	// The date and time _(in ISO-8601 format)_ when the CompoundAggregation was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the CompoundAggregation was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// This is a boolean True / False flag.
	//
	// If set to TRUE, the calculation will be evaluated even if the referenced
	// aggregation has no usage data.
	EvaluateNullAggregations bool `json:"evaluateNullAggregations"`
	// The unique identifier (UUID) of the user who last modified this
	// CompoundAggregation.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Descriptive name for the Aggregation.
	Name string `json:"name"`
	// This field represents the unique identifier (UUID) of the Product that is
	// associated with the CompoundAggregation.
	ProductID string `json:"productId"`
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
	Rounding CompoundAggregationResponseRounding `json:"rounding"`
	// _(Optional)_. Used when creating a segmented Aggregation, which segments the
	// usage data collected by a single Meter. Works together with `segmentedFields`.
	//
	// Contains the values that are to be used as the segments, read from the fields in
	// the meter pointed at by `segmentedFields`.
	Segments []map[string]string `json:"segments"`
	// User defined or following the _Unified Code for Units of Measure_ (UCUM).
	//
	// Used as the label for billing, indicating to your customers what they are being
	// charged for.
	Unit string                          `json:"unit"`
	JSON compoundAggregationResponseJSON `json:"-"`
}

// compoundAggregationResponseJSON contains the JSON metadata for the struct
// [CompoundAggregationResponse]
type compoundAggregationResponseJSON struct {
	ID                       apijson.Field
	Version                  apijson.Field
	AccountingProductID      apijson.Field
	Calculation              apijson.Field
	Code                     apijson.Field
	CreatedBy                apijson.Field
	CustomFields             apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	EvaluateNullAggregations apijson.Field
	LastModifiedBy           apijson.Field
	Name                     apijson.Field
	ProductID                apijson.Field
	QuantityPerUnit          apijson.Field
	Rounding                 apijson.Field
	Segments                 apijson.Field
	Unit                     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CompoundAggregationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compoundAggregationResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CompoundAggregationResponseCustomFieldsUnion interface {
	ImplementsCompoundAggregationResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CompoundAggregationResponseCustomFieldsUnion)(nil)).Elem(),
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
type CompoundAggregationResponseRounding string

const (
	CompoundAggregationResponseRoundingUp      CompoundAggregationResponseRounding = "UP"
	CompoundAggregationResponseRoundingDown    CompoundAggregationResponseRounding = "DOWN"
	CompoundAggregationResponseRoundingNearest CompoundAggregationResponseRounding = "NEAREST"
	CompoundAggregationResponseRoundingNone    CompoundAggregationResponseRounding = "NONE"
)

func (r CompoundAggregationResponseRounding) IsKnown() bool {
	switch r {
	case CompoundAggregationResponseRoundingUp, CompoundAggregationResponseRoundingDown, CompoundAggregationResponseRoundingNearest, CompoundAggregationResponseRoundingNone:
		return true
	}
	return false
}

type CompoundAggregationNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// String that represents the formula for the calculation. This formula determines
	// how the CompoundAggregation value is calculated. The calculation can reference
	// simple Aggregations or Custom Fields. This field is required when creating or
	// updating a CompoundAggregation.
	//
	// **NOTE:** If a simple Aggregation referenced by a Compound Aggregation has a
	// **Quantity per unit** defined or a **Rounding** defined, these will not be
	// factored into the value used by the calculation. For example, if the simple
	// Aggregation referenced has a base value of 100 and has **Quantity per unit** set
	// at 10, the Compound Aggregation calculation _will use the base value of 100 not
	// 10_.
	Calculation param.Field[string] `json:"calculation,required"`
	// Descriptive name for the Aggregation.
	Name param.Field[string] `json:"name,required"`
	// Defines how much of a quantity equates to 1 unit. Used when setting the price
	// per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at
	// rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at
	// $0.25 per unit.
	//
	// **Note:** If `quantityPerUnit` is set to a value other than one, `rounding` is
	// typically set to `"UP"`.
	QuantityPerUnit param.Field[float64] `json:"quantityPerUnit,required"`
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
	Rounding param.Field[CompoundAggregationNewParamsRounding] `json:"rounding,required"`
	// User defined label for units shown for Bill line items, indicating to your
	// customers what they are being charged for.
	Unit param.Field[string] `json:"unit,required"`
	// Optional Product ID this Aggregation should be attributed to for accounting
	// purposes
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Code of the new Aggregation. A unique short code to identify the Aggregation.
	Code         param.Field[string]                                                   `json:"code"`
	CustomFields param.Field[map[string]CompoundAggregationNewParamsCustomFieldsUnion] `json:"customFields"`
	// Boolean True / False flag:
	//
	//   - **TRUE** - set to TRUE if you want to allow null values from the simple
	//     Aggregations referenced in the Compound Aggregation to be passed in. Simple
	//     Aggregations based on Meter Target Fields where no usage data is available
	//     will have null values.
	//   - **FALSE** Default.
	//
	// **Note:** If any of the simple Aggregations you reference in a Compound
	// Aggregation calculation might have null values, you must set their Default Value
	// to 0. This ensures that any null values passed into the Compound Aggregation are
	// passed in correctly with value = 0.
	EvaluateNullAggregations param.Field[bool] `json:"evaluateNullAggregations"`
	// Unique identifier (UUID) of the Product the CompoundAggregation belongs to.
	//
	// **Note:** Omit this parameter if you want to create a _Global_
	// CompoundAggregation.
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

func (r CompoundAggregationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
type CompoundAggregationNewParamsRounding string

const (
	CompoundAggregationNewParamsRoundingUp      CompoundAggregationNewParamsRounding = "UP"
	CompoundAggregationNewParamsRoundingDown    CompoundAggregationNewParamsRounding = "DOWN"
	CompoundAggregationNewParamsRoundingNearest CompoundAggregationNewParamsRounding = "NEAREST"
	CompoundAggregationNewParamsRoundingNone    CompoundAggregationNewParamsRounding = "NONE"
)

func (r CompoundAggregationNewParamsRounding) IsKnown() bool {
	switch r {
	case CompoundAggregationNewParamsRoundingUp, CompoundAggregationNewParamsRoundingDown, CompoundAggregationNewParamsRoundingNearest, CompoundAggregationNewParamsRoundingNone:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CompoundAggregationNewParamsCustomFieldsUnion interface {
	ImplementsCompoundAggregationNewParamsCustomFieldsUnion()
}

type CompoundAggregationGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type CompoundAggregationUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// String that represents the formula for the calculation. This formula determines
	// how the CompoundAggregation value is calculated. The calculation can reference
	// simple Aggregations or Custom Fields. This field is required when creating or
	// updating a CompoundAggregation.
	//
	// **NOTE:** If a simple Aggregation referenced by a Compound Aggregation has a
	// **Quantity per unit** defined or a **Rounding** defined, these will not be
	// factored into the value used by the calculation. For example, if the simple
	// Aggregation referenced has a base value of 100 and has **Quantity per unit** set
	// at 10, the Compound Aggregation calculation _will use the base value of 100 not
	// 10_.
	Calculation param.Field[string] `json:"calculation,required"`
	// Descriptive name for the Aggregation.
	Name param.Field[string] `json:"name,required"`
	// Defines how much of a quantity equates to 1 unit. Used when setting the price
	// per unit for billing purposes - if charging for kilobytes per second (KiBy/s) at
	// rate of $0.25 per 500 KiBy/s, then set quantityPerUnit to 500 and price Plan at
	// $0.25 per unit.
	//
	// **Note:** If `quantityPerUnit` is set to a value other than one, `rounding` is
	// typically set to `"UP"`.
	QuantityPerUnit param.Field[float64] `json:"quantityPerUnit,required"`
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
	Rounding param.Field[CompoundAggregationUpdateParamsRounding] `json:"rounding,required"`
	// User defined label for units shown for Bill line items, indicating to your
	// customers what they are being charged for.
	Unit param.Field[string] `json:"unit,required"`
	// Optional Product ID this Aggregation should be attributed to for accounting
	// purposes
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Code of the new Aggregation. A unique short code to identify the Aggregation.
	Code         param.Field[string]                                                      `json:"code"`
	CustomFields param.Field[map[string]CompoundAggregationUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// Boolean True / False flag:
	//
	//   - **TRUE** - set to TRUE if you want to allow null values from the simple
	//     Aggregations referenced in the Compound Aggregation to be passed in. Simple
	//     Aggregations based on Meter Target Fields where no usage data is available
	//     will have null values.
	//   - **FALSE** Default.
	//
	// **Note:** If any of the simple Aggregations you reference in a Compound
	// Aggregation calculation might have null values, you must set their Default Value
	// to 0. This ensures that any null values passed into the Compound Aggregation are
	// passed in correctly with value = 0.
	EvaluateNullAggregations param.Field[bool] `json:"evaluateNullAggregations"`
	// Unique identifier (UUID) of the Product the CompoundAggregation belongs to.
	//
	// **Note:** Omit this parameter if you want to create a _Global_
	// CompoundAggregation.
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

func (r CompoundAggregationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
type CompoundAggregationUpdateParamsRounding string

const (
	CompoundAggregationUpdateParamsRoundingUp      CompoundAggregationUpdateParamsRounding = "UP"
	CompoundAggregationUpdateParamsRoundingDown    CompoundAggregationUpdateParamsRounding = "DOWN"
	CompoundAggregationUpdateParamsRoundingNearest CompoundAggregationUpdateParamsRounding = "NEAREST"
	CompoundAggregationUpdateParamsRoundingNone    CompoundAggregationUpdateParamsRounding = "NONE"
)

func (r CompoundAggregationUpdateParamsRounding) IsKnown() bool {
	switch r {
	case CompoundAggregationUpdateParamsRoundingUp, CompoundAggregationUpdateParamsRoundingDown, CompoundAggregationUpdateParamsRoundingNearest, CompoundAggregationUpdateParamsRoundingNone:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CompoundAggregationUpdateParamsCustomFieldsUnion interface {
	ImplementsCompoundAggregationUpdateParamsCustomFieldsUnion()
}

type CompoundAggregationListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// An optional parameter to retrieve specific CompoundAggregations based on their
	// short codes.
	Codes param.Field[[]string] `query:"codes"`
	// An optional parameter to retrieve specific CompoundAggregations based on their
	// unique identifiers (UUIDs).
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// CompoundAggregations in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of CompoundAggregations to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// An optional parameter to filter the CompoundAggregations based on specific
	// Product unique identifiers (UUIDs).
	ProductID param.Field[[]string] `query:"productId"`
}

// URLQuery serializes [CompoundAggregationListParams]'s query parameters as
// `url.Values`.
func (r CompoundAggregationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CompoundAggregationDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
