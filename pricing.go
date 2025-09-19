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
	"github.com/m3ter-com/m3ter-sdk-go/shared"
)

// PricingService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPricingService] method instead.
type PricingService struct {
	Options []option.RequestOption
}

// NewPricingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPricingService(opts ...option.RequestOption) (r *PricingService) {
	r = &PricingService{}
	r.Options = opts
	return
}

// Create a new Pricing.
//
// **Note:** Either `planId` or `planTemplateId` request parameters are required
// for this call to be valid. If you omit both, then you will receive a validation
// error.
func (r *PricingService) New(ctx context.Context, params PricingNewParams, opts ...option.RequestOption) (res *PricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/pricings", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Pricing with the given UUID.
func (r *PricingService) Get(ctx context.Context, id string, query PricingGetParams, opts ...option.RequestOption) (res *PricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/pricings/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Pricing for the given UUID.
//
// **Note:** Either `planId` or `planTemplateId` request parameters are required
// for this call to be valid. If you omit both, then you will receive a validation
// error.
func (r *PricingService) Update(ctx context.Context, id string, params PricingUpdateParams, opts ...option.RequestOption) (res *PricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/pricings/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Pricings filtered by date, Plan ID, PlanTemplate ID, or
// Pricing ID.
func (r *PricingService) List(ctx context.Context, params PricingListParams, opts ...option.RequestOption) (res *pagination.Cursor[PricingResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/pricings", params.OrgID)
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

// Retrieve a list of Pricings filtered by date, Plan ID, PlanTemplate ID, or
// Pricing ID.
func (r *PricingService) ListAutoPaging(ctx context.Context, params PricingListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PricingResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Pricing with the given UUID.
func (r *PricingService) Delete(ctx context.Context, id string, body PricingDeleteParams, opts ...option.RequestOption) (res *PricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/pricings/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PricingResponse struct {
	// The UUID of the entity.
	ID                  string `json:"id,required"`
	AccountingProductID string `json:"accountingProductId"`
	// UUID of the Aggregation used to create the Pricing. Use this when creating a
	// Pricing for a segmented aggregation.
	AggregationID   string                         `json:"aggregationId"`
	AggregationType PricingResponseAggregationType `json:"aggregationType"`
	// Unique short code for the Pricing.
	Code string `json:"code"`
	// UUID of the Compound Aggregation used to create the Pricing.
	CompoundAggregationID string `json:"compoundAggregationId"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// Controls whether or not charge rates under a set of pricing bands configured for
	// a Pricing are applied according to each separate band or at the highest band
	// reached.
	//
	// The default value is **TRUE**.
	//
	//   - When TRUE, at billing charge rates are applied according to each separate
	//     band.
	//
	//   - When FALSE, at billing charge rates are applied according to highest band
	//     reached.
	Cumulative bool `json:"cumulative"`
	// Displayed on Bill line items.
	Description string `json:"description"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The end date _(in ISO-8601 format)_ for when the Pricing ceases to be active for
	// the Plan or Plan Template.
	//
	// If not specified, the Pricing remains active indefinitely.
	EndDate time.Time `json:"endDate" format:"date-time"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The minimum spend amount per billing cycle for end customer Accounts on a Plan
	// to which the Pricing is applied.
	MinimumSpend float64 `json:"minimumSpend"`
	// The default value is **FALSE**.
	//
	// - When TRUE, minimum spend is billed at the start of each billing period.
	//
	// - When FALSE, minimum spend is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at Organization level for minimum spend
	// billing in arrears/in advance.
	MinimumSpendBillInAdvance bool `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription string `json:"minimumSpendDescription"`
	// The Prepayment/Balance overage pricing in pricing bands for the case of a
	// **Tiered** pricing structure.
	OveragePricingBands []shared.PricingBand `json:"overagePricingBands"`
	// UUID of the Plan the Pricing is created for.
	PlanID string `json:"planId"`
	// UUID of the Plan Template the Pricing was created for.
	PlanTemplateID string               `json:"planTemplateId"`
	PricingBands   []shared.PricingBand `json:"pricingBands"`
	// Name of the segment for which you are defining a Pricing.
	//
	// For each segment in a segmented aggregation, make a separate call using
	// `aggregationId` parameter to update a Pricing.
	Segment       map[string]string `json:"segment"`
	SegmentString string            `json:"segmentString"`
	// The start date _(in ISO-8601 format)_ for when the Pricing starts to be active
	// for the Plan of Plan Template.
	StartDate time.Time `json:"startDate" format:"date-time"`
	// The default value is **FALSE**.
	//
	//   - If TRUE, usage accumulates over the entire period the priced Plan is active
	//     for the account, and is not reset for pricing band rates at the start of each
	//     billing period.
	//
	//   - If FALSE, usage does not accumulate, and is reset for pricing bands at the
	//     start of each billing period.
	TiersSpanPlan bool `json:"tiersSpanPlan"`
	//   - **DEBIT**. Default setting. The amount calculated using the Pricing is added
	//     to the bill as a debit.
	//
	//   - **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the
	//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
	//     will be capped at the total of other line items for the _same_ Product.
	//
	//   - **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the
	//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
	//     will be capped at the total of other line items for the entire bill, which
	//     might include other Products the Account consumes.
	Type PricingResponseType `json:"type"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64               `json:"version"`
	JSON    pricingResponseJSON `json:"-"`
}

// pricingResponseJSON contains the JSON metadata for the struct [PricingResponse]
type pricingResponseJSON struct {
	ID                        apijson.Field
	AccountingProductID       apijson.Field
	AggregationID             apijson.Field
	AggregationType           apijson.Field
	Code                      apijson.Field
	CompoundAggregationID     apijson.Field
	CreatedBy                 apijson.Field
	Cumulative                apijson.Field
	Description               apijson.Field
	DtCreated                 apijson.Field
	DtLastModified            apijson.Field
	EndDate                   apijson.Field
	LastModifiedBy            apijson.Field
	MinimumSpend              apijson.Field
	MinimumSpendBillInAdvance apijson.Field
	MinimumSpendDescription   apijson.Field
	OveragePricingBands       apijson.Field
	PlanID                    apijson.Field
	PlanTemplateID            apijson.Field
	PricingBands              apijson.Field
	Segment                   apijson.Field
	SegmentString             apijson.Field
	StartDate                 apijson.Field
	TiersSpanPlan             apijson.Field
	Type                      apijson.Field
	Version                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PricingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricingResponseJSON) RawJSON() string {
	return r.raw
}

type PricingResponseAggregationType string

const (
	PricingResponseAggregationTypeSimple   PricingResponseAggregationType = "SIMPLE"
	PricingResponseAggregationTypeCompound PricingResponseAggregationType = "COMPOUND"
)

func (r PricingResponseAggregationType) IsKnown() bool {
	switch r {
	case PricingResponseAggregationTypeSimple, PricingResponseAggregationTypeCompound:
		return true
	}
	return false
}

//   - **DEBIT**. Default setting. The amount calculated using the Pricing is added
//     to the bill as a debit.
//
//   - **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the
//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
//     will be capped at the total of other line items for the _same_ Product.
//
//   - **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the
//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
//     will be capped at the total of other line items for the entire bill, which
//     might include other Products the Account consumes.
type PricingResponseType string

const (
	PricingResponseTypeDebit         PricingResponseType = "DEBIT"
	PricingResponseTypeProductCredit PricingResponseType = "PRODUCT_CREDIT"
	PricingResponseTypeGlobalCredit  PricingResponseType = "GLOBAL_CREDIT"
)

func (r PricingResponseType) IsKnown() bool {
	switch r {
	case PricingResponseTypeDebit, PricingResponseTypeProductCredit, PricingResponseTypeGlobalCredit:
		return true
	}
	return false
}

type PricingNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID        param.Field[string]                    `path:"orgId,required"`
	PricingBands param.Field[[]shared.PricingBandParam] `json:"pricingBands,required"`
	// The start date _(in ISO-8601 format)_ for when the Pricing starts to be active
	// for the Plan of Plan Template._(Required)_
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// Optional Product ID this Pricing should be attributed to for accounting purposes
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// UUID of the Aggregation used to create the Pricing. Use this when creating a
	// Pricing for a segmented aggregation.
	AggregationID param.Field[string] `json:"aggregationId"`
	// Unique short code for the Pricing.
	Code param.Field[string] `json:"code"`
	// UUID of the Compound Aggregation used to create the Pricing.
	CompoundAggregationID param.Field[string] `json:"compoundAggregationId"`
	// Controls whether or not charge rates under a set of pricing bands configured for
	// a Pricing are applied according to each separate band or at the highest band
	// reached.
	//
	// _(Optional)_. The default value is **FALSE**.
	//
	//   - When TRUE, at billing charge rates are applied according to each separate
	//     band.
	//
	//   - When FALSE, at billing charge rates are applied according to highest band
	//     reached.
	//
	// **NOTE:** Use the `cumulative` parameter to create the type of Pricing you
	// require. For example, for Tiered Pricing set to **TRUE**; for Volume Pricing,
	// set to **FALSE**.
	Cumulative param.Field[bool] `json:"cumulative"`
	// Displayed on Bill line items.
	Description param.Field[string] `json:"description"`
	// The end date _(in ISO-8601 format)_ for when the Pricing ceases to be active for
	// the Plan or Plan Template.
	//
	// _(Optional)_ If not specified, the Pricing remains active indefinitely.
	EndDate param.Field[time.Time] `json:"endDate" format:"date-time"`
	// The minimum spend amount per billing cycle for end customer Accounts on a Plan
	// to which the Pricing is applied.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// The default value is **FALSE**.
	//
	// - When TRUE, minimum spend is billed at the start of each billing period.
	//
	// - When FALSE, minimum spend is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at Organization level for minimum spend
	// billing in arrears/in advance.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// Specify Prepayment/Balance overage pricing in pricing bands for the case of a
	// **Tiered** pricing structure.
	OveragePricingBands param.Field[[]shared.PricingBandParam] `json:"overagePricingBands"`
	// UUID of the Plan the Pricing is created for.
	PlanID param.Field[string] `json:"planId"`
	// UUID of the Plan Template the Pricing is created for.
	PlanTemplateID param.Field[string] `json:"planTemplateId"`
	// Specifies the segment value which you are defining a Pricing for using this
	// call:
	//
	//   - For each segment value defined on a Segmented Aggregation you must create a
	//     separate Pricing and use the appropriate `aggregationId` parameter for the
	//     call.
	//   - If you specify a segment value that has not been defined for the Aggregation,
	//     you'll receive an error.
	//   - If you've defined segment values for the Aggregation using a single wildcard
	//     or multiple wildcards, you can create Pricing for these wildcard segment
	//     values also.
	//
	// For more details on creating Pricings for segment values on a Segmented
	// Aggregation using this call, together with some examples, see the
	// [Using API Call to Create Segmented Pricings](https://www.m3ter.com/docs/guides/plans-and-pricing/pricing-plans/pricing-plans-using-segmented-aggregations#using-api-call-to-create-a-segmented-pricing)
	// in our User Documentation.
	Segment param.Field[map[string]string] `json:"segment"`
	// The default value is **FALSE**.
	//
	//   - If TRUE, usage accumulates over the entire period the priced Plan is active
	//     for the account, and is not reset for pricing band rates at the start of each
	//     billing period.
	//
	//   - If FALSE, usage does not accumulate, and is reset for pricing bands at the
	//     start of each billing period.
	TiersSpanPlan param.Field[bool] `json:"tiersSpanPlan"`
	//   - **DEBIT**. Default setting. The amount calculated using the Pricing is added
	//     to the bill as a debit.
	//
	//   - **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the
	//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
	//     will be capped at the total of other line items for the same Product.
	//
	//   - **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the
	//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
	//     will be capped at the total of other line items for the entire bill, which
	//     might include other Products the Account consumes.
	Type param.Field[PricingNewParamsType] `json:"type"`
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

func (r PricingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

//   - **DEBIT**. Default setting. The amount calculated using the Pricing is added
//     to the bill as a debit.
//
//   - **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the
//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
//     will be capped at the total of other line items for the same Product.
//
//   - **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the
//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
//     will be capped at the total of other line items for the entire bill, which
//     might include other Products the Account consumes.
type PricingNewParamsType string

const (
	PricingNewParamsTypeDebit         PricingNewParamsType = "DEBIT"
	PricingNewParamsTypeProductCredit PricingNewParamsType = "PRODUCT_CREDIT"
	PricingNewParamsTypeGlobalCredit  PricingNewParamsType = "GLOBAL_CREDIT"
)

func (r PricingNewParamsType) IsKnown() bool {
	switch r {
	case PricingNewParamsTypeDebit, PricingNewParamsTypeProductCredit, PricingNewParamsTypeGlobalCredit:
		return true
	}
	return false
}

type PricingGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type PricingUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID        param.Field[string]                    `path:"orgId,required"`
	PricingBands param.Field[[]shared.PricingBandParam] `json:"pricingBands,required"`
	// The start date _(in ISO-8601 format)_ for when the Pricing starts to be active
	// for the Plan of Plan Template._(Required)_
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// Optional Product ID this Pricing should be attributed to for accounting purposes
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// UUID of the Aggregation used to create the Pricing. Use this when creating a
	// Pricing for a segmented aggregation.
	AggregationID param.Field[string] `json:"aggregationId"`
	// Unique short code for the Pricing.
	Code param.Field[string] `json:"code"`
	// UUID of the Compound Aggregation used to create the Pricing.
	CompoundAggregationID param.Field[string] `json:"compoundAggregationId"`
	// Controls whether or not charge rates under a set of pricing bands configured for
	// a Pricing are applied according to each separate band or at the highest band
	// reached.
	//
	// _(Optional)_. The default value is **FALSE**.
	//
	//   - When TRUE, at billing charge rates are applied according to each separate
	//     band.
	//
	//   - When FALSE, at billing charge rates are applied according to highest band
	//     reached.
	//
	// **NOTE:** Use the `cumulative` parameter to create the type of Pricing you
	// require. For example, for Tiered Pricing set to **TRUE**; for Volume Pricing,
	// set to **FALSE**.
	Cumulative param.Field[bool] `json:"cumulative"`
	// Displayed on Bill line items.
	Description param.Field[string] `json:"description"`
	// The end date _(in ISO-8601 format)_ for when the Pricing ceases to be active for
	// the Plan or Plan Template.
	//
	// _(Optional)_ If not specified, the Pricing remains active indefinitely.
	EndDate param.Field[time.Time] `json:"endDate" format:"date-time"`
	// The minimum spend amount per billing cycle for end customer Accounts on a Plan
	// to which the Pricing is applied.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// The default value is **FALSE**.
	//
	// - When TRUE, minimum spend is billed at the start of each billing period.
	//
	// - When FALSE, minimum spend is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at Organization level for minimum spend
	// billing in arrears/in advance.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// Specify Prepayment/Balance overage pricing in pricing bands for the case of a
	// **Tiered** pricing structure.
	OveragePricingBands param.Field[[]shared.PricingBandParam] `json:"overagePricingBands"`
	// UUID of the Plan the Pricing is created for.
	PlanID param.Field[string] `json:"planId"`
	// UUID of the Plan Template the Pricing is created for.
	PlanTemplateID param.Field[string] `json:"planTemplateId"`
	// Specifies the segment value which you are defining a Pricing for using this
	// call:
	//
	//   - For each segment value defined on a Segmented Aggregation you must create a
	//     separate Pricing and use the appropriate `aggregationId` parameter for the
	//     call.
	//   - If you specify a segment value that has not been defined for the Aggregation,
	//     you'll receive an error.
	//   - If you've defined segment values for the Aggregation using a single wildcard
	//     or multiple wildcards, you can create Pricing for these wildcard segment
	//     values also.
	//
	// For more details on creating Pricings for segment values on a Segmented
	// Aggregation using this call, together with some examples, see the
	// [Using API Call to Create Segmented Pricings](https://www.m3ter.com/docs/guides/plans-and-pricing/pricing-plans/pricing-plans-using-segmented-aggregations#using-api-call-to-create-a-segmented-pricing)
	// in our User Documentation.
	Segment param.Field[map[string]string] `json:"segment"`
	// The default value is **FALSE**.
	//
	//   - If TRUE, usage accumulates over the entire period the priced Plan is active
	//     for the account, and is not reset for pricing band rates at the start of each
	//     billing period.
	//
	//   - If FALSE, usage does not accumulate, and is reset for pricing bands at the
	//     start of each billing period.
	TiersSpanPlan param.Field[bool] `json:"tiersSpanPlan"`
	//   - **DEBIT**. Default setting. The amount calculated using the Pricing is added
	//     to the bill as a debit.
	//
	//   - **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the
	//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
	//     will be capped at the total of other line items for the same Product.
	//
	//   - **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the
	//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
	//     will be capped at the total of other line items for the entire bill, which
	//     might include other Products the Account consumes.
	Type param.Field[PricingUpdateParamsType] `json:"type"`
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

func (r PricingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

//   - **DEBIT**. Default setting. The amount calculated using the Pricing is added
//     to the bill as a debit.
//
//   - **PRODUCT_CREDIT**. The amount calculated using the Pricing is added to the
//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
//     will be capped at the total of other line items for the same Product.
//
//   - **GLOBAL_CREDIT**. The amount calculated using the Pricing is added to the
//     bill as a credit _(negative amount)_. To prevent negative billing, the bill
//     will be capped at the total of other line items for the entire bill, which
//     might include other Products the Account consumes.
type PricingUpdateParamsType string

const (
	PricingUpdateParamsTypeDebit         PricingUpdateParamsType = "DEBIT"
	PricingUpdateParamsTypeProductCredit PricingUpdateParamsType = "PRODUCT_CREDIT"
	PricingUpdateParamsTypeGlobalCredit  PricingUpdateParamsType = "GLOBAL_CREDIT"
)

func (r PricingUpdateParamsType) IsKnown() bool {
	switch r {
	case PricingUpdateParamsTypeDebit, PricingUpdateParamsTypeProductCredit, PricingUpdateParamsTypeGlobalCredit:
		return true
	}
	return false
}

type PricingListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// UUID of the Aggregation to retrieve pricings for
	AggregationID param.Field[string] `query:"aggregationId"`
	// Date on which to retrieve active Pricings.
	Date param.Field[string] `query:"date"`
	// List of Pricing IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi-page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Pricings to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// UUID of the Plan to retrieve Pricings for.
	PlanID param.Field[string] `query:"planId"`
	// UUID of the PlanTemplate to retrieve Pricings for.
	PlanTemplateID param.Field[string] `query:"planTemplateId"`
}

// URLQuery serializes [PricingListParams]'s query parameters as `url.Values`.
func (r PricingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PricingDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
