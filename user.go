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

// UserService contains methods and other services that help with interacting with
// the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options     []option.RequestOption
	Invitations *UserInvitationService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Invitations = NewUserInvitationService(opts...)
	return
}

// Retrieve the OrgUser with the given UUID.
//
// Retrieves detailed information for a specific user within an Organization, using
// their unique identifier (UUID).
func (r *UserService) Get(ctx context.Context, id string, query UserGetParams, opts ...option.RequestOption) (res *UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the OrgUser with the given UUID.
//
// Updates the details for a specific user within an Organization using their
// unique identifier (UUID). Use this endpoint when you need to modify user
// information such as their permission policy.
func (r *UserService) Update(ctx context.Context, id string, params UserUpdateParams, opts ...option.RequestOption) (res *UserResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of OrgUsers.
//
// Retrieves a list of all users within a specified Organization. Use this endpoint
// to get an overview of all users and their basic details. The list can be
// paginated for easier management.
func (r *UserService) List(ctx context.Context, params UserListParams, opts ...option.RequestOption) (res *pagination.Cursor[UserResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users", params.OrgID)
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

// Retrieve a list of OrgUsers.
//
// Retrieves a list of all users within a specified Organization. Use this endpoint
// to get an overview of all users and their basic details. The list can be
// paginated for easier management.
func (r *UserService) ListAutoPaging(ctx context.Context, params UserListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[UserResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Retrieve the permissions for the OrgUser with the given UUID.
//
// Retrieves a list of all permissions associated with a specific user in an
// Organization using their UUID. The list can be paginated for easier management.
func (r *UserService) GetPermissions(ctx context.Context, id string, params UserGetPermissionsParams, opts ...option.RequestOption) (res *PermissionPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s/permissions", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Retrieve a list of User Groups for an OrgUser.
//
// Retrieves a list of all User Groups that a specific user belongs to within an
// Organization. The list can be paginated for easier management.
//
// **Notes:**
//
//   - **User Groups as Resource Groups**. A User Group is a Resource Group - one
//     used to group resources of type `user`. You can use the _Create ResourceGroup_
//     call detailed in the
//     [ResourceGroup](https://www.m3ter.com/docs/api#tag/ResourceGroup) section to
//     create a User Resource Group, and then use the _Add Item_ and _Remove Item_
//     calls to manage which Users belong to the User Resource Group.
//   - **Using the `inherited` parameter for the Retrieve OrgUser Groups call**.
//     Resource Groups can be nested, which means a User Resource Group can contain
//     another User Resource Group as a member. You can use the `inherited` parameter
//     with this _Retrieve OrgUser Groups_ call as a _QUERY PARAMETER_ to control
//     which User Resource Groups are returned:
//
//   - If the user specified belongs to a User Resource Group that is nested as part
//     of another User Resource Group:
//   - If `inherited = TRUE`, then any Groups the user belongs to AND any parent
//     Groups those Groups belong to as nested Groups are returned.
//   - If `inherited = FALSE`, then only those User Resource Groups to which the
//     user belongs are returned.
func (r *UserService) GetUserGroups(ctx context.Context, id string, params UserGetUserGroupsParams, opts ...option.RequestOption) (res *ResourceGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s/usergroups", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Retrieve information about the current user
func (r *UserService) Me(ctx context.Context, query UserMeParams, opts ...option.RequestOption) (res *UserMeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/me", query.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resend temporary password for user
func (r *UserService) ResendPassword(ctx context.Context, id string, body UserResendPasswordParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/users/%s/password/resend", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, nil, opts...)
	return
}

type UserResponse struct {
	// The unique identifier (UUID) of this user.
	ID string `json:"id"`
	// The user's contact telephone number.
	ContactNumber string `json:"contactNumber"`
	// The user who created this user.
	CreatedBy string `json:"createdBy"`
	// The date and time _(in ISO-8601 format)_ when the user was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the user's access will end. Used
	// to set or update the date and time a user's access expires.
	DtEndAccess time.Time `json:"dtEndAccess" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the user was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The email address for this user.
	Email string `json:"email"`
	// The date and time _(in ISO 8601 format)_ when this user first accepted the the
	// m3ter terms and conditions.
	FirstAcceptedTermsAndConditions time.Time `json:"firstAcceptedTermsAndConditions" format:"date-time"`
	// The first name of the user.
	FirstName string `json:"firstName"`
	// The date and time _(in ISO 8601 format)_ when this user last accepted the the
	// m3ter terms and conditions.
	LastAcceptedTermsAndConditions time.Time `json:"lastAcceptedTermsAndConditions" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this user record.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The surname of the user.
	LastName string `json:"lastName"`
	// An array listing the Organizations where this user has access.
	Organizations []string `json:"organizations"`
	// An array of permission statements for the user. Each permission statement
	// defines a specific permission for the user.
	PermissionPolicy []PermissionStatementResponse `json:"permissionPolicy"`
	// Indicates whether this is a m3ter Support user.
	SupportUser bool `json:"supportUser"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64            `json:"version"`
	JSON    userResponseJSON `json:"-"`
}

// userResponseJSON contains the JSON metadata for the struct [UserResponse]
type userResponseJSON struct {
	ID                              apijson.Field
	ContactNumber                   apijson.Field
	CreatedBy                       apijson.Field
	DtCreated                       apijson.Field
	DtEndAccess                     apijson.Field
	DtLastModified                  apijson.Field
	Email                           apijson.Field
	FirstAcceptedTermsAndConditions apijson.Field
	FirstName                       apijson.Field
	LastAcceptedTermsAndConditions  apijson.Field
	LastModifiedBy                  apijson.Field
	LastName                        apijson.Field
	Organizations                   apijson.Field
	PermissionPolicy                apijson.Field
	SupportUser                     apijson.Field
	Version                         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *UserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeResponse struct {
	Organization UserMeResponseOrganization `json:"organization"`
	ServiceUser  UserMeResponseServiceUser  `json:"serviceUser"`
	User         UserMeResponseUser         `json:"user"`
	JSON         userMeResponseJSON         `json:"-"`
}

// userMeResponseJSON contains the JSON metadata for the struct [UserMeResponse]
type userMeResponseJSON struct {
	Organization apijson.Field
	ServiceUser  apijson.Field
	User         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserMeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeResponseJSON) RawJSON() string {
	return r.raw
}

type UserMeResponseOrganization struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version               int64  `json:"version,required"`
	AddressLine1          string `json:"addressLine1"`
	AddressLine2          string `json:"addressLine2"`
	AddressLine3          string `json:"addressLine3"`
	AddressLine4          string `json:"addressLine4"`
	BillingContactUserId1 string `json:"billingContactUserId1"`
	BillingContactUserId2 string `json:"billingContactUserId2"`
	BillingContactUserId3 string `json:"billingContactUserId3"`
	Country               string `json:"country"`
	// The id of the user who created this organization.
	CreatedBy  string `json:"createdBy"`
	CustomerID string `json:"customerId"`
	// The DateTime when the organization was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the organization was last modified.
	DtLastModified          time.Time `json:"dtLastModified" format:"date-time"`
	InvoiceGeneralReference string    `json:"invoiceGeneralReference"`
	// The id of the user who last modified this organization.
	LastModifiedBy      string                           `json:"lastModifiedBy"`
	Locality            string                           `json:"locality"`
	OrganizationName    string                           `json:"organizationName"`
	OrgID               string                           `json:"orgId"`
	PostCode            string                           `json:"postCode"`
	PurchaseOrderNumber string                           `json:"purchaseOrderNumber"`
	Region              string                           `json:"region"`
	ShortName           string                           `json:"shortName"`
	Status              UserMeResponseOrganizationStatus `json:"status"`
	TaxID               string                           `json:"taxId"`
	Type                UserMeResponseOrganizationType   `json:"type"`
	JSON                userMeResponseOrganizationJSON   `json:"-"`
}

// userMeResponseOrganizationJSON contains the JSON metadata for the struct
// [UserMeResponseOrganization]
type userMeResponseOrganizationJSON struct {
	ID                      apijson.Field
	Version                 apijson.Field
	AddressLine1            apijson.Field
	AddressLine2            apijson.Field
	AddressLine3            apijson.Field
	AddressLine4            apijson.Field
	BillingContactUserId1   apijson.Field
	BillingContactUserId2   apijson.Field
	BillingContactUserId3   apijson.Field
	Country                 apijson.Field
	CreatedBy               apijson.Field
	CustomerID              apijson.Field
	DtCreated               apijson.Field
	DtLastModified          apijson.Field
	InvoiceGeneralReference apijson.Field
	LastModifiedBy          apijson.Field
	Locality                apijson.Field
	OrganizationName        apijson.Field
	OrgID                   apijson.Field
	PostCode                apijson.Field
	PurchaseOrderNumber     apijson.Field
	Region                  apijson.Field
	ShortName               apijson.Field
	Status                  apijson.Field
	TaxID                   apijson.Field
	Type                    apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *UserMeResponseOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeResponseOrganizationJSON) RawJSON() string {
	return r.raw
}

type UserMeResponseOrganizationStatus string

const (
	UserMeResponseOrganizationStatusActive   UserMeResponseOrganizationStatus = "ACTIVE"
	UserMeResponseOrganizationStatusArchived UserMeResponseOrganizationStatus = "ARCHIVED"
)

func (r UserMeResponseOrganizationStatus) IsKnown() bool {
	switch r {
	case UserMeResponseOrganizationStatusActive, UserMeResponseOrganizationStatusArchived:
		return true
	}
	return false
}

type UserMeResponseOrganizationType string

const (
	UserMeResponseOrganizationTypeProduction UserMeResponseOrganizationType = "PRODUCTION"
	UserMeResponseOrganizationTypeSandbox    UserMeResponseOrganizationType = "SANDBOX"
)

func (r UserMeResponseOrganizationType) IsKnown() bool {
	switch r {
	case UserMeResponseOrganizationTypeProduction, UserMeResponseOrganizationTypeSandbox:
		return true
	}
	return false
}

type UserMeResponseServiceUser struct {
	ID string `json:"id"`
	// The id of the user who created this service user.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the service user was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the service user was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this service user.
	LastModifiedBy string                        `json:"lastModifiedBy"`
	Name           string                        `json:"name"`
	Version        int64                         `json:"version"`
	JSON           userMeResponseServiceUserJSON `json:"-"`
}

// userMeResponseServiceUserJSON contains the JSON metadata for the struct
// [UserMeResponseServiceUser]
type userMeResponseServiceUserJSON struct {
	ID             apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserMeResponseServiceUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeResponseServiceUserJSON) RawJSON() string {
	return r.raw
}

type UserMeResponseUser struct {
	// The unique identifier (UUID) of this user.
	ID string `json:"id"`
	// The user's contact telephone number.
	ContactNumber string `json:"contactNumber"`
	// The user who created this user.
	CreatedBy string `json:"createdBy"`
	// The date and time _(in ISO-8601 format)_ when the user was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the user was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The email address for this user.
	Email string `json:"email"`
	// The date and time _(in ISO 8601 format)_ when this user first accepted the the
	// m3ter terms and conditions.
	FirstAcceptedTermsAndConditions time.Time `json:"firstAcceptedTermsAndConditions" format:"date-time"`
	// The first name of the user.
	FirstName string `json:"firstName"`
	// The date and time _(in ISO 8601 format)_ when this user last accepted the the
	// m3ter terms and conditions.
	LastAcceptedTermsAndConditions time.Time `json:"lastAcceptedTermsAndConditions" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this user record.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The surname of the user.
	LastName string `json:"lastName"`
	// An array listing the Organizations where this user has access.
	Organizations []string `json:"organizations"`
	// Indicates whether this is a m3ter Support user.
	SupportUser bool `json:"supportUser"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                  `json:"version"`
	JSON    userMeResponseUserJSON `json:"-"`
}

// userMeResponseUserJSON contains the JSON metadata for the struct
// [UserMeResponseUser]
type userMeResponseUserJSON struct {
	ID                              apijson.Field
	ContactNumber                   apijson.Field
	CreatedBy                       apijson.Field
	DtCreated                       apijson.Field
	DtLastModified                  apijson.Field
	Email                           apijson.Field
	FirstAcceptedTermsAndConditions apijson.Field
	FirstName                       apijson.Field
	LastAcceptedTermsAndConditions  apijson.Field
	LastModifiedBy                  apijson.Field
	LastName                        apijson.Field
	Organizations                   apijson.Field
	SupportUser                     apijson.Field
	Version                         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *UserMeResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMeResponseUserJSON) RawJSON() string {
	return r.raw
}

type UserGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type UserUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The date and time _(in ISO 8601 format)_ when the user's access will end. Use
	// this to set or update the expiration of the user's access.
	DtEndAccess param.Field[time.Time] `json:"dtEndAccess" format:"date-time"`
	// An array of permission statements for the user. Each permission statement
	// defines a specific permission for the user.
	//
	// See
	// [Understanding, Creating, and Managing Permission Policies](https://www.m3ter.com/docs/guides/organization-and-access-management/creating-and-managing-permissions)
	// for more information.
	PermissionPolicy param.Field[[]PermissionStatementResponseParam] `json:"permissionPolicy"`
	// The version number of the entity:
	//
	//   - **Newly created entity:** On initial Create, version is set at 1 and listed in
	//     the response.
	//   - **Update Entity:** On Update, version is required and must match the existing
	//     version because a check is performed to ensure sequential versioning is
	//     preserved. Version is incremented by 1 and listed in the response.
	Version param.Field[int64] `json:"version"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// list of ids to retrieve
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// OrgUsers in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of OrgUsers to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGetPermissionsParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Permission Policies in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Permission Policies to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [UserGetPermissionsParams]'s query parameters as
// `url.Values`.
func (r UserGetPermissionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserGetUserGroupsParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// User Groups in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of User Groups to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [UserGetUserGroupsParams]'s query parameters as
// `url.Values`.
func (r UserGetUserGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserMeParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type UserResendPasswordParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
