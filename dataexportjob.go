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

// DataExportJobService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataExportJobService] method instead.
type DataExportJobService struct {
	Options []option.RequestOption
}

// NewDataExportJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataExportJobService(opts ...option.RequestOption) (r *DataExportJobService) {
	r = &DataExportJobService{}
	r.Options = opts
	return
}

// Retrieve an Export Job for the given UUID.
//
// The response returns:
//
//   - The source type for the data exported by the Export Job: one of USAGE or
//     OPERATIONAL.
//   - The status of the Export Job.
func (r *DataExportJobService) Get(ctx context.Context, id string, query DataExportJobGetParams, opts ...option.RequestOption) (res *DataExportJob, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/jobs/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of Export Job entities.
func (r *DataExportJobService) List(ctx context.Context, params DataExportJobListParams, opts ...option.RequestOption) (res *pagination.Cursor[DataExportJob], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/jobs", params.OrgID)
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

// Retrieve a list of Export Job entities.
func (r *DataExportJobService) ListAutoPaging(ctx context.Context, params DataExportJobListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DataExportJob] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Returns a presigned download URL for data export file download based on the
// `jobId` provided.
//
// If you omit `destinationIds` when setting up your
// [Ad-Hoc data exports](https://www.m3ter.com/docs/api#tag/ExportAdHoc) or
// [Scheduled data exports](https://www.m3ter.com/docs/api#tag/ExportSchedule),
// then the data is not copied to a destination but is available for you to
// download using the returned download URL.
//
// **Constraints:**
//
// - Only valid for Export jobs ran in the past 24 hours.
// - The download URL is time-bound and is only valid for 15 minutes.
//
// **NOTE!** This ExportDestination endpoint is available in Beta release version.
// Beta release features are functional but may be incomplete, and there is no
// commitment at this stage to particular functionality or timelines.
func (r *DataExportJobService) GetDownloadURL(ctx context.Context, jobID string, query DataExportJobGetDownloadURLParams, opts ...option.RequestOption) (res *DataExportJobGetDownloadURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/jobs/%s/getdownloadurl", query.OrgID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DataExportJob struct {
	// The id of the Export Job.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// When the data Export Job was created.
	DateCreated time.Time `json:"dateCreated" format:"date-time"`
	// The id of the data Export Schedule.
	ScheduleID string                  `json:"scheduleId"`
	SourceType DataExportJobSourceType `json:"sourceType"`
	// When the data Export Job started running
	StartedAt time.Time           `json:"startedAt" format:"date-time"`
	Status    DataExportJobStatus `json:"status"`
	JSON      dataExportJobJSON   `json:"-"`
}

// dataExportJobJSON contains the JSON metadata for the struct [DataExportJob]
type dataExportJobJSON struct {
	ID          apijson.Field
	Version     apijson.Field
	DateCreated apijson.Field
	ScheduleID  apijson.Field
	SourceType  apijson.Field
	StartedAt   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataExportJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportJobJSON) RawJSON() string {
	return r.raw
}

type DataExportJobSourceType string

const (
	DataExportJobSourceTypeUsage       DataExportJobSourceType = "USAGE"
	DataExportJobSourceTypeOperational DataExportJobSourceType = "OPERATIONAL"
)

func (r DataExportJobSourceType) IsKnown() bool {
	switch r {
	case DataExportJobSourceTypeUsage, DataExportJobSourceTypeOperational:
		return true
	}
	return false
}

type DataExportJobStatus string

const (
	DataExportJobStatusPending   DataExportJobStatus = "PENDING"
	DataExportJobStatusRunning   DataExportJobStatus = "RUNNING"
	DataExportJobStatusSucceeded DataExportJobStatus = "SUCCEEDED"
	DataExportJobStatusFailed    DataExportJobStatus = "FAILED"
)

func (r DataExportJobStatus) IsKnown() bool {
	switch r {
	case DataExportJobStatusPending, DataExportJobStatusRunning, DataExportJobStatusSucceeded, DataExportJobStatusFailed:
		return true
	}
	return false
}

// It contains details for downloading an export file
type DataExportJobGetDownloadURLResponse struct {
	// The expiration time of the URL
	ExpirationTime time.Time `json:"expirationTime" format:"date-time"`
	// The presigned download URL
	URL  string                                  `json:"url"`
	JSON dataExportJobGetDownloadURLResponseJSON `json:"-"`
}

// dataExportJobGetDownloadURLResponseJSON contains the JSON metadata for the
// struct [DataExportJobGetDownloadURLResponse]
type dataExportJobGetDownloadURLResponseJSON struct {
	ExpirationTime apijson.Field
	URL            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportJobGetDownloadURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportJobGetDownloadURLResponseJSON) RawJSON() string {
	return r.raw
}

type DataExportJobGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type DataExportJobListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Include only Job entities created before this date. Format:
	// yyyy-MM-dd'T'HH:mm:ss'Z'
	DateCreatedEnd param.Field[string] `query:"dateCreatedEnd"`
	// Include only Job entities created on or after this date. Format:
	// yyyy-MM-dd'T'HH:mm:ss'Z'
	DateCreatedStart param.Field[string] `query:"dateCreatedStart"`
	// List Job entities for the given UUIDs
	IDs param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Jobs to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// List Job entities for the schedule UUID
	ScheduleID param.Field[string] `query:"scheduleId"`
	// List Job entities for the status
	Status param.Field[DataExportJobListParamsStatus] `query:"status"`
}

// URLQuery serializes [DataExportJobListParams]'s query parameters as
// `url.Values`.
func (r DataExportJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// List Job entities for the status
type DataExportJobListParamsStatus string

const (
	DataExportJobListParamsStatusPending   DataExportJobListParamsStatus = "PENDING"
	DataExportJobListParamsStatusRunning   DataExportJobListParamsStatus = "RUNNING"
	DataExportJobListParamsStatusSucceeded DataExportJobListParamsStatus = "SUCCEEDED"
	DataExportJobListParamsStatusFailed    DataExportJobListParamsStatus = "FAILED"
)

func (r DataExportJobListParamsStatus) IsKnown() bool {
	switch r {
	case DataExportJobListParamsStatusPending, DataExportJobListParamsStatusRunning, DataExportJobListParamsStatusSucceeded, DataExportJobListParamsStatusFailed:
		return true
	}
	return false
}

type DataExportJobGetDownloadURLParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
