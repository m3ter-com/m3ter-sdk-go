// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
)

// UsageFileUploadJobService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageFileUploadJobService] method instead.
type UsageFileUploadJobService struct {
	Options []option.RequestOption
}

// NewUsageFileUploadJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUsageFileUploadJobService(opts ...option.RequestOption) (r *UsageFileUploadJobService) {
	r = &UsageFileUploadJobService{}
	r.Options = opts
	return
}

// Get the file upload job response using the UUID of the file upload job.
//
// Part of the file upload service for measurements ingest.
func (r *UsageFileUploadJobService) Get(ctx context.Context, id string, query UsageFileUploadJobGetParams, opts ...option.RequestOption) (res *FileUploadJobResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/fileuploads/measurements/jobs/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the File Upload jobs. Part of the File Upload service for measurements
// ingest:
//
//   - You can use the `dateCreatedStart` and `dateCreatedEnd` optional Query
//     parameters to define a date range to filter the File Uploads jobs returned for
//     this call.
//   - If `dateCreatedStart` and `dateCreatedEnd` Query parameters are not used, then
//     all File Upload jobs are returned.
func (r *UsageFileUploadJobService) List(ctx context.Context, params UsageFileUploadJobListParams, opts ...option.RequestOption) (res *pagination.Cursor[FileUploadJobResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/fileuploads/measurements/jobs", params.OrgID)
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

// Lists the File Upload jobs. Part of the File Upload service for measurements
// ingest:
//
//   - You can use the `dateCreatedStart` and `dateCreatedEnd` optional Query
//     parameters to define a date range to filter the File Uploads jobs returned for
//     this call.
//   - If `dateCreatedStart` and `dateCreatedEnd` Query parameters are not used, then
//     all File Upload jobs are returned.
func (r *UsageFileUploadJobService) ListAutoPaging(ctx context.Context, params UsageFileUploadJobListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[FileUploadJobResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Use the original file upload job id to obtain a download URL, which you can then
// use to retrieve the file you originally uploaded to the file upload service:
//
//   - A download URL is returned together with a download job id.
//   - You can then use a `GET` using the returned download URL as the endpoint to
//     retrieve the file you originally uploaded.
//
// Part of the file upload service for submitting measurements data files.
func (r *UsageFileUploadJobService) GetOriginalDownloadURL(ctx context.Context, id string, query UsageFileUploadJobGetOriginalDownloadURLParams, opts ...option.RequestOption) (res *UsageFileUploadJobGetOriginalDownloadURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/fileuploads/measurements/jobs/%s/original", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing the upload job details.
type FileUploadJobResponse struct {
	// UUID of the file upload job.
	ID string `json:"id"`
	// The size of the body in bytes. For example: `"contentLength": 485`, where 485 is
	// the size in bytes of the file uploaded.
	ContentLength int64 `json:"contentLength"`
	// The number of rows that failed processing during ingest.
	FailedRows int64 `json:"failedRows"`
	// The name of the measurements file for the upload job.
	FileName string `json:"fileName"`
	// The number of rows that were processed during ingest.
	ProcessedRows int64 `json:"processedRows"`
	// The status of the file upload job.
	Status FileUploadJobResponseStatus `json:"status"`
	// The total number of rows in the file.
	TotalRows int64 `json:"totalRows"`
	// The upload date for the upload job _(in ISO-8601 format)_.
	UploadDate string `json:"uploadDate"`
	// The version number. Default value when newly created is one.
	Version int64                     `json:"version"`
	JSON    fileUploadJobResponseJSON `json:"-"`
}

// fileUploadJobResponseJSON contains the JSON metadata for the struct
// [FileUploadJobResponse]
type fileUploadJobResponseJSON struct {
	ID            apijson.Field
	ContentLength apijson.Field
	FailedRows    apijson.Field
	FileName      apijson.Field
	ProcessedRows apijson.Field
	Status        apijson.Field
	TotalRows     apijson.Field
	UploadDate    apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FileUploadJobResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileUploadJobResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the file upload job.
type FileUploadJobResponseStatus string

const (
	FileUploadJobResponseStatusNotUploaded FileUploadJobResponseStatus = "notUploaded"
	FileUploadJobResponseStatusRunning     FileUploadJobResponseStatus = "running"
	FileUploadJobResponseStatusFailed      FileUploadJobResponseStatus = "failed"
	FileUploadJobResponseStatusSucceeded   FileUploadJobResponseStatus = "succeeded"
)

func (r FileUploadJobResponseStatus) IsKnown() bool {
	switch r {
	case FileUploadJobResponseStatusNotUploaded, FileUploadJobResponseStatusRunning, FileUploadJobResponseStatusFailed, FileUploadJobResponseStatusSucceeded:
		return true
	}
	return false
}

// It contains details for downloading a file
type UsageFileUploadJobGetOriginalDownloadURLResponse struct {
	// The headers
	Headers map[string]string `json:"headers"`
	// UUID of the download job
	JobID string `json:"jobId"`
	// The URL
	URL  string                                               `json:"url"`
	JSON usageFileUploadJobGetOriginalDownloadURLResponseJSON `json:"-"`
}

// usageFileUploadJobGetOriginalDownloadURLResponseJSON contains the JSON metadata
// for the struct [UsageFileUploadJobGetOriginalDownloadURLResponse]
type usageFileUploadJobGetOriginalDownloadURLResponseJSON struct {
	Headers     apijson.Field
	JobID       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageFileUploadJobGetOriginalDownloadURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageFileUploadJobGetOriginalDownloadURLResponseJSON) RawJSON() string {
	return r.raw
}

type UsageFileUploadJobGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type UsageFileUploadJobListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Include only File Upload jobs created before this date. Required format is
	// ISO-8601: yyyy-MM-dd'T'HH:mm:ss'Z'
	DateCreatedEnd param.Field[string] `query:"dateCreatedEnd"`
	// Include only File Upload jobs created on or after this date. Required format is
	// ISO-8601: yyyy-MM-dd'T'HH:mm:ss'Z'
	DateCreatedStart param.Field[string] `query:"dateCreatedStart"`
	// <<deprecated>>
	FileKey param.Field[string] `query:"fileKey"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of File Upload jobs to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [UsageFileUploadJobListParams]'s query parameters as
// `url.Values`.
func (r UsageFileUploadJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UsageFileUploadJobGetOriginalDownloadURLParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
