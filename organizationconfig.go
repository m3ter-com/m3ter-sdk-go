// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// OrganizationConfigService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationConfigService] method instead.
type OrganizationConfigService struct {
	Options []option.RequestOption
}

// NewOrganizationConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationConfigService(opts ...option.RequestOption) (r *OrganizationConfigService) {
	r = &OrganizationConfigService{}
	r.Options = opts
	return
}

// Retrieve the Organization-wide configuration details.
func (r *OrganizationConfigService) Get(ctx context.Context, orgID string, opts ...option.RequestOption) (res *OrganizationConfig, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/organizationconfig", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Organization-wide configuration details.
func (r *OrganizationConfigService) Update(ctx context.Context, orgID string, body OrganizationConfigUpdateParams, opts ...option.RequestOption) (res *OrganizationConfig, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/organizationconfig", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type OrganizationConfig struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Grace period before bills are auto-approved. Used in combination with the field
	// `autoApproveBillsGracePeriodUnit`.
	AutoApproveBillsGracePeriod int64 `json:"autoApproveBillsGracePeriod"`
	AutoApproveBillsGracePeriodUnit OrganizationConfigAutoApproveBillsGracePeriodUnit `json:"autoApproveBillsGracePeriodUnit"`
	// Specifies whether to auto-generate statements once Bills are _approved_ or
	// _locked_. It will not auto-generate if a bill is in _pending_ state.
	//
	// The default value is **None**.
	//
	// - **None**. Statements will not be auto-generated.
	// - **JSON**. Statements are auto-generated in JSON format.
	// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
	AutoGenerateStatementMode OrganizationConfigAutoGenerateStatementMode `json:"autoGenerateStatementMode"`
	// Prefix to be used for sequential invoice numbers. This will be combined with the
	// `sequenceStartNumber`.
	BillPrefix string `json:"billPrefix"`
	// Specifies whether commitments _(prepayments)_ are billed in advance at the start
	// of each billing period, or billed in arrears at the end of each billing period.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	CommitmentFeeBillInAdvance bool `json:"commitmentFeeBillInAdvance"`
	// Specifies whether to consolidate different billing frequencies onto the same
	// bill.
	//
	// - **TRUE** - consolidate different billing frequencies onto the same bill.
	// - **FALSE** - bills are not consolidated.
	ConsolidateBills bool `json:"consolidateBills"`
	// The id of the user who created this organization config.
	CreatedBy string `json:"createdBy"`
	// The order in which any Prepayment or Balance credit amounts on Accounts are to
	// be drawn-down against for billing. Four options:
	//
	//   - `"PREPAYMENT","BALANCE"`. Draw-down against Prepayment credit before Balance
	//     credit.
	//   - `"BALANCE","PREPAYMENT"`. Draw-down against Balance credit before Prepayment
	//     credit.
	//   - `"PREPAYMENT"`. Only draw-down against Prepayment credit.
	//   - `"BALANCE"`. Only draw-down against Balance credit.
	CreditApplicationOrder []OrganizationConfigCreditApplicationOrder `json:"creditApplicationOrder"`
	// The currency code for the currency used in this Organization. For example: USD,
	// GBP, or EUR.
	Currency string `json:"currency"`
	// Currency conversion rates from Bill currency to Organization currency.
	//
	// For example, if Account is billed in GBP and Organization is set to USD, Bill
	// line items are calculated in GBP and then converted to USD using the defined
	// rate.
	CurrencyConversions []OrganizationConfigCurrencyConversion `json:"currencyConversions"`
	// The first bill date _(in ISO-8601 format)_ for daily billing periods.
	DayEpoch time.Time `json:"dayEpoch" format:"date"`
	// The number of days after the Bill generation date shown on Bills as the due
	// date.
	DaysBeforeBillDue int64 `json:"daysBeforeBillDue"`
	// Organization level default `statementDefinitionId` to be used when there is no
	// statement definition linked to the account.
	//
	// Statement definitions are used to generate bill statements, which are
	// informative backing sheets to invoices.
	DefaultStatementDefinitionID string `json:"defaultStatementDefinitionId"`
	// The DateTime when the organization config was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the organization config was last modified _(in ISO-8601
	// format)_.
	DtLastModified      time.Time                             `json:"dtLastModified" format:"date-time"`
	ExternalInvoiceDate OrganizationConfigExternalInvoiceDate `json:"externalInvoiceDate"`
	// The id of the user who last modified this organization config.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Specifies whether minimum spend amounts are billed in advance at the start of
	// each billing period, or billed in arrears at the end of each billing period.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	MinimumSpendBillInAdvance bool `json:"minimumSpendBillInAdvance"`
	// The first bill date _(in ISO-8601 format)_ for monthly billing periods.
	MonthEpoch time.Time `json:"monthEpoch" format:"date"`
	// Specifies the required interval for updating bills.
	//
	//   - **For portions of an hour (minutes)**. Two options: **0.25** (15 minutes) and
	//     **0.5** (30 minutes).
	//   - **For full hours.** Eight possible values: **1**, **2**, **3**, **4**, **6**,
	//     **8**, **12**, or **24**.
	//   - **Default.** The default is **0**, which disables scheduling.
	ScheduledBillInterval float64 `json:"scheduledBillInterval"`
	// The starting number to be used for sequential invoice numbers. This will be
	// combined with the `billPrefix`.
	SequenceStartNumber int64 `json:"sequenceStartNumber"`
	// Specifies whether the standing charge is billed in advance at the start of each
	// billing period, or billed in arrears at the end of each billing period.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	StandingChargeBillInAdvance bool `json:"standingChargeBillInAdvance"`
	// Specifies whether to supress generating bills that have no line items.
	//
	// - **TRUE** - prevents generating bills with no line items.
	// - **FALSE** - bills are still generated even when they have no line items.
	SuppressedEmptyBills bool `json:"suppressedEmptyBills"`
	// The timezone for the Organization.
	Timezone string `json:"timezone"`
	// The first bill date _(in ISO-8601 format)_ for weekly billing periods.
	WeekEpoch time.Time `json:"weekEpoch" format:"date"`
	// The first bill date _(in ISO-8601 format)_ for yearly billing periods.
	YearEpoch time.Time              `json:"yearEpoch" format:"date"`
	JSON      organizationConfigJSON `json:"-"`
}

// organizationConfigJSON contains the JSON metadata for the struct
// [OrganizationConfig]
type organizationConfigJSON struct {
	ID                              apijson.Field
	Version                         apijson.Field
	AutoApproveBillsGracePeriod     apijson.Field
	AutoApproveBillsGracePeriodUnit apijson.Field
	AutoGenerateStatementMode       apijson.Field
	BillPrefix                      apijson.Field
	CommitmentFeeBillInAdvance      apijson.Field
	ConsolidateBills                apijson.Field
	CreatedBy                       apijson.Field
	CreditApplicationOrder          apijson.Field
	Currency                        apijson.Field
	CurrencyConversions             apijson.Field
	DayEpoch                        apijson.Field
	DaysBeforeBillDue               apijson.Field
	DefaultStatementDefinitionID    apijson.Field
	DtCreated                       apijson.Field
	DtLastModified                  apijson.Field
	ExternalInvoiceDate             apijson.Field
	LastModifiedBy                  apijson.Field
	MinimumSpendBillInAdvance       apijson.Field
	MonthEpoch                      apijson.Field
	ScheduledBillInterval           apijson.Field
	SequenceStartNumber             apijson.Field
	StandingChargeBillInAdvance     apijson.Field
	SuppressedEmptyBills            apijson.Field
	Timezone                        apijson.Field
	WeekEpoch                       apijson.Field
	YearEpoch                       apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *OrganizationConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationConfigJSON) RawJSON() string {
	return r.raw
}

type OrganizationConfigAutoApproveBillsGracePeriodUnit string

const (
	OrganizationConfigAutoApproveBillsGracePeriodUnitMinutes OrganizationConfigAutoApproveBillsGracePeriodUnit = "MINUTES"
	OrganizationConfigAutoApproveBillsGracePeriodUnitHours   OrganizationConfigAutoApproveBillsGracePeriodUnit = "HOURS"
	OrganizationConfigAutoApproveBillsGracePeriodUnitDays    OrganizationConfigAutoApproveBillsGracePeriodUnit = "DAYS"
)

func (r OrganizationConfigAutoApproveBillsGracePeriodUnit) IsKnown() bool {
	switch r {
	case OrganizationConfigAutoApproveBillsGracePeriodUnitMinutes, OrganizationConfigAutoApproveBillsGracePeriodUnitHours, OrganizationConfigAutoApproveBillsGracePeriodUnitDays:
		return true
	}
	return false
}

// Specifies whether to auto-generate statements once Bills are _approved_ or
// _locked_. It will not auto-generate if a bill is in _pending_ state.
//
// The default value is **None**.
//
// - **None**. Statements will not be auto-generated.
// - **JSON**. Statements are auto-generated in JSON format.
// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
type OrganizationConfigAutoGenerateStatementMode string

const (
	OrganizationConfigAutoGenerateStatementModeNone       OrganizationConfigAutoGenerateStatementMode = "NONE"
	OrganizationConfigAutoGenerateStatementModeJson       OrganizationConfigAutoGenerateStatementMode = "JSON"
	OrganizationConfigAutoGenerateStatementModeJsonAndCsv OrganizationConfigAutoGenerateStatementMode = "JSON_AND_CSV"
)

func (r OrganizationConfigAutoGenerateStatementMode) IsKnown() bool {
	switch r {
	case OrganizationConfigAutoGenerateStatementModeNone, OrganizationConfigAutoGenerateStatementModeJson, OrganizationConfigAutoGenerateStatementModeJsonAndCsv:
		return true
	}
	return false
}

type OrganizationConfigCreditApplicationOrder string

const (
	OrganizationConfigCreditApplicationOrderPrepayment OrganizationConfigCreditApplicationOrder = "PREPAYMENT"
	OrganizationConfigCreditApplicationOrderBalance    OrganizationConfigCreditApplicationOrder = "BALANCE"
)

func (r OrganizationConfigCreditApplicationOrder) IsKnown() bool {
	switch r {
	case OrganizationConfigCreditApplicationOrderPrepayment, OrganizationConfigCreditApplicationOrderBalance:
		return true
	}
	return false
}

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type OrganizationConfigCurrencyConversion struct {
	// Currency to convert from. For example: GBP.
	From string `json:"from,required"`
	// Currency to convert to. For example: USD.
	To string `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier float64                                  `json:"multiplier"`
	JSON       organizationConfigCurrencyConversionJSON `json:"-"`
}

// organizationConfigCurrencyConversionJSON contains the JSON metadata for the
// struct [OrganizationConfigCurrencyConversion]
type organizationConfigCurrencyConversionJSON struct {
	From        apijson.Field
	To          apijson.Field
	Multiplier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationConfigCurrencyConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationConfigCurrencyConversionJSON) RawJSON() string {
	return r.raw
}

type OrganizationConfigExternalInvoiceDate string

const (
	OrganizationConfigExternalInvoiceDateLastDayOfArrears     OrganizationConfigExternalInvoiceDate = "LAST_DAY_OF_ARREARS"
	OrganizationConfigExternalInvoiceDateFirstDayOfNextPeriod OrganizationConfigExternalInvoiceDate = "FIRST_DAY_OF_NEXT_PERIOD"
)

func (r OrganizationConfigExternalInvoiceDate) IsKnown() bool {
	switch r {
	case OrganizationConfigExternalInvoiceDateLastDayOfArrears, OrganizationConfigExternalInvoiceDateFirstDayOfNextPeriod:
		return true
	}
	return false
}

type OrganizationConfigUpdateParams struct {
	// The currency code for the Organization. For example: USD, GBP, or EUR:
	//
	//   - This defines the _billing currency_ for the Organization. You can override
	//     this by selecting a different billing currency at individual Account level.
	//   - You must first define the currencies you want to use in your Organization. See
	//     the [Currency](https://www.m3ter.com/docs/api#tag/Currency) section in this
	//     API Reference.
	//
	// **Note:** If you use a different currency as the _pricing currency_ for Plans to
	// set charge rates for Product consumption by an Account, you must define a
	// currency conversion rate from the pricing currency to the billing currency
	// before you run billing for the Account, otherwise billing will fail. See below
	// for the `currencyConversions` request parameter.
	Currency param.Field[string] `json:"currency,required"`
	// Optional setting that defines the billing cycle date for Accounts that are
	// billed daily. Defines the date of the first Bill:
	//
	//   - For example, suppose the Plan you attach to an Account is configured for daily
	//     billing frequency and will apply to the Account from January 1st, 2022 until
	//     June 30th, 2022. If you set a `dayEpoch` date of January 2nd, 2022, then the
	//     first Bill is created for the Account on that date and subsequent Bills are
	//     created for the Account each day following through to the end of the billing
	//     service period.
	//   - The date is in ISO-8601 format.
	DayEpoch param.Field[string] `json:"dayEpoch,required"`
	// Enter the number of days after the Bill generation date that you want to show on
	// Bills as the due date.
	//
	// **Note:** If you define `daysBeforeBillDue` at individual Account level, this
	// will take precedence over any `daysBeforeBillDue` setting defined at
	// Organization level.
	DaysBeforeBillDue param.Field[int64] `json:"daysBeforeBillDue,required"`
	// Optional setting that defines the billing cycle date for Accounts that are
	// billed monthly. Defines the date of the first Bill and then acts as reference
	// for when subsequent Bills are created for the Account:
	//
	//   - For example, suppose the Plan you attach to an Account is configured for
	//     monthly billing frequency and will apply to the Account from January 1st, 2022
	//     until June 30th, 2022. If you set a `monthEpoch` date of January 15th, 2022,
	//     then the first Bill is created for the Account on that date and subsequent
	//     Bills are created for the Account on the 15th of each month following through
	//     to the end of the billing service period - February 15th, March 15th, and so
	//     on.
	//   - The date is in ISO-8601 format.
	MonthEpoch param.Field[string] `json:"monthEpoch,required"`
	// Sets the timezone for the Organization.
	Timezone param.Field[string] `json:"timezone,required"`
	// Optional setting that defines the billing cycle date for Accounts that are
	// billed weekly. Defines the date of the first Bill and then acts as reference for
	// when subsequent Bills are created for the Account:
	//
	//   - For example, suppose the Plan you attach to an Account is configured for
	//     weekly billing frequency and will apply to the Account from January 1st, 2022
	//     until June 30th, 2022. If you set a `weekEpoch` date of January 15th, 2022,
	//     which falls on a Saturday, then the first Bill is created for the Account on
	//     that date and subsequent Bills are created for the Account on Saturday of each
	//     week following through to the end of the billing service period.
	//   - The date is in ISO-8601 format.
	WeekEpoch param.Field[string] `json:"weekEpoch,required"`
	// Optional setting that defines the billing cycle date for Accounts that are
	// billed yearly. Defines the date of the first Bill and then acts as reference for
	// when subsequent Bills are created for the Account:
	//
	//   - For example, suppose the Plan you attach to an Account is configured for
	//     yearly billing frequency and will apply to the Account from January 1st, 2022
	//     until January 15th, 2028. If you set a `yearEpoch` date of January 1st, 2023,
	//     then the first Bill is created for the Account on that date and subsequent
	//     Bills are created for the Account on January 1st of each year following
	//     through to the end of the billing service period - January 1st, 2023, January
	//     1st, 2024 and so on.
	//   - The date is in ISO-8601 format.
	YearEpoch param.Field[string] `json:"yearEpoch,required"`
	// Grace period before bills are auto-approved. Used in combination with
	// `autoApproveBillsGracePeriodUnit` parameter.
	//
	// **Note:** When used in combination with `autoApproveBillsGracePeriodUnit`
	// enables auto-approval of Bills for Organization, which occurs when the specified
	// time period has elapsed after Bill generation.
	AutoApproveBillsGracePeriod param.Field[int64] `json:"autoApproveBillsGracePeriod"`
	// Time unit of grace period before bills are auto-approved. Used in combination
	// with `autoApproveBillsGracePeriod` parameter. Allowed options are MINUTES,
	// HOURS, or DAYS.
	//
	// **Note:** When used in combination with `autoApproveBillsGracePeriod` enables
	// auto-approval of Bills for Organization, which occurs when the specified time
	// period has elapsed after Bill generation.
	AutoApproveBillsGracePeriodUnit param.Field[string] `json:"autoApproveBillsGracePeriodUnit"`
	// Specify whether to auto-generate statements once Bills are _approved_ or
	// _locked_. It will not auto-generate if a bill is in _pending_ state.
	//
	// The default value is **None**.
	//
	// - **None**. Statements will not be auto-generated.
	// - **JSON**. Statements are auto-generated in JSON format.
	// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
	AutoGenerateStatementMode param.Field[OrganizationConfigUpdateParamsAutoGenerateStatementMode] `json:"autoGenerateStatementMode"`
	// Prefix to be used for sequential invoice numbers. This will be combined with the
	// `sequenceStartNumber`.
	//
	// **NOTES:**
	//
	//   - If you do not define a `billPrefix`, a default will be used in the Console for
	//     the Bill **REFERENCE** number. This default will concatenate **INV-** with the
	//     last four characters of the `billId`.
	//   - If you do not define a `billPrefix`, the Bill response schema for API calls
	//     that retrieve Bill data will not contain a `sequentialInvoiceNumber`.
	BillPrefix param.Field[string] `json:"billPrefix"`
	// Boolean setting to specify whether commitments _(prepayments)_ are billed in
	// advance at the start of each billing period, or billed in arrears at the end of
	// each billing period.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	CommitmentFeeBillInAdvance param.Field[bool] `json:"commitmentFeeBillInAdvance"`
	// Boolean setting to consolidate different billing frequencies onto the same bill.
	//
	// - **TRUE** - consolidate different billing frequencies onto the same bill.
	// - **FALSE** - bills are not consolidated.
	ConsolidateBills param.Field[bool] `json:"consolidateBills"`
	// Define the order in which any Prepayment or Balance amounts on Accounts are to
	// be drawn-down against for billing. Four options:
	//
	//   - `"PREPAYMENT","BALANCE"`. Draw-down against Prepayment credit before Balance
	//     credit.
	//   - `"BALANCE","PREPAYMENT"`. Draw-down against Balance credit before Prepayment
	//     credit.
	//   - `"PREPAYMENT"`. Only draw-down against Prepayment credit.
	//   - `"BALANCE"`. Only draw-down against Balance credit.
	//
	// **NOTE:** You can override this Organization-level setting for credit
	// application order at the level of an individual Account.
	CreditApplicationOrder param.Field[[]OrganizationConfigUpdateParamsCreditApplicationOrder] `json:"creditApplicationOrder"`
	// Define currency conversion rates from _pricing currency_ to _billing currency_:
	//
	//   - You can use the `currency` request parameter with this call to define the
	//     billing currency for your Organization - see above.
	//   - You can also define a billing currency at the individual Account level and
	//     this will override the Organization billing currency.
	//   - A Plan used to set Product consumption charge rates on an Account might use a
	//     different pricing currency. At billing, charges are calculated in the pricing
	//     currency and then converted into billing currency amounts to appear on Bills.
	//     If you haven't defined a currency conversion rate from pricing to billing
	//     currency, billing will fail for the Account.
	CurrencyConversions param.Field[[]OrganizationConfigUpdateParamsCurrencyConversion] `json:"currencyConversions"`
	// Organization level default `statementDefinitionId` to be used when there is no
	// statement definition linked to the account.
	//
	// Statement definitions are used to generate bill statements, which are
	// informative backing sheets to invoices.
	DefaultStatementDefinitionID param.Field[string] `json:"defaultStatementDefinitionId"`
	// Date to use for the invoice date. Allowed values are `FIRST_DAY_OF_NEXT_PERIOD`
	// or `LAST_DAY_OF_ARREARS`.
	ExternalInvoiceDate param.Field[string] `json:"externalInvoiceDate"`
	// Boolean setting to specify whether minimum spend amounts are billed in advance
	// at the start of each billing period, or billed in arrears at the end of each
	// billing period.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Sets the required interval for updating bills. It is an optional parameter that
	// can be set as:
	//
	//   - **For portions of an hour (minutes)**. Two options: **0.25** (15 minutes) and
	//     **0.5** (30 minutes).
	//   - **For full hours.** Enter **1** for every hour, **2** for every two hours, and
	//     so on. Eight options: **1**, **2**, **3**, **4**, **6**, **8**, **12**, or
	//     **24**.
	//   - **Default.** The default is **0**, which disables scheduling.
	ScheduledBillInterval param.Field[float64] `json:"scheduledBillInterval"`
	// The starting number to be used for sequential invoice numbers. This will be
	// combined with the `billPrefix`.
	//
	// For example, if you define `billPrefix` to be **INVOICE-** and you set the
	// `seqenceStartNumber` as **100**, the first Bill created after updating your
	// Organization Configuration will have a `sequentialInvoiceNumber` assigned of
	// **INVOICE-101**. Subsequent Bills created will be numbered in time sequence for
	// their initial creation date/time.
	SequenceStartNumber param.Field[int64] `json:"sequenceStartNumber"`
	// Boolean setting to specify whether the standing charge is billed in advance at
	// the start of each billing period, or billed in arrears at the end of each
	// billing period.
	//
	// - **TRUE** - bill in advance _(start of each billing period)_.
	// - **FALSE** - bill in arrears _(end of each billing period)_.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Boolean setting that supresses generating bills that have no line items.
	//
	// - **TRUE** - prevents generating bills with no line items.
	// - **FALSE** - bills are still generated even when they have no line items.
	SuppressedEmptyBills param.Field[bool] `json:"suppressedEmptyBills"`
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

func (r OrganizationConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify whether to auto-generate statements once Bills are _approved_ or
// _locked_. It will not auto-generate if a bill is in _pending_ state.
//
// The default value is **None**.
//
// - **None**. Statements will not be auto-generated.
// - **JSON**. Statements are auto-generated in JSON format.
// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
type OrganizationConfigUpdateParamsAutoGenerateStatementMode string

const (
	OrganizationConfigUpdateParamsAutoGenerateStatementModeNone       OrganizationConfigUpdateParamsAutoGenerateStatementMode = "NONE"
	OrganizationConfigUpdateParamsAutoGenerateStatementModeJson       OrganizationConfigUpdateParamsAutoGenerateStatementMode = "JSON"
	OrganizationConfigUpdateParamsAutoGenerateStatementModeJsonAndCsv OrganizationConfigUpdateParamsAutoGenerateStatementMode = "JSON_AND_CSV"
)

func (r OrganizationConfigUpdateParamsAutoGenerateStatementMode) IsKnown() bool {
	switch r {
	case OrganizationConfigUpdateParamsAutoGenerateStatementModeNone, OrganizationConfigUpdateParamsAutoGenerateStatementModeJson, OrganizationConfigUpdateParamsAutoGenerateStatementModeJsonAndCsv:
		return true
	}
	return false
}

type OrganizationConfigUpdateParamsCreditApplicationOrder string

const (
	OrganizationConfigUpdateParamsCreditApplicationOrderPrepayment OrganizationConfigUpdateParamsCreditApplicationOrder = "PREPAYMENT"
	OrganizationConfigUpdateParamsCreditApplicationOrderBalance    OrganizationConfigUpdateParamsCreditApplicationOrder = "BALANCE"
)

func (r OrganizationConfigUpdateParamsCreditApplicationOrder) IsKnown() bool {
	switch r {
	case OrganizationConfigUpdateParamsCreditApplicationOrderPrepayment, OrganizationConfigUpdateParamsCreditApplicationOrderBalance:
		return true
	}
	return false
}

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type OrganizationConfigUpdateParamsCurrencyConversion struct {
	// Currency to convert from. For example: GBP.
	From param.Field[string] `json:"from,required"`
	// Currency to convert to. For example: USD.
	To param.Field[string] `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier param.Field[float64] `json:"multiplier"`
}

func (r OrganizationConfigUpdateParamsCurrencyConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
