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

// PlanGroupService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanGroupService] method instead.
type PlanGroupService struct {
	Options []option.RequestOption
}

// NewPlanGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlanGroupService(opts ...option.RequestOption) (r *PlanGroupService) {
	r = &PlanGroupService{}
	r.Options = opts
	return
}

// Create a new PlanGroup. This endpoint creates a new PlanGroup within the
// specified organization.
func (r *PlanGroupService) New(ctx context.Context, params PlanGroupNewParams, opts ...option.RequestOption) (res *PlanGroup, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plangroups", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a specific PlanGroup with the given UUID.
//
// This endpoint retrieves detailed information about a specific PlanGroup
// identified by the given UUID within a specific organization.
func (r *PlanGroupService) Get(ctx context.Context, id string, query PlanGroupGetParams, opts ...option.RequestOption) (res *PlanGroup, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plangroups/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the PlanGroup with the given UUID.
//
// This endpoint updates the details of a specific PlanGroup identified by the
// given UUID within a specific organization. This allows modifications to existing
// PlanGroup attributes.
//
// **Note:** If you have created Custom Fields for a PlanGroup, when you use this
// endpoint to update the PlanGroup use the `customFields` parameter to preserve
// those Custom Fields. If you omit them from the update request, they will be
// lost.
func (r *PlanGroupService) Update(ctx context.Context, id string, params PlanGroupUpdateParams, opts ...option.RequestOption) (res *PlanGroup, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plangroups/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of PlanGroups.
//
// Retrieves a list of PlanGroups within the specified organization. You can
// optionally filter by Account IDs or PlanGroup IDs, and also paginate the results
// for easier management.
func (r *PlanGroupService) List(ctx context.Context, params PlanGroupListParams, opts ...option.RequestOption) (res *pagination.Cursor[PlanGroup], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plangroups", params.OrgID)
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

// Retrieve a list of PlanGroups.
//
// Retrieves a list of PlanGroups within the specified organization. You can
// optionally filter by Account IDs or PlanGroup IDs, and also paginate the results
// for easier management.
func (r *PlanGroupService) ListAutoPaging(ctx context.Context, params PlanGroupListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PlanGroup] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a PlanGroup with the given UUID.
//
// This endpoint deletes the PlanGroup identified by the given UUID within a
// specific organization. This operation is irreversible and removes the PlanGroup
// along with any associated settings.
func (r *PlanGroupService) Delete(ctx context.Context, id string, body PlanGroupDeleteParams, opts ...option.RequestOption) (res *PlanGroup, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/plangroups/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PlanGroup struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Optional. This PlanGroup was created as bespoke for the associated Account with
	// this Account ID.
	AccountID string `json:"accountId"`
	// The short code representing the PlanGroup.
	Code string `json:"code"`
	// The unique identifier (UUID) for the user who created the PlanGroup.
	CreatedBy string `json:"createdBy"`
	// Currency code for the PlanGroup (For example, USD).
	Currency string `json:"currency"`
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
	CustomFields map[string]PlanGroupCustomFieldsUnion `json:"customFields"`
	// The date and time _(in ISO 8601 format)_ when the PlanGroup was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the PlanGroup was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the PlanGroup.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The minimum spend amount for the PlanGroup.
	MinimumSpend float64 `json:"minimumSpend"`
	// Optional. Product ID to attribute the PlanGroup's minimum spend for accounting
	// purposes.
	MinimumSpendAccountingProductID string `json:"minimumSpendAccountingProductId"`
	// A boolean flag that determines when the minimum spend is billed. This flag
	// overrides the setting at Organizational level for minimum spend billing in
	// arrears/in advance.
	//
	// - **TRUE** - minimum spend is billed at the start of each billing period.
	// - **FALSE** - minimum spend is billed at the end of each billing period.
	MinimumSpendBillInAdvance bool `json:"minimumSpendBillInAdvance"`
	// Description of the minimum spend, displayed on the bill line item.
	MinimumSpendDescription string `json:"minimumSpendDescription"`
	// The name of the PlanGroup.
	Name string `json:"name"`
	// Standing charge amount for the PlanGroup.
	StandingCharge float64 `json:"standingCharge"`
	// Optional. Product ID to attribute the PlanGroup's standing charge for accounting
	// purposes.
	StandingChargeAccountingProductID string `json:"standingChargeAccountingProductId"`
	// A boolean flag that determines when the standing charge is billed. This flag
	// overrides the setting at Organizational level for standing charge billing in
	// arrears/in advance.
	//
	// - **TRUE** - standing charge is billed at the start of each billing period.
	// - **FALSE** - standing charge is billed at the end of each billing period.
	StandingChargeBillInAdvance bool `json:"standingChargeBillInAdvance"`
	// Description of the standing charge, displayed on the bill line item.
	StandingChargeDescription string        `json:"standingChargeDescription"`
	JSON                      planGroupJSON `json:"-"`
}

// planGroupJSON contains the JSON metadata for the struct [PlanGroup]
type planGroupJSON struct {
	ID                                apijson.Field
	Version                           apijson.Field
	AccountID                         apijson.Field
	Code                              apijson.Field
	CreatedBy                         apijson.Field
	Currency                          apijson.Field
	CustomFields                      apijson.Field
	DtCreated                         apijson.Field
	DtLastModified                    apijson.Field
	LastModifiedBy                    apijson.Field
	MinimumSpend                      apijson.Field
	MinimumSpendAccountingProductID   apijson.Field
	MinimumSpendBillInAdvance         apijson.Field
	MinimumSpendDescription           apijson.Field
	Name                              apijson.Field
	StandingCharge                    apijson.Field
	StandingChargeAccountingProductID apijson.Field
	StandingChargeBillInAdvance       apijson.Field
	StandingChargeDescription         apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *PlanGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type PlanGroupCustomFieldsUnion interface {
	ImplementsPlanGroupCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanGroupCustomFieldsUnion)(nil)).Elem(),
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

type PlanGroupNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Currency code for the PlanGroup (For example, USD).
	Currency param.Field[string] `json:"currency,required"`
	// The name of the PlanGroup.
	Name param.Field[string] `json:"name,required"`
	// Optional. This PlanGroup is created as bespoke for the associated Account with
	// this Account ID.
	AccountID param.Field[string] `json:"accountId"`
	// The short code representing the PlanGroup.
	Code param.Field[string] `json:"code"`
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
	CustomFields param.Field[map[string]PlanGroupNewParamsCustomFieldsUnion] `json:"customFields"`
	// The minimum spend amount for the PlanGroup.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// Optional. Product ID to attribute the PlanGroup's minimum spend for accounting
	// purposes.
	MinimumSpendAccountingProductID param.Field[string] `json:"minimumSpendAccountingProductId"`
	// A boolean flag that determines when the minimum spend is billed. This flag
	// overrides the setting at Organizational level for minimum spend billing in
	// arrears/in advance.
	//
	// - **TRUE** - minimum spend is billed at the start of each billing period.
	// - **FALSE** - minimum spend is billed at the end of each billing period.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Description of the minimum spend, displayed on the bill line item.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// Standing charge amount for the PlanGroup.
	StandingCharge param.Field[float64] `json:"standingCharge"`
	// Optional. Product ID to attribute the PlanGroup's standing charge for accounting
	// purposes.
	StandingChargeAccountingProductID param.Field[string] `json:"standingChargeAccountingProductId"`
	// A boolean flag that determines when the standing charge is billed. This flag
	// overrides the setting at Organizational level for standing charge billing in
	// arrears/in advance.
	//
	// - **TRUE** - standing charge is billed at the start of each billing period.
	// - **FALSE** - standing charge is billed at the end of each billing period.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Description of the standing charge, displayed on the bill line item.
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

func (r PlanGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PlanGroupNewParamsCustomFieldsUnion interface {
	ImplementsPlanGroupNewParamsCustomFieldsUnion()
}

type PlanGroupGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type PlanGroupUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Currency code for the PlanGroup (For example, USD).
	Currency param.Field[string] `json:"currency,required"`
	// The name of the PlanGroup.
	Name param.Field[string] `json:"name,required"`
	// Optional. This PlanGroup is created as bespoke for the associated Account with
	// this Account ID.
	AccountID param.Field[string] `json:"accountId"`
	// The short code representing the PlanGroup.
	Code param.Field[string] `json:"code"`
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
	CustomFields param.Field[map[string]PlanGroupUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// The minimum spend amount for the PlanGroup.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// Optional. Product ID to attribute the PlanGroup's minimum spend for accounting
	// purposes.
	MinimumSpendAccountingProductID param.Field[string] `json:"minimumSpendAccountingProductId"`
	// A boolean flag that determines when the minimum spend is billed. This flag
	// overrides the setting at Organizational level for minimum spend billing in
	// arrears/in advance.
	//
	// - **TRUE** - minimum spend is billed at the start of each billing period.
	// - **FALSE** - minimum spend is billed at the end of each billing period.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Description of the minimum spend, displayed on the bill line item.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// Standing charge amount for the PlanGroup.
	StandingCharge param.Field[float64] `json:"standingCharge"`
	// Optional. Product ID to attribute the PlanGroup's standing charge for accounting
	// purposes.
	StandingChargeAccountingProductID param.Field[string] `json:"standingChargeAccountingProductId"`
	// A boolean flag that determines when the standing charge is billed. This flag
	// overrides the setting at Organizational level for standing charge billing in
	// arrears/in advance.
	//
	// - **TRUE** - standing charge is billed at the start of each billing period.
	// - **FALSE** - standing charge is billed at the end of each billing period.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Description of the standing charge, displayed on the bill line item.
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

func (r PlanGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PlanGroupUpdateParamsCustomFieldsUnion interface {
	ImplementsPlanGroupUpdateParamsCustomFieldsUnion()
}

type PlanGroupListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// Optional filter. The list of Account IDs to which the PlanGroups belong.
	AccountID param.Field[[]string] `query:"accountId"`
	// Optional filter. The list of PlanGroup IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// PlanGroups in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of PlanGroups to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [PlanGroupListParams]'s query parameters as `url.Values`.
func (r PlanGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanGroupDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
