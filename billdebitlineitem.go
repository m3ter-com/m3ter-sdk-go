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

// BillDebitLineItemService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillDebitLineItemService] method instead.
type BillDebitLineItemService struct {
	Options []option.RequestOption
}

// NewBillDebitLineItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBillDebitLineItemService(opts ...option.RequestOption) (r *BillDebitLineItemService) {
	r = &BillDebitLineItemService{}
	r.Options = opts
	return
}

// Create a new Debit line item for the given bill.
//
// When creating Debit line items for Bills, use the Debit Reasons created for your
// Organization. See [DebitReason](https://www.m3ter.com/docs/api#tag/DebitReason).
func (r *BillDebitLineItemService) New(ctx context.Context, billID string, params BillDebitLineItemNewParams, opts ...option.RequestOption) (res *DebitLineItemResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/debitlineitems", params.OrgID, billID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Debit line item with the given UUID.
func (r *BillDebitLineItemService) Get(ctx context.Context, billID string, id string, query BillDebitLineItemGetParams, opts ...option.RequestOption) (res *DebitLineItemResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/debitlineitems/%s", query.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Debit line item with the given UUID.
func (r *BillDebitLineItemService) Update(ctx context.Context, billID string, id string, params BillDebitLineItemUpdateParams, opts ...option.RequestOption) (res *DebitLineItemResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/debitlineitems/%s", params.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List the Debit line items for the given bill.
func (r *BillDebitLineItemService) List(ctx context.Context, billID string, params BillDebitLineItemListParams, opts ...option.RequestOption) (res *pagination.Cursor[DebitLineItemResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/debitlineitems", params.OrgID, billID)
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

// List the Debit line items for the given bill.
func (r *BillDebitLineItemService) ListAutoPaging(ctx context.Context, billID string, params BillDebitLineItemListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DebitLineItemResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, billID, params, opts...))
}

// Delete the Debit line item with the given UUID.
func (r *BillDebitLineItemService) Delete(ctx context.Context, billID string, id string, body BillDebitLineItemDeleteParams, opts ...option.RequestOption) (res *DebitLineItemResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/debitlineitems/%s", body.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DebitLineItemResponse struct {
	// The UUID of the entity.
	ID     string  `json:"id" api:"required"`
	Amount float64 `json:"amount"`
	// The ID of the user who created this line item.
	CreatedBy string `json:"createdBy"`
	// The UUID of the debit reason for this debit line item.
	DebitReasonID string `json:"debitReasonId"`
	Description   string `json:"description"`
	// The DateTime when the line item was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the line item was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this line item.
	LastModifiedBy         string                            `json:"lastModifiedBy"`
	LineItemType           DebitLineItemResponseLineItemType `json:"lineItemType"`
	ProductID              string                            `json:"productId"`
	ReferencedBillID       string                            `json:"referencedBillId"`
	ReferencedLineItemID   string                            `json:"referencedLineItemId"`
	ServicePeriodEndDate   time.Time                         `json:"servicePeriodEndDate" format:"date-time"`
	ServicePeriodStartDate time.Time                         `json:"servicePeriodStartDate" format:"date-time"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                     `json:"version"`
	JSON    debitLineItemResponseJSON `json:"-"`
}

// debitLineItemResponseJSON contains the JSON metadata for the struct
// [DebitLineItemResponse]
type debitLineItemResponseJSON struct {
	ID                     apijson.Field
	Amount                 apijson.Field
	CreatedBy              apijson.Field
	DebitReasonID          apijson.Field
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

func (r *DebitLineItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r debitLineItemResponseJSON) RawJSON() string {
	return r.raw
}

type DebitLineItemResponseLineItemType string

const (
	DebitLineItemResponseLineItemTypeStandingCharge            DebitLineItemResponseLineItemType = "STANDING_CHARGE"
	DebitLineItemResponseLineItemTypeUsage                     DebitLineItemResponseLineItemType = "USAGE"
	DebitLineItemResponseLineItemTypeCounterRunningTotalCharge DebitLineItemResponseLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	DebitLineItemResponseLineItemTypeCounterAdjustmentDebit    DebitLineItemResponseLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	DebitLineItemResponseLineItemTypeCounterAdjustmentCredit   DebitLineItemResponseLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	DebitLineItemResponseLineItemTypeUsageCredit               DebitLineItemResponseLineItemType = "USAGE_CREDIT"
	DebitLineItemResponseLineItemTypeMinimumSpend              DebitLineItemResponseLineItemType = "MINIMUM_SPEND"
	DebitLineItemResponseLineItemTypeMinimumSpendRefund        DebitLineItemResponseLineItemType = "MINIMUM_SPEND_REFUND"
	DebitLineItemResponseLineItemTypeCreditDeduction           DebitLineItemResponseLineItemType = "CREDIT_DEDUCTION"
	DebitLineItemResponseLineItemTypeManualAdjustment          DebitLineItemResponseLineItemType = "MANUAL_ADJUSTMENT"
	DebitLineItemResponseLineItemTypeCreditMemo                DebitLineItemResponseLineItemType = "CREDIT_MEMO"
	DebitLineItemResponseLineItemTypeDebitMemo                 DebitLineItemResponseLineItemType = "DEBIT_MEMO"
	DebitLineItemResponseLineItemTypeCommitmentConsumed        DebitLineItemResponseLineItemType = "COMMITMENT_CONSUMED"
	DebitLineItemResponseLineItemTypeCommitmentFee             DebitLineItemResponseLineItemType = "COMMITMENT_FEE"
	DebitLineItemResponseLineItemTypeOverageSurcharge          DebitLineItemResponseLineItemType = "OVERAGE_SURCHARGE"
	DebitLineItemResponseLineItemTypeOverageUsage              DebitLineItemResponseLineItemType = "OVERAGE_USAGE"
	DebitLineItemResponseLineItemTypeBalanceConsumed           DebitLineItemResponseLineItemType = "BALANCE_CONSUMED"
	DebitLineItemResponseLineItemTypeBalanceFee                DebitLineItemResponseLineItemType = "BALANCE_FEE"
	DebitLineItemResponseLineItemTypeAdHoc                     DebitLineItemResponseLineItemType = "AD_HOC"
)

func (r DebitLineItemResponseLineItemType) IsKnown() bool {
	switch r {
	case DebitLineItemResponseLineItemTypeStandingCharge, DebitLineItemResponseLineItemTypeUsage, DebitLineItemResponseLineItemTypeCounterRunningTotalCharge, DebitLineItemResponseLineItemTypeCounterAdjustmentDebit, DebitLineItemResponseLineItemTypeCounterAdjustmentCredit, DebitLineItemResponseLineItemTypeUsageCredit, DebitLineItemResponseLineItemTypeMinimumSpend, DebitLineItemResponseLineItemTypeMinimumSpendRefund, DebitLineItemResponseLineItemTypeCreditDeduction, DebitLineItemResponseLineItemTypeManualAdjustment, DebitLineItemResponseLineItemTypeCreditMemo, DebitLineItemResponseLineItemTypeDebitMemo, DebitLineItemResponseLineItemTypeCommitmentConsumed, DebitLineItemResponseLineItemTypeCommitmentFee, DebitLineItemResponseLineItemTypeOverageSurcharge, DebitLineItemResponseLineItemTypeOverageUsage, DebitLineItemResponseLineItemTypeBalanceConsumed, DebitLineItemResponseLineItemTypeBalanceFee, DebitLineItemResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillDebitLineItemNewParams struct {
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
	// The ID of the Debit Reason given for this debit line item.
	DebitReasonID param.Field[string]                                 `json:"debitReasonId"`
	LineItemType  param.Field[BillDebitLineItemNewParamsLineItemType] `json:"lineItemType"`
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

func (r BillDebitLineItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillDebitLineItemNewParamsLineItemType string

const (
	BillDebitLineItemNewParamsLineItemTypeStandingCharge            BillDebitLineItemNewParamsLineItemType = "STANDING_CHARGE"
	BillDebitLineItemNewParamsLineItemTypeUsage                     BillDebitLineItemNewParamsLineItemType = "USAGE"
	BillDebitLineItemNewParamsLineItemTypeCounterRunningTotalCharge BillDebitLineItemNewParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BillDebitLineItemNewParamsLineItemTypeCounterAdjustmentDebit    BillDebitLineItemNewParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BillDebitLineItemNewParamsLineItemTypeCounterAdjustmentCredit   BillDebitLineItemNewParamsLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	BillDebitLineItemNewParamsLineItemTypeUsageCredit               BillDebitLineItemNewParamsLineItemType = "USAGE_CREDIT"
	BillDebitLineItemNewParamsLineItemTypeMinimumSpend              BillDebitLineItemNewParamsLineItemType = "MINIMUM_SPEND"
	BillDebitLineItemNewParamsLineItemTypeMinimumSpendRefund        BillDebitLineItemNewParamsLineItemType = "MINIMUM_SPEND_REFUND"
	BillDebitLineItemNewParamsLineItemTypeCreditDeduction           BillDebitLineItemNewParamsLineItemType = "CREDIT_DEDUCTION"
	BillDebitLineItemNewParamsLineItemTypeManualAdjustment          BillDebitLineItemNewParamsLineItemType = "MANUAL_ADJUSTMENT"
	BillDebitLineItemNewParamsLineItemTypeCreditMemo                BillDebitLineItemNewParamsLineItemType = "CREDIT_MEMO"
	BillDebitLineItemNewParamsLineItemTypeDebitMemo                 BillDebitLineItemNewParamsLineItemType = "DEBIT_MEMO"
	BillDebitLineItemNewParamsLineItemTypeCommitmentConsumed        BillDebitLineItemNewParamsLineItemType = "COMMITMENT_CONSUMED"
	BillDebitLineItemNewParamsLineItemTypeCommitmentFee             BillDebitLineItemNewParamsLineItemType = "COMMITMENT_FEE"
	BillDebitLineItemNewParamsLineItemTypeOverageSurcharge          BillDebitLineItemNewParamsLineItemType = "OVERAGE_SURCHARGE"
	BillDebitLineItemNewParamsLineItemTypeOverageUsage              BillDebitLineItemNewParamsLineItemType = "OVERAGE_USAGE"
	BillDebitLineItemNewParamsLineItemTypeBalanceConsumed           BillDebitLineItemNewParamsLineItemType = "BALANCE_CONSUMED"
	BillDebitLineItemNewParamsLineItemTypeBalanceFee                BillDebitLineItemNewParamsLineItemType = "BALANCE_FEE"
	BillDebitLineItemNewParamsLineItemTypeAdHoc                     BillDebitLineItemNewParamsLineItemType = "AD_HOC"
)

func (r BillDebitLineItemNewParamsLineItemType) IsKnown() bool {
	switch r {
	case BillDebitLineItemNewParamsLineItemTypeStandingCharge, BillDebitLineItemNewParamsLineItemTypeUsage, BillDebitLineItemNewParamsLineItemTypeCounterRunningTotalCharge, BillDebitLineItemNewParamsLineItemTypeCounterAdjustmentDebit, BillDebitLineItemNewParamsLineItemTypeCounterAdjustmentCredit, BillDebitLineItemNewParamsLineItemTypeUsageCredit, BillDebitLineItemNewParamsLineItemTypeMinimumSpend, BillDebitLineItemNewParamsLineItemTypeMinimumSpendRefund, BillDebitLineItemNewParamsLineItemTypeCreditDeduction, BillDebitLineItemNewParamsLineItemTypeManualAdjustment, BillDebitLineItemNewParamsLineItemTypeCreditMemo, BillDebitLineItemNewParamsLineItemTypeDebitMemo, BillDebitLineItemNewParamsLineItemTypeCommitmentConsumed, BillDebitLineItemNewParamsLineItemTypeCommitmentFee, BillDebitLineItemNewParamsLineItemTypeOverageSurcharge, BillDebitLineItemNewParamsLineItemTypeOverageUsage, BillDebitLineItemNewParamsLineItemTypeBalanceConsumed, BillDebitLineItemNewParamsLineItemTypeBalanceFee, BillDebitLineItemNewParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillDebitLineItemGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type BillDebitLineItemUpdateParams struct {
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
	// The ID of the Debit Reason given for this debit line item.
	DebitReasonID param.Field[string]                                    `json:"debitReasonId"`
	LineItemType  param.Field[BillDebitLineItemUpdateParamsLineItemType] `json:"lineItemType"`
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

func (r BillDebitLineItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillDebitLineItemUpdateParamsLineItemType string

const (
	BillDebitLineItemUpdateParamsLineItemTypeStandingCharge            BillDebitLineItemUpdateParamsLineItemType = "STANDING_CHARGE"
	BillDebitLineItemUpdateParamsLineItemTypeUsage                     BillDebitLineItemUpdateParamsLineItemType = "USAGE"
	BillDebitLineItemUpdateParamsLineItemTypeCounterRunningTotalCharge BillDebitLineItemUpdateParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BillDebitLineItemUpdateParamsLineItemTypeCounterAdjustmentDebit    BillDebitLineItemUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BillDebitLineItemUpdateParamsLineItemTypeCounterAdjustmentCredit   BillDebitLineItemUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	BillDebitLineItemUpdateParamsLineItemTypeUsageCredit               BillDebitLineItemUpdateParamsLineItemType = "USAGE_CREDIT"
	BillDebitLineItemUpdateParamsLineItemTypeMinimumSpend              BillDebitLineItemUpdateParamsLineItemType = "MINIMUM_SPEND"
	BillDebitLineItemUpdateParamsLineItemTypeMinimumSpendRefund        BillDebitLineItemUpdateParamsLineItemType = "MINIMUM_SPEND_REFUND"
	BillDebitLineItemUpdateParamsLineItemTypeCreditDeduction           BillDebitLineItemUpdateParamsLineItemType = "CREDIT_DEDUCTION"
	BillDebitLineItemUpdateParamsLineItemTypeManualAdjustment          BillDebitLineItemUpdateParamsLineItemType = "MANUAL_ADJUSTMENT"
	BillDebitLineItemUpdateParamsLineItemTypeCreditMemo                BillDebitLineItemUpdateParamsLineItemType = "CREDIT_MEMO"
	BillDebitLineItemUpdateParamsLineItemTypeDebitMemo                 BillDebitLineItemUpdateParamsLineItemType = "DEBIT_MEMO"
	BillDebitLineItemUpdateParamsLineItemTypeCommitmentConsumed        BillDebitLineItemUpdateParamsLineItemType = "COMMITMENT_CONSUMED"
	BillDebitLineItemUpdateParamsLineItemTypeCommitmentFee             BillDebitLineItemUpdateParamsLineItemType = "COMMITMENT_FEE"
	BillDebitLineItemUpdateParamsLineItemTypeOverageSurcharge          BillDebitLineItemUpdateParamsLineItemType = "OVERAGE_SURCHARGE"
	BillDebitLineItemUpdateParamsLineItemTypeOverageUsage              BillDebitLineItemUpdateParamsLineItemType = "OVERAGE_USAGE"
	BillDebitLineItemUpdateParamsLineItemTypeBalanceConsumed           BillDebitLineItemUpdateParamsLineItemType = "BALANCE_CONSUMED"
	BillDebitLineItemUpdateParamsLineItemTypeBalanceFee                BillDebitLineItemUpdateParamsLineItemType = "BALANCE_FEE"
	BillDebitLineItemUpdateParamsLineItemTypeAdHoc                     BillDebitLineItemUpdateParamsLineItemType = "AD_HOC"
)

func (r BillDebitLineItemUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case BillDebitLineItemUpdateParamsLineItemTypeStandingCharge, BillDebitLineItemUpdateParamsLineItemTypeUsage, BillDebitLineItemUpdateParamsLineItemTypeCounterRunningTotalCharge, BillDebitLineItemUpdateParamsLineItemTypeCounterAdjustmentDebit, BillDebitLineItemUpdateParamsLineItemTypeCounterAdjustmentCredit, BillDebitLineItemUpdateParamsLineItemTypeUsageCredit, BillDebitLineItemUpdateParamsLineItemTypeMinimumSpend, BillDebitLineItemUpdateParamsLineItemTypeMinimumSpendRefund, BillDebitLineItemUpdateParamsLineItemTypeCreditDeduction, BillDebitLineItemUpdateParamsLineItemTypeManualAdjustment, BillDebitLineItemUpdateParamsLineItemTypeCreditMemo, BillDebitLineItemUpdateParamsLineItemTypeDebitMemo, BillDebitLineItemUpdateParamsLineItemTypeCommitmentConsumed, BillDebitLineItemUpdateParamsLineItemTypeCommitmentFee, BillDebitLineItemUpdateParamsLineItemTypeOverageSurcharge, BillDebitLineItemUpdateParamsLineItemTypeOverageUsage, BillDebitLineItemUpdateParamsLineItemTypeBalanceConsumed, BillDebitLineItemUpdateParamsLineItemTypeBalanceFee, BillDebitLineItemUpdateParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillDebitLineItemListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of line items to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [BillDebitLineItemListParams]'s query parameters as
// `url.Values`.
func (r BillDebitLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillDebitLineItemDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
