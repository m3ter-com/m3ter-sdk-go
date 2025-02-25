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

// BillJobService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillJobService] method instead.
type BillJobService struct {
	Options []option.RequestOption
}

// NewBillJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBillJobService(opts ...option.RequestOption) (r *BillJobService) {
	r = &BillJobService{}
	r.Options = opts
	return
}

// Create a new BillJob to handle asynchronous bill calculations for a specific
// Organization.
//
// This operation allows you to initiate the processing of bills according to
// specified parameters. For example, create a BillJob to run only those bills
// where `billingFrequency` is `MONTHLY`. Note that if you want to run a BillJob
// for all billing frequencies, simply omit the `billingFrequency` request
// parameter.
//
// Once created, the BillJob's progress can be monitored:
//
//   - In the Running Tasks panel in the m3ter Console - for more details, see
//     [Running Bills Manually](https://www.m3ter.com/docs/guides/billing-and-usage-data/running-viewing-and-managing-bills/running-bills-and-viewing-bill-details#running-bills-manually)
//   - Queried using the
//     [List BillJobs](https://www.m3ter.com/docs/api#tag/BillJob/operation/ListBillJobs)
//     operation.
//
// **NOTES:**
//
//   - **Consolidated bills**. If you've already run billing with the Consolidate
//     bills option disabled for your Organization but you then enable it, subsequent
//     Bills for specific bill dates will now start afresh and not update earlier
//     non-consolidated Bills for the same bill date. To avoid any billing conflicts,
//     you might want to archive these earlier versions or delete them entirely.
//   - **Maximum concurrent BillJobs**. If you already have 10 BillJobs currently
//     running, and try to create another one, you'll get an HTTP 429 response (Too
//     many requests). When one of the existing BillJobs has completed, you'll be
//     able to submit another job
func (r *BillJobService) New(ctx context.Context, params BillJobNewParams, opts ...option.RequestOption) (res *BillJob, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/billjobs", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Bill Job for the given UUID.
func (r *BillJobService) Get(ctx context.Context, id string, query BillJobGetParams, opts ...option.RequestOption) (res *BillJob, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/billjobs/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of BillJobs.
//
// This endpoint retrieves a list of BillJobs for a specified organization. The
// list can be paginated for easier management, and allows you to query and filter
// based on various parameters, such as BillJob `status` and whether or not BillJob
// remains `active`.
func (r *BillJobService) List(ctx context.Context, params BillJobListParams, opts ...option.RequestOption) (res *pagination.Cursor[BillJob], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/billjobs", params.OrgID)
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

// Retrieve a list of BillJobs.
//
// This endpoint retrieves a list of BillJobs for a specified organization. The
// list can be paginated for easier management, and allows you to query and filter
// based on various parameters, such as BillJob `status` and whether or not BillJob
// remains `active`.
func (r *BillJobService) ListAutoPaging(ctx context.Context, params BillJobListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[BillJob] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Cancel an ongoing BillJob for the given Organization and BillJob UUID.
//
// This endpoint allows you to halt the processing of a specific BillJob, which
// might be necessary if there are changes in billing requirements or other
// operational considerations.
func (r *BillJobService) Cancel(ctx context.Context, id string, body BillJobCancelParams, opts ...option.RequestOption) (res *BillJob, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/billjobs/%s/cancel", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create a new BillJob specifically to recalculate existing bills for a given
// Organization.
//
// This operation is essential when adjustments or corrections are required in
// previously calculated bills. The recalculated bills when the BillJob is complete
// can be checked in the m3ter Console Bill Management page or queried by using the
// [List Bills](https://www.m3ter.com/docs/api#tag/Bill/operation/ListBills)
// operation.
//
// **NOTE:**
//
//   - **Response Schema**. The response schema for this call is dynamic. This means
//     that the response might not contain all of the parameters listed. If set to
//     null,the parameter is hidden to help simplify the output as well as to reduce
//     its size and improve performance.
func (r *BillJobService) Recalculate(ctx context.Context, params BillJobRecalculateParams, opts ...option.RequestOption) (res *BillJob, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/billjobs/recalculate", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BillJob struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// An array of UUIDs representing the end customer Accounts associated with the
	// BillJob.
	AccountIDs []string `json:"accountIds"`
	// The specific billing date _(in ISO 8601 format)_, determining when the Bill was
	// generated.
	//
	// For example: `"2023-01-24"`.
	BillDate time.Time `json:"billDate" format:"date"`
	// How often Bills are issued - used in conjunction with `billingFrequency`.
	//
	// For example, if `billingFrequency` is set to Monthly and `billFrequencyInterval`
	// is set to 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// An array of Bill IDs related to the BillJob, providing references to the
	// specific Bills generated.
	BillIDs []string `json:"billIds"`
	// Defines how often Bills are generated.
	//
	//   - **Daily**. Starting at midnight each day, covering a twenty-four hour period
	//     following.
	//
	//   - **Weekly**. Starting at midnight on a Monday morning covering the seven-day
	//     period following.
	//
	//   - **Monthly**. Starting at midnight on the morning of the first day of each
	//     month covering the entire calendar month following.
	//
	//   - **Annually**. Starting at midnight on the morning of the first day of each
	//     year covering the entire calendar year following.
	//
	//   - **Ad_Hoc**. Use this setting when a custom billing schedule is used for
	//     billing an Account, such as for billing of Prepayment/Commitment fees using a
	//     custom billing schedule.
	BillingFrequency BillJobBillingFrequency `json:"billingFrequency"`
	// The unique identifier (UUID) for the user who created the BillJob.
	CreatedBy string `json:"createdBy"`
	// An array of currency conversion rates from Bill currency to Organization
	// currency. For example, if Account is billed in GBP and Organization is set to
	// USD, Bill line items are calculated in GBP and then converted to USD using the
	// defined rate.
	CurrencyConversions []BillJobCurrencyConversion `json:"currencyConversions"`
	// The starting date _(epoch)_ for Daily billing frequency _(in ISO 8601 format)_,
	// determining the first Bill date for daily Bills.
	DayEpoch time.Time `json:"dayEpoch" format:"date"`
	// The date and time _(in ISO 8601 format)_ when the BillJob was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the BillJob was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The due date _(in ISO 8601 format)_ for payment of the Bill.
	//
	// For example: `"2023-02-24"`.
	DueDate time.Time `json:"dueDate" format:"date"`
	// For accounting purposes, the date set at Organization level to use for external
	// invoicing with respect to billing periods - two options:
	//
	//   - `FIRST_DAY_OF_NEXT_PERIOD` _(Default)_. Used when you want to recognize usage
	//     revenue in the following period.
	//   - `LAST_DAY_OF_ARREARS`. Used when you want to recognize usage revenue in the
	//     same period that it's consumed, instead of in the following period.
	//
	// For example, if the retrieved Bill was on a monthly billing frequency and the
	// billing period for the Bill is September 2023 and the _External invoice date_ is
	// set at `FIRST_DAY_OF_NEXT_PERIOD`, then the `externalInvoiceDate` will be
	// `"2023-10-01"`.
	ExternalInvoiceDate time.Time `json:"externalInvoiceDate" format:"date"`
	// Specifies the date _(in ISO 8601 format)_ of the last day in the billing period,
	// defining the time range for the associated Bills.
	//
	// For example: `"2023-03-24"`.
	LastDateInBillingPeriod time.Time `json:"lastDateInBillingPeriod" format:"date"`
	// The unique identifier (UUID) for the user who last modified this BillJob.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The starting date _(epoch)_ for Monthly billing frequency _(in ISO 8601
	// format)_, determining the first Bill date for monthly Bills.
	MonthEpoch time.Time `json:"monthEpoch" format:"date"`
	// The number of pending actions or calculations within the BillJob.
	Pending int64 `json:"pending"`
	// The current status of the BillJob, indicating its progress or completion state.
	Status BillJobStatus `json:"status"`
	// The currency code used for the Bill, such as USD, GBP, or EUR.
	TargetCurrency string `json:"targetCurrency"`
	// Specifies the time zone used for the generated Bills, ensuring alignment with
	// the local time zone.
	Timezone string `json:"timezone"`
	// The total number of Bills or calculations related to the BillJob.
	Total int64 `json:"total"`
	// Specifies the type of BillJob.
	//
	// - **CREATE** Returned for a _Create BillJob_ call.
	// - **RECALCULATE** Returned for a successful _Create Recalculation BillJob_ call.
	Type BillJobType `json:"type"`
	// The starting date _(epoch)_ for Weekly billing frequency _(in ISO 8601 format)_,
	// determining the first Bill date for weekly Bills.
	WeekEpoch time.Time `json:"weekEpoch" format:"date"`
	// The starting date _(epoch)_ for Yearly billing frequency _(in ISO 8601 format)_,
	// determining the first Bill date for yearly Bills.
	YearEpoch time.Time   `json:"yearEpoch" format:"date"`
	JSON      billJobJSON `json:"-"`
}

// billJobJSON contains the JSON metadata for the struct [BillJob]
type billJobJSON struct {
	ID                      apijson.Field
	Version                 apijson.Field
	AccountIDs              apijson.Field
	BillDate                apijson.Field
	BillFrequencyInterval   apijson.Field
	BillIDs                 apijson.Field
	BillingFrequency        apijson.Field
	CreatedBy               apijson.Field
	CurrencyConversions     apijson.Field
	DayEpoch                apijson.Field
	DtCreated               apijson.Field
	DtLastModified          apijson.Field
	DueDate                 apijson.Field
	ExternalInvoiceDate     apijson.Field
	LastDateInBillingPeriod apijson.Field
	LastModifiedBy          apijson.Field
	MonthEpoch              apijson.Field
	Pending                 apijson.Field
	Status                  apijson.Field
	TargetCurrency          apijson.Field
	Timezone                apijson.Field
	Total                   apijson.Field
	Type                    apijson.Field
	WeekEpoch               apijson.Field
	YearEpoch               apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BillJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billJobJSON) RawJSON() string {
	return r.raw
}

// Defines how often Bills are generated.
//
//   - **Daily**. Starting at midnight each day, covering a twenty-four hour period
//     following.
//
//   - **Weekly**. Starting at midnight on a Monday morning covering the seven-day
//     period following.
//
//   - **Monthly**. Starting at midnight on the morning of the first day of each
//     month covering the entire calendar month following.
//
//   - **Annually**. Starting at midnight on the morning of the first day of each
//     year covering the entire calendar year following.
//
//   - **Ad_Hoc**. Use this setting when a custom billing schedule is used for
//     billing an Account, such as for billing of Prepayment/Commitment fees using a
//     custom billing schedule.
type BillJobBillingFrequency string

const (
	BillJobBillingFrequencyDaily    BillJobBillingFrequency = "DAILY"
	BillJobBillingFrequencyWeekly   BillJobBillingFrequency = "WEEKLY"
	BillJobBillingFrequencyMonthly  BillJobBillingFrequency = "MONTHLY"
	BillJobBillingFrequencyAnnually BillJobBillingFrequency = "ANNUALLY"
	BillJobBillingFrequencyAdHoc    BillJobBillingFrequency = "AD_HOC"
)

func (r BillJobBillingFrequency) IsKnown() bool {
	switch r {
	case BillJobBillingFrequencyDaily, BillJobBillingFrequencyWeekly, BillJobBillingFrequencyMonthly, BillJobBillingFrequencyAnnually, BillJobBillingFrequencyAdHoc:
		return true
	}
	return false
}

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type BillJobCurrencyConversion struct {
	// Currency to convert from. For example: GBP.
	From string `json:"from,required"`
	// Currency to convert to. For example: USD.
	To string `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier float64                       `json:"multiplier"`
	JSON       billJobCurrencyConversionJSON `json:"-"`
}

// billJobCurrencyConversionJSON contains the JSON metadata for the struct
// [BillJobCurrencyConversion]
type billJobCurrencyConversionJSON struct {
	From        apijson.Field
	To          apijson.Field
	Multiplier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillJobCurrencyConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billJobCurrencyConversionJSON) RawJSON() string {
	return r.raw
}

// The current status of the BillJob, indicating its progress or completion state.
type BillJobStatus string

const (
	BillJobStatusPending      BillJobStatus = "PENDING"
	BillJobStatusInitializing BillJobStatus = "INITIALIZING"
	BillJobStatusRunning      BillJobStatus = "RUNNING"
	BillJobStatusComplete     BillJobStatus = "COMPLETE"
	BillJobStatusCancelled    BillJobStatus = "CANCELLED"
)

func (r BillJobStatus) IsKnown() bool {
	switch r {
	case BillJobStatusPending, BillJobStatusInitializing, BillJobStatusRunning, BillJobStatusComplete, BillJobStatusCancelled:
		return true
	}
	return false
}

// Specifies the type of BillJob.
//
// - **CREATE** Returned for a _Create BillJob_ call.
// - **RECALCULATE** Returned for a successful _Create Recalculation BillJob_ call.
type BillJobType string

const (
	BillJobTypeCreate      BillJobType = "CREATE"
	BillJobTypeRecalculate BillJobType = "RECALCULATE"
)

func (r BillJobType) IsKnown() bool {
	switch r {
	case BillJobTypeCreate, BillJobTypeRecalculate:
		return true
	}
	return false
}

type BillJobNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// An array of UUIDs representing the end customer Accounts associated with the
	// BillJob.
	AccountIDs param.Field[[]string] `json:"accountIds"`
	// The specific billing date _(in ISO 8601 format)_, determining when the Bill was
	// generated.
	//
	// For example: `"2023-01-24"`.
	BillDate param.Field[time.Time] `json:"billDate" format:"date"`
	// How often Bills are issued - used in conjunction with `billingFrequency`.
	//
	// For example, if `billingFrequency` is set to Monthly and `billFrequencyInterval`
	// is set to 3, Bills are issued every three months.
	BillFrequencyInterval param.Field[int64] `json:"billFrequencyInterval"`
	// How often Bills are generated.
	//
	//   - **Daily**. Starting at midnight each day, covering a twenty-four hour period
	//     following.
	//
	//   - **Weekly**. Starting at midnight on a Monday morning covering the seven-day
	//     period following.
	//
	//   - **Monthly**. Starting at midnight on the morning of the first day of each
	//     month covering the entire calendar month following.
	//
	//   - **Annually**. Starting at midnight on the morning of the first day of each
	//     year covering the entire calendar year following.
	//
	//   - **Ad_Hoc**. Use this setting when a custom billing schedule is used for
	//     billing an Account, such as for billing of Prepayment/Commitment fees using a
	//     custom billing schedule.
	BillingFrequency param.Field[BillJobNewParamsBillingFrequency] `json:"billingFrequency"`
	// An array of currency conversion rates from Bill currency to Organization
	// currency. For example, if Account is billed in GBP and Organization is set to
	// USD, Bill line items are calculated in GBP and then converted to USD using the
	// defined rate.
	CurrencyConversions param.Field[[]BillJobNewParamsCurrencyConversion] `json:"currencyConversions"`
	// The starting date _(epoch)_ for Daily billing frequency _(in ISO 8601 format)_,
	// determining the first Bill date for daily Bills.
	DayEpoch param.Field[time.Time] `json:"dayEpoch" format:"date"`
	// The due date _(in ISO 8601 format)_ for payment of the Bill.
	//
	// For example: `"2023-02-24"`.
	DueDate param.Field[time.Time] `json:"dueDate" format:"date"`
	// For accounting purposes, the date set at Organization level to use for external
	// invoicing with respect to billing periods - two options:
	//
	//   - `FIRST_DAY_OF_NEXT_PERIOD` _(Default)_. Used when you want to recognize usage
	//     revenue in the following period.
	//   - `LAST_DAY_OF_ARREARS`. Used when you want to recognize usage revenue in the
	//     same period that it's consumed, instead of in the following period.
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
	ExternalInvoiceDate param.Field[time.Time] `json:"externalInvoiceDate" format:"date"`
	// Specifies the date _(in ISO 8601 format)_ of the last day in the billing period,
	// defining the time range for the associated Bills.
	//
	// For example: `"2023-03-24"`.
	LastDateInBillingPeriod param.Field[time.Time] `json:"lastDateInBillingPeriod" format:"date"`
	// The starting date _(epoch)_ for Monthly billing frequency _(in ISO 8601
	// format)_, determining the first Bill date for monthly Bills.
	MonthEpoch param.Field[time.Time] `json:"monthEpoch" format:"date"`
	// The currency code used for the Bill, such as USD, GBP, or EUR.
	TargetCurrency param.Field[string] `json:"targetCurrency"`
	// Specifies the time zone used for the generated Bills, ensuring alignment with
	// the local time zone.
	Timezone param.Field[string] `json:"timezone"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - _do not use
	//     for Create_. On initial Create, version is set at 1 and listed in the
	//     response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
	// The starting date _(epoch)_ for Weekly billing frequency _(in ISO 8601 format)_,
	// determining the first Bill date for weekly Bills.
	WeekEpoch param.Field[time.Time] `json:"weekEpoch" format:"date"`
	// The starting date _(epoch)_ for Yearly billing frequency _(in ISO 8601 format)_,
	// determining the first Bill date for yearly Bills.
	YearEpoch param.Field[time.Time] `json:"yearEpoch" format:"date"`
}

func (r BillJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often Bills are generated.
//
//   - **Daily**. Starting at midnight each day, covering a twenty-four hour period
//     following.
//
//   - **Weekly**. Starting at midnight on a Monday morning covering the seven-day
//     period following.
//
//   - **Monthly**. Starting at midnight on the morning of the first day of each
//     month covering the entire calendar month following.
//
//   - **Annually**. Starting at midnight on the morning of the first day of each
//     year covering the entire calendar year following.
//
//   - **Ad_Hoc**. Use this setting when a custom billing schedule is used for
//     billing an Account, such as for billing of Prepayment/Commitment fees using a
//     custom billing schedule.
type BillJobNewParamsBillingFrequency string

const (
	BillJobNewParamsBillingFrequencyDaily    BillJobNewParamsBillingFrequency = "DAILY"
	BillJobNewParamsBillingFrequencyWeekly   BillJobNewParamsBillingFrequency = "WEEKLY"
	BillJobNewParamsBillingFrequencyMonthly  BillJobNewParamsBillingFrequency = "MONTHLY"
	BillJobNewParamsBillingFrequencyAnnually BillJobNewParamsBillingFrequency = "ANNUALLY"
	BillJobNewParamsBillingFrequencyAdHoc    BillJobNewParamsBillingFrequency = "AD_HOC"
)

func (r BillJobNewParamsBillingFrequency) IsKnown() bool {
	switch r {
	case BillJobNewParamsBillingFrequencyDaily, BillJobNewParamsBillingFrequencyWeekly, BillJobNewParamsBillingFrequencyMonthly, BillJobNewParamsBillingFrequencyAnnually, BillJobNewParamsBillingFrequencyAdHoc:
		return true
	}
	return false
}

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type BillJobNewParamsCurrencyConversion struct {
	// Currency to convert from. For example: GBP.
	From param.Field[string] `json:"from,required"`
	// Currency to convert to. For example: USD.
	To param.Field[string] `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier param.Field[float64] `json:"multiplier"`
}

func (r BillJobNewParamsCurrencyConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BillJobGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillJobListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Boolean filter to retrieve only active BillJobs and exclude completed or
	// cancelled BillJobs from the results.
	//
	// - TRUE - only active BillJobs.
	// - FALSE - all BillJobs including completed and cancelled BillJobs.
	Active param.Field[string] `query:"active"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// BillJobs in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of BillJobs to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// Filter BillJobs by specific status. Allows for targeted retrieval of BillJobs
	// based on their current processing status.
	//
	// Possible states are:
	//
	// - PENDING
	// - INITIALIZING
	// - RUNNING
	// - COMPLETE
	// - CANCELLED
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [BillJobListParams]'s query parameters as `url.Values`.
func (r BillJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BillJobCancelParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillJobRecalculateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The array of unique identifiers (UUIDs) for the Bills which are to be
	// recalculated.
	BillIDs param.Field[[]string] `json:"billIds,required"`
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

func (r BillJobRecalculateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
