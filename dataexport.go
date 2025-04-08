// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// DataExportService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataExportService] method instead.
type DataExportService struct {
	Options      []option.RequestOption
	Destinations *DataExportDestinationService
	Jobs         *DataExportJobService
	Schedules    *DataExportScheduleService
}

// NewDataExportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataExportService(opts ...option.RequestOption) (r *DataExportService) {
	r = &DataExportService{}
	r.Options = opts
	r.Destinations = NewDataExportDestinationService(opts...)
	r.Jobs = NewDataExportJobService(opts...)
	r.Schedules = NewDataExportScheduleService(opts...)
	return
}

// Trigger an ad-hoc Data Export. Each ad-hoc Export can be configured for
// exporting _only one of_ either Usage or Operational data:
//
// **Operational Data Exports**.
//
//   - **Entity Types**. Use the `operationalDataTypes` parameter to specify the
//     entities whose operational data you want to include in the ad-hoc export.
//   - **Export Files**. For each of the entity types you select, when the ad-hoc
//     export runs a separate file is compiled containing the operational data for
//     all entities of that type that exist in your Organization
//
// **Usage Data Exports**.
//
//   - **Meters/Accounts**. Select the Meters and Accounts whose usage data you want
//     to include in the ad-hoc export.
//   - **Aggregated or non-aggregated data**:
//
//  1. If you _don't want to aggregate_ the usage data collected by the selected
//     Meters, use **ORIGINAL** for `aggregationFrequency`, which is the _default_.
//     This means the raw usage data collected by any type of Data Fields and the
//     values for any Derived Fields on the selected Meters will be included in the
//     ad-hoc export.
//  2. If you _do want to aggregate_ the usage data collected by the selected
//     Meters, use one of the other options for `aggregationFrequency`: **HOUR**,
//     **DAY**, **WEEK**, or **MONTH**. You _must_ then also specified an
//     `aggregation` method to be used on the usage data before export. Importantly,
//     if you do aggregate usage data, only the usage data collected by any numeric
//     Data Fields on the selected Meters - those of type **MEASURE**, **INCOME**,
//     or **COST** - will be included in the ad-hoc export.
//
// **Date Range for Operational Data Exports**. To restrict the operational data
// included in the export by a date/time range, use the `startDate` and `endDate`
// date/time request parameters to specify the period. Constraints:
//
//   - `startDate` must be before `endDate`.
//   - `startDate` with no `endDate` is valid.
//   - No `startDate` with `endDate` is valid.
//   - `endDate` must be before present date/time.
//   - Both are optional and if neither is defined, the export includes all data for
//     selected entities.
//
// **Date Range for Usage Data Exports**. To restrict the usage data included in
// the export by date/time range, use the `timePeriod` request parameter to define
// a set date range. Alternatively, define a custom date range using the
// `startDate` and `endDate` date/time parameters:
//
//   - Both `startDate` and `endDate` are required.
//   - You cannot use a `startDate` earlier than 35 days in the past.
//   - The `endDate` is valid up to tomorrow at 00:00.
//   - You must define a Date Range using **timePeriod** or **startDate/endDate**,
//     but they are mutually exclusive and you cannot use them together.
//
// **NOTE:** You can use the ExportJob `id` returned to check the status of the
// triggered ad-hoc export. See the
// [ExportJob](https://www.m3ter.com/docs/api#tag/ExportJob) section of this API
// Reference.
func (r *DataExportService) NewAdhoc(ctx context.Context, params DataExportNewAdhocParams, opts ...option.RequestOption) (res *AdHocResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/dataexports/adhoc", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AdHocOperationalDataRequestParam struct {
	// The list of the operational data types should be exported for.
	OperationalDataTypes param.Field[[]AdHocOperationalDataRequestOperationalDataType] `json:"operationalDataTypes,required"`
	SourceType           param.Field[AdHocOperationalDataRequestSourceType]            `json:"sourceType,required"`
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

func (r AdHocOperationalDataRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdHocOperationalDataRequestParam) implementsDataExportNewAdhocParamsBodyUnion() {}

type AdHocOperationalDataRequestOperationalDataType string

const (
	AdHocOperationalDataRequestOperationalDataTypeBills                AdHocOperationalDataRequestOperationalDataType = "BILLS"
	AdHocOperationalDataRequestOperationalDataTypeCommitments          AdHocOperationalDataRequestOperationalDataType = "COMMITMENTS"
	AdHocOperationalDataRequestOperationalDataTypeAccounts             AdHocOperationalDataRequestOperationalDataType = "ACCOUNTS"
	AdHocOperationalDataRequestOperationalDataTypeBalances             AdHocOperationalDataRequestOperationalDataType = "BALANCES"
	AdHocOperationalDataRequestOperationalDataTypeContracts            AdHocOperationalDataRequestOperationalDataType = "CONTRACTS"
	AdHocOperationalDataRequestOperationalDataTypeAccountPlans         AdHocOperationalDataRequestOperationalDataType = "ACCOUNT_PLANS"
	AdHocOperationalDataRequestOperationalDataTypeAggregations         AdHocOperationalDataRequestOperationalDataType = "AGGREGATIONS"
	AdHocOperationalDataRequestOperationalDataTypePlans                AdHocOperationalDataRequestOperationalDataType = "PLANS"
	AdHocOperationalDataRequestOperationalDataTypePricing              AdHocOperationalDataRequestOperationalDataType = "PRICING"
	AdHocOperationalDataRequestOperationalDataTypePricingBands         AdHocOperationalDataRequestOperationalDataType = "PRICING_BANDS"
	AdHocOperationalDataRequestOperationalDataTypeBillLineItems        AdHocOperationalDataRequestOperationalDataType = "BILL_LINE_ITEMS"
	AdHocOperationalDataRequestOperationalDataTypeMeters               AdHocOperationalDataRequestOperationalDataType = "METERS"
	AdHocOperationalDataRequestOperationalDataTypeProducts             AdHocOperationalDataRequestOperationalDataType = "PRODUCTS"
	AdHocOperationalDataRequestOperationalDataTypeCompoundAggregations AdHocOperationalDataRequestOperationalDataType = "COMPOUND_AGGREGATIONS"
	AdHocOperationalDataRequestOperationalDataTypePlanGroups           AdHocOperationalDataRequestOperationalDataType = "PLAN_GROUPS"
	AdHocOperationalDataRequestOperationalDataTypePlanGroupLinks       AdHocOperationalDataRequestOperationalDataType = "PLAN_GROUP_LINKS"
	AdHocOperationalDataRequestOperationalDataTypePlanTemplates        AdHocOperationalDataRequestOperationalDataType = "PLAN_TEMPLATES"
	AdHocOperationalDataRequestOperationalDataTypeBalanceTransactions  AdHocOperationalDataRequestOperationalDataType = "BALANCE_TRANSACTIONS"
)

func (r AdHocOperationalDataRequestOperationalDataType) IsKnown() bool {
	switch r {
	case AdHocOperationalDataRequestOperationalDataTypeBills, AdHocOperationalDataRequestOperationalDataTypeCommitments, AdHocOperationalDataRequestOperationalDataTypeAccounts, AdHocOperationalDataRequestOperationalDataTypeBalances, AdHocOperationalDataRequestOperationalDataTypeContracts, AdHocOperationalDataRequestOperationalDataTypeAccountPlans, AdHocOperationalDataRequestOperationalDataTypeAggregations, AdHocOperationalDataRequestOperationalDataTypePlans, AdHocOperationalDataRequestOperationalDataTypePricing, AdHocOperationalDataRequestOperationalDataTypePricingBands, AdHocOperationalDataRequestOperationalDataTypeBillLineItems, AdHocOperationalDataRequestOperationalDataTypeMeters, AdHocOperationalDataRequestOperationalDataTypeProducts, AdHocOperationalDataRequestOperationalDataTypeCompoundAggregations, AdHocOperationalDataRequestOperationalDataTypePlanGroups, AdHocOperationalDataRequestOperationalDataTypePlanGroupLinks, AdHocOperationalDataRequestOperationalDataTypePlanTemplates, AdHocOperationalDataRequestOperationalDataTypeBalanceTransactions:
		return true
	}
	return false
}

type AdHocOperationalDataRequestSourceType string

const (
	AdHocOperationalDataRequestSourceTypeUsage       AdHocOperationalDataRequestSourceType = "USAGE"
	AdHocOperationalDataRequestSourceTypeOperational AdHocOperationalDataRequestSourceType = "OPERATIONAL"
)

func (r AdHocOperationalDataRequestSourceType) IsKnown() bool {
	switch r {
	case AdHocOperationalDataRequestSourceTypeUsage, AdHocOperationalDataRequestSourceTypeOperational:
		return true
	}
	return false
}

// Response containing data export ad-hoc jobId
type AdHocResponse struct {
	// The id of the job
	JobID string            `json:"jobId"`
	JSON  adHocResponseJSON `json:"-"`
}

// adHocResponseJSON contains the JSON metadata for the struct [AdHocResponse]
type adHocResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdHocResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adHocResponseJSON) RawJSON() string {
	return r.raw
}

type AdHocUsageDataRequestParam struct {
	// Specifies the time period for the aggregation of usage data included each time
	// the Data Export Schedule runs:
	//
	//   - **ORIGINAL**. Usage data is _not aggregated_. If you select to not aggregate,
	//     then raw usage data measurements collected by all Data Field types and any
	//     Derived Fields on the selected Meters are included in the export. This is the
	//     _Default_.
	//
	// If you want to aggregate usage data for the Export Schedule you must define an
	// `aggregationFrequency`:
	//
	// - **HOUR**. Aggregated hourly.
	// - **DAY**. Aggregated daily.
	// - **WEEK**. Aggregated weekly.
	// - **MONTH**. Aggregated monthly.
	//
	//   - If you select to aggregate usage data for a Export Schedule, then only the
	//     aggregated usage data collected by numeric Data Fields of type **MEASURE**,
	//     **INCOME**, or **COST** on selected Meters are included in the export.
	//
	// **NOTE**: If you define an `aggregationFrequency` other than **ORIGINAL** and do
	// not define an `aggregation` method, then you'll receive and error.
	AggregationFrequency param.Field[AdHocUsageDataRequestAggregationFrequency] `json:"aggregationFrequency,required"`
	SourceType           param.Field[AdHocUsageDataRequestSourceType]           `json:"sourceType,required"`
	// List of account IDs for which the usage data will be exported.
	AccountIDs param.Field[[]string] `json:"accountIds"`
	// Specifies the aggregation method applied to usage data collected in the numeric
	// Data Fields of Meters included for the Data Export Schedule - that is, Data
	// Fields of type **MEASURE**, **INCOME**, or **COST**:
	//
	//   - **SUM**. Adds the values.
	//   - **MIN**. Uses the minimum value.
	//   - **MAX**. Uses the maximum value.
	//   - **COUNT**. Counts the number of values.
	//   - **LATEST**. Uses the most recent value. Note: Based on the timestamp `ts`
	//     value of usage data measurement submissions. If using this method, please
	//     ensure _distinct_ `ts` values are used for usage data measurement submissions.
	Aggregation param.Field[AdHocUsageDataRequestAggregation] `json:"aggregation"`
	// List of meter IDs for which the usage data will be exported.
	MeterIDs param.Field[[]string] `json:"meterIds"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export
	//     runs.
	//   - **YESTERDAY**. Data collected for the day before the export runs - that is,
	//     the 24 hour period from midnight to midnight of the day before.
	//   - **WEEK_TO_DATE**. Data collected for the period covering the current week to
	//     the date and time the export runs, and weeks run Monday to Monday.
	//   - **CURRENT_MONTH**. Data collected for the current month in which the export is
	//     ran up to and including the date and time the export runs.
	//   - **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export
	//     is ran.
	//   - **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export
	//     is ran.
	//   - **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks
	//     run Monday to Monday.
	//   - **PREVIOUS_MONTH**. Data collected for the previous full month period.
	//
	// For more details and examples, see the
	// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
	// section in our main User Documentation.
	TimePeriod param.Field[AdHocUsageDataRequestTimePeriod] `json:"timePeriod"`
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

func (r AdHocUsageDataRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdHocUsageDataRequestParam) implementsDataExportNewAdhocParamsBodyUnion() {}

// Specifies the time period for the aggregation of usage data included each time
// the Data Export Schedule runs:
//
//   - **ORIGINAL**. Usage data is _not aggregated_. If you select to not aggregate,
//     then raw usage data measurements collected by all Data Field types and any
//     Derived Fields on the selected Meters are included in the export. This is the
//     _Default_.
//
// If you want to aggregate usage data for the Export Schedule you must define an
// `aggregationFrequency`:
//
// - **HOUR**. Aggregated hourly.
// - **DAY**. Aggregated daily.
// - **WEEK**. Aggregated weekly.
// - **MONTH**. Aggregated monthly.
//
//   - If you select to aggregate usage data for a Export Schedule, then only the
//     aggregated usage data collected by numeric Data Fields of type **MEASURE**,
//     **INCOME**, or **COST** on selected Meters are included in the export.
//
// **NOTE**: If you define an `aggregationFrequency` other than **ORIGINAL** and do
// not define an `aggregation` method, then you'll receive and error.
type AdHocUsageDataRequestAggregationFrequency string

const (
	AdHocUsageDataRequestAggregationFrequencyOriginal AdHocUsageDataRequestAggregationFrequency = "ORIGINAL"
	AdHocUsageDataRequestAggregationFrequencyHour     AdHocUsageDataRequestAggregationFrequency = "HOUR"
	AdHocUsageDataRequestAggregationFrequencyDay      AdHocUsageDataRequestAggregationFrequency = "DAY"
	AdHocUsageDataRequestAggregationFrequencyWeek     AdHocUsageDataRequestAggregationFrequency = "WEEK"
	AdHocUsageDataRequestAggregationFrequencyMonth    AdHocUsageDataRequestAggregationFrequency = "MONTH"
)

func (r AdHocUsageDataRequestAggregationFrequency) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestAggregationFrequencyOriginal, AdHocUsageDataRequestAggregationFrequencyHour, AdHocUsageDataRequestAggregationFrequencyDay, AdHocUsageDataRequestAggregationFrequencyWeek, AdHocUsageDataRequestAggregationFrequencyMonth:
		return true
	}
	return false
}

type AdHocUsageDataRequestSourceType string

const (
	AdHocUsageDataRequestSourceTypeUsage       AdHocUsageDataRequestSourceType = "USAGE"
	AdHocUsageDataRequestSourceTypeOperational AdHocUsageDataRequestSourceType = "OPERATIONAL"
)

func (r AdHocUsageDataRequestSourceType) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestSourceTypeUsage, AdHocUsageDataRequestSourceTypeOperational:
		return true
	}
	return false
}

// Specifies the aggregation method applied to usage data collected in the numeric
// Data Fields of Meters included for the Data Export Schedule - that is, Data
// Fields of type **MEASURE**, **INCOME**, or **COST**:
//
//   - **SUM**. Adds the values.
//   - **MIN**. Uses the minimum value.
//   - **MAX**. Uses the maximum value.
//   - **COUNT**. Counts the number of values.
//   - **LATEST**. Uses the most recent value. Note: Based on the timestamp `ts`
//     value of usage data measurement submissions. If using this method, please
//     ensure _distinct_ `ts` values are used for usage data measurement submissions.
type AdHocUsageDataRequestAggregation string

const (
	AdHocUsageDataRequestAggregationSum    AdHocUsageDataRequestAggregation = "SUM"
	AdHocUsageDataRequestAggregationMin    AdHocUsageDataRequestAggregation = "MIN"
	AdHocUsageDataRequestAggregationMax    AdHocUsageDataRequestAggregation = "MAX"
	AdHocUsageDataRequestAggregationCount  AdHocUsageDataRequestAggregation = "COUNT"
	AdHocUsageDataRequestAggregationLatest AdHocUsageDataRequestAggregation = "LATEST"
	AdHocUsageDataRequestAggregationMean   AdHocUsageDataRequestAggregation = "MEAN"
)

func (r AdHocUsageDataRequestAggregation) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestAggregationSum, AdHocUsageDataRequestAggregationMin, AdHocUsageDataRequestAggregationMax, AdHocUsageDataRequestAggregationCount, AdHocUsageDataRequestAggregationLatest, AdHocUsageDataRequestAggregationMean:
		return true
	}
	return false
}

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export
//     runs.
//   - **YESTERDAY**. Data collected for the day before the export runs - that is,
//     the 24 hour period from midnight to midnight of the day before.
//   - **WEEK_TO_DATE**. Data collected for the period covering the current week to
//     the date and time the export runs, and weeks run Monday to Monday.
//   - **CURRENT_MONTH**. Data collected for the current month in which the export is
//     ran up to and including the date and time the export runs.
//   - **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export
//     is ran.
//   - **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export
//     is ran.
//   - **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks
//     run Monday to Monday.
//   - **PREVIOUS_MONTH**. Data collected for the previous full month period.
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type AdHocUsageDataRequestTimePeriod string

const (
	AdHocUsageDataRequestTimePeriodToday         AdHocUsageDataRequestTimePeriod = "TODAY"
	AdHocUsageDataRequestTimePeriodYesterday     AdHocUsageDataRequestTimePeriod = "YESTERDAY"
	AdHocUsageDataRequestTimePeriodWeekToDate    AdHocUsageDataRequestTimePeriod = "WEEK_TO_DATE"
	AdHocUsageDataRequestTimePeriodCurrentMonth  AdHocUsageDataRequestTimePeriod = "CURRENT_MONTH"
	AdHocUsageDataRequestTimePeriodLast30Days    AdHocUsageDataRequestTimePeriod = "LAST_30_DAYS"
	AdHocUsageDataRequestTimePeriodLast35Days    AdHocUsageDataRequestTimePeriod = "LAST_35_DAYS"
	AdHocUsageDataRequestTimePeriodPreviousWeek  AdHocUsageDataRequestTimePeriod = "PREVIOUS_WEEK"
	AdHocUsageDataRequestTimePeriodPreviousMonth AdHocUsageDataRequestTimePeriod = "PREVIOUS_MONTH"
)

func (r AdHocUsageDataRequestTimePeriod) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestTimePeriodToday, AdHocUsageDataRequestTimePeriodYesterday, AdHocUsageDataRequestTimePeriodWeekToDate, AdHocUsageDataRequestTimePeriodCurrentMonth, AdHocUsageDataRequestTimePeriodLast30Days, AdHocUsageDataRequestTimePeriodLast35Days, AdHocUsageDataRequestTimePeriodPreviousWeek, AdHocUsageDataRequestTimePeriodPreviousMonth:
		return true
	}
	return false
}

type DataExportNewAdhocParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Request representing an operational data export configuration.
	Body DataExportNewAdhocParamsBodyUnion `json:"body,required"`
}

func (r DataExportNewAdhocParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request representing an operational data export configuration.
type DataExportNewAdhocParamsBody struct {
	SourceType param.Field[DataExportNewAdhocParamsBodySourceType] `json:"sourceType,required"`
	AccountIDs param.Field[interface{}]                            `json:"accountIds"`
	// Specifies the aggregation method applied to usage data collected in the numeric
	// Data Fields of Meters included for the Data Export Schedule - that is, Data
	// Fields of type **MEASURE**, **INCOME**, or **COST**:
	//
	//   - **SUM**. Adds the values.
	//   - **MIN**. Uses the minimum value.
	//   - **MAX**. Uses the maximum value.
	//   - **COUNT**. Counts the number of values.
	//   - **LATEST**. Uses the most recent value. Note: Based on the timestamp `ts`
	//     value of usage data measurement submissions. If using this method, please
	//     ensure _distinct_ `ts` values are used for usage data measurement submissions.
	Aggregation param.Field[DataExportNewAdhocParamsBodyAggregation] `json:"aggregation"`
	// Specifies the time period for the aggregation of usage data included each time
	// the Data Export Schedule runs:
	//
	//   - **ORIGINAL**. Usage data is _not aggregated_. If you select to not aggregate,
	//     then raw usage data measurements collected by all Data Field types and any
	//     Derived Fields on the selected Meters are included in the export. This is the
	//     _Default_.
	//
	// If you want to aggregate usage data for the Export Schedule you must define an
	// `aggregationFrequency`:
	//
	// - **HOUR**. Aggregated hourly.
	// - **DAY**. Aggregated daily.
	// - **WEEK**. Aggregated weekly.
	// - **MONTH**. Aggregated monthly.
	//
	//   - If you select to aggregate usage data for a Export Schedule, then only the
	//     aggregated usage data collected by numeric Data Fields of type **MEASURE**,
	//     **INCOME**, or **COST** on selected Meters are included in the export.
	//
	// **NOTE**: If you define an `aggregationFrequency` other than **ORIGINAL** and do
	// not define an `aggregation` method, then you'll receive and error.
	AggregationFrequency param.Field[DataExportNewAdhocParamsBodyAggregationFrequency] `json:"aggregationFrequency"`
	MeterIDs             param.Field[interface{}]                                      `json:"meterIds"`
	OperationalDataTypes param.Field[interface{}]                                      `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export
	//     runs.
	//   - **YESTERDAY**. Data collected for the day before the export runs - that is,
	//     the 24 hour period from midnight to midnight of the day before.
	//   - **WEEK_TO_DATE**. Data collected for the period covering the current week to
	//     the date and time the export runs, and weeks run Monday to Monday.
	//   - **CURRENT_MONTH**. Data collected for the current month in which the export is
	//     ran up to and including the date and time the export runs.
	//   - **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export
	//     is ran.
	//   - **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export
	//     is ran.
	//   - **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks
	//     run Monday to Monday.
	//   - **PREVIOUS_MONTH**. Data collected for the previous full month period.
	//
	// For more details and examples, see the
	// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
	// section in our main User Documentation.
	TimePeriod param.Field[DataExportNewAdhocParamsBodyTimePeriod] `json:"timePeriod"`
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

func (r DataExportNewAdhocParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportNewAdhocParamsBody) implementsDataExportNewAdhocParamsBodyUnion() {}

// Request representing an operational data export configuration.
//
// Satisfied by [AdHocOperationalDataRequestParam], [AdHocUsageDataRequestParam],
// [DataExportNewAdhocParamsBody].
type DataExportNewAdhocParamsBodyUnion interface {
	implementsDataExportNewAdhocParamsBodyUnion()
}

type DataExportNewAdhocParamsBodySourceType string

const (
	DataExportNewAdhocParamsBodySourceTypeUsage       DataExportNewAdhocParamsBodySourceType = "USAGE"
	DataExportNewAdhocParamsBodySourceTypeOperational DataExportNewAdhocParamsBodySourceType = "OPERATIONAL"
)

func (r DataExportNewAdhocParamsBodySourceType) IsKnown() bool {
	switch r {
	case DataExportNewAdhocParamsBodySourceTypeUsage, DataExportNewAdhocParamsBodySourceTypeOperational:
		return true
	}
	return false
}

// Specifies the aggregation method applied to usage data collected in the numeric
// Data Fields of Meters included for the Data Export Schedule - that is, Data
// Fields of type **MEASURE**, **INCOME**, or **COST**:
//
//   - **SUM**. Adds the values.
//   - **MIN**. Uses the minimum value.
//   - **MAX**. Uses the maximum value.
//   - **COUNT**. Counts the number of values.
//   - **LATEST**. Uses the most recent value. Note: Based on the timestamp `ts`
//     value of usage data measurement submissions. If using this method, please
//     ensure _distinct_ `ts` values are used for usage data measurement submissions.
type DataExportNewAdhocParamsBodyAggregation string

const (
	DataExportNewAdhocParamsBodyAggregationSum    DataExportNewAdhocParamsBodyAggregation = "SUM"
	DataExportNewAdhocParamsBodyAggregationMin    DataExportNewAdhocParamsBodyAggregation = "MIN"
	DataExportNewAdhocParamsBodyAggregationMax    DataExportNewAdhocParamsBodyAggregation = "MAX"
	DataExportNewAdhocParamsBodyAggregationCount  DataExportNewAdhocParamsBodyAggregation = "COUNT"
	DataExportNewAdhocParamsBodyAggregationLatest DataExportNewAdhocParamsBodyAggregation = "LATEST"
	DataExportNewAdhocParamsBodyAggregationMean   DataExportNewAdhocParamsBodyAggregation = "MEAN"
)

func (r DataExportNewAdhocParamsBodyAggregation) IsKnown() bool {
	switch r {
	case DataExportNewAdhocParamsBodyAggregationSum, DataExportNewAdhocParamsBodyAggregationMin, DataExportNewAdhocParamsBodyAggregationMax, DataExportNewAdhocParamsBodyAggregationCount, DataExportNewAdhocParamsBodyAggregationLatest, DataExportNewAdhocParamsBodyAggregationMean:
		return true
	}
	return false
}

// Specifies the time period for the aggregation of usage data included each time
// the Data Export Schedule runs:
//
//   - **ORIGINAL**. Usage data is _not aggregated_. If you select to not aggregate,
//     then raw usage data measurements collected by all Data Field types and any
//     Derived Fields on the selected Meters are included in the export. This is the
//     _Default_.
//
// If you want to aggregate usage data for the Export Schedule you must define an
// `aggregationFrequency`:
//
// - **HOUR**. Aggregated hourly.
// - **DAY**. Aggregated daily.
// - **WEEK**. Aggregated weekly.
// - **MONTH**. Aggregated monthly.
//
//   - If you select to aggregate usage data for a Export Schedule, then only the
//     aggregated usage data collected by numeric Data Fields of type **MEASURE**,
//     **INCOME**, or **COST** on selected Meters are included in the export.
//
// **NOTE**: If you define an `aggregationFrequency` other than **ORIGINAL** and do
// not define an `aggregation` method, then you'll receive and error.
type DataExportNewAdhocParamsBodyAggregationFrequency string

const (
	DataExportNewAdhocParamsBodyAggregationFrequencyOriginal DataExportNewAdhocParamsBodyAggregationFrequency = "ORIGINAL"
	DataExportNewAdhocParamsBodyAggregationFrequencyHour     DataExportNewAdhocParamsBodyAggregationFrequency = "HOUR"
	DataExportNewAdhocParamsBodyAggregationFrequencyDay      DataExportNewAdhocParamsBodyAggregationFrequency = "DAY"
	DataExportNewAdhocParamsBodyAggregationFrequencyWeek     DataExportNewAdhocParamsBodyAggregationFrequency = "WEEK"
	DataExportNewAdhocParamsBodyAggregationFrequencyMonth    DataExportNewAdhocParamsBodyAggregationFrequency = "MONTH"
)

func (r DataExportNewAdhocParamsBodyAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportNewAdhocParamsBodyAggregationFrequencyOriginal, DataExportNewAdhocParamsBodyAggregationFrequencyHour, DataExportNewAdhocParamsBodyAggregationFrequencyDay, DataExportNewAdhocParamsBodyAggregationFrequencyWeek, DataExportNewAdhocParamsBodyAggregationFrequencyMonth:
		return true
	}
	return false
}

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export
//     runs.
//   - **YESTERDAY**. Data collected for the day before the export runs - that is,
//     the 24 hour period from midnight to midnight of the day before.
//   - **WEEK_TO_DATE**. Data collected for the period covering the current week to
//     the date and time the export runs, and weeks run Monday to Monday.
//   - **CURRENT_MONTH**. Data collected for the current month in which the export is
//     ran up to and including the date and time the export runs.
//   - **LAST_30_DAYS**. Data collected for the 30 days prior to the date the export
//     is ran.
//   - **LAST_35_DAYS**. Data collected for the 35 days prior to the date the export
//     is ran.
//   - **PREVIOUS_WEEK**. Data collected for the previous full week period, and weeks
//     run Monday to Monday.
//   - **PREVIOUS_MONTH**. Data collected for the previous full month period.
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportNewAdhocParamsBodyTimePeriod string

const (
	DataExportNewAdhocParamsBodyTimePeriodToday         DataExportNewAdhocParamsBodyTimePeriod = "TODAY"
	DataExportNewAdhocParamsBodyTimePeriodYesterday     DataExportNewAdhocParamsBodyTimePeriod = "YESTERDAY"
	DataExportNewAdhocParamsBodyTimePeriodWeekToDate    DataExportNewAdhocParamsBodyTimePeriod = "WEEK_TO_DATE"
	DataExportNewAdhocParamsBodyTimePeriodCurrentMonth  DataExportNewAdhocParamsBodyTimePeriod = "CURRENT_MONTH"
	DataExportNewAdhocParamsBodyTimePeriodLast30Days    DataExportNewAdhocParamsBodyTimePeriod = "LAST_30_DAYS"
	DataExportNewAdhocParamsBodyTimePeriodLast35Days    DataExportNewAdhocParamsBodyTimePeriod = "LAST_35_DAYS"
	DataExportNewAdhocParamsBodyTimePeriodPreviousWeek  DataExportNewAdhocParamsBodyTimePeriod = "PREVIOUS_WEEK"
	DataExportNewAdhocParamsBodyTimePeriodPreviousMonth DataExportNewAdhocParamsBodyTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportNewAdhocParamsBodyTimePeriod) IsKnown() bool {
	switch r {
	case DataExportNewAdhocParamsBodyTimePeriodToday, DataExportNewAdhocParamsBodyTimePeriodYesterday, DataExportNewAdhocParamsBodyTimePeriodWeekToDate, DataExportNewAdhocParamsBodyTimePeriodCurrentMonth, DataExportNewAdhocParamsBodyTimePeriodLast30Days, DataExportNewAdhocParamsBodyTimePeriodLast35Days, DataExportNewAdhocParamsBodyTimePeriodPreviousWeek, DataExportNewAdhocParamsBodyTimePeriodPreviousMonth:
		return true
	}
	return false
}
