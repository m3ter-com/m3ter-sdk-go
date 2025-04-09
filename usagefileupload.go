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

// UsageFileUploadService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageFileUploadService] method instead.
type UsageFileUploadService struct {
	Options []option.RequestOption
	Jobs    *UsageFileUploadJobService
}

// NewUsageFileUploadService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsageFileUploadService(opts ...option.RequestOption) (r *UsageFileUploadService) {
	r = &UsageFileUploadService{}
	r.Options = opts
	r.Jobs = NewUsageFileUploadJobService(opts...)
	return
}

// Generate a URL for uploading a file containing measurements to the platform in
// preparation for the measurements it contains to be ingested:
//
//   - An upload URL is returned together with an upload job id:
//   - You can then upload your data measurements file using a `PUT` request using
//     the returned upload URL as the endpoint.
//   - You can use the returned upload job id with other calls to the File Upload
//     Service for any follow-up or troubleshooting.
//
// **Important:**
//
// - The `contentLength` request parameter is required.
// - The upload URL is time limited - it is valid for **_one_** minute.
//
// Part of the file upload service for submitting measurements data files.
func (r *UsageFileUploadService) GenerateUploadURL(ctx context.Context, params UsageFileUploadGenerateUploadURLParams, opts ...option.RequestOption) (res *UsageFileUploadGenerateUploadURLResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/fileuploads/measurements/generateUploadUrl", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Response containing the upload job URL details
type UsageFileUploadGenerateUploadURLResponse struct {
	// The headers
	Headers map[string]string `json:"headers"`
	// UUID of the upload job
	JobID string `json:"jobId"`
	// The URL
	URL  string                                       `json:"url"`
	JSON usageFileUploadGenerateUploadURLResponseJSON `json:"-"`
}

// usageFileUploadGenerateUploadURLResponseJSON contains the JSON metadata for the
// struct [UsageFileUploadGenerateUploadURLResponse]
type usageFileUploadGenerateUploadURLResponseJSON struct {
	Headers     apijson.Field
	JobID       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageFileUploadGenerateUploadURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageFileUploadGenerateUploadURLResponseJSON) RawJSON() string {
	return r.raw
}

type UsageFileUploadGenerateUploadURLParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The media type of the entity body sent, for example:
	// `"contentType":"text/json"`.
	//
	// **NOTE:** Currently only a JSON formatted file type is supported by the File
	// Upload Service.
	ContentType param.Field[string] `json:"contentType,required"`
	// The name of the measurements file to be uploaded.
	FileName param.Field[string] `json:"fileName,required"`
	// The size of the body in bytes. For example: `"contentLength": 485`, where 485 is
	// the size in bytes of the file to upload.
	//
	// **NOTE:** Required.
	ContentLength param.Field[int64] `json:"contentLength"`
}

func (r UsageFileUploadGenerateUploadURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
