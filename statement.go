// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// Endpoints for billing operations such as creating, updating,
// listing,downloading, and deleting Bills.
//
// Bills are generated for an Account, and are calculated in accordance with the
// usage-based pricing Plans applied for the Products the Account consumes. These
// endpoints enable interaction with the billing system, allowing you to obtain
// billing details and insights into the consumption patterns and charges of your
// end-customer Accounts.
//
// StatementService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementService] method instead.
type StatementService struct {
	Options []option.RequestOption
	// Endpoints for creating, retrieving, listing, and cancelling statement jobs.
	//
	// StatementJobs are tasks to asynchronously calculate and generate a bill
	// statement.
	//
	// Bill statements are informative backing sheets to invoices. They provide a
	// breakdown of the usage charges that appear on the bill, helping your end
	// customers better understand those charges, and gain a clearer picture of their
	// usage over the billing period.
	StatementJobs *StatementStatementJobService
	// Endpoints for listing, creating, updating, retrieving, or deleting Statement
	// Definitions.
	//
	// Bill statements are informative backing sheets to invoices. They provide a
	// breakdown of the usage charges that appear on the bill, helping your end
	// customers better understand those charges, and gain a clearer picture of their
	// usage over the billing period.
	//
	// Statement Definitions specify the way billed usage will be aggregated and
	// compiled in the Statement. For example, if you are billing customers monthly,
	// you might want to breakdown the usage responsible for the monthly charge on a
	// Bill into weekly portions in Bill statements.
	StatementDefinitions *StatementStatementDefinitionService
}

// NewStatementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStatementService(opts ...option.RequestOption) (r *StatementService) {
	r = &StatementService{}
	r.Options = opts
	r.StatementJobs = NewStatementStatementJobService(opts...)
	r.StatementDefinitions = NewStatementStatementDefinitionService(opts...)
	return
}

// Generate a specific Bill Statement for the provided Bill UUID in CSV format.
//
// Bill Statements are backing sheets to the invoices sent to your customers. Bill
// Statements provide a breakdown of the usage responsible for the usage charge
// line items shown on invoices.
//
// The response to this call returns a pre-signed `downloadUrl`, which you then use
// with a `GET` call to obtain the Bill statement in CSV format.
func (r *StatementService) NewCsv(ctx context.Context, id string, body StatementNewCsvParams, opts ...option.RequestOption) (res *ObjectURLResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/statement/csv", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve a specific Bill Statement for the given Bill UUID in CSV format.
//
// Bill Statements are backing sheets to the invoices sent to your customers. Bill
// Statements provide a breakdown of the usage responsible for the usage charge
// line items shown on invoices.
//
// The response includes a pre-signed `downloadUrl`, which must be used with a
// separate `GET` call to download the actual Bill Statement. This ensures secure
// access to the requested information.
func (r *StatementService) GetCsv(ctx context.Context, id string, query StatementGetCsvParams, opts ...option.RequestOption) (res *ObjectURLResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/statement/csv", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a Bill Statement in JSON format for a given Bill ID.
//
// Bill Statements are backing sheets to the invoices sent to your customers. Bill
// Statements provide a breakdown of the usage responsible for the usage charge
// line items shown on invoices.
//
// The response to this call returns a pre-signed `downloadUrl`, which you use with
// a `GET` call to obtain the Bill Statement.
func (r *StatementService) GetJson(ctx context.Context, id string, query StatementGetJsonParams, opts ...option.RequestOption) (res *ObjectURLResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/bills/%s/statement/json", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ObjectURLResponse struct {
	// The pre-signed download URL.
	DownloadURL string                `json:"downloadUrl" format:"url"`
	JSON        objectURLResponseJSON `json:"-"`
}

// objectURLResponseJSON contains the JSON metadata for the struct
// [ObjectURLResponse]
type objectURLResponseJSON struct {
	DownloadURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectURLResponseJSON) RawJSON() string {
	return r.raw
}

type StatementDefinitionResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// This specifies how often the Statement should aggregate data.
	AggregationFrequency StatementDefinitionResponseAggregationFrequency `json:"aggregationFrequency"`
	// The unique identifier (UUID) of the user who created this StatementDefinition.
	CreatedBy string `json:"createdBy"`
	// An array of objects, each representing a Dimension data field from a Meter _(for
	// Meters that have Dimensions setup)_.
	Dimensions []StatementDefinitionResponseDimension `json:"dimensions"`
	// The date and time _(in ISO-8601 format)_ when the StatementDefinition was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the StatementDefinition was last
	// modified.
	DtLastModified         time.Time `json:"dtLastModified" format:"date-time"`
	GenerateSlimStatements bool      `json:"generateSlimStatements"`
	// A Boolean indicating whether to include the price per unit in the Statement.
	//
	// - TRUE - includes the price per unit.
	// - FALSE - excludes the price per unit.
	IncludePricePerUnit bool `json:"includePricePerUnit"`
	// The unique identifier (UUID) of the user who last modified this
	// StatementDefinition.
	LastModifiedBy string `json:"lastModifiedBy"`
	// An array of objects, each representing a Measure data field from a Meter.
	Measures []StatementDefinitionResponseMeasure `json:"measures"`
	// Descriptive name for the StatementDefinition providing context and information.
	Name string `json:"name"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                           `json:"version"`
	JSON    statementDefinitionResponseJSON `json:"-"`
}

// statementDefinitionResponseJSON contains the JSON metadata for the struct
// [StatementDefinitionResponse]
type statementDefinitionResponseJSON struct {
	ID                     apijson.Field
	AggregationFrequency   apijson.Field
	CreatedBy              apijson.Field
	Dimensions             apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	GenerateSlimStatements apijson.Field
	IncludePricePerUnit    apijson.Field
	LastModifiedBy         apijson.Field
	Measures               apijson.Field
	Name                   apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *StatementDefinitionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementDefinitionResponseJSON) RawJSON() string {
	return r.raw
}

// This specifies how often the Statement should aggregate data.
type StatementDefinitionResponseAggregationFrequency string

const (
	StatementDefinitionResponseAggregationFrequencyDay         StatementDefinitionResponseAggregationFrequency = "DAY"
	StatementDefinitionResponseAggregationFrequencyWeek        StatementDefinitionResponseAggregationFrequency = "WEEK"
	StatementDefinitionResponseAggregationFrequencyMonth       StatementDefinitionResponseAggregationFrequency = "MONTH"
	StatementDefinitionResponseAggregationFrequencyQuarter     StatementDefinitionResponseAggregationFrequency = "QUARTER"
	StatementDefinitionResponseAggregationFrequencyYear        StatementDefinitionResponseAggregationFrequency = "YEAR"
	StatementDefinitionResponseAggregationFrequencyWholePeriod StatementDefinitionResponseAggregationFrequency = "WHOLE_PERIOD"
)

func (r StatementDefinitionResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case StatementDefinitionResponseAggregationFrequencyDay, StatementDefinitionResponseAggregationFrequencyWeek, StatementDefinitionResponseAggregationFrequencyMonth, StatementDefinitionResponseAggregationFrequencyQuarter, StatementDefinitionResponseAggregationFrequencyYear, StatementDefinitionResponseAggregationFrequencyWholePeriod:
		return true
	}
	return false
}

// A Dimension belonging to a Meter.
type StatementDefinitionResponseDimension struct {
	// The value of a Dimension to use as a filter. Use "\*" as a wildcard to filter on
	// all Dimension values.
	Filter []string `json:"filter" api:"required"`
	// The name of the Dimension to target in the Meter.
	Name string `json:"name" api:"required"`
	// The Dimension attribute to target.
	Attributes []string `json:"attributes"`
	// The unique identifier (UUID) of the Meter containing this Dimension.
	MeterID string                                   `json:"meterId"`
	JSON    statementDefinitionResponseDimensionJSON `json:"-"`
}

// statementDefinitionResponseDimensionJSON contains the JSON metadata for the
// struct [StatementDefinitionResponseDimension]
type statementDefinitionResponseDimensionJSON struct {
	Filter      apijson.Field
	Name        apijson.Field
	Attributes  apijson.Field
	MeterID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatementDefinitionResponseDimension) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementDefinitionResponseDimensionJSON) RawJSON() string {
	return r.raw
}

type StatementDefinitionResponseMeasure struct {
	// A list of Aggregations to apply to the Measure.
	Aggregations []StatementDefinitionResponseMeasuresAggregation `json:"aggregations"`
	// The unique identifier (UUID) of the Meter containing this Measure.
	MeterID string `json:"meterId"`
	// The name of a Measure data field _(or blank to indicate a wildcard, i.e. all
	// fields)_. Default value is blank.
	Name string                                 `json:"name"`
	JSON statementDefinitionResponseMeasureJSON `json:"-"`
}

// statementDefinitionResponseMeasureJSON contains the JSON metadata for the struct
// [StatementDefinitionResponseMeasure]
type statementDefinitionResponseMeasureJSON struct {
	Aggregations apijson.Field
	MeterID      apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *StatementDefinitionResponseMeasure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementDefinitionResponseMeasureJSON) RawJSON() string {
	return r.raw
}

// Specifies the computation method applied to usage data collected in
// `targetField`. Aggregation unit value depends on the **Category** configured for
// the selected targetField.
//
//   - **SUM**. Adds the values. Can be applied to a **Measure**, **Income**, or
//     **Cost** `targetField`.
//
//   - **MIN**. Uses the minimum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **MAX**. Uses the maximum value. Can be applied to a **Measure**, **Income**,
//     or **Cost** `targetField`.
//
//   - **COUNT**. Counts the number of values. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`.
//
//   - **LATEST**. Uses the most recent value. Can be applied to a **Measure**,
//     **Income**, or **Cost** `targetField`. Note: Based on the timestamp `ts` value
//     of usage data measurement submissions. If using this method, please ensure
//     _distinct_ `ts` values are used for usage data measurement submissions.
//
//   - **MEAN**. Uses the arithmetic mean of the values. Can be applied to a
//     **Measure**, **Income**, or **Cost** `targetField`.
//
//   - **UNIQUE**. Uses unique values and returns a count of the number of unique
//     values. Can be applied to a **Metadata** `targetField`.
type StatementDefinitionResponseMeasuresAggregation string

const (
	StatementDefinitionResponseMeasuresAggregationSum       StatementDefinitionResponseMeasuresAggregation = "SUM"
	StatementDefinitionResponseMeasuresAggregationMin       StatementDefinitionResponseMeasuresAggregation = "MIN"
	StatementDefinitionResponseMeasuresAggregationMax       StatementDefinitionResponseMeasuresAggregation = "MAX"
	StatementDefinitionResponseMeasuresAggregationCount     StatementDefinitionResponseMeasuresAggregation = "COUNT"
	StatementDefinitionResponseMeasuresAggregationLatest    StatementDefinitionResponseMeasuresAggregation = "LATEST"
	StatementDefinitionResponseMeasuresAggregationMean      StatementDefinitionResponseMeasuresAggregation = "MEAN"
	StatementDefinitionResponseMeasuresAggregationUnique    StatementDefinitionResponseMeasuresAggregation = "UNIQUE"
	StatementDefinitionResponseMeasuresAggregationCustomSql StatementDefinitionResponseMeasuresAggregation = "CUSTOM_SQL"
)

func (r StatementDefinitionResponseMeasuresAggregation) IsKnown() bool {
	switch r {
	case StatementDefinitionResponseMeasuresAggregationSum, StatementDefinitionResponseMeasuresAggregationMin, StatementDefinitionResponseMeasuresAggregationMax, StatementDefinitionResponseMeasuresAggregationCount, StatementDefinitionResponseMeasuresAggregationLatest, StatementDefinitionResponseMeasuresAggregationMean, StatementDefinitionResponseMeasuresAggregationUnique, StatementDefinitionResponseMeasuresAggregationCustomSql:
		return true
	}
	return false
}

type StatementJobResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// The unique identifier (UUID) of the bill associated with the StatementJob.
	BillID string `json:"billId"`
	// The unique identifier (UUID) of the user who created this StatementJob.
	CreatedBy          string                                 `json:"createdBy"`
	CsvStatementStatus StatementJobResponseCsvStatementStatus `json:"csvStatementStatus"`
	// The date and time _(in ISO-8601 format)_ when the StatementJob was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the StatementJob was last
	// modified.
	DtLastModified time.Time                   `json:"dtLastModified" format:"date-time"`
	Filters        StatementJobResponseFilters `json:"filters"`
	// A Boolean value indicating whether the generated statement includes a CSV
	// format.
	//
	// - TRUE - includes the statement in CSV format.
	// - FALSE - no CSV format statement.
	IncludeCsvFormat    bool                                    `json:"includeCsvFormat"`
	JsonStatementStatus StatementJobResponseJsonStatementStatus `json:"jsonStatementStatus"`
	// The unique identifier (UUID) of the user who last modified this StatementJob.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The unique identifier (UUID) of your Organization. The Organization represents
	// your company as a direct customer of our service.
	OrgID                    string `json:"orgId"`
	PresignedCsvStatementURL string `json:"presignedCsvStatementUrl"`
	// The URL to access the generated statement in JSON format. This URL is temporary
	// and has a limited lifetime.
	PresignedJsonStatementURL string `json:"presignedJsonStatementUrl"`
	// The current status of the StatementJob. The status helps track the progress and
	// outcome of a StatementJob.
	StatementJobStatus StatementJobResponseStatementJobStatus `json:"statementJobStatus"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                    `json:"version"`
	JSON    statementJobResponseJSON `json:"-"`
}

// statementJobResponseJSON contains the JSON metadata for the struct
// [StatementJobResponse]
type statementJobResponseJSON struct {
	ID                        apijson.Field
	BillID                    apijson.Field
	CreatedBy                 apijson.Field
	CsvStatementStatus        apijson.Field
	DtCreated                 apijson.Field
	DtLastModified            apijson.Field
	Filters                   apijson.Field
	IncludeCsvFormat          apijson.Field
	JsonStatementStatus       apijson.Field
	LastModifiedBy            apijson.Field
	OrgID                     apijson.Field
	PresignedCsvStatementURL  apijson.Field
	PresignedJsonStatementURL apijson.Field
	StatementJobStatus        apijson.Field
	Version                   apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *StatementJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementJobResponseJSON) RawJSON() string {
	return r.raw
}

type StatementJobResponseCsvStatementStatus string

const (
	StatementJobResponseCsvStatementStatusLatest      StatementJobResponseCsvStatementStatus = "LATEST"
	StatementJobResponseCsvStatementStatusStale       StatementJobResponseCsvStatementStatus = "STALE"
	StatementJobResponseCsvStatementStatusInvalidated StatementJobResponseCsvStatementStatus = "INVALIDATED"
)

func (r StatementJobResponseCsvStatementStatus) IsKnown() bool {
	switch r {
	case StatementJobResponseCsvStatementStatusLatest, StatementJobResponseCsvStatementStatusStale, StatementJobResponseCsvStatementStatusInvalidated:
		return true
	}
	return false
}

type StatementJobResponseFilters struct {
	// Include usage line items whose meterId matches one of these values.
	MeterIDs []string                        `json:"meterIds"`
	JSON     statementJobResponseFiltersJSON `json:"-"`
}

// statementJobResponseFiltersJSON contains the JSON metadata for the struct
// [StatementJobResponseFilters]
type statementJobResponseFiltersJSON struct {
	MeterIDs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StatementJobResponseFilters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r statementJobResponseFiltersJSON) RawJSON() string {
	return r.raw
}

type StatementJobResponseJsonStatementStatus string

const (
	StatementJobResponseJsonStatementStatusLatest      StatementJobResponseJsonStatementStatus = "LATEST"
	StatementJobResponseJsonStatementStatusStale       StatementJobResponseJsonStatementStatus = "STALE"
	StatementJobResponseJsonStatementStatusInvalidated StatementJobResponseJsonStatementStatus = "INVALIDATED"
)

func (r StatementJobResponseJsonStatementStatus) IsKnown() bool {
	switch r {
	case StatementJobResponseJsonStatementStatusLatest, StatementJobResponseJsonStatementStatusStale, StatementJobResponseJsonStatementStatusInvalidated:
		return true
	}
	return false
}

// The current status of the StatementJob. The status helps track the progress and
// outcome of a StatementJob.
type StatementJobResponseStatementJobStatus string

const (
	StatementJobResponseStatementJobStatusPending   StatementJobResponseStatementJobStatus = "PENDING"
	StatementJobResponseStatementJobStatusRunning   StatementJobResponseStatementJobStatus = "RUNNING"
	StatementJobResponseStatementJobStatusComplete  StatementJobResponseStatementJobStatus = "COMPLETE"
	StatementJobResponseStatementJobStatusCancelled StatementJobResponseStatementJobStatus = "CANCELLED"
	StatementJobResponseStatementJobStatusFailed    StatementJobResponseStatementJobStatus = "FAILED"
)

func (r StatementJobResponseStatementJobStatus) IsKnown() bool {
	switch r {
	case StatementJobResponseStatementJobStatusPending, StatementJobResponseStatementJobStatusRunning, StatementJobResponseStatementJobStatusComplete, StatementJobResponseStatementJobStatusCancelled, StatementJobResponseStatementJobStatusFailed:
		return true
	}
	return false
}

type StatementNewCsvParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type StatementGetCsvParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type StatementGetJsonParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
