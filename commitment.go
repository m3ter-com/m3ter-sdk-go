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

// CommitmentService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommitmentService] method instead.
type CommitmentService struct {
	Options []option.RequestOption
}

// NewCommitmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCommitmentService(opts ...option.RequestOption) (r *CommitmentService) {
	r = &CommitmentService{}
	r.Options = opts
	return
}

// Create a new Commitment.
//
// Creates a new Commitment for an Organization. The request body must include all
// the necessary details such as the agreed amount, overage surcharge percentage,
// and the associated account and product details.
//
// **Note:** If some of the agreed Commitment amount remains unpaid at the start of
// an end-customer contract period, when you create a Commitment for an Account you
// can set up billing for the outstanding amount in one of two ways:
//
//   - Select a Product _Plan to bill with_. Use the `billingPlanId` request
//     parameter to select the Plan used for billing.
//   - Define a _schedule of billing dates_. Omit a `billingPlanId` and use the
//     `feeDates` request parameter to define a precise schedule of bill dates and
//     amounts.
func (r *CommitmentService) New(ctx context.Context, params CommitmentNewParams, opts ...option.RequestOption) (res *CommitmentResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/commitments", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a specific Commitment.
//
// Retrieve the details of the Commitment with the given UUID. It provides
// comprehensive information about the Commitment, such as the agreed amount,
// overage surcharge percentage, and other related details.
func (r *CommitmentService) Get(ctx context.Context, id string, query CommitmentGetParams, opts ...option.RequestOption) (res *CommitmentResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/commitments/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a specific Commitment.
//
// Update the details of the Commitment with the given UUID. Use this endpoint to
// adjust Commitment parameters such as the fixed amount, overage surcharge
// percentage, or associated contract details.
func (r *CommitmentService) Update(ctx context.Context, id string, params CommitmentUpdateParams, opts ...option.RequestOption) (res *CommitmentResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/commitments/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Commitments.
//
// Retrieves a list of all Commitments associated with an Organization. This
// endpoint supports pagination and includes various query parameters to filter the
// Commitments based on Account, Product, date, and end dates.
func (r *CommitmentService) List(ctx context.Context, params CommitmentListParams, opts ...option.RequestOption) (res *pagination.Cursor[CommitmentResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/commitments", params.OrgID)
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

// Retrieve a list of Commitments.
//
// Retrieves a list of all Commitments associated with an Organization. This
// endpoint supports pagination and includes various query parameters to filter the
// Commitments based on Account, Product, date, and end dates.
func (r *CommitmentService) ListAutoPaging(ctx context.Context, params CommitmentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CommitmentResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Remove a specific Commitment.
//
// Deletes the Commitment with the given UUID. Use this endpoint when a Commitment
// is no longer valid or needs to be removed from the system.
func (r *CommitmentService) Delete(ctx context.Context, id string, body CommitmentDeleteParams, opts ...option.RequestOption) (res *CommitmentResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/commitments/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Search for commitment entities.
//
// This endpoint executes a search query for Commitments based on the user
// specified search criteria. The search query is customizable, allowing for
// complex nested conditions and sorting. The returned list of Commitments can be
// paginated for easier management.
func (r *CommitmentService) Search(ctx context.Context, params CommitmentSearchParams, opts ...option.RequestOption) (res *CommitmentSearchResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/commitments/search", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type CommitmentFee struct {
	Amount                 float64           `json:"amount,required"`
	Date                   time.Time         `json:"date,required" format:"date"`
	ServicePeriodEndDate   time.Time         `json:"servicePeriodEndDate,required" format:"date-time"`
	ServicePeriodStartDate time.Time         `json:"servicePeriodStartDate,required" format:"date-time"`
	JSON                   commitmentFeeJSON `json:"-"`
}

// commitmentFeeJSON contains the JSON metadata for the struct [CommitmentFee]
type commitmentFeeJSON struct {
	Amount                 apijson.Field
	Date                   apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CommitmentFee) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitmentFeeJSON) RawJSON() string {
	return r.raw
}

type CommitmentFeeParam struct {
	Amount                 param.Field[float64]   `json:"amount,required"`
	Date                   param.Field[time.Time] `json:"date,required" format:"date"`
	ServicePeriodEndDate   param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
}

func (r CommitmentFeeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommitmentResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The unique identifier (UUID) for the end customer Account the Commitment is
	// added to.
	AccountID string `json:"accountId"`
	// The unique identifier (UUID) for the Product linked to the Commitment for
	// accounting purposes.
	AccountingProductID string `json:"accountingProductId"`
	// The total amount that the customer has committed to pay.
	Amount float64 `json:"amount"`
	// The amount to be billed in the first invoice.
	AmountFirstBill float64 `json:"amountFirstBill"`
	// The amount that the customer has already paid upfront at the start of the
	// Commitment service period.
	AmountPrePaid float64 `json:"amountPrePaid"`
	// The total amount of the Commitment that the customer has spent so far.
	AmountSpent float64 `json:"amountSpent"`
	// The starting date _(in ISO-8601 date format)_ from which the billing cycles are
	// calculated.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// How often the Commitment fees are applied to bills. For example, if the plan
	// being used to bill for Commitment fees is set to issue bills every three months
	// and the `billingInterval` is set to 2, then the Commitment fees are applied
	// every six months.
	BillingInterval int64 `json:"billingInterval"`
	// The offset for when the Commitment fees are first applied to bills on the
	// Account. For example, if bills are issued every three months and the
	// `billingOffset` is 0, then the charge is applied to the first bill (at three
	// months); if set to 1, it's applied to the next bill (at six months), and so on.
	BillingOffset int64 `json:"billingOffset"`
	// The unique identifier (UUID) for the Product Plan used for billing Commitment
	// fees due.
	BillingPlanID string `json:"billingPlanId"`
	// If the Account is either a Parent or a Child Account, this specifies the Account
	// hierarchy billing mode. The mode determines how billing will be handled and
	// shown on bills for charges due on the Parent Account, and charges due on Child
	// Accounts:
	//
	// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
	//
	// - **Parent Summary** - single bill line item for all Accounts.
	//
	// - **Child** - the Child Account is billed.
	ChildBillingMode CommitmentResponseChildBillingMode `json:"childBillingMode"`
	// A boolean value indicating whether the Commitment fee is billed in advance
	// _(start of each billing period)_ or arrears _(end of each billing period)_.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	CommitmentFeeBillInAdvance bool `json:"commitmentFeeBillInAdvance"`
	// A textual description of the Commitment fee.
	CommitmentFeeDescription string `json:"commitmentFeeDescription"`
	// A textual description of the Commitment usage.
	CommitmentUsageDescription string `json:"commitmentUsageDescription"`
	// The unique identifier (UUID) for a Contract you've created for the Account and
	// to which the Commitment has been added.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) of the user who created this Commitment.
	CreatedBy string `json:"createdBy"`
	// The currency used for the Commitment. For example, 'USD'.
	Currency string `json:"currency"`
	// Optional Product ID this Commitment's consumptions should be attributed to for
	// accounting purposes.
	DrawdownsAccountingProductID string `json:"drawdownsAccountingProductId"`
	// The date and time _(in ISO-8601 format)_ when the Commitment was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the Commitment was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The end date of the Commitment period in ISO-8601 format.
	EndDate time.Time `json:"endDate" format:"date"`
	// Used for billing any outstanding Commitment fees _on a schedule_.
	//
	// An array defining a series of bill dates and amounts covering specified service
	// periods:
	//
	//   - `date` - the billing date _(in ISO-8601 format)_.
	//   - `amount` - the billed amount.
	//   - `servicePeriodStartDate` and `servicePeriodEndDate` - defines the service
	//     period the bill covers _(in ISO-8601 format)_.
	FeeDates []CommitmentFee `json:"feeDates"`
	// Optional Product ID this Commitment's fees should be attributed to for
	// accounting purposes.
	FeesAccountingProductID string `json:"feesAccountingProductId"`
	// The unique identifier (UUID) of the user who last modified this Commitment.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Specifies the line item charge types that can draw-down at billing against the
	// Commitment amount. Options are:
	//
	// - `MINIMUM_SPEND`
	// - `STANDING_CHARGE`
	// - `USAGE`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	LineItemTypes []CommitmentResponseLineItemType `json:"lineItemTypes"`
	// A textual description of the overage charges.
	OverageDescription string `json:"overageDescription"`
	// The percentage surcharge applied to the usage charges that exceed the Commitment
	// amount.
	OverageSurchargePercent float64 `json:"overageSurchargePercent"`
	// A list of unique identifiers (UUIDs) for Products the Account consumes. Charges
	// due for these Products will be made available for draw-down against the
	// Commitment.
	//
	// **Note:** If not used, then charges due for all Products the Account consumes
	// will be made available for draw-down against the Commitment.
	ProductIDs []string `json:"productIds"`
	// A boolean value indicating whether the overage usage is billed separately or
	// together. If overage usage is separated and a Commitment amount has been
	// consumed by an Account, any subsequent line items on Bills against the Account
	// for usage will show as separate "overage usage" charges, not simply as "usage"
	// charges:
	//
	// - **TRUE** - billed separately.
	// - **FALSE** - billed together.
	SeparateOverageUsage bool `json:"separateOverageUsage"`
	// The start date of the Commitment period in ISO-8601 format.
	StartDate time.Time `json:"startDate" format:"date"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                  `json:"version"`
	JSON    commitmentResponseJSON `json:"-"`
}

// commitmentResponseJSON contains the JSON metadata for the struct
// [CommitmentResponse]
type commitmentResponseJSON struct {
	ID                           apijson.Field
	AccountID                    apijson.Field
	AccountingProductID          apijson.Field
	Amount                       apijson.Field
	AmountFirstBill              apijson.Field
	AmountPrePaid                apijson.Field
	AmountSpent                  apijson.Field
	BillEpoch                    apijson.Field
	BillingInterval              apijson.Field
	BillingOffset                apijson.Field
	BillingPlanID                apijson.Field
	ChildBillingMode             apijson.Field
	CommitmentFeeBillInAdvance   apijson.Field
	CommitmentFeeDescription     apijson.Field
	CommitmentUsageDescription   apijson.Field
	ContractID                   apijson.Field
	CreatedBy                    apijson.Field
	Currency                     apijson.Field
	DrawdownsAccountingProductID apijson.Field
	DtCreated                    apijson.Field
	DtLastModified               apijson.Field
	EndDate                      apijson.Field
	FeeDates                     apijson.Field
	FeesAccountingProductID      apijson.Field
	LastModifiedBy               apijson.Field
	LineItemTypes                apijson.Field
	OverageDescription           apijson.Field
	OverageSurchargePercent      apijson.Field
	ProductIDs                   apijson.Field
	SeparateOverageUsage         apijson.Field
	StartDate                    apijson.Field
	Version                      apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *CommitmentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitmentResponseJSON) RawJSON() string {
	return r.raw
}

// If the Account is either a Parent or a Child Account, this specifies the Account
// hierarchy billing mode. The mode determines how billing will be handled and
// shown on bills for charges due on the Parent Account, and charges due on Child
// Accounts:
//
// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
//
// - **Parent Summary** - single bill line item for all Accounts.
//
// - **Child** - the Child Account is billed.
type CommitmentResponseChildBillingMode string

const (
	CommitmentResponseChildBillingModeParentSummary   CommitmentResponseChildBillingMode = "PARENT_SUMMARY"
	CommitmentResponseChildBillingModeParentBreakdown CommitmentResponseChildBillingMode = "PARENT_BREAKDOWN"
	CommitmentResponseChildBillingModeChild           CommitmentResponseChildBillingMode = "CHILD"
)

func (r CommitmentResponseChildBillingMode) IsKnown() bool {
	switch r {
	case CommitmentResponseChildBillingModeParentSummary, CommitmentResponseChildBillingModeParentBreakdown, CommitmentResponseChildBillingModeChild:
		return true
	}
	return false
}

// Available line item types for Commitments
type CommitmentResponseLineItemType string

const (
	CommitmentResponseLineItemTypeStandingCharge            CommitmentResponseLineItemType = "STANDING_CHARGE"
	CommitmentResponseLineItemTypeUsage                     CommitmentResponseLineItemType = "USAGE"
	CommitmentResponseLineItemTypeMinimumSpend              CommitmentResponseLineItemType = "MINIMUM_SPEND"
	CommitmentResponseLineItemTypeCounterRunningTotalCharge CommitmentResponseLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	CommitmentResponseLineItemTypeCounterAdjustmentDebit    CommitmentResponseLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
)

func (r CommitmentResponseLineItemType) IsKnown() bool {
	switch r {
	case CommitmentResponseLineItemTypeStandingCharge, CommitmentResponseLineItemTypeUsage, CommitmentResponseLineItemTypeMinimumSpend, CommitmentResponseLineItemTypeCounterRunningTotalCharge, CommitmentResponseLineItemTypeCounterAdjustmentDebit:
		return true
	}
	return false
}

type CommitmentSearchResponse struct {
	Data      []CommitmentResponse         `json:"data"`
	NextToken string                       `json:"nextToken"`
	JSON      commitmentSearchResponseJSON `json:"-"`
}

// commitmentSearchResponseJSON contains the JSON metadata for the struct
// [CommitmentSearchResponse]
type commitmentSearchResponseJSON struct {
	Data        apijson.Field
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitmentSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitmentSearchResponseJSON) RawJSON() string {
	return r.raw
}

type CommitmentNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) for the end customer Account the Commitment is
	// added to.
	AccountID param.Field[string] `json:"accountId,required"`
	// The total amount that the customer has committed to pay.
	Amount param.Field[float64] `json:"amount,required"`
	// The currency used for the Commitment. For example: USD.
	Currency param.Field[string] `json:"currency,required"`
	// The end date of the Commitment period in ISO-8601 format.
	//
	// **Note:** End date is exclusive - if you set an end date of June 1st 2022, then
	// the Commitment ceases to be active for the Account at midnight on May 31st 2022,
	// and any Prepayment fees due are calculated up to that point in time, NOT up to
	// midnight on June 1st
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date"`
	// The start date of the Commitment period in ISO-8601 format.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date"`
	// The unique identifier (UUID) for the Product linked to the Commitment for
	// accounting purposes. _(Optional)_
	//
	// **NOTE:** If you're planning to set up an integration for sending Bills to an
	// external accounts receivable system, please check requirements for your chosen
	// system. Some systems, such as NetSuite, require a Product to be linked with any
	// Bill line items associated with Account Commitments, and the integration will
	// fail if this is not present
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// The amount to be billed in the first invoice.
	AmountFirstBill param.Field[float64] `json:"amountFirstBill"`
	// The amount that the customer has already paid upfront at the start of the
	// Commitment service period.
	AmountPrePaid param.Field[float64] `json:"amountPrePaid"`
	// The starting date _(in ISO-8601 date format)_ from which the billing cycles are
	// calculated.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
	// How often the Commitment fees are applied to bills. For example, if the plan
	// being used to bill for Commitment fees is set to issue bills every three months
	// and the `billingInterval` is set to 2, then the Commitment fees are applied
	// every six months.
	BillingInterval param.Field[int64] `json:"billingInterval"`
	// Defines an offset for when the Commitment fees are first applied to bills on the
	// Account. For example, if bills are issued every three months and the
	// `billingOffset` is 0, then the charge is applied to the first bill (at three
	// months); if set to 1, it's applied to the next bill (at six months), and so on.
	BillingOffset param.Field[int64] `json:"billingOffset"`
	// The unique identifier (UUID) for the Product Plan used for billing Commitment
	// fees due.
	BillingPlanID param.Field[string] `json:"billingPlanId"`
	// If the Account is either a Parent or a Child Account, this specifies the Account
	// hierarchy billing mode. The mode determines how billing will be handled and
	// shown on bills for charges due on the Parent Account, and charges due on Child
	// Accounts:
	//
	// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
	//
	// - **Parent Summary** - single bill line item for all Accounts.
	//
	// - **Child** - the Child Account is billed.
	ChildBillingMode param.Field[CommitmentNewParamsChildBillingMode] `json:"childBillingMode"`
	// A boolean value indicating whether the Commitment fee is billed in advance
	// _(start of each billing period)_ or arrears _(end of each billing period)_.
	//
	// If no value is supplied, then the Organization Configuration value is used.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	CommitmentFeeBillInAdvance param.Field[bool] `json:"commitmentFeeBillInAdvance"`
	// A textual description of the Commitment fee.
	CommitmentFeeDescription param.Field[string] `json:"commitmentFeeDescription"`
	// A textual description of the Commitment usage.
	CommitmentUsageDescription param.Field[string] `json:"commitmentUsageDescription"`
	// The unique identifier (UUID) for a Contract you've created for the Account -
	// used to add the Commitment to this Contract.
	//
	// **Note:** If you associate the Commitment with a Contract you must ensure the
	// Account Plan attached to the Account has the same Contract associated with it.
	// If the Account Plan Contract and Commitment Contract do not match, then at
	// billing the Commitment amount will not be drawn-down against.
	ContractID param.Field[string] `json:"contractId"`
	// Optional Product ID this Commitment's consumptions should be attributed to for
	// accounting purposes.
	DrawdownsAccountingProductID param.Field[string] `json:"drawdownsAccountingProductId"`
	// Used for billing any outstanding Commitment fees _on a schedule_.
	//
	// Create an array to define a series of bill dates and amounts covering specified
	// service periods:
	//
	//   - `date` - the billing date _(in ISO-8601 format)_.
	//   - `amount` - the billed amount.
	//   - `servicePeriodStartDate` and `servicePeriodEndDate` - defines the service
	//     period the bill covers _(in ISO-8601 format)_.
	//
	// **Notes:**
	//
	//   - If you try to set `servicePeriodStartDate` _after_ `servicePeriodEndDate`,
	//     you'll receive an error.
	//   - You can set `servicePeriodStartDate` and `servicePeriodEndDate` to the _same
	//     date_ without receiving an error, but _please be sure_ your Commitment billing
	//     use case requires this.
	FeeDates param.Field[[]CommitmentFeeParam] `json:"feeDates"`
	// Optional Product ID this Commitment's fees should be attributed to for
	// accounting purposes.
	FeesAccountingProductID param.Field[string] `json:"feesAccountingProductId"`
	// Specify the line item charge types that can draw-down at billing against the
	// Commitment amount. Options are:
	//
	// - `MINIMUM_SPEND`
	// - `STANDING_CHARGE`
	// - `USAGE`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	//
	// **NOTE:** If no charge types are specified, by default _all types_ can draw-down
	// against the Commitment amount at billing.
	LineItemTypes param.Field[[]CommitmentNewParamsLineItemType] `json:"lineItemTypes"`
	// A textual description of the overage charges.
	OverageDescription param.Field[string] `json:"overageDescription"`
	// The percentage surcharge applied to usage charges that exceed the Commitment
	// amount.
	//
	// **Note:** You can enter a _negative percentage_ if you want to give a discount
	// rate for usage to end customers who exceed their Commitment amount
	OverageSurchargePercent param.Field[float64] `json:"overageSurchargePercent"`
	// A list of unique identifiers (UUIDs) for Products the Account consumes. Charges
	// due for these Products will be made available for draw-down against the
	// Commitment.
	//
	// **Note:** If not used, then charges due for all Products the Account consumes
	// will be made available for draw-down against the Commitment.
	ProductIDs param.Field[[]string] `json:"productIds"`
	// A boolean value indicating whether the overage usage is billed separately or
	// together. If overage usage is separated and a Commitment amount has been
	// consumed by an Account, any subsequent line items on Bills against the Account
	// for usage will show as separate "overage usage" charges, not simply as "usage"
	// charges:
	//
	// - **TRUE** - billed separately.
	// - **FALSE** - billed together.
	//
	// **Notes:**
	//
	//   - Can be used only if no value or 0 has been defined for the
	//     `overageSurchargePercent` parameter. If you try to separate overage usage when
	//     a value other than 0 has been defined for `overageSurchargePercent`, you'll
	//     receive an error.
	//   - If a priced Plan is used to bill any outstanding Commitment fees due and the
	//     Plan is set up with overage pricing on a _tiered pricing structure_ and you
	//     enable separate bill line items for overage usage, then overage usage charges
	//     will be rated according to the overage pricing defined for the tiered pricing
	//     on the Plan.
	SeparateOverageUsage param.Field[bool] `json:"separateOverageUsage"`
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

func (r CommitmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Account is either a Parent or a Child Account, this specifies the Account
// hierarchy billing mode. The mode determines how billing will be handled and
// shown on bills for charges due on the Parent Account, and charges due on Child
// Accounts:
//
// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
//
// - **Parent Summary** - single bill line item for all Accounts.
//
// - **Child** - the Child Account is billed.
type CommitmentNewParamsChildBillingMode string

const (
	CommitmentNewParamsChildBillingModeParentSummary   CommitmentNewParamsChildBillingMode = "PARENT_SUMMARY"
	CommitmentNewParamsChildBillingModeParentBreakdown CommitmentNewParamsChildBillingMode = "PARENT_BREAKDOWN"
	CommitmentNewParamsChildBillingModeChild           CommitmentNewParamsChildBillingMode = "CHILD"
)

func (r CommitmentNewParamsChildBillingMode) IsKnown() bool {
	switch r {
	case CommitmentNewParamsChildBillingModeParentSummary, CommitmentNewParamsChildBillingModeParentBreakdown, CommitmentNewParamsChildBillingModeChild:
		return true
	}
	return false
}

// Available line item types for Commitments
type CommitmentNewParamsLineItemType string

const (
	CommitmentNewParamsLineItemTypeStandingCharge            CommitmentNewParamsLineItemType = "STANDING_CHARGE"
	CommitmentNewParamsLineItemTypeUsage                     CommitmentNewParamsLineItemType = "USAGE"
	CommitmentNewParamsLineItemTypeMinimumSpend              CommitmentNewParamsLineItemType = "MINIMUM_SPEND"
	CommitmentNewParamsLineItemTypeCounterRunningTotalCharge CommitmentNewParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	CommitmentNewParamsLineItemTypeCounterAdjustmentDebit    CommitmentNewParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
)

func (r CommitmentNewParamsLineItemType) IsKnown() bool {
	switch r {
	case CommitmentNewParamsLineItemTypeStandingCharge, CommitmentNewParamsLineItemTypeUsage, CommitmentNewParamsLineItemTypeMinimumSpend, CommitmentNewParamsLineItemTypeCounterRunningTotalCharge, CommitmentNewParamsLineItemTypeCounterAdjustmentDebit:
		return true
	}
	return false
}

type CommitmentGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type CommitmentUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) for the end customer Account the Commitment is
	// added to.
	AccountID param.Field[string] `json:"accountId,required"`
	// The total amount that the customer has committed to pay.
	Amount param.Field[float64] `json:"amount,required"`
	// The currency used for the Commitment. For example: USD.
	Currency param.Field[string] `json:"currency,required"`
	// The end date of the Commitment period in ISO-8601 format.
	//
	// **Note:** End date is exclusive - if you set an end date of June 1st 2022, then
	// the Commitment ceases to be active for the Account at midnight on May 31st 2022,
	// and any Prepayment fees due are calculated up to that point in time, NOT up to
	// midnight on June 1st
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date"`
	// The start date of the Commitment period in ISO-8601 format.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date"`
	// The unique identifier (UUID) for the Product linked to the Commitment for
	// accounting purposes. _(Optional)_
	//
	// **NOTE:** If you're planning to set up an integration for sending Bills to an
	// external accounts receivable system, please check requirements for your chosen
	// system. Some systems, such as NetSuite, require a Product to be linked with any
	// Bill line items associated with Account Commitments, and the integration will
	// fail if this is not present
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// The amount to be billed in the first invoice.
	AmountFirstBill param.Field[float64] `json:"amountFirstBill"`
	// The amount that the customer has already paid upfront at the start of the
	// Commitment service period.
	AmountPrePaid param.Field[float64] `json:"amountPrePaid"`
	// The starting date _(in ISO-8601 date format)_ from which the billing cycles are
	// calculated.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
	// How often the Commitment fees are applied to bills. For example, if the plan
	// being used to bill for Commitment fees is set to issue bills every three months
	// and the `billingInterval` is set to 2, then the Commitment fees are applied
	// every six months.
	BillingInterval param.Field[int64] `json:"billingInterval"`
	// Defines an offset for when the Commitment fees are first applied to bills on the
	// Account. For example, if bills are issued every three months and the
	// `billingOffset` is 0, then the charge is applied to the first bill (at three
	// months); if set to 1, it's applied to the next bill (at six months), and so on.
	BillingOffset param.Field[int64] `json:"billingOffset"`
	// The unique identifier (UUID) for the Product Plan used for billing Commitment
	// fees due.
	BillingPlanID param.Field[string] `json:"billingPlanId"`
	// If the Account is either a Parent or a Child Account, this specifies the Account
	// hierarchy billing mode. The mode determines how billing will be handled and
	// shown on bills for charges due on the Parent Account, and charges due on Child
	// Accounts:
	//
	// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
	//
	// - **Parent Summary** - single bill line item for all Accounts.
	//
	// - **Child** - the Child Account is billed.
	ChildBillingMode param.Field[CommitmentUpdateParamsChildBillingMode] `json:"childBillingMode"`
	// A boolean value indicating whether the Commitment fee is billed in advance
	// _(start of each billing period)_ or arrears _(end of each billing period)_.
	//
	// If no value is supplied, then the Organization Configuration value is used.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	CommitmentFeeBillInAdvance param.Field[bool] `json:"commitmentFeeBillInAdvance"`
	// A textual description of the Commitment fee.
	CommitmentFeeDescription param.Field[string] `json:"commitmentFeeDescription"`
	// A textual description of the Commitment usage.
	CommitmentUsageDescription param.Field[string] `json:"commitmentUsageDescription"`
	// The unique identifier (UUID) for a Contract you've created for the Account -
	// used to add the Commitment to this Contract.
	//
	// **Note:** If you associate the Commitment with a Contract you must ensure the
	// Account Plan attached to the Account has the same Contract associated with it.
	// If the Account Plan Contract and Commitment Contract do not match, then at
	// billing the Commitment amount will not be drawn-down against.
	ContractID param.Field[string] `json:"contractId"`
	// Optional Product ID this Commitment's consumptions should be attributed to for
	// accounting purposes.
	DrawdownsAccountingProductID param.Field[string] `json:"drawdownsAccountingProductId"`
	// Used for billing any outstanding Commitment fees _on a schedule_.
	//
	// Create an array to define a series of bill dates and amounts covering specified
	// service periods:
	//
	//   - `date` - the billing date _(in ISO-8601 format)_.
	//   - `amount` - the billed amount.
	//   - `servicePeriodStartDate` and `servicePeriodEndDate` - defines the service
	//     period the bill covers _(in ISO-8601 format)_.
	//
	// **Notes:**
	//
	//   - If you try to set `servicePeriodStartDate` _after_ `servicePeriodEndDate`,
	//     you'll receive an error.
	//   - You can set `servicePeriodStartDate` and `servicePeriodEndDate` to the _same
	//     date_ without receiving an error, but _please be sure_ your Commitment billing
	//     use case requires this.
	FeeDates param.Field[[]CommitmentFeeParam] `json:"feeDates"`
	// Optional Product ID this Commitment's fees should be attributed to for
	// accounting purposes.
	FeesAccountingProductID param.Field[string] `json:"feesAccountingProductId"`
	// Specify the line item charge types that can draw-down at billing against the
	// Commitment amount. Options are:
	//
	// - `MINIMUM_SPEND`
	// - `STANDING_CHARGE`
	// - `USAGE`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	//
	// **NOTE:** If no charge types are specified, by default _all types_ can draw-down
	// against the Commitment amount at billing.
	LineItemTypes param.Field[[]CommitmentUpdateParamsLineItemType] `json:"lineItemTypes"`
	// A textual description of the overage charges.
	OverageDescription param.Field[string] `json:"overageDescription"`
	// The percentage surcharge applied to usage charges that exceed the Commitment
	// amount.
	//
	// **Note:** You can enter a _negative percentage_ if you want to give a discount
	// rate for usage to end customers who exceed their Commitment amount
	OverageSurchargePercent param.Field[float64] `json:"overageSurchargePercent"`
	// A list of unique identifiers (UUIDs) for Products the Account consumes. Charges
	// due for these Products will be made available for draw-down against the
	// Commitment.
	//
	// **Note:** If not used, then charges due for all Products the Account consumes
	// will be made available for draw-down against the Commitment.
	ProductIDs param.Field[[]string] `json:"productIds"`
	// A boolean value indicating whether the overage usage is billed separately or
	// together. If overage usage is separated and a Commitment amount has been
	// consumed by an Account, any subsequent line items on Bills against the Account
	// for usage will show as separate "overage usage" charges, not simply as "usage"
	// charges:
	//
	// - **TRUE** - billed separately.
	// - **FALSE** - billed together.
	//
	// **Notes:**
	//
	//   - Can be used only if no value or 0 has been defined for the
	//     `overageSurchargePercent` parameter. If you try to separate overage usage when
	//     a value other than 0 has been defined for `overageSurchargePercent`, you'll
	//     receive an error.
	//   - If a priced Plan is used to bill any outstanding Commitment fees due and the
	//     Plan is set up with overage pricing on a _tiered pricing structure_ and you
	//     enable separate bill line items for overage usage, then overage usage charges
	//     will be rated according to the overage pricing defined for the tiered pricing
	//     on the Plan.
	SeparateOverageUsage param.Field[bool] `json:"separateOverageUsage"`
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

func (r CommitmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Account is either a Parent or a Child Account, this specifies the Account
// hierarchy billing mode. The mode determines how billing will be handled and
// shown on bills for charges due on the Parent Account, and charges due on Child
// Accounts:
//
// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
//
// - **Parent Summary** - single bill line item for all Accounts.
//
// - **Child** - the Child Account is billed.
type CommitmentUpdateParamsChildBillingMode string

const (
	CommitmentUpdateParamsChildBillingModeParentSummary   CommitmentUpdateParamsChildBillingMode = "PARENT_SUMMARY"
	CommitmentUpdateParamsChildBillingModeParentBreakdown CommitmentUpdateParamsChildBillingMode = "PARENT_BREAKDOWN"
	CommitmentUpdateParamsChildBillingModeChild           CommitmentUpdateParamsChildBillingMode = "CHILD"
)

func (r CommitmentUpdateParamsChildBillingMode) IsKnown() bool {
	switch r {
	case CommitmentUpdateParamsChildBillingModeParentSummary, CommitmentUpdateParamsChildBillingModeParentBreakdown, CommitmentUpdateParamsChildBillingModeChild:
		return true
	}
	return false
}

// Available line item types for Commitments
type CommitmentUpdateParamsLineItemType string

const (
	CommitmentUpdateParamsLineItemTypeStandingCharge            CommitmentUpdateParamsLineItemType = "STANDING_CHARGE"
	CommitmentUpdateParamsLineItemTypeUsage                     CommitmentUpdateParamsLineItemType = "USAGE"
	CommitmentUpdateParamsLineItemTypeMinimumSpend              CommitmentUpdateParamsLineItemType = "MINIMUM_SPEND"
	CommitmentUpdateParamsLineItemTypeCounterRunningTotalCharge CommitmentUpdateParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	CommitmentUpdateParamsLineItemTypeCounterAdjustmentDebit    CommitmentUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
)

func (r CommitmentUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case CommitmentUpdateParamsLineItemTypeStandingCharge, CommitmentUpdateParamsLineItemTypeUsage, CommitmentUpdateParamsLineItemTypeMinimumSpend, CommitmentUpdateParamsLineItemTypeCounterRunningTotalCharge, CommitmentUpdateParamsLineItemTypeCounterAdjustmentDebit:
		return true
	}
	return false
}

type CommitmentListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) for the Account. This parameter helps filter the
	// Commitments related to a specific end-customer Account.
	AccountID  param.Field[string] `query:"accountId"`
	ContractID param.Field[string] `query:"contractId"`
	// A date _(in ISO-8601 format)_ to filter Commitments which are active on this
	// specific date.
	Date param.Field[string] `query:"date"`
	// A date _(in ISO-8601 format)_ used to filter Commitments. Only Commitments with
	// end dates before this date will be included.
	EndDateEnd param.Field[string] `query:"endDateEnd"`
	// A date _(in ISO-8601 format)_ used to filter Commitments. Only Commitments with
	// end dates on or after this date will be included.
	EndDateStart param.Field[string] `query:"endDateStart"`
	// A list of unique identifiers (UUIDs) for the Commitments to retrieve. Use this
	// to fetch specific Commitments in a single request.
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Commitments in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Commitments to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// The unique identifier (UUID) for the Product. This parameter helps filter the
	// Commitments related to a specific Product.
	ProductID param.Field[string] `query:"productId"`
}

// URLQuery serializes [CommitmentListParams]'s query parameters as `url.Values`.
func (r CommitmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CommitmentDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type CommitmentSearchParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// `fromDocument` for multi page retrievals.
	FromDocument param.Field[int64] `query:"fromDocument"`
	// Search Operator to be used while querying search.
	Operator param.Field[CommitmentSearchParamsOperator] `query:"operator"`
	// Number of Commitments to retrieve per page.
	//
	// **NOTE:** If not defined, default is 10.
	PageSize param.Field[int64] `query:"pageSize"`
	// Query for data using special syntax:
	//
	// - Query parameters should be delimited using $ (dollar sign).
	// - Allowed comparators are:
	//   - (greater than) >
	//   - (greater than or equal to) >=
	//   - (equal to) :
	//   - (less than) <
	//   - (less than or equal to) <=
	//   - (match phrase/prefix) ~
	//   - Allowed parameters: startDate, endDate, contractId, accountId, productId,
	//     productIds, id, createdBy, dtCreated, lastModifiedBy, ids.
	//   - Query example:
	//   - searchQuery=startDate>2023-01-01$accountId:062085ab-a301-4f21-a081-411020864452.
	//   - This query is translated into: find commitments where the startDate is older
	//     than 2023-01-01 AND the accountId is equal to
	//     062085ab-a301-4f21-a081-411020864452.
	//
	// **Note:** Using the ~ match phrase/prefix comparator. For best results, we
	// recommend treating this as a "starts with" comparator for your search query.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// Name of the parameter on which sorting is performed. Use any field available on
	// the Commitment entity to sort by, such as `accountId`, `endDate`, and so on.
	SortBy param.Field[string] `query:"sortBy"`
	// Sorting order.
	SortOrder param.Field[CommitmentSearchParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [CommitmentSearchParams]'s query parameters as `url.Values`.
func (r CommitmentSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Search Operator to be used while querying search.
type CommitmentSearchParamsOperator string

const (
	CommitmentSearchParamsOperatorAnd CommitmentSearchParamsOperator = "AND"
	CommitmentSearchParamsOperatorOr  CommitmentSearchParamsOperator = "OR"
)

func (r CommitmentSearchParamsOperator) IsKnown() bool {
	switch r {
	case CommitmentSearchParamsOperatorAnd, CommitmentSearchParamsOperatorOr:
		return true
	}
	return false
}

// Sorting order.
type CommitmentSearchParamsSortOrder string

const (
	CommitmentSearchParamsSortOrderAsc  CommitmentSearchParamsSortOrder = "ASC"
	CommitmentSearchParamsSortOrderDesc CommitmentSearchParamsSortOrder = "DESC"
)

func (r CommitmentSearchParamsSortOrder) IsKnown() bool {
	switch r {
	case CommitmentSearchParamsSortOrderAsc, CommitmentSearchParamsSortOrderDesc:
		return true
	}
	return false
}
