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
	"github.com/m3ter-com/m3ter-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// PlanService contains methods and other services that help with interacting with
// the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanService] method instead.
type PlanService struct {
	Options []option.RequestOption
}

// NewPlanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPlanService(opts ...option.RequestOption) (r *PlanService) {
	r = &PlanService{}
	r.Options = opts
	return
}

// Create a new Plan.
func (r *PlanService) New(ctx context.Context, params PlanNewParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plans", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Plan with the given UUID.
func (r *PlanService) Get(ctx context.Context, id string, query PlanGetParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plans/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Plan with the given UUID.
//
// **Note:** If you have created Custom Fields for a Plan, when you use this
// endpoint to update the Plan use the `customFields` parameter to preserve those
// Custom Fields. If you omit them from the update request, they will be lost.
func (r *PlanService) Update(ctx context.Context, id string, params PlanUpdateParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plans/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Plans that can be filtered by Product, Account, or Plan ID.
func (r *PlanService) List(ctx context.Context, params PlanListParams, opts ...option.RequestOption) (res *pagination.Cursor[Plan], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plans", params.OrgID)
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

// Retrieve a list of Plans that can be filtered by Product, Account, or Plan ID.
func (r *PlanService) ListAutoPaging(ctx context.Context, params PlanListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Plan] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Plan with the given UUID.
func (r *PlanService) Delete(ctx context.Context, id string, body PlanDeleteParams, opts ...option.RequestOption) (res *Plan, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plans/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Plan struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// _(Optional)_. The Account ID for which this plan was created as custom/bespoke.
	// A custom/bespoke Plan can only be attached to the specified Account.
	AccountID string `json:"accountId"`
	// TRUE/FALSE flag indicating whether the plan is custom/bespoke for a particular
	// Account.
	Bespoke bool `json:"bespoke"`
	// Unique short code reference for the Plan.
	Code string `json:"code"`
	// The id of the user who created this plan.
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
	CustomFields map[string]PlanCustomFieldsUnion `json:"customFields"`
	// The DateTime _(in ISO-8601 format)_ when the plan was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the plan was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this plan.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The product minimum spend amount per billing cycle for end customer Accounts on
	// a priced Plan.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	MinimumSpend float64 `json:"minimumSpend"`
	// Optional Product ID this plan's minimum spend should be attributed to for
	// accounting purposes
	MinimumSpendAccountingProductID string `json:"minimumSpendAccountingProductId"`
	// When TRUE, minimum spend is billed at the start of each billing period.
	//
	// When FALSE, minimum spend is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at PlanTemplate level for minimum spend
	// billing in arrears/in advance.
	MinimumSpendBillInAdvance bool `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription string `json:"minimumSpendDescription"`
	// Descriptive name for the Plan.
	Name string `json:"name"`
	// Assigns a rank or position to the Plan in your order of pricing plans - lower
	// numbers represent more basic pricing plans; higher numbers represent more
	// premium pricing plans.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	//
	// **NOTE:** **DEPRECATED** - no longer used.
	Ordinal int64 `json:"ordinal"`
	// UUID of the PlanTemplate the Plan belongs to.
	PlanTemplateID string `json:"planTemplateId"`
	// UUID of the Product the Plan belongs to.
	ProductID string `json:"productId"`
	// The standing charge applied to bills for end customers. This is prorated.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	StandingCharge float64 `json:"standingCharge"`
	// Optional Product ID this plan's standing charge should be attributed to for
	// accounting purposes
	StandingChargeAccountingProductID string `json:"standingChargeAccountingProductId"`
	// When TRUE, standing charge is billed at the start of each billing period.
	//
	// When FALSE, standing charge is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at PlanTemplate level for standing charge
	// billing in arrears/in advance.
	StandingChargeBillInAdvance bool `json:"standingChargeBillInAdvance"`
	// Standing charge description _(displayed on the bill line item)_.
	StandingChargeDescription string   `json:"standingChargeDescription"`
	JSON                      planJSON `json:"-"`
}

// planJSON contains the JSON metadata for the struct [Plan]
type planJSON struct {
	ID                                apijson.Field
	Version                           apijson.Field
	AccountID                         apijson.Field
	Bespoke                           apijson.Field
	Code                              apijson.Field
	CreatedBy                         apijson.Field
	CustomFields                      apijson.Field
	DtCreated                         apijson.Field
	DtLastModified                    apijson.Field
	LastModifiedBy                    apijson.Field
	MinimumSpend                      apijson.Field
	MinimumSpendAccountingProductID   apijson.Field
	MinimumSpendBillInAdvance         apijson.Field
	MinimumSpendDescription           apijson.Field
	Name                              apijson.Field
	Ordinal                           apijson.Field
	PlanTemplateID                    apijson.Field
	ProductID                         apijson.Field
	StandingCharge                    apijson.Field
	StandingChargeAccountingProductID apijson.Field
	StandingChargeBillInAdvance       apijson.Field
	StandingChargeDescription         apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *Plan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type PlanCustomFieldsUnion interface {
	ImplementsPlanCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanCustomFieldsUnion)(nil)).Elem(),
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

type PlanNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Unique short code reference for the Plan.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name for the Plan.
	Name param.Field[string] `json:"name,required"`
	// UUID of the PlanTemplate the Plan belongs to.
	PlanTemplateID param.Field[string] `json:"planTemplateId,required"`
	// _(Optional)_. Used to specify an Account for which the Plan will be a
	// custom/bespoke Plan:
	//
	//   - Use when first creating a Plan.
	//   - A custom/bespoke Plan can only be attached to the specified Account.
	//   - Once created, a custom/bespoke Plan cannot be updated to be made a
	//     custom/bespoke Plan for a different Account.
	AccountID param.Field[string] `json:"accountId"`
	// TRUE/FALSE flag indicating whether the plan is a custom/bespoke Plan for a
	// particular Account:
	//
	//   - When creating a Plan, use the `accountId` request parameter to specify the
	//     Account for which the Plan will be custom/bespoke.
	//   - A custom/bespoke Plan can only be attached to the specified Account.
	Bespoke param.Field[bool] `json:"bespoke"`
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
	CustomFields param.Field[map[string]PlanNewParamsCustomFieldsUnion] `json:"customFields"`
	// The product minimum spend amount per billing cycle for end customer Accounts on
	// a priced Plan.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// Optional Product ID this plan's minimum spend should be attributed to for
	// accounting purposes
	MinimumSpendAccountingProductID param.Field[string] `json:"minimumSpendAccountingProductId"`
	// When TRUE, minimum spend is billed at the start of each billing period.
	//
	// When FALSE, minimum spend is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at PlanTemplate level for minimum spend
	// billing in arrears/in advance.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// Assigns a rank or position to the Plan in your order of pricing plans - lower
	// numbers represent more basic pricing plans; higher numbers represent more
	// premium pricing plans.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	//
	// **NOTE: DEPRECATED** - do not use.
	Ordinal param.Field[int64] `json:"ordinal"`
	// The standing charge applied to bills for end customers. This is prorated.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	StandingCharge param.Field[float64] `json:"standingCharge"`
	// Optional Product ID this plan's standing charge should be attributed to for
	// accounting purposes
	StandingChargeAccountingProductID param.Field[string] `json:"standingChargeAccountingProductId"`
	// When TRUE, standing charge is billed at the start of each billing period.
	//
	// When FALSE, standing charge is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at PlanTemplate level for standing charge
	// billing in arrears/in advance.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Standing charge description _(displayed on the bill line item)_.
	StandingChargeDescription param.Field[string] `json:"standingChargeDescription"`
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

func (r PlanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PlanNewParamsCustomFieldsUnion interface {
	ImplementsPlanNewParamsCustomFieldsUnion()
}

type PlanGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type PlanUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Unique short code reference for the Plan.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name for the Plan.
	Name param.Field[string] `json:"name,required"`
	// UUID of the PlanTemplate the Plan belongs to.
	PlanTemplateID param.Field[string] `json:"planTemplateId,required"`
	// _(Optional)_. Used to specify an Account for which the Plan will be a
	// custom/bespoke Plan:
	//
	//   - Use when first creating a Plan.
	//   - A custom/bespoke Plan can only be attached to the specified Account.
	//   - Once created, a custom/bespoke Plan cannot be updated to be made a
	//     custom/bespoke Plan for a different Account.
	AccountID param.Field[string] `json:"accountId"`
	// TRUE/FALSE flag indicating whether the plan is a custom/bespoke Plan for a
	// particular Account:
	//
	//   - When creating a Plan, use the `accountId` request parameter to specify the
	//     Account for which the Plan will be custom/bespoke.
	//   - A custom/bespoke Plan can only be attached to the specified Account.
	Bespoke param.Field[bool] `json:"bespoke"`
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
	CustomFields param.Field[map[string]PlanUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// The product minimum spend amount per billing cycle for end customer Accounts on
	// a priced Plan.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// Optional Product ID this plan's minimum spend should be attributed to for
	// accounting purposes
	MinimumSpendAccountingProductID param.Field[string] `json:"minimumSpendAccountingProductId"`
	// When TRUE, minimum spend is billed at the start of each billing period.
	//
	// When FALSE, minimum spend is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at PlanTemplate level for minimum spend
	// billing in arrears/in advance.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// Assigns a rank or position to the Plan in your order of pricing plans - lower
	// numbers represent more basic pricing plans; higher numbers represent more
	// premium pricing plans.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	//
	// **NOTE: DEPRECATED** - do not use.
	Ordinal param.Field[int64] `json:"ordinal"`
	// The standing charge applied to bills for end customers. This is prorated.
	//
	// _(Optional)_. Overrides PlanTemplate value.
	StandingCharge param.Field[float64] `json:"standingCharge"`
	// Optional Product ID this plan's standing charge should be attributed to for
	// accounting purposes
	StandingChargeAccountingProductID param.Field[string] `json:"standingChargeAccountingProductId"`
	// When TRUE, standing charge is billed at the start of each billing period.
	//
	// When FALSE, standing charge is billed at the end of each billing period.
	//
	// _(Optional)_. Overrides the setting at PlanTemplate level for standing charge
	// billing in arrears/in advance.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Standing charge description _(displayed on the bill line item)_.
	StandingChargeDescription param.Field[string] `json:"standingChargeDescription"`
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

func (r PlanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PlanUpdateParamsCustomFieldsUnion interface {
	ImplementsPlanUpdateParamsCustomFieldsUnion()
}

type PlanListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// List of Account IDs the Plan belongs to.
	AccountID param.Field[[]string] `query:"accountId"`
	// List of Plan IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi-page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Plans to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// UUID of the Product to retrieve Plans for.
	ProductID param.Field[string] `query:"productId"`
}

// URLQuery serializes [PlanListParams]'s query parameters as `url.Values`.
func (r PlanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
