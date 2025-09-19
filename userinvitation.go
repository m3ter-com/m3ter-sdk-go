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
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
)

// UserInvitationService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserInvitationService] method instead.
type UserInvitationService struct {
	Options []option.RequestOption
}

// NewUserInvitationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserInvitationService(opts ...option.RequestOption) (r *UserInvitationService) {
	r = &UserInvitationService{}
	r.Options = opts
	return
}

// Invite a new user to your Organization.
//
// This sends an email to someone inviting them to join your m3ter Organization.
func (r *UserInvitationService) New(ctx context.Context, params UserInvitationNewParams, opts ...option.RequestOption) (res *InvitationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/invitations", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the specified invitation with the given UUID.
func (r *UserInvitationService) Get(ctx context.Context, id string, query UserInvitationGetParams, opts ...option.RequestOption) (res *InvitationResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/invitations/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of all invitations in the Organization.
func (r *UserInvitationService) List(ctx context.Context, params UserInvitationListParams, opts ...option.RequestOption) (res *pagination.Cursor[InvitationResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/invitations", params.OrgID)
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

// Retrieve a list of all invitations in the Organization.
func (r *UserInvitationService) ListAutoPaging(ctx context.Context, params UserInvitationListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[InvitationResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

type InvitationResponse struct {
	// The UUID of the invitation.
	ID string `json:"id,required"`
	// Boolean indicating whether the user has accepted the invitation.
	//
	// - TRUE - the invite has been accepted.
	// - FALSE - the invite has not yet been accepted.
	Accepted bool `json:"accepted,required"`
	// The date that access will end for the user _(in ISO-8601 format)_. If this is
	// blank, there is no end date meaning that the user has permanent access.
	DtEndAccess time.Time `json:"dtEndAccess,required" format:"date-time"`
	// The date when the invite expires _(in ISO-8601 format)_. After this date the
	// invited user can no longer accept the invite. By default, any invite is valid
	// for 30 days from the date the invite is sent.
	DtExpiry time.Time `json:"dtExpiry,required" format:"date-time"`
	// The email address of the invitee. The invitation will be sent to this email
	// address.
	Email string `json:"email,required"`
	// The first name of the invitee.
	FirstName string `json:"firstName,required"`
	// The UUID of the user who sent the invite.
	InvitingPrincipalID string `json:"invitingPrincipalId,required"`
	// The surname of the invitee.
	LastName string `json:"lastName,required"`
	// The IDs of the permission policies the invited user has been assigned. This
	// controls the access rights and privileges that this user will have when working
	// in the m3ter Organization.
	PermissionPolicyIDs []string `json:"permissionPolicyIds,required"`
	// The version number. Default value when newly created is one.
	Version int64 `json:"version,required"`
	// The UUID of the user who created the invitation.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the invitation was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the invitation was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The UUID of the user who last modified the invitation.
	LastModifiedBy string                 `json:"lastModifiedBy"`
	JSON           invitationResponseJSON `json:"-"`
}

// invitationResponseJSON contains the JSON metadata for the struct
// [InvitationResponse]
type invitationResponseJSON struct {
	ID                  apijson.Field
	Accepted            apijson.Field
	DtEndAccess         apijson.Field
	DtExpiry            apijson.Field
	Email               apijson.Field
	FirstName           apijson.Field
	InvitingPrincipalID apijson.Field
	LastName            apijson.Field
	PermissionPolicyIDs apijson.Field
	Version             apijson.Field
	CreatedBy           apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	LastModifiedBy      apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvitationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r invitationResponseJSON) RawJSON() string {
	return r.raw
}

type UserInvitationNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID         param.Field[string] `path:"orgId,required"`
	Email         param.Field[string] `json:"email,required" format:"email"`
	FirstName     param.Field[string] `json:"firstName,required"`
	LastName      param.Field[string] `json:"lastName,required"`
	ContactNumber param.Field[string] `json:"contactNumber"`
	// The date when access will end for the user _(in ISO-8601 format)_. Leave blank
	// for no end date, which gives the user permanent access.
	DtEndAccess param.Field[time.Time] `json:"dtEndAccess" format:"date-time"`
	// The date when the invite expires _(in ISO-8601 format)_. After this date the
	// invited user can no longer accept the invite. By default, any invite is valid
	// for 30 days from the date the invite is sent.
	DtExpiry  param.Field[time.Time] `json:"dtExpiry" format:"date-time"`
	M3terUser param.Field[bool]      `json:"m3terUser"`
	// The IDs of the permission policies the invited user has been assigned. This
	// controls the access rights and privileges that this user will have when working
	// in the m3ter Organization.
	PermissionPolicyIDs param.Field[[]string] `json:"permissionPolicyIds"`
	Version             param.Field[int64]    `json:"version"`
}

func (r UserInvitationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserInvitationGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type UserInvitationListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of invitations to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [UserInvitationListParams]'s query parameters as
// `url.Values`.
func (r UserInvitationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
