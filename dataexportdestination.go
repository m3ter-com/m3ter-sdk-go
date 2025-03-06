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

// DataExportDestinationService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataExportDestinationService] method instead.
type DataExportDestinationService struct {
	Options []option.RequestOption
}

// NewDataExportDestinationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDataExportDestinationService(opts ...option.RequestOption) (r *DataExportDestinationService) {
	r = &DataExportDestinationService{}
	r.Options = opts
	return
}

// Create a new Export Destination to use for your Data Export Schedules or Ad-Hoc
// Data Exports.
//
// **NOTE:** Currently, you can only create Export Destinations using an S3 bucket
// on your AWS Account.
func (r *DataExportDestinationService) New(ctx context.Context, params DataExportDestinationNewParams, opts ...option.RequestOption) (res *DataExportDestinationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/destinations", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve an Export Destination for the given UUID.
func (r *DataExportDestinationService) Get(ctx context.Context, id string, query DataExportDestinationGetParams, opts ...option.RequestOption) (res *DataExportDestinationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/destinations/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an Export Destination for the given UUID.
//
// **NOTE:** Currently, only Export Destinations using an S3 bucket on your AWS
// Account are supported.
func (r *DataExportDestinationService) Update(ctx context.Context, id string, params DataExportDestinationUpdateParams, opts ...option.RequestOption) (res *DataExportDestinationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/destinations/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Export Destination entities. You can filter the list of
// Destinations returned by UUID.
func (r *DataExportDestinationService) List(ctx context.Context, params DataExportDestinationListParams, opts ...option.RequestOption) (res *pagination.Cursor[DataExportDestinationResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/destinations", params.OrgID)
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

// Retrieve a list of Export Destination entities. You can filter the list of
// Destinations returned by UUID.
func (r *DataExportDestinationService) ListAutoPaging(ctx context.Context, params DataExportDestinationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DataExportDestinationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete an Export Destination for the given UUID.
//
// **NOTE:** If you attempt to delete an Export Destination that is currently
// linked to a Data Export Schedule, an error message is returned and you won't be
// able to delete the Destination.
func (r *DataExportDestinationService) Delete(ctx context.Context, id string, body DataExportDestinationDeleteParams, opts ...option.RequestOption) (res *DataExportDestinationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/dataexports/destinations/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DataExportDestinationResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The code of the data Export Destination.
	Code string `json:"code"`
	// The id of the user who created the Export Destination.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Export Destination was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Export Destination was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified the Export Destination.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data Export Destination.
	Name string                            `json:"name"`
	JSON dataExportDestinationResponseJSON `json:"-"`
}

// dataExportDestinationResponseJSON contains the JSON metadata for the struct
// [DataExportDestinationResponse]
type dataExportDestinationResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationResponseJSON) RawJSON() string {
	return r.raw
}

type DataExportDestinationNewResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Name of the S3 bucket for the Export Destination.
	BucketName string `json:"bucketName"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// Specify how you want the file path to be structured in your bucket destination -
	// by Time first (Default) or Type first.
	//
	// Type is dependent on whether the Export is for Usage data or Operational data:
	//
	//   - **Usage.** Type is `measurements`.
	//   - **Operational.** Type is one of the entities for which operational data
	//     exports are available, such as `account`, `commitment`, `meter`, and so on.
	//
	// Example for Usage Data Export using .CSV format:
	//
	//   - Time first:
	//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	//   - Type first:
	//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	PartitionOrder DataExportDestinationNewResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                               `json:"prefix"`
	JSON   dataExportDestinationNewResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationNewResponseJSON contains the JSON metadata for the struct
// [DataExportDestinationNewResponse]
type dataExportDestinationNewResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specify how you want the file path to be structured in your bucket destination -
// by Time first (Default) or Type first.
//
// Type is dependent on whether the Export is for Usage data or Operational data:
//
//   - **Usage.** Type is `measurements`.
//   - **Operational.** Type is one of the entities for which operational data
//     exports are available, such as `account`, `commitment`, `meter`, and so on.
//
// Example for Usage Data Export using .CSV format:
//
//   - Time first:
//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
//   - Type first:
//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
type DataExportDestinationNewResponsePartitionOrder string

const (
	DataExportDestinationNewResponsePartitionOrderTypeFirst DataExportDestinationNewResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationNewResponsePartitionOrderTimeFirst DataExportDestinationNewResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationNewResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationNewResponsePartitionOrderTypeFirst, DataExportDestinationNewResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationGetResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Name of the S3 bucket for the Export Destination.
	BucketName string `json:"bucketName"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// Specify how you want the file path to be structured in your bucket destination -
	// by Time first (Default) or Type first.
	//
	// Type is dependent on whether the Export is for Usage data or Operational data:
	//
	//   - **Usage.** Type is `measurements`.
	//   - **Operational.** Type is one of the entities for which operational data
	//     exports are available, such as `account`, `commitment`, `meter`, and so on.
	//
	// Example for Usage Data Export using .CSV format:
	//
	//   - Time first:
	//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	//   - Type first:
	//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	PartitionOrder DataExportDestinationGetResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                               `json:"prefix"`
	JSON   dataExportDestinationGetResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationGetResponseJSON contains the JSON metadata for the struct
// [DataExportDestinationGetResponse]
type dataExportDestinationGetResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Specify how you want the file path to be structured in your bucket destination -
// by Time first (Default) or Type first.
//
// Type is dependent on whether the Export is for Usage data or Operational data:
//
//   - **Usage.** Type is `measurements`.
//   - **Operational.** Type is one of the entities for which operational data
//     exports are available, such as `account`, `commitment`, `meter`, and so on.
//
// Example for Usage Data Export using .CSV format:
//
//   - Time first:
//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
//   - Type first:
//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
type DataExportDestinationGetResponsePartitionOrder string

const (
	DataExportDestinationGetResponsePartitionOrderTypeFirst DataExportDestinationGetResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationGetResponsePartitionOrderTimeFirst DataExportDestinationGetResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationGetResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationGetResponsePartitionOrderTypeFirst, DataExportDestinationGetResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationUpdateResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Name of the S3 bucket for the Export Destination.
	BucketName string `json:"bucketName"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// Specify how you want the file path to be structured in your bucket destination -
	// by Time first (Default) or Type first.
	//
	// Type is dependent on whether the Export is for Usage data or Operational data:
	//
	//   - **Usage.** Type is `measurements`.
	//   - **Operational.** Type is one of the entities for which operational data
	//     exports are available, such as `account`, `commitment`, `meter`, and so on.
	//
	// Example for Usage Data Export using .CSV format:
	//
	//   - Time first:
	//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	//   - Type first:
	//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	PartitionOrder DataExportDestinationUpdateResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                                  `json:"prefix"`
	JSON   dataExportDestinationUpdateResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationUpdateResponseJSON contains the JSON metadata for the
// struct [DataExportDestinationUpdateResponse]
type dataExportDestinationUpdateResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Specify how you want the file path to be structured in your bucket destination -
// by Time first (Default) or Type first.
//
// Type is dependent on whether the Export is for Usage data or Operational data:
//
//   - **Usage.** Type is `measurements`.
//   - **Operational.** Type is one of the entities for which operational data
//     exports are available, such as `account`, `commitment`, `meter`, and so on.
//
// Example for Usage Data Export using .CSV format:
//
//   - Time first:
//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
//   - Type first:
//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
type DataExportDestinationUpdateResponsePartitionOrder string

const (
	DataExportDestinationUpdateResponsePartitionOrderTypeFirst DataExportDestinationUpdateResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationUpdateResponsePartitionOrderTimeFirst DataExportDestinationUpdateResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationUpdateResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateResponsePartitionOrderTypeFirst, DataExportDestinationUpdateResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationDeleteResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Name of the S3 bucket for the Export Destination.
	BucketName string `json:"bucketName"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// Specify how you want the file path to be structured in your bucket destination -
	// by Time first (Default) or Type first.
	//
	// Type is dependent on whether the Export is for Usage data or Operational data:
	//
	//   - **Usage.** Type is `measurements`.
	//   - **Operational.** Type is one of the entities for which operational data
	//     exports are available, such as `account`, `commitment`, `meter`, and so on.
	//
	// Example for Usage Data Export using .CSV format:
	//
	//   - Time first:
	//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	//   - Type first:
	//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	PartitionOrder DataExportDestinationDeleteResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                                  `json:"prefix"`
	JSON   dataExportDestinationDeleteResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationDeleteResponseJSON contains the JSON metadata for the
// struct [DataExportDestinationDeleteResponse]
type dataExportDestinationDeleteResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Specify how you want the file path to be structured in your bucket destination -
// by Time first (Default) or Type first.
//
// Type is dependent on whether the Export is for Usage data or Operational data:
//
//   - **Usage.** Type is `measurements`.
//   - **Operational.** Type is one of the entities for which operational data
//     exports are available, such as `account`, `commitment`, `meter`, and so on.
//
// Example for Usage Data Export using .CSV format:
//
//   - Time first:
//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
//   - Type first:
//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
type DataExportDestinationDeleteResponsePartitionOrder string

const (
	DataExportDestinationDeleteResponsePartitionOrderTypeFirst DataExportDestinationDeleteResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationDeleteResponsePartitionOrderTimeFirst DataExportDestinationDeleteResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationDeleteResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationDeleteResponsePartitionOrderTypeFirst, DataExportDestinationDeleteResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Name of the S3 bucket for the Export Destination.
	BucketName param.Field[string] `json:"bucketName,required"`
	// The code of the Export Destination.
	Code param.Field[string] `json:"code,required"`
	// To enable m3ter to upload a Data Exports to your S3 bucket, the service has to
	// assume an IAM role with PutObject permission for the specified `bucketName`.
	// Create a suitable IAM role in your AWS account and enter ARN:
	//
	// **Formatting Constraints:**
	//
	//   - The IAM role ARN must begin with "arn:aws:iam".
	//   - The general format required is:
	//     "arn:aws:iam::<aws account id>:role/<role name>". For example:
	//     "arn:aws:iam::922609978421:role/IAMRole636".
	//   - If the `iamRoleArn` used doesn't comply with this format, then an error
	//     message will be returned.
	//
	// **More Details:** For more details and examples of the Permission and Trust
	// Policies you can use to create the required IAM Role ARN, see
	// [Creating Data Export Destinations](https://www.m3ter.com/docs/guides/data-exports/creating-data-export-destinations)
	// in our main User documentation.
	IamRoleArn param.Field[string] `json:"iamRoleArn,required"`
	// The name of the Export Destination.
	Name param.Field[string] `json:"name,required"`
	// Specify how you want the file path to be structured in your bucket destination -
	// by Time first (Default) or Type first.
	//
	// Type is dependent on whether the Export is for Usage data or Operational data:
	//
	//   - **Usage.** Type is `measurements`.
	//   - **Operational.** Type is one of the entities for which operational data
	//     exports are available, such as `account`, `commitment`, `meter`, and so on.
	//
	// Example for Usage Data Export using .CSV format:
	//
	//   - Time first:
	//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	//   - Type first:
	//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	PartitionOrder param.Field[DataExportDestinationNewParamsPartitionOrder] `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If you omit a
	// `prefix`, then the root of the bucket will be used.
	Prefix param.Field[string] `json:"prefix"`
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

func (r DataExportDestinationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how you want the file path to be structured in your bucket destination -
// by Time first (Default) or Type first.
//
// Type is dependent on whether the Export is for Usage data or Operational data:
//
//   - **Usage.** Type is `measurements`.
//   - **Operational.** Type is one of the entities for which operational data
//     exports are available, such as `account`, `commitment`, `meter`, and so on.
//
// Example for Usage Data Export using .CSV format:
//
//   - Time first:
//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
//   - Type first:
//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
type DataExportDestinationNewParamsPartitionOrder string

const (
	DataExportDestinationNewParamsPartitionOrderTypeFirst DataExportDestinationNewParamsPartitionOrder = "TYPE_FIRST"
	DataExportDestinationNewParamsPartitionOrderTimeFirst DataExportDestinationNewParamsPartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationNewParamsPartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationNewParamsPartitionOrderTypeFirst, DataExportDestinationNewParamsPartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type DataExportDestinationUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Name of the S3 bucket for the Export Destination.
	BucketName param.Field[string] `json:"bucketName,required"`
	// The code of the Export Destination.
	Code param.Field[string] `json:"code,required"`
	// To enable m3ter to upload a Data Exports to your S3 bucket, the service has to
	// assume an IAM role with PutObject permission for the specified `bucketName`.
	// Create a suitable IAM role in your AWS account and enter ARN:
	//
	// **Formatting Constraints:**
	//
	//   - The IAM role ARN must begin with "arn:aws:iam".
	//   - The general format required is:
	//     "arn:aws:iam::<aws account id>:role/<role name>". For example:
	//     "arn:aws:iam::922609978421:role/IAMRole636".
	//   - If the `iamRoleArn` used doesn't comply with this format, then an error
	//     message will be returned.
	//
	// **More Details:** For more details and examples of the Permission and Trust
	// Policies you can use to create the required IAM Role ARN, see
	// [Creating Data Export Destinations](https://www.m3ter.com/docs/guides/data-exports/creating-data-export-destinations)
	// in our main User documentation.
	IamRoleArn param.Field[string] `json:"iamRoleArn,required"`
	// The name of the Export Destination.
	Name param.Field[string] `json:"name,required"`
	// Specify how you want the file path to be structured in your bucket destination -
	// by Time first (Default) or Type first.
	//
	// Type is dependent on whether the Export is for Usage data or Operational data:
	//
	//   - **Usage.** Type is `measurements`.
	//   - **Operational.** Type is one of the entities for which operational data
	//     exports are available, such as `account`, `commitment`, `meter`, and so on.
	//
	// Example for Usage Data Export using .CSV format:
	//
	//   - Time first:
	//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	//   - Type first:
	//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
	PartitionOrder param.Field[DataExportDestinationUpdateParamsPartitionOrder] `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If you omit a
	// `prefix`, then the root of the bucket will be used.
	Prefix param.Field[string] `json:"prefix"`
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

func (r DataExportDestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how you want the file path to be structured in your bucket destination -
// by Time first (Default) or Type first.
//
// Type is dependent on whether the Export is for Usage data or Operational data:
//
//   - **Usage.** Type is `measurements`.
//   - **Operational.** Type is one of the entities for which operational data
//     exports are available, such as `account`, `commitment`, `meter`, and so on.
//
// Example for Usage Data Export using .CSV format:
//
//   - Time first:
//     `{bucketName}/{prefix}/orgId={orgId}/date=2025-01-27/hour=10/type=measurements/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
//   - Type first:
//     `{bucketName}/{prefix}/orgId={orgId}/type=measurements/date=2025-01-27/hour=10/b9a317a6-860a-40f9-9bf4-e65c44c72c94_measurements.csv.gz`
type DataExportDestinationUpdateParamsPartitionOrder string

const (
	DataExportDestinationUpdateParamsPartitionOrderTypeFirst DataExportDestinationUpdateParamsPartitionOrder = "TYPE_FIRST"
	DataExportDestinationUpdateParamsPartitionOrderTimeFirst DataExportDestinationUpdateParamsPartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationUpdateParamsPartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateParamsPartitionOrderTypeFirst, DataExportDestinationUpdateParamsPartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// List of Export Destination UUIDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of returned Export Destinations to list per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [DataExportDestinationListParams]'s query parameters as
// `url.Values`.
func (r DataExportDestinationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataExportDestinationDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
