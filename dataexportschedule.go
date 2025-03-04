// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// DataExportScheduleService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataExportScheduleService] method instead.
type DataExportScheduleService struct {
	Options []option.RequestOption
}

// NewDataExportScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataExportScheduleService(opts ...option.RequestOption) (r *DataExportScheduleService) {
	r = &DataExportScheduleService{}
	r.Options = opts
	return
}

// Create a new Data Export Schedule. Each Schedule can be configured for exporting
// _only one_ of either Usage or Operational data:
//
// **Operational Data Exports**.
//
//   - Use the `operationalDataTypes` parameter to specify the entities whose
//     operational data you want to include in the export each time the Export
//     Schedule runs.
//   - For each of the entity types you select, each time the Export Schedule runs a
//     separate file is compiled containing the operational data for all entities of
//     that type that exist in your Organization
//
// **Usage Data Exports**.
//
//   - Select the Meters and Accounts whose usage data you want to include in the
//     export each time the Export Schedule runs.
//   - If _don't want to aggregate_ the usage data collected by the selected Meters,
//     use **ORIGINAL** for `aggregationFrequency`, which is the _default_. This
//     means the raw usage data collected by any type of Data Fields and the values
//     for any Derived Fields on the selected Meters will be included in the export.
//   - If you _do want to aggregate_ the usage data collected by the selected Meters,
//     use one of the other options for `aggregationFrequency`: **HOUR**, **DAY**,
//     **WEEK**, or **MONTH**. You _must_ then also specified an `aggregation` method
//     to be used on the usage data before export. Importantly, if you do aggregate
//     usage data, only the usage data collected by any numeric Data Fields on the
//     selected Meters - those of type **MEASURE**, **INCOME**, or **COST** - will be
//     included in the export each time the Export Schedule runs.
func (r *DataExportScheduleService) New(ctx context.Context, params DataExportScheduleNewParams, opts ...option.RequestOption) (res *DataExportScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/schedules", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Data Export Schedule for the given UUID. Each Schedule can be
// configured for exporting _only one_ of either Usage or Operational data.
func (r *DataExportScheduleService) Get(ctx context.Context, id string, query DataExportScheduleGetParams, opts ...option.RequestOption) (res *DataExportScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/schedules/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Data Export Schedule for the given UUID. Each Schedule can be
// configured for exporting _only one_ of either Usage or Operational data:
//
// **Operational Data Exports**.
//
//   - Use the `operationalDataTypes` parameter to specify the entities whose
//     operational data you want to include in the export each time the Export
//     Schedule runs.
//   - For each of the entity types you select, each time the Export Schedule runs a
//     separate file is compiled containing the operational data for all entities of
//     that type that exist in your Organization
//
// **Usage Data Exports**.
//
//   - Select the Meters and Accounts whose usage data you want to include in the
//     export each time the Export Schedule runs.
//   - If _don't want to aggregate_ the usage data collected by the selected Meters,
//     use **ORIGINAL** for `aggregationFrequency`, which is the _default_. This
//     means the raw usage data collected by any type of Data Fields and the values
//     for any Derived Fields on the selected Meters will be included in the export.
//   - If you _do want to aggregate_ the usage data collected by the selected Meters,
//     use one of the other options for `aggregationFrequency`: **HOUR**, **DAY**,
//     **WEEK**, or **MONTH**. You _must_ then also specified an `aggregation` method
//     to be used on the usage data before export. Importantly, if you do aggregate
//     usage data, only the usage data collected by any numeric Data Fields on the
//     selected Meters - those of type **MEASURE**, **INCOME**, or **COST** - will be
//     included in the export each time the Export Schedule runs.
func (r *DataExportScheduleService) Update(ctx context.Context, id string, params DataExportScheduleUpdateParams, opts ...option.RequestOption) (res *DataExportScheduleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/schedules/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Data Export Schedules created for your Organization. You can
// filter the response by Schedules `ids`.
//
// The response will contain an array for both the operational and usage Data
// Export Schedules in your Organization.
func (r *DataExportScheduleService) List(ctx context.Context, params DataExportScheduleListParams, opts ...option.RequestOption) (res *pagination.Cursor[DataExportScheduleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/schedules", params.OrgID)
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

// Retrieve a list of Data Export Schedules created for your Organization. You can
// filter the response by Schedules `ids`.
//
// The response will contain an array for both the operational and usage Data
// Export Schedules in your Organization.
func (r *DataExportScheduleService) ListAutoPaging(ctx context.Context, params DataExportScheduleListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DataExportScheduleListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Data Export Schedule for the given UUID. Each Schedule can be
// configured for exporting _only one_ of either Usage or Operational data.
func (r *DataExportScheduleService) Delete(ctx context.Context, id string, body DataExportScheduleDeleteParams, opts ...option.RequestOption) (res *DataExportScheduleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/schedules/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type OperationalDataExportScheduleRequestParam struct {
	// A list of the entities whose operational data is included in the data export.
	OperationalDataTypes param.Field[[]OperationalDataExportScheduleRequestOperationalDataType] `json:"operationalDataTypes,required"`
	SourceType           param.Field[OperationalDataExportScheduleRequestSourceType]            `json:"sourceType,required"`
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

func (r OperationalDataExportScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OperationalDataExportScheduleRequestParam) implementsDataExportScheduleNewParamsBodyUnion() {}

func (r OperationalDataExportScheduleRequestParam) implementsDataExportScheduleUpdateParamsBodyUnion() {
}

type OperationalDataExportScheduleRequestOperationalDataType string

const (
	OperationalDataExportScheduleRequestOperationalDataTypeBills                OperationalDataExportScheduleRequestOperationalDataType = "BILLS"
	OperationalDataExportScheduleRequestOperationalDataTypeCommitments          OperationalDataExportScheduleRequestOperationalDataType = "COMMITMENTS"
	OperationalDataExportScheduleRequestOperationalDataTypeAccounts             OperationalDataExportScheduleRequestOperationalDataType = "ACCOUNTS"
	OperationalDataExportScheduleRequestOperationalDataTypeBalances             OperationalDataExportScheduleRequestOperationalDataType = "BALANCES"
	OperationalDataExportScheduleRequestOperationalDataTypeContracts            OperationalDataExportScheduleRequestOperationalDataType = "CONTRACTS"
	OperationalDataExportScheduleRequestOperationalDataTypeAccountPlans         OperationalDataExportScheduleRequestOperationalDataType = "ACCOUNT_PLANS"
	OperationalDataExportScheduleRequestOperationalDataTypeAggregations         OperationalDataExportScheduleRequestOperationalDataType = "AGGREGATIONS"
	OperationalDataExportScheduleRequestOperationalDataTypePlans                OperationalDataExportScheduleRequestOperationalDataType = "PLANS"
	OperationalDataExportScheduleRequestOperationalDataTypePricing              OperationalDataExportScheduleRequestOperationalDataType = "PRICING"
	OperationalDataExportScheduleRequestOperationalDataTypePricingBands         OperationalDataExportScheduleRequestOperationalDataType = "PRICING_BANDS"
	OperationalDataExportScheduleRequestOperationalDataTypeBillLineItems        OperationalDataExportScheduleRequestOperationalDataType = "BILL_LINE_ITEMS"
	OperationalDataExportScheduleRequestOperationalDataTypeMeters               OperationalDataExportScheduleRequestOperationalDataType = "METERS"
	OperationalDataExportScheduleRequestOperationalDataTypeProducts             OperationalDataExportScheduleRequestOperationalDataType = "PRODUCTS"
	OperationalDataExportScheduleRequestOperationalDataTypeCompoundAggregations OperationalDataExportScheduleRequestOperationalDataType = "COMPOUND_AGGREGATIONS"
	OperationalDataExportScheduleRequestOperationalDataTypePlanGroups           OperationalDataExportScheduleRequestOperationalDataType = "PLAN_GROUPS"
	OperationalDataExportScheduleRequestOperationalDataTypePlanGroupLinks       OperationalDataExportScheduleRequestOperationalDataType = "PLAN_GROUP_LINKS"
	OperationalDataExportScheduleRequestOperationalDataTypePlanTemplates        OperationalDataExportScheduleRequestOperationalDataType = "PLAN_TEMPLATES"
	OperationalDataExportScheduleRequestOperationalDataTypeBalanceTransactions  OperationalDataExportScheduleRequestOperationalDataType = "BALANCE_TRANSACTIONS"
)

func (r OperationalDataExportScheduleRequestOperationalDataType) IsKnown() bool {
	switch r {
	case OperationalDataExportScheduleRequestOperationalDataTypeBills, OperationalDataExportScheduleRequestOperationalDataTypeCommitments, OperationalDataExportScheduleRequestOperationalDataTypeAccounts, OperationalDataExportScheduleRequestOperationalDataTypeBalances, OperationalDataExportScheduleRequestOperationalDataTypeContracts, OperationalDataExportScheduleRequestOperationalDataTypeAccountPlans, OperationalDataExportScheduleRequestOperationalDataTypeAggregations, OperationalDataExportScheduleRequestOperationalDataTypePlans, OperationalDataExportScheduleRequestOperationalDataTypePricing, OperationalDataExportScheduleRequestOperationalDataTypePricingBands, OperationalDataExportScheduleRequestOperationalDataTypeBillLineItems, OperationalDataExportScheduleRequestOperationalDataTypeMeters, OperationalDataExportScheduleRequestOperationalDataTypeProducts, OperationalDataExportScheduleRequestOperationalDataTypeCompoundAggregations, OperationalDataExportScheduleRequestOperationalDataTypePlanGroups, OperationalDataExportScheduleRequestOperationalDataTypePlanGroupLinks, OperationalDataExportScheduleRequestOperationalDataTypePlanTemplates, OperationalDataExportScheduleRequestOperationalDataTypeBalanceTransactions:
		return true
	}
	return false
}

type OperationalDataExportScheduleRequestSourceType string

const (
	OperationalDataExportScheduleRequestSourceTypeUsage       OperationalDataExportScheduleRequestSourceType = "USAGE"
	OperationalDataExportScheduleRequestSourceTypeOperational OperationalDataExportScheduleRequestSourceType = "OPERATIONAL"
)

func (r OperationalDataExportScheduleRequestSourceType) IsKnown() bool {
	switch r {
	case OperationalDataExportScheduleRequestSourceTypeUsage, OperationalDataExportScheduleRequestSourceTypeOperational:
		return true
	}
	return false
}

type OperationalDataExportScheduleResponse struct {
	// The id of the schedule.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A list of the entities whose operational data is included in the data export.
	OperationalDataTypes []OperationalDataExportScheduleResponseOperationalDataType `json:"operationalDataTypes"`
	JSON                 operationalDataExportScheduleResponseJSON                  `json:"-"`
}

// operationalDataExportScheduleResponseJSON contains the JSON metadata for the
// struct [OperationalDataExportScheduleResponse]
type operationalDataExportScheduleResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	OperationalDataTypes apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *OperationalDataExportScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationalDataExportScheduleResponseJSON) RawJSON() string {
	return r.raw
}

func (r OperationalDataExportScheduleResponse) implementsDataExportScheduleNewResponse() {}

func (r OperationalDataExportScheduleResponse) implementsDataExportScheduleGetResponse() {}

func (r OperationalDataExportScheduleResponse) implementsDataExportScheduleUpdateResponse() {}

func (r OperationalDataExportScheduleResponse) implementsDataExportScheduleDeleteResponse() {}

type OperationalDataExportScheduleResponseOperationalDataType string

const (
	OperationalDataExportScheduleResponseOperationalDataTypeBills                OperationalDataExportScheduleResponseOperationalDataType = "BILLS"
	OperationalDataExportScheduleResponseOperationalDataTypeCommitments          OperationalDataExportScheduleResponseOperationalDataType = "COMMITMENTS"
	OperationalDataExportScheduleResponseOperationalDataTypeAccounts             OperationalDataExportScheduleResponseOperationalDataType = "ACCOUNTS"
	OperationalDataExportScheduleResponseOperationalDataTypeBalances             OperationalDataExportScheduleResponseOperationalDataType = "BALANCES"
	OperationalDataExportScheduleResponseOperationalDataTypeContracts            OperationalDataExportScheduleResponseOperationalDataType = "CONTRACTS"
	OperationalDataExportScheduleResponseOperationalDataTypeAccountPlans         OperationalDataExportScheduleResponseOperationalDataType = "ACCOUNT_PLANS"
	OperationalDataExportScheduleResponseOperationalDataTypeAggregations         OperationalDataExportScheduleResponseOperationalDataType = "AGGREGATIONS"
	OperationalDataExportScheduleResponseOperationalDataTypePlans                OperationalDataExportScheduleResponseOperationalDataType = "PLANS"
	OperationalDataExportScheduleResponseOperationalDataTypePricing              OperationalDataExportScheduleResponseOperationalDataType = "PRICING"
	OperationalDataExportScheduleResponseOperationalDataTypePricingBands         OperationalDataExportScheduleResponseOperationalDataType = "PRICING_BANDS"
	OperationalDataExportScheduleResponseOperationalDataTypeBillLineItems        OperationalDataExportScheduleResponseOperationalDataType = "BILL_LINE_ITEMS"
	OperationalDataExportScheduleResponseOperationalDataTypeMeters               OperationalDataExportScheduleResponseOperationalDataType = "METERS"
	OperationalDataExportScheduleResponseOperationalDataTypeProducts             OperationalDataExportScheduleResponseOperationalDataType = "PRODUCTS"
	OperationalDataExportScheduleResponseOperationalDataTypeCompoundAggregations OperationalDataExportScheduleResponseOperationalDataType = "COMPOUND_AGGREGATIONS"
	OperationalDataExportScheduleResponseOperationalDataTypePlanGroups           OperationalDataExportScheduleResponseOperationalDataType = "PLAN_GROUPS"
	OperationalDataExportScheduleResponseOperationalDataTypePlanGroupLinks       OperationalDataExportScheduleResponseOperationalDataType = "PLAN_GROUP_LINKS"
	OperationalDataExportScheduleResponseOperationalDataTypePlanTemplates        OperationalDataExportScheduleResponseOperationalDataType = "PLAN_TEMPLATES"
	OperationalDataExportScheduleResponseOperationalDataTypeBalanceTransactions  OperationalDataExportScheduleResponseOperationalDataType = "BALANCE_TRANSACTIONS"
)

func (r OperationalDataExportScheduleResponseOperationalDataType) IsKnown() bool {
	switch r {
	case OperationalDataExportScheduleResponseOperationalDataTypeBills, OperationalDataExportScheduleResponseOperationalDataTypeCommitments, OperationalDataExportScheduleResponseOperationalDataTypeAccounts, OperationalDataExportScheduleResponseOperationalDataTypeBalances, OperationalDataExportScheduleResponseOperationalDataTypeContracts, OperationalDataExportScheduleResponseOperationalDataTypeAccountPlans, OperationalDataExportScheduleResponseOperationalDataTypeAggregations, OperationalDataExportScheduleResponseOperationalDataTypePlans, OperationalDataExportScheduleResponseOperationalDataTypePricing, OperationalDataExportScheduleResponseOperationalDataTypePricingBands, OperationalDataExportScheduleResponseOperationalDataTypeBillLineItems, OperationalDataExportScheduleResponseOperationalDataTypeMeters, OperationalDataExportScheduleResponseOperationalDataTypeProducts, OperationalDataExportScheduleResponseOperationalDataTypeCompoundAggregations, OperationalDataExportScheduleResponseOperationalDataTypePlanGroups, OperationalDataExportScheduleResponseOperationalDataTypePlanGroupLinks, OperationalDataExportScheduleResponseOperationalDataTypePlanTemplates, OperationalDataExportScheduleResponseOperationalDataTypeBalanceTransactions:
		return true
	}
	return false
}

type UsageDataExportScheduleRequestParam struct {
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
	AggregationFrequency param.Field[UsageDataExportScheduleRequestAggregationFrequency] `json:"aggregationFrequency,required"`
	SourceType           param.Field[UsageDataExportScheduleRequestSourceType]           `json:"sourceType,required"`
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
	TimePeriod param.Field[UsageDataExportScheduleRequestTimePeriod] `json:"timePeriod,required"`
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
	Aggregation param.Field[UsageDataExportScheduleRequestAggregation] `json:"aggregation"`
	// List of meter IDs for which the usage data will be exported.
	MeterIDs param.Field[[]string] `json:"meterIds"`
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

func (r UsageDataExportScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDataExportScheduleRequestParam) implementsDataExportScheduleNewParamsBodyUnion() {}

func (r UsageDataExportScheduleRequestParam) implementsDataExportScheduleUpdateParamsBodyUnion() {}

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
type UsageDataExportScheduleRequestAggregationFrequency string

const (
	UsageDataExportScheduleRequestAggregationFrequencyOriginal UsageDataExportScheduleRequestAggregationFrequency = "ORIGINAL"
	UsageDataExportScheduleRequestAggregationFrequencyHour     UsageDataExportScheduleRequestAggregationFrequency = "HOUR"
	UsageDataExportScheduleRequestAggregationFrequencyDay      UsageDataExportScheduleRequestAggregationFrequency = "DAY"
	UsageDataExportScheduleRequestAggregationFrequencyWeek     UsageDataExportScheduleRequestAggregationFrequency = "WEEK"
	UsageDataExportScheduleRequestAggregationFrequencyMonth    UsageDataExportScheduleRequestAggregationFrequency = "MONTH"
)

func (r UsageDataExportScheduleRequestAggregationFrequency) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestAggregationFrequencyOriginal, UsageDataExportScheduleRequestAggregationFrequencyHour, UsageDataExportScheduleRequestAggregationFrequencyDay, UsageDataExportScheduleRequestAggregationFrequencyWeek, UsageDataExportScheduleRequestAggregationFrequencyMonth:
		return true
	}
	return false
}

type UsageDataExportScheduleRequestSourceType string

const (
	UsageDataExportScheduleRequestSourceTypeUsage       UsageDataExportScheduleRequestSourceType = "USAGE"
	UsageDataExportScheduleRequestSourceTypeOperational UsageDataExportScheduleRequestSourceType = "OPERATIONAL"
)

func (r UsageDataExportScheduleRequestSourceType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestSourceTypeUsage, UsageDataExportScheduleRequestSourceTypeOperational:
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
type UsageDataExportScheduleRequestTimePeriod string

const (
	UsageDataExportScheduleRequestTimePeriodToday         UsageDataExportScheduleRequestTimePeriod = "TODAY"
	UsageDataExportScheduleRequestTimePeriodYesterday     UsageDataExportScheduleRequestTimePeriod = "YESTERDAY"
	UsageDataExportScheduleRequestTimePeriodWeekToDate    UsageDataExportScheduleRequestTimePeriod = "WEEK_TO_DATE"
	UsageDataExportScheduleRequestTimePeriodCurrentMonth  UsageDataExportScheduleRequestTimePeriod = "CURRENT_MONTH"
	UsageDataExportScheduleRequestTimePeriodLast30Days    UsageDataExportScheduleRequestTimePeriod = "LAST_30_DAYS"
	UsageDataExportScheduleRequestTimePeriodLast35Days    UsageDataExportScheduleRequestTimePeriod = "LAST_35_DAYS"
	UsageDataExportScheduleRequestTimePeriodPreviousWeek  UsageDataExportScheduleRequestTimePeriod = "PREVIOUS_WEEK"
	UsageDataExportScheduleRequestTimePeriodPreviousMonth UsageDataExportScheduleRequestTimePeriod = "PREVIOUS_MONTH"
)

func (r UsageDataExportScheduleRequestTimePeriod) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestTimePeriodToday, UsageDataExportScheduleRequestTimePeriodYesterday, UsageDataExportScheduleRequestTimePeriodWeekToDate, UsageDataExportScheduleRequestTimePeriodCurrentMonth, UsageDataExportScheduleRequestTimePeriodLast30Days, UsageDataExportScheduleRequestTimePeriodLast35Days, UsageDataExportScheduleRequestTimePeriodPreviousWeek, UsageDataExportScheduleRequestTimePeriodPreviousMonth:
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
type UsageDataExportScheduleRequestAggregation string

const (
	UsageDataExportScheduleRequestAggregationSum    UsageDataExportScheduleRequestAggregation = "SUM"
	UsageDataExportScheduleRequestAggregationMin    UsageDataExportScheduleRequestAggregation = "MIN"
	UsageDataExportScheduleRequestAggregationMax    UsageDataExportScheduleRequestAggregation = "MAX"
	UsageDataExportScheduleRequestAggregationCount  UsageDataExportScheduleRequestAggregation = "COUNT"
	UsageDataExportScheduleRequestAggregationLatest UsageDataExportScheduleRequestAggregation = "LATEST"
	UsageDataExportScheduleRequestAggregationMean   UsageDataExportScheduleRequestAggregation = "MEAN"
)

func (r UsageDataExportScheduleRequestAggregation) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestAggregationSum, UsageDataExportScheduleRequestAggregationMin, UsageDataExportScheduleRequestAggregationMax, UsageDataExportScheduleRequestAggregationCount, UsageDataExportScheduleRequestAggregationLatest, UsageDataExportScheduleRequestAggregationMean:
		return true
	}
	return false
}

type UsageDataExportScheduleResponse struct {
	// The id of the schedule
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// List of account IDs for which the usage data will be exported.
	AccountIDs []string `json:"accountIds"`
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
	Aggregation UsageDataExportScheduleResponseAggregation `json:"aggregation"`
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
	AggregationFrequency UsageDataExportScheduleResponseAggregationFrequency `json:"aggregationFrequency"`
	// List of meter IDs for which the usage data will be exported.
	MeterIDs []string `json:"meterIds"`
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
	TimePeriod UsageDataExportScheduleResponseTimePeriod `json:"timePeriod"`
	JSON       usageDataExportScheduleResponseJSON       `json:"-"`
}

// usageDataExportScheduleResponseJSON contains the JSON metadata for the struct
// [UsageDataExportScheduleResponse]
type usageDataExportScheduleResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	AccountIDs           apijson.Field
	Aggregation          apijson.Field
	AggregationFrequency apijson.Field
	MeterIDs             apijson.Field
	TimePeriod           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UsageDataExportScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDataExportScheduleResponseJSON) RawJSON() string {
	return r.raw
}

func (r UsageDataExportScheduleResponse) implementsDataExportScheduleNewResponse() {}

func (r UsageDataExportScheduleResponse) implementsDataExportScheduleGetResponse() {}

func (r UsageDataExportScheduleResponse) implementsDataExportScheduleUpdateResponse() {}

func (r UsageDataExportScheduleResponse) implementsDataExportScheduleDeleteResponse() {}

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
type UsageDataExportScheduleResponseAggregation string

const (
	UsageDataExportScheduleResponseAggregationSum    UsageDataExportScheduleResponseAggregation = "SUM"
	UsageDataExportScheduleResponseAggregationMin    UsageDataExportScheduleResponseAggregation = "MIN"
	UsageDataExportScheduleResponseAggregationMax    UsageDataExportScheduleResponseAggregation = "MAX"
	UsageDataExportScheduleResponseAggregationCount  UsageDataExportScheduleResponseAggregation = "COUNT"
	UsageDataExportScheduleResponseAggregationLatest UsageDataExportScheduleResponseAggregation = "LATEST"
	UsageDataExportScheduleResponseAggregationMean   UsageDataExportScheduleResponseAggregation = "MEAN"
)

func (r UsageDataExportScheduleResponseAggregation) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseAggregationSum, UsageDataExportScheduleResponseAggregationMin, UsageDataExportScheduleResponseAggregationMax, UsageDataExportScheduleResponseAggregationCount, UsageDataExportScheduleResponseAggregationLatest, UsageDataExportScheduleResponseAggregationMean:
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
type UsageDataExportScheduleResponseAggregationFrequency string

const (
	UsageDataExportScheduleResponseAggregationFrequencyOriginal UsageDataExportScheduleResponseAggregationFrequency = "ORIGINAL"
	UsageDataExportScheduleResponseAggregationFrequencyHour     UsageDataExportScheduleResponseAggregationFrequency = "HOUR"
	UsageDataExportScheduleResponseAggregationFrequencyDay      UsageDataExportScheduleResponseAggregationFrequency = "DAY"
	UsageDataExportScheduleResponseAggregationFrequencyWeek     UsageDataExportScheduleResponseAggregationFrequency = "WEEK"
	UsageDataExportScheduleResponseAggregationFrequencyMonth    UsageDataExportScheduleResponseAggregationFrequency = "MONTH"
)

func (r UsageDataExportScheduleResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseAggregationFrequencyOriginal, UsageDataExportScheduleResponseAggregationFrequencyHour, UsageDataExportScheduleResponseAggregationFrequencyDay, UsageDataExportScheduleResponseAggregationFrequencyWeek, UsageDataExportScheduleResponseAggregationFrequencyMonth:
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
type UsageDataExportScheduleResponseTimePeriod string

const (
	UsageDataExportScheduleResponseTimePeriodToday         UsageDataExportScheduleResponseTimePeriod = "TODAY"
	UsageDataExportScheduleResponseTimePeriodYesterday     UsageDataExportScheduleResponseTimePeriod = "YESTERDAY"
	UsageDataExportScheduleResponseTimePeriodWeekToDate    UsageDataExportScheduleResponseTimePeriod = "WEEK_TO_DATE"
	UsageDataExportScheduleResponseTimePeriodCurrentMonth  UsageDataExportScheduleResponseTimePeriod = "CURRENT_MONTH"
	UsageDataExportScheduleResponseTimePeriodLast30Days    UsageDataExportScheduleResponseTimePeriod = "LAST_30_DAYS"
	UsageDataExportScheduleResponseTimePeriodLast35Days    UsageDataExportScheduleResponseTimePeriod = "LAST_35_DAYS"
	UsageDataExportScheduleResponseTimePeriodPreviousWeek  UsageDataExportScheduleResponseTimePeriod = "PREVIOUS_WEEK"
	UsageDataExportScheduleResponseTimePeriodPreviousMonth UsageDataExportScheduleResponseTimePeriod = "PREVIOUS_MONTH"
)

func (r UsageDataExportScheduleResponseTimePeriod) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseTimePeriodToday, UsageDataExportScheduleResponseTimePeriodYesterday, UsageDataExportScheduleResponseTimePeriodWeekToDate, UsageDataExportScheduleResponseTimePeriodCurrentMonth, UsageDataExportScheduleResponseTimePeriodLast30Days, UsageDataExportScheduleResponseTimePeriodLast35Days, UsageDataExportScheduleResponseTimePeriodPreviousWeek, UsageDataExportScheduleResponseTimePeriodPreviousMonth:
		return true
	}
	return false
}

// Response representing an operational data export configuration.
type DataExportScheduleNewResponse struct {
	// The id of the schedule.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// This field can have the runtime type of [[]string].
	AccountIDs interface{} `json:"accountIds"`
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
	Aggregation DataExportScheduleNewResponseAggregation `json:"aggregation"`
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
	AggregationFrequency DataExportScheduleNewResponseAggregationFrequency `json:"aggregationFrequency"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
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
	TimePeriod DataExportScheduleNewResponseTimePeriod `json:"timePeriod"`
	JSON       dataExportScheduleNewResponseJSON       `json:"-"`
	union      DataExportScheduleNewResponseUnion
}

// dataExportScheduleNewResponseJSON contains the JSON metadata for the struct
// [DataExportScheduleNewResponse]
type dataExportScheduleNewResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	AccountIDs           apijson.Field
	Aggregation          apijson.Field
	AggregationFrequency apijson.Field
	MeterIDs             apijson.Field
	OperationalDataTypes apijson.Field
	TimePeriod           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r dataExportScheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportScheduleNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportScheduleNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [OperationalDataExportScheduleResponse],
// [UsageDataExportScheduleResponse].
func (r DataExportScheduleNewResponse) AsUnion() DataExportScheduleNewResponseUnion {
	return r.union
}

// Response representing an operational data export configuration.
//
// Union satisfied by [OperationalDataExportScheduleResponse] or
// [UsageDataExportScheduleResponse].
type DataExportScheduleNewResponseUnion interface {
	implementsDataExportScheduleNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportScheduleNewResponseUnion)(nil)).Elem(),
		"sourceType",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationalDataExportScheduleResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponse{}),
		},
	)
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
type DataExportScheduleNewResponseAggregation string

const (
	DataExportScheduleNewResponseAggregationSum    DataExportScheduleNewResponseAggregation = "SUM"
	DataExportScheduleNewResponseAggregationMin    DataExportScheduleNewResponseAggregation = "MIN"
	DataExportScheduleNewResponseAggregationMax    DataExportScheduleNewResponseAggregation = "MAX"
	DataExportScheduleNewResponseAggregationCount  DataExportScheduleNewResponseAggregation = "COUNT"
	DataExportScheduleNewResponseAggregationLatest DataExportScheduleNewResponseAggregation = "LATEST"
	DataExportScheduleNewResponseAggregationMean   DataExportScheduleNewResponseAggregation = "MEAN"
)

func (r DataExportScheduleNewResponseAggregation) IsKnown() bool {
	switch r {
	case DataExportScheduleNewResponseAggregationSum, DataExportScheduleNewResponseAggregationMin, DataExportScheduleNewResponseAggregationMax, DataExportScheduleNewResponseAggregationCount, DataExportScheduleNewResponseAggregationLatest, DataExportScheduleNewResponseAggregationMean:
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
type DataExportScheduleNewResponseAggregationFrequency string

const (
	DataExportScheduleNewResponseAggregationFrequencyOriginal DataExportScheduleNewResponseAggregationFrequency = "ORIGINAL"
	DataExportScheduleNewResponseAggregationFrequencyHour     DataExportScheduleNewResponseAggregationFrequency = "HOUR"
	DataExportScheduleNewResponseAggregationFrequencyDay      DataExportScheduleNewResponseAggregationFrequency = "DAY"
	DataExportScheduleNewResponseAggregationFrequencyWeek     DataExportScheduleNewResponseAggregationFrequency = "WEEK"
	DataExportScheduleNewResponseAggregationFrequencyMonth    DataExportScheduleNewResponseAggregationFrequency = "MONTH"
)

func (r DataExportScheduleNewResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportScheduleNewResponseAggregationFrequencyOriginal, DataExportScheduleNewResponseAggregationFrequencyHour, DataExportScheduleNewResponseAggregationFrequencyDay, DataExportScheduleNewResponseAggregationFrequencyWeek, DataExportScheduleNewResponseAggregationFrequencyMonth:
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
type DataExportScheduleNewResponseTimePeriod string

const (
	DataExportScheduleNewResponseTimePeriodToday         DataExportScheduleNewResponseTimePeriod = "TODAY"
	DataExportScheduleNewResponseTimePeriodYesterday     DataExportScheduleNewResponseTimePeriod = "YESTERDAY"
	DataExportScheduleNewResponseTimePeriodWeekToDate    DataExportScheduleNewResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleNewResponseTimePeriodCurrentMonth  DataExportScheduleNewResponseTimePeriod = "CURRENT_MONTH"
	DataExportScheduleNewResponseTimePeriodLast30Days    DataExportScheduleNewResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleNewResponseTimePeriodLast35Days    DataExportScheduleNewResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleNewResponseTimePeriodPreviousWeek  DataExportScheduleNewResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleNewResponseTimePeriodPreviousMonth DataExportScheduleNewResponseTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportScheduleNewResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleNewResponseTimePeriodToday, DataExportScheduleNewResponseTimePeriodYesterday, DataExportScheduleNewResponseTimePeriodWeekToDate, DataExportScheduleNewResponseTimePeriodCurrentMonth, DataExportScheduleNewResponseTimePeriodLast30Days, DataExportScheduleNewResponseTimePeriodLast35Days, DataExportScheduleNewResponseTimePeriodPreviousWeek, DataExportScheduleNewResponseTimePeriodPreviousMonth:
		return true
	}
	return false
}

// Response representing an operational data export configuration.
type DataExportScheduleGetResponse struct {
	// The id of the schedule.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// This field can have the runtime type of [[]string].
	AccountIDs interface{} `json:"accountIds"`
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
	Aggregation DataExportScheduleGetResponseAggregation `json:"aggregation"`
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
	AggregationFrequency DataExportScheduleGetResponseAggregationFrequency `json:"aggregationFrequency"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
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
	TimePeriod DataExportScheduleGetResponseTimePeriod `json:"timePeriod"`
	JSON       dataExportScheduleGetResponseJSON       `json:"-"`
	union      DataExportScheduleGetResponseUnion
}

// dataExportScheduleGetResponseJSON contains the JSON metadata for the struct
// [DataExportScheduleGetResponse]
type dataExportScheduleGetResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	AccountIDs           apijson.Field
	Aggregation          apijson.Field
	AggregationFrequency apijson.Field
	MeterIDs             apijson.Field
	OperationalDataTypes apijson.Field
	TimePeriod           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r dataExportScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportScheduleGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportScheduleGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [OperationalDataExportScheduleResponse],
// [UsageDataExportScheduleResponse].
func (r DataExportScheduleGetResponse) AsUnion() DataExportScheduleGetResponseUnion {
	return r.union
}

// Response representing an operational data export configuration.
//
// Union satisfied by [OperationalDataExportScheduleResponse] or
// [UsageDataExportScheduleResponse].
type DataExportScheduleGetResponseUnion interface {
	implementsDataExportScheduleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportScheduleGetResponseUnion)(nil)).Elem(),
		"sourceType",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationalDataExportScheduleResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponse{}),
		},
	)
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
type DataExportScheduleGetResponseAggregation string

const (
	DataExportScheduleGetResponseAggregationSum    DataExportScheduleGetResponseAggregation = "SUM"
	DataExportScheduleGetResponseAggregationMin    DataExportScheduleGetResponseAggregation = "MIN"
	DataExportScheduleGetResponseAggregationMax    DataExportScheduleGetResponseAggregation = "MAX"
	DataExportScheduleGetResponseAggregationCount  DataExportScheduleGetResponseAggregation = "COUNT"
	DataExportScheduleGetResponseAggregationLatest DataExportScheduleGetResponseAggregation = "LATEST"
	DataExportScheduleGetResponseAggregationMean   DataExportScheduleGetResponseAggregation = "MEAN"
)

func (r DataExportScheduleGetResponseAggregation) IsKnown() bool {
	switch r {
	case DataExportScheduleGetResponseAggregationSum, DataExportScheduleGetResponseAggregationMin, DataExportScheduleGetResponseAggregationMax, DataExportScheduleGetResponseAggregationCount, DataExportScheduleGetResponseAggregationLatest, DataExportScheduleGetResponseAggregationMean:
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
type DataExportScheduleGetResponseAggregationFrequency string

const (
	DataExportScheduleGetResponseAggregationFrequencyOriginal DataExportScheduleGetResponseAggregationFrequency = "ORIGINAL"
	DataExportScheduleGetResponseAggregationFrequencyHour     DataExportScheduleGetResponseAggregationFrequency = "HOUR"
	DataExportScheduleGetResponseAggregationFrequencyDay      DataExportScheduleGetResponseAggregationFrequency = "DAY"
	DataExportScheduleGetResponseAggregationFrequencyWeek     DataExportScheduleGetResponseAggregationFrequency = "WEEK"
	DataExportScheduleGetResponseAggregationFrequencyMonth    DataExportScheduleGetResponseAggregationFrequency = "MONTH"
)

func (r DataExportScheduleGetResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportScheduleGetResponseAggregationFrequencyOriginal, DataExportScheduleGetResponseAggregationFrequencyHour, DataExportScheduleGetResponseAggregationFrequencyDay, DataExportScheduleGetResponseAggregationFrequencyWeek, DataExportScheduleGetResponseAggregationFrequencyMonth:
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
type DataExportScheduleGetResponseTimePeriod string

const (
	DataExportScheduleGetResponseTimePeriodToday         DataExportScheduleGetResponseTimePeriod = "TODAY"
	DataExportScheduleGetResponseTimePeriodYesterday     DataExportScheduleGetResponseTimePeriod = "YESTERDAY"
	DataExportScheduleGetResponseTimePeriodWeekToDate    DataExportScheduleGetResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleGetResponseTimePeriodCurrentMonth  DataExportScheduleGetResponseTimePeriod = "CURRENT_MONTH"
	DataExportScheduleGetResponseTimePeriodLast30Days    DataExportScheduleGetResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleGetResponseTimePeriodLast35Days    DataExportScheduleGetResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleGetResponseTimePeriodPreviousWeek  DataExportScheduleGetResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleGetResponseTimePeriodPreviousMonth DataExportScheduleGetResponseTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportScheduleGetResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleGetResponseTimePeriodToday, DataExportScheduleGetResponseTimePeriodYesterday, DataExportScheduleGetResponseTimePeriodWeekToDate, DataExportScheduleGetResponseTimePeriodCurrentMonth, DataExportScheduleGetResponseTimePeriodLast30Days, DataExportScheduleGetResponseTimePeriodLast35Days, DataExportScheduleGetResponseTimePeriodPreviousWeek, DataExportScheduleGetResponseTimePeriodPreviousMonth:
		return true
	}
	return false
}

// Response representing an operational data export configuration.
type DataExportScheduleUpdateResponse struct {
	// The id of the schedule.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// This field can have the runtime type of [[]string].
	AccountIDs interface{} `json:"accountIds"`
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
	Aggregation DataExportScheduleUpdateResponseAggregation `json:"aggregation"`
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
	AggregationFrequency DataExportScheduleUpdateResponseAggregationFrequency `json:"aggregationFrequency"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
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
	TimePeriod DataExportScheduleUpdateResponseTimePeriod `json:"timePeriod"`
	JSON       dataExportScheduleUpdateResponseJSON       `json:"-"`
	union      DataExportScheduleUpdateResponseUnion
}

// dataExportScheduleUpdateResponseJSON contains the JSON metadata for the struct
// [DataExportScheduleUpdateResponse]
type dataExportScheduleUpdateResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	AccountIDs           apijson.Field
	Aggregation          apijson.Field
	AggregationFrequency apijson.Field
	MeterIDs             apijson.Field
	OperationalDataTypes apijson.Field
	TimePeriod           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r dataExportScheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportScheduleUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportScheduleUpdateResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [OperationalDataExportScheduleResponse],
// [UsageDataExportScheduleResponse].
func (r DataExportScheduleUpdateResponse) AsUnion() DataExportScheduleUpdateResponseUnion {
	return r.union
}

// Response representing an operational data export configuration.
//
// Union satisfied by [OperationalDataExportScheduleResponse] or
// [UsageDataExportScheduleResponse].
type DataExportScheduleUpdateResponseUnion interface {
	implementsDataExportScheduleUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportScheduleUpdateResponseUnion)(nil)).Elem(),
		"sourceType",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationalDataExportScheduleResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponse{}),
		},
	)
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
type DataExportScheduleUpdateResponseAggregation string

const (
	DataExportScheduleUpdateResponseAggregationSum    DataExportScheduleUpdateResponseAggregation = "SUM"
	DataExportScheduleUpdateResponseAggregationMin    DataExportScheduleUpdateResponseAggregation = "MIN"
	DataExportScheduleUpdateResponseAggregationMax    DataExportScheduleUpdateResponseAggregation = "MAX"
	DataExportScheduleUpdateResponseAggregationCount  DataExportScheduleUpdateResponseAggregation = "COUNT"
	DataExportScheduleUpdateResponseAggregationLatest DataExportScheduleUpdateResponseAggregation = "LATEST"
	DataExportScheduleUpdateResponseAggregationMean   DataExportScheduleUpdateResponseAggregation = "MEAN"
)

func (r DataExportScheduleUpdateResponseAggregation) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateResponseAggregationSum, DataExportScheduleUpdateResponseAggregationMin, DataExportScheduleUpdateResponseAggregationMax, DataExportScheduleUpdateResponseAggregationCount, DataExportScheduleUpdateResponseAggregationLatest, DataExportScheduleUpdateResponseAggregationMean:
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
type DataExportScheduleUpdateResponseAggregationFrequency string

const (
	DataExportScheduleUpdateResponseAggregationFrequencyOriginal DataExportScheduleUpdateResponseAggregationFrequency = "ORIGINAL"
	DataExportScheduleUpdateResponseAggregationFrequencyHour     DataExportScheduleUpdateResponseAggregationFrequency = "HOUR"
	DataExportScheduleUpdateResponseAggregationFrequencyDay      DataExportScheduleUpdateResponseAggregationFrequency = "DAY"
	DataExportScheduleUpdateResponseAggregationFrequencyWeek     DataExportScheduleUpdateResponseAggregationFrequency = "WEEK"
	DataExportScheduleUpdateResponseAggregationFrequencyMonth    DataExportScheduleUpdateResponseAggregationFrequency = "MONTH"
)

func (r DataExportScheduleUpdateResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateResponseAggregationFrequencyOriginal, DataExportScheduleUpdateResponseAggregationFrequencyHour, DataExportScheduleUpdateResponseAggregationFrequencyDay, DataExportScheduleUpdateResponseAggregationFrequencyWeek, DataExportScheduleUpdateResponseAggregationFrequencyMonth:
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
type DataExportScheduleUpdateResponseTimePeriod string

const (
	DataExportScheduleUpdateResponseTimePeriodToday         DataExportScheduleUpdateResponseTimePeriod = "TODAY"
	DataExportScheduleUpdateResponseTimePeriodYesterday     DataExportScheduleUpdateResponseTimePeriod = "YESTERDAY"
	DataExportScheduleUpdateResponseTimePeriodWeekToDate    DataExportScheduleUpdateResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleUpdateResponseTimePeriodCurrentMonth  DataExportScheduleUpdateResponseTimePeriod = "CURRENT_MONTH"
	DataExportScheduleUpdateResponseTimePeriodLast30Days    DataExportScheduleUpdateResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleUpdateResponseTimePeriodLast35Days    DataExportScheduleUpdateResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleUpdateResponseTimePeriodPreviousWeek  DataExportScheduleUpdateResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleUpdateResponseTimePeriodPreviousMonth DataExportScheduleUpdateResponseTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportScheduleUpdateResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateResponseTimePeriodToday, DataExportScheduleUpdateResponseTimePeriodYesterday, DataExportScheduleUpdateResponseTimePeriodWeekToDate, DataExportScheduleUpdateResponseTimePeriodCurrentMonth, DataExportScheduleUpdateResponseTimePeriodLast30Days, DataExportScheduleUpdateResponseTimePeriodLast35Days, DataExportScheduleUpdateResponseTimePeriodPreviousWeek, DataExportScheduleUpdateResponseTimePeriodPreviousMonth:
		return true
	}
	return false
}

type DataExportScheduleListResponse struct {
	// The id of the Data Export Schedule.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Unique short code of the Data Export Schedule.
	Code string `json:"code"`
	// The id of the user who created this Schedule.
	CreatedBy string `json:"createdBy"`
	// The Export Destination ids.
	DestinationIDs []string `json:"destinationIds"`
	// The DateTime when the Data Export Schedule was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Schedule was last modified.
	DtLastModified   time.Time                                      `json:"dtLastModified" format:"date-time"`
	ExportFileFormat DataExportScheduleListResponseExportFileFormat `json:"exportFileFormat"`
	// The id of the user who last modified this Data Export Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Data Export Schedule.
	Name string `json:"name"`
	// Defines the Schedule frequency for the Data Export to run in Hours or Days. Used
	// in conjunction with the `scheduleType` parameter.
	Period       int64                                      `json:"period"`
	ScheduleType DataExportScheduleListResponseScheduleType `json:"scheduleType"`
	SourceType   DataExportScheduleListResponseSourceType   `json:"sourceType"`
	JSON         dataExportScheduleListResponseJSON         `json:"-"`
}

// dataExportScheduleListResponseJSON contains the JSON metadata for the struct
// [DataExportScheduleListResponse]
type dataExportScheduleListResponseJSON struct {
	ID               apijson.Field
	Version          apijson.Field
	Code             apijson.Field
	CreatedBy        apijson.Field
	DestinationIDs   apijson.Field
	DtCreated        apijson.Field
	DtLastModified   apijson.Field
	ExportFileFormat apijson.Field
	LastModifiedBy   apijson.Field
	Name             apijson.Field
	Period           apijson.Field
	ScheduleType     apijson.Field
	SourceType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DataExportScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportScheduleListResponseJSON) RawJSON() string {
	return r.raw
}

type DataExportScheduleListResponseExportFileFormat string

const (
	DataExportScheduleListResponseExportFileFormatCsv  DataExportScheduleListResponseExportFileFormat = "CSV"
	DataExportScheduleListResponseExportFileFormatJson DataExportScheduleListResponseExportFileFormat = "JSON"
)

func (r DataExportScheduleListResponseExportFileFormat) IsKnown() bool {
	switch r {
	case DataExportScheduleListResponseExportFileFormatCsv, DataExportScheduleListResponseExportFileFormatJson:
		return true
	}
	return false
}

type DataExportScheduleListResponseScheduleType string

const (
	DataExportScheduleListResponseScheduleTypeHourly DataExportScheduleListResponseScheduleType = "HOURLY"
	DataExportScheduleListResponseScheduleTypeDaily  DataExportScheduleListResponseScheduleType = "DAILY"
	DataExportScheduleListResponseScheduleTypeAdHoc  DataExportScheduleListResponseScheduleType = "AD_HOC"
)

func (r DataExportScheduleListResponseScheduleType) IsKnown() bool {
	switch r {
	case DataExportScheduleListResponseScheduleTypeHourly, DataExportScheduleListResponseScheduleTypeDaily, DataExportScheduleListResponseScheduleTypeAdHoc:
		return true
	}
	return false
}

type DataExportScheduleListResponseSourceType string

const (
	DataExportScheduleListResponseSourceTypeUsage       DataExportScheduleListResponseSourceType = "USAGE"
	DataExportScheduleListResponseSourceTypeOperational DataExportScheduleListResponseSourceType = "OPERATIONAL"
)

func (r DataExportScheduleListResponseSourceType) IsKnown() bool {
	switch r {
	case DataExportScheduleListResponseSourceTypeUsage, DataExportScheduleListResponseSourceTypeOperational:
		return true
	}
	return false
}

// Response representing an operational data export configuration.
type DataExportScheduleDeleteResponse struct {
	// The id of the schedule.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// This field can have the runtime type of [[]string].
	AccountIDs interface{} `json:"accountIds"`
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
	Aggregation DataExportScheduleDeleteResponseAggregation `json:"aggregation"`
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
	AggregationFrequency DataExportScheduleDeleteResponseAggregationFrequency `json:"aggregationFrequency"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
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
	TimePeriod DataExportScheduleDeleteResponseTimePeriod `json:"timePeriod"`
	JSON       dataExportScheduleDeleteResponseJSON       `json:"-"`
	union      DataExportScheduleDeleteResponseUnion
}

// dataExportScheduleDeleteResponseJSON contains the JSON metadata for the struct
// [DataExportScheduleDeleteResponse]
type dataExportScheduleDeleteResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	AccountIDs           apijson.Field
	Aggregation          apijson.Field
	AggregationFrequency apijson.Field
	MeterIDs             apijson.Field
	OperationalDataTypes apijson.Field
	TimePeriod           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r dataExportScheduleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportScheduleDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportScheduleDeleteResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [OperationalDataExportScheduleResponse],
// [UsageDataExportScheduleResponse].
func (r DataExportScheduleDeleteResponse) AsUnion() DataExportScheduleDeleteResponseUnion {
	return r.union
}

// Response representing an operational data export configuration.
//
// Union satisfied by [OperationalDataExportScheduleResponse] or
// [UsageDataExportScheduleResponse].
type DataExportScheduleDeleteResponseUnion interface {
	implementsDataExportScheduleDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportScheduleDeleteResponseUnion)(nil)).Elem(),
		"sourceType",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationalDataExportScheduleResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponse{}),
		},
	)
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
type DataExportScheduleDeleteResponseAggregation string

const (
	DataExportScheduleDeleteResponseAggregationSum    DataExportScheduleDeleteResponseAggregation = "SUM"
	DataExportScheduleDeleteResponseAggregationMin    DataExportScheduleDeleteResponseAggregation = "MIN"
	DataExportScheduleDeleteResponseAggregationMax    DataExportScheduleDeleteResponseAggregation = "MAX"
	DataExportScheduleDeleteResponseAggregationCount  DataExportScheduleDeleteResponseAggregation = "COUNT"
	DataExportScheduleDeleteResponseAggregationLatest DataExportScheduleDeleteResponseAggregation = "LATEST"
	DataExportScheduleDeleteResponseAggregationMean   DataExportScheduleDeleteResponseAggregation = "MEAN"
)

func (r DataExportScheduleDeleteResponseAggregation) IsKnown() bool {
	switch r {
	case DataExportScheduleDeleteResponseAggregationSum, DataExportScheduleDeleteResponseAggregationMin, DataExportScheduleDeleteResponseAggregationMax, DataExportScheduleDeleteResponseAggregationCount, DataExportScheduleDeleteResponseAggregationLatest, DataExportScheduleDeleteResponseAggregationMean:
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
type DataExportScheduleDeleteResponseAggregationFrequency string

const (
	DataExportScheduleDeleteResponseAggregationFrequencyOriginal DataExportScheduleDeleteResponseAggregationFrequency = "ORIGINAL"
	DataExportScheduleDeleteResponseAggregationFrequencyHour     DataExportScheduleDeleteResponseAggregationFrequency = "HOUR"
	DataExportScheduleDeleteResponseAggregationFrequencyDay      DataExportScheduleDeleteResponseAggregationFrequency = "DAY"
	DataExportScheduleDeleteResponseAggregationFrequencyWeek     DataExportScheduleDeleteResponseAggregationFrequency = "WEEK"
	DataExportScheduleDeleteResponseAggregationFrequencyMonth    DataExportScheduleDeleteResponseAggregationFrequency = "MONTH"
)

func (r DataExportScheduleDeleteResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportScheduleDeleteResponseAggregationFrequencyOriginal, DataExportScheduleDeleteResponseAggregationFrequencyHour, DataExportScheduleDeleteResponseAggregationFrequencyDay, DataExportScheduleDeleteResponseAggregationFrequencyWeek, DataExportScheduleDeleteResponseAggregationFrequencyMonth:
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
type DataExportScheduleDeleteResponseTimePeriod string

const (
	DataExportScheduleDeleteResponseTimePeriodToday         DataExportScheduleDeleteResponseTimePeriod = "TODAY"
	DataExportScheduleDeleteResponseTimePeriodYesterday     DataExportScheduleDeleteResponseTimePeriod = "YESTERDAY"
	DataExportScheduleDeleteResponseTimePeriodWeekToDate    DataExportScheduleDeleteResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleDeleteResponseTimePeriodCurrentMonth  DataExportScheduleDeleteResponseTimePeriod = "CURRENT_MONTH"
	DataExportScheduleDeleteResponseTimePeriodLast30Days    DataExportScheduleDeleteResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleDeleteResponseTimePeriodLast35Days    DataExportScheduleDeleteResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleDeleteResponseTimePeriodPreviousWeek  DataExportScheduleDeleteResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleDeleteResponseTimePeriodPreviousMonth DataExportScheduleDeleteResponseTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportScheduleDeleteResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleDeleteResponseTimePeriodToday, DataExportScheduleDeleteResponseTimePeriodYesterday, DataExportScheduleDeleteResponseTimePeriodWeekToDate, DataExportScheduleDeleteResponseTimePeriodCurrentMonth, DataExportScheduleDeleteResponseTimePeriodLast30Days, DataExportScheduleDeleteResponseTimePeriodLast35Days, DataExportScheduleDeleteResponseTimePeriodPreviousWeek, DataExportScheduleDeleteResponseTimePeriodPreviousMonth:
		return true
	}
	return false
}

type DataExportScheduleNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Request representing an operational schedule configuration.
	Body DataExportScheduleNewParamsBodyUnion `json:"body,required"`
}

func (r DataExportScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request representing an operational schedule configuration.
type DataExportScheduleNewParamsBody struct {
	SourceType param.Field[DataExportScheduleNewParamsBodySourceType] `json:"sourceType,required"`
	AccountIDs param.Field[interface{}]                               `json:"accountIds"`
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
	Aggregation param.Field[DataExportScheduleNewParamsBodyAggregation] `json:"aggregation"`
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
	AggregationFrequency param.Field[DataExportScheduleNewParamsBodyAggregationFrequency] `json:"aggregationFrequency"`
	MeterIDs             param.Field[interface{}]                                         `json:"meterIds"`
	OperationalDataTypes param.Field[interface{}]                                         `json:"operationalDataTypes"`
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
	TimePeriod param.Field[DataExportScheduleNewParamsBodyTimePeriod] `json:"timePeriod"`
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

func (r DataExportScheduleNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportScheduleNewParamsBody) implementsDataExportScheduleNewParamsBodyUnion() {}

// Request representing an operational schedule configuration.
//
// Satisfied by [OperationalDataExportScheduleRequestParam],
// [UsageDataExportScheduleRequestParam], [DataExportScheduleNewParamsBody].
type DataExportScheduleNewParamsBodyUnion interface {
	implementsDataExportScheduleNewParamsBodyUnion()
}

type DataExportScheduleNewParamsBodySourceType string

const (
	DataExportScheduleNewParamsBodySourceTypeUsage       DataExportScheduleNewParamsBodySourceType = "USAGE"
	DataExportScheduleNewParamsBodySourceTypeOperational DataExportScheduleNewParamsBodySourceType = "OPERATIONAL"
)

func (r DataExportScheduleNewParamsBodySourceType) IsKnown() bool {
	switch r {
	case DataExportScheduleNewParamsBodySourceTypeUsage, DataExportScheduleNewParamsBodySourceTypeOperational:
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
type DataExportScheduleNewParamsBodyAggregation string

const (
	DataExportScheduleNewParamsBodyAggregationSum    DataExportScheduleNewParamsBodyAggregation = "SUM"
	DataExportScheduleNewParamsBodyAggregationMin    DataExportScheduleNewParamsBodyAggregation = "MIN"
	DataExportScheduleNewParamsBodyAggregationMax    DataExportScheduleNewParamsBodyAggregation = "MAX"
	DataExportScheduleNewParamsBodyAggregationCount  DataExportScheduleNewParamsBodyAggregation = "COUNT"
	DataExportScheduleNewParamsBodyAggregationLatest DataExportScheduleNewParamsBodyAggregation = "LATEST"
	DataExportScheduleNewParamsBodyAggregationMean   DataExportScheduleNewParamsBodyAggregation = "MEAN"
)

func (r DataExportScheduleNewParamsBodyAggregation) IsKnown() bool {
	switch r {
	case DataExportScheduleNewParamsBodyAggregationSum, DataExportScheduleNewParamsBodyAggregationMin, DataExportScheduleNewParamsBodyAggregationMax, DataExportScheduleNewParamsBodyAggregationCount, DataExportScheduleNewParamsBodyAggregationLatest, DataExportScheduleNewParamsBodyAggregationMean:
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
type DataExportScheduleNewParamsBodyAggregationFrequency string

const (
	DataExportScheduleNewParamsBodyAggregationFrequencyOriginal DataExportScheduleNewParamsBodyAggregationFrequency = "ORIGINAL"
	DataExportScheduleNewParamsBodyAggregationFrequencyHour     DataExportScheduleNewParamsBodyAggregationFrequency = "HOUR"
	DataExportScheduleNewParamsBodyAggregationFrequencyDay      DataExportScheduleNewParamsBodyAggregationFrequency = "DAY"
	DataExportScheduleNewParamsBodyAggregationFrequencyWeek     DataExportScheduleNewParamsBodyAggregationFrequency = "WEEK"
	DataExportScheduleNewParamsBodyAggregationFrequencyMonth    DataExportScheduleNewParamsBodyAggregationFrequency = "MONTH"
)

func (r DataExportScheduleNewParamsBodyAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportScheduleNewParamsBodyAggregationFrequencyOriginal, DataExportScheduleNewParamsBodyAggregationFrequencyHour, DataExportScheduleNewParamsBodyAggregationFrequencyDay, DataExportScheduleNewParamsBodyAggregationFrequencyWeek, DataExportScheduleNewParamsBodyAggregationFrequencyMonth:
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
type DataExportScheduleNewParamsBodyTimePeriod string

const (
	DataExportScheduleNewParamsBodyTimePeriodToday         DataExportScheduleNewParamsBodyTimePeriod = "TODAY"
	DataExportScheduleNewParamsBodyTimePeriodYesterday     DataExportScheduleNewParamsBodyTimePeriod = "YESTERDAY"
	DataExportScheduleNewParamsBodyTimePeriodWeekToDate    DataExportScheduleNewParamsBodyTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleNewParamsBodyTimePeriodCurrentMonth  DataExportScheduleNewParamsBodyTimePeriod = "CURRENT_MONTH"
	DataExportScheduleNewParamsBodyTimePeriodLast30Days    DataExportScheduleNewParamsBodyTimePeriod = "LAST_30_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodLast35Days    DataExportScheduleNewParamsBodyTimePeriod = "LAST_35_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodPreviousWeek  DataExportScheduleNewParamsBodyTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleNewParamsBodyTimePeriodPreviousMonth DataExportScheduleNewParamsBodyTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportScheduleNewParamsBodyTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleNewParamsBodyTimePeriodToday, DataExportScheduleNewParamsBodyTimePeriodYesterday, DataExportScheduleNewParamsBodyTimePeriodWeekToDate, DataExportScheduleNewParamsBodyTimePeriodCurrentMonth, DataExportScheduleNewParamsBodyTimePeriodLast30Days, DataExportScheduleNewParamsBodyTimePeriodLast35Days, DataExportScheduleNewParamsBodyTimePeriodPreviousWeek, DataExportScheduleNewParamsBodyTimePeriodPreviousMonth:
		return true
	}
	return false
}

type DataExportScheduleGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type DataExportScheduleUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Request representing an operational schedule configuration.
	Body DataExportScheduleUpdateParamsBodyUnion `json:"body,required"`
}

func (r DataExportScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request representing an operational schedule configuration.
type DataExportScheduleUpdateParamsBody struct {
	SourceType param.Field[DataExportScheduleUpdateParamsBodySourceType] `json:"sourceType,required"`
	AccountIDs param.Field[interface{}]                                  `json:"accountIds"`
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
	Aggregation param.Field[DataExportScheduleUpdateParamsBodyAggregation] `json:"aggregation"`
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
	AggregationFrequency param.Field[DataExportScheduleUpdateParamsBodyAggregationFrequency] `json:"aggregationFrequency"`
	MeterIDs             param.Field[interface{}]                                            `json:"meterIds"`
	OperationalDataTypes param.Field[interface{}]                                            `json:"operationalDataTypes"`
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
	TimePeriod param.Field[DataExportScheduleUpdateParamsBodyTimePeriod] `json:"timePeriod"`
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

func (r DataExportScheduleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportScheduleUpdateParamsBody) implementsDataExportScheduleUpdateParamsBodyUnion() {}

// Request representing an operational schedule configuration.
//
// Satisfied by [OperationalDataExportScheduleRequestParam],
// [UsageDataExportScheduleRequestParam], [DataExportScheduleUpdateParamsBody].
type DataExportScheduleUpdateParamsBodyUnion interface {
	implementsDataExportScheduleUpdateParamsBodyUnion()
}

type DataExportScheduleUpdateParamsBodySourceType string

const (
	DataExportScheduleUpdateParamsBodySourceTypeUsage       DataExportScheduleUpdateParamsBodySourceType = "USAGE"
	DataExportScheduleUpdateParamsBodySourceTypeOperational DataExportScheduleUpdateParamsBodySourceType = "OPERATIONAL"
)

func (r DataExportScheduleUpdateParamsBodySourceType) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateParamsBodySourceTypeUsage, DataExportScheduleUpdateParamsBodySourceTypeOperational:
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
type DataExportScheduleUpdateParamsBodyAggregation string

const (
	DataExportScheduleUpdateParamsBodyAggregationSum    DataExportScheduleUpdateParamsBodyAggregation = "SUM"
	DataExportScheduleUpdateParamsBodyAggregationMin    DataExportScheduleUpdateParamsBodyAggregation = "MIN"
	DataExportScheduleUpdateParamsBodyAggregationMax    DataExportScheduleUpdateParamsBodyAggregation = "MAX"
	DataExportScheduleUpdateParamsBodyAggregationCount  DataExportScheduleUpdateParamsBodyAggregation = "COUNT"
	DataExportScheduleUpdateParamsBodyAggregationLatest DataExportScheduleUpdateParamsBodyAggregation = "LATEST"
	DataExportScheduleUpdateParamsBodyAggregationMean   DataExportScheduleUpdateParamsBodyAggregation = "MEAN"
)

func (r DataExportScheduleUpdateParamsBodyAggregation) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateParamsBodyAggregationSum, DataExportScheduleUpdateParamsBodyAggregationMin, DataExportScheduleUpdateParamsBodyAggregationMax, DataExportScheduleUpdateParamsBodyAggregationCount, DataExportScheduleUpdateParamsBodyAggregationLatest, DataExportScheduleUpdateParamsBodyAggregationMean:
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
type DataExportScheduleUpdateParamsBodyAggregationFrequency string

const (
	DataExportScheduleUpdateParamsBodyAggregationFrequencyOriginal DataExportScheduleUpdateParamsBodyAggregationFrequency = "ORIGINAL"
	DataExportScheduleUpdateParamsBodyAggregationFrequencyHour     DataExportScheduleUpdateParamsBodyAggregationFrequency = "HOUR"
	DataExportScheduleUpdateParamsBodyAggregationFrequencyDay      DataExportScheduleUpdateParamsBodyAggregationFrequency = "DAY"
	DataExportScheduleUpdateParamsBodyAggregationFrequencyWeek     DataExportScheduleUpdateParamsBodyAggregationFrequency = "WEEK"
	DataExportScheduleUpdateParamsBodyAggregationFrequencyMonth    DataExportScheduleUpdateParamsBodyAggregationFrequency = "MONTH"
)

func (r DataExportScheduleUpdateParamsBodyAggregationFrequency) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateParamsBodyAggregationFrequencyOriginal, DataExportScheduleUpdateParamsBodyAggregationFrequencyHour, DataExportScheduleUpdateParamsBodyAggregationFrequencyDay, DataExportScheduleUpdateParamsBodyAggregationFrequencyWeek, DataExportScheduleUpdateParamsBodyAggregationFrequencyMonth:
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
type DataExportScheduleUpdateParamsBodyTimePeriod string

const (
	DataExportScheduleUpdateParamsBodyTimePeriodToday         DataExportScheduleUpdateParamsBodyTimePeriod = "TODAY"
	DataExportScheduleUpdateParamsBodyTimePeriodYesterday     DataExportScheduleUpdateParamsBodyTimePeriod = "YESTERDAY"
	DataExportScheduleUpdateParamsBodyTimePeriodWeekToDate    DataExportScheduleUpdateParamsBodyTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleUpdateParamsBodyTimePeriodCurrentMonth  DataExportScheduleUpdateParamsBodyTimePeriod = "CURRENT_MONTH"
	DataExportScheduleUpdateParamsBodyTimePeriodLast30Days    DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_30_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodLast35Days    DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_35_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodPreviousWeek  DataExportScheduleUpdateParamsBodyTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleUpdateParamsBodyTimePeriodPreviousMonth DataExportScheduleUpdateParamsBodyTimePeriod = "PREVIOUS_MONTH"
)

func (r DataExportScheduleUpdateParamsBodyTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateParamsBodyTimePeriodToday, DataExportScheduleUpdateParamsBodyTimePeriodYesterday, DataExportScheduleUpdateParamsBodyTimePeriodWeekToDate, DataExportScheduleUpdateParamsBodyTimePeriodCurrentMonth, DataExportScheduleUpdateParamsBodyTimePeriodLast30Days, DataExportScheduleUpdateParamsBodyTimePeriodLast35Days, DataExportScheduleUpdateParamsBodyTimePeriodPreviousWeek, DataExportScheduleUpdateParamsBodyTimePeriodPreviousMonth:
		return true
	}
	return false
}

type DataExportScheduleListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Data Export Schedule IDs to filter the returned list by.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of schedules to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [DataExportScheduleListParams]'s query parameters as
// `url.Values`.
func (r DataExportScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportScheduleDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
