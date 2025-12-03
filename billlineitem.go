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

// BillLineItemService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillLineItemService] method instead.
type BillLineItemService struct {
	Options []option.RequestOption
}

// NewBillLineItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillLineItemService(opts ...option.RequestOption) (r *BillLineItemService) {
	r = &BillLineItemService{}
	r.Options = opts
	return
}

// Retrieves a specific line item within a Bill.
//
// This endpoint retrieves the line item given by its unique identifier (UUID) from
// a specific Bill.
func (r *BillLineItemService) Get(ctx context.Context, billID string, id string, params BillLineItemGetParams, opts ...option.RequestOption) (res *LineItemResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/lineitems/%s", params.OrgID, billID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Lists all the line items for a specific Bill.
//
// This endpoint retrieves a list of line items for the given Bill within the
// specified Organization. The list can also be paginated for easier management.
// The line items returned in the list include individual charges, discounts, or
// adjustments within a Bill.
func (r *BillLineItemService) List(ctx context.Context, billID string, params BillLineItemListParams, opts ...option.RequestOption) (res *pagination.Cursor[LineItemResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/lineitems", params.OrgID, billID)
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

// Lists all the line items for a specific Bill.
//
// This endpoint retrieves a list of line items for the given Bill within the
// specified Organization. The list can also be paginated for easier management.
// The line items returned in the list include individual charges, discounts, or
// adjustments within a Bill.
func (r *BillLineItemService) ListAutoPaging(ctx context.Context, billID string, params BillLineItemListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[LineItemResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, billID, params, opts...))
}

type LineItemResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The code of the Accounting Product associated with this line item.
	AccountingProductCode string `json:"accountingProductCode"`
	// The unique identifier (UUID) for the associated Accounting Product.
	AccountingProductID string `json:"accountingProductId"`
	// The name of the Accounting Product associated with this line item.
	AccountingProductName string                 `json:"accountingProductName"`
	Additional            map[string]interface{} `json:"additional"`
	// A unique identifier (UUID) for the Aggregation that contributes to this Bill
	// line item.
	AggregationID string `json:"aggregationId"`
	// Represents the average unit price calculated across all pricing bands or tiers
	// for this line item.
	AverageUnitPrice float64 `json:"averageUnitPrice"`
	// The unique identifier (UUID) for the Balance associated with this line item.
	BalanceID string `json:"balanceId"`
	// Array containing the pricing band information, which shows the details for each
	// pricing band or tier.
	BandUsage []LineItemResponseBandUsage `json:"bandUsage"`
	// The unique identifier (UUID) for the Bill that includes this line item.
	BillID string `json:"billId"`
	// The unique identifier (UUID) for the Charge associated with this line item.
	ChargeID string `json:"chargeId"`
	// The unique identifier (UUID) of the Commitment associated with this line item.
	CommitmentID string `json:"commitmentId"`
	// A unique identifier (UUID) for the Compound Aggregation, if applicable.
	CompoundAggregationID string `json:"compoundAggregationId"`
	// The unique identifier (UUID) for the Contract associated with this line item.
	ContractID string `json:"contractId"`
	// The currency conversion rate _(if used)_ for the line item.
	ConversionRate float64 `json:"conversionRate"`
	// The subtotal amount for this line item after currency conversion, if applicable.
	ConvertedSubtotal float64 `json:"convertedSubtotal"`
	// The unique identifier (UUID) for the Counter associated with this line item.
	CounterID string `json:"counterId"`
	// The ID of the user who created this line item.
	CreatedBy string `json:"createdBy"`
	// The unique identifier (UUID) for the type of credit applied to this line item.
	CreditTypeID string `json:"creditTypeId"`
	// The currency in which the line item is billed, represented as a currency code.
	// For example, USD, GBP, or EUR.
	Currency    string `json:"currency"`
	Description string `json:"description"`
	// The DateTime when the line item was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the line item was last modified.
	DtLastModified time.Time         `json:"dtLastModified" format:"date-time"`
	Group          map[string]string `json:"group"`
	// Boolean flag indicating whether the Bill line item has associated statement
	// usage in JSON format. When a Bill statement is generated, usage line items have
	// their usage stored in JSON format.
	//
	// Deprecated: deprecated
	JsonUsageGenerated bool `json:"jsonUsageGenerated"`
	// The ID of the user who last modified this line item.
	LastModifiedBy string                       `json:"lastModifiedBy"`
	LineItemType   LineItemResponseLineItemType `json:"lineItemType"`
	// The unique identifier (UUID) of the Meter responsible for tracking usage.
	MeterID string `json:"meterId"`
	// The unique identifier (UUID) of the Plan Group associated with this line item.
	PlanGroupID string `json:"planGroupId"`
	// A unique identifier (UUID) for the billing Plan associated with this line item.
	PlanID string `json:"planId"`
	// The unique identifier (UUID) of the Pricing used for this line item,
	PricingID string `json:"pricingId"`
	// The code of the Product associated with this line item.
	ProductCode string `json:"productCode"`
	ProductID   string `json:"productId"`
	// The name of the Product associated with this line item.
	ProductName string `json:"productName"`
	// The amount of the product or service used in this line item.
	Quantity float64 `json:"quantity"`
	// A unique identifier (UUID) for the reason or justification for this line item,
	// if applicable.
	ReasonID             string `json:"reasonId"`
	ReferencedBillID     string `json:"referencedBillId"`
	ReferencedLineItemID string `json:"referencedLineItemId"`
	// Specifies the segment name or identifier when segmented Aggregation is used.
	// This is relevant for more complex billing structures.
	Segment map[string]string `json:"segment"`
	// The line item sequence number.
	SequenceNumber         int64     `json:"sequenceNumber"`
	ServicePeriodEndDate   time.Time `json:"servicePeriodEndDate" format:"date-time"`
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// The subtotal amount when not currency converted _(in the cases where currency
	// conversion is required)_.
	Subtotal float64 `json:"subtotal"`
	// Specifies the unit type. For example: **MB**, **GB**, **api_calls**, and so on.
	Unit string `json:"unit"`
	// The number of units rated in the line item, each of which is of the type
	// specified in the `unit` field. For example: 400 api_calls.
	//
	// In this example, the unit type of **api_calls** is read from the `unit` field.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                `json:"version"`
	JSON    lineItemResponseJSON `json:"-"`
}

// lineItemResponseJSON contains the JSON metadata for the struct
// [LineItemResponse]
type lineItemResponseJSON struct {
	ID                     apijson.Field
	AccountingProductCode  apijson.Field
	AccountingProductID    apijson.Field
	AccountingProductName  apijson.Field
	Additional             apijson.Field
	AggregationID          apijson.Field
	AverageUnitPrice       apijson.Field
	BalanceID              apijson.Field
	BandUsage              apijson.Field
	BillID                 apijson.Field
	ChargeID               apijson.Field
	CommitmentID           apijson.Field
	CompoundAggregationID  apijson.Field
	ContractID             apijson.Field
	ConversionRate         apijson.Field
	ConvertedSubtotal      apijson.Field
	CounterID              apijson.Field
	CreatedBy              apijson.Field
	CreditTypeID           apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	Group                  apijson.Field
	JsonUsageGenerated     apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	MeterID                apijson.Field
	PlanGroupID            apijson.Field
	PlanID                 apijson.Field
	PricingID              apijson.Field
	ProductCode            apijson.Field
	ProductID              apijson.Field
	ProductName            apijson.Field
	Quantity               apijson.Field
	ReasonID               apijson.Field
	ReferencedBillID       apijson.Field
	ReferencedLineItemID   apijson.Field
	Segment                apijson.Field
	SequenceNumber         apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	Subtotal               apijson.Field
	Unit                   apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *LineItemResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lineItemResponseJSON) RawJSON() string {
	return r.raw
}

// Array containing the pricing band information, which shows the details for each
// pricing band or tier.
type LineItemResponseBandUsage struct {
	// Usage amount within the band.
	BandQuantity float64 `json:"bandQuantity"`
	// Subtotal amount for the band.
	BandSubtotal float64 `json:"bandSubtotal"`
	// The number of units used within the band.
	BandUnits             float64 `json:"bandUnits"`
	ConvertedBandSubtotal float64 `json:"convertedBandSubtotal"`
	// The UUID of the credit type.
	CreditTypeID string `json:"creditTypeId"`
	// Fixed price is a charge entered for certain pricing types such as Stairstep,
	// Custom Tiered, and Custom Volume. It is a set price and not dependent on usage.
	FixedPrice float64 `json:"fixedPrice"`
	// The lower limit _(start)_ of the pricing band.
	LowerLimit float64 `json:"lowerLimit"`
	// The UUID for the pricing band.
	PricingBandID string `json:"pricingBandId"`
	// The price per unit in the band.
	UnitPrice float64 `json:"unitPrice"`
	// The subtotal of the unit usage.
	UnitSubtotal float64                       `json:"unitSubtotal"`
	JSON         lineItemResponseBandUsageJSON `json:"-"`
}

// lineItemResponseBandUsageJSON contains the JSON metadata for the struct
// [LineItemResponseBandUsage]
type lineItemResponseBandUsageJSON struct {
	BandQuantity          apijson.Field
	BandSubtotal          apijson.Field
	BandUnits             apijson.Field
	ConvertedBandSubtotal apijson.Field
	CreditTypeID          apijson.Field
	FixedPrice            apijson.Field
	LowerLimit            apijson.Field
	PricingBandID         apijson.Field
	UnitPrice             apijson.Field
	UnitSubtotal          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *LineItemResponseBandUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lineItemResponseBandUsageJSON) RawJSON() string {
	return r.raw
}

type LineItemResponseLineItemType string

const (
	LineItemResponseLineItemTypeStandingCharge            LineItemResponseLineItemType = "STANDING_CHARGE"
	LineItemResponseLineItemTypeUsage                     LineItemResponseLineItemType = "USAGE"
	LineItemResponseLineItemTypeCounterRunningTotalCharge LineItemResponseLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	LineItemResponseLineItemTypeCounterAdjustmentDebit    LineItemResponseLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	LineItemResponseLineItemTypeCounterAdjustmentCredit   LineItemResponseLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	LineItemResponseLineItemTypeUsageCredit               LineItemResponseLineItemType = "USAGE_CREDIT"
	LineItemResponseLineItemTypeMinimumSpend              LineItemResponseLineItemType = "MINIMUM_SPEND"
	LineItemResponseLineItemTypeMinimumSpendRefund        LineItemResponseLineItemType = "MINIMUM_SPEND_REFUND"
	LineItemResponseLineItemTypeCreditDeduction           LineItemResponseLineItemType = "CREDIT_DEDUCTION"
	LineItemResponseLineItemTypeManualAdjustment          LineItemResponseLineItemType = "MANUAL_ADJUSTMENT"
	LineItemResponseLineItemTypeCreditMemo                LineItemResponseLineItemType = "CREDIT_MEMO"
	LineItemResponseLineItemTypeDebitMemo                 LineItemResponseLineItemType = "DEBIT_MEMO"
	LineItemResponseLineItemTypeCommitmentConsumed        LineItemResponseLineItemType = "COMMITMENT_CONSUMED"
	LineItemResponseLineItemTypeCommitmentFee             LineItemResponseLineItemType = "COMMITMENT_FEE"
	LineItemResponseLineItemTypeOverageSurcharge          LineItemResponseLineItemType = "OVERAGE_SURCHARGE"
	LineItemResponseLineItemTypeOverageUsage              LineItemResponseLineItemType = "OVERAGE_USAGE"
	LineItemResponseLineItemTypeBalanceConsumed           LineItemResponseLineItemType = "BALANCE_CONSUMED"
	LineItemResponseLineItemTypeBalanceFee                LineItemResponseLineItemType = "BALANCE_FEE"
	LineItemResponseLineItemTypeAdHoc                     LineItemResponseLineItemType = "AD_HOC"
)

func (r LineItemResponseLineItemType) IsKnown() bool {
	switch r {
	case LineItemResponseLineItemTypeStandingCharge, LineItemResponseLineItemTypeUsage, LineItemResponseLineItemTypeCounterRunningTotalCharge, LineItemResponseLineItemTypeCounterAdjustmentDebit, LineItemResponseLineItemTypeCounterAdjustmentCredit, LineItemResponseLineItemTypeUsageCredit, LineItemResponseLineItemTypeMinimumSpend, LineItemResponseLineItemTypeMinimumSpendRefund, LineItemResponseLineItemTypeCreditDeduction, LineItemResponseLineItemTypeManualAdjustment, LineItemResponseLineItemTypeCreditMemo, LineItemResponseLineItemTypeDebitMemo, LineItemResponseLineItemTypeCommitmentConsumed, LineItemResponseLineItemTypeCommitmentFee, LineItemResponseLineItemTypeOverageSurcharge, LineItemResponseLineItemTypeOverageUsage, LineItemResponseLineItemTypeBalanceConsumed, LineItemResponseLineItemTypeBalanceFee, LineItemResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

type BillLineItemGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Comma separated list of additional fields.
	Additional param.Field[[]string] `query:"additional"`
}

// URLQuery serializes [BillLineItemGetParams]'s query parameters as `url.Values`.
func (r BillLineItemGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillLineItemListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Comma separated list of additional fields.
	Additional param.Field[[]string] `query:"additional"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// line items in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of line items to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [BillLineItemListParams]'s query parameters as `url.Values`.
func (r BillLineItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
