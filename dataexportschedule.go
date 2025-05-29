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
//   - You can use the `dimensionFilters` parameter to filter the usage data returned
//     for export by adding specific values of non-numeric Dimension data fields on
//     included Meters. Only the data collected for the values you've added for the
//     selected Dimension fields will be included in the export.
//   - You can use the `aggregations` to apply aggregation methods the usage data
//     returned for export. This restricts the range of usage data returned for
//     export to only the data collected by aggregated fields on selected Meters.
//     Nothing is returned for any non-aggregated fields on Meters. The usage data
//     for Meter fields is returned as the values resulting from applying the
//     selected aggregation method. See the
//     [Aggregations for Queries - Options and Consequences](https://www.m3ter.com/docs/guides/data-explorer/usage-data-explorer-v2#aggregations-for-queries---understanding-options-and-consequences)
//     for more details.
//   - If you've applied `aggregations` to the usage returned for export, you can
//     then use the `groups` parameter to group the data by _Account_, _Dimension_,
//     or _Time_.
//
// Request and Response schema:
//
//   - Use the selector under the `sourceType` parameter to expose the relevant
//     request and response schema for the source data type.
//
// Request and Response samples:
//
//   - Use the **Example** selector to show the relevant request and response samples
//     for source data type.
func (r *DataExportScheduleService) New(ctx context.Context, params DataExportScheduleNewParams, opts ...option.RequestOption) (res *DataExportScheduleNewResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/dataexports/schedules", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Data Export Schedule for the given UUID. Each Schedule can be
// configured for exporting _only one_ of either Usage or Operational data.
func (r *DataExportScheduleService) Get(ctx context.Context, id string, query DataExportScheduleGetParams, opts ...option.RequestOption) (res *DataExportScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
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
//   - You can use the `dimensionFilters` parameter to filter the usage data returned
//     for export by adding specific values of non-numeric Dimension data fields on
//     included Meters. Only the data collected for the values you've added for the
//     selected Dimension fields will be included in the export.
//   - You can use the `aggregations` to apply aggregation methods the usage data
//     returned for export. This restricts the range of usage data returned for
//     export to only the data collected by aggregated fields on selected Meters.
//     Nothing is returned for any non-aggregated fields on Meters. The usage data
//     for Meter fields is returned as the values resulting from applying the
//     selected aggregation method. See the
//     [Aggregations for Queries - Options and Consequences](https://www.m3ter.com/docs/guides/data-explorer/usage-data-explorer-v2#aggregations-for-queries---understanding-options-and-consequences)
//     for more details.
//   - If you've applied `aggregations` to the usage returned for export, you can
//     then use the `groups` parameter to group the data by _Account_, _Dimension_,
//     or _Time_.
func (r *DataExportScheduleService) Update(ctx context.Context, id string, params DataExportScheduleUpdateParams, opts ...option.RequestOption) (res *DataExportScheduleUpdateResponse, err error) {
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
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
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
	path := fmt.Sprintf("organizations/%s/dataexports/schedules/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type OperationalDataExportScheduleRequestParam struct {
	// A list of the entities whose operational data is included in the data export.
	OperationalDataTypes param.Field[[]OperationalDataExportScheduleRequestOperationalDataType] `json:"operationalDataTypes,required"`
	// The type of data to export. Possible values are: OPERATIONAL
	SourceType param.Field[OperationalDataExportScheduleRequestSourceType] `json:"sourceType,required"`
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

// The type of data to export. Possible values are: OPERATIONAL
type OperationalDataExportScheduleRequestSourceType string

const (
	OperationalDataExportScheduleRequestSourceTypeOperational OperationalDataExportScheduleRequestSourceType = "OPERATIONAL"
)

func (r OperationalDataExportScheduleRequestSourceType) IsKnown() bool {
	switch r {
	case OperationalDataExportScheduleRequestSourceTypeOperational:
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
	// The type of data to export. Possible values are: USAGE
	SourceType param.Field[UsageDataExportScheduleRequestSourceType] `json:"sourceType,required"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
	//
	// For more details and examples, see the
	// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
	// section in our main User Documentation.
	TimePeriod param.Field[UsageDataExportScheduleRequestTimePeriod] `json:"timePeriod,required"`
	// List of account IDs to export
	AccountIDs param.Field[[]string] `json:"accountIds"`
	// List of aggregations to apply
	Aggregations param.Field[[]UsageDataExportScheduleRequestAggregationParam] `json:"aggregations"`
	// List of dimension filters to apply
	DimensionFilters param.Field[[]UsageDataExportScheduleRequestDimensionFilterParam] `json:"dimensionFilters"`
	// List of groups to apply
	Groups param.Field[[]UsageDataExportScheduleRequestGroupsUnionParam] `json:"groups"`
	// List of meter IDs to export
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

// The type of data to export. Possible values are: USAGE
type UsageDataExportScheduleRequestSourceType string

const (
	UsageDataExportScheduleRequestSourceTypeUsage UsageDataExportScheduleRequestSourceType = "USAGE"
)

func (r UsageDataExportScheduleRequestSourceType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestSourceTypeUsage:
		return true
	}
	return false
}

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type UsageDataExportScheduleRequestTimePeriod string

const (
	UsageDataExportScheduleRequestTimePeriodToday           UsageDataExportScheduleRequestTimePeriod = "TODAY"
	UsageDataExportScheduleRequestTimePeriodYesterday       UsageDataExportScheduleRequestTimePeriod = "YESTERDAY"
	UsageDataExportScheduleRequestTimePeriodWeekToDate      UsageDataExportScheduleRequestTimePeriod = "WEEK_TO_DATE"
	UsageDataExportScheduleRequestTimePeriodMonthToDate     UsageDataExportScheduleRequestTimePeriod = "MONTH_TO_DATE"
	UsageDataExportScheduleRequestTimePeriodYearToDate      UsageDataExportScheduleRequestTimePeriod = "YEAR_TO_DATE"
	UsageDataExportScheduleRequestTimePeriodPreviousWeek    UsageDataExportScheduleRequestTimePeriod = "PREVIOUS_WEEK"
	UsageDataExportScheduleRequestTimePeriodPreviousMonth   UsageDataExportScheduleRequestTimePeriod = "PREVIOUS_MONTH"
	UsageDataExportScheduleRequestTimePeriodPreviousQuarter UsageDataExportScheduleRequestTimePeriod = "PREVIOUS_QUARTER"
	UsageDataExportScheduleRequestTimePeriodPreviousYear    UsageDataExportScheduleRequestTimePeriod = "PREVIOUS_YEAR"
	UsageDataExportScheduleRequestTimePeriodLast12Hours     UsageDataExportScheduleRequestTimePeriod = "LAST_12_HOURS"
	UsageDataExportScheduleRequestTimePeriodLast7Days       UsageDataExportScheduleRequestTimePeriod = "LAST_7_DAYS"
	UsageDataExportScheduleRequestTimePeriodLast30Days      UsageDataExportScheduleRequestTimePeriod = "LAST_30_DAYS"
	UsageDataExportScheduleRequestTimePeriodLast35Days      UsageDataExportScheduleRequestTimePeriod = "LAST_35_DAYS"
	UsageDataExportScheduleRequestTimePeriodLast90Days      UsageDataExportScheduleRequestTimePeriod = "LAST_90_DAYS"
	UsageDataExportScheduleRequestTimePeriodLast120Days     UsageDataExportScheduleRequestTimePeriod = "LAST_120_DAYS"
	UsageDataExportScheduleRequestTimePeriodLastYear        UsageDataExportScheduleRequestTimePeriod = "LAST_YEAR"
)

func (r UsageDataExportScheduleRequestTimePeriod) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestTimePeriodToday, UsageDataExportScheduleRequestTimePeriodYesterday, UsageDataExportScheduleRequestTimePeriodWeekToDate, UsageDataExportScheduleRequestTimePeriodMonthToDate, UsageDataExportScheduleRequestTimePeriodYearToDate, UsageDataExportScheduleRequestTimePeriodPreviousWeek, UsageDataExportScheduleRequestTimePeriodPreviousMonth, UsageDataExportScheduleRequestTimePeriodPreviousQuarter, UsageDataExportScheduleRequestTimePeriodPreviousYear, UsageDataExportScheduleRequestTimePeriodLast12Hours, UsageDataExportScheduleRequestTimePeriodLast7Days, UsageDataExportScheduleRequestTimePeriodLast30Days, UsageDataExportScheduleRequestTimePeriodLast35Days, UsageDataExportScheduleRequestTimePeriodLast90Days, UsageDataExportScheduleRequestTimePeriodLast120Days, UsageDataExportScheduleRequestTimePeriodLastYear:
		return true
	}
	return false
}

type UsageDataExportScheduleRequestAggregationParam struct {
	// Field code
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Type of field
	FieldType param.Field[UsageDataExportScheduleRequestAggregationsFieldType] `json:"fieldType,required"`
	// Aggregation function
	Function param.Field[UsageDataExportScheduleRequestAggregationsFunction] `json:"function,required"`
	// Meter ID
	MeterID param.Field[string] `json:"meterId,required"`
}

func (r UsageDataExportScheduleRequestAggregationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of field
type UsageDataExportScheduleRequestAggregationsFieldType string

const (
	UsageDataExportScheduleRequestAggregationsFieldTypeDimension UsageDataExportScheduleRequestAggregationsFieldType = "DIMENSION"
	UsageDataExportScheduleRequestAggregationsFieldTypeMeasure   UsageDataExportScheduleRequestAggregationsFieldType = "MEASURE"
)

func (r UsageDataExportScheduleRequestAggregationsFieldType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestAggregationsFieldTypeDimension, UsageDataExportScheduleRequestAggregationsFieldTypeMeasure:
		return true
	}
	return false
}

// Aggregation function
type UsageDataExportScheduleRequestAggregationsFunction string

const (
	UsageDataExportScheduleRequestAggregationsFunctionSum    UsageDataExportScheduleRequestAggregationsFunction = "SUM"
	UsageDataExportScheduleRequestAggregationsFunctionMin    UsageDataExportScheduleRequestAggregationsFunction = "MIN"
	UsageDataExportScheduleRequestAggregationsFunctionMax    UsageDataExportScheduleRequestAggregationsFunction = "MAX"
	UsageDataExportScheduleRequestAggregationsFunctionCount  UsageDataExportScheduleRequestAggregationsFunction = "COUNT"
	UsageDataExportScheduleRequestAggregationsFunctionLatest UsageDataExportScheduleRequestAggregationsFunction = "LATEST"
	UsageDataExportScheduleRequestAggregationsFunctionMean   UsageDataExportScheduleRequestAggregationsFunction = "MEAN"
	UsageDataExportScheduleRequestAggregationsFunctionUnique UsageDataExportScheduleRequestAggregationsFunction = "UNIQUE"
)

func (r UsageDataExportScheduleRequestAggregationsFunction) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestAggregationsFunctionSum, UsageDataExportScheduleRequestAggregationsFunctionMin, UsageDataExportScheduleRequestAggregationsFunctionMax, UsageDataExportScheduleRequestAggregationsFunctionCount, UsageDataExportScheduleRequestAggregationsFunctionLatest, UsageDataExportScheduleRequestAggregationsFunctionMean, UsageDataExportScheduleRequestAggregationsFunctionUnique:
		return true
	}
	return false
}

type UsageDataExportScheduleRequestDimensionFilterParam struct {
	// Field code
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Meter ID
	MeterID param.Field[string] `json:"meterId,required"`
	// Values to filter by
	Values param.Field[[]string] `json:"values,required"`
}

func (r UsageDataExportScheduleRequestDimensionFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Group by a field
//
// Satisfied by
// [UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupParam],
// [UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupParam],
// [UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupParam].
type UsageDataExportScheduleRequestGroupsUnionParam interface {
	implementsUsageDataExportScheduleRequestGroupsUnionParam()
}

// Group by account
type UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupParam struct {
	GroupType param.Field[UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupType] `json:"groupType"`
	DataExplorerAccountGroupParam
}

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupParam) implementsUsageDataExportScheduleRequestGroupsUnionParam() {
}

type UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupType string

const (
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount   UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupType = "ACCOUNT"
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupType = "DIMENSION"
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeTime      UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupType = "TIME"
)

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount, UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension, UsageDataExportScheduleRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by dimension
type UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupParam struct {
	GroupType param.Field[UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupType] `json:"groupType"`
	DataExplorerDimensionGroupParam
}

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupParam) implementsUsageDataExportScheduleRequestGroupsUnionParam() {
}

type UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupType string

const (
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount   UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupType = "ACCOUNT"
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupType = "DIMENSION"
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime      UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupType = "TIME"
)

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount, UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension, UsageDataExportScheduleRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by time
type UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupParam struct {
	GroupType param.Field[UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupType] `json:"groupType"`
	DataExplorerTimeGroupParam
}

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupParam) implementsUsageDataExportScheduleRequestGroupsUnionParam() {
}

type UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupType string

const (
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount   UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupType = "ACCOUNT"
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupType = "DIMENSION"
	UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeTime      UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupType = "TIME"
)

func (r UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount, UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension, UsageDataExportScheduleRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeTime:
		return true
	}
	return false
}

type UsageDataExportScheduleResponse struct {
	// The id of the schedule configuration.
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
	// List of aggregations to apply
	Aggregations []UsageDataExportScheduleResponseAggregation `json:"aggregations"`
	// List of dimension filters to apply
	DimensionFilters []UsageDataExportScheduleResponseDimensionFilter `json:"dimensionFilters"`
	// List of groups to apply
	Groups []UsageDataExportScheduleResponseGroup `json:"groups"`
	// List of meter IDs for which the usage data will be exported.
	MeterIDs []string `json:"meterIds"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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
	ID               apijson.Field
	Version          apijson.Field
	AccountIDs       apijson.Field
	Aggregations     apijson.Field
	DimensionFilters apijson.Field
	Groups           apijson.Field
	MeterIDs         apijson.Field
	TimePeriod       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
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

type UsageDataExportScheduleResponseAggregation struct {
	// Field code
	FieldCode string `json:"fieldCode,required"`
	// Type of field
	FieldType UsageDataExportScheduleResponseAggregationsFieldType `json:"fieldType,required"`
	// Aggregation function
	Function UsageDataExportScheduleResponseAggregationsFunction `json:"function,required"`
	// Meter ID
	MeterID string                                         `json:"meterId,required"`
	JSON    usageDataExportScheduleResponseAggregationJSON `json:"-"`
}

// usageDataExportScheduleResponseAggregationJSON contains the JSON metadata for
// the struct [UsageDataExportScheduleResponseAggregation]
type usageDataExportScheduleResponseAggregationJSON struct {
	FieldCode   apijson.Field
	FieldType   apijson.Field
	Function    apijson.Field
	MeterID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageDataExportScheduleResponseAggregation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDataExportScheduleResponseAggregationJSON) RawJSON() string {
	return r.raw
}

// Type of field
type UsageDataExportScheduleResponseAggregationsFieldType string

const (
	UsageDataExportScheduleResponseAggregationsFieldTypeDimension UsageDataExportScheduleResponseAggregationsFieldType = "DIMENSION"
	UsageDataExportScheduleResponseAggregationsFieldTypeMeasure   UsageDataExportScheduleResponseAggregationsFieldType = "MEASURE"
)

func (r UsageDataExportScheduleResponseAggregationsFieldType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseAggregationsFieldTypeDimension, UsageDataExportScheduleResponseAggregationsFieldTypeMeasure:
		return true
	}
	return false
}

// Aggregation function
type UsageDataExportScheduleResponseAggregationsFunction string

const (
	UsageDataExportScheduleResponseAggregationsFunctionSum    UsageDataExportScheduleResponseAggregationsFunction = "SUM"
	UsageDataExportScheduleResponseAggregationsFunctionMin    UsageDataExportScheduleResponseAggregationsFunction = "MIN"
	UsageDataExportScheduleResponseAggregationsFunctionMax    UsageDataExportScheduleResponseAggregationsFunction = "MAX"
	UsageDataExportScheduleResponseAggregationsFunctionCount  UsageDataExportScheduleResponseAggregationsFunction = "COUNT"
	UsageDataExportScheduleResponseAggregationsFunctionLatest UsageDataExportScheduleResponseAggregationsFunction = "LATEST"
	UsageDataExportScheduleResponseAggregationsFunctionMean   UsageDataExportScheduleResponseAggregationsFunction = "MEAN"
	UsageDataExportScheduleResponseAggregationsFunctionUnique UsageDataExportScheduleResponseAggregationsFunction = "UNIQUE"
)

func (r UsageDataExportScheduleResponseAggregationsFunction) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseAggregationsFunctionSum, UsageDataExportScheduleResponseAggregationsFunctionMin, UsageDataExportScheduleResponseAggregationsFunctionMax, UsageDataExportScheduleResponseAggregationsFunctionCount, UsageDataExportScheduleResponseAggregationsFunctionLatest, UsageDataExportScheduleResponseAggregationsFunctionMean, UsageDataExportScheduleResponseAggregationsFunctionUnique:
		return true
	}
	return false
}

type UsageDataExportScheduleResponseDimensionFilter struct {
	// Field code
	FieldCode string `json:"fieldCode,required"`
	// Meter ID
	MeterID string `json:"meterId,required"`
	// Values to filter by
	Values []string                                           `json:"values,required"`
	JSON   usageDataExportScheduleResponseDimensionFilterJSON `json:"-"`
}

// usageDataExportScheduleResponseDimensionFilterJSON contains the JSON metadata
// for the struct [UsageDataExportScheduleResponseDimensionFilter]
type usageDataExportScheduleResponseDimensionFilterJSON struct {
	FieldCode   apijson.Field
	MeterID     apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageDataExportScheduleResponseDimensionFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDataExportScheduleResponseDimensionFilterJSON) RawJSON() string {
	return r.raw
}

// Group by a field
type UsageDataExportScheduleResponseGroup struct {
	// Field code to group by
	FieldCode string `json:"fieldCode"`
	// Frequency of usage data
	Frequency UsageDataExportScheduleResponseGroupsFrequency `json:"frequency"`
	GroupType UsageDataExportScheduleResponseGroupsGroupType `json:"groupType"`
	// Meter ID to group by
	MeterID string                                   `json:"meterId"`
	JSON    usageDataExportScheduleResponseGroupJSON `json:"-"`
	union   UsageDataExportScheduleResponseGroupsUnion
}

// usageDataExportScheduleResponseGroupJSON contains the JSON metadata for the
// struct [UsageDataExportScheduleResponseGroup]
type usageDataExportScheduleResponseGroupJSON struct {
	FieldCode   apijson.Field
	Frequency   apijson.Field
	GroupType   apijson.Field
	MeterID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r usageDataExportScheduleResponseGroupJSON) RawJSON() string {
	return r.raw
}

func (r *UsageDataExportScheduleResponseGroup) UnmarshalJSON(data []byte) (err error) {
	*r = UsageDataExportScheduleResponseGroup{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UsageDataExportScheduleResponseGroupsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup],
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup],
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup].
func (r UsageDataExportScheduleResponseGroup) AsUnion() UsageDataExportScheduleResponseGroupsUnion {
	return r.union
}

// Group by a field
//
// Union satisfied by
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup],
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup] or
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup].
type UsageDataExportScheduleResponseGroupsUnion interface {
	implementsUsageDataExportScheduleResponseGroup()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UsageDataExportScheduleResponseGroupsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup{}),
		},
	)
}

// Group by account
type UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup struct {
	GroupType UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupType `json:"groupType"`
	JSON      usageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupJSON      `json:"-"`
	DataExplorerAccountGroup
}

// usageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupJSON
// contains the JSON metadata for the struct
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup]
type usageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupJSON struct {
	GroupType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupJSON) RawJSON() string {
	return r.raw
}

func (r UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroup) implementsUsageDataExportScheduleResponseGroup() {
}

type UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupType string

const (
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount   UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupType = "ACCOUNT"
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupType = "DIMENSION"
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupTypeTime      UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupType = "TIME"
)

func (r UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount, UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension, UsageDataExportScheduleResponseGroupsDataExportsDataExplorerAccountGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by dimension
type UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup struct {
	GroupType UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupType `json:"groupType"`
	JSON      usageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupJSON      `json:"-"`
	DataExplorerDimensionGroup
}

// usageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupJSON
// contains the JSON metadata for the struct
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup]
type usageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupJSON struct {
	GroupType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupJSON) RawJSON() string {
	return r.raw
}

func (r UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroup) implementsUsageDataExportScheduleResponseGroup() {
}

type UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupType string

const (
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount   UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupType = "ACCOUNT"
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupType = "DIMENSION"
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime      UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupType = "TIME"
)

func (r UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount, UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension, UsageDataExportScheduleResponseGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by time
type UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup struct {
	GroupType UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupType `json:"groupType"`
	JSON      usageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupJSON      `json:"-"`
	DataExplorerTimeGroup
}

// usageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupJSON
// contains the JSON metadata for the struct
// [UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup]
type usageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupJSON struct {
	GroupType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupJSON) RawJSON() string {
	return r.raw
}

func (r UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroup) implementsUsageDataExportScheduleResponseGroup() {
}

type UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupType string

const (
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount   UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupType = "ACCOUNT"
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupType = "DIMENSION"
	UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupTypeTime      UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupType = "TIME"
)

func (r UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount, UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension, UsageDataExportScheduleResponseGroupsDataExportsDataExplorerTimeGroupGroupTypeTime:
		return true
	}
	return false
}

// Frequency of usage data
type UsageDataExportScheduleResponseGroupsFrequency string

const (
	UsageDataExportScheduleResponseGroupsFrequencyDay     UsageDataExportScheduleResponseGroupsFrequency = "DAY"
	UsageDataExportScheduleResponseGroupsFrequencyHour    UsageDataExportScheduleResponseGroupsFrequency = "HOUR"
	UsageDataExportScheduleResponseGroupsFrequencyWeek    UsageDataExportScheduleResponseGroupsFrequency = "WEEK"
	UsageDataExportScheduleResponseGroupsFrequencyMonth   UsageDataExportScheduleResponseGroupsFrequency = "MONTH"
	UsageDataExportScheduleResponseGroupsFrequencyQuarter UsageDataExportScheduleResponseGroupsFrequency = "QUARTER"
)

func (r UsageDataExportScheduleResponseGroupsFrequency) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseGroupsFrequencyDay, UsageDataExportScheduleResponseGroupsFrequencyHour, UsageDataExportScheduleResponseGroupsFrequencyWeek, UsageDataExportScheduleResponseGroupsFrequencyMonth, UsageDataExportScheduleResponseGroupsFrequencyQuarter:
		return true
	}
	return false
}

type UsageDataExportScheduleResponseGroupsGroupType string

const (
	UsageDataExportScheduleResponseGroupsGroupTypeAccount   UsageDataExportScheduleResponseGroupsGroupType = "ACCOUNT"
	UsageDataExportScheduleResponseGroupsGroupTypeDimension UsageDataExportScheduleResponseGroupsGroupType = "DIMENSION"
	UsageDataExportScheduleResponseGroupsGroupTypeTime      UsageDataExportScheduleResponseGroupsGroupType = "TIME"
)

func (r UsageDataExportScheduleResponseGroupsGroupType) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseGroupsGroupTypeAccount, UsageDataExportScheduleResponseGroupsGroupTypeDimension, UsageDataExportScheduleResponseGroupsGroupTypeTime:
		return true
	}
	return false
}

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type UsageDataExportScheduleResponseTimePeriod string

const (
	UsageDataExportScheduleResponseTimePeriodToday           UsageDataExportScheduleResponseTimePeriod = "TODAY"
	UsageDataExportScheduleResponseTimePeriodYesterday       UsageDataExportScheduleResponseTimePeriod = "YESTERDAY"
	UsageDataExportScheduleResponseTimePeriodWeekToDate      UsageDataExportScheduleResponseTimePeriod = "WEEK_TO_DATE"
	UsageDataExportScheduleResponseTimePeriodMonthToDate     UsageDataExportScheduleResponseTimePeriod = "MONTH_TO_DATE"
	UsageDataExportScheduleResponseTimePeriodYearToDate      UsageDataExportScheduleResponseTimePeriod = "YEAR_TO_DATE"
	UsageDataExportScheduleResponseTimePeriodPreviousWeek    UsageDataExportScheduleResponseTimePeriod = "PREVIOUS_WEEK"
	UsageDataExportScheduleResponseTimePeriodPreviousMonth   UsageDataExportScheduleResponseTimePeriod = "PREVIOUS_MONTH"
	UsageDataExportScheduleResponseTimePeriodPreviousQuarter UsageDataExportScheduleResponseTimePeriod = "PREVIOUS_QUARTER"
	UsageDataExportScheduleResponseTimePeriodPreviousYear    UsageDataExportScheduleResponseTimePeriod = "PREVIOUS_YEAR"
	UsageDataExportScheduleResponseTimePeriodLast12Hours     UsageDataExportScheduleResponseTimePeriod = "LAST_12_HOURS"
	UsageDataExportScheduleResponseTimePeriodLast7Days       UsageDataExportScheduleResponseTimePeriod = "LAST_7_DAYS"
	UsageDataExportScheduleResponseTimePeriodLast30Days      UsageDataExportScheduleResponseTimePeriod = "LAST_30_DAYS"
	UsageDataExportScheduleResponseTimePeriodLast35Days      UsageDataExportScheduleResponseTimePeriod = "LAST_35_DAYS"
	UsageDataExportScheduleResponseTimePeriodLast90Days      UsageDataExportScheduleResponseTimePeriod = "LAST_90_DAYS"
	UsageDataExportScheduleResponseTimePeriodLast120Days     UsageDataExportScheduleResponseTimePeriod = "LAST_120_DAYS"
	UsageDataExportScheduleResponseTimePeriodLastYear        UsageDataExportScheduleResponseTimePeriod = "LAST_YEAR"
)

func (r UsageDataExportScheduleResponseTimePeriod) IsKnown() bool {
	switch r {
	case UsageDataExportScheduleResponseTimePeriodToday, UsageDataExportScheduleResponseTimePeriodYesterday, UsageDataExportScheduleResponseTimePeriodWeekToDate, UsageDataExportScheduleResponseTimePeriodMonthToDate, UsageDataExportScheduleResponseTimePeriodYearToDate, UsageDataExportScheduleResponseTimePeriodPreviousWeek, UsageDataExportScheduleResponseTimePeriodPreviousMonth, UsageDataExportScheduleResponseTimePeriodPreviousQuarter, UsageDataExportScheduleResponseTimePeriodPreviousYear, UsageDataExportScheduleResponseTimePeriodLast12Hours, UsageDataExportScheduleResponseTimePeriodLast7Days, UsageDataExportScheduleResponseTimePeriodLast30Days, UsageDataExportScheduleResponseTimePeriodLast35Days, UsageDataExportScheduleResponseTimePeriodLast90Days, UsageDataExportScheduleResponseTimePeriodLast120Days, UsageDataExportScheduleResponseTimePeriodLastYear:
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
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseAggregation].
	Aggregations interface{} `json:"aggregations"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseDimensionFilter].
	DimensionFilters interface{} `json:"dimensionFilters"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseGroup].
	Groups interface{} `json:"groups"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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
	Aggregations         apijson.Field
	DimensionFilters     apijson.Field
	Groups               apijson.Field
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

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportScheduleNewResponseTimePeriod string

const (
	DataExportScheduleNewResponseTimePeriodToday           DataExportScheduleNewResponseTimePeriod = "TODAY"
	DataExportScheduleNewResponseTimePeriodYesterday       DataExportScheduleNewResponseTimePeriod = "YESTERDAY"
	DataExportScheduleNewResponseTimePeriodWeekToDate      DataExportScheduleNewResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleNewResponseTimePeriodMonthToDate     DataExportScheduleNewResponseTimePeriod = "MONTH_TO_DATE"
	DataExportScheduleNewResponseTimePeriodYearToDate      DataExportScheduleNewResponseTimePeriod = "YEAR_TO_DATE"
	DataExportScheduleNewResponseTimePeriodPreviousWeek    DataExportScheduleNewResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleNewResponseTimePeriodPreviousMonth   DataExportScheduleNewResponseTimePeriod = "PREVIOUS_MONTH"
	DataExportScheduleNewResponseTimePeriodPreviousQuarter DataExportScheduleNewResponseTimePeriod = "PREVIOUS_QUARTER"
	DataExportScheduleNewResponseTimePeriodPreviousYear    DataExportScheduleNewResponseTimePeriod = "PREVIOUS_YEAR"
	DataExportScheduleNewResponseTimePeriodLast12Hours     DataExportScheduleNewResponseTimePeriod = "LAST_12_HOURS"
	DataExportScheduleNewResponseTimePeriodLast7Days       DataExportScheduleNewResponseTimePeriod = "LAST_7_DAYS"
	DataExportScheduleNewResponseTimePeriodLast30Days      DataExportScheduleNewResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleNewResponseTimePeriodLast35Days      DataExportScheduleNewResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleNewResponseTimePeriodLast90Days      DataExportScheduleNewResponseTimePeriod = "LAST_90_DAYS"
	DataExportScheduleNewResponseTimePeriodLast120Days     DataExportScheduleNewResponseTimePeriod = "LAST_120_DAYS"
	DataExportScheduleNewResponseTimePeriodLastYear        DataExportScheduleNewResponseTimePeriod = "LAST_YEAR"
)

func (r DataExportScheduleNewResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleNewResponseTimePeriodToday, DataExportScheduleNewResponseTimePeriodYesterday, DataExportScheduleNewResponseTimePeriodWeekToDate, DataExportScheduleNewResponseTimePeriodMonthToDate, DataExportScheduleNewResponseTimePeriodYearToDate, DataExportScheduleNewResponseTimePeriodPreviousWeek, DataExportScheduleNewResponseTimePeriodPreviousMonth, DataExportScheduleNewResponseTimePeriodPreviousQuarter, DataExportScheduleNewResponseTimePeriodPreviousYear, DataExportScheduleNewResponseTimePeriodLast12Hours, DataExportScheduleNewResponseTimePeriodLast7Days, DataExportScheduleNewResponseTimePeriodLast30Days, DataExportScheduleNewResponseTimePeriodLast35Days, DataExportScheduleNewResponseTimePeriodLast90Days, DataExportScheduleNewResponseTimePeriodLast120Days, DataExportScheduleNewResponseTimePeriodLastYear:
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
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseAggregation].
	Aggregations interface{} `json:"aggregations"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseDimensionFilter].
	DimensionFilters interface{} `json:"dimensionFilters"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseGroup].
	Groups interface{} `json:"groups"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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
	Aggregations         apijson.Field
	DimensionFilters     apijson.Field
	Groups               apijson.Field
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

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportScheduleGetResponseTimePeriod string

const (
	DataExportScheduleGetResponseTimePeriodToday           DataExportScheduleGetResponseTimePeriod = "TODAY"
	DataExportScheduleGetResponseTimePeriodYesterday       DataExportScheduleGetResponseTimePeriod = "YESTERDAY"
	DataExportScheduleGetResponseTimePeriodWeekToDate      DataExportScheduleGetResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleGetResponseTimePeriodMonthToDate     DataExportScheduleGetResponseTimePeriod = "MONTH_TO_DATE"
	DataExportScheduleGetResponseTimePeriodYearToDate      DataExportScheduleGetResponseTimePeriod = "YEAR_TO_DATE"
	DataExportScheduleGetResponseTimePeriodPreviousWeek    DataExportScheduleGetResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleGetResponseTimePeriodPreviousMonth   DataExportScheduleGetResponseTimePeriod = "PREVIOUS_MONTH"
	DataExportScheduleGetResponseTimePeriodPreviousQuarter DataExportScheduleGetResponseTimePeriod = "PREVIOUS_QUARTER"
	DataExportScheduleGetResponseTimePeriodPreviousYear    DataExportScheduleGetResponseTimePeriod = "PREVIOUS_YEAR"
	DataExportScheduleGetResponseTimePeriodLast12Hours     DataExportScheduleGetResponseTimePeriod = "LAST_12_HOURS"
	DataExportScheduleGetResponseTimePeriodLast7Days       DataExportScheduleGetResponseTimePeriod = "LAST_7_DAYS"
	DataExportScheduleGetResponseTimePeriodLast30Days      DataExportScheduleGetResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleGetResponseTimePeriodLast35Days      DataExportScheduleGetResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleGetResponseTimePeriodLast90Days      DataExportScheduleGetResponseTimePeriod = "LAST_90_DAYS"
	DataExportScheduleGetResponseTimePeriodLast120Days     DataExportScheduleGetResponseTimePeriod = "LAST_120_DAYS"
	DataExportScheduleGetResponseTimePeriodLastYear        DataExportScheduleGetResponseTimePeriod = "LAST_YEAR"
)

func (r DataExportScheduleGetResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleGetResponseTimePeriodToday, DataExportScheduleGetResponseTimePeriodYesterday, DataExportScheduleGetResponseTimePeriodWeekToDate, DataExportScheduleGetResponseTimePeriodMonthToDate, DataExportScheduleGetResponseTimePeriodYearToDate, DataExportScheduleGetResponseTimePeriodPreviousWeek, DataExportScheduleGetResponseTimePeriodPreviousMonth, DataExportScheduleGetResponseTimePeriodPreviousQuarter, DataExportScheduleGetResponseTimePeriodPreviousYear, DataExportScheduleGetResponseTimePeriodLast12Hours, DataExportScheduleGetResponseTimePeriodLast7Days, DataExportScheduleGetResponseTimePeriodLast30Days, DataExportScheduleGetResponseTimePeriodLast35Days, DataExportScheduleGetResponseTimePeriodLast90Days, DataExportScheduleGetResponseTimePeriodLast120Days, DataExportScheduleGetResponseTimePeriodLastYear:
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
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseAggregation].
	Aggregations interface{} `json:"aggregations"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseDimensionFilter].
	DimensionFilters interface{} `json:"dimensionFilters"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseGroup].
	Groups interface{} `json:"groups"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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
	Aggregations         apijson.Field
	DimensionFilters     apijson.Field
	Groups               apijson.Field
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

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportScheduleUpdateResponseTimePeriod string

const (
	DataExportScheduleUpdateResponseTimePeriodToday           DataExportScheduleUpdateResponseTimePeriod = "TODAY"
	DataExportScheduleUpdateResponseTimePeriodYesterday       DataExportScheduleUpdateResponseTimePeriod = "YESTERDAY"
	DataExportScheduleUpdateResponseTimePeriodWeekToDate      DataExportScheduleUpdateResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleUpdateResponseTimePeriodMonthToDate     DataExportScheduleUpdateResponseTimePeriod = "MONTH_TO_DATE"
	DataExportScheduleUpdateResponseTimePeriodYearToDate      DataExportScheduleUpdateResponseTimePeriod = "YEAR_TO_DATE"
	DataExportScheduleUpdateResponseTimePeriodPreviousWeek    DataExportScheduleUpdateResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleUpdateResponseTimePeriodPreviousMonth   DataExportScheduleUpdateResponseTimePeriod = "PREVIOUS_MONTH"
	DataExportScheduleUpdateResponseTimePeriodPreviousQuarter DataExportScheduleUpdateResponseTimePeriod = "PREVIOUS_QUARTER"
	DataExportScheduleUpdateResponseTimePeriodPreviousYear    DataExportScheduleUpdateResponseTimePeriod = "PREVIOUS_YEAR"
	DataExportScheduleUpdateResponseTimePeriodLast12Hours     DataExportScheduleUpdateResponseTimePeriod = "LAST_12_HOURS"
	DataExportScheduleUpdateResponseTimePeriodLast7Days       DataExportScheduleUpdateResponseTimePeriod = "LAST_7_DAYS"
	DataExportScheduleUpdateResponseTimePeriodLast30Days      DataExportScheduleUpdateResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleUpdateResponseTimePeriodLast35Days      DataExportScheduleUpdateResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleUpdateResponseTimePeriodLast90Days      DataExportScheduleUpdateResponseTimePeriod = "LAST_90_DAYS"
	DataExportScheduleUpdateResponseTimePeriodLast120Days     DataExportScheduleUpdateResponseTimePeriod = "LAST_120_DAYS"
	DataExportScheduleUpdateResponseTimePeriodLastYear        DataExportScheduleUpdateResponseTimePeriod = "LAST_YEAR"
)

func (r DataExportScheduleUpdateResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateResponseTimePeriodToday, DataExportScheduleUpdateResponseTimePeriodYesterday, DataExportScheduleUpdateResponseTimePeriodWeekToDate, DataExportScheduleUpdateResponseTimePeriodMonthToDate, DataExportScheduleUpdateResponseTimePeriodYearToDate, DataExportScheduleUpdateResponseTimePeriodPreviousWeek, DataExportScheduleUpdateResponseTimePeriodPreviousMonth, DataExportScheduleUpdateResponseTimePeriodPreviousQuarter, DataExportScheduleUpdateResponseTimePeriodPreviousYear, DataExportScheduleUpdateResponseTimePeriodLast12Hours, DataExportScheduleUpdateResponseTimePeriodLast7Days, DataExportScheduleUpdateResponseTimePeriodLast30Days, DataExportScheduleUpdateResponseTimePeriodLast35Days, DataExportScheduleUpdateResponseTimePeriodLast90Days, DataExportScheduleUpdateResponseTimePeriodLast120Days, DataExportScheduleUpdateResponseTimePeriodLastYear:
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
	// Defines the Schedule frequency for the Data Export to run in Hours, Days, or
	// Minutes. Used in conjunction with the `scheduleType` parameter.
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
	DataExportScheduleListResponseScheduleTypeHour   DataExportScheduleListResponseScheduleType = "HOUR"
	DataExportScheduleListResponseScheduleTypeDay    DataExportScheduleListResponseScheduleType = "DAY"
	DataExportScheduleListResponseScheduleTypeMinute DataExportScheduleListResponseScheduleType = "MINUTE"
	DataExportScheduleListResponseScheduleTypeAdHoc  DataExportScheduleListResponseScheduleType = "AD_HOC"
)

func (r DataExportScheduleListResponseScheduleType) IsKnown() bool {
	switch r {
	case DataExportScheduleListResponseScheduleTypeHour, DataExportScheduleListResponseScheduleTypeDay, DataExportScheduleListResponseScheduleTypeMinute, DataExportScheduleListResponseScheduleTypeAdHoc:
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
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseAggregation].
	Aggregations interface{} `json:"aggregations"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseDimensionFilter].
	DimensionFilters interface{} `json:"dimensionFilters"`
	// This field can have the runtime type of
	// [[]UsageDataExportScheduleResponseGroup].
	Groups interface{} `json:"groups"`
	// This field can have the runtime type of [[]string].
	MeterIDs interface{} `json:"meterIds"`
	// This field can have the runtime type of
	// [[]OperationalDataExportScheduleResponseOperationalDataType].
	OperationalDataTypes interface{} `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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
	Aggregations         apijson.Field
	DimensionFilters     apijson.Field
	Groups               apijson.Field
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

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportScheduleDeleteResponseTimePeriod string

const (
	DataExportScheduleDeleteResponseTimePeriodToday           DataExportScheduleDeleteResponseTimePeriod = "TODAY"
	DataExportScheduleDeleteResponseTimePeriodYesterday       DataExportScheduleDeleteResponseTimePeriod = "YESTERDAY"
	DataExportScheduleDeleteResponseTimePeriodWeekToDate      DataExportScheduleDeleteResponseTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleDeleteResponseTimePeriodMonthToDate     DataExportScheduleDeleteResponseTimePeriod = "MONTH_TO_DATE"
	DataExportScheduleDeleteResponseTimePeriodYearToDate      DataExportScheduleDeleteResponseTimePeriod = "YEAR_TO_DATE"
	DataExportScheduleDeleteResponseTimePeriodPreviousWeek    DataExportScheduleDeleteResponseTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleDeleteResponseTimePeriodPreviousMonth   DataExportScheduleDeleteResponseTimePeriod = "PREVIOUS_MONTH"
	DataExportScheduleDeleteResponseTimePeriodPreviousQuarter DataExportScheduleDeleteResponseTimePeriod = "PREVIOUS_QUARTER"
	DataExportScheduleDeleteResponseTimePeriodPreviousYear    DataExportScheduleDeleteResponseTimePeriod = "PREVIOUS_YEAR"
	DataExportScheduleDeleteResponseTimePeriodLast12Hours     DataExportScheduleDeleteResponseTimePeriod = "LAST_12_HOURS"
	DataExportScheduleDeleteResponseTimePeriodLast7Days       DataExportScheduleDeleteResponseTimePeriod = "LAST_7_DAYS"
	DataExportScheduleDeleteResponseTimePeriodLast30Days      DataExportScheduleDeleteResponseTimePeriod = "LAST_30_DAYS"
	DataExportScheduleDeleteResponseTimePeriodLast35Days      DataExportScheduleDeleteResponseTimePeriod = "LAST_35_DAYS"
	DataExportScheduleDeleteResponseTimePeriodLast90Days      DataExportScheduleDeleteResponseTimePeriod = "LAST_90_DAYS"
	DataExportScheduleDeleteResponseTimePeriodLast120Days     DataExportScheduleDeleteResponseTimePeriod = "LAST_120_DAYS"
	DataExportScheduleDeleteResponseTimePeriodLastYear        DataExportScheduleDeleteResponseTimePeriod = "LAST_YEAR"
)

func (r DataExportScheduleDeleteResponseTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleDeleteResponseTimePeriodToday, DataExportScheduleDeleteResponseTimePeriodYesterday, DataExportScheduleDeleteResponseTimePeriodWeekToDate, DataExportScheduleDeleteResponseTimePeriodMonthToDate, DataExportScheduleDeleteResponseTimePeriodYearToDate, DataExportScheduleDeleteResponseTimePeriodPreviousWeek, DataExportScheduleDeleteResponseTimePeriodPreviousMonth, DataExportScheduleDeleteResponseTimePeriodPreviousQuarter, DataExportScheduleDeleteResponseTimePeriodPreviousYear, DataExportScheduleDeleteResponseTimePeriodLast12Hours, DataExportScheduleDeleteResponseTimePeriodLast7Days, DataExportScheduleDeleteResponseTimePeriodLast30Days, DataExportScheduleDeleteResponseTimePeriodLast35Days, DataExportScheduleDeleteResponseTimePeriodLast90Days, DataExportScheduleDeleteResponseTimePeriodLast120Days, DataExportScheduleDeleteResponseTimePeriodLastYear:
		return true
	}
	return false
}

type DataExportScheduleNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Request representing an operational schedule configuration.
	Body DataExportScheduleNewParamsBodyUnion `json:"body,required"`
}

func (r DataExportScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request representing an operational schedule configuration.
type DataExportScheduleNewParamsBody struct {
	// The type of data to export. Possible values are: OPERATIONAL
	SourceType           param.Field[DataExportScheduleNewParamsBodySourceType] `json:"sourceType,required"`
	AccountIDs           param.Field[interface{}]                               `json:"accountIds"`
	Aggregations         param.Field[interface{}]                               `json:"aggregations"`
	DimensionFilters     param.Field[interface{}]                               `json:"dimensionFilters"`
	Groups               param.Field[interface{}]                               `json:"groups"`
	MeterIDs             param.Field[interface{}]                               `json:"meterIds"`
	OperationalDataTypes param.Field[interface{}]                               `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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

// The type of data to export. Possible values are: OPERATIONAL
type DataExportScheduleNewParamsBodySourceType string

const (
	DataExportScheduleNewParamsBodySourceTypeOperational DataExportScheduleNewParamsBodySourceType = "OPERATIONAL"
	DataExportScheduleNewParamsBodySourceTypeUsage       DataExportScheduleNewParamsBodySourceType = "USAGE"
)

func (r DataExportScheduleNewParamsBodySourceType) IsKnown() bool {
	switch r {
	case DataExportScheduleNewParamsBodySourceTypeOperational, DataExportScheduleNewParamsBodySourceTypeUsage:
		return true
	}
	return false
}

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportScheduleNewParamsBodyTimePeriod string

const (
	DataExportScheduleNewParamsBodyTimePeriodToday           DataExportScheduleNewParamsBodyTimePeriod = "TODAY"
	DataExportScheduleNewParamsBodyTimePeriodYesterday       DataExportScheduleNewParamsBodyTimePeriod = "YESTERDAY"
	DataExportScheduleNewParamsBodyTimePeriodWeekToDate      DataExportScheduleNewParamsBodyTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleNewParamsBodyTimePeriodMonthToDate     DataExportScheduleNewParamsBodyTimePeriod = "MONTH_TO_DATE"
	DataExportScheduleNewParamsBodyTimePeriodYearToDate      DataExportScheduleNewParamsBodyTimePeriod = "YEAR_TO_DATE"
	DataExportScheduleNewParamsBodyTimePeriodPreviousWeek    DataExportScheduleNewParamsBodyTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleNewParamsBodyTimePeriodPreviousMonth   DataExportScheduleNewParamsBodyTimePeriod = "PREVIOUS_MONTH"
	DataExportScheduleNewParamsBodyTimePeriodPreviousQuarter DataExportScheduleNewParamsBodyTimePeriod = "PREVIOUS_QUARTER"
	DataExportScheduleNewParamsBodyTimePeriodPreviousYear    DataExportScheduleNewParamsBodyTimePeriod = "PREVIOUS_YEAR"
	DataExportScheduleNewParamsBodyTimePeriodLast12Hours     DataExportScheduleNewParamsBodyTimePeriod = "LAST_12_HOURS"
	DataExportScheduleNewParamsBodyTimePeriodLast7Days       DataExportScheduleNewParamsBodyTimePeriod = "LAST_7_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodLast30Days      DataExportScheduleNewParamsBodyTimePeriod = "LAST_30_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodLast35Days      DataExportScheduleNewParamsBodyTimePeriod = "LAST_35_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodLast90Days      DataExportScheduleNewParamsBodyTimePeriod = "LAST_90_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodLast120Days     DataExportScheduleNewParamsBodyTimePeriod = "LAST_120_DAYS"
	DataExportScheduleNewParamsBodyTimePeriodLastYear        DataExportScheduleNewParamsBodyTimePeriod = "LAST_YEAR"
)

func (r DataExportScheduleNewParamsBodyTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleNewParamsBodyTimePeriodToday, DataExportScheduleNewParamsBodyTimePeriodYesterday, DataExportScheduleNewParamsBodyTimePeriodWeekToDate, DataExportScheduleNewParamsBodyTimePeriodMonthToDate, DataExportScheduleNewParamsBodyTimePeriodYearToDate, DataExportScheduleNewParamsBodyTimePeriodPreviousWeek, DataExportScheduleNewParamsBodyTimePeriodPreviousMonth, DataExportScheduleNewParamsBodyTimePeriodPreviousQuarter, DataExportScheduleNewParamsBodyTimePeriodPreviousYear, DataExportScheduleNewParamsBodyTimePeriodLast12Hours, DataExportScheduleNewParamsBodyTimePeriodLast7Days, DataExportScheduleNewParamsBodyTimePeriodLast30Days, DataExportScheduleNewParamsBodyTimePeriodLast35Days, DataExportScheduleNewParamsBodyTimePeriodLast90Days, DataExportScheduleNewParamsBodyTimePeriodLast120Days, DataExportScheduleNewParamsBodyTimePeriodLastYear:
		return true
	}
	return false
}

type DataExportScheduleGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type DataExportScheduleUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Request representing an operational schedule configuration.
	Body DataExportScheduleUpdateParamsBodyUnion `json:"body,required"`
}

func (r DataExportScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request representing an operational schedule configuration.
type DataExportScheduleUpdateParamsBody struct {
	// The type of data to export. Possible values are: OPERATIONAL
	SourceType           param.Field[DataExportScheduleUpdateParamsBodySourceType] `json:"sourceType,required"`
	AccountIDs           param.Field[interface{}]                                  `json:"accountIds"`
	Aggregations         param.Field[interface{}]                                  `json:"aggregations"`
	DimensionFilters     param.Field[interface{}]                                  `json:"dimensionFilters"`
	Groups               param.Field[interface{}]                                  `json:"groups"`
	MeterIDs             param.Field[interface{}]                                  `json:"meterIds"`
	OperationalDataTypes param.Field[interface{}]                                  `json:"operationalDataTypes"`
	// Define a time period to control the range of usage data you want the data export
	// to contain when it runs:
	//
	//   - **TODAY**. Data collected for the current day up until the time the export is
	//     scheduled to run.
	//   - **YESTERDAY**. Data collected for the day before the export runs under the
	//     schedule - that is, the 24 hour period from midnight to midnight of the day
	//     before.
	//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
	//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
	//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
	//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
	//     export will contain data for the period running from Monday, June 3rd 2024 to
	//     midnight on Sunday, June 9th 2024.
	//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
	//     the period covering the current week, month, or year period. For example if
	//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
	//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
	//     all usage data collected starting Monday October 14th 2024 through to the
	//     Wednesday at 10 a.m. UTC of the current week.
	//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
	//     of the hour in which the export is scheduled to run.
	//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
	//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
	//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
	//     and the export is scheduled to run for any time on June 15th 2024, it will
	//     contain usage data collected for the previous 30 days - starting May 16th 2024
	//     through to midnight on June 14th 2024
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

// The type of data to export. Possible values are: OPERATIONAL
type DataExportScheduleUpdateParamsBodySourceType string

const (
	DataExportScheduleUpdateParamsBodySourceTypeOperational DataExportScheduleUpdateParamsBodySourceType = "OPERATIONAL"
	DataExportScheduleUpdateParamsBodySourceTypeUsage       DataExportScheduleUpdateParamsBodySourceType = "USAGE"
)

func (r DataExportScheduleUpdateParamsBodySourceType) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateParamsBodySourceTypeOperational, DataExportScheduleUpdateParamsBodySourceTypeUsage:
		return true
	}
	return false
}

// Define a time period to control the range of usage data you want the data export
// to contain when it runs:
//
//   - **TODAY**. Data collected for the current day up until the time the export is
//     scheduled to run.
//   - **YESTERDAY**. Data collected for the day before the export runs under the
//     schedule - that is, the 24 hour period from midnight to midnight of the day
//     before.
//   - **PREVIOUS_WEEK**, **PREVIOUS_MONTH**, **PREVIOUS_QUARTER**,
//     **PREVIOUS_YEAR**. Data collected for the previous full week, month, quarter,
//     or year period. For example if **PREVIOUS_WEEK**, weeks run Monday to Monday -
//     if the export is scheduled to run on June 12th 2024, which is a Wednesday, the
//     export will contain data for the period running from Monday, June 3rd 2024 to
//     midnight on Sunday, June 9th 2024.
//   - **WEEK_TO_DATE**, **MONTH_TO_DATE**, or **YEAR_TO_DATE**. Data collected for
//     the period covering the current week, month, or year period. For example if
//     **WEEK_TO_DATE**, weeks run Monday to Monday - if the Export is scheduled to
//     run at 10 a.m. UTC on October 16th 2024, which is a Wednesday, it will contain
//     all usage data collected starting Monday October 14th 2024 through to the
//     Wednesday at 10 a.m. UTC of the current week.
//   - **LAST_12_HOURS**. Data collected for the twelve hour period up to the start
//     of the hour in which the export is scheduled to run.
//   - **LAST_7_DAYS**, **LAST_30_DAYS**, **LAST_35_DAYS**, **LAST_90_DAYS**,
//     **LAST_120_DAYS** **LAST_YEAR**. Data collected for the selected period prior
//     to the date the export is scheduled to run. For example if **LAST_30_DAYS**
//     and the export is scheduled to run for any time on June 15th 2024, it will
//     contain usage data collected for the previous 30 days - starting May 16th 2024
//     through to midnight on June 14th 2024
//
// For more details and examples, see the
// [Time Period](https://www.m3ter.com/docs/guides/data-exports/creating-export-schedules#time-period)
// section in our main User Documentation.
type DataExportScheduleUpdateParamsBodyTimePeriod string

const (
	DataExportScheduleUpdateParamsBodyTimePeriodToday           DataExportScheduleUpdateParamsBodyTimePeriod = "TODAY"
	DataExportScheduleUpdateParamsBodyTimePeriodYesterday       DataExportScheduleUpdateParamsBodyTimePeriod = "YESTERDAY"
	DataExportScheduleUpdateParamsBodyTimePeriodWeekToDate      DataExportScheduleUpdateParamsBodyTimePeriod = "WEEK_TO_DATE"
	DataExportScheduleUpdateParamsBodyTimePeriodMonthToDate     DataExportScheduleUpdateParamsBodyTimePeriod = "MONTH_TO_DATE"
	DataExportScheduleUpdateParamsBodyTimePeriodYearToDate      DataExportScheduleUpdateParamsBodyTimePeriod = "YEAR_TO_DATE"
	DataExportScheduleUpdateParamsBodyTimePeriodPreviousWeek    DataExportScheduleUpdateParamsBodyTimePeriod = "PREVIOUS_WEEK"
	DataExportScheduleUpdateParamsBodyTimePeriodPreviousMonth   DataExportScheduleUpdateParamsBodyTimePeriod = "PREVIOUS_MONTH"
	DataExportScheduleUpdateParamsBodyTimePeriodPreviousQuarter DataExportScheduleUpdateParamsBodyTimePeriod = "PREVIOUS_QUARTER"
	DataExportScheduleUpdateParamsBodyTimePeriodPreviousYear    DataExportScheduleUpdateParamsBodyTimePeriod = "PREVIOUS_YEAR"
	DataExportScheduleUpdateParamsBodyTimePeriodLast12Hours     DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_12_HOURS"
	DataExportScheduleUpdateParamsBodyTimePeriodLast7Days       DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_7_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodLast30Days      DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_30_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodLast35Days      DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_35_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodLast90Days      DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_90_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodLast120Days     DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_120_DAYS"
	DataExportScheduleUpdateParamsBodyTimePeriodLastYear        DataExportScheduleUpdateParamsBodyTimePeriod = "LAST_YEAR"
)

func (r DataExportScheduleUpdateParamsBodyTimePeriod) IsKnown() bool {
	switch r {
	case DataExportScheduleUpdateParamsBodyTimePeriodToday, DataExportScheduleUpdateParamsBodyTimePeriodYesterday, DataExportScheduleUpdateParamsBodyTimePeriodWeekToDate, DataExportScheduleUpdateParamsBodyTimePeriodMonthToDate, DataExportScheduleUpdateParamsBodyTimePeriodYearToDate, DataExportScheduleUpdateParamsBodyTimePeriodPreviousWeek, DataExportScheduleUpdateParamsBodyTimePeriodPreviousMonth, DataExportScheduleUpdateParamsBodyTimePeriodPreviousQuarter, DataExportScheduleUpdateParamsBodyTimePeriodPreviousYear, DataExportScheduleUpdateParamsBodyTimePeriodLast12Hours, DataExportScheduleUpdateParamsBodyTimePeriodLast7Days, DataExportScheduleUpdateParamsBodyTimePeriodLast30Days, DataExportScheduleUpdateParamsBodyTimePeriodLast35Days, DataExportScheduleUpdateParamsBodyTimePeriodLast90Days, DataExportScheduleUpdateParamsBodyTimePeriodLast120Days, DataExportScheduleUpdateParamsBodyTimePeriodLastYear:
		return true
	}
	return false
}

type DataExportScheduleListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
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
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
