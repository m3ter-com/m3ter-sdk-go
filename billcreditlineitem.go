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
)

// BillCreditLineItemService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillCreditLineItemService] method instead.
type BillCreditLineItemService struct {
	Options []option.RequestOption
}

// NewBillCreditLineItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBillCreditLineItemService(opts ...option.RequestOption) (r *BillCreditLineItemService) {
	r = &BillCreditLineItemService{}
	r.Options = opts
	return
}

// Create a new Credit line item for the given Bill.
//
// When creating Credit line items for Bills, use the Credit Reasons created for
// your Organization. See
// [CreditReason](https://www.m3ter.com/docs/api#tag/CreditReason).
func (r *BillCreditLineItemService) New(ctx context.Context, billID string, params BillCreditLineItemNewParams, opts ...option.RequestOption) (res *CreditLineItem, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if billID == "" {
		err = errors.New("missing required billId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/creditlineitems", params.OrgID, billID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Credit line item with the given UUID.
func (r *BillCreditLineItemService) Get(ctx context.Context, billID string, id string, query BillCreditLineItemGetParams, opts ...option.RequestOption) (res *CreditLineItem, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if billID == "" {
		err = errors.New("missing required billId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/creditlineitems/%s", query.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Credit line item with the given UUID.
func (r *BillCreditLineItemService) Update(ctx context.Context, billID string, id string, params BillCreditLineItemUpdateParams, opts ...option.RequestOption) (res *CreditLineItem, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if billID == "" {
		err = errors.New("missing required billId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/creditlineitems/%s", params.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List the Credit line items for the given Bill.
func (r *BillCreditLineItemService) List(ctx context.Context, billID string, params BillCreditLineItemListParams, opts ...option.RequestOption) (res *pagination.Cursor[CreditLineItem], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if billID == "" {
		err = errors.New("missing required billId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/creditlineitems", params.OrgID, billID)
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

// List the Credit line items for the given Bill.
func (r *BillCreditLineItemService) ListAutoPaging(ctx context.Context, billID string, params BillCreditLineItemListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CreditLineItem] {
	return pagination.NewCursorAutoPager(r.List(ctx, billID, params, opts...))
}

// Delete the Credit line item with the given UUID.
func (r *BillCreditLineItemService) Delete(ctx context.Context, billID string, id string, body BillCreditLineItemDeleteParams, opts ...option.RequestOption) (res *CreditLineItem, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if billID == "" {
		err = errors.New("missing required billId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/creditlineitems/%s", body.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreditLineItem struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The amount for the line item.
	Amount float64 `json:"amount,required"`
	// The description of the line item.
	Description string `json:"description,required"`
	// The UUID of the Product.
	ProductID string `json:"productId,required"`
	// The UUID of the bill for the line item.
	ReferencedBillID string `json:"referencedBillId,required"`
	// The UUID of the line item.
	ReferencedLineItemID string `json:"referencedLineItemId,required"`
	// The service period end date in ISO-8601 format. _(exclusive of the ending
	// date)_.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date in ISO-8601 format. _(inclusive of the starting
	// date)_.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate,required" format:"date-time"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The id of the user who created this credit line item.
	CreatedBy string `json:"createdBy"`
	// The UUID of the credit reason for this credit line item.
	CreditReasonID string `json:"creditReasonId"`
	// The DateTime when the credit line item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the credit line item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this credit line item.
	LastModifiedBy string             `json:"lastModifiedBy"`
	JSON           creditLineItemJSON `json:"-"`
}

// creditLineItemJSON contains the JSON metadata for the struct [CreditLineItem]
type creditLineItemJSON struct {
	ID                     apijson.Field
	Amount                 apijson.Field
	Description            apijson.Field
	ProductID              apijson.Field
	ReferencedBillID       apijson.Field
	ReferencedLineItemID   apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	Version                apijson.Field
	CreatedBy              apijson.Field
	CreditReasonID         apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CreditLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLineItemJSON) RawJSON() string {
	return r.raw
}

type BillCreditLineItemNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The amount for the line item.
	Amount param.Field[float64] `json:"amount,required"`
	// The description for the line item.
	Description param.Field[string] `json:"description,required"`
	// The UUID of the Product.
	ProductID param.Field[string] `json:"productId,required"`
	// The UUID of the bill for the line item.
	ReferencedBillID param.Field[string] `json:"referencedBillId,required"`
	// The UUID of the line item.
	ReferencedLineItemID param.Field[string] `json:"referencedLineItemId,required"`
	// The service period end date in ISO-8601 format._(exclusive of the ending date)_.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date in ISO-8601 format. _(inclusive of the starting
	// date)_.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// The UUID of the credit reason.
	CreditReasonID param.Field[string]                                  `json:"creditReasonId"`
	LineItemType   param.Field[BillCreditLineItemNewParamsLineItemType] `json:"lineItemType"`
	// The UUID of the line item reason.
	ReasonID param.Field[string] `json:"reasonId"`
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

func (r BillCreditLineItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillCreditLineItemNewParamsLineItemType string

const (
	BillCreditLineItemNewParamsLineItemTypeStandingCharge            BillCreditLineItemNewParamsLineItemType = "STANDING_CHARGE"
	BillCreditLineItemNewParamsLineItemTypeUsage                     BillCreditLineItemNewParamsLineItemType = "USAGE"
	BillCreditLineItemNewParamsLineItemTypeCounterRunningTotalCharge BillCreditLineItemNewParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BillCreditLineItemNewParamsLineItemTypeCounterAdjustmentDebit    BillCreditLineItemNewParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BillCreditLineItemNewParamsLineItemTypeCounterAdjustmentCredit   BillCreditLineItemNewParamsLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	BillCreditLineItemNewParamsLineItemTypeUsageCredit               BillCreditLineItemNewParamsLineItemType = "USAGE_CREDIT"
	BillCreditLineItemNewParamsLineItemTypeMinimumSpend              BillCreditLineItemNewParamsLineItemType = "MINIMUM_SPEND"
	BillCreditLineItemNewParamsLineItemTypeMinimumSpendRefund        BillCreditLineItemNewParamsLineItemType = "MINIMUM_SPEND_REFUND"
	BillCreditLineItemNewParamsLineItemTypeCreditDeduction           BillCreditLineItemNewParamsLineItemType = "CREDIT_DEDUCTION"
	BillCreditLineItemNewParamsLineItemTypeManualAdjustment          BillCreditLineItemNewParamsLineItemType = "MANUAL_ADJUSTMENT"
	BillCreditLineItemNewParamsLineItemTypeCreditMemo                BillCreditLineItemNewParamsLineItemType = "CREDIT_MEMO"
	BillCreditLineItemNewParamsLineItemTypeDebitMemo                 BillCreditLineItemNewParamsLineItemType = "DEBIT_MEMO"
	BillCreditLineItemNewParamsLineItemTypeCommitmentConsumed        BillCreditLineItemNewParamsLineItemType = "COMMITMENT_CONSUMED"
	BillCreditLineItemNewParamsLineItemTypeCommitmentFee             BillCreditLineItemNewParamsLineItemType = "COMMITMENT_FEE"
	BillCreditLineItemNewParamsLineItemTypeOverageSurcharge          BillCreditLineItemNewParamsLineItemType = "OVERAGE_SURCHARGE"
	BillCreditLineItemNewParamsLineItemTypeOverageUsage              BillCreditLineItemNewParamsLineItemType = "OVERAGE_USAGE"
	BillCreditLineItemNewParamsLineItemTypeBalanceConsumed           BillCreditLineItemNewParamsLineItemType = "BALANCE_CONSUMED"
	BillCreditLineItemNewParamsLineItemTypeBalanceFee                BillCreditLineItemNewParamsLineItemType = "BALANCE_FEE"
)

func (r BillCreditLineItemNewParamsLineItemType) IsKnown() bool {
	switch r {
	case BillCreditLineItemNewParamsLineItemTypeStandingCharge, BillCreditLineItemNewParamsLineItemTypeUsage, BillCreditLineItemNewParamsLineItemTypeCounterRunningTotalCharge, BillCreditLineItemNewParamsLineItemTypeCounterAdjustmentDebit, BillCreditLineItemNewParamsLineItemTypeCounterAdjustmentCredit, BillCreditLineItemNewParamsLineItemTypeUsageCredit, BillCreditLineItemNewParamsLineItemTypeMinimumSpend, BillCreditLineItemNewParamsLineItemTypeMinimumSpendRefund, BillCreditLineItemNewParamsLineItemTypeCreditDeduction, BillCreditLineItemNewParamsLineItemTypeManualAdjustment, BillCreditLineItemNewParamsLineItemTypeCreditMemo, BillCreditLineItemNewParamsLineItemTypeDebitMemo, BillCreditLineItemNewParamsLineItemTypeCommitmentConsumed, BillCreditLineItemNewParamsLineItemTypeCommitmentFee, BillCreditLineItemNewParamsLineItemTypeOverageSurcharge, BillCreditLineItemNewParamsLineItemTypeOverageUsage, BillCreditLineItemNewParamsLineItemTypeBalanceConsumed, BillCreditLineItemNewParamsLineItemTypeBalanceFee:
		return true
	}
	return false
}

type BillCreditLineItemGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillCreditLineItemUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The amount for the line item.
	Amount param.Field[float64] `json:"amount,required"`
	// The description for the line item.
	Description param.Field[string] `json:"description,required"`
	// The UUID of the Product.
	ProductID param.Field[string] `json:"productId,required"`
	// The UUID of the bill for the line item.
	ReferencedBillID param.Field[string] `json:"referencedBillId,required"`
	// The UUID of the line item.
	ReferencedLineItemID param.Field[string] `json:"referencedLineItemId,required"`
	// The service period end date in ISO-8601 format._(exclusive of the ending date)_.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date in ISO-8601 format. _(inclusive of the starting
	// date)_.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// The UUID of the credit reason.
	CreditReasonID param.Field[string]                                     `json:"creditReasonId"`
	LineItemType   param.Field[BillCreditLineItemUpdateParamsLineItemType] `json:"lineItemType"`
	// The UUID of the line item reason.
	ReasonID param.Field[string] `json:"reasonId"`
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

func (r BillCreditLineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillCreditLineItemUpdateParamsLineItemType string

const (
	BillCreditLineItemUpdateParamsLineItemTypeStandingCharge            BillCreditLineItemUpdateParamsLineItemType = "STANDING_CHARGE"
	BillCreditLineItemUpdateParamsLineItemTypeUsage                     BillCreditLineItemUpdateParamsLineItemType = "USAGE"
	BillCreditLineItemUpdateParamsLineItemTypeCounterRunningTotalCharge BillCreditLineItemUpdateParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BillCreditLineItemUpdateParamsLineItemTypeCounterAdjustmentDebit    BillCreditLineItemUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BillCreditLineItemUpdateParamsLineItemTypeCounterAdjustmentCredit   BillCreditLineItemUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	BillCreditLineItemUpdateParamsLineItemTypeUsageCredit               BillCreditLineItemUpdateParamsLineItemType = "USAGE_CREDIT"
	BillCreditLineItemUpdateParamsLineItemTypeMinimumSpend              BillCreditLineItemUpdateParamsLineItemType = "MINIMUM_SPEND"
	BillCreditLineItemUpdateParamsLineItemTypeMinimumSpendRefund        BillCreditLineItemUpdateParamsLineItemType = "MINIMUM_SPEND_REFUND"
	BillCreditLineItemUpdateParamsLineItemTypeCreditDeduction           BillCreditLineItemUpdateParamsLineItemType = "CREDIT_DEDUCTION"
	BillCreditLineItemUpdateParamsLineItemTypeManualAdjustment          BillCreditLineItemUpdateParamsLineItemType = "MANUAL_ADJUSTMENT"
	BillCreditLineItemUpdateParamsLineItemTypeCreditMemo                BillCreditLineItemUpdateParamsLineItemType = "CREDIT_MEMO"
	BillCreditLineItemUpdateParamsLineItemTypeDebitMemo                 BillCreditLineItemUpdateParamsLineItemType = "DEBIT_MEMO"
	BillCreditLineItemUpdateParamsLineItemTypeCommitmentConsumed        BillCreditLineItemUpdateParamsLineItemType = "COMMITMENT_CONSUMED"
	BillCreditLineItemUpdateParamsLineItemTypeCommitmentFee             BillCreditLineItemUpdateParamsLineItemType = "COMMITMENT_FEE"
	BillCreditLineItemUpdateParamsLineItemTypeOverageSurcharge          BillCreditLineItemUpdateParamsLineItemType = "OVERAGE_SURCHARGE"
	BillCreditLineItemUpdateParamsLineItemTypeOverageUsage              BillCreditLineItemUpdateParamsLineItemType = "OVERAGE_USAGE"
	BillCreditLineItemUpdateParamsLineItemTypeBalanceConsumed           BillCreditLineItemUpdateParamsLineItemType = "BALANCE_CONSUMED"
	BillCreditLineItemUpdateParamsLineItemTypeBalanceFee                BillCreditLineItemUpdateParamsLineItemType = "BALANCE_FEE"
)

func (r BillCreditLineItemUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case BillCreditLineItemUpdateParamsLineItemTypeStandingCharge, BillCreditLineItemUpdateParamsLineItemTypeUsage, BillCreditLineItemUpdateParamsLineItemTypeCounterRunningTotalCharge, BillCreditLineItemUpdateParamsLineItemTypeCounterAdjustmentDebit, BillCreditLineItemUpdateParamsLineItemTypeCounterAdjustmentCredit, BillCreditLineItemUpdateParamsLineItemTypeUsageCredit, BillCreditLineItemUpdateParamsLineItemTypeMinimumSpend, BillCreditLineItemUpdateParamsLineItemTypeMinimumSpendRefund, BillCreditLineItemUpdateParamsLineItemTypeCreditDeduction, BillCreditLineItemUpdateParamsLineItemTypeManualAdjustment, BillCreditLineItemUpdateParamsLineItemTypeCreditMemo, BillCreditLineItemUpdateParamsLineItemTypeDebitMemo, BillCreditLineItemUpdateParamsLineItemTypeCommitmentConsumed, BillCreditLineItemUpdateParamsLineItemTypeCommitmentFee, BillCreditLineItemUpdateParamsLineItemTypeOverageSurcharge, BillCreditLineItemUpdateParamsLineItemTypeOverageUsage, BillCreditLineItemUpdateParamsLineItemTypeBalanceConsumed, BillCreditLineItemUpdateParamsLineItemTypeBalanceFee:
		return true
	}
	return false
}

type BillCreditLineItemListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Line Items to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [BillCreditLineItemListParams]'s query parameters as
// `url.Values`.
func (r BillCreditLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillCreditLineItemDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
