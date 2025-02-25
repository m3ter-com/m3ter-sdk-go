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

// PermissionPolicyService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPermissionPolicyService] method instead.
type PermissionPolicyService struct {
	Options []option.RequestOption
}

// NewPermissionPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPermissionPolicyService(opts ...option.RequestOption) (r *PermissionPolicyService) {
	r = &PermissionPolicyService{}
	r.Options = opts
	return
}

// Create a new Permission Policy
//
// **NOTE:** When you set up a policy statement for this call using the
// `permissionPolicy` request parameter to specify the `effect`, `action`, and
// `resource`, you must use all lower case and the format as shown in this example
// for a Permission Policy statement that grants full CRUD access to all meters:
//
// ```
// "permissionPolicy" : [
//
//	{
//		"effect" : "allow",
//		"action" : [
//		"config:create",
//		"config:delete",
//		"config:retrieve",
//		"config:update"
//		]
//		"resource" : [
//		"config:meter/*"
//		]
//	}
//
// ]
// ```
//
// For more details and further examples, see
// [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/organization-and-access-management/creating-and-managing-permissions#permission-policy-statements---available-actions-and-resources)
// in our main Documentation.
func (r *PermissionPolicyService) New(ctx context.Context, params PermissionPolicyNewParams, opts ...option.RequestOption) (res *PermissionPolicy, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the permission policy for the UUID
func (r *PermissionPolicyService) Get(ctx context.Context, id string, query PermissionPolicyGetParams, opts ...option.RequestOption) (res *PermissionPolicy, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Permission Policy for the UUID
//
// **NOTE:** When you set up a policy statement for this call to specify the
// `effect`, `action`, and `resource`, you must use all lower case and the format
// as shown in this example - a Permission Policy statement that grants full CRUD
// access to all meters:
//
// ```
// "permissionPolicy" : [
//
//	{
//		"effect" : "allow",
//		"action" : [
//		"config:create",
//		"config:delete",
//		"config:retrieve",
//		"config:update"
//		]
//		"resource" : [
//		"config:meter/*"
//		]
//	}
//
// ]
// ```
//
// For more details and further examples, see
// [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/organization-and-access-management/creating-and-managing-permissions#permission-policy-statements---available-actions-and-resources)
// in our main Documentation.
func (r *PermissionPolicyService) Update(ctx context.Context, id string, params PermissionPolicyUpdateParams, opts ...option.RequestOption) (res *PermissionPolicy, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of PermissionPolicy entities
func (r *PermissionPolicyService) List(ctx context.Context, params PermissionPolicyListParams, opts ...option.RequestOption) (res *pagination.Cursor[PermissionPolicy], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies", params.OrgID)
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

// Retrieve a list of PermissionPolicy entities
func (r *PermissionPolicyService) ListAutoPaging(ctx context.Context, params PermissionPolicyListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PermissionPolicy] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the PermissionPolicy for the UUID
func (r *PermissionPolicyService) Delete(ctx context.Context, id string, body PermissionPolicyDeleteParams, opts ...option.RequestOption) (res *PermissionPolicy, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a permission policy to a service user.
func (r *PermissionPolicyService) AddToServiceUser(ctx context.Context, permissionPolicyID string, params PermissionPolicyAddToServiceUserParams, opts ...option.RequestOption) (res *PermissionPolicyAddToServiceUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/addtoserviceuser", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add a permission policy to support users for an organization.
func (r *PermissionPolicyService) AddToSupportUser(ctx context.Context, permissionPolicyID string, params PermissionPolicyAddToSupportUserParams, opts ...option.RequestOption) (res *PermissionPolicyAddToSupportUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/addtosupportusers", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add a permission policy to a user.
func (r *PermissionPolicyService) AddToUser(ctx context.Context, permissionPolicyID string, params PermissionPolicyAddToUserParams, opts ...option.RequestOption) (res *PermissionPolicyAddToUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/addtouser", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add a permission Policy to a user group
func (r *PermissionPolicyService) AddToUserGroup(ctx context.Context, permissionPolicyID string, params PermissionPolicyAddToUserGroupParams, opts ...option.RequestOption) (res *PermissionPolicyAddToUserGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/addtousergroup", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove a permission policy from a service user.
func (r *PermissionPolicyService) RemoveFromServiceUser(ctx context.Context, permissionPolicyID string, params PermissionPolicyRemoveFromServiceUserParams, opts ...option.RequestOption) (res *PermissionPolicyRemoveFromServiceUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/removefromserviceuser", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove a permission policy from support users for an organization.
func (r *PermissionPolicyService) RemoveFromSupportUser(ctx context.Context, permissionPolicyID string, body PermissionPolicyRemoveFromSupportUserParams, opts ...option.RequestOption) (res *PermissionPolicyRemoveFromSupportUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/removefromsupportusers", body.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Remove a permission policy from a user.
func (r *PermissionPolicyService) RemoveFromUser(ctx context.Context, permissionPolicyID string, params PermissionPolicyRemoveFromUserParams, opts ...option.RequestOption) (res *PermissionPolicyRemoveFromUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/removefromuser", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove a permission policy from a user group.
func (r *PermissionPolicyService) RemoveFromUserGroup(ctx context.Context, permissionPolicyID string, params PermissionPolicyRemoveFromUserGroupParams, opts ...option.RequestOption) (res *PermissionPolicyRemoveFromUserGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if permissionPolicyID == "" {
		err = errors.New("missing required permissionPolicyId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/permissionpolicies/%s/removefromusergroup", params.OrgID, permissionPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type PermissionPolicy struct {
	// The unique identifier (UUID) for this Permission Policy.
	ID string `json:"id"`
	// The unique identifier (UUID) of the user who created this Permission Policy.
	CreatedBy string `json:"createdBy"`
	// The date and time _(in ISO-8601 format)_ when the Permission Policy was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the Permission Policy was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this Permission
	// Policy.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Indicates whether this is a system generated Managed Permission Policy.
	ManagedPolicy bool `json:"managedPolicy"`
	// The name of the Permission Policy.
	Name string `json:"name"`
	// Array containing the Permission Policies information.
	PermissionPolicy []PermissionPolicyPermissionPolicy `json:"permissionPolicy"`
	// The version number. Default value when newly created is one.
	Version int64                `json:"version"`
	JSON    permissionPolicyJSON `json:"-"`
}

// permissionPolicyJSON contains the JSON metadata for the struct
// [PermissionPolicy]
type permissionPolicyJSON struct {
	ID               apijson.Field
	CreatedBy        apijson.Field
	DtCreated        apijson.Field
	DtLastModified   apijson.Field
	LastModifiedBy   apijson.Field
	ManagedPolicy    apijson.Field
	Name             apijson.Field
	PermissionPolicy apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PermissionPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyPermissionPolicy struct {
	// The actions available to users who are assigned the Permission Policy - what
	// they can do or cannot do with respect to the specified resource.
	//
	// **NOTE:** Use lower case and a colon-separated format, for example, if you want
	// to confer full CRUD, use:
	//
	// ```
	// "config:create",
	// "config:delete",
	// "config:retrieve",
	// "config:update"
	// ```
	Action []PermissionPolicyPermissionPolicyAction `json:"action,required"`
	// Specifies whether or not the user is allowed to perform the action on the
	// resource.
	//
	// **NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll
	// receive an error.
	Effect PermissionPolicyPermissionPolicyEffect `json:"effect,required"`
	// See
	// [Statements - Available Resources](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions#statements---available-resources)
	// for a listing of available resources for Permission Policy statements.
	Resource []string                             `json:"resource,required"`
	JSON     permissionPolicyPermissionPolicyJSON `json:"-"`
}

// permissionPolicyPermissionPolicyJSON contains the JSON metadata for the struct
// [PermissionPolicyPermissionPolicy]
type permissionPolicyPermissionPolicyJSON struct {
	Action      apijson.Field
	Effect      apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionPolicyPermissionPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyPermissionPolicyJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyPermissionPolicyAction string

const (
	PermissionPolicyPermissionPolicyActionAll                    PermissionPolicyPermissionPolicyAction = "ALL"
	PermissionPolicyPermissionPolicyActionConfigCreate           PermissionPolicyPermissionPolicyAction = "CONFIG_CREATE"
	PermissionPolicyPermissionPolicyActionConfigRetrieve         PermissionPolicyPermissionPolicyAction = "CONFIG_RETRIEVE"
	PermissionPolicyPermissionPolicyActionConfigUpdate           PermissionPolicyPermissionPolicyAction = "CONFIG_UPDATE"
	PermissionPolicyPermissionPolicyActionConfigDelete           PermissionPolicyPermissionPolicyAction = "CONFIG_DELETE"
	PermissionPolicyPermissionPolicyActionConfigExport           PermissionPolicyPermissionPolicyAction = "CONFIG_EXPORT"
	PermissionPolicyPermissionPolicyActionAnalyticsQuery         PermissionPolicyPermissionPolicyAction = "ANALYTICS_QUERY"
	PermissionPolicyPermissionPolicyActionMeasurementsUpload     PermissionPolicyPermissionPolicyAction = "MEASUREMENTS_UPLOAD"
	PermissionPolicyPermissionPolicyActionMeasurementsFileupload PermissionPolicyPermissionPolicyAction = "MEASUREMENTS_FILEUPLOAD"
	PermissionPolicyPermissionPolicyActionMeasurementsRetrieve   PermissionPolicyPermissionPolicyAction = "MEASUREMENTS_RETRIEVE"
	PermissionPolicyPermissionPolicyActionMeasurementsExport     PermissionPolicyPermissionPolicyAction = "MEASUREMENTS_EXPORT"
	PermissionPolicyPermissionPolicyActionForecastRetrieve       PermissionPolicyPermissionPolicyAction = "FORECAST_RETRIEVE"
	PermissionPolicyPermissionPolicyActionHealthscoresRetrieve   PermissionPolicyPermissionPolicyAction = "HEALTHSCORES_RETRIEVE"
	PermissionPolicyPermissionPolicyActionAnomaliesRetrieve      PermissionPolicyPermissionPolicyAction = "ANOMALIES_RETRIEVE"
	PermissionPolicyPermissionPolicyActionExportsDownload        PermissionPolicyPermissionPolicyAction = "EXPORTS_DOWNLOAD"
)

func (r PermissionPolicyPermissionPolicyAction) IsKnown() bool {
	switch r {
	case PermissionPolicyPermissionPolicyActionAll, PermissionPolicyPermissionPolicyActionConfigCreate, PermissionPolicyPermissionPolicyActionConfigRetrieve, PermissionPolicyPermissionPolicyActionConfigUpdate, PermissionPolicyPermissionPolicyActionConfigDelete, PermissionPolicyPermissionPolicyActionConfigExport, PermissionPolicyPermissionPolicyActionAnalyticsQuery, PermissionPolicyPermissionPolicyActionMeasurementsUpload, PermissionPolicyPermissionPolicyActionMeasurementsFileupload, PermissionPolicyPermissionPolicyActionMeasurementsRetrieve, PermissionPolicyPermissionPolicyActionMeasurementsExport, PermissionPolicyPermissionPolicyActionForecastRetrieve, PermissionPolicyPermissionPolicyActionHealthscoresRetrieve, PermissionPolicyPermissionPolicyActionAnomaliesRetrieve, PermissionPolicyPermissionPolicyActionExportsDownload:
		return true
	}
	return false
}

// Specifies whether or not the user is allowed to perform the action on the
// resource.
//
// **NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll
// receive an error.
type PermissionPolicyPermissionPolicyEffect string

const (
	PermissionPolicyPermissionPolicyEffectAllow PermissionPolicyPermissionPolicyEffect = "ALLOW"
	PermissionPolicyPermissionPolicyEffectDeny  PermissionPolicyPermissionPolicyEffect = "DENY"
)

func (r PermissionPolicyPermissionPolicyEffect) IsKnown() bool {
	switch r {
	case PermissionPolicyPermissionPolicyEffectAllow, PermissionPolicyPermissionPolicyEffectDeny:
		return true
	}
	return false
}

type PermissionPolicyAddToServiceUserResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                                `json:"lastModifiedBy"`
	PermissionPolicyID string                                                `json:"permissionPolicyId"`
	PrincipalID        string                                                `json:"principalId"`
	PrincipalType      PermissionPolicyAddToServiceUserResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                        `json:"version"`
	JSON    permissionPolicyAddToServiceUserResponseJSON `json:"-"`
}

// permissionPolicyAddToServiceUserResponseJSON contains the JSON metadata for the
// struct [PermissionPolicyAddToServiceUserResponse]
type permissionPolicyAddToServiceUserResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyAddToServiceUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyAddToServiceUserResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyAddToServiceUserResponsePrincipalType string

const (
	PermissionPolicyAddToServiceUserResponsePrincipalTypeUser         PermissionPolicyAddToServiceUserResponsePrincipalType = "USER"
	PermissionPolicyAddToServiceUserResponsePrincipalTypeUsergroup    PermissionPolicyAddToServiceUserResponsePrincipalType = "USERGROUP"
	PermissionPolicyAddToServiceUserResponsePrincipalTypeServiceuser  PermissionPolicyAddToServiceUserResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyAddToServiceUserResponsePrincipalTypeSupportusers PermissionPolicyAddToServiceUserResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyAddToServiceUserResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyAddToServiceUserResponsePrincipalTypeUser, PermissionPolicyAddToServiceUserResponsePrincipalTypeUsergroup, PermissionPolicyAddToServiceUserResponsePrincipalTypeServiceuser, PermissionPolicyAddToServiceUserResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyAddToSupportUserResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                                `json:"lastModifiedBy"`
	PermissionPolicyID string                                                `json:"permissionPolicyId"`
	PrincipalID        string                                                `json:"principalId"`
	PrincipalType      PermissionPolicyAddToSupportUserResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                        `json:"version"`
	JSON    permissionPolicyAddToSupportUserResponseJSON `json:"-"`
}

// permissionPolicyAddToSupportUserResponseJSON contains the JSON metadata for the
// struct [PermissionPolicyAddToSupportUserResponse]
type permissionPolicyAddToSupportUserResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyAddToSupportUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyAddToSupportUserResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyAddToSupportUserResponsePrincipalType string

const (
	PermissionPolicyAddToSupportUserResponsePrincipalTypeUser         PermissionPolicyAddToSupportUserResponsePrincipalType = "USER"
	PermissionPolicyAddToSupportUserResponsePrincipalTypeUsergroup    PermissionPolicyAddToSupportUserResponsePrincipalType = "USERGROUP"
	PermissionPolicyAddToSupportUserResponsePrincipalTypeServiceuser  PermissionPolicyAddToSupportUserResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyAddToSupportUserResponsePrincipalTypeSupportusers PermissionPolicyAddToSupportUserResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyAddToSupportUserResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyAddToSupportUserResponsePrincipalTypeUser, PermissionPolicyAddToSupportUserResponsePrincipalTypeUsergroup, PermissionPolicyAddToSupportUserResponsePrincipalTypeServiceuser, PermissionPolicyAddToSupportUserResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyAddToUserResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                         `json:"lastModifiedBy"`
	PermissionPolicyID string                                         `json:"permissionPolicyId"`
	PrincipalID        string                                         `json:"principalId"`
	PrincipalType      PermissionPolicyAddToUserResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                 `json:"version"`
	JSON    permissionPolicyAddToUserResponseJSON `json:"-"`
}

// permissionPolicyAddToUserResponseJSON contains the JSON metadata for the struct
// [PermissionPolicyAddToUserResponse]
type permissionPolicyAddToUserResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyAddToUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyAddToUserResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyAddToUserResponsePrincipalType string

const (
	PermissionPolicyAddToUserResponsePrincipalTypeUser         PermissionPolicyAddToUserResponsePrincipalType = "USER"
	PermissionPolicyAddToUserResponsePrincipalTypeUsergroup    PermissionPolicyAddToUserResponsePrincipalType = "USERGROUP"
	PermissionPolicyAddToUserResponsePrincipalTypeServiceuser  PermissionPolicyAddToUserResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyAddToUserResponsePrincipalTypeSupportusers PermissionPolicyAddToUserResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyAddToUserResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyAddToUserResponsePrincipalTypeUser, PermissionPolicyAddToUserResponsePrincipalTypeUsergroup, PermissionPolicyAddToUserResponsePrincipalTypeServiceuser, PermissionPolicyAddToUserResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyAddToUserGroupResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                              `json:"lastModifiedBy"`
	PermissionPolicyID string                                              `json:"permissionPolicyId"`
	PrincipalID        string                                              `json:"principalId"`
	PrincipalType      PermissionPolicyAddToUserGroupResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                      `json:"version"`
	JSON    permissionPolicyAddToUserGroupResponseJSON `json:"-"`
}

// permissionPolicyAddToUserGroupResponseJSON contains the JSON metadata for the
// struct [PermissionPolicyAddToUserGroupResponse]
type permissionPolicyAddToUserGroupResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyAddToUserGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyAddToUserGroupResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyAddToUserGroupResponsePrincipalType string

const (
	PermissionPolicyAddToUserGroupResponsePrincipalTypeUser         PermissionPolicyAddToUserGroupResponsePrincipalType = "USER"
	PermissionPolicyAddToUserGroupResponsePrincipalTypeUsergroup    PermissionPolicyAddToUserGroupResponsePrincipalType = "USERGROUP"
	PermissionPolicyAddToUserGroupResponsePrincipalTypeServiceuser  PermissionPolicyAddToUserGroupResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyAddToUserGroupResponsePrincipalTypeSupportusers PermissionPolicyAddToUserGroupResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyAddToUserGroupResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyAddToUserGroupResponsePrincipalTypeUser, PermissionPolicyAddToUserGroupResponsePrincipalTypeUsergroup, PermissionPolicyAddToUserGroupResponsePrincipalTypeServiceuser, PermissionPolicyAddToUserGroupResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyRemoveFromServiceUserResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                                     `json:"lastModifiedBy"`
	PermissionPolicyID string                                                     `json:"permissionPolicyId"`
	PrincipalID        string                                                     `json:"principalId"`
	PrincipalType      PermissionPolicyRemoveFromServiceUserResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                             `json:"version"`
	JSON    permissionPolicyRemoveFromServiceUserResponseJSON `json:"-"`
}

// permissionPolicyRemoveFromServiceUserResponseJSON contains the JSON metadata for
// the struct [PermissionPolicyRemoveFromServiceUserResponse]
type permissionPolicyRemoveFromServiceUserResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyRemoveFromServiceUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyRemoveFromServiceUserResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyRemoveFromServiceUserResponsePrincipalType string

const (
	PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeUser         PermissionPolicyRemoveFromServiceUserResponsePrincipalType = "USER"
	PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeUsergroup    PermissionPolicyRemoveFromServiceUserResponsePrincipalType = "USERGROUP"
	PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeServiceuser  PermissionPolicyRemoveFromServiceUserResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeSupportusers PermissionPolicyRemoveFromServiceUserResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyRemoveFromServiceUserResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeUser, PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeUsergroup, PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeServiceuser, PermissionPolicyRemoveFromServiceUserResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyRemoveFromSupportUserResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                                     `json:"lastModifiedBy"`
	PermissionPolicyID string                                                     `json:"permissionPolicyId"`
	PrincipalID        string                                                     `json:"principalId"`
	PrincipalType      PermissionPolicyRemoveFromSupportUserResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                             `json:"version"`
	JSON    permissionPolicyRemoveFromSupportUserResponseJSON `json:"-"`
}

// permissionPolicyRemoveFromSupportUserResponseJSON contains the JSON metadata for
// the struct [PermissionPolicyRemoveFromSupportUserResponse]
type permissionPolicyRemoveFromSupportUserResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyRemoveFromSupportUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyRemoveFromSupportUserResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyRemoveFromSupportUserResponsePrincipalType string

const (
	PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeUser         PermissionPolicyRemoveFromSupportUserResponsePrincipalType = "USER"
	PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeUsergroup    PermissionPolicyRemoveFromSupportUserResponsePrincipalType = "USERGROUP"
	PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeServiceuser  PermissionPolicyRemoveFromSupportUserResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeSupportusers PermissionPolicyRemoveFromSupportUserResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyRemoveFromSupportUserResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeUser, PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeUsergroup, PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeServiceuser, PermissionPolicyRemoveFromSupportUserResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyRemoveFromUserResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                              `json:"lastModifiedBy"`
	PermissionPolicyID string                                              `json:"permissionPolicyId"`
	PrincipalID        string                                              `json:"principalId"`
	PrincipalType      PermissionPolicyRemoveFromUserResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                      `json:"version"`
	JSON    permissionPolicyRemoveFromUserResponseJSON `json:"-"`
}

// permissionPolicyRemoveFromUserResponseJSON contains the JSON metadata for the
// struct [PermissionPolicyRemoveFromUserResponse]
type permissionPolicyRemoveFromUserResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyRemoveFromUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyRemoveFromUserResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyRemoveFromUserResponsePrincipalType string

const (
	PermissionPolicyRemoveFromUserResponsePrincipalTypeUser         PermissionPolicyRemoveFromUserResponsePrincipalType = "USER"
	PermissionPolicyRemoveFromUserResponsePrincipalTypeUsergroup    PermissionPolicyRemoveFromUserResponsePrincipalType = "USERGROUP"
	PermissionPolicyRemoveFromUserResponsePrincipalTypeServiceuser  PermissionPolicyRemoveFromUserResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyRemoveFromUserResponsePrincipalTypeSupportusers PermissionPolicyRemoveFromUserResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyRemoveFromUserResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyRemoveFromUserResponsePrincipalTypeUser, PermissionPolicyRemoveFromUserResponsePrincipalTypeUsergroup, PermissionPolicyRemoveFromUserResponsePrincipalTypeServiceuser, PermissionPolicyRemoveFromUserResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyRemoveFromUserGroupResponse struct {
	ID string `json:"id"`
	// The id of the user who created this principal permission.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the principal permission was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this principal permission.
	LastModifiedBy     string                                                   `json:"lastModifiedBy"`
	PermissionPolicyID string                                                   `json:"permissionPolicyId"`
	PrincipalID        string                                                   `json:"principalId"`
	PrincipalType      PermissionPolicyRemoveFromUserGroupResponsePrincipalType `json:"principalType"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                           `json:"version"`
	JSON    permissionPolicyRemoveFromUserGroupResponseJSON `json:"-"`
}

// permissionPolicyRemoveFromUserGroupResponseJSON contains the JSON metadata for
// the struct [PermissionPolicyRemoveFromUserGroupResponse]
type permissionPolicyRemoveFromUserGroupResponseJSON struct {
	ID                 apijson.Field
	CreatedBy          apijson.Field
	DtCreated          apijson.Field
	DtLastModified     apijson.Field
	LastModifiedBy     apijson.Field
	PermissionPolicyID apijson.Field
	PrincipalID        apijson.Field
	PrincipalType      apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PermissionPolicyRemoveFromUserGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionPolicyRemoveFromUserGroupResponseJSON) RawJSON() string {
	return r.raw
}

type PermissionPolicyRemoveFromUserGroupResponsePrincipalType string

const (
	PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeUser         PermissionPolicyRemoveFromUserGroupResponsePrincipalType = "USER"
	PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeUsergroup    PermissionPolicyRemoveFromUserGroupResponsePrincipalType = "USERGROUP"
	PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeServiceuser  PermissionPolicyRemoveFromUserGroupResponsePrincipalType = "SERVICEUSER"
	PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeSupportusers PermissionPolicyRemoveFromUserGroupResponsePrincipalType = "SUPPORTUSERS"
)

func (r PermissionPolicyRemoveFromUserGroupResponsePrincipalType) IsKnown() bool {
	switch r {
	case PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeUser, PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeUsergroup, PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeServiceuser, PermissionPolicyRemoveFromUserGroupResponsePrincipalTypeSupportusers:
		return true
	}
	return false
}

type PermissionPolicyNewParams struct {
	OrgID            param.Field[string]                                      `path:"orgId,required"`
	Name             param.Field[string]                                      `json:"name,required"`
	PermissionPolicy param.Field[[]PermissionPolicyNewParamsPermissionPolicy] `json:"permissionPolicy,required"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - do not use
	//     for Create. On initial Create, version is set at 1 and listed in the response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r PermissionPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyNewParamsPermissionPolicy struct {
	// The actions available to users who are assigned the Permission Policy - what
	// they can do or cannot do with respect to the specified resource.
	//
	// **NOTE:** Use lower case and a colon-separated format, for example, if you want
	// to confer full CRUD, use:
	//
	// ```
	// "config:create",
	// "config:delete",
	// "config:retrieve",
	// "config:update"
	// ```
	Action param.Field[[]PermissionPolicyNewParamsPermissionPolicyAction] `json:"action,required"`
	// Specifies whether or not the user is allowed to perform the action on the
	// resource.
	//
	// **NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll
	// receive an error.
	Effect param.Field[PermissionPolicyNewParamsPermissionPolicyEffect] `json:"effect,required"`
	// See
	// [Statements - Available Resources](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions#statements---available-resources)
	// for a listing of available resources for Permission Policy statements.
	Resource param.Field[[]string] `json:"resource,required"`
}

func (r PermissionPolicyNewParamsPermissionPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyNewParamsPermissionPolicyAction string

const (
	PermissionPolicyNewParamsPermissionPolicyActionAll                    PermissionPolicyNewParamsPermissionPolicyAction = "ALL"
	PermissionPolicyNewParamsPermissionPolicyActionConfigCreate           PermissionPolicyNewParamsPermissionPolicyAction = "CONFIG_CREATE"
	PermissionPolicyNewParamsPermissionPolicyActionConfigRetrieve         PermissionPolicyNewParamsPermissionPolicyAction = "CONFIG_RETRIEVE"
	PermissionPolicyNewParamsPermissionPolicyActionConfigUpdate           PermissionPolicyNewParamsPermissionPolicyAction = "CONFIG_UPDATE"
	PermissionPolicyNewParamsPermissionPolicyActionConfigDelete           PermissionPolicyNewParamsPermissionPolicyAction = "CONFIG_DELETE"
	PermissionPolicyNewParamsPermissionPolicyActionConfigExport           PermissionPolicyNewParamsPermissionPolicyAction = "CONFIG_EXPORT"
	PermissionPolicyNewParamsPermissionPolicyActionAnalyticsQuery         PermissionPolicyNewParamsPermissionPolicyAction = "ANALYTICS_QUERY"
	PermissionPolicyNewParamsPermissionPolicyActionMeasurementsUpload     PermissionPolicyNewParamsPermissionPolicyAction = "MEASUREMENTS_UPLOAD"
	PermissionPolicyNewParamsPermissionPolicyActionMeasurementsFileupload PermissionPolicyNewParamsPermissionPolicyAction = "MEASUREMENTS_FILEUPLOAD"
	PermissionPolicyNewParamsPermissionPolicyActionMeasurementsRetrieve   PermissionPolicyNewParamsPermissionPolicyAction = "MEASUREMENTS_RETRIEVE"
	PermissionPolicyNewParamsPermissionPolicyActionMeasurementsExport     PermissionPolicyNewParamsPermissionPolicyAction = "MEASUREMENTS_EXPORT"
	PermissionPolicyNewParamsPermissionPolicyActionForecastRetrieve       PermissionPolicyNewParamsPermissionPolicyAction = "FORECAST_RETRIEVE"
	PermissionPolicyNewParamsPermissionPolicyActionHealthscoresRetrieve   PermissionPolicyNewParamsPermissionPolicyAction = "HEALTHSCORES_RETRIEVE"
	PermissionPolicyNewParamsPermissionPolicyActionAnomaliesRetrieve      PermissionPolicyNewParamsPermissionPolicyAction = "ANOMALIES_RETRIEVE"
	PermissionPolicyNewParamsPermissionPolicyActionExportsDownload        PermissionPolicyNewParamsPermissionPolicyAction = "EXPORTS_DOWNLOAD"
)

func (r PermissionPolicyNewParamsPermissionPolicyAction) IsKnown() bool {
	switch r {
	case PermissionPolicyNewParamsPermissionPolicyActionAll, PermissionPolicyNewParamsPermissionPolicyActionConfigCreate, PermissionPolicyNewParamsPermissionPolicyActionConfigRetrieve, PermissionPolicyNewParamsPermissionPolicyActionConfigUpdate, PermissionPolicyNewParamsPermissionPolicyActionConfigDelete, PermissionPolicyNewParamsPermissionPolicyActionConfigExport, PermissionPolicyNewParamsPermissionPolicyActionAnalyticsQuery, PermissionPolicyNewParamsPermissionPolicyActionMeasurementsUpload, PermissionPolicyNewParamsPermissionPolicyActionMeasurementsFileupload, PermissionPolicyNewParamsPermissionPolicyActionMeasurementsRetrieve, PermissionPolicyNewParamsPermissionPolicyActionMeasurementsExport, PermissionPolicyNewParamsPermissionPolicyActionForecastRetrieve, PermissionPolicyNewParamsPermissionPolicyActionHealthscoresRetrieve, PermissionPolicyNewParamsPermissionPolicyActionAnomaliesRetrieve, PermissionPolicyNewParamsPermissionPolicyActionExportsDownload:
		return true
	}
	return false
}

// Specifies whether or not the user is allowed to perform the action on the
// resource.
//
// **NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll
// receive an error.
type PermissionPolicyNewParamsPermissionPolicyEffect string

const (
	PermissionPolicyNewParamsPermissionPolicyEffectAllow PermissionPolicyNewParamsPermissionPolicyEffect = "ALLOW"
	PermissionPolicyNewParamsPermissionPolicyEffectDeny  PermissionPolicyNewParamsPermissionPolicyEffect = "DENY"
)

func (r PermissionPolicyNewParamsPermissionPolicyEffect) IsKnown() bool {
	switch r {
	case PermissionPolicyNewParamsPermissionPolicyEffectAllow, PermissionPolicyNewParamsPermissionPolicyEffectDeny:
		return true
	}
	return false
}

type PermissionPolicyGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type PermissionPolicyUpdateParams struct {
	OrgID            param.Field[string]                                         `path:"orgId,required"`
	Name             param.Field[string]                                         `json:"name,required"`
	PermissionPolicy param.Field[[]PermissionPolicyUpdateParamsPermissionPolicy] `json:"permissionPolicy,required"`
	// The version number of the entity:
	//
	//   - **Create entity:** Not valid for initial insertion of new entity - do not use
	//     for Create. On initial Create, version is set at 1 and listed in the response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r PermissionPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyUpdateParamsPermissionPolicy struct {
	// The actions available to users who are assigned the Permission Policy - what
	// they can do or cannot do with respect to the specified resource.
	//
	// **NOTE:** Use lower case and a colon-separated format, for example, if you want
	// to confer full CRUD, use:
	//
	// ```
	// "config:create",
	// "config:delete",
	// "config:retrieve",
	// "config:update"
	// ```
	Action param.Field[[]PermissionPolicyUpdateParamsPermissionPolicyAction] `json:"action,required"`
	// Specifies whether or not the user is allowed to perform the action on the
	// resource.
	//
	// **NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll
	// receive an error.
	Effect param.Field[PermissionPolicyUpdateParamsPermissionPolicyEffect] `json:"effect,required"`
	// See
	// [Statements - Available Resources](https://www.m3ter.com/docs/guides/managing-organization-and-users/creating-and-managing-permissions#statements---available-resources)
	// for a listing of available resources for Permission Policy statements.
	Resource param.Field[[]string] `json:"resource,required"`
}

func (r PermissionPolicyUpdateParamsPermissionPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyUpdateParamsPermissionPolicyAction string

const (
	PermissionPolicyUpdateParamsPermissionPolicyActionAll                    PermissionPolicyUpdateParamsPermissionPolicyAction = "ALL"
	PermissionPolicyUpdateParamsPermissionPolicyActionConfigCreate           PermissionPolicyUpdateParamsPermissionPolicyAction = "CONFIG_CREATE"
	PermissionPolicyUpdateParamsPermissionPolicyActionConfigRetrieve         PermissionPolicyUpdateParamsPermissionPolicyAction = "CONFIG_RETRIEVE"
	PermissionPolicyUpdateParamsPermissionPolicyActionConfigUpdate           PermissionPolicyUpdateParamsPermissionPolicyAction = "CONFIG_UPDATE"
	PermissionPolicyUpdateParamsPermissionPolicyActionConfigDelete           PermissionPolicyUpdateParamsPermissionPolicyAction = "CONFIG_DELETE"
	PermissionPolicyUpdateParamsPermissionPolicyActionConfigExport           PermissionPolicyUpdateParamsPermissionPolicyAction = "CONFIG_EXPORT"
	PermissionPolicyUpdateParamsPermissionPolicyActionAnalyticsQuery         PermissionPolicyUpdateParamsPermissionPolicyAction = "ANALYTICS_QUERY"
	PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsUpload     PermissionPolicyUpdateParamsPermissionPolicyAction = "MEASUREMENTS_UPLOAD"
	PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsFileupload PermissionPolicyUpdateParamsPermissionPolicyAction = "MEASUREMENTS_FILEUPLOAD"
	PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsRetrieve   PermissionPolicyUpdateParamsPermissionPolicyAction = "MEASUREMENTS_RETRIEVE"
	PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsExport     PermissionPolicyUpdateParamsPermissionPolicyAction = "MEASUREMENTS_EXPORT"
	PermissionPolicyUpdateParamsPermissionPolicyActionForecastRetrieve       PermissionPolicyUpdateParamsPermissionPolicyAction = "FORECAST_RETRIEVE"
	PermissionPolicyUpdateParamsPermissionPolicyActionHealthscoresRetrieve   PermissionPolicyUpdateParamsPermissionPolicyAction = "HEALTHSCORES_RETRIEVE"
	PermissionPolicyUpdateParamsPermissionPolicyActionAnomaliesRetrieve      PermissionPolicyUpdateParamsPermissionPolicyAction = "ANOMALIES_RETRIEVE"
	PermissionPolicyUpdateParamsPermissionPolicyActionExportsDownload        PermissionPolicyUpdateParamsPermissionPolicyAction = "EXPORTS_DOWNLOAD"
)

func (r PermissionPolicyUpdateParamsPermissionPolicyAction) IsKnown() bool {
	switch r {
	case PermissionPolicyUpdateParamsPermissionPolicyActionAll, PermissionPolicyUpdateParamsPermissionPolicyActionConfigCreate, PermissionPolicyUpdateParamsPermissionPolicyActionConfigRetrieve, PermissionPolicyUpdateParamsPermissionPolicyActionConfigUpdate, PermissionPolicyUpdateParamsPermissionPolicyActionConfigDelete, PermissionPolicyUpdateParamsPermissionPolicyActionConfigExport, PermissionPolicyUpdateParamsPermissionPolicyActionAnalyticsQuery, PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsUpload, PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsFileupload, PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsRetrieve, PermissionPolicyUpdateParamsPermissionPolicyActionMeasurementsExport, PermissionPolicyUpdateParamsPermissionPolicyActionForecastRetrieve, PermissionPolicyUpdateParamsPermissionPolicyActionHealthscoresRetrieve, PermissionPolicyUpdateParamsPermissionPolicyActionAnomaliesRetrieve, PermissionPolicyUpdateParamsPermissionPolicyActionExportsDownload:
		return true
	}
	return false
}

// Specifies whether or not the user is allowed to perform the action on the
// resource.
//
// **NOTE:** Use lower case, for example: `"allow"`. If you use upper case, you'll
// receive an error.
type PermissionPolicyUpdateParamsPermissionPolicyEffect string

const (
	PermissionPolicyUpdateParamsPermissionPolicyEffectAllow PermissionPolicyUpdateParamsPermissionPolicyEffect = "ALLOW"
	PermissionPolicyUpdateParamsPermissionPolicyEffectDeny  PermissionPolicyUpdateParamsPermissionPolicyEffect = "DENY"
)

func (r PermissionPolicyUpdateParamsPermissionPolicyEffect) IsKnown() bool {
	switch r {
	case PermissionPolicyUpdateParamsPermissionPolicyEffectAllow, PermissionPolicyUpdateParamsPermissionPolicyEffectDeny:
		return true
	}
	return false
}

type PermissionPolicyListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of permission polices to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [PermissionPolicyListParams]'s query parameters as
// `url.Values`.
func (r PermissionPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PermissionPolicyDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type PermissionPolicyAddToServiceUserParams struct {
	OrgID       param.Field[string] `path:"orgId,required"`
	PrincipalID param.Field[string] `json:"principalId,required"`
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

func (r PermissionPolicyAddToServiceUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyAddToSupportUserParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
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

func (r PermissionPolicyAddToSupportUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyAddToUserParams struct {
	OrgID       param.Field[string] `path:"orgId,required"`
	PrincipalID param.Field[string] `json:"principalId,required"`
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

func (r PermissionPolicyAddToUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyAddToUserGroupParams struct {
	OrgID       param.Field[string] `path:"orgId,required"`
	PrincipalID param.Field[string] `json:"principalId,required"`
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

func (r PermissionPolicyAddToUserGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyRemoveFromServiceUserParams struct {
	OrgID       param.Field[string] `path:"orgId,required"`
	PrincipalID param.Field[string] `json:"principalId,required"`
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

func (r PermissionPolicyRemoveFromServiceUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyRemoveFromSupportUserParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type PermissionPolicyRemoveFromUserParams struct {
	OrgID       param.Field[string] `path:"orgId,required"`
	PrincipalID param.Field[string] `json:"principalId,required"`
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

func (r PermissionPolicyRemoveFromUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionPolicyRemoveFromUserGroupParams struct {
	OrgID       param.Field[string] `path:"orgId,required"`
	PrincipalID param.Field[string] `json:"principalId,required"`
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

func (r PermissionPolicyRemoveFromUserGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
