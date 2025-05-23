// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
	"github.com/tidwall/gjson"
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
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
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
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
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
	path := fmt.Sprintf("organizations/%s/dataexports/destinations/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DataExportDestinationGoogleCloudStorageRequestParam struct {
	// The export destination bucket name.
	BucketName param.Field[string] `json:"bucketName,required"`
	// The export destination Web Identity Federation poolId.
	PoolID param.Field[string] `json:"poolId,required"`
	// The export destination GCP projectNumber.
	ProjectNumber param.Field[string] `json:"projectNumber,required"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID param.Field[string] `json:"providerId,required"`
	// The type of destination to create. Possible values are: GCS
	DestinationType param.Field[DataExportDestinationGoogleCloudStorageRequestDestinationType] `json:"destinationType"`
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
	PartitionOrder param.Field[DataExportDestinationGoogleCloudStorageRequestPartitionOrder] `json:"partitionOrder"`
	// The export destination prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The export destination service account email.
	ServiceAccountEmail param.Field[string] `json:"serviceAccountEmail"`
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

func (r DataExportDestinationGoogleCloudStorageRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportDestinationGoogleCloudStorageRequestParam) implementsDataExportDestinationNewParamsBodyUnion() {
}

func (r DataExportDestinationGoogleCloudStorageRequestParam) implementsDataExportDestinationUpdateParamsBodyUnion() {
}

// The type of destination to create. Possible values are: GCS
type DataExportDestinationGoogleCloudStorageRequestDestinationType string

const (
	DataExportDestinationGoogleCloudStorageRequestDestinationTypeGcs DataExportDestinationGoogleCloudStorageRequestDestinationType = "GCS"
)

func (r DataExportDestinationGoogleCloudStorageRequestDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationGoogleCloudStorageRequestDestinationTypeGcs:
		return true
	}
	return false
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
type DataExportDestinationGoogleCloudStorageRequestPartitionOrder string

const (
	DataExportDestinationGoogleCloudStorageRequestPartitionOrderTypeFirst DataExportDestinationGoogleCloudStorageRequestPartitionOrder = "TYPE_FIRST"
	DataExportDestinationGoogleCloudStorageRequestPartitionOrderTimeFirst DataExportDestinationGoogleCloudStorageRequestPartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationGoogleCloudStorageRequestPartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationGoogleCloudStorageRequestPartitionOrderTypeFirst, DataExportDestinationGoogleCloudStorageRequestPartitionOrderTimeFirst:
		return true
	}
	return false
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
	CreatedBy       string                                       `json:"createdBy"`
	DestinationType DataExportDestinationResponseDestinationType `json:"destinationType"`
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
	ID              apijson.Field
	Version         apijson.Field
	Code            apijson.Field
	CreatedBy       apijson.Field
	DestinationType apijson.Field
	DtCreated       apijson.Field
	DtLastModified  apijson.Field
	LastModifiedBy  apijson.Field
	Name            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DataExportDestinationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationResponseJSON) RawJSON() string {
	return r.raw
}

type DataExportDestinationResponseDestinationType string

const (
	DataExportDestinationResponseDestinationTypeS3  DataExportDestinationResponseDestinationType = "S3"
	DataExportDestinationResponseDestinationTypeGcs DataExportDestinationResponseDestinationType = "GCS"
)

func (r DataExportDestinationResponseDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationResponseDestinationTypeS3, DataExportDestinationResponseDestinationTypeGcs:
		return true
	}
	return false
}

type DataExportDestinationS3RequestParam struct {
	// Name of the S3 bucket for the Export Destination.
	BucketName param.Field[string] `json:"bucketName,required"`
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
	// The type of destination to create. Possible values are: S3
	DestinationType param.Field[DataExportDestinationS3RequestDestinationType] `json:"destinationType"`
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
	PartitionOrder param.Field[DataExportDestinationS3RequestPartitionOrder] `json:"partitionOrder"`
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

func (r DataExportDestinationS3RequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportDestinationS3RequestParam) implementsDataExportDestinationNewParamsBodyUnion() {}

func (r DataExportDestinationS3RequestParam) implementsDataExportDestinationUpdateParamsBodyUnion() {}

// The type of destination to create. Possible values are: S3
type DataExportDestinationS3RequestDestinationType string

const (
	DataExportDestinationS3RequestDestinationTypeS3 DataExportDestinationS3RequestDestinationType = "S3"
)

func (r DataExportDestinationS3RequestDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationS3RequestDestinationTypeS3:
		return true
	}
	return false
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
type DataExportDestinationS3RequestPartitionOrder string

const (
	DataExportDestinationS3RequestPartitionOrderTypeFirst DataExportDestinationS3RequestPartitionOrder = "TYPE_FIRST"
	DataExportDestinationS3RequestPartitionOrderTimeFirst DataExportDestinationS3RequestPartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationS3RequestPartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationS3RequestPartitionOrderTypeFirst, DataExportDestinationS3RequestPartitionOrderTimeFirst:
		return true
	}
	return false
}

// The response containing the details of an Google Cloud Storage export
// destination.
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
	// The code of the data Export Destination.
	Code string `json:"code"`
	// The id of the user who created the Export Destination.
	CreatedBy       string                                          `json:"createdBy"`
	DestinationType DataExportDestinationNewResponseDestinationType `json:"destinationType"`
	// The DateTime when the Export Destination was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Export Destination was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// The id of the user who last modified the Export Destination.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data Export Destination.
	Name string `json:"name"`
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
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                               `json:"serviceAccountEmail"`
	JSON                dataExportDestinationNewResponseJSON `json:"-"`
	union               DataExportDestinationNewResponseUnion
}

// dataExportDestinationNewResponseJSON contains the JSON metadata for the struct
// [DataExportDestinationNewResponse]
type dataExportDestinationNewResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	Code                apijson.Field
	CreatedBy           apijson.Field
	DestinationType     apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	IamRoleArn          apijson.Field
	LastModifiedBy      apijson.Field
	Name                apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r dataExportDestinationNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportDestinationNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportDestinationNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportDestinationNewResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DataExportDestinationNewResponseExportDestinationS3Response],
// [DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse].
func (r DataExportDestinationNewResponse) AsUnion() DataExportDestinationNewResponseUnion {
	return r.union
}

// The response containing the details of an Google Cloud Storage export
// destination.
//
// Union satisfied by [DataExportDestinationNewResponseExportDestinationS3Response]
// or
// [DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse].
type DataExportDestinationNewResponseUnion interface {
	implementsDataExportDestinationNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportDestinationNewResponseUnion)(nil)).Elem(),
		"destinationType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationNewResponseExportDestinationS3Response{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationNewResponseExportDestinationS3Response{}),
			DiscriminatorValue: "GCS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "GCS",
		},
	)
}

type DataExportDestinationNewResponseExportDestinationS3Response struct {
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
	PartitionOrder DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                                                          `json:"prefix"`
	JSON   dataExportDestinationNewResponseExportDestinationS3ResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationNewResponseExportDestinationS3ResponseJSON contains the
// JSON metadata for the struct
// [DataExportDestinationNewResponseExportDestinationS3Response]
type dataExportDestinationNewResponseExportDestinationS3ResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationNewResponseExportDestinationS3Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationNewResponseExportDestinationS3ResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationNewResponseExportDestinationS3Response) implementsDataExportDestinationNewResponse() {
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
type DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrder string

const (
	DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrderTypeFirst DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrderTimeFirst DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrderTypeFirst, DataExportDestinationNewResponseExportDestinationS3ResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

// The response containing the details of an Google Cloud Storage export
// destination.
type DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The bucket name.
	BucketName string `json:"bucketName"`
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
	PartitionOrder DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrder `json:"partitionOrder"`
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// The prefix.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                                                                          `json:"serviceAccountEmail"`
	JSON                dataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponseJSON
// contains the JSON metadata for the struct
// [DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse]
type dataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponse) implementsDataExportDestinationNewResponse() {
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
type DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrder string

const (
	DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst, DataExportDestinationNewResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationNewResponseDestinationType string

const (
	DataExportDestinationNewResponseDestinationTypeS3  DataExportDestinationNewResponseDestinationType = "S3"
	DataExportDestinationNewResponseDestinationTypeGcs DataExportDestinationNewResponseDestinationType = "GCS"
)

func (r DataExportDestinationNewResponseDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationNewResponseDestinationTypeS3, DataExportDestinationNewResponseDestinationTypeGcs:
		return true
	}
	return false
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

// The response containing the details of an Google Cloud Storage export
// destination.
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
	// The code of the data Export Destination.
	Code string `json:"code"`
	// The id of the user who created the Export Destination.
	CreatedBy       string                                          `json:"createdBy"`
	DestinationType DataExportDestinationGetResponseDestinationType `json:"destinationType"`
	// The DateTime when the Export Destination was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Export Destination was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// The id of the user who last modified the Export Destination.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data Export Destination.
	Name string `json:"name"`
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
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                               `json:"serviceAccountEmail"`
	JSON                dataExportDestinationGetResponseJSON `json:"-"`
	union               DataExportDestinationGetResponseUnion
}

// dataExportDestinationGetResponseJSON contains the JSON metadata for the struct
// [DataExportDestinationGetResponse]
type dataExportDestinationGetResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	Code                apijson.Field
	CreatedBy           apijson.Field
	DestinationType     apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	IamRoleArn          apijson.Field
	LastModifiedBy      apijson.Field
	Name                apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r dataExportDestinationGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportDestinationGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportDestinationGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportDestinationGetResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DataExportDestinationGetResponseExportDestinationS3Response],
// [DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse].
func (r DataExportDestinationGetResponse) AsUnion() DataExportDestinationGetResponseUnion {
	return r.union
}

// The response containing the details of an Google Cloud Storage export
// destination.
//
// Union satisfied by [DataExportDestinationGetResponseExportDestinationS3Response]
// or
// [DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse].
type DataExportDestinationGetResponseUnion interface {
	implementsDataExportDestinationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportDestinationGetResponseUnion)(nil)).Elem(),
		"destinationType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationGetResponseExportDestinationS3Response{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationGetResponseExportDestinationS3Response{}),
			DiscriminatorValue: "GCS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "GCS",
		},
	)
}

type DataExportDestinationGetResponseExportDestinationS3Response struct {
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
	PartitionOrder DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                                                          `json:"prefix"`
	JSON   dataExportDestinationGetResponseExportDestinationS3ResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationGetResponseExportDestinationS3ResponseJSON contains the
// JSON metadata for the struct
// [DataExportDestinationGetResponseExportDestinationS3Response]
type dataExportDestinationGetResponseExportDestinationS3ResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationGetResponseExportDestinationS3Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationGetResponseExportDestinationS3ResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationGetResponseExportDestinationS3Response) implementsDataExportDestinationGetResponse() {
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
type DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrder string

const (
	DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrderTypeFirst DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrderTimeFirst DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrderTypeFirst, DataExportDestinationGetResponseExportDestinationS3ResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

// The response containing the details of an Google Cloud Storage export
// destination.
type DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The bucket name.
	BucketName string `json:"bucketName"`
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
	PartitionOrder DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrder `json:"partitionOrder"`
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// The prefix.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                                                                          `json:"serviceAccountEmail"`
	JSON                dataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponseJSON
// contains the JSON metadata for the struct
// [DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse]
type dataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponse) implementsDataExportDestinationGetResponse() {
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
type DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrder string

const (
	DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst, DataExportDestinationGetResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationGetResponseDestinationType string

const (
	DataExportDestinationGetResponseDestinationTypeS3  DataExportDestinationGetResponseDestinationType = "S3"
	DataExportDestinationGetResponseDestinationTypeGcs DataExportDestinationGetResponseDestinationType = "GCS"
)

func (r DataExportDestinationGetResponseDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationGetResponseDestinationTypeS3, DataExportDestinationGetResponseDestinationTypeGcs:
		return true
	}
	return false
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

// The response containing the details of an Google Cloud Storage export
// destination.
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
	// The code of the data Export Destination.
	Code string `json:"code"`
	// The id of the user who created the Export Destination.
	CreatedBy       string                                             `json:"createdBy"`
	DestinationType DataExportDestinationUpdateResponseDestinationType `json:"destinationType"`
	// The DateTime when the Export Destination was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Export Destination was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// The id of the user who last modified the Export Destination.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data Export Destination.
	Name string `json:"name"`
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
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                                  `json:"serviceAccountEmail"`
	JSON                dataExportDestinationUpdateResponseJSON `json:"-"`
	union               DataExportDestinationUpdateResponseUnion
}

// dataExportDestinationUpdateResponseJSON contains the JSON metadata for the
// struct [DataExportDestinationUpdateResponse]
type dataExportDestinationUpdateResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	Code                apijson.Field
	CreatedBy           apijson.Field
	DestinationType     apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	IamRoleArn          apijson.Field
	LastModifiedBy      apijson.Field
	Name                apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r dataExportDestinationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportDestinationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportDestinationUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportDestinationUpdateResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DataExportDestinationUpdateResponseExportDestinationS3Response],
// [DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse].
func (r DataExportDestinationUpdateResponse) AsUnion() DataExportDestinationUpdateResponseUnion {
	return r.union
}

// The response containing the details of an Google Cloud Storage export
// destination.
//
// Union satisfied by
// [DataExportDestinationUpdateResponseExportDestinationS3Response] or
// [DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse].
type DataExportDestinationUpdateResponseUnion interface {
	implementsDataExportDestinationUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportDestinationUpdateResponseUnion)(nil)).Elem(),
		"destinationType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationUpdateResponseExportDestinationS3Response{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationUpdateResponseExportDestinationS3Response{}),
			DiscriminatorValue: "GCS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "GCS",
		},
	)
}

type DataExportDestinationUpdateResponseExportDestinationS3Response struct {
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
	PartitionOrder DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                                                             `json:"prefix"`
	JSON   dataExportDestinationUpdateResponseExportDestinationS3ResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationUpdateResponseExportDestinationS3ResponseJSON contains the
// JSON metadata for the struct
// [DataExportDestinationUpdateResponseExportDestinationS3Response]
type dataExportDestinationUpdateResponseExportDestinationS3ResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationUpdateResponseExportDestinationS3Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationUpdateResponseExportDestinationS3ResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationUpdateResponseExportDestinationS3Response) implementsDataExportDestinationUpdateResponse() {
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
type DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrder string

const (
	DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrderTypeFirst DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrderTimeFirst DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrderTypeFirst, DataExportDestinationUpdateResponseExportDestinationS3ResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

// The response containing the details of an Google Cloud Storage export
// destination.
type DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The bucket name.
	BucketName string `json:"bucketName"`
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
	PartitionOrder DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrder `json:"partitionOrder"`
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// The prefix.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                                                                             `json:"serviceAccountEmail"`
	JSON                dataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponseJSON
// contains the JSON metadata for the struct
// [DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse]
type dataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponse) implementsDataExportDestinationUpdateResponse() {
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
type DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrder string

const (
	DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst, DataExportDestinationUpdateResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationUpdateResponseDestinationType string

const (
	DataExportDestinationUpdateResponseDestinationTypeS3  DataExportDestinationUpdateResponseDestinationType = "S3"
	DataExportDestinationUpdateResponseDestinationTypeGcs DataExportDestinationUpdateResponseDestinationType = "GCS"
)

func (r DataExportDestinationUpdateResponseDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateResponseDestinationTypeS3, DataExportDestinationUpdateResponseDestinationTypeGcs:
		return true
	}
	return false
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

// The response containing the details of an Google Cloud Storage export
// destination.
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
	// The code of the data Export Destination.
	Code string `json:"code"`
	// The id of the user who created the Export Destination.
	CreatedBy       string                                             `json:"createdBy"`
	DestinationType DataExportDestinationDeleteResponseDestinationType `json:"destinationType"`
	// The DateTime when the Export Destination was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Export Destination was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The specified IAM role ARN with PutObject permission for the specified
	// `bucketName`, which allows the service to upload Data Exports to your S3 bucket.
	IamRoleArn string `json:"iamRoleArn"`
	// The id of the user who last modified the Export Destination.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data Export Destination.
	Name string `json:"name"`
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
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                                  `json:"serviceAccountEmail"`
	JSON                dataExportDestinationDeleteResponseJSON `json:"-"`
	union               DataExportDestinationDeleteResponseUnion
}

// dataExportDestinationDeleteResponseJSON contains the JSON metadata for the
// struct [DataExportDestinationDeleteResponse]
type dataExportDestinationDeleteResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	Code                apijson.Field
	CreatedBy           apijson.Field
	DestinationType     apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	IamRoleArn          apijson.Field
	LastModifiedBy      apijson.Field
	Name                apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r dataExportDestinationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DataExportDestinationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DataExportDestinationDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DataExportDestinationDeleteResponseUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DataExportDestinationDeleteResponseExportDestinationS3Response],
// [DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse].
func (r DataExportDestinationDeleteResponse) AsUnion() DataExportDestinationDeleteResponseUnion {
	return r.union
}

// The response containing the details of an Google Cloud Storage export
// destination.
//
// Union satisfied by
// [DataExportDestinationDeleteResponseExportDestinationS3Response] or
// [DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse].
type DataExportDestinationDeleteResponseUnion interface {
	implementsDataExportDestinationDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DataExportDestinationDeleteResponseUnion)(nil)).Elem(),
		"destinationType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationDeleteResponseExportDestinationS3Response{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationDeleteResponseExportDestinationS3Response{}),
			DiscriminatorValue: "GCS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "S3",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse{}),
			DiscriminatorValue: "GCS",
		},
	)
}

type DataExportDestinationDeleteResponseExportDestinationS3Response struct {
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
	PartitionOrder DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrder `json:"partitionOrder"`
	// Location in specified S3 bucket for the Export Destination. If no `prefix` is
	// specified, then the root of the bucket is used.
	Prefix string                                                             `json:"prefix"`
	JSON   dataExportDestinationDeleteResponseExportDestinationS3ResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationDeleteResponseExportDestinationS3ResponseJSON contains the
// JSON metadata for the struct
// [DataExportDestinationDeleteResponseExportDestinationS3Response]
type dataExportDestinationDeleteResponseExportDestinationS3ResponseJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	BucketName     apijson.Field
	IamRoleArn     apijson.Field
	PartitionOrder apijson.Field
	Prefix         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DataExportDestinationDeleteResponseExportDestinationS3Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationDeleteResponseExportDestinationS3ResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationDeleteResponseExportDestinationS3Response) implementsDataExportDestinationDeleteResponse() {
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
type DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrder string

const (
	DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrderTypeFirst DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrderTimeFirst DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrderTypeFirst, DataExportDestinationDeleteResponseExportDestinationS3ResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

// The response containing the details of an Google Cloud Storage export
// destination.
type DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The bucket name.
	BucketName string `json:"bucketName"`
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
	PartitionOrder DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrder `json:"partitionOrder"`
	// The export destination Web Identity Federation poolId.
	PoolID string `json:"poolId"`
	// The prefix.
	Prefix string `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber string `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID string `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail string                                                                             `json:"serviceAccountEmail"`
	JSON                dataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponseJSON `json:"-"`
	DataExportDestinationResponse
}

// dataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponseJSON
// contains the JSON metadata for the struct
// [DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse]
type dataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	BucketName          apijson.Field
	PartitionOrder      apijson.Field
	PoolID              apijson.Field
	Prefix              apijson.Field
	ProjectNumber       apijson.Field
	ProviderID          apijson.Field
	ServiceAccountEmail apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponseJSON) RawJSON() string {
	return r.raw
}

func (r DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponse) implementsDataExportDestinationDeleteResponse() {
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
type DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrder string

const (
	DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TYPE_FIRST"
	DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTypeFirst, DataExportDestinationDeleteResponseExportDestinationGoogleCloudStorageResponsePartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationDeleteResponseDestinationType string

const (
	DataExportDestinationDeleteResponseDestinationTypeS3  DataExportDestinationDeleteResponseDestinationType = "S3"
	DataExportDestinationDeleteResponseDestinationTypeGcs DataExportDestinationDeleteResponseDestinationType = "GCS"
)

func (r DataExportDestinationDeleteResponseDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationDeleteResponseDestinationTypeS3, DataExportDestinationDeleteResponseDestinationTypeGcs:
		return true
	}
	return false
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
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string]                     `path:"orgId,required"`
	Body  DataExportDestinationNewParamsBodyUnion `json:"body,required"`
}

func (r DataExportDestinationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DataExportDestinationNewParamsBody struct {
	// Name of the S3 bucket for the Export Destination.
	BucketName param.Field[string] `json:"bucketName,required"`
	// The type of destination to create. Possible values are: S3
	DestinationType param.Field[DataExportDestinationNewParamsBodyDestinationType] `json:"destinationType"`
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
	IamRoleArn param.Field[string] `json:"iamRoleArn"`
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
	PartitionOrder param.Field[DataExportDestinationNewParamsBodyPartitionOrder] `json:"partitionOrder"`
	// The export destination Web Identity Federation poolId.
	PoolID param.Field[string] `json:"poolId"`
	// Location in specified S3 bucket for the Export Destination. If you omit a
	// `prefix`, then the root of the bucket will be used.
	Prefix param.Field[string] `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber param.Field[string] `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID param.Field[string] `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail param.Field[string] `json:"serviceAccountEmail"`
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

func (r DataExportDestinationNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportDestinationNewParamsBody) implementsDataExportDestinationNewParamsBodyUnion() {}

// Satisfied by [DataExportDestinationS3RequestParam],
// [DataExportDestinationGoogleCloudStorageRequestParam],
// [DataExportDestinationNewParamsBody].
type DataExportDestinationNewParamsBodyUnion interface {
	implementsDataExportDestinationNewParamsBodyUnion()
}

// The type of destination to create. Possible values are: S3
type DataExportDestinationNewParamsBodyDestinationType string

const (
	DataExportDestinationNewParamsBodyDestinationTypeS3  DataExportDestinationNewParamsBodyDestinationType = "S3"
	DataExportDestinationNewParamsBodyDestinationTypeGcs DataExportDestinationNewParamsBodyDestinationType = "GCS"
)

func (r DataExportDestinationNewParamsBodyDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationNewParamsBodyDestinationTypeS3, DataExportDestinationNewParamsBodyDestinationTypeGcs:
		return true
	}
	return false
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
type DataExportDestinationNewParamsBodyPartitionOrder string

const (
	DataExportDestinationNewParamsBodyPartitionOrderTypeFirst DataExportDestinationNewParamsBodyPartitionOrder = "TYPE_FIRST"
	DataExportDestinationNewParamsBodyPartitionOrderTimeFirst DataExportDestinationNewParamsBodyPartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationNewParamsBodyPartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationNewParamsBodyPartitionOrderTypeFirst, DataExportDestinationNewParamsBodyPartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type DataExportDestinationUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string]                        `path:"orgId,required"`
	Body  DataExportDestinationUpdateParamsBodyUnion `json:"body,required"`
}

func (r DataExportDestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DataExportDestinationUpdateParamsBody struct {
	// Name of the S3 bucket for the Export Destination.
	BucketName param.Field[string] `json:"bucketName,required"`
	// The type of destination to create. Possible values are: S3
	DestinationType param.Field[DataExportDestinationUpdateParamsBodyDestinationType] `json:"destinationType"`
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
	IamRoleArn param.Field[string] `json:"iamRoleArn"`
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
	PartitionOrder param.Field[DataExportDestinationUpdateParamsBodyPartitionOrder] `json:"partitionOrder"`
	// The export destination Web Identity Federation poolId.
	PoolID param.Field[string] `json:"poolId"`
	// Location in specified S3 bucket for the Export Destination. If you omit a
	// `prefix`, then the root of the bucket will be used.
	Prefix param.Field[string] `json:"prefix"`
	// The export destination GCP projectNumber.
	ProjectNumber param.Field[string] `json:"projectNumber"`
	// The export destination Web Identity Federation identity providerId.
	ProviderID param.Field[string] `json:"providerId"`
	// The export destination service account email.
	ServiceAccountEmail param.Field[string] `json:"serviceAccountEmail"`
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

func (r DataExportDestinationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DataExportDestinationUpdateParamsBody) implementsDataExportDestinationUpdateParamsBodyUnion() {
}

// Satisfied by [DataExportDestinationS3RequestParam],
// [DataExportDestinationGoogleCloudStorageRequestParam],
// [DataExportDestinationUpdateParamsBody].
type DataExportDestinationUpdateParamsBodyUnion interface {
	implementsDataExportDestinationUpdateParamsBodyUnion()
}

// The type of destination to create. Possible values are: S3
type DataExportDestinationUpdateParamsBodyDestinationType string

const (
	DataExportDestinationUpdateParamsBodyDestinationTypeS3  DataExportDestinationUpdateParamsBodyDestinationType = "S3"
	DataExportDestinationUpdateParamsBodyDestinationTypeGcs DataExportDestinationUpdateParamsBodyDestinationType = "GCS"
)

func (r DataExportDestinationUpdateParamsBodyDestinationType) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateParamsBodyDestinationTypeS3, DataExportDestinationUpdateParamsBodyDestinationTypeGcs:
		return true
	}
	return false
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
type DataExportDestinationUpdateParamsBodyPartitionOrder string

const (
	DataExportDestinationUpdateParamsBodyPartitionOrderTypeFirst DataExportDestinationUpdateParamsBodyPartitionOrder = "TYPE_FIRST"
	DataExportDestinationUpdateParamsBodyPartitionOrderTimeFirst DataExportDestinationUpdateParamsBodyPartitionOrder = "TIME_FIRST"
)

func (r DataExportDestinationUpdateParamsBodyPartitionOrder) IsKnown() bool {
	switch r {
	case DataExportDestinationUpdateParamsBodyPartitionOrderTypeFirst, DataExportDestinationUpdateParamsBodyPartitionOrderTimeFirst:
		return true
	}
	return false
}

type DataExportDestinationListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
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
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
