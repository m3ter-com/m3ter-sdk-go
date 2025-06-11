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

// StatementStatementDefinitionService contains methods and other services that
// help with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementStatementDefinitionService] method instead.
type StatementStatementDefinitionService struct {
	Options []option.RequestOption
}

// NewStatementStatementDefinitionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewStatementStatementDefinitionService(opts ...option.RequestOption) (r *StatementStatementDefinitionService) {
	r = &StatementStatementDefinitionService{}
	r.Options = opts
	return
}

// Create a new StatementDefinition.
//
// This endpoint creates a new StatementDefinition within the specified
// Organization. The details of the StatementDefinition are provided in the request
// body.
func (r *StatementStatementDefinitionService) New(ctx context.Context, params StatementStatementDefinitionNewParams, opts ...option.RequestOption) (res *StatementDefinitionResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementdefinitions", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a StatementDefinition with the given UUID.
//
// Retrieves the details of a specific StatementDefinition for the specified
// Organization, using its unique identifier (UUID). This endpoint is useful when
// you want to retrieve the complete details of a single StatementDefinition.
func (r *StatementStatementDefinitionService) Get(ctx context.Context, id string, query StatementStatementDefinitionGetParams, opts ...option.RequestOption) (res *StatementDefinitionResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementdefinitions/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update StatementDefinition for the given UUID.
//
// Update the details of a specific StatementDefinition for the specified
// Organization, using its unique identifier (UUID). The updated details for the
// StatementDefinition should be sent in the request body.
func (r *StatementStatementDefinitionService) Update(ctx context.Context, id string, params StatementStatementDefinitionUpdateParams, opts ...option.RequestOption) (res *StatementDefinitionResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementdefinitions/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of StatementDefinitions.
//
// This endpoint retrieves a list of all the StatementDefinitions within a
// specified Organization. The list can be paginated for easier management.
func (r *StatementStatementDefinitionService) List(ctx context.Context, params StatementStatementDefinitionListParams, opts ...option.RequestOption) (res *pagination.Cursor[StatementDefinitionResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/statementdefinitions", params.OrgID)
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

// Retrieve a list of StatementDefinitions.
//
// This endpoint retrieves a list of all the StatementDefinitions within a
// specified Organization. The list can be paginated for easier management.
func (r *StatementStatementDefinitionService) ListAutoPaging(ctx context.Context, params StatementStatementDefinitionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[StatementDefinitionResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a StatementDefinition with the given UUID.
//
// This endpoint deletes a specific StatementDefinition within a specified
// Organization, using the StatementDefinition UUID.
func (r *StatementStatementDefinitionService) Delete(ctx context.Context, id string, body StatementStatementDefinitionDeleteParams, opts ...option.RequestOption) (res *StatementDefinitionResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementdefinitions/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type StatementDefinitionResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
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
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
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
	Name string                          `json:"name"`
	JSON statementDefinitionResponseJSON `json:"-"`
}

// statementDefinitionResponseJSON contains the JSON metadata for the struct
// [StatementDefinitionResponse]
type statementDefinitionResponseJSON struct {
	ID                   apijson.Field
	Version              apijson.Field
	AggregationFrequency apijson.Field
	CreatedBy            apijson.Field
	Dimensions           apijson.Field
	DtCreated            apijson.Field
	DtLastModified       apijson.Field
	IncludePricePerUnit  apijson.Field
	LastModifiedBy       apijson.Field
	Measures             apijson.Field
	Name                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
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
	StatementDefinitionResponseAggregationFrequencyOriginal    StatementDefinitionResponseAggregationFrequency = "ORIGINAL"
	StatementDefinitionResponseAggregationFrequencyHour        StatementDefinitionResponseAggregationFrequency = "HOUR"
	StatementDefinitionResponseAggregationFrequencyDay         StatementDefinitionResponseAggregationFrequency = "DAY"
	StatementDefinitionResponseAggregationFrequencyWeek        StatementDefinitionResponseAggregationFrequency = "WEEK"
	StatementDefinitionResponseAggregationFrequencyMonth       StatementDefinitionResponseAggregationFrequency = "MONTH"
	StatementDefinitionResponseAggregationFrequencyQuarter     StatementDefinitionResponseAggregationFrequency = "QUARTER"
	StatementDefinitionResponseAggregationFrequencyYear        StatementDefinitionResponseAggregationFrequency = "YEAR"
	StatementDefinitionResponseAggregationFrequencyWholePeriod StatementDefinitionResponseAggregationFrequency = "WHOLE_PERIOD"
)

func (r StatementDefinitionResponseAggregationFrequency) IsKnown() bool {
	switch r {
	case StatementDefinitionResponseAggregationFrequencyOriginal, StatementDefinitionResponseAggregationFrequencyHour, StatementDefinitionResponseAggregationFrequencyDay, StatementDefinitionResponseAggregationFrequencyWeek, StatementDefinitionResponseAggregationFrequencyMonth, StatementDefinitionResponseAggregationFrequencyQuarter, StatementDefinitionResponseAggregationFrequencyYear, StatementDefinitionResponseAggregationFrequencyWholePeriod:
		return true
	}
	return false
}

// A Dimension belonging to a Meter.
type StatementDefinitionResponseDimension struct {
	// Attributes belonging to the dimension
	DimensionAttributes []string `json:"dimensionAttributes"`
	// The name of a dimension
	DimensionName string                                   `json:"dimensionName"`
	JSON          statementDefinitionResponseDimensionJSON `json:"-"`
}

// statementDefinitionResponseDimensionJSON contains the JSON metadata for the
// struct [StatementDefinitionResponseDimension]
type statementDefinitionResponseDimensionJSON struct {
	DimensionAttributes apijson.Field
	DimensionName       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
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
	StatementDefinitionResponseMeasuresAggregationSum    StatementDefinitionResponseMeasuresAggregation = "SUM"
	StatementDefinitionResponseMeasuresAggregationMin    StatementDefinitionResponseMeasuresAggregation = "MIN"
	StatementDefinitionResponseMeasuresAggregationMax    StatementDefinitionResponseMeasuresAggregation = "MAX"
	StatementDefinitionResponseMeasuresAggregationCount  StatementDefinitionResponseMeasuresAggregation = "COUNT"
	StatementDefinitionResponseMeasuresAggregationLatest StatementDefinitionResponseMeasuresAggregation = "LATEST"
	StatementDefinitionResponseMeasuresAggregationMean   StatementDefinitionResponseMeasuresAggregation = "MEAN"
	StatementDefinitionResponseMeasuresAggregationUnique StatementDefinitionResponseMeasuresAggregation = "UNIQUE"
)

func (r StatementDefinitionResponseMeasuresAggregation) IsKnown() bool {
	switch r {
	case StatementDefinitionResponseMeasuresAggregationSum, StatementDefinitionResponseMeasuresAggregationMin, StatementDefinitionResponseMeasuresAggregationMax, StatementDefinitionResponseMeasuresAggregationCount, StatementDefinitionResponseMeasuresAggregationLatest, StatementDefinitionResponseMeasuresAggregationMean, StatementDefinitionResponseMeasuresAggregationUnique:
		return true
	}
	return false
}

type StatementStatementDefinitionNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// This specifies how often the Statement should aggregate data.
	AggregationFrequency param.Field[StatementStatementDefinitionNewParamsAggregationFrequency] `json:"aggregationFrequency,required"`
	// An array of objects, each representing a Dimension data field from a Meter _(for
	// Meters that have Dimensions setup)_.
	Dimensions param.Field[[]StatementStatementDefinitionNewParamsDimension] `json:"dimensions"`
	// A Boolean indicating whether to include the price per unit in the Statement.
	//
	// - TRUE - includes the price per unit.
	// - FALSE - excludes the price per unit.
	IncludePricePerUnit param.Field[bool] `json:"includePricePerUnit"`
	// An array of objects, each representing a Measure data field from a Meter.
	Measures param.Field[[]StatementStatementDefinitionNewParamsMeasure] `json:"measures"`
	// Descriptive name for the StatementDefinition providing context and information.
	Name param.Field[string] `json:"name"`
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

func (r StatementStatementDefinitionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This specifies how often the Statement should aggregate data.
type StatementStatementDefinitionNewParamsAggregationFrequency string

const (
	StatementStatementDefinitionNewParamsAggregationFrequencyOriginal    StatementStatementDefinitionNewParamsAggregationFrequency = "ORIGINAL"
	StatementStatementDefinitionNewParamsAggregationFrequencyHour        StatementStatementDefinitionNewParamsAggregationFrequency = "HOUR"
	StatementStatementDefinitionNewParamsAggregationFrequencyDay         StatementStatementDefinitionNewParamsAggregationFrequency = "DAY"
	StatementStatementDefinitionNewParamsAggregationFrequencyWeek        StatementStatementDefinitionNewParamsAggregationFrequency = "WEEK"
	StatementStatementDefinitionNewParamsAggregationFrequencyMonth       StatementStatementDefinitionNewParamsAggregationFrequency = "MONTH"
	StatementStatementDefinitionNewParamsAggregationFrequencyQuarter     StatementStatementDefinitionNewParamsAggregationFrequency = "QUARTER"
	StatementStatementDefinitionNewParamsAggregationFrequencyYear        StatementStatementDefinitionNewParamsAggregationFrequency = "YEAR"
	StatementStatementDefinitionNewParamsAggregationFrequencyWholePeriod StatementStatementDefinitionNewParamsAggregationFrequency = "WHOLE_PERIOD"
)

func (r StatementStatementDefinitionNewParamsAggregationFrequency) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionNewParamsAggregationFrequencyOriginal, StatementStatementDefinitionNewParamsAggregationFrequencyHour, StatementStatementDefinitionNewParamsAggregationFrequencyDay, StatementStatementDefinitionNewParamsAggregationFrequencyWeek, StatementStatementDefinitionNewParamsAggregationFrequencyMonth, StatementStatementDefinitionNewParamsAggregationFrequencyQuarter, StatementStatementDefinitionNewParamsAggregationFrequencyYear, StatementStatementDefinitionNewParamsAggregationFrequencyWholePeriod:
		return true
	}
	return false
}

// A Dimension belonging to a Meter.
type StatementStatementDefinitionNewParamsDimension struct {
	// Attributes belonging to the dimension
	DimensionAttributes param.Field[[]string] `json:"dimensionAttributes"`
	// The name of a dimension
	DimensionName param.Field[string] `json:"dimensionName"`
}

func (r StatementStatementDefinitionNewParamsDimension) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StatementStatementDefinitionNewParamsMeasure struct {
	// A list of Aggregations to apply to the Measure.
	Aggregations param.Field[[]StatementStatementDefinitionNewParamsMeasuresAggregation] `json:"aggregations"`
	// The unique identifier (UUID) of the Meter containing this Measure.
	MeterID param.Field[string] `json:"meterId"`
	// The name of a Measure data field _(or blank to indicate a wildcard, i.e. all
	// fields)_. Default value is blank.
	Name param.Field[string] `json:"name"`
}

func (r StatementStatementDefinitionNewParamsMeasure) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
type StatementStatementDefinitionNewParamsMeasuresAggregation string

const (
	StatementStatementDefinitionNewParamsMeasuresAggregationSum    StatementStatementDefinitionNewParamsMeasuresAggregation = "SUM"
	StatementStatementDefinitionNewParamsMeasuresAggregationMin    StatementStatementDefinitionNewParamsMeasuresAggregation = "MIN"
	StatementStatementDefinitionNewParamsMeasuresAggregationMax    StatementStatementDefinitionNewParamsMeasuresAggregation = "MAX"
	StatementStatementDefinitionNewParamsMeasuresAggregationCount  StatementStatementDefinitionNewParamsMeasuresAggregation = "COUNT"
	StatementStatementDefinitionNewParamsMeasuresAggregationLatest StatementStatementDefinitionNewParamsMeasuresAggregation = "LATEST"
	StatementStatementDefinitionNewParamsMeasuresAggregationMean   StatementStatementDefinitionNewParamsMeasuresAggregation = "MEAN"
	StatementStatementDefinitionNewParamsMeasuresAggregationUnique StatementStatementDefinitionNewParamsMeasuresAggregation = "UNIQUE"
)

func (r StatementStatementDefinitionNewParamsMeasuresAggregation) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionNewParamsMeasuresAggregationSum, StatementStatementDefinitionNewParamsMeasuresAggregationMin, StatementStatementDefinitionNewParamsMeasuresAggregationMax, StatementStatementDefinitionNewParamsMeasuresAggregationCount, StatementStatementDefinitionNewParamsMeasuresAggregationLatest, StatementStatementDefinitionNewParamsMeasuresAggregationMean, StatementStatementDefinitionNewParamsMeasuresAggregationUnique:
		return true
	}
	return false
}

type StatementStatementDefinitionGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type StatementStatementDefinitionUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// This specifies how often the Statement should aggregate data.
	AggregationFrequency param.Field[StatementStatementDefinitionUpdateParamsAggregationFrequency] `json:"aggregationFrequency,required"`
	// An array of objects, each representing a Dimension data field from a Meter _(for
	// Meters that have Dimensions setup)_.
	Dimensions param.Field[[]StatementStatementDefinitionUpdateParamsDimension] `json:"dimensions"`
	// A Boolean indicating whether to include the price per unit in the Statement.
	//
	// - TRUE - includes the price per unit.
	// - FALSE - excludes the price per unit.
	IncludePricePerUnit param.Field[bool] `json:"includePricePerUnit"`
	// An array of objects, each representing a Measure data field from a Meter.
	Measures param.Field[[]StatementStatementDefinitionUpdateParamsMeasure] `json:"measures"`
	// Descriptive name for the StatementDefinition providing context and information.
	Name param.Field[string] `json:"name"`
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

func (r StatementStatementDefinitionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This specifies how often the Statement should aggregate data.
type StatementStatementDefinitionUpdateParamsAggregationFrequency string

const (
	StatementStatementDefinitionUpdateParamsAggregationFrequencyOriginal    StatementStatementDefinitionUpdateParamsAggregationFrequency = "ORIGINAL"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyHour        StatementStatementDefinitionUpdateParamsAggregationFrequency = "HOUR"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyDay         StatementStatementDefinitionUpdateParamsAggregationFrequency = "DAY"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyWeek        StatementStatementDefinitionUpdateParamsAggregationFrequency = "WEEK"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyMonth       StatementStatementDefinitionUpdateParamsAggregationFrequency = "MONTH"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyQuarter     StatementStatementDefinitionUpdateParamsAggregationFrequency = "QUARTER"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyYear        StatementStatementDefinitionUpdateParamsAggregationFrequency = "YEAR"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyWholePeriod StatementStatementDefinitionUpdateParamsAggregationFrequency = "WHOLE_PERIOD"
)

func (r StatementStatementDefinitionUpdateParamsAggregationFrequency) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionUpdateParamsAggregationFrequencyOriginal, StatementStatementDefinitionUpdateParamsAggregationFrequencyHour, StatementStatementDefinitionUpdateParamsAggregationFrequencyDay, StatementStatementDefinitionUpdateParamsAggregationFrequencyWeek, StatementStatementDefinitionUpdateParamsAggregationFrequencyMonth, StatementStatementDefinitionUpdateParamsAggregationFrequencyQuarter, StatementStatementDefinitionUpdateParamsAggregationFrequencyYear, StatementStatementDefinitionUpdateParamsAggregationFrequencyWholePeriod:
		return true
	}
	return false
}

// A Dimension belonging to a Meter.
type StatementStatementDefinitionUpdateParamsDimension struct {
	// Attributes belonging to the dimension
	DimensionAttributes param.Field[[]string] `json:"dimensionAttributes"`
	// The name of a dimension
	DimensionName param.Field[string] `json:"dimensionName"`
}

func (r StatementStatementDefinitionUpdateParamsDimension) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StatementStatementDefinitionUpdateParamsMeasure struct {
	// A list of Aggregations to apply to the Measure.
	Aggregations param.Field[[]StatementStatementDefinitionUpdateParamsMeasuresAggregation] `json:"aggregations"`
	// The unique identifier (UUID) of the Meter containing this Measure.
	MeterID param.Field[string] `json:"meterId"`
	// The name of a Measure data field _(or blank to indicate a wildcard, i.e. all
	// fields)_. Default value is blank.
	Name param.Field[string] `json:"name"`
}

func (r StatementStatementDefinitionUpdateParamsMeasure) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
type StatementStatementDefinitionUpdateParamsMeasuresAggregation string

const (
	StatementStatementDefinitionUpdateParamsMeasuresAggregationSum    StatementStatementDefinitionUpdateParamsMeasuresAggregation = "SUM"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationMin    StatementStatementDefinitionUpdateParamsMeasuresAggregation = "MIN"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationMax    StatementStatementDefinitionUpdateParamsMeasuresAggregation = "MAX"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationCount  StatementStatementDefinitionUpdateParamsMeasuresAggregation = "COUNT"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationLatest StatementStatementDefinitionUpdateParamsMeasuresAggregation = "LATEST"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationMean   StatementStatementDefinitionUpdateParamsMeasuresAggregation = "MEAN"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationUnique StatementStatementDefinitionUpdateParamsMeasuresAggregation = "UNIQUE"
)

func (r StatementStatementDefinitionUpdateParamsMeasuresAggregation) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionUpdateParamsMeasuresAggregationSum, StatementStatementDefinitionUpdateParamsMeasuresAggregationMin, StatementStatementDefinitionUpdateParamsMeasuresAggregationMax, StatementStatementDefinitionUpdateParamsMeasuresAggregationCount, StatementStatementDefinitionUpdateParamsMeasuresAggregationLatest, StatementStatementDefinitionUpdateParamsMeasuresAggregationMean, StatementStatementDefinitionUpdateParamsMeasuresAggregationUnique:
		return true
	}
	return false
}

type StatementStatementDefinitionListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// StatementDefinitions in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of StatementDefinitions to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [StatementStatementDefinitionListParams]'s query parameters
// as `url.Values`.
func (r StatementStatementDefinitionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatementStatementDefinitionDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
