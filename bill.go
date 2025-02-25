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

// BillService contains methods and other services that help with interacting with
// the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillService] method instead.
type BillService struct {
	Options         []option.RequestOption
	CreditLineItems *BillCreditLineItemService
	DebitLineItems  *BillDebitLineItemService
	LineItems       *BillLineItemService
}

// NewBillService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillService(opts ...option.RequestOption) (r *BillService) {
	r = &BillService{}
	r.Options = opts
	r.CreditLineItems = NewBillCreditLineItemService(opts...)
	r.DebitLineItems = NewBillDebitLineItemService(opts...)
	r.LineItems = NewBillLineItemService(opts...)
	return
}

// Retrieve the Bill with the given UUID.
//
// This endpoint retrieves the Bill with the given unique identifier (UUID) and
// specific Organization.
func (r *BillService) Get(ctx context.Context, id string, query BillGetParams, opts ...option.RequestOption) (res *Bill, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of Bills.
//
// This endpoint retrieves a list of all Bills for the given Account within the
// specified Organization. Optional filters can be applied such as by date range,
// lock status, or other attributes. The list can also be paginated for easier
// management.
func (r *BillService) List(ctx context.Context, params BillListParams, opts ...option.RequestOption) (res *pagination.Cursor[Bill], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills", params.OrgID)
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

// Retrieve a list of Bills.
//
// This endpoint retrieves a list of all Bills for the given Account within the
// specified Organization. Optional filters can be applied such as by date range,
// lock status, or other attributes. The list can also be paginated for easier
// management.
func (r *BillService) ListAutoPaging(ctx context.Context, params BillListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Bill] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Bill with the given UUID.
//
// This endpoint deletes the specified Bill with the given unique identifier. Use
// with caution since deleted Bills cannot be recovered. Suitable for removing
// incorrect or obsolete Bills, and for Bills that have not been sent to customers.
// Where end-customer invoices for Bills have been sent to customers, Bills should
// not be deleted to ensure you have an audit trail of how the invoice was created.
func (r *BillService) Delete(ctx context.Context, id string, body BillDeleteParams, opts ...option.RequestOption) (res *Bill, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Approve multiple Bills for the specified Organization based on the given
// criteria.
//
// This endpoint allows you to change currently _Pending_ Bills to _Approved_
// status for further processing.
//
// Query Parameters:
//
// - Use `accountIds` to approve Bills for specifed Accounts.
//
// Request Body Schema Parameter:
//
// - Use `billIds` to specify a collection of Bills for batch approval.
//
// **Important!** If you use the `billIds` Request Body Schema parameter, any Query
// parameters you might have also used are ignored when the call is processed.
func (r *BillService) Approve(ctx context.Context, params BillApproveParams, opts ...option.RequestOption) (res *BillApproveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/approve", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the latest Bill for the given Account.
func (r *BillService) LatestByAccount(ctx context.Context, accountID string, query BillLatestByAccountParams, opts ...option.RequestOption) (res *Bill, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/latest/%s", query.OrgID, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lock a Bill for the given UUID
func (r *BillService) Lock(ctx context.Context, id string, body BillLockParams, opts ...option.RequestOption) (res *Bill, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/lock", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Search for bill entities
func (r *BillService) Search(ctx context.Context, params BillSearchParams, opts ...option.RequestOption) (res *BillSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/search", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update Bill Status for the given UUID
func (r *BillService) UpdateStatus(ctx context.Context, id string, params BillUpdateStatusParams, opts ...option.RequestOption) (res *Bill, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/status", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type Bill struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version               int64                `json:"version,required"`
	AccountCode           string               `json:"accountCode"`
	AccountID             string               `json:"accountId"`
	BillDate              time.Time            `json:"billDate" format:"date"`
	BillFrequencyInterval int64                `json:"billFrequencyInterval"`
	BillingFrequency      BillBillingFrequency `json:"billingFrequency"`
	BillJobID             string               `json:"billJobId"`
	// The sum total for the Bill.
	BillTotal float64 `json:"billTotal"`
	// The unique identifier (UUID) for the user who created the Bill.
	CreatedBy   string    `json:"createdBy"`
	CreatedDate time.Time `json:"createdDate" format:"date-time"`
	// Flag to indicate that the statement in CSV format has been generated for the
	// Bill.
	//
	// - **TRUE** - CSV statement has been generated.
	// - **FALSE** - no CSV statement generated.
	CsvStatementGenerated bool                     `json:"csvStatementGenerated"`
	Currency              string                   `json:"currency"`
	CurrencyConversions   []BillCurrencyConversion `json:"currencyConversions"`
	// The date and time _(in ISO 8601 format)_ when the Bill was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the Bill was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	DueDate        time.Time `json:"dueDate" format:"date"`
	EndDate        time.Time `json:"endDate" format:"date"`
	EndDateTimeUtc time.Time `json:"endDateTimeUTC" format:"date-time"`
	// For accounting purposes, the date set at Organization level to use for external
	// invoicing with respect to billing periods - two options:
	//
	// - `FIRST_DAY_OF_NEXT_PERIOD` _(Default)_.
	// - `LAST_DAY_OF_ARREARS`
	//
	// For example, if the retrieved Bill was on a monthly billing frequency and the
	// billing period for the Bill is September 2023 and the _External invoice date_ is
	// set at `FIRST_DAY_OF_NEXT_PERIOD`, then the `externalInvoiceDate` will be
	// `"2023-10-01"`.
	//
	// **NOTE:** To change the `externalInvoiceDate` setting for your Organization, you
	// can use the
	// [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/GetOrganizationConfig)
	// call.
	ExternalInvoiceDate time.Time `json:"externalInvoiceDate" format:"date"`
	// The reference ID to use for external invoice.
	ExternalInvoiceReference string `json:"externalInvoiceReference"`
	// Flag to indicate that the statement in JSON format has been generated for the
	// Bill.
	//
	// - **TRUE** - JSON statement has been generated.
	// - **FALSE** - no JSON statement generated.
	JsonStatementGenerated bool      `json:"jsonStatementGenerated"`
	LastCalculatedDate     time.Time `json:"lastCalculatedDate" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified this Bill.
	LastModifiedBy string `json:"lastModifiedBy"`
	// An array of the Bill line items.
	LineItems []BillLineItem `json:"lineItems"`
	Locked    bool           `json:"locked"`
	// Purchase Order number linked to the Account the Bill is for.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// The sequential invoice number of the Bill.
	//
	// **NOTE:** If you have not defined a `billPrefix` for your Organization, a
	// `sequentialInvoiceNumber` is not returned in the response. See
	// [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/UpdateOrganizationConfig)
	SequentialInvoiceNumber string     `json:"sequentialInvoiceNumber"`
	StartDate               time.Time  `json:"startDate" format:"date"`
	StartDateTimeUtc        time.Time  `json:"startDateTimeUTC" format:"date-time"`
	Status                  BillStatus `json:"status"`
	Timezone                string     `json:"timezone"`
	JSON                    billJSON   `json:"-"`
}

// billJSON contains the JSON metadata for the struct [Bill]
type billJSON struct {
	ID                       apijson.Field
	Version                  apijson.Field
	AccountCode              apijson.Field
	AccountID                apijson.Field
	BillDate                 apijson.Field
	BillFrequencyInterval    apijson.Field
	BillingFrequency         apijson.Field
	BillJobID                apijson.Field
	BillTotal                apijson.Field
	CreatedBy                apijson.Field
	CreatedDate              apijson.Field
	CsvStatementGenerated    apijson.Field
	Currency                 apijson.Field
	CurrencyConversions      apijson.Field
	DtCreated                apijson.Field
	DtLastModified           apijson.Field
	DueDate                  apijson.Field
	EndDate                  apijson.Field
	EndDateTimeUtc           apijson.Field
	ExternalInvoiceDate      apijson.Field
	ExternalInvoiceReference apijson.Field
	JsonStatementGenerated   apijson.Field
	LastCalculatedDate       apijson.Field
	LastModifiedBy           apijson.Field
	LineItems                apijson.Field
	Locked                   apijson.Field
	PurchaseOrderNumber      apijson.Field
	SequentialInvoiceNumber  apijson.Field
	StartDate                apijson.Field
	StartDateTimeUtc         apijson.Field
	Status                   apijson.Field
	Timezone                 apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *Bill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billJSON) RawJSON() string {
	return r.raw
}

type BillBillingFrequency string

const (
	BillBillingFrequencyDaily    BillBillingFrequency = "DAILY"
	BillBillingFrequencyWeekly   BillBillingFrequency = "WEEKLY"
	BillBillingFrequencyMonthly  BillBillingFrequency = "MONTHLY"
	BillBillingFrequencyAnnually BillBillingFrequency = "ANNUALLY"
	BillBillingFrequencyAdHoc    BillBillingFrequency = "AD_HOC"
	BillBillingFrequencyMixed    BillBillingFrequency = "MIXED"
)

func (r BillBillingFrequency) IsKnown() bool {
	switch r {
	case BillBillingFrequencyDaily, BillBillingFrequencyWeekly, BillBillingFrequencyMonthly, BillBillingFrequencyAnnually, BillBillingFrequencyAdHoc, BillBillingFrequencyMixed:
		return true
	}
	return false
}

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type BillCurrencyConversion struct {
	// Currency to convert from. For example: GBP.
	From string `json:"from,required"`
	// Currency to convert to. For example: USD.
	To string `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier float64                    `json:"multiplier"`
	JSON       billCurrencyConversionJSON `json:"-"`
}

// billCurrencyConversionJSON contains the JSON metadata for the struct
// [BillCurrencyConversion]
type billCurrencyConversionJSON struct {
	From        apijson.Field
	To          apijson.Field
	Multiplier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillCurrencyConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billCurrencyConversionJSON) RawJSON() string {
	return r.raw
}

type BillLineItem struct {
	// The average unit price across all tiers / pricing bands.
	AverageUnitPrice float64 `json:"averageUnitPrice,required"`
	// The currency conversion rate if currency conversion is required for the line
	// item.
	ConversionRate float64 `json:"conversionRate,required"`
	// The converted subtotal amount if currency conversions have been used.
	ConvertedSubtotal float64 `json:"convertedSubtotal,required"`
	// The currency code for the currency used in the line item. For example: USD, GBP,
	// or EUR.
	Currency string `json:"currency,required"`
	// Line item description.
	Description  string                    `json:"description,required"`
	LineItemType BillLineItemsLineItemType `json:"lineItemType,required"`
	// The amount of usage for the line item.
	Quantity float64 `json:"quantity,required"`
	// The subtotal amount for the line item, before any currency conversions.
	Subtotal float64 `json:"subtotal,required"`
	// The unit for the usage data in thie line item. For example: **GB** of disk
	// storage space.
	Unit string `json:"unit,required"`
	// The number of units used for the line item.
	Units float64 `json:"units,required"`
	// The UUID for the line item.
	ID string `json:"id"`
	// The Aggregation ID used for the line item.
	AggregationID string `json:"aggregationId"`
	BalanceID     string `json:"balanceId"`
	// If part of a Parent/Child account billing hierarchy, this is the code for the
	// child Account.
	ChildAccountCode string `json:"childAccountCode"`
	// If part of a Parent/Child account billing hierarchy, this is the child Account
	// UUID.
	ChildAccountID string `json:"childAccountId"`
	// If Commitments _(prepayments)_ are used in the line item, this shows the
	// Commitment UUID.
	CommitmentID string `json:"commitmentId"`
	// The Compound Aggregation ID for the line item if a Compound Aggregation has been
	// used.
	CompoundAggregationID string `json:"compoundAggregationId"`
	// The UUID for the Contract used in the line item.
	ContractID   string            `json:"contractId"`
	CounterID    string            `json:"counterId"`
	CreditTypeID string            `json:"creditTypeId"`
	Group        map[string]string `json:"group"`
	// The UUID of the Meter used in the line item.
	MeterID string `json:"meterId"`
	// The UUID of the PlanGroup, provided the line item used a PlanGroup.
	PlanGroupID string `json:"planGroupId"`
	// The ID of the Plan used for the line item.
	PlanID string `json:"planId"`
	// The UUID of the Pricing used on the line item.
	PricingID   string `json:"pricingId"`
	ProductCode string `json:"productCode"`
	// The UUID of the Product for the line item.
	ProductID string `json:"productId"`
	// The name of the Product for the line item.
	ProductName          string `json:"productName"`
	ReasonID             string `json:"reasonId"`
	ReferencedBillID     string `json:"referencedBillId"`
	ReferencedLineItemID string `json:"referencedLineItemId"`
	// Applies only when segmented Aggregations have been used. The Segment to which
	// the usage data in this line item belongs.
	Segment map[string]string `json:"segment"`
	// The number used for sequential invoices.
	SequenceNumber int64 `json:"sequenceNumber"`
	// The ending date _(exclusive)_ for the service period _(in ISO 8601 format)_.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The starting date _(inclusive)_ for the service period _(in ISO 8601 format)_.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Shows the usage by pricing band for tiered pricing structures.
	UsagePerPricingBand []BillLineItemsUsagePerPricingBand `json:"usagePerPricingBand"`
	JSON                billLineItemJSON                   `json:"-"`
}

// billLineItemJSON contains the JSON metadata for the struct [BillLineItem]
type billLineItemJSON struct {
	AverageUnitPrice       apijson.Field
	ConversionRate         apijson.Field
	ConvertedSubtotal      apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	LineItemType           apijson.Field
	Quantity               apijson.Field
	Subtotal               apijson.Field
	Unit                   apijson.Field
	Units                  apijson.Field
	ID                     apijson.Field
	AggregationID          apijson.Field
	BalanceID              apijson.Field
	ChildAccountCode       apijson.Field
	ChildAccountID         apijson.Field
	CommitmentID           apijson.Field
	CompoundAggregationID  apijson.Field
	ContractID             apijson.Field
	CounterID              apijson.Field
	CreditTypeID           apijson.Field
	Group                  apijson.Field
	MeterID                apijson.Field
	PlanGroupID            apijson.Field
	PlanID                 apijson.Field
	PricingID              apijson.Field
	ProductCode            apijson.Field
	ProductID              apijson.Field
	ProductName            apijson.Field
	ReasonID               apijson.Field
	ReferencedBillID       apijson.Field
	ReferencedLineItemID   apijson.Field
	Segment                apijson.Field
	SequenceNumber         apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UsagePerPricingBand    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BillLineItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billLineItemJSON) RawJSON() string {
	return r.raw
}

type BillLineItemsLineItemType string

const (
	BillLineItemsLineItemTypeStandingCharge            BillLineItemsLineItemType = "STANDING_CHARGE"
	BillLineItemsLineItemTypeUsage                     BillLineItemsLineItemType = "USAGE"
	BillLineItemsLineItemTypeCounterRunningTotalCharge BillLineItemsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BillLineItemsLineItemTypeCounterAdjustmentDebit    BillLineItemsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BillLineItemsLineItemTypeCounterAdjustmentCredit   BillLineItemsLineItemType = "COUNTER_ADJUSTMENT_CREDIT"
	BillLineItemsLineItemTypeUsageCredit               BillLineItemsLineItemType = "USAGE_CREDIT"
	BillLineItemsLineItemTypeMinimumSpend              BillLineItemsLineItemType = "MINIMUM_SPEND"
	BillLineItemsLineItemTypeMinimumSpendRefund        BillLineItemsLineItemType = "MINIMUM_SPEND_REFUND"
	BillLineItemsLineItemTypeCreditDeduction           BillLineItemsLineItemType = "CREDIT_DEDUCTION"
	BillLineItemsLineItemTypeManualAdjustment          BillLineItemsLineItemType = "MANUAL_ADJUSTMENT"
	BillLineItemsLineItemTypeCreditMemo                BillLineItemsLineItemType = "CREDIT_MEMO"
	BillLineItemsLineItemTypeDebitMemo                 BillLineItemsLineItemType = "DEBIT_MEMO"
	BillLineItemsLineItemTypeCommitmentConsumed        BillLineItemsLineItemType = "COMMITMENT_CONSUMED"
	BillLineItemsLineItemTypeCommitmentFee             BillLineItemsLineItemType = "COMMITMENT_FEE"
	BillLineItemsLineItemTypeOverageSurcharge          BillLineItemsLineItemType = "OVERAGE_SURCHARGE"
	BillLineItemsLineItemTypeOverageUsage              BillLineItemsLineItemType = "OVERAGE_USAGE"
	BillLineItemsLineItemTypeBalanceConsumed           BillLineItemsLineItemType = "BALANCE_CONSUMED"
	BillLineItemsLineItemTypeBalanceFee                BillLineItemsLineItemType = "BALANCE_FEE"
)

func (r BillLineItemsLineItemType) IsKnown() bool {
	switch r {
	case BillLineItemsLineItemTypeStandingCharge, BillLineItemsLineItemTypeUsage, BillLineItemsLineItemTypeCounterRunningTotalCharge, BillLineItemsLineItemTypeCounterAdjustmentDebit, BillLineItemsLineItemTypeCounterAdjustmentCredit, BillLineItemsLineItemTypeUsageCredit, BillLineItemsLineItemTypeMinimumSpend, BillLineItemsLineItemTypeMinimumSpendRefund, BillLineItemsLineItemTypeCreditDeduction, BillLineItemsLineItemTypeManualAdjustment, BillLineItemsLineItemTypeCreditMemo, BillLineItemsLineItemTypeDebitMemo, BillLineItemsLineItemTypeCommitmentConsumed, BillLineItemsLineItemTypeCommitmentFee, BillLineItemsLineItemTypeOverageSurcharge, BillLineItemsLineItemTypeOverageUsage, BillLineItemsLineItemTypeBalanceConsumed, BillLineItemsLineItemTypeBalanceFee:
		return true
	}
	return false
}

// Array containing the pricing band information, which shows the details for each
// pricing band or tier.
type BillLineItemsUsagePerPricingBand struct {
	// Usage amount within the band.
	BandQuantity float64 `json:"bandQuantity"`
	// Subtotal amount for the band.
	BandSubtotal float64 `json:"bandSubtotal"`
	// The number of units used within the band.
	BandUnits float64 `json:"bandUnits"`
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
	UnitSubtotal float64                              `json:"unitSubtotal"`
	JSON         billLineItemsUsagePerPricingBandJSON `json:"-"`
}

// billLineItemsUsagePerPricingBandJSON contains the JSON metadata for the struct
// [BillLineItemsUsagePerPricingBand]
type billLineItemsUsagePerPricingBandJSON struct {
	BandQuantity  apijson.Field
	BandSubtotal  apijson.Field
	BandUnits     apijson.Field
	CreditTypeID  apijson.Field
	FixedPrice    apijson.Field
	LowerLimit    apijson.Field
	PricingBandID apijson.Field
	UnitPrice     apijson.Field
	UnitSubtotal  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BillLineItemsUsagePerPricingBand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billLineItemsUsagePerPricingBandJSON) RawJSON() string {
	return r.raw
}

type BillStatus string

const (
	BillStatusPending  BillStatus = "PENDING"
	BillStatusApproved BillStatus = "APPROVED"
)

func (r BillStatus) IsKnown() bool {
	switch r {
	case BillStatusPending, BillStatusApproved:
		return true
	}
	return false
}

type BillApproveResponse struct {
	// A message indicating the success or failure of the Bills' approval, along with
	// relevant details.
	Message string                  `json:"message"`
	JSON    billApproveResponseJSON `json:"-"`
}

// billApproveResponseJSON contains the JSON metadata for the struct
// [BillApproveResponse]
type billApproveResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillApproveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billApproveResponseJSON) RawJSON() string {
	return r.raw
}

type BillSearchResponse struct {
	// An array containing the list of requested Bills.
	Data []Bill `json:"data"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Bills in a paginated list.
	NextToken string                 `json:"nextToken"`
	JSON      billSearchResponseJSON `json:"-"`
}

// billSearchResponseJSON contains the JSON metadata for the struct
// [BillSearchResponse]
type billSearchResponseJSON struct {
	Data        apijson.Field
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSearchResponseJSON) RawJSON() string {
	return r.raw
}

type BillGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Optional filter. An Account ID - returns the Bills for the single specified
	// Account.
	AccountID param.Field[string] `query:"accountId"`
	// The specific date in ISO 8601 format for which you want to retrieve Bills.
	BillDate param.Field[string] `query:"billDate"`
	// Only include Bills with bill dates earlier than this date.
	BillDateEnd param.Field[string] `query:"billDateEnd"`
	// Only include Bills with bill dates equal to or later than this date.
	BillDateStart    param.Field[string] `query:"billDateStart"`
	BillingFrequency param.Field[string] `query:"billingFrequency"`
	// Exclude Line Items
	ExcludeLineItems param.Field[bool] `query:"excludeLineItems"`
	// Only include Bills with external invoice dates earlier than this date.
	ExternalInvoiceDateEnd param.Field[string] `query:"externalInvoiceDateEnd"`
	// Only include Bills with external invoice dates equal to or later than this date.
	ExternalInvoiceDateStart param.Field[string] `query:"externalInvoiceDateStart"`
	// Optional filter. The list of Bill IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// Include Bill Total
	IncludeBillTotal param.Field[bool] `query:"includeBillTotal"`
	// Boolean flag specifying whether to include Bills with "locked" status.
	//
	// - **TRUE** - the list inlcudes "locked" Bills.
	// - **FALSE** - excludes "locked" Bills from the list.
	Locked param.Field[bool] `query:"locked"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Bills in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Bills to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// Only include Bills having the given status
	Status param.Field[BillListParamsStatus] `query:"status"`
}

// URLQuery serializes [BillListParams]'s query parameters as `url.Values`.
func (r BillListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only include Bills having the given status
type BillListParamsStatus string

const (
	BillListParamsStatusPending  BillListParamsStatus = "PENDING"
	BillListParamsStatusApproved BillListParamsStatus = "APPROVED"
)

func (r BillListParamsStatus) IsKnown() bool {
	switch r {
	case BillListParamsStatusPending, BillListParamsStatusApproved:
		return true
	}
	return false
}

type BillDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillApproveParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Use to specify a collection of Bills by their IDs for batch approval
	BillIDs param.Field[[]string] `json:"billIds,required"`
	// List of Account IDs to filter Bills. This allows you to approve Bills for
	// specific Accounts within the Organization.
	AccountIDs param.Field[string] `query:"accountIds"`
	// End date for filtering Bills by external invoice date. Includes Bills with dates
	// earlier than this date.
	ExternalInvoiceDateEnd param.Field[string] `query:"externalInvoiceDateEnd"`
	// Start date for filtering Bills by external invoice date. Includes Bills with
	// dates equal to or later than this date.
	ExternalInvoiceDateStart param.Field[string] `query:"externalInvoiceDateStart"`
}

func (r BillApproveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BillApproveParams]'s query parameters as `url.Values`.
func (r BillApproveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillLatestByAccountParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillLockParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillSearchParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// fromDocument for multi page retrievals
	FromDocument param.Field[int64] `query:"fromDocument"`
	// Search Operator to be used while querying search
	Operator param.Field[BillSearchParamsOperator] `query:"operator"`
	// Number of Commitments to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// Query for data using special syntax. Query parameters should be delimited using
	// $.Allowed comparators are > (greater than), >= (grater than or equal), : (equal), < (less than), <= (less than or equal), ~ (contains). Allowed parameters: accountId, locked, billDate, startDate, endDate, dueDate, billingFrequency, externalInvoiceDateStart, externalInvoiceDateEnd, id, createdBy, dtCreated, lastModifiedBy, ids.Query example: searchQuery=startDate>2023-01-01$accountId:999cb15f-3e8a-4146-b4be-28d0aaedf275.
	// This query is translated into: find bills that startDate is older than
	// 2023-01-01 AND accountId is equal to 999cb15f-3e8a-4146-b4be-28d0aaedf275.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// Name of the parameter on which sorting is performed
	SortBy param.Field[string] `query:"sortBy"`
	// Sorting order
	SortOrder param.Field[BillSearchParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [BillSearchParams]'s query parameters as `url.Values`.
func (r BillSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Search Operator to be used while querying search
type BillSearchParamsOperator string

const (
	BillSearchParamsOperatorAnd BillSearchParamsOperator = "AND"
	BillSearchParamsOperatorOr  BillSearchParamsOperator = "OR"
)

func (r BillSearchParamsOperator) IsKnown() bool {
	switch r {
	case BillSearchParamsOperatorAnd, BillSearchParamsOperatorOr:
		return true
	}
	return false
}

// Sorting order
type BillSearchParamsSortOrder string

const (
	BillSearchParamsSortOrderAsc  BillSearchParamsSortOrder = "ASC"
	BillSearchParamsSortOrderDesc BillSearchParamsSortOrder = "DESC"
)

func (r BillSearchParamsSortOrder) IsKnown() bool {
	switch r {
	case BillSearchParamsSortOrderAsc, BillSearchParamsSortOrderDesc:
		return true
	}
	return false
}

type BillUpdateStatusParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The new status you want to assign to the Bill. Must be one "Pending" or
	// "Approved".
	Status param.Field[BillUpdateStatusParamsStatus] `json:"status,required"`
}

func (r BillUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The new status you want to assign to the Bill. Must be one "Pending" or
// "Approved".
type BillUpdateStatusParamsStatus string

const (
	BillUpdateStatusParamsStatusPending  BillUpdateStatusParamsStatus = "PENDING"
	BillUpdateStatusParamsStatusApproved BillUpdateStatusParamsStatus = "APPROVED"
)

func (r BillUpdateStatusParamsStatus) IsKnown() bool {
	switch r {
	case BillUpdateStatusParamsStatusPending, BillUpdateStatusParamsStatusApproved:
		return true
	}
	return false
}
