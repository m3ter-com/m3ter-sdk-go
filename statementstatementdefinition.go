// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("organizations/%s/statementdefinitions/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type StatementStatementDefinitionNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// This specifies how often the Statement should aggregate data.
	AggregationFrequency param.Field[StatementStatementDefinitionNewParamsAggregationFrequency] `json:"aggregationFrequency,required"`
	// An array of objects, each representing a Dimension data field from a Meter _(for
	// Meters that have Dimensions setup)_.
	Dimensions             param.Field[[]StatementStatementDefinitionNewParamsDimension] `json:"dimensions"`
	GenerateSlimStatements param.Field[bool]                                             `json:"generateSlimStatements"`
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
	StatementStatementDefinitionNewParamsAggregationFrequencyDay         StatementStatementDefinitionNewParamsAggregationFrequency = "DAY"
	StatementStatementDefinitionNewParamsAggregationFrequencyWeek        StatementStatementDefinitionNewParamsAggregationFrequency = "WEEK"
	StatementStatementDefinitionNewParamsAggregationFrequencyMonth       StatementStatementDefinitionNewParamsAggregationFrequency = "MONTH"
	StatementStatementDefinitionNewParamsAggregationFrequencyQuarter     StatementStatementDefinitionNewParamsAggregationFrequency = "QUARTER"
	StatementStatementDefinitionNewParamsAggregationFrequencyYear        StatementStatementDefinitionNewParamsAggregationFrequency = "YEAR"
	StatementStatementDefinitionNewParamsAggregationFrequencyWholePeriod StatementStatementDefinitionNewParamsAggregationFrequency = "WHOLE_PERIOD"
)

func (r StatementStatementDefinitionNewParamsAggregationFrequency) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionNewParamsAggregationFrequencyDay, StatementStatementDefinitionNewParamsAggregationFrequencyWeek, StatementStatementDefinitionNewParamsAggregationFrequencyMonth, StatementStatementDefinitionNewParamsAggregationFrequencyQuarter, StatementStatementDefinitionNewParamsAggregationFrequencyYear, StatementStatementDefinitionNewParamsAggregationFrequencyWholePeriod:
		return true
	}
	return false
}

// A Dimension belonging to a Meter.
type StatementStatementDefinitionNewParamsDimension struct {
	// The value of a Dimension to use as a filter. Use "\*" as a wildcard to filter on
	// all Dimension values.
	Filter param.Field[[]string] `json:"filter,required"`
	// The name of the Dimension to target in the Meter.
	Name param.Field[string] `json:"name,required"`
	// The Dimension attribute to target.
	Attributes param.Field[[]string] `json:"attributes"`
	// The unique identifier (UUID) of the Meter containing this Dimension.
	MeterID param.Field[string] `json:"meterId"`
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
	StatementStatementDefinitionNewParamsMeasuresAggregationSum       StatementStatementDefinitionNewParamsMeasuresAggregation = "SUM"
	StatementStatementDefinitionNewParamsMeasuresAggregationMin       StatementStatementDefinitionNewParamsMeasuresAggregation = "MIN"
	StatementStatementDefinitionNewParamsMeasuresAggregationMax       StatementStatementDefinitionNewParamsMeasuresAggregation = "MAX"
	StatementStatementDefinitionNewParamsMeasuresAggregationCount     StatementStatementDefinitionNewParamsMeasuresAggregation = "COUNT"
	StatementStatementDefinitionNewParamsMeasuresAggregationLatest    StatementStatementDefinitionNewParamsMeasuresAggregation = "LATEST"
	StatementStatementDefinitionNewParamsMeasuresAggregationMean      StatementStatementDefinitionNewParamsMeasuresAggregation = "MEAN"
	StatementStatementDefinitionNewParamsMeasuresAggregationUnique    StatementStatementDefinitionNewParamsMeasuresAggregation = "UNIQUE"
	StatementStatementDefinitionNewParamsMeasuresAggregationCustomSql StatementStatementDefinitionNewParamsMeasuresAggregation = "CUSTOM_SQL"
)

func (r StatementStatementDefinitionNewParamsMeasuresAggregation) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionNewParamsMeasuresAggregationSum, StatementStatementDefinitionNewParamsMeasuresAggregationMin, StatementStatementDefinitionNewParamsMeasuresAggregationMax, StatementStatementDefinitionNewParamsMeasuresAggregationCount, StatementStatementDefinitionNewParamsMeasuresAggregationLatest, StatementStatementDefinitionNewParamsMeasuresAggregationMean, StatementStatementDefinitionNewParamsMeasuresAggregationUnique, StatementStatementDefinitionNewParamsMeasuresAggregationCustomSql:
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
	Dimensions             param.Field[[]StatementStatementDefinitionUpdateParamsDimension] `json:"dimensions"`
	GenerateSlimStatements param.Field[bool]                                                `json:"generateSlimStatements"`
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
	StatementStatementDefinitionUpdateParamsAggregationFrequencyDay         StatementStatementDefinitionUpdateParamsAggregationFrequency = "DAY"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyWeek        StatementStatementDefinitionUpdateParamsAggregationFrequency = "WEEK"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyMonth       StatementStatementDefinitionUpdateParamsAggregationFrequency = "MONTH"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyQuarter     StatementStatementDefinitionUpdateParamsAggregationFrequency = "QUARTER"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyYear        StatementStatementDefinitionUpdateParamsAggregationFrequency = "YEAR"
	StatementStatementDefinitionUpdateParamsAggregationFrequencyWholePeriod StatementStatementDefinitionUpdateParamsAggregationFrequency = "WHOLE_PERIOD"
)

func (r StatementStatementDefinitionUpdateParamsAggregationFrequency) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionUpdateParamsAggregationFrequencyDay, StatementStatementDefinitionUpdateParamsAggregationFrequencyWeek, StatementStatementDefinitionUpdateParamsAggregationFrequencyMonth, StatementStatementDefinitionUpdateParamsAggregationFrequencyQuarter, StatementStatementDefinitionUpdateParamsAggregationFrequencyYear, StatementStatementDefinitionUpdateParamsAggregationFrequencyWholePeriod:
		return true
	}
	return false
}

// A Dimension belonging to a Meter.
type StatementStatementDefinitionUpdateParamsDimension struct {
	// The value of a Dimension to use as a filter. Use "\*" as a wildcard to filter on
	// all Dimension values.
	Filter param.Field[[]string] `json:"filter,required"`
	// The name of the Dimension to target in the Meter.
	Name param.Field[string] `json:"name,required"`
	// The Dimension attribute to target.
	Attributes param.Field[[]string] `json:"attributes"`
	// The unique identifier (UUID) of the Meter containing this Dimension.
	MeterID param.Field[string] `json:"meterId"`
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
	StatementStatementDefinitionUpdateParamsMeasuresAggregationSum       StatementStatementDefinitionUpdateParamsMeasuresAggregation = "SUM"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationMin       StatementStatementDefinitionUpdateParamsMeasuresAggregation = "MIN"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationMax       StatementStatementDefinitionUpdateParamsMeasuresAggregation = "MAX"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationCount     StatementStatementDefinitionUpdateParamsMeasuresAggregation = "COUNT"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationLatest    StatementStatementDefinitionUpdateParamsMeasuresAggregation = "LATEST"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationMean      StatementStatementDefinitionUpdateParamsMeasuresAggregation = "MEAN"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationUnique    StatementStatementDefinitionUpdateParamsMeasuresAggregation = "UNIQUE"
	StatementStatementDefinitionUpdateParamsMeasuresAggregationCustomSql StatementStatementDefinitionUpdateParamsMeasuresAggregation = "CUSTOM_SQL"
)

func (r StatementStatementDefinitionUpdateParamsMeasuresAggregation) IsKnown() bool {
	switch r {
	case StatementStatementDefinitionUpdateParamsMeasuresAggregationSum, StatementStatementDefinitionUpdateParamsMeasuresAggregationMin, StatementStatementDefinitionUpdateParamsMeasuresAggregationMax, StatementStatementDefinitionUpdateParamsMeasuresAggregationCount, StatementStatementDefinitionUpdateParamsMeasuresAggregationLatest, StatementStatementDefinitionUpdateParamsMeasuresAggregationMean, StatementStatementDefinitionUpdateParamsMeasuresAggregationUnique, StatementStatementDefinitionUpdateParamsMeasuresAggregationCustomSql:
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
