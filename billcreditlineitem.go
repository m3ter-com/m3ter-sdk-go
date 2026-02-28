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

// Endpoints for Credit line item related operations such as creation, update, list
// and delete. These are line items on Bills that are specifically related to
// Credits.
//
// You use the Credit Reasons created for your Organization when you create Credit
// line items for Bills. See
// [CreditReason](https://www.m3ter.com/docs/api#tag/CreditReason).
//
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
func (r *BillCreditLineItemService) New(ctx context.Context, billID string, params BillCreditLineItemNewParams, opts ...option.RequestOption) (res *CreditLineItemResponse, err error) {
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
	if billID == "" {
		err = errors.New("missing required billId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/creditlineitems", params.OrgID, billID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Credit line item with the given UUID.
func (r *BillCreditLineItemService) Get(ctx context.Context, billID string, id string, query BillCreditLineItemGetParams, opts ...option.RequestOption) (res *CreditLineItemResponse, err error) {
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
func (r *BillCreditLineItemService) Update(ctx context.Context, billID string, id string, params BillCreditLineItemUpdateParams, opts ...option.RequestOption) (res *CreditLineItemResponse, err error) {
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
func (r *BillCreditLineItemService) List(ctx context.Context, billID string, params BillCreditLineItemListParams, opts ...option.RequestOption) (res *pagination.Cursor[CreditLineItemResponse], err error) {
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
func (r *BillCreditLineItemService) ListAutoPaging(ctx context.Context, billID string, params BillCreditLineItemListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CreditLineItemResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, billID, params, opts...))
}

// Delete the Credit line item with the given UUID.
func (r *BillCreditLineItemService) Delete(ctx context.Context, billID string, id string, body BillCreditLineItemDeleteParams, opts ...option.RequestOption) (res *CreditLineItemResponse, err error) {
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

type CreditLineItemResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// The credit amount.
	Amount float64 `json:"amount"`
	// The ID of the user who created this line item.
	CreatedBy string `json:"createdBy"`
	// The UUID of the credit reason for this credit line item.
	CreditReasonID string `json:"creditReasonId"`
	Description    string `json:"description"`
	// The DateTime when the line item was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the line item was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this line item.
	LastModifiedBy         string                             `json:"lastModifiedBy"`
	LineItemType           CreditLineItemResponseLineItemType `json:"lineItemType"`
	ProductID              string                             `json:"productId"`
	ReferencedBillID       string                             `json:"referencedBillId"`
	ReferencedLineItemID   string                             `json:"referencedLineItemId"`
	ServicePeriodEndDate   time.Time                          `json:"servicePeriodEndDate" format:"date-time"`
	ServicePeriodStartDate time.Time                          `json:"servicePeriodStartDate" format:"date-time"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                      `json:"version"`
	JSON    creditLineItemResponseJSON `json:"-"`
}

// creditLineItemResponseJSON contains the JSON metadata for the struct
// [CreditLineItemResponse]
type creditLineItemResponseJSON struct {
	ID                     apijson.Field
	Amount                 apijson.Field
	CreatedBy              apijson.Field
	CreditReasonID         apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	ProductID              apijson.Field
	ReferencedBillID       apijson.Field
	ReferencedLineItemID   apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CreditLineItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLineItemResponseJSON) RawJSON() string {
	return r.raw
}

type CreditLineItemResponseLineItemType string

const (
	CreditLineItemResponseLineItemTypeStandingCharge            CreditLineItemResponseLineItemType = "STANDING_CHARGE"
	CreditLineItemResponseLineItemTypeUsage                     CreditLineItemResponseLineItemType = "USAGE"
	CreditLineItemResponseLineItemTypeCounterRunningTotalCharge CreditLineItemResponseLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	CreditLineItemResponseLineItemTypeCounterAdjustmentDebit    CreditLineItemResponseLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	CreditLineItemResponseLineItemTypeCounterAdjustmentCredit   CreditLineItemResponseLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	CreditLineItemResponseLineItemTypeUsageCredit               CreditLineItemResponseLineItemType = "USAGE_CREDIT"
	CreditLineItemResponseLineItemTypeMinimumSpend              CreditLineItemResponseLineItemType = "MINIMUM_SPEND"
	CreditLineItemResponseLineItemTypeMinimumSpendRefund        CreditLineItemResponseLineItemType = "MINIMUM_SPEND_REFUND"
	CreditLineItemResponseLineItemTypeCreditDeduction           CreditLineItemResponseLineItemType = "CREDIT_DEDUCTION"
	CreditLineItemResponseLineItemTypeManualAdjustment          CreditLineItemResponseLineItemType = "MANUAL_ADJUSTMENT"
	CreditLineItemResponseLineItemTypeCreditMemo                CreditLineItemResponseLineItemType = "CREDIT_MEMO"
	CreditLineItemResponseLineItemTypeDebitMemo                 CreditLineItemResponseLineItemType = "DEBIT_MEMO"
	CreditLineItemResponseLineItemTypeCommitmentConsumed        CreditLineItemResponseLineItemType = "COMMITMENT_CONSUMED"
	CreditLineItemResponseLineItemTypeCommitmentFee             CreditLineItemResponseLineItemType = "COMMITMENT_FEE"
	CreditLineItemResponseLineItemTypeOverageSurcharge          CreditLineItemResponseLineItemType = "OVERAGE_SURCHARGE"
	CreditLineItemResponseLineItemTypeOverageUsage              CreditLineItemResponseLineItemType = "OVERAGE_USAGE"
	CreditLineItemResponseLineItemTypeBalanceConsumed           CreditLineItemResponseLineItemType = "BALANCE_CONSUMED"
	CreditLineItemResponseLineItemTypeBalanceFee                CreditLineItemResponseLineItemType = "BALANCE_FEE"
	CreditLineItemResponseLineItemTypeAdHoc                     CreditLineItemResponseLineItemType = "AD_HOC"
)

func (r CreditLineItemResponseLineItemType) IsKnown() bool {
	switch r {
	case CreditLineItemResponseLineItemTypeStandingCharge, CreditLineItemResponseLineItemTypeUsage, CreditLineItemResponseLineItemTypeCounterRunningTotalCharge, CreditLineItemResponseLineItemTypeCounterAdjustmentDebit, CreditLineItemResponseLineItemTypeCounterAdjustmentCredit, CreditLineItemResponseLineItemTypeUsageCredit, CreditLineItemResponseLineItemTypeMinimumSpend, CreditLineItemResponseLineItemTypeMinimumSpendRefund, CreditLineItemResponseLineItemTypeCreditDeduction, CreditLineItemResponseLineItemTypeManualAdjustment, CreditLineItemResponseLineItemTypeCreditMemo, CreditLineItemResponseLineItemTypeDebitMemo, CreditLineItemResponseLineItemTypeCommitmentConsumed, CreditLineItemResponseLineItemTypeCommitmentFee, CreditLineItemResponseLineItemTypeOverageSurcharge, CreditLineItemResponseLineItemTypeOverageUsage, CreditLineItemResponseLineItemTypeBalanceConsumed, CreditLineItemResponseLineItemTypeBalanceFee, CreditLineItemResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillCreditLineItemNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID               param.Field[string] `path:"orgId" api:"required"`
	AccountingProductID param.Field[string] `json:"accountingProductId" api:"required"`
	// The amount for the line item.
	Amount param.Field[float64] `json:"amount" api:"required"`
	// The description for the line item.
	Description param.Field[string] `json:"description" api:"required"`
	// The UUID of the Product.
	ProductID param.Field[string] `json:"productId" api:"required"`
	// The UUID of the bill for the line item.
	ReferencedBillID param.Field[string] `json:"referencedBillId" api:"required"`
	// The UUID of the line item.
	ReferencedLineItemID param.Field[string] `json:"referencedLineItemId" api:"required"`
	// The service period end date in ISO-8601 format._(exclusive of the ending date)_.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate" api:"required" format:"date-time"`
	// The service period start date in ISO-8601 format. _(inclusive of the starting
	// date)_.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate" api:"required" format:"date-time"`
	AmountToApplyOnBill    param.Field[float64]   `json:"amountToApplyOnBill"`
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
	BillCreditLineItemNewParamsLineItemTypeAdHoc                     BillCreditLineItemNewParamsLineItemType = "AD_HOC"
)

func (r BillCreditLineItemNewParamsLineItemType) IsKnown() bool {
	switch r {
	case BillCreditLineItemNewParamsLineItemTypeStandingCharge, BillCreditLineItemNewParamsLineItemTypeUsage, BillCreditLineItemNewParamsLineItemTypeCounterRunningTotalCharge, BillCreditLineItemNewParamsLineItemTypeCounterAdjustmentDebit, BillCreditLineItemNewParamsLineItemTypeCounterAdjustmentCredit, BillCreditLineItemNewParamsLineItemTypeUsageCredit, BillCreditLineItemNewParamsLineItemTypeMinimumSpend, BillCreditLineItemNewParamsLineItemTypeMinimumSpendRefund, BillCreditLineItemNewParamsLineItemTypeCreditDeduction, BillCreditLineItemNewParamsLineItemTypeManualAdjustment, BillCreditLineItemNewParamsLineItemTypeCreditMemo, BillCreditLineItemNewParamsLineItemTypeDebitMemo, BillCreditLineItemNewParamsLineItemTypeCommitmentConsumed, BillCreditLineItemNewParamsLineItemTypeCommitmentFee, BillCreditLineItemNewParamsLineItemTypeOverageSurcharge, BillCreditLineItemNewParamsLineItemTypeOverageUsage, BillCreditLineItemNewParamsLineItemTypeBalanceConsumed, BillCreditLineItemNewParamsLineItemTypeBalanceFee, BillCreditLineItemNewParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillCreditLineItemGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type BillCreditLineItemUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID               param.Field[string] `path:"orgId" api:"required"`
	AccountingProductID param.Field[string] `json:"accountingProductId" api:"required"`
	// The amount for the line item.
	Amount param.Field[float64] `json:"amount" api:"required"`
	// The description for the line item.
	Description param.Field[string] `json:"description" api:"required"`
	// The UUID of the Product.
	ProductID param.Field[string] `json:"productId" api:"required"`
	// The UUID of the bill for the line item.
	ReferencedBillID param.Field[string] `json:"referencedBillId" api:"required"`
	// The UUID of the line item.
	ReferencedLineItemID param.Field[string] `json:"referencedLineItemId" api:"required"`
	// The service period end date in ISO-8601 format._(exclusive of the ending date)_.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate" api:"required" format:"date-time"`
	// The service period start date in ISO-8601 format. _(inclusive of the starting
	// date)_.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate" api:"required" format:"date-time"`
	AmountToApplyOnBill    param.Field[float64]   `json:"amountToApplyOnBill"`
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
	BillCreditLineItemUpdateParamsLineItemTypeAdHoc                     BillCreditLineItemUpdateParamsLineItemType = "AD_HOC"
)

func (r BillCreditLineItemUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case BillCreditLineItemUpdateParamsLineItemTypeStandingCharge, BillCreditLineItemUpdateParamsLineItemTypeUsage, BillCreditLineItemUpdateParamsLineItemTypeCounterRunningTotalCharge, BillCreditLineItemUpdateParamsLineItemTypeCounterAdjustmentDebit, BillCreditLineItemUpdateParamsLineItemTypeCounterAdjustmentCredit, BillCreditLineItemUpdateParamsLineItemTypeUsageCredit, BillCreditLineItemUpdateParamsLineItemTypeMinimumSpend, BillCreditLineItemUpdateParamsLineItemTypeMinimumSpendRefund, BillCreditLineItemUpdateParamsLineItemTypeCreditDeduction, BillCreditLineItemUpdateParamsLineItemTypeManualAdjustment, BillCreditLineItemUpdateParamsLineItemTypeCreditMemo, BillCreditLineItemUpdateParamsLineItemTypeDebitMemo, BillCreditLineItemUpdateParamsLineItemTypeCommitmentConsumed, BillCreditLineItemUpdateParamsLineItemTypeCommitmentFee, BillCreditLineItemUpdateParamsLineItemTypeOverageSurcharge, BillCreditLineItemUpdateParamsLineItemTypeOverageUsage, BillCreditLineItemUpdateParamsLineItemTypeBalanceConsumed, BillCreditLineItemUpdateParamsLineItemTypeBalanceFee, BillCreditLineItemUpdateParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillCreditLineItemListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
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
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
