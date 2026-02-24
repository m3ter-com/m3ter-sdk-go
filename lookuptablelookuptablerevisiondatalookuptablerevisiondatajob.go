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

// LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService contains
// methods and other services that help with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupTableLookupTableRevisionDataLookupTableRevisionDataJobService]
// method instead.
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService struct {
	Options []option.RequestOption
}

// NewLookupTableLookupTableRevisionDataLookupTableRevisionDataJobService generates
// a new service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewLookupTableLookupTableRevisionDataLookupTableRevisionDataJobService(opts ...option.RequestOption) (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService) {
	r = &LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService{}
	r.Options = opts
	return
}

// Get the Lookup Table Revision Data job Response for given job id.
//
// **NOTE:** Use the
// [List LookupTableRevisionData Jobs](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/ListLookupTableRevisionDataJobs)
// endpoint to list the Data job Responses for a specific Revision.
func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService) Get(ctx context.Context, lookupTableID string, lookupTableRevisionID string, id string, query LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if lookupTableRevisionID == "" {
		err = errors.New("missing required lookupTableRevisionId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/jobs/%s", query.OrgID, lookupTableID, lookupTableRevisionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the Lookup Table Revision Data job Responses for the given Lookup Table
// Revision.
//
// There are four types of Revision Data jobs:
//
//   - **COPY**. Job runs when you use the
//     [Copy LookupTableRevisionData](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/CopyLookupTableRevisionData)
//     endpoint which returns the `jobId`.
//   - **UPLOAD**. Job runs when you use the
//     [Generate LookupTableRevisionData Upload URL](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GenerateLookupTableDataUploadUrl)
//     endpoint which returns the `jobId`.
//   - **DOWNLOAD**. Job runs when you use the
//     [](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/TriggerLookupTableRevisionDataDownloadJob)
//     endpoint which returns the `jobId`.
//   - **ARCHIVE**. Job runs when you either manually change a DRAFT Revision to
//     PUBLISHED using the
//     [Update LookupTableRevision Status](https://www.m3ter.com/docs/api#tag/LookupTableRevision/operation/UpdateLookupTableRevisionStatus)
//     endpoint or you publish a DRAFT Revision and the existing PUBLISHED Revision
//     is archived.
//
// **NOTE:** This endpoint returns the id of each Data job. You then use:
//
//   - The
//     [Get LookupTableRevisionData Job Response](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionDataJobResponse)
//     endpoint to retrieve a specific Data job Response.
//   - The
//     [Delete LookupTableRevisionData Job Response](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/DeleteLookupTableRevisionDataJobResponse)
//     to delete a specific Data job Response.
func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService) List(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListParams, opts ...option.RequestOption) (res *pagination.Cursor[LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponse], err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if lookupTableRevisionID == "" {
		err = errors.New("missing required lookupTableRevisionId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/jobs", params.OrgID, lookupTableID, lookupTableRevisionID)
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

// List the Lookup Table Revision Data job Responses for the given Lookup Table
// Revision.
//
// There are four types of Revision Data jobs:
//
//   - **COPY**. Job runs when you use the
//     [Copy LookupTableRevisionData](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/CopyLookupTableRevisionData)
//     endpoint which returns the `jobId`.
//   - **UPLOAD**. Job runs when you use the
//     [Generate LookupTableRevisionData Upload URL](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GenerateLookupTableDataUploadUrl)
//     endpoint which returns the `jobId`.
//   - **DOWNLOAD**. Job runs when you use the
//     [](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/TriggerLookupTableRevisionDataDownloadJob)
//     endpoint which returns the `jobId`.
//   - **ARCHIVE**. Job runs when you either manually change a DRAFT Revision to
//     PUBLISHED using the
//     [Update LookupTableRevision Status](https://www.m3ter.com/docs/api#tag/LookupTableRevision/operation/UpdateLookupTableRevisionStatus)
//     endpoint or you publish a DRAFT Revision and the existing PUBLISHED Revision
//     is archived.
//
// **NOTE:** This endpoint returns the id of each Data job. You then use:
//
//   - The
//     [Get LookupTableRevisionData Job Response](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionDataJobResponse)
//     endpoint to retrieve a specific Data job Response.
//   - The
//     [Delete LookupTableRevisionData Job Response](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/DeleteLookupTableRevisionDataJobResponse)
//     to delete a specific Data job Response.
func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService) ListAutoPaging(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, lookupTableID, lookupTableRevisionID, params, opts...))
}

// Delete the LookupTableRevisionData Job Response for given job id.
//
// **NOTE:** Use the
// [List LookupTableRevisionData Jobs](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/ListLookupTableRevisionDataJobs)
// endpoint to list the Data job Responses for a specific Revision.
func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService) Delete(ctx context.Context, lookupTableID string, lookupTableRevisionID string, id string, body LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if lookupTableRevisionID == "" {
		err = errors.New("missing required lookupTableRevisionId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/jobs/%s", body.OrgID, lookupTableID, lookupTableRevisionID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Trigger an URL job to download the Lookup Table Revision Data. The URL download
// Data `jobId` is returned and you can then use the
// [List LookupTableRevisionData Jobs](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/ListLookupTableRevisionDataJobs)
// endpoint or the
// [Get LookupTableRevisionData Job Response](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionDataJobResponse)
// endpoint to retrieve the URL and perform the Revision data Download.
func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService) Download(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponse, err error) {
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
	if lookupTableID == "" {
		err = errors.New("missing required lookupTableId parameter")
		return
	}
	if lookupTableRevisionID == "" {
		err = errors.New("missing required lookupTableRevisionId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/jobs/download", params.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Response containing the LookupTableRevisionData job details
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponse struct {
	// UUID of the Lookup Table Revision Data job.
	ID string `json:"id"`
	// UUID of the destination Lookup Table Revision if the Lookup Table Revision Data
	// job is a COPY job.
	DestinationLookupTableRevisionID string `json:"destinationLookupTableRevisionId"`
	// The download Url if the Lookup Table Revision Data job is a DOWNLOAD job.
	DownloadURL string `json:"downloadUrl"`
	// The download Url expiry if the Lookup Table Revision Data job is a DOWNLOAD job.
	DownloadURLExpiry string `json:"downloadUrlExpiry"`
	// The failure reason if the Lookup Table Revision Data job failed.
	FailureReason string `json:"failureReason"`
	// The file name for a Lookup Table Revision Data UPLOAD or DOWNLOAD job.
	FileName string `json:"fileName"`
	// The date when the Lookup Table Revision Data job was created.
	JobDate string `json:"jobDate"`
	// UUID of the Lookup Table.
	LookupTableID string `json:"lookupTableId"`
	// UUID of the Lookup Table Revision.
	LookupTableRevisionID string `json:"lookupTableRevisionId"`
	// The status of a job
	Status LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatus `json:"status"`
	Type   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType   `json:"type"`
	// Version of the Lookup Table Revision Data job.
	Version int64                                                                       `json:"version"`
	JSON    lookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseJSON
// contains the JSON metadata for the struct
// [LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponse]
type lookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseJSON struct {
	ID                               apijson.Field
	DestinationLookupTableRevisionID apijson.Field
	DownloadURL                      apijson.Field
	DownloadURLExpiry                apijson.Field
	FailureReason                    apijson.Field
	FileName                         apijson.Field
	JobDate                          apijson.Field
	LookupTableID                    apijson.Field
	LookupTableRevisionID            apijson.Field
	Status                           apijson.Field
	Type                             apijson.Field
	Version                          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseJSON) RawJSON() string {
	return r.raw
}

// The status of a job
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatus string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatusPending   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatus = "PENDING"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatusFailed    LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatus = "FAILED"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatusSucceeded LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatus = "SUCCEEDED"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatus) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatusPending, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatusFailed, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseStatusSucceeded:
		return true
	}
	return false
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeCopy     LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType = "COPY"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeUpload   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType = "UPLOAD"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeDownload LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType = "DOWNLOAD"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeArchive  LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType = "ARCHIVE"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseType) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeCopy, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeUpload, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeDownload, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetResponseTypeArchive:
		return true
	}
	return false
}

// Response containing the LookupTableRevisionData job details
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponse struct {
	// UUID of the Lookup Table Revision Data job.
	ID string `json:"id"`
	// UUID of the destination Lookup Table Revision if the Lookup Table Revision Data
	// job is a COPY job.
	DestinationLookupTableRevisionID string `json:"destinationLookupTableRevisionId"`
	// The download Url if the Lookup Table Revision Data job is a DOWNLOAD job.
	DownloadURL string `json:"downloadUrl"`
	// The download Url expiry if the Lookup Table Revision Data job is a DOWNLOAD job.
	DownloadURLExpiry string `json:"downloadUrlExpiry"`
	// The failure reason if the Lookup Table Revision Data job failed.
	FailureReason string `json:"failureReason"`
	// The file name for a Lookup Table Revision Data UPLOAD or DOWNLOAD job.
	FileName string `json:"fileName"`
	// The date when the Lookup Table Revision Data job was created.
	JobDate string `json:"jobDate"`
	// UUID of the Lookup Table.
	LookupTableID string `json:"lookupTableId"`
	// UUID of the Lookup Table Revision.
	LookupTableRevisionID string `json:"lookupTableRevisionId"`
	// The status of a job
	Status LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatus `json:"status"`
	Type   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType   `json:"type"`
	// Version of the Lookup Table Revision Data job.
	Version int64                                                                        `json:"version"`
	JSON    lookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseJSON
// contains the JSON metadata for the struct
// [LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponse]
type lookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseJSON struct {
	ID                               apijson.Field
	DestinationLookupTableRevisionID apijson.Field
	DownloadURL                      apijson.Field
	DownloadURLExpiry                apijson.Field
	FailureReason                    apijson.Field
	FileName                         apijson.Field
	JobDate                          apijson.Field
	LookupTableID                    apijson.Field
	LookupTableRevisionID            apijson.Field
	Status                           apijson.Field
	Type                             apijson.Field
	Version                          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseJSON) RawJSON() string {
	return r.raw
}

// The status of a job
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatus string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatusPending   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatus = "PENDING"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatusFailed    LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatus = "FAILED"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatusSucceeded LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatus = "SUCCEEDED"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatus) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatusPending, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatusFailed, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseStatusSucceeded:
		return true
	}
	return false
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeCopy     LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType = "COPY"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeUpload   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType = "UPLOAD"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeDownload LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType = "DOWNLOAD"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeArchive  LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType = "ARCHIVE"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseType) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeCopy, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeUpload, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeDownload, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListResponseTypeArchive:
		return true
	}
	return false
}

// Response containing the LookupTableRevisionData job details
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponse struct {
	// UUID of the Lookup Table Revision Data job.
	ID string `json:"id"`
	// UUID of the destination Lookup Table Revision if the Lookup Table Revision Data
	// job is a COPY job.
	DestinationLookupTableRevisionID string `json:"destinationLookupTableRevisionId"`
	// The download Url if the Lookup Table Revision Data job is a DOWNLOAD job.
	DownloadURL string `json:"downloadUrl"`
	// The download Url expiry if the Lookup Table Revision Data job is a DOWNLOAD job.
	DownloadURLExpiry string `json:"downloadUrlExpiry"`
	// The failure reason if the Lookup Table Revision Data job failed.
	FailureReason string `json:"failureReason"`
	// The file name for a Lookup Table Revision Data UPLOAD or DOWNLOAD job.
	FileName string `json:"fileName"`
	// The date when the Lookup Table Revision Data job was created.
	JobDate string `json:"jobDate"`
	// UUID of the Lookup Table.
	LookupTableID string `json:"lookupTableId"`
	// UUID of the Lookup Table Revision.
	LookupTableRevisionID string `json:"lookupTableRevisionId"`
	// The status of a job
	Status LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatus `json:"status"`
	Type   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType   `json:"type"`
	// Version of the Lookup Table Revision Data job.
	Version int64                                                                          `json:"version"`
	JSON    lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseJSON
// contains the JSON metadata for the struct
// [LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponse]
type lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseJSON struct {
	ID                               apijson.Field
	DestinationLookupTableRevisionID apijson.Field
	DownloadURL                      apijson.Field
	DownloadURLExpiry                apijson.Field
	FailureReason                    apijson.Field
	FileName                         apijson.Field
	JobDate                          apijson.Field
	LookupTableID                    apijson.Field
	LookupTableRevisionID            apijson.Field
	Status                           apijson.Field
	Type                             apijson.Field
	Version                          apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The status of a job
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatus string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatusPending   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatus = "PENDING"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatusFailed    LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatus = "FAILED"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatusSucceeded LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatus = "SUCCEEDED"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatus) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatusPending, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatusFailed, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseStatusSucceeded:
		return true
	}
	return false
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeCopy     LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType = "COPY"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeUpload   LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType = "UPLOAD"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeDownload LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType = "DOWNLOAD"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeArchive  LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType = "ARCHIVE"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseType) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeCopy, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeUpload, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeDownload, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteResponseTypeArchive:
		return true
	}
	return false
}

// Response containing the download job details
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponse struct {
	// UUID of the download job
	JobID string                                                                           `json:"jobId"`
	JSON  lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponseJSON
// contains the JSON metadata for the struct
// [LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponse]
type lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponseJSON struct {
	JobID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadResponseJSON) RawJSON() string {
	return r.raw
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// The number of Lookup Table Revision Data Job Responses to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes
// [LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListParams]'s query
// parameters as `url.Values`.
func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The content type
	ContentType param.Field[LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentType] `json:"contentType" api:"required"`
}

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The content type
type LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentType string

const (
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentTypeApplicationJSONL LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentType = "application/jsonl"
	LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentTypeTextCsv          LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentType = "text/csv"
)

func (r LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentType) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentTypeApplicationJSONL, LookupTableLookupTableRevisionDataLookupTableRevisionDataJobDownloadParamsContentTypeTextCsv:
		return true
	}
	return false
}
