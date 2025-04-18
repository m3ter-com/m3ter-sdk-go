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
	"github.com/m3ter-com/m3ter-sdk-go/shared"
)

// CounterPricingService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCounterPricingService] method instead.
type CounterPricingService struct {
	Options []option.RequestOption
}

// NewCounterPricingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCounterPricingService(opts ...option.RequestOption) (r *CounterPricingService) {
	r = &CounterPricingService{}
	r.Options = opts
	return
}

// Create a new CounterPricing.
//
// **Note:** Either `planId` or `planTemplateId` request parameters are required
// for this call to be valid. If you omit both, then you will receive a validation
// error.
func (r *CounterPricingService) New(ctx context.Context, params CounterPricingNewParams, opts ...option.RequestOption) (res *CounterPricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/counterpricings", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a CounterPricing for the given UUID.
func (r *CounterPricingService) Get(ctx context.Context, id string, query CounterPricingGetParams, opts ...option.RequestOption) (res *CounterPricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/counterpricings/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update CounterPricing for the given UUID.
//
// **Note:** Either `planId` or `planTemplateId` request parameters are required
// for this call to be valid. If you omit both, then you will receive a validation
// error.
func (r *CounterPricingService) Update(ctx context.Context, id string, params CounterPricingUpdateParams, opts ...option.RequestOption) (res *CounterPricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/counterpricings/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of CounterPricing entities filtered by date, Plan ID, Plan
// Template ID, or CounterPricing ID.
func (r *CounterPricingService) List(ctx context.Context, params CounterPricingListParams, opts ...option.RequestOption) (res *pagination.Cursor[CounterPricingResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/counterpricings", params.OrgID)
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

// Retrieve a list of CounterPricing entities filtered by date, Plan ID, Plan
// Template ID, or CounterPricing ID.
func (r *CounterPricingService) ListAutoPaging(ctx context.Context, params CounterPricingListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CounterPricingResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a CounterPricing for the given UUID.
func (r *CounterPricingService) Delete(ctx context.Context, id string, body CounterPricingDeleteParams, opts ...option.RequestOption) (res *CounterPricingResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/counterpricings/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CounterPricingResponse struct {
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
	// Unique short code for the Pricing.
	Code string `json:"code"`
	// UUID of the Counter used to create the pricing.
	CounterID string `json:"counterId"`
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
	// UUID of the Plan the Pricing is created for.
	PlanID string `json:"planId"`
	// UUID of the Plan Template the Pricing was created for.
	PlanTemplateID string               `json:"planTemplateId"`
	PricingBands   []shared.PricingBand `json:"pricingBands"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter adjustment credits are prorated and are billed according to
	//     the number of days in billing period.
	//
	//   - When FALSE, counter adjustment credits are not prorated and are billed for the
	//     entire billing period.
	ProRateAdjustmentCredit bool `json:"proRateAdjustmentCredit"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter adjustment debits are prorated and are billed according to
	//     the number of days in billing period.
	//
	//   - When FALSE, counter adjustment debits are not prorated and are billed for the
	//     entire billing period.
	ProRateAdjustmentDebit bool `json:"proRateAdjustmentDebit"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter running total charges are prorated and are billed according
	//     to the number of days in billing period.
	//
	//   - When FALSE, counter running total charges are not prorated and are billed for
	//     the entire billing period.
	ProRateRunningTotal bool `json:"proRateRunningTotal"`
	// The default value is **TRUE**.
	//
	// - When TRUE, running totals are billed at the start of each billing period.
	//
	// - When FALSE, running totals are billed at the end of each billing period.
	RunningTotalBillInAdvance bool `json:"runningTotalBillInAdvance"`
	// The start date _(in ISO-8601 format)_ for when the Pricing starts to be active
	// for the Plan of Plan Template.
	StartDate time.Time                  `json:"startDate" format:"date-time"`
	JSON      counterPricingResponseJSON `json:"-"`
}

// counterPricingResponseJSON contains the JSON metadata for the struct
// [CounterPricingResponse]
type counterPricingResponseJSON struct {
	ID                        apijson.Field
	Version                   apijson.Field
	AccountingProductID       apijson.Field
	Code                      apijson.Field
	CounterID                 apijson.Field
	CreatedBy                 apijson.Field
	Cumulative                apijson.Field
	Description               apijson.Field
	DtCreated                 apijson.Field
	DtLastModified            apijson.Field
	EndDate                   apijson.Field
	LastModifiedBy            apijson.Field
	PlanID                    apijson.Field
	PlanTemplateID            apijson.Field
	PricingBands              apijson.Field
	ProRateAdjustmentCredit   apijson.Field
	ProRateAdjustmentDebit    apijson.Field
	ProRateRunningTotal       apijson.Field
	RunningTotalBillInAdvance apijson.Field
	StartDate                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CounterPricingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r counterPricingResponseJSON) RawJSON() string {
	return r.raw
}

type CounterPricingNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// UUID of the Counter used to create the pricing.
	CounterID    param.Field[string]                    `json:"counterId,required"`
	PricingBands param.Field[[]shared.PricingBandParam] `json:"pricingBands,required"`
	// The start date _(in ISO-8601 format)_ for when the Pricing starts to be active
	// for the Plan of Plan Template._(Required)_
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// Optional Product ID this Pricing should be attributed to for accounting purposes
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Unique short code for the Pricing.
	Code param.Field[string] `json:"code"`
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
	// UUID of the Plan the Pricing is created for.
	PlanID param.Field[string] `json:"planId"`
	// UUID of the Plan Template the Pricing is created for.
	PlanTemplateID param.Field[string] `json:"planTemplateId"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter adjustment credits are prorated and are billed according to
	//     the number of days in billing period.
	//
	//   - When FALSE, counter adjustment credits are not prorated and are billed for the
	//     entire billing period.
	//
	// _(Optional)_.
	ProRateAdjustmentCredit param.Field[bool] `json:"proRateAdjustmentCredit"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter adjustment debits are prorated and are billed according to
	//     the number of days in billing period.
	//
	//   - When FALSE, counter adjustment debits are not prorated and are billed for the
	//     entire billing period.
	//
	// _(Optional)_.
	ProRateAdjustmentDebit param.Field[bool] `json:"proRateAdjustmentDebit"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter running total charges are prorated and are billed according
	//     to the number of days in billing period.
	//
	//   - When FALSE, counter running total charges are not prorated and are billed for
	//     the entire billing period.
	//
	// _(Optional)_.
	ProRateRunningTotal param.Field[bool] `json:"proRateRunningTotal"`
	// The default value is **TRUE**.
	//
	// - When TRUE, running totals are billed at the start of each billing period.
	//
	// - When FALSE, running totals are billed at the end of each billing period.
	//
	// _(Optional)_.
	RunningTotalBillInAdvance param.Field[bool] `json:"runningTotalBillInAdvance"`
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

func (r CounterPricingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterPricingGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type CounterPricingUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// UUID of the Counter used to create the pricing.
	CounterID    param.Field[string]                    `json:"counterId,required"`
	PricingBands param.Field[[]shared.PricingBandParam] `json:"pricingBands,required"`
	// The start date _(in ISO-8601 format)_ for when the Pricing starts to be active
	// for the Plan of Plan Template._(Required)_
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// Optional Product ID this Pricing should be attributed to for accounting purposes
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Unique short code for the Pricing.
	Code param.Field[string] `json:"code"`
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
	// UUID of the Plan the Pricing is created for.
	PlanID param.Field[string] `json:"planId"`
	// UUID of the Plan Template the Pricing is created for.
	PlanTemplateID param.Field[string] `json:"planTemplateId"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter adjustment credits are prorated and are billed according to
	//     the number of days in billing period.
	//
	//   - When FALSE, counter adjustment credits are not prorated and are billed for the
	//     entire billing period.
	//
	// _(Optional)_.
	ProRateAdjustmentCredit param.Field[bool] `json:"proRateAdjustmentCredit"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter adjustment debits are prorated and are billed according to
	//     the number of days in billing period.
	//
	//   - When FALSE, counter adjustment debits are not prorated and are billed for the
	//     entire billing period.
	//
	// _(Optional)_.
	ProRateAdjustmentDebit param.Field[bool] `json:"proRateAdjustmentDebit"`
	// The default value is **TRUE**.
	//
	//   - When TRUE, counter running total charges are prorated and are billed according
	//     to the number of days in billing period.
	//
	//   - When FALSE, counter running total charges are not prorated and are billed for
	//     the entire billing period.
	//
	// _(Optional)_.
	ProRateRunningTotal param.Field[bool] `json:"proRateRunningTotal"`
	// The default value is **TRUE**.
	//
	// - When TRUE, running totals are billed at the start of each billing period.
	//
	// - When FALSE, running totals are billed at the end of each billing period.
	//
	// _(Optional)_.
	RunningTotalBillInAdvance param.Field[bool] `json:"runningTotalBillInAdvance"`
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

func (r CounterPricingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterPricingListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Date on which to retrieve active CounterPricings.
	Date param.Field[string] `query:"date"`
	// List of CounterPricing IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of CounterPricings to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// UUID of the Plan to retrieve CounterPricings for.
	PlanID param.Field[string] `query:"planId"`
	// UUID of the Plan Template to retrieve CounterPricings for.
	PlanTemplateID param.Field[string] `query:"planTemplateId"`
}

// URLQuery serializes [CounterPricingListParams]'s query parameters as
// `url.Values`.
func (r CounterPricingListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CounterPricingDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
