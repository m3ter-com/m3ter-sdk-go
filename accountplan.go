// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// AccountPlanService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPlanService] method instead.
type AccountPlanService struct {
	Options []option.RequestOption
}

// NewAccountPlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountPlanService(opts ...option.RequestOption) (r *AccountPlanService) {
	r = &AccountPlanService{}
	r.Options = opts
	return
}

// Create a new AccountPlan or AccountPlanGroup.
//
// This endpoint creates a new AccountPlan or AccountPlanGroup for a specific
// Account in your Organization. The details of the new AccountPlan or
// AccountPlanGroup should be supplied in the request body.
//
// **Note:** You cannot use this call to create _both_ an AccountPlan and
// AccountPlanGroup for an Account at the same time. If you want to create both for
// an Account, you must submit two separate calls.
func (r *AccountPlanService) New(ctx context.Context, params AccountPlanNewParams, opts ...option.RequestOption) (res *AccountPlanResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accountplans", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the AccountPlan or AccountPlanGroup details corresponding to the given
// UUID.
func (r *AccountPlanService) Get(ctx context.Context, id string, query AccountPlanGetParams, opts ...option.RequestOption) (res *AccountPlanResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accountplans/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the AccountPlan or AccountPlanGroup with the given UUID.
//
// This endpoint updates a new AccountPlan or AccountPlanGroup for a specific
// Account in your Organization. The updated information should be provided in the
// request body.
//
// **Notes:**
//
//   - You cannot use this call to update _both_ an AccountPlan and AccountPlanGroup
//     for an Account at the same time. If you want to update an AccounPlan and an
//     AccountPlanGroup attached to an Account, you must submit two separate calls.
//   - If you have created Custom Fields for an AccountPlan, when you use this
//     endpoint to update the AccountPlan use the `customFields` parameter to
//     preserve those Custom Fields. If you omit them from the update request, they
//     will be lost.
func (r *AccountPlanService) Update(ctx context.Context, id string, params AccountPlanUpdateParams, opts ...option.RequestOption) (res *AccountPlanResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/accountplans/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieves a list of AccountPlan and AccountPlanGroup entities for the specified
// Organization. The list can be paginated for easier management, and supports
// filtering with various query parameters.
func (r *AccountPlanService) List(ctx context.Context, params AccountPlanListParams, opts ...option.RequestOption) (res *pagination.Cursor[AccountPlanResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/accountplans", params.OrgID)
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

// Retrieves a list of AccountPlan and AccountPlanGroup entities for the specified
// Organization. The list can be paginated for easier management, and supports
// filtering with various query parameters.
func (r *AccountPlanService) ListAutoPaging(ctx context.Context, params AccountPlanListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AccountPlanResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the AccountPlan or AccountPlanGroup with the given UUID.
//
// This endpoint deletes an AccountPlan or AccountPlanGroup that has been attached
// to a specific Account in your Organization.
func (r *AccountPlanService) Delete(ctx context.Context, id string, body AccountPlanDeleteParams, opts ...option.RequestOption) (res *AccountPlanResponse, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/accountplans/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountPlanResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The unique identifier (UUID) for the Account to which the AccountPlan or
	// AccounPlanGroup is attached.
	AccountID string `json:"accountId"`
	// The initial date for creating the first bill against the Account for charges due
	// under the AccountPlan or AccountPlanGroup. All subsequent bill creation dates
	// are calculated from this date. If left empty, the first bill date definedfor the
	// Account is used. The date is in ISO-8601 format.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// If the Account is either a Parent or a Child Account, this specifies the Account
	// hierarchy billing mode. The mode determines how billing will be handled and
	// shown on bills for charges due on the Parent Account, and charges due on Child
	// Accounts:
	//
	// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
	//
	// - **Parent Summary** - single bill line item for all Accounts.
	//
	// - **Child** - the Child Account is billed.
	ChildBillingMode AccountPlanResponseChildBillingMode `json:"childBillingMode"`
	// The unique short code of the AccountPlan or AccountPlanGroup.
	Code string `json:"code"`
	// The unique identifier (UUID) for the Contract to which the Plan or Plan Group
	// attached to the Account has been added.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) for the user who created the AccountPlan or
	// AccountPlanGroup.
	CreatedBy string `json:"createdBy"`
	// User defined fields enabling you to attach custom data. The value for a custom
	// field can be either a string or a number.
	//
	// If `customFields` can also be defined for this entity at the Organizational
	// level,`customField` values defined at individual level override values of
	// `customFields` with the same name defined at Organization level.
	//
	// See
	// [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields)
	// in the m3ter documentation for more information.
	CustomFields map[string]AccountPlanResponseCustomFieldsUnion `json:"customFields"`
	// The date and time _(in ISO 8601 format)_ when the AccountPlan or
	// AccountPlanGroup was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the AccountPlan or
	// AccountPlanGroup was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The end date _(in ISO-8601 format)_ for when the AccountPlan or AccountPlanGroup
	// ceases to be active for the Account. If not specified, the AccountPlan or
	// AccountPlanGroup remains active indefinitely.
	EndDate time.Time `json:"endDate" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the AccountPlan or
	// AccountPlanGroup.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The unique identifier (UUID) of the Plan Group that has been attached to the
	// Account to create the AccountPlanGroup.
	PlanGroupID string `json:"planGroupId"`
	// The unique identifier (UUID) of the Plan that has been attached to the Account
	// to create the AccountPlan.
	PlanID string `json:"planId"`
	// The unique identifier (UUID) for the Product associated with the AccountPlan.
	//
	// **Note:** Not present in response for AccountPlanGroup - Plan Groups can contain
	// multiple Plans belonging to different Products.
	ProductID string `json:"productId"`
	// The start date _(in ISO-8601 format)_ for the when the AccountPlan or
	// AccountPlanGroup starts to be active for the Account.
	StartDate time.Time `json:"startDate" format:"date-time"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                   `json:"version"`
	JSON    accountPlanResponseJSON `json:"-"`
}

// accountPlanResponseJSON contains the JSON metadata for the struct
// [AccountPlanResponse]
type accountPlanResponseJSON struct {
	ID               apijson.Field
	AccountID        apijson.Field
	BillEpoch        apijson.Field
	ChildBillingMode apijson.Field
	Code             apijson.Field
	ContractID       apijson.Field
	CreatedBy        apijson.Field
	CustomFields     apijson.Field
	DtCreated        apijson.Field
	DtLastModified   apijson.Field
	EndDate          apijson.Field
	LastModifiedBy   apijson.Field
	PlanGroupID      apijson.Field
	PlanID           apijson.Field
	ProductID        apijson.Field
	StartDate        apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPlanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPlanResponseJSON) RawJSON() string {
	return r.raw
}

// If the Account is either a Parent or a Child Account, this specifies the Account
// hierarchy billing mode. The mode determines how billing will be handled and
// shown on bills for charges due on the Parent Account, and charges due on Child
// Accounts:
//
// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
//
// - **Parent Summary** - single bill line item for all Accounts.
//
// - **Child** - the Child Account is billed.
type AccountPlanResponseChildBillingMode string

const (
	AccountPlanResponseChildBillingModeParentSummary   AccountPlanResponseChildBillingMode = "PARENT_SUMMARY"
	AccountPlanResponseChildBillingModeParentBreakdown AccountPlanResponseChildBillingMode = "PARENT_BREAKDOWN"
	AccountPlanResponseChildBillingModeChild           AccountPlanResponseChildBillingMode = "CHILD"
)

func (r AccountPlanResponseChildBillingMode) IsKnown() bool {
	switch r {
	case AccountPlanResponseChildBillingModeParentSummary, AccountPlanResponseChildBillingModeParentBreakdown, AccountPlanResponseChildBillingModeChild:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type AccountPlanResponseCustomFieldsUnion interface {
	ImplementsAccountPlanResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountPlanResponseCustomFieldsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type AccountPlanNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) for the Account.
	AccountID param.Field[string] `json:"accountId,required"`
	// The start date _(in ISO-8601 format)_ for the AccountPlan or AccountPlanGroup
	// becoming active for the Account.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// Optional setting to define a _billing cycle date_, which acts as a reference for
	// when in the applied billing frequency period bills are created:
	//
	//   - For example, if you attach a Plan to an Account where the Plan is configured
	//     for monthly billing frequency and you've defined the period the Plan will
	//     apply to the Account to be from January 1st, 2022 until January 1st, 2023. You
	//     then set a `billEpoch` date of February 15th, 2022. The first Bill will be
	//     created for the Account on February 15th, and subsequent Bills created on the
	//     15th of the months following for the remainder of the billing period - March
	//     15th, April 15th, and so on.
	//   - If not defined, then the `billEpoch` date set for the Account will be used
	//     instead.
	//   - The date is in ISO-8601 format.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
	// If the Account is either a Parent or a Child Account, this specifies the Account
	// hierarchy billing mode. The mode determines how billing will be handled and
	// shown on bills for charges due on the Parent Account, and charges due on Child
	// Accounts:
	//
	// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
	//
	// - **Parent Summary** - single bill line item for all Accounts.
	//
	// - **Child** - the Child Account is billed.
	ChildBillingMode param.Field[AccountPlanNewParamsChildBillingMode] `json:"childBillingMode"`
	// A unique short code for the AccountPlan or AccountPlanGroup.
	Code param.Field[string] `json:"code"`
	// The unique identifier (UUID) for a Contract to which you want to add the Plan or
	// Plan Group being attached to the Account.
	ContractID param.Field[string] `json:"contractId"`
	// User defined fields enabling you to attach custom data. The value for a custom
	// field can be either a string or a number.
	//
	// If `customFields` can also be defined for this entity at the Organizational
	// level, `customField` values defined at individual level override values of
	// `customFields` with the same name defined at Organization level.
	//
	// See
	// [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields)
	// in the m3ter documentation for more information.
	CustomFields param.Field[map[string]AccountPlanNewParamsCustomFieldsUnion] `json:"customFields"`
	// The end date _(in ISO-8601 format)_ for when the AccountPlan or AccountPlanGroup
	// ceases to be active for the Account. If not specified, the AccountPlan or
	// AccountPlanGroup remains active indefinitely.
	EndDate param.Field[time.Time] `json:"endDate" format:"date-time"`
	// The unique identifier (UUID) of the PlanGroup to be attached to the Account to
	// create an AccountPlanGroup.
	//
	// **Note:** Exclusive of the `planId` request parameter - exactly one of `planId`
	// or `planGroupId` must be used per call.
	PlanGroupID param.Field[string] `json:"planGroupId"`
	// The unique identifier (UUID) of the Plan to be attached to the Account to create
	// an AccountPlan.
	//
	// **Note:** Exclusive of the `planGroupId` request parameter - exactly one of
	// `planId` or `planGroupId` must be used per call.
	PlanID param.Field[string] `json:"planId"`
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

func (r AccountPlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Account is either a Parent or a Child Account, this specifies the Account
// hierarchy billing mode. The mode determines how billing will be handled and
// shown on bills for charges due on the Parent Account, and charges due on Child
// Accounts:
//
// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
//
// - **Parent Summary** - single bill line item for all Accounts.
//
// - **Child** - the Child Account is billed.
type AccountPlanNewParamsChildBillingMode string

const (
	AccountPlanNewParamsChildBillingModeParentSummary   AccountPlanNewParamsChildBillingMode = "PARENT_SUMMARY"
	AccountPlanNewParamsChildBillingModeParentBreakdown AccountPlanNewParamsChildBillingMode = "PARENT_BREAKDOWN"
	AccountPlanNewParamsChildBillingModeChild           AccountPlanNewParamsChildBillingMode = "CHILD"
)

func (r AccountPlanNewParamsChildBillingMode) IsKnown() bool {
	switch r {
	case AccountPlanNewParamsChildBillingModeParentSummary, AccountPlanNewParamsChildBillingModeParentBreakdown, AccountPlanNewParamsChildBillingModeChild:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type AccountPlanNewParamsCustomFieldsUnion interface {
	ImplementsAccountPlanNewParamsCustomFieldsUnion()
}

type AccountPlanGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type AccountPlanUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) for the Account.
	AccountID param.Field[string] `json:"accountId,required"`
	// The start date _(in ISO-8601 format)_ for the AccountPlan or AccountPlanGroup
	// becoming active for the Account.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// Optional setting to define a _billing cycle date_, which acts as a reference for
	// when in the applied billing frequency period bills are created:
	//
	//   - For example, if you attach a Plan to an Account where the Plan is configured
	//     for monthly billing frequency and you've defined the period the Plan will
	//     apply to the Account to be from January 1st, 2022 until January 1st, 2023. You
	//     then set a `billEpoch` date of February 15th, 2022. The first Bill will be
	//     created for the Account on February 15th, and subsequent Bills created on the
	//     15th of the months following for the remainder of the billing period - March
	//     15th, April 15th, and so on.
	//   - If not defined, then the `billEpoch` date set for the Account will be used
	//     instead.
	//   - The date is in ISO-8601 format.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
	// If the Account is either a Parent or a Child Account, this specifies the Account
	// hierarchy billing mode. The mode determines how billing will be handled and
	// shown on bills for charges due on the Parent Account, and charges due on Child
	// Accounts:
	//
	// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
	//
	// - **Parent Summary** - single bill line item for all Accounts.
	//
	// - **Child** - the Child Account is billed.
	ChildBillingMode param.Field[AccountPlanUpdateParamsChildBillingMode] `json:"childBillingMode"`
	// A unique short code for the AccountPlan or AccountPlanGroup.
	Code param.Field[string] `json:"code"`
	// The unique identifier (UUID) for a Contract to which you want to add the Plan or
	// Plan Group being attached to the Account.
	ContractID param.Field[string] `json:"contractId"`
	// User defined fields enabling you to attach custom data. The value for a custom
	// field can be either a string or a number.
	//
	// If `customFields` can also be defined for this entity at the Organizational
	// level, `customField` values defined at individual level override values of
	// `customFields` with the same name defined at Organization level.
	//
	// See
	// [Working with Custom Fields](https://www.m3ter.com/docs/guides/creating-and-managing-products/working-with-custom-fields)
	// in the m3ter documentation for more information.
	CustomFields param.Field[map[string]AccountPlanUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// The end date _(in ISO-8601 format)_ for when the AccountPlan or AccountPlanGroup
	// ceases to be active for the Account. If not specified, the AccountPlan or
	// AccountPlanGroup remains active indefinitely.
	EndDate param.Field[time.Time] `json:"endDate" format:"date-time"`
	// The unique identifier (UUID) of the PlanGroup to be attached to the Account to
	// create an AccountPlanGroup.
	//
	// **Note:** Exclusive of the `planId` request parameter - exactly one of `planId`
	// or `planGroupId` must be used per call.
	PlanGroupID param.Field[string] `json:"planGroupId"`
	// The unique identifier (UUID) of the Plan to be attached to the Account to create
	// an AccountPlan.
	//
	// **Note:** Exclusive of the `planGroupId` request parameter - exactly one of
	// `planId` or `planGroupId` must be used per call.
	PlanID param.Field[string] `json:"planId"`
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

func (r AccountPlanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the Account is either a Parent or a Child Account, this specifies the Account
// hierarchy billing mode. The mode determines how billing will be handled and
// shown on bills for charges due on the Parent Account, and charges due on Child
// Accounts:
//
// - **Parent Breakdown** - a separate bill line item per Account. Default setting.
//
// - **Parent Summary** - single bill line item for all Accounts.
//
// - **Child** - the Child Account is billed.
type AccountPlanUpdateParamsChildBillingMode string

const (
	AccountPlanUpdateParamsChildBillingModeParentSummary   AccountPlanUpdateParamsChildBillingMode = "PARENT_SUMMARY"
	AccountPlanUpdateParamsChildBillingModeParentBreakdown AccountPlanUpdateParamsChildBillingMode = "PARENT_BREAKDOWN"
	AccountPlanUpdateParamsChildBillingModeChild           AccountPlanUpdateParamsChildBillingMode = "CHILD"
)

func (r AccountPlanUpdateParamsChildBillingMode) IsKnown() bool {
	switch r {
	case AccountPlanUpdateParamsChildBillingModeParentSummary, AccountPlanUpdateParamsChildBillingModeParentBreakdown, AccountPlanUpdateParamsChildBillingModeChild:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type AccountPlanUpdateParamsCustomFieldsUnion interface {
	ImplementsAccountPlanUpdateParamsCustomFieldsUnion()
}

type AccountPlanListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The unique identifier (UUID) for the Account whose AccountPlans and
	// AccountPlanGroups you want to retrieve.
	//
	// **NOTE:** Only returns the currently active AccountPlans and AccountPlanGroups
	// for the specified Account. Use in combination with the `includeall` query
	// parameter to return both active and inactive.
	Account param.Field[string] `query:"account"`
	// The unique identifier (UUID) of the Contract which the AccountPlans you want to
	// retrieve have been linked to.
	//
	// **NOTE:** Does not return AccountPlanGroups that have been linked to the
	// Contract.
	Contract param.Field[string] `query:"contract"`
	// The specific date for which you want to retrieve AccountPlans and
	// AccountPlanGroups.
	//
	// **NOTE:** Returns both active and inactive AccountPlans and AccountPlanGroups
	// for the specified date.
	Date param.Field[string] `query:"date"`
	// A list of unique identifiers (UUIDs) for specific AccountPlans and
	// AccountPlanGroups you want to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// A Boolean flag that specifies whether to include both active and inactive
	// AccountPlans and AccountPlanGroups in the list.
	//
	//   - **TRUE** - both active and inactive AccountPlans and AccountPlanGroups are
	//     included in the list.
	//   - **FALSE** - only active AccountPlans and AccountPlanGroups are retrieved in
	//     the list.
	Includeall param.Field[bool] `query:"includeall"`
	// The `nextToken` for retrieving the next page of AccountPlans and
	// AccountPlanGroups. It is used to fetch the next page of AccountPlans and
	// AccountPlanGroups in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// The maximum number of AccountPlans and AccountPlanGroups to return per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// The unique identifier (UUID) for the Plan whose associated AccountPlans you want
	// to retrieve.
	//
	// **NOTE:** Does not return AccountPlanGroups if you use a `planGroupId`.
	Plan param.Field[string] `query:"plan"`
	// The unique identifier (UUID) for the Product whose associated AccountPlans you
	// want to retrieve.
	//
	// **NOTE:** You cannot use the `product` query parameter as a single filter
	// condition, but must always use it in combination with the `account` query
	// parameter.
	Product param.Field[string] `query:"product"`
}

// URLQuery serializes [AccountPlanListParams]'s query parameters as `url.Values`.
func (r AccountPlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountPlanDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
