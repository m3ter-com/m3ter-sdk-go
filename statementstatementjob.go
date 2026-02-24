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

// StatementStatementJobService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementStatementJobService] method instead.
type StatementStatementJobService struct {
	Options []option.RequestOption
}

// NewStatementStatementJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStatementStatementJobService(opts ...option.RequestOption) (r *StatementStatementJobService) {
	r = &StatementStatementJobService{}
	r.Options = opts
	return
}

// This endpoint creates a StatementJob for a single bill within an Organization
// using the Bill UUID.
//
// The Bill Statement is generated asynchronously:
//
//   - The default format for generating the Statement is in JSON format and
//     according to the Bill Statement Definition you've specified at either
//     Organization level or Account level.
//   - If you also want to generate the Statement in CSV format, use the
//     `includeCsvFormat` request body parameter.
//   - The response body provides a time-bound pre-signed URL, which you can use to
//     download the JSON format Statement.
//   - When you have generated a Statement for a Bill, you can also obtain a
//     time-bound pre-signed download URL using either the
//     [Retrieve Bill Statement in JSON Format](https://www.m3ter.com/docs/api#tag/Bill/operation/GetBillJsonStatement)
//     and
//     [Retrieve Bill Statement in CSV Format](https://www.m3ter.com/docs/api#tag/Bill/operation/GetBillCsvStatement)
//     calls found in the [Bill](https://www.m3ter.com/docs/api#tag/Bill) section of
//     this API Reference.
//
// **Notes:**
//
//   - If the response to the Create StatementJob call shows the `statementJobStatus`
//     as `PENDING` or `RUNNING`, you will not receive the pre-signed URL in the
//     response. Wait a few minutes to allow the StatementJob to complete and then
//     use the
//     [Get StatmentJob](https://www.m3ter.com/docs/api#tag/StatementJob/operation/GetStatementJob)
//     call in this section to obtain the pre-signed download URL for the generated
//     Bill Statement.
//   - When you have submitted a StatementJob and a Bill Statement has been
//     generated, you can also download the Statement directly from a Bill Details
//     page in the Console. See
//     [Working with Bill Statements](https://www.m3ter.com/docs/guides/billing-and-usage-data/running-viewing-and-managing-bills/working-with-bill-statements)
//     in our user Documentation.
func (r *StatementStatementJobService) New(ctx context.Context, params StatementStatementJobNewParams, opts ...option.RequestOption) (res *StatementJobResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementjobs", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves the details of a specific StatementJob using its UUID.
//
// Use this call to obtain the time-bound pre-signed download URL for the generated
// Bill Statement if the initial
// [Create StatementJob](https://www.m3ter.com/docs/api#tag/StatementJob/operation/CreateStatementJob)
// returned a response showing the `statementJobStatus` not yet complete and as
// `PENDING` or `RUNNING`.
//
// **Note:** When you have submitted a StatementJob and a Bill Statement has been
// generated, you can also download the Statement directly from a Bill Details page
// in the Console. See
// [Working with Bill Statements](https://www.m3ter.com/docs/guides/billing-and-usage-data/running-viewing-and-managing-bills/working-with-bill-statements)
// in our user Documentation.
func (r *StatementStatementJobService) Get(ctx context.Context, id string, query StatementStatementJobGetParams, opts ...option.RequestOption) (res *StatementJobResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementjobs/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of StatementJobs.
//
// Retrieves a list of all StatementJobs for a specific Organization. You can
// filter the results based on:
//
// - StatementJob status.
// - Whether StatementJob is neither completed nor cancelled but remains active.
// - The ID of the Bill the StatementJob is associated with.
//
// You can also paginate the results for easier management.
//
// **WARNING!**
//
//   - You can use only one of the valid Query parameters: `active`, `status`, or
//     `billId` in any call. If you use more than one of these Query parameters in
//     the same call, then a 400 Bad Request is returned with an error message.
func (r *StatementStatementJobService) List(ctx context.Context, params StatementStatementJobListParams, opts ...option.RequestOption) (res *pagination.Cursor[StatementJobResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/statementjobs", params.OrgID)
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

// Retrieve a list of StatementJobs.
//
// Retrieves a list of all StatementJobs for a specific Organization. You can
// filter the results based on:
//
// - StatementJob status.
// - Whether StatementJob is neither completed nor cancelled but remains active.
// - The ID of the Bill the StatementJob is associated with.
//
// You can also paginate the results for easier management.
//
// **WARNING!**
//
//   - You can use only one of the valid Query parameters: `active`, `status`, or
//     `billId` in any call. If you use more than one of these Query parameters in
//     the same call, then a 400 Bad Request is returned with an error message.
func (r *StatementStatementJobService) ListAutoPaging(ctx context.Context, params StatementStatementJobListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[StatementJobResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Cancel the StatementJob with the given UUID.
//
// Use this endpoint to halt the execution of a specific StatementJob identified by
// its UUID. This operation may be useful if you need to stop a StatementJob due to
// unforeseen issues or changes.
func (r *StatementStatementJobService) Cancel(ctx context.Context, id string, body StatementStatementJobCancelParams, opts ...option.RequestOption) (res *StatementJobResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementjobs/%s/cancel", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create a batch of StatementJobs for multiple bills.
//
// Initiate the creation of multiple StatementJobs asynchronously for the list of
// bills with the given UUIDs:
//
//   - The default format for generating Bill Statements is in JSON format and
//     according to the Bill Statement Definition you've specified at either
//     Organization level or Account level.
//   - If you also want to generate the Statements in CSV format, use the
//     `includeCsvFormat` request body parameter.
//   - The response body provides a time-bound pre-signed URL, which you can use to
//     download the JSON format Statement.
//   - When you have generated a Statement for a Bill, you can also obtain a
//     time-bound pre-signed download URL using either the
//     [Retrieve Bill Statement in JSON Format](https://www.m3ter.com/docs/api#tag/Bill/operation/GetBillJsonStatement)
//     and
//     [Retrieve Bill Statement in CSV Format](https://www.m3ter.com/docs/api#tag/Bill/operation/GetBillCsvStatement)
//     calls found in the [Bill](https://www.m3ter.com/docs/api#tag/Bill) section of
//     this API Reference.
//
// **Notes:**
//
//   - If the response to the Create StatementJob call shows the `statementJobStatus`
//     as `PENDING` or `RUNNING`, you will not receive the pre-signed URL in the
//     response. Wait a few minutes to allow the StatementJob to complete and then
//     use the
//     [Get StatmentJob](https://www.m3ter.com/docs/api#tag/StatementJob/operation/GetStatementJob)
//     call in this section to obtain the pre-signed download URL for the generated
//     Bill Statement.
//   - When you have submitted a StatementJob and a Bill Statement has been
//     generated, you can also download the Statement directly from a Bill Details
//     page in the Console. See
//     [Working with Bill Statements](https://www.m3ter.com/docs/guides/billing-and-usage-data/running-viewing-and-managing-bills/working-with-bill-statements)
//     in our user Documentation.
func (r *StatementStatementJobService) NewBatch(ctx context.Context, params StatementStatementJobNewBatchParams, opts ...option.RequestOption) (res *[]StatementJobResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/statementjobs/batch", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type StatementStatementJobNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The unique identifier (UUID) of the bill associated with the StatementJob.
	BillID  param.Field[string]                                `json:"billId" api:"required"`
	Filters param.Field[StatementStatementJobNewParamsFilters] `json:"filters"`
	// A Boolean value indicating whether the generated statement includes a CSV
	// format.
	//
	// - TRUE - includes the statement in CSV format.
	// - FALSE - no CSV format statement.
	IncludeCsvFormat param.Field[bool] `json:"includeCsvFormat"`
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

func (r StatementStatementJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StatementStatementJobNewParamsFilters struct {
	// Include usage line items whose meterId matches one of these values.
	MeterIDs param.Field[[]string] `json:"meterIds"`
}

func (r StatementStatementJobNewParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StatementStatementJobGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type StatementStatementJobListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Boolean filter on whether to only retrieve active _(i.e. not
	// completed/cancelled)_ StatementJobs.
	//
	// - TRUE - only active StatementJobs retrieved.
	// - FALSE - all StatementJobs retrieved.
	Active param.Field[string] `query:"active"`
	// Filter Statement Jobs by billId
	BillID param.Field[string] `query:"billId"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// StatementJobs in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of StatementJobs to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// Filter using the StatementJobs status. Possible values:
	//
	// - `PENDING`
	// - `RUNNING`
	// - `COMPLETE`
	// - `CANCELLED`
	// - `FAILED`
	Status param.Field[string] `query:"status"`
}

// URLQuery serializes [StatementStatementJobListParams]'s query parameters as
// `url.Values`.
func (r StatementStatementJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StatementStatementJobCancelParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type StatementStatementJobNewBatchParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The list of unique identifiers (UUIDs) of the bills associated with the
	// StatementJob.
	BillIDs param.Field[[]string]                                   `json:"billIds" api:"required"`
	Filters param.Field[StatementStatementJobNewBatchParamsFilters] `json:"filters"`
	// A Boolean value indicating whether the generated statement includes a CSV
	// format.
	//
	// - TRUE - includes the statement in CSV format.
	// - FALSE - no CSV format statement.
	IncludeCsvFormat param.Field[bool] `json:"includeCsvFormat"`
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

func (r StatementStatementJobNewBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StatementStatementJobNewBatchParamsFilters struct {
	// Include usage line items whose meterId matches one of these values.
	MeterIDs param.Field[[]string] `json:"meterIds"`
}

func (r StatementStatementJobNewBatchParamsFilters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
