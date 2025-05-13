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
)

// UsageService contains methods and other services that help with interacting with
// the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options     []option.RequestOption
	FileUploads *UsageFileUploadService
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r *UsageService) {
	r = &UsageService{}
	r.Options = opts
	r.FileUploads = NewUsageFileUploadService(opts...)
	return
}

// Returns a presigned download URL for failed ingest file download based on the
// file path provided.
//
// If a usage data ingest measurement you submit to the m3ter platform fails, an
// `ingest.validation.failure` Event is generated. Use this call to obtain a
// download URL which you can then use to download a file containing details of
// what went wrong with the attempted usage data measurement ingest, and allowing
// you to follow-up and resolve the issue.
//
// To obtain the `file` query parameter:
//
//   - Use the
//     [List Events](https://www.m3ter.com/docs/api#tag/Events/operation/ListEventFields)
//     call with the `ingest.validation.failure` for the `eventName` query parameter.
//   - The response contains a `getDownloadUrl` response parameter and this contains
//     the file path you can use to obtain the failed ingest file download URL.
//
// **Notes:**
//
//   - The presigned Url returned to use for failed ingest file download is
//     time-bound and expires after 5 minutes.
//   - If you make a List Events call for `ingest.validation.failure` Events in your
//     Organization, then you can perform this **GET** call using the full URL
//     returned for any ingest failure Event to obtain a failed ingest file download
//     URL for the Event.
func (r *UsageService) GetFailedIngestDownloadURL(ctx context.Context, params UsageGetFailedIngestDownloadURLParams, opts ...option.RequestOption) (res *DownloadURLResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/measurements/failedIngest/getDownloadUrl", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Query and filter usage data collected for your Organization.
//
// You can use several parameters to filter the range of usage data returned:
//
//   - **Time period.** Use `startDate` and `endDate` to define a period. The query
//     references the `timestamp` values of usage data submissions for applying the
//     defined time period, and not the time submissions were `receivedAt` by the
//     platform. Only usage data with a `timestamp` that falls in the defined time
//     period are returned.(Required)
//   - **Meters.** Specify the Meters you want the query to return data for.
//   - **Accounts.** Specify the Accounts you want the query to return data for.
//   - **Dimension Filters.** Specify values for Dimension data fields on included
//     Meters. Only data that match the specified Dimension field values will be
//     returned for the query.
//
// You can apply Aggregations functions to the usage data returned for the query.
// If you apply Aggregations, you can select to group the data by:
//
// - **Account**
// - **Time**
// - **Dimension**
func (r *UsageService) Query(ctx context.Context, params UsageQueryParams, opts ...option.RequestOption) (res *UsageQueryResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/usage/query", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Submit a measurement or multiple measurements to the m3ter platform. The maximum
// size of the payload needs to be less than 512,000 bytes.
//
// **NOTES:**
//
//   - **Non-existent Accounts.** The `account` request parameter is required.
//     However, if you want to submit a usage data measurement for an Account which
//     does not yet exist in your Organization, you can use an `account` code for a
//     non-existent Account. A new skeleton Account will be automatically created.
//     The usage data measurement is accepted and ingested as data belonging to the
//     new auto-created Account. At a later date, you can edit the Account's
//     Code,??Name, and??e-mail address. For more details, see
//     [Submitting Usage Data for Non-Existent Accounts](https://www.m3ter.com/docs/guides/billing-and-usage-data/submitting-usage-data/submitting-usage-data-for-non-existent-accounts)
//     in our main documentation.
//   - **Usage Data Adjustments.** If you need to make corrections for billing
//     retrospectively against an Account, you can use date/time values in the past
//     for the `ts` (timestamp) request parameter to submit positive or negative
//     usage data amounts to correct and reconcile earlier billing anomalies. For
//     more details, see
//     [Submitting Usage Data Adjustments Using Timestamp](https://www.m3ter.com/docs/guides/billing-and-usage-data/submitting-usage-data/submitting-usage-data-adjustments-using-timestamp)
//     in our main documentation.
//   - **Ingest Validation Failure Events.** After the intial submission of a usage
//     data measurement to the Ingest API, a data enrichment stage is performed to
//     check for any errors in the usage data measurement, such as a missing field.
//     If an error is identified, this might result in the submission being rejected.
//     In these cases, an _ingest validation failure_ Event is generated, which you
//     can review on the
//     [Ingest Events](https://www.m3ter.com/docs/guides/billing-and-usage-data/submitting-usage-data/reviewing-and-resolving-ingest-events)
//     page in the Console. See also the
//     [Events](https://www.m3ter.com/docs/api#tag/Events) section in this API
//     Reference.
//
// **IMPORTANT! - Use of PII:** The use of any of your end-customers' Personally
// Identifiable Information (PII) in m3ter is restricted to a few fields on the
// **Account** entity. Please ensure that any measurements you submit do not
// contain any end-customer PII data. See the
// [Introduction section](https://www.m3ter.com/docs/api#section/Introduction)
// above for more details.
func (r *UsageService) Submit(ctx context.Context, params UsageSubmitParams, opts ...option.RequestOption) (res *SubmitMeasurementsResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://ingest.m3ter.com/")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/measurements", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// It contains details for downloading a file
type DownloadURLResponse struct {
	// The presigned download URL
	URL  string                  `json:"url"`
	JSON downloadURLResponseJSON `json:"-"`
}

// downloadURLResponseJSON contains the JSON metadata for the struct
// [DownloadURLResponse]
type downloadURLResponseJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadURLResponseJSON) RawJSON() string {
	return r.raw
}

type MeasurementRequestParam struct {
	// Code of the Account the measurement is for.
	Account param.Field[string] `json:"account,required"`
	// Short code identifying the Meter the measurement is for.
	Meter param.Field[string] `json:"meter,required"`
	// Timestamp for the measurement _(in ISO 8601 format)_.
	Ts param.Field[time.Time] `json:"ts,required" format:"date-time"`
	// 'cost' values
	Cost param.Field[map[string]float64] `json:"cost"`
	// End timestamp for the measurement _(in ISO 8601 format)_. _(Optional)_.
	//
	// Can be used in the case a usage event needs to have an explicit start and end
	// rather than being instantaneous.
	Ets param.Field[time.Time] `json:"ets" format:"date-time"`
	// 'income' values
	Income param.Field[map[string]float64] `json:"income"`
	// 'measure' values
	Measure param.Field[map[string]float64] `json:"measure"`
	// 'metadata' values
	Metadata param.Field[map[string]string] `json:"metadata"`
	// 'other' values
	Other param.Field[map[string]string] `json:"other"`
	// Unique ID for this measurement.
	Uid param.Field[string] `json:"uid"`
	// 'what' values
	What param.Field[map[string]string] `json:"what"`
	// 'where' values
	Where param.Field[map[string]string] `json:"where"`
	// 'who' values
	Who param.Field[map[string]string] `json:"who"`
}

func (r MeasurementRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubmitMeasurementsRequestParam struct {
	// Request containing the usage data measurements for submission.
	Measurements param.Field[[]MeasurementRequestParam] `json:"measurements,required"`
}

func (r SubmitMeasurementsRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubmitMeasurementsResponse struct {
	// `accepted` is returned when successful.
	Result string                         `json:"result"`
	JSON   submitMeasurementsResponseJSON `json:"-"`
}

// submitMeasurementsResponseJSON contains the JSON metadata for the struct
// [SubmitMeasurementsResponse]
type submitMeasurementsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmitMeasurementsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submitMeasurementsResponseJSON) RawJSON() string {
	return r.raw
}

type UsageQueryResponse struct {
	Data []map[string]interface{} `json:"data"`
	// Boolean flag to indicate whether or not there are more data available for the
	// query than are returned:
	//
	// - If there are more data, then TRUE.
	// - If there are no more data, then FALSE.
	//
	// **NOTES:**
	//
	//   - The limit on the size of the return is 20000 data items. If the query returns
	//     more than this limit, only 20000 items are returned with most recent first and
	//     `hasMoreData` will be TRUE.
	//   - If you have set `limit` in your query request at fewer than the number
	//     returned by the query, then `hasMoreData` will be TRUE in the response.
	HasMoreData bool                   `json:"hasMoreData"`
	JSON        usageQueryResponseJSON `json:"-"`
}

// usageQueryResponseJSON contains the JSON metadata for the struct
// [UsageQueryResponse]
type usageQueryResponseJSON struct {
	Data        apijson.Field
	HasMoreData apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageQueryResponseJSON) RawJSON() string {
	return r.raw
}

type UsageGetFailedIngestDownloadURLParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The file path
	File param.Field[string] `query:"file"`
}

// URLQuery serializes [UsageGetFailedIngestDownloadURLParams]'s query parameters
// as `url.Values`.
func (r UsageGetFailedIngestDownloadURLParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UsageQueryParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Specify the Accounts you want the query to return usage data for.
	AccountIDs param.Field[[]string] `json:"accountIds"`
	// Define the Aggregation functions you want to apply to data fields on included
	// Meters:
	//
	// - **SUM**. Adds the values.
	// - **MIN**. Uses the minimum value.
	// - **MAX**. Uses the maximum value.
	// - **COUNT**. Counts the number of values.
	// - **LATEST**. Uses the most recent value.
	// - **MEAN**. Uses the arithmetic mean of the values.
	// - **UNIQUE**. Uses a count of the number of unique values.
	//
	// **NOTE!** The Aggregation functions that can be applied depend on the data field
	// type:
	//
	//   - **Measure** fields. `SUM`, `MIN`, `MAX`, `COUNT`, `LATEST`, or `MEAN`
	//     functions can be applied.
	//   - **Dimension** field. `COUNT` or `UNIQUE` functions can be applied.
	Aggregations param.Field[[]UsageQueryParamsAggregation] `json:"aggregations"`
	// Define Dimension filters you want to apply for the query.
	//
	// Specify values for Dimension data fields on included Meters. Only data that
	// match the specified Dimension field values will be returned for the query.
	DimensionFilters param.Field[[]UsageQueryParamsDimensionFilter] `json:"dimensionFilters"`
	// The exclusive end date to define a time period to filter by. (_ISO 8601
	// formatted_)
	EndDate param.Field[time.Time] `json:"endDate" format:"date-time"`
	// If you've applied Aggregations for your query, specify any grouping you want to
	// impose on the returned data:
	//
	//   - **Account**
	//   - **Time** - group by frequency. Five options: `DAY`, `HOUR`, `WEEK`, `MONTH`,
	//     or `QUARTER`.
	//   - **Dimension** - group by Meter and data field.
	//
	// **NOTE:** If you attempt to impose grouping for a query that doesn't apply
	// Aggregations, you'll receive an error.
	Groups param.Field[[]UsageQueryParamsGroupUnion] `json:"groups"`
	// Define a limit for the number of usage data items you want the query to return,
	// starting with the most recently received data item.
	Limit param.Field[int64] `json:"limit"`
	// Specify the Meters you want the query to return usage data for.
	MeterIDs param.Field[[]string] `json:"meterIds"`
	// The inclusive start date to define a time period to filter by. (_ISO 8601
	// formatted_)
	StartDate param.Field[time.Time] `json:"startDate" format:"date-time"`
}

func (r UsageQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UsageQueryParamsAggregation struct {
	// Field code
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Type of field
	FieldType param.Field[UsageQueryParamsAggregationsFieldType] `json:"fieldType,required"`
	// Aggregation function
	Function param.Field[UsageQueryParamsAggregationsFunction] `json:"function,required"`
	// Meter ID
	MeterID param.Field[string] `json:"meterId,required"`
}

func (r UsageQueryParamsAggregation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of field
type UsageQueryParamsAggregationsFieldType string

const (
	UsageQueryParamsAggregationsFieldTypeDimension UsageQueryParamsAggregationsFieldType = "DIMENSION"
	UsageQueryParamsAggregationsFieldTypeMeasure   UsageQueryParamsAggregationsFieldType = "MEASURE"
)

func (r UsageQueryParamsAggregationsFieldType) IsKnown() bool {
	switch r {
	case UsageQueryParamsAggregationsFieldTypeDimension, UsageQueryParamsAggregationsFieldTypeMeasure:
		return true
	}
	return false
}

// Aggregation function
type UsageQueryParamsAggregationsFunction string

const (
	UsageQueryParamsAggregationsFunctionSum    UsageQueryParamsAggregationsFunction = "SUM"
	UsageQueryParamsAggregationsFunctionMin    UsageQueryParamsAggregationsFunction = "MIN"
	UsageQueryParamsAggregationsFunctionMax    UsageQueryParamsAggregationsFunction = "MAX"
	UsageQueryParamsAggregationsFunctionCount  UsageQueryParamsAggregationsFunction = "COUNT"
	UsageQueryParamsAggregationsFunctionLatest UsageQueryParamsAggregationsFunction = "LATEST"
	UsageQueryParamsAggregationsFunctionMean   UsageQueryParamsAggregationsFunction = "MEAN"
	UsageQueryParamsAggregationsFunctionUnique UsageQueryParamsAggregationsFunction = "UNIQUE"
)

func (r UsageQueryParamsAggregationsFunction) IsKnown() bool {
	switch r {
	case UsageQueryParamsAggregationsFunctionSum, UsageQueryParamsAggregationsFunctionMin, UsageQueryParamsAggregationsFunctionMax, UsageQueryParamsAggregationsFunctionCount, UsageQueryParamsAggregationsFunctionLatest, UsageQueryParamsAggregationsFunctionMean, UsageQueryParamsAggregationsFunctionUnique:
		return true
	}
	return false
}

type UsageQueryParamsDimensionFilter struct {
	// Field code
	FieldCode param.Field[string] `json:"fieldCode,required"`
	// Meter ID
	MeterID param.Field[string] `json:"meterId,required"`
	// Values to filter by
	Values param.Field[[]string] `json:"values,required"`
}

func (r UsageQueryParamsDimensionFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Group by a field
//
// Satisfied by [UsageQueryParamsGroupsDataExportsDataExplorerAccountGroup],
// [UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroup],
// [UsageQueryParamsGroupsDataExportsDataExplorerTimeGroup].
type UsageQueryParamsGroupUnion interface {
	implementsUsageQueryParamsGroupUnion()
}

// Group by account
type UsageQueryParamsGroupsDataExportsDataExplorerAccountGroup struct {
	GroupType param.Field[UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupType] `json:"groupType"`
	DataExplorerAccountGroupParam
}

func (r UsageQueryParamsGroupsDataExportsDataExplorerAccountGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageQueryParamsGroupsDataExportsDataExplorerAccountGroup) implementsUsageQueryParamsGroupUnion() {
}

type UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupType string

const (
	UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount   UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupType = "ACCOUNT"
	UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupType = "DIMENSION"
	UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupTypeTime      UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupType = "TIME"
)

func (r UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupType) IsKnown() bool {
	switch r {
	case UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupTypeAccount, UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupTypeDimension, UsageQueryParamsGroupsDataExportsDataExplorerAccountGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by dimension
type UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroup struct {
	GroupType param.Field[UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupType] `json:"groupType"`
	DataExplorerDimensionGroupParam
}

func (r UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroup) implementsUsageQueryParamsGroupUnion() {
}

type UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupType string

const (
	UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount   UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupType = "ACCOUNT"
	UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupType = "DIMENSION"
	UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime      UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupType = "TIME"
)

func (r UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupType) IsKnown() bool {
	switch r {
	case UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupTypeAccount, UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupTypeDimension, UsageQueryParamsGroupsDataExportsDataExplorerDimensionGroupGroupTypeTime:
		return true
	}
	return false
}

// Group by time
type UsageQueryParamsGroupsDataExportsDataExplorerTimeGroup struct {
	GroupType param.Field[UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupType] `json:"groupType"`
	DataExplorerTimeGroupParam
}

func (r UsageQueryParamsGroupsDataExportsDataExplorerTimeGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UsageQueryParamsGroupsDataExportsDataExplorerTimeGroup) implementsUsageQueryParamsGroupUnion() {
}

type UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupType string

const (
	UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount   UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupType = "ACCOUNT"
	UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupType = "DIMENSION"
	UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupTypeTime      UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupType = "TIME"
)

func (r UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupType) IsKnown() bool {
	switch r {
	case UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupTypeAccount, UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupTypeDimension, UsageQueryParamsGroupsDataExportsDataExplorerTimeGroupGroupTypeTime:
		return true
	}
	return false
}

type UsageSubmitParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID                     param.Field[string]            `path:"orgId,required"`
	SubmitMeasurementsRequest SubmitMeasurementsRequestParam `json:"submit_measurements_request,required"`
}

func (r UsageSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubmitMeasurementsRequest)
}
