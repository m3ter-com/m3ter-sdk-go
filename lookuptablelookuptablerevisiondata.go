// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// LookupTableLookupTableRevisionDataService contains methods and other services
// that help with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLookupTableLookupTableRevisionDataService] method instead.
type LookupTableLookupTableRevisionDataService struct {
	Options                     []option.RequestOption
	LookupTableRevisionDataJobs *LookupTableLookupTableRevisionDataLookupTableRevisionDataJobService
}

// NewLookupTableLookupTableRevisionDataService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewLookupTableLookupTableRevisionDataService(opts ...option.RequestOption) (r *LookupTableLookupTableRevisionDataService) {
	r = &LookupTableLookupTableRevisionDataService{}
	r.Options = opts
	r.LookupTableRevisionDataJobs = NewLookupTableLookupTableRevisionDataLookupTableRevisionDataJobService(opts...)
	return
}

// List Lookup Table Revision Data items for the given UUID.
func (r *LookupTableLookupTableRevisionDataService) Get(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataGetParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataGetResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data", params.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Create/update the Lookup Table Revision Data for the given UUID.
func (r *LookupTableLookupTableRevisionDataService) Update(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataUpdateParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataUpdateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data", params.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete the Lookup Table Revision Data for the given UUID.
func (r *LookupTableLookupTableRevisionDataService) Delete(ctx context.Context, lookupTableID string, lookupTableRevisionID string, body LookupTableLookupTableRevisionDataDeleteParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataDeleteResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data", body.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a URL which you can use to download the data for the specified archived
// Lookup Table Revision:
//
//   - The `contentType` request parameter is required.
//   - The returned URL is presigned - you can copy it into a browser and the data
//     file is downloaded locally.
//   - The upload URL is time limited - the `expiry` time is given in the response
//     and the URL is valid for **_one hour_**.
func (r *LookupTableLookupTableRevisionDataService) Archieve(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataArchieveParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataArchieveResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/archived", params.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Copy the Lookup Table Revision Data from a source Revision to an optional target
// Revision:
//
//   - If you omit a target `revisionId`, then the source Revision and its Data is
//     duplicated. The new Revision is given the source Revision's name appended with
//     "Copy" but is assigned a new unique id.
//   - If you specify a target `revisionId` to copy the source Revision and its Data
//     to, you must ensure that the target Revision has a Data schema that matches
//     the source Revision's Data schema otherwise you'll receive an error
func (r *LookupTableLookupTableRevisionDataService) Copy(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataCopyParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataCopyResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/copy", params.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a Lookup Table Revision Data entry by lookup key.
//
// **NOTES:**
//
//   - To obtain the lookup key for a Revision's data items, use the
//     [Get LookupTableRevisionData](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionData)
//     endpoint in this section and use the `additional=lookupKey` query parameter.
//   - If the Revision's Data schema uses multiple key fields, enter these as a
//     comma-separated list for the `lookupKey` path parameter: .../key1,key2,key3
//     and so on. Importantly, multiple keys must be _entered in the same order_ as
//     they are configured in the Revision's Data schema.
func (r *LookupTableLookupTableRevisionDataService) DeleteKey(ctx context.Context, lookupTableID string, lookupTableRevisionID string, lookupKey string, params LookupTableLookupTableRevisionDataDeleteKeyParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataDeleteKeyResponse, err error) {
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
	if lookupKey == "" {
		err = errors.New("missing required lookupKey parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/%s", params.OrgID, lookupTableID, lookupTableRevisionID, lookupKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Generate a URL which can be used to upload a data file for creating or updating
// the Lookup Table Revision's data:
//
//   - An upload URL is returned together with an UPLOAD `jobId`.
//   - You can then upload your data file using a PUT request using the returned
//     upload URL as the endpoint. For the PUT request, map the headers returned and
//     their values and in the request body select the specified CSV or JSONL file
//     containing the Revision Data to upload.
//   - You can use the returned UPLOAD `jobId` with the
//     [List LookupTableRevisionData Jobs](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/ListLookupTableRevisionDataJobs)
//     or the
//     [Get LookupTableRevisionData Job Response](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionDataJobResponse)
//     endpoints for any follow-up or troubleshooting.
//
// **Important:**
//
// - The `contentLength` request parameter is required.
// - The upload URL is time limited - it is valid for **_one minute_**.
func (r *LookupTableLookupTableRevisionDataService) GenerateDownloadURL(ctx context.Context, lookupTableID string, lookupTableRevisionID string, params LookupTableLookupTableRevisionDataGenerateDownloadURLParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataGenerateDownloadURLResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/generateuploadurl", params.OrgID, lookupTableID, lookupTableRevisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Lookup Table Revision Data item for the given lookup key.
//
// **NOTES:**
//
//   - To obtain the lookup key for a Revision's data items, use the
//     [Get LookupTableRevisionData](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionData)
//     endpoint in this section and use the `additional=lookupKey` query parameter.
//   - If the Revision's Data schema uses multiple key fields, enter these as a
//     comma-separated list for the `lookupKey` path parameter: .../key1,key2,key3
//     and so on. Importantly, multiple keys must be _entered in the same order_ as
//     they are configured in the Revision's Data schema.
func (r *LookupTableLookupTableRevisionDataService) GetKey(ctx context.Context, lookupTableID string, lookupTableRevisionID string, lookupKey string, query LookupTableLookupTableRevisionDataGetKeyParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataGetKeyResponse, err error) {
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
	if lookupKey == "" {
		err = errors.New("missing required lookupKey parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/%s", query.OrgID, lookupTableID, lookupTableRevisionID, lookupKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create/update a Lookup Table Revision Data item by lookup key.
//
// **NOTES:**
//
//   - To obtain the lookup key for a Revision's data items, use the
//     [Get LookupTableRevisionData](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/GetLookupTableRevisionData)
//     endpoint in this section and use the `additional=lookupKey` query parameter.
//   - If the Revision's Data schema uses multiple key fields, enter these as a
//     comma-separated list for the `lookupKey` path parameter: .../key1,key2,key3
//     and so on. Importantly, multiple keys must be _entered in the same order_ as
//     they are configured in the Revision's Data schema.
func (r *LookupTableLookupTableRevisionDataService) UpdateKey(ctx context.Context, lookupTableID string, lookupTableRevisionID string, lookupKey string, params LookupTableLookupTableRevisionDataUpdateKeyParams, opts ...option.RequestOption) (res *LookupTableLookupTableRevisionDataUpdateKeyResponse, err error) {
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
	if lookupKey == "" {
		err = errors.New("missing required lookupKey parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/lookuptables/%s/revisions/%s/data/%s", params.OrgID, lookupTableID, lookupTableRevisionID, lookupKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Response containing data for a Lookup Table Revision
type LookupTableLookupTableRevisionDataGetResponse struct {
	// The Lookup Table Revision Data.
	Items []map[string]interface{} `json:"items" api:"required"`
	// The id of the user who created the Lookup Table Revision Data.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Lookup Table Revision Data was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table Revision Data was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified the Lookup Table Revision Data.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version of the Lookup Table Revision Data.
	Version int64                                             `json:"version"`
	JSON    lookupTableLookupTableRevisionDataGetResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataGetResponseJSON contains the JSON metadata for
// the struct [LookupTableLookupTableRevisionDataGetResponse]
type lookupTableLookupTableRevisionDataGetResponseJSON struct {
	Items          apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataGetResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing data for a Lookup Table Revision
type LookupTableLookupTableRevisionDataUpdateResponse struct {
	// The Lookup Table Revision Data.
	Items []map[string]interface{} `json:"items" api:"required"`
	// The id of the user who created the Lookup Table Revision Data.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Lookup Table Revision Data was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table Revision Data was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified the Lookup Table Revision Data.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version of the Lookup Table Revision Data.
	Version int64                                                `json:"version"`
	JSON    lookupTableLookupTableRevisionDataUpdateResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataUpdateResponseJSON contains the JSON metadata
// for the struct [LookupTableLookupTableRevisionDataUpdateResponse]
type lookupTableLookupTableRevisionDataUpdateResponseJSON struct {
	Items          apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type LookupTableLookupTableRevisionDataDeleteResponse struct {
	// The result of the Lookup Table Revision Data delete action.
	Result string                                               `json:"result"`
	JSON   lookupTableLookupTableRevisionDataDeleteResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataDeleteResponseJSON contains the JSON metadata
// for the struct [LookupTableLookupTableRevisionDataDeleteResponse]
type lookupTableLookupTableRevisionDataDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing the archive job URL details
type LookupTableLookupTableRevisionDataArchieveResponse struct {
	// The URL expiry time
	Expiry time.Time `json:"expiry" format:"date-time"`
	// The presigned URL
	URL  string                                                 `json:"url"`
	JSON lookupTableLookupTableRevisionDataArchieveResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataArchieveResponseJSON contains the JSON
// metadata for the struct [LookupTableLookupTableRevisionDataArchieveResponse]
type lookupTableLookupTableRevisionDataArchieveResponseJSON struct {
	Expiry      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataArchieveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataArchieveResponseJSON) RawJSON() string {
	return r.raw
}

type LookupTableLookupTableRevisionDataCopyResponse struct {
	// UUID of the Revision Data copy job.
	JobID string `json:"jobId"`
	// The ID of the destination Revision.
	RevisionID string                                             `json:"revisionId"`
	JSON       lookupTableLookupTableRevisionDataCopyResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataCopyResponseJSON contains the JSON metadata
// for the struct [LookupTableLookupTableRevisionDataCopyResponse]
type lookupTableLookupTableRevisionDataCopyResponseJSON struct {
	JobID       apijson.Field
	RevisionID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataCopyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataCopyResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing data for a Lookup Table Revision
type LookupTableLookupTableRevisionDataDeleteKeyResponse struct {
	// The Lookup Table Revision Data.
	Items []map[string]interface{} `json:"items" api:"required"`
	// The id of the user who created the Lookup Table Revision Data.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Lookup Table Revision Data was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table Revision Data was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified the Lookup Table Revision Data.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version of the Lookup Table Revision Data.
	Version int64                                                   `json:"version"`
	JSON    lookupTableLookupTableRevisionDataDeleteKeyResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataDeleteKeyResponseJSON contains the JSON
// metadata for the struct [LookupTableLookupTableRevisionDataDeleteKeyResponse]
type lookupTableLookupTableRevisionDataDeleteKeyResponseJSON struct {
	Items          apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataDeleteKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataDeleteKeyResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing the upload job URL details
type LookupTableLookupTableRevisionDataGenerateDownloadURLResponse struct {
	// The headers
	Headers map[string]string `json:"headers"`
	// UUID of the upload job
	JobID string `json:"jobId"`
	// The presigned URL
	URL  string                                                            `json:"url"`
	JSON lookupTableLookupTableRevisionDataGenerateDownloadURLResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataGenerateDownloadURLResponseJSON contains the
// JSON metadata for the struct
// [LookupTableLookupTableRevisionDataGenerateDownloadURLResponse]
type lookupTableLookupTableRevisionDataGenerateDownloadURLResponseJSON struct {
	Headers     apijson.Field
	JobID       apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataGenerateDownloadURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataGenerateDownloadURLResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing data for a Lookup Table Revision
type LookupTableLookupTableRevisionDataGetKeyResponse struct {
	// The Lookup Table Revision Data.
	Items []map[string]interface{} `json:"items" api:"required"`
	// The id of the user who created the Lookup Table Revision Data.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Lookup Table Revision Data was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table Revision Data was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified the Lookup Table Revision Data.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version of the Lookup Table Revision Data.
	Version int64                                                `json:"version"`
	JSON    lookupTableLookupTableRevisionDataGetKeyResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataGetKeyResponseJSON contains the JSON metadata
// for the struct [LookupTableLookupTableRevisionDataGetKeyResponse]
type lookupTableLookupTableRevisionDataGetKeyResponseJSON struct {
	Items          apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataGetKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataGetKeyResponseJSON) RawJSON() string {
	return r.raw
}

// Response containing data for a Lookup Table Revision
type LookupTableLookupTableRevisionDataUpdateKeyResponse struct {
	// The Lookup Table Revision Data.
	Items []map[string]interface{} `json:"items" api:"required"`
	// The id of the user who created the Lookup Table Revision Data.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Lookup Table Revision Data was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Lookup Table Revision Data was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified the Lookup Table Revision Data.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version of the Lookup Table Revision Data.
	Version int64                                                   `json:"version"`
	JSON    lookupTableLookupTableRevisionDataUpdateKeyResponseJSON `json:"-"`
}

// lookupTableLookupTableRevisionDataUpdateKeyResponseJSON contains the JSON
// metadata for the struct [LookupTableLookupTableRevisionDataUpdateKeyResponse]
type lookupTableLookupTableRevisionDataUpdateKeyResponseJSON struct {
	Items          apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LookupTableLookupTableRevisionDataUpdateKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lookupTableLookupTableRevisionDataUpdateKeyResponseJSON) RawJSON() string {
	return r.raw
}

type LookupTableLookupTableRevisionDataGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// Comma separated list of additional fields. For example, you can use
	// `additional=lookupKey` to get the lookup key returned for each Data item. You
	// can then use a lookup key for the Get/Upsert/Delete data entry endpoints in this
	// section.
	Additional param.Field[[]string] `query:"additional"`
	// The maximum number of Data items to return. Defaults to 2000. You can set this
	// to return fewer items if required.
	//
	// If you expect the Revision to contain more than 2000 Data items, you can use the
	// [Trigger Downlad URL Job](https://www.m3ter.com/docs/api#tag/LookupTableRevisionData/operation/TriggerDownloadJob)
	// to download the Lookup Table Revision Data.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [LookupTableLookupTableRevisionDataGetParams]'s query
// parameters as `url.Values`.
func (r LookupTableLookupTableRevisionDataGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupTableLookupTableRevisionDataUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The data for a lookup table revision
	Items param.Field[[]map[string]interface{}] `json:"items" api:"required"`
	// Comma separated list of additional fields. For example, you can use
	// `additional=lookupKey` to get the lookup key returned for each Data item. You
	// can then use a lookup key for the Get/Upsert/Delete data entry endpoints in this
	// section.
	Additional param.Field[[]string] `query:"additional"`
	// The version of the LookupTableRevisionData.
	Version param.Field[int64] `json:"version"`
}

func (r LookupTableLookupTableRevisionDataUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [LookupTableLookupTableRevisionDataUpdateParams]'s query
// parameters as `url.Values`.
func (r LookupTableLookupTableRevisionDataUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LookupTableLookupTableRevisionDataDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type LookupTableLookupTableRevisionDataArchieveParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The content type
	ContentType param.Field[LookupTableLookupTableRevisionDataArchieveParamsContentType] `json:"contentType" api:"required"`
}

func (r LookupTableLookupTableRevisionDataArchieveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The content type
type LookupTableLookupTableRevisionDataArchieveParamsContentType string

const (
	LookupTableLookupTableRevisionDataArchieveParamsContentTypeApplicationJSONL LookupTableLookupTableRevisionDataArchieveParamsContentType = "application/jsonl"
	LookupTableLookupTableRevisionDataArchieveParamsContentTypeTextCsv          LookupTableLookupTableRevisionDataArchieveParamsContentType = "text/csv"
)

func (r LookupTableLookupTableRevisionDataArchieveParamsContentType) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataArchieveParamsContentTypeApplicationJSONL, LookupTableLookupTableRevisionDataArchieveParamsContentTypeTextCsv:
		return true
	}
	return false
}

type LookupTableLookupTableRevisionDataCopyParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The target Revision id that the source Revision's data will be copied to.
	// _(Optional)_
	RevisionID param.Field[string] `json:"revisionId"`
}

func (r LookupTableLookupTableRevisionDataCopyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LookupTableLookupTableRevisionDataDeleteKeyParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The version of the Lookup Table Revision Data.
	Version param.Field[int64] `json:"version"`
}

func (r LookupTableLookupTableRevisionDataDeleteKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LookupTableLookupTableRevisionDataGenerateDownloadURLParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The size of the file body in bytes. For example: `"contentLength": 485`, where
	// 485 is the size in bytes of the file to upload.
	ContentLength param.Field[int64] `json:"contentLength" api:"required"`
	// The content type
	ContentType param.Field[LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentType] `json:"contentType" api:"required"`
	// The name of the file to be uploaded.
	FileName param.Field[string] `json:"fileName" api:"required"`
	// Version of the Lookup Table Revision Data.
	Version param.Field[int64] `json:"version"`
}

func (r LookupTableLookupTableRevisionDataGenerateDownloadURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The content type
type LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentType string

const (
	LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentTypeApplicationJSONL LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentType = "application/jsonl"
	LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentTypeTextCsv          LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentType = "text/csv"
)

func (r LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentType) IsKnown() bool {
	switch r {
	case LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentTypeApplicationJSONL, LookupTableLookupTableRevisionDataGenerateDownloadURLParamsContentTypeTextCsv:
		return true
	}
	return false
}

type LookupTableLookupTableRevisionDataGetKeyParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type LookupTableLookupTableRevisionDataUpdateKeyParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The item you want to upsert
	Item param.Field[map[string]interface{}] `json:"item" api:"required"`
	// Comma separated list of additional fields. For example, you can use
	// `additional=lookupKey` to get the lookup key returned for the Data item.
	Additional param.Field[[]string] `query:"additional"`
	// The version of the LookupTableRevisionData.
	Version param.Field[int64] `json:"version"`
}

func (r LookupTableLookupTableRevisionDataUpdateKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [LookupTableLookupTableRevisionDataUpdateKeyParams]'s query
// parameters as `url.Values`.
func (r LookupTableLookupTableRevisionDataUpdateKeyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
