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
//   - Select the Meters and Accounts whose usage data you want to include in the
//     ad-hoc export.
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
// **Date Range for Operational Data Exports**. To restrict the operational data
// included in the export by a date/time range, use the `startDate` and `endDate`
// date/time request parameters to specify the period. Constraints:
//
// - `startDate` with no `endDate` is valid.
// - `endDate` with no `startDate` is valid.
// - If both are set,`startDate` must be before `endDate`.
// - `endDate` must be before now UTC.
//
// **Date Range for Usage Data Exports**. To restrict the usage data included in
// the export by date/time range, use the `startDate` and `endDate` date/time
// parameters:
//
//   - Both `startDate` and `endDate` are required.
//   - `endDate` must be after `startDate`.
//   - `endDate` cannot be after tomorrow at midnight UTC. For example if today is
//     May 20th 2025, you can only choose `endDate` to be equal or before
//     2025-05-21T00:00:00.000Z.
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
	// The type of data to export. Possible values are: OPERATIONAL
	SourceType param.Field[AdHocOperationalDataRequestSourceType] `json:"sourceType,required"`
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

// The type of data to export. Possible values are: OPERATIONAL
type AdHocOperationalDataRequestSourceType string

const (
	AdHocOperationalDataRequestSourceTypeOperational AdHocOperationalDataRequestSourceType = "OPERATIONAL"
)

func (r AdHocOperationalDataRequestSourceType) IsKnown() bool {
	switch r {
	case AdHocOperationalDataRequestSourceTypeOperational:
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
	// The type of data to export. Possible values are: USAGE
	SourceType param.Field[AdHocUsageDataRequestSourceType] `json:"sourceType,required"`
	// List of account IDs for which the usage data will be exported.
	AccountIDs param.Field[[]string] `json:"accountIds"`
	// List of aggregations to apply
	Aggregations param.Field[[]AdHocUsageDataRequestAggregationParam] `json:"aggregations"`
	// List of dimension filters to apply
	DimensionFilters param.Field[[]AdHocUsageDataRequestDimensionFilterParam] `json:"dimensionFilters"`
	// List of groups to apply
	Groups param.Field[[]AdHocUsageDataRequestGroupsUnionParam] `json:"groups"`
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

func (r AdHocUsageDataRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdHocUsageDataRequestParam) implementsDataExportNewAdhocParamsBodyUnion() {}

// The type of data to export. Possible values are: USAGE
type AdHocUsageDataRequestSourceType string

const (
	AdHocUsageDataRequestSourceTypeUsage AdHocUsageDataRequestSourceType = "USAGE"
)

func (r AdHocUsageDataRequestSourceType) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestSourceTypeUsage:
		return true
	}
	return false
}

type AdHocUsageDataRequestAggregationParam struct {
	// Field code
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Type of field
	FieldType param.Field[AdHocUsageDataRequestAggregationsFieldType] `json:"fieldType,required"`
	// Aggregation function
	Function param.Field[AdHocUsageDataRequestAggregationsFunction] `json:"function,required"`
	// Meter ID
	MeterID param.Field[string] `json:"meterId,required"`
}

func (r AdHocUsageDataRequestAggregationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of field
type AdHocUsageDataRequestAggregationsFieldType string

const (
	AdHocUsageDataRequestAggregationsFieldTypeDimension AdHocUsageDataRequestAggregationsFieldType = "DIMENSION"
	AdHocUsageDataRequestAggregationsFieldTypeMeasure   AdHocUsageDataRequestAggregationsFieldType = "MEASURE"
)

func (r AdHocUsageDataRequestAggregationsFieldType) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestAggregationsFieldTypeDimension, AdHocUsageDataRequestAggregationsFieldTypeMeasure:
		return true
	}
	return false
}

// Aggregation function
type AdHocUsageDataRequestAggregationsFunction string

const (
	AdHocUsageDataRequestAggregationsFunctionSum    AdHocUsageDataRequestAggregationsFunction = "SUM"
	AdHocUsageDataRequestAggregationsFunctionMin    AdHocUsageDataRequestAggregationsFunction = "MIN"
	AdHocUsageDataRequestAggregationsFunctionMax    AdHocUsageDataRequestAggregationsFunction = "MAX"
	AdHocUsageDataRequestAggregationsFunctionCount  AdHocUsageDataRequestAggregationsFunction = "COUNT"
	AdHocUsageDataRequestAggregationsFunctionLatest AdHocUsageDataRequestAggregationsFunction = "LATEST"
	AdHocUsageDataRequestAggregationsFunctionMean   AdHocUsageDataRequestAggregationsFunction = "MEAN"
	AdHocUsageDataRequestAggregationsFunctionUnique AdHocUsageDataRequestAggregationsFunction = "UNIQUE"
)

func (r AdHocUsageDataRequestAggregationsFunction) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestAggregationsFunctionSum, AdHocUsageDataRequestAggregationsFunctionMin, AdHocUsageDataRequestAggregationsFunctionMax, AdHocUsageDataRequestAggregationsFunctionCount, AdHocUsageDataRequestAggregationsFunctionLatest, AdHocUsageDataRequestAggregationsFunctionMean, AdHocUsageDataRequestAggregationsFunctionUnique:
		return true
	}
	return false
}

type AdHocUsageDataRequestDimensionFilterParam struct {
	// Field code
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Meter ID
	MeterID param.Field[string] `json:"meterId,required"`
	// Values to filter by
	Values param.Field[[]string] `json:"values,required"`
}

func (r AdHocUsageDataRequestDimensionFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Group by a field
//
// Satisfied by
// [AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupParam],
// [AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupParam],
// [AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupParam].
type AdHocUsageDataRequestGroupsUnionParam interface {
	implementsAdHocUsageDataRequestGroupsUnionParam()
}

// Group by account
type AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupParam struct {
	GroupType param.Field[AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupType] `json:"groupType"`
	DataExplorerAccountGroupParam
}

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupParam) implementsAdHocUsageDataRequestGroupsUnionParam() {
}

type AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupType string

const (
	AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount   AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupType = "ACCOUNT"
	AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupType = "DIMENSION"
	AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeTime      AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupType = "TIME"
)

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupType) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount, AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension, AdHocUsageDataRequestGroupsDataExportsDataExplorerAccountGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by dimension
type AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupParam struct {
	GroupType param.Field[AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupType] `json:"groupType"`
	DataExplorerDimensionGroupParam
}

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupParam) implementsAdHocUsageDataRequestGroupsUnionParam() {
}

type AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupType string

const (
	AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount   AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupType = "ACCOUNT"
	AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupType = "DIMENSION"
	AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime      AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupType = "TIME"
)

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupType) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount, AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension, AdHocUsageDataRequestGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by time
type AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupParam struct {
	GroupType param.Field[AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupType] `json:"groupType"`
	DataExplorerTimeGroupParam
}

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupParam) implementsAdHocUsageDataRequestGroupsUnionParam() {
}

type AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupType string

const (
	AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount   AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupType = "ACCOUNT"
	AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupType = "DIMENSION"
	AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeTime      AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupType = "TIME"
)

func (r AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupType) IsKnown() bool {
	switch r {
	case AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount, AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension, AdHocUsageDataRequestGroupsDataExportsDataExplorerTimeGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by account
type DataExplorerAccountGroup struct {
	GroupType DataExplorerAccountGroupGroupType `json:"groupType"`
	JSON      dataExplorerAccountGroupJSON      `json:"-"`
}

// dataExplorerAccountGroupJSON contains the JSON metadata for the struct
// [DataExplorerAccountGroup]
type dataExplorerAccountGroupJSON struct {
	GroupType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataExplorerAccountGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExplorerAccountGroupJSON) RawJSON() string {
	return r.raw
}

type DataExplorerAccountGroupGroupType string

const (
	DataExplorerAccountGroupGroupTypeAccount   DataExplorerAccountGroupGroupType = "ACCOUNT"
	DataExplorerAccountGroupGroupTypeDimension DataExplorerAccountGroupGroupType = "DIMENSION"
	DataExplorerAccountGroupGroupTypeTime      DataExplorerAccountGroupGroupType = "TIME"
)

func (r DataExplorerAccountGroupGroupType) IsKnown() bool {
	switch r {
	case DataExplorerAccountGroupGroupTypeAccount, DataExplorerAccountGroupGroupTypeDimension, DataExplorerAccountGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by account
type DataExplorerAccountGroupParam struct {
	GroupType param.Field[DataExplorerAccountGroupGroupType] `json:"groupType"`
}

func (r DataExplorerAccountGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Group by dimension
type DataExplorerDimensionGroup struct {
	// Field code to group by
	FieldCode string `json:"fieldCode,required"`
	// Meter ID to group by
	MeterID   string                              `json:"meterId,required"`
	GroupType DataExplorerDimensionGroupGroupType `json:"groupType"`
	JSON      dataExplorerDimensionGroupJSON      `json:"-"`
}

// dataExplorerDimensionGroupJSON contains the JSON metadata for the struct
// [DataExplorerDimensionGroup]
type dataExplorerDimensionGroupJSON struct {
	FieldCode   apijson.Field
	MeterID     apijson.Field
	GroupType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataExplorerDimensionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExplorerDimensionGroupJSON) RawJSON() string {
	return r.raw
}

type DataExplorerDimensionGroupGroupType string

const (
	DataExplorerDimensionGroupGroupTypeAccount   DataExplorerDimensionGroupGroupType = "ACCOUNT"
	DataExplorerDimensionGroupGroupTypeDimension DataExplorerDimensionGroupGroupType = "DIMENSION"
	DataExplorerDimensionGroupGroupTypeTime      DataExplorerDimensionGroupGroupType = "TIME"
)

func (r DataExplorerDimensionGroupGroupType) IsKnown() bool {
	switch r {
	case DataExplorerDimensionGroupGroupTypeAccount, DataExplorerDimensionGroupGroupTypeDimension, DataExplorerDimensionGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by dimension
type DataExplorerDimensionGroupParam struct {
	// Field code to group by
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Meter ID to group by
	MeterID   param.Field[string]                              `json:"meterId,required"`
	GroupType param.Field[DataExplorerDimensionGroupGroupType] `json:"groupType"`
}

func (r DataExplorerDimensionGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Group by time
type DataExplorerTimeGroup struct {
	// Frequency of usage data
	Frequency DataExplorerTimeGroupFrequency `json:"frequency,required"`
	GroupType DataExplorerTimeGroupGroupType `json:"groupType"`
	JSON      dataExplorerTimeGroupJSON      `json:"-"`
}

// dataExplorerTimeGroupJSON contains the JSON metadata for the struct
// [DataExplorerTimeGroup]
type dataExplorerTimeGroupJSON struct {
	Frequency   apijson.Field
	GroupType   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataExplorerTimeGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExplorerTimeGroupJSON) RawJSON() string {
	return r.raw
}

// Frequency of usage data
type DataExplorerTimeGroupFrequency string

const (
	DataExplorerTimeGroupFrequencyDay     DataExplorerTimeGroupFrequency = "DAY"
	DataExplorerTimeGroupFrequencyHour    DataExplorerTimeGroupFrequency = "HOUR"
	DataExplorerTimeGroupFrequencyWeek    DataExplorerTimeGroupFrequency = "WEEK"
	DataExplorerTimeGroupFrequencyMonth   DataExplorerTimeGroupFrequency = "MONTH"
	DataExplorerTimeGroupFrequencyQuarter DataExplorerTimeGroupFrequency = "QUARTER"
)

func (r DataExplorerTimeGroupFrequency) IsKnown() bool {
	switch r {
	case DataExplorerTimeGroupFrequencyDay, DataExplorerTimeGroupFrequencyHour, DataExplorerTimeGroupFrequencyWeek, DataExplorerTimeGroupFrequencyMonth, DataExplorerTimeGroupFrequencyQuarter:
		return true
	}
	return false
}

type DataExplorerTimeGroupGroupType string

const (
	DataExplorerTimeGroupGroupTypeAccount   DataExplorerTimeGroupGroupType = "ACCOUNT"
	DataExplorerTimeGroupGroupTypeDimension DataExplorerTimeGroupGroupType = "DIMENSION"
	DataExplorerTimeGroupGroupTypeTime      DataExplorerTimeGroupGroupType = "TIME"
)

func (r DataExplorerTimeGroupGroupType) IsKnown() bool {
	switch r {
	case DataExplorerTimeGroupGroupTypeAccount, DataExplorerTimeGroupGroupTypeDimension, DataExplorerTimeGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by time
type DataExplorerTimeGroupParam struct {
	// Frequency of usage data
	Frequency param.Field[DataExplorerTimeGroupFrequency] `json:"frequency,required"`
	GroupType param.Field[DataExplorerTimeGroupGroupType] `json:"groupType"`
}

func (r DataExplorerTimeGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// The type of data to export. Possible values are: OPERATIONAL
	SourceType           param.Field[DataExportNewAdhocParamsBodySourceType] `json:"sourceType,required"`
	AccountIDs           param.Field[interface{}]                            `json:"accountIds"`
	Aggregations         param.Field[interface{}]                            `json:"aggregations"`
	DimensionFilters     param.Field[interface{}]                            `json:"dimensionFilters"`
	Groups               param.Field[interface{}]                            `json:"groups"`
	MeterIDs             param.Field[interface{}]                            `json:"meterIds"`
	OperationalDataTypes param.Field[interface{}]                            `json:"operationalDataTypes"`
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

// The type of data to export. Possible values are: OPERATIONAL
type DataExportNewAdhocParamsBodySourceType string

const (
	DataExportNewAdhocParamsBodySourceTypeOperational DataExportNewAdhocParamsBodySourceType = "OPERATIONAL"
	DataExportNewAdhocParamsBodySourceTypeUsage       DataExportNewAdhocParamsBodySourceType = "USAGE"
)

func (r DataExportNewAdhocParamsBodySourceType) IsKnown() bool {
	switch r {
	case DataExportNewAdhocParamsBodySourceTypeOperational, DataExportNewAdhocParamsBodySourceTypeUsage:
		return true
	}
	return false
}
