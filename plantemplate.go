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

// PlanTemplateService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanTemplateService] method instead.
type PlanTemplateService struct {
	Options []option.RequestOption
}

// NewPlanTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlanTemplateService(opts ...option.RequestOption) (r *PlanTemplateService) {
	r = &PlanTemplateService{}
	r.Options = opts
	return
}

// Create a new PlanTemplate.
//
// This endpoint creates a new PlanTemplate within a specific Organization,
// identified by its unique UUID. The request body should contain the necessary
// information for the new PlanTemplate.
func (r *PlanTemplateService) New(ctx context.Context, params PlanTemplateNewParams, opts ...option.RequestOption) (res *PlanTemplateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plantemplates", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a specific PlanTemplate.
//
// This endpoint allows you to retrieve a specific PlanTemplate within a specific
// Organization, both identified by their unique identifiers (UUIDs).
func (r *PlanTemplateService) Get(ctx context.Context, id string, query PlanTemplateGetParams, opts ...option.RequestOption) (res *PlanTemplateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plantemplates/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific PlanTemplate.
//
// This endpoint enables you to update a specific PlanTemplate within a specific
// Organization, both identified by their unique identifiers (UUIDs). The request
// body should contain the updated information for the PlanTemplate.
//
// **Note:** If you have created Custom Fields for a Plan Template, when you use
// this endpoint to update the Plan Template use the `customFields` parameter to
// preserve those Custom Fields. If you omit them from the update request, they
// will be lost.
func (r *PlanTemplateService) Update(ctx context.Context, id string, params PlanTemplateUpdateParams, opts ...option.RequestOption) (res *PlanTemplateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plantemplates/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of PlanTemplates.
//
// This endpoint enables you to retrieve a paginated list of PlanTemplates
// belonging to a specific Organization, identified by its UUID. You can filter the
// list by PlanTemplate IDs or Product IDs for more focused retrieval.
func (r *PlanTemplateService) List(ctx context.Context, params PlanTemplateListParams, opts ...option.RequestOption) (res *pagination.Cursor[PlanTemplateResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/plantemplates", params.OrgID)
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

// Retrieve a list of PlanTemplates.
//
// This endpoint enables you to retrieve a paginated list of PlanTemplates
// belonging to a specific Organization, identified by its UUID. You can filter the
// list by PlanTemplate IDs or Product IDs for more focused retrieval.
func (r *PlanTemplateService) ListAutoPaging(ctx context.Context, params PlanTemplateListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PlanTemplateResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a specific PlanTemplate.
//
// This endpoint enables you to delete a specific PlanTemplate within a specific
// Organization, both identified by their unique identifiers (UUIDs).
func (r *PlanTemplateService) Delete(ctx context.Context, id string, body PlanTemplateDeleteParams, opts ...option.RequestOption) (res *PlanTemplateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plantemplates/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PlanTemplateResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Determines the frequency at which bills are generated.
	//
	//   - **Daily**. Starting at midnight each day, covering the twenty-four hour period
	//     following.
	//
	//   - **Weekly**. Starting at midnight on a Monday, covering the seven-day period
	//     following.
	//
	//   - **Monthly**. Starting at midnight on the first day of each month, covering the
	//     entire calendar month following.
	//
	//   - **Annually**. Starting at midnight on first day of each year covering the
	//     entire calendar year following.
	BillFrequency PlanTemplateResponseBillFrequency `json:"billFrequency"`
	// How often bills are issued. For example, if `billFrequency` is Monthly and
	// `billFrequencyInterval` is 3, bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// A unique, short code reference for the PlanTemplate. This code should not
	// contain control characters or spaces.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this PlanTemplate.
	CreatedBy string `json:"createdBy"`
	// The ISO currency code for the pricing currency used by Plans based on the Plan
	// Template to define charge rates for Product consumption - for example USD, GBP,
	// EUR.
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
	CustomFields map[string]PlanTemplateResponseCustomFieldsUnion `json:"customFields"`
	// The date and time _(in ISO-8601 format)_ when the PlanTemplate was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the PlanTemplate was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this PlanTemplate.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The Product minimum spend amount per billing cycle for end customer Accounts on
	// a pricing Plan based on the PlanTemplate. This must be a non-negative number.
	MinimumSpend float64 `json:"minimumSpend"`
	// A boolean that determines when the minimum spend is billed.
	//
	// - TRUE - minimum spend is billed at the start of each billing period.
	// - FALSE - minimum spend is billed at the end of each billing period.
	//
	// Overrides the setting at Organizational level for minimum spend billing in
	// arrears/in advance.
	MinimumSpendBillInAdvance bool `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription string `json:"minimumSpendDescription"`
	// Descriptive name for the PlanTemplate.
	Name string `json:"name"`
	// The ranking of the PlanTemplate among your pricing plans. Lower numbers
	// represent more basic plans, while higher numbers represent premium plans. This
	// must be a non-negative integer.
	//
	// **NOTE:** **DEPRECATED** - no longer used.
	Ordinal int64 `json:"ordinal"`
	// The unique identifier (UUID) of the Product associated with this PlanTemplate.
	ProductID string `json:"productId"`
	// The fixed charge _(standing charge)_ applied to customer bills. This charge is
	// prorated and must be a non-negative number.
	StandingCharge float64 `json:"standingCharge"`
	// A boolean that determines when the standing charge is billed.
	//
	// - TRUE - standing charge is billed at the start of each billing period.
	// - FALSE - standing charge is billed at the end of each billing period.
	//
	// Overrides the setting at Organizational level for standing charge billing in
	// arrears/in advance.
	StandingChargeBillInAdvance bool `json:"standingChargeBillInAdvance"`
	// Standing charge description _(displayed on the bill line item)_.
	StandingChargeDescription string `json:"standingChargeDescription"`
	// How often the standing charge is applied. For example, if the bill is issued
	// every three months and `standingChargeInterval` is 2, then the standing charge
	// is applied every six months.
	StandingChargeInterval int64 `json:"standingChargeInterval"`
	// Defines an offset for when the standing charge is first applied. For example, if
	// the bill is issued every three months and the `standingChargeOfset` is 0, then
	// the charge is applied to the first bill _(at three months)_; if 1, it would be
	// applied to the second bill _(at six months)_, and so on.
	StandingChargeOffset int64                    `json:"standingChargeOffset"`
	JSON                 planTemplateResponseJSON `json:"-"`
}

// planTemplateResponseJSON contains the JSON metadata for the struct
// [PlanTemplateResponse]
type planTemplateResponseJSON struct {
	ID                          apijson.Field
	Version                     apijson.Field
	BillFrequency               apijson.Field
	BillFrequencyInterval       apijson.Field
	Code                        apijson.Field
	CreatedBy                   apijson.Field
	Currency                    apijson.Field
	CustomFields                apijson.Field
	DtCreated                   apijson.Field
	DtLastModified              apijson.Field
	LastModifiedBy              apijson.Field
	MinimumSpend                apijson.Field
	MinimumSpendBillInAdvance   apijson.Field
	MinimumSpendDescription     apijson.Field
	Name                        apijson.Field
	Ordinal                     apijson.Field
	ProductID                   apijson.Field
	StandingCharge              apijson.Field
	StandingChargeBillInAdvance apijson.Field
	StandingChargeDescription   apijson.Field
	StandingChargeInterval      apijson.Field
	StandingChargeOffset        apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *PlanTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planTemplateResponseJSON) RawJSON() string {
	return r.raw
}

// Determines the frequency at which bills are generated.
//
//   - **Daily**. Starting at midnight each day, covering the twenty-four hour period
//     following.
//
//   - **Weekly**. Starting at midnight on a Monday, covering the seven-day period
//     following.
//
//   - **Monthly**. Starting at midnight on the first day of each month, covering the
//     entire calendar month following.
//
//   - **Annually**. Starting at midnight on first day of each year covering the
//     entire calendar year following.
type PlanTemplateResponseBillFrequency string

const (
	PlanTemplateResponseBillFrequencyDaily    PlanTemplateResponseBillFrequency = "DAILY"
	PlanTemplateResponseBillFrequencyWeekly   PlanTemplateResponseBillFrequency = "WEEKLY"
	PlanTemplateResponseBillFrequencyMonthly  PlanTemplateResponseBillFrequency = "MONTHLY"
	PlanTemplateResponseBillFrequencyAnnually PlanTemplateResponseBillFrequency = "ANNUALLY"
	PlanTemplateResponseBillFrequencyAdHoc    PlanTemplateResponseBillFrequency = "AD_HOC"
	PlanTemplateResponseBillFrequencyMixed    PlanTemplateResponseBillFrequency = "MIXED"
)

func (r PlanTemplateResponseBillFrequency) IsKnown() bool {
	switch r {
	case PlanTemplateResponseBillFrequencyDaily, PlanTemplateResponseBillFrequencyWeekly, PlanTemplateResponseBillFrequencyMonthly, PlanTemplateResponseBillFrequencyAnnually, PlanTemplateResponseBillFrequencyAdHoc, PlanTemplateResponseBillFrequencyMixed:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type PlanTemplateResponseCustomFieldsUnion interface {
	ImplementsPlanTemplateResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PlanTemplateResponseCustomFieldsUnion)(nil)).Elem(),
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

type PlanTemplateNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Determines the frequency at which bills are generated.
	//
	//   - **Daily**. Starting at midnight each day, covering the twenty-four hour period
	//     following.
	//
	//   - **Weekly**. Starting at midnight on a Monday, covering the seven-day period
	//     following.
	//
	//   - **Monthly**. Starting at midnight on the first day of each month, covering the
	//     entire calendar month following.
	//
	//   - **Annually**. Starting at midnight on first day of each year covering the
	//     entire calendar year following.
	BillFrequency param.Field[PlanTemplateNewParamsBillFrequency] `json:"billFrequency,required"`
	// The ISO currency code for the currency used to charge end users - for example
	// USD, GBP, EUR. This defines the _pricing currency_ and is inherited by any Plans
	// based on the Plan Template.
	//
	// **Notes:**
	//
	//   - You can define a currency at Organization-level or Account-level to be used as
	//     the _billing currency_. This can be a different currency to that used for the
	//     Plan as the _pricing currency_.
	//   - If the billing currency for an Account is different to the pricing currency
	//     used by a Plan attached to the Account, you must ensure a _currency conversion
	//     rate_ is defined for your Organization to convert the pricing currency into
	//     the billing currency at billing, otherwise Bills will fail for the Account.
	//   - To define any required currency conversion rates, use the
	//     `currencyConversions` request body parameter for the
	//     [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/UpdateOrganizationConfig)
	//     call.
	Currency param.Field[string] `json:"currency,required"`
	// Descriptive name for the PlanTemplate.
	Name param.Field[string] `json:"name,required"`
	// The unique identifier (UUID) of the Product associated with this PlanTemplate.
	ProductID param.Field[string] `json:"productId,required"`
	// The fixed charge _(standing charge)_ applied to customer bills. This charge is
	// prorated and must be a non-negative number.
	StandingCharge param.Field[float64] `json:"standingCharge,required"`
	// How often bills are issued. For example, if `billFrequency` is Monthly and
	// `billFrequencyInterval` is 3, bills are issued every three months.
	BillFrequencyInterval param.Field[int64] `json:"billFrequencyInterval"`
	// A unique, short code reference for the PlanTemplate. This code should not
	// contain control characters or spaces.
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
	CustomFields param.Field[map[string]PlanTemplateNewParamsCustomFieldsUnion] `json:"customFields"`
	// The Product minimum spend amount per billing cycle for end customer Accounts on
	// a pricing Plan based on the PlanTemplate. This must be a non-negative number.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// A boolean that determines when the minimum spend is billed.
	//
	// - TRUE - minimum spend is billed at the start of each billing period.
	// - FALSE - minimum spend is billed at the end of each billing period.
	//
	// Overrides the setting at Organizational level for minimum spend billing in
	// arrears/in advance.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// The ranking of the PlanTemplate among your pricing plans. Lower numbers
	// represent more basic plans, while higher numbers represent premium plans. This
	// must be a non-negative integer.
	//
	// **NOTE: DEPRECATED** - do not use.
	Ordinal param.Field[int64] `json:"ordinal"`
	// A boolean that determines when the standing charge is billed.
	//
	// - TRUE - standing charge is billed at the start of each billing period.
	// - FALSE - standing charge is billed at the end of each billing period.
	//
	// Overrides the setting at Organizational level for standing charge billing in
	// arrears/in advance.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Standing charge description _(displayed on the bill line item)_.
	StandingChargeDescription param.Field[string] `json:"standingChargeDescription"`
	// How often the standing charge is applied. For example, if the bill is issued
	// every three months and `standingChargeInterval` is 2, then the standing charge
	// is applied every six months.
	StandingChargeInterval param.Field[int64] `json:"standingChargeInterval"`
	// Defines an offset for when the standing charge is first applied. For example, if
	// the bill is issued every three months and the `standingChargeOfset` is 0, then
	// the charge is applied to the first bill _(at three months)_; if 1, it would be
	// applied to the second bill _(at six months)_, and so on.
	StandingChargeOffset param.Field[int64] `json:"standingChargeOffset"`
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

func (r PlanTemplateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the frequency at which bills are generated.
//
//   - **Daily**. Starting at midnight each day, covering the twenty-four hour period
//     following.
//
//   - **Weekly**. Starting at midnight on a Monday, covering the seven-day period
//     following.
//
//   - **Monthly**. Starting at midnight on the first day of each month, covering the
//     entire calendar month following.
//
//   - **Annually**. Starting at midnight on first day of each year covering the
//     entire calendar year following.
type PlanTemplateNewParamsBillFrequency string

const (
	PlanTemplateNewParamsBillFrequencyDaily    PlanTemplateNewParamsBillFrequency = "DAILY"
	PlanTemplateNewParamsBillFrequencyWeekly   PlanTemplateNewParamsBillFrequency = "WEEKLY"
	PlanTemplateNewParamsBillFrequencyMonthly  PlanTemplateNewParamsBillFrequency = "MONTHLY"
	PlanTemplateNewParamsBillFrequencyAnnually PlanTemplateNewParamsBillFrequency = "ANNUALLY"
	PlanTemplateNewParamsBillFrequencyAdHoc    PlanTemplateNewParamsBillFrequency = "AD_HOC"
	PlanTemplateNewParamsBillFrequencyMixed    PlanTemplateNewParamsBillFrequency = "MIXED"
)

func (r PlanTemplateNewParamsBillFrequency) IsKnown() bool {
	switch r {
	case PlanTemplateNewParamsBillFrequencyDaily, PlanTemplateNewParamsBillFrequencyWeekly, PlanTemplateNewParamsBillFrequencyMonthly, PlanTemplateNewParamsBillFrequencyAnnually, PlanTemplateNewParamsBillFrequencyAdHoc, PlanTemplateNewParamsBillFrequencyMixed:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PlanTemplateNewParamsCustomFieldsUnion interface {
	ImplementsPlanTemplateNewParamsCustomFieldsUnion()
}

type PlanTemplateGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type PlanTemplateUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Determines the frequency at which bills are generated.
	//
	//   - **Daily**. Starting at midnight each day, covering the twenty-four hour period
	//     following.
	//
	//   - **Weekly**. Starting at midnight on a Monday, covering the seven-day period
	//     following.
	//
	//   - **Monthly**. Starting at midnight on the first day of each month, covering the
	//     entire calendar month following.
	//
	//   - **Annually**. Starting at midnight on first day of each year covering the
	//     entire calendar year following.
	BillFrequency param.Field[PlanTemplateUpdateParamsBillFrequency] `json:"billFrequency,required"`
	// The ISO currency code for the currency used to charge end users - for example
	// USD, GBP, EUR. This defines the _pricing currency_ and is inherited by any Plans
	// based on the Plan Template.
	//
	// **Notes:**
	//
	//   - You can define a currency at Organization-level or Account-level to be used as
	//     the _billing currency_. This can be a different currency to that used for the
	//     Plan as the _pricing currency_.
	//   - If the billing currency for an Account is different to the pricing currency
	//     used by a Plan attached to the Account, you must ensure a _currency conversion
	//     rate_ is defined for your Organization to convert the pricing currency into
	//     the billing currency at billing, otherwise Bills will fail for the Account.
	//   - To define any required currency conversion rates, use the
	//     `currencyConversions` request body parameter for the
	//     [Update OrganizationConfig](https://www.m3ter.com/docs/api#tag/OrganizationConfig/operation/UpdateOrganizationConfig)
	//     call.
	Currency param.Field[string] `json:"currency,required"`
	// Descriptive name for the PlanTemplate.
	Name param.Field[string] `json:"name,required"`
	// The unique identifier (UUID) of the Product associated with this PlanTemplate.
	ProductID param.Field[string] `json:"productId,required"`
	// The fixed charge _(standing charge)_ applied to customer bills. This charge is
	// prorated and must be a non-negative number.
	StandingCharge param.Field[float64] `json:"standingCharge,required"`
	// How often bills are issued. For example, if `billFrequency` is Monthly and
	// `billFrequencyInterval` is 3, bills are issued every three months.
	BillFrequencyInterval param.Field[int64] `json:"billFrequencyInterval"`
	// A unique, short code reference for the PlanTemplate. This code should not
	// contain control characters or spaces.
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
	CustomFields param.Field[map[string]PlanTemplateUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// The Product minimum spend amount per billing cycle for end customer Accounts on
	// a pricing Plan based on the PlanTemplate. This must be a non-negative number.
	MinimumSpend param.Field[float64] `json:"minimumSpend"`
	// A boolean that determines when the minimum spend is billed.
	//
	// - TRUE - minimum spend is billed at the start of each billing period.
	// - FALSE - minimum spend is billed at the end of each billing period.
	//
	// Overrides the setting at Organizational level for minimum spend billing in
	// arrears/in advance.
	MinimumSpendBillInAdvance param.Field[bool] `json:"minimumSpendBillInAdvance"`
	// Minimum spend description _(displayed on the bill line item)_.
	MinimumSpendDescription param.Field[string] `json:"minimumSpendDescription"`
	// The ranking of the PlanTemplate among your pricing plans. Lower numbers
	// represent more basic plans, while higher numbers represent premium plans. This
	// must be a non-negative integer.
	//
	// **NOTE: DEPRECATED** - do not use.
	Ordinal param.Field[int64] `json:"ordinal"`
	// A boolean that determines when the standing charge is billed.
	//
	// - TRUE - standing charge is billed at the start of each billing period.
	// - FALSE - standing charge is billed at the end of each billing period.
	//
	// Overrides the setting at Organizational level for standing charge billing in
	// arrears/in advance.
	StandingChargeBillInAdvance param.Field[bool] `json:"standingChargeBillInAdvance"`
	// Standing charge description _(displayed on the bill line item)_.
	StandingChargeDescription param.Field[string] `json:"standingChargeDescription"`
	// How often the standing charge is applied. For example, if the bill is issued
	// every three months and `standingChargeInterval` is 2, then the standing charge
	// is applied every six months.
	StandingChargeInterval param.Field[int64] `json:"standingChargeInterval"`
	// Defines an offset for when the standing charge is first applied. For example, if
	// the bill is issued every three months and the `standingChargeOfset` is 0, then
	// the charge is applied to the first bill _(at three months)_; if 1, it would be
	// applied to the second bill _(at six months)_, and so on.
	StandingChargeOffset param.Field[int64] `json:"standingChargeOffset"`
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

func (r PlanTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the frequency at which bills are generated.
//
//   - **Daily**. Starting at midnight each day, covering the twenty-four hour period
//     following.
//
//   - **Weekly**. Starting at midnight on a Monday, covering the seven-day period
//     following.
//
//   - **Monthly**. Starting at midnight on the first day of each month, covering the
//     entire calendar month following.
//
//   - **Annually**. Starting at midnight on first day of each year covering the
//     entire calendar year following.
type PlanTemplateUpdateParamsBillFrequency string

const (
	PlanTemplateUpdateParamsBillFrequencyDaily    PlanTemplateUpdateParamsBillFrequency = "DAILY"
	PlanTemplateUpdateParamsBillFrequencyWeekly   PlanTemplateUpdateParamsBillFrequency = "WEEKLY"
	PlanTemplateUpdateParamsBillFrequencyMonthly  PlanTemplateUpdateParamsBillFrequency = "MONTHLY"
	PlanTemplateUpdateParamsBillFrequencyAnnually PlanTemplateUpdateParamsBillFrequency = "ANNUALLY"
	PlanTemplateUpdateParamsBillFrequencyAdHoc    PlanTemplateUpdateParamsBillFrequency = "AD_HOC"
	PlanTemplateUpdateParamsBillFrequencyMixed    PlanTemplateUpdateParamsBillFrequency = "MIXED"
)

func (r PlanTemplateUpdateParamsBillFrequency) IsKnown() bool {
	switch r {
	case PlanTemplateUpdateParamsBillFrequencyDaily, PlanTemplateUpdateParamsBillFrequencyWeekly, PlanTemplateUpdateParamsBillFrequencyMonthly, PlanTemplateUpdateParamsBillFrequencyAnnually, PlanTemplateUpdateParamsBillFrequencyAdHoc, PlanTemplateUpdateParamsBillFrequencyMixed:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type PlanTemplateUpdateParamsCustomFieldsUnion interface {
	ImplementsPlanTemplateUpdateParamsCustomFieldsUnion()
}

type PlanTemplateListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// List of specific PlanTemplate UUIDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// PlanTemplates in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of PlanTemplates to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
	// The unique identifiers (UUIDs) of the Products to retrieve associated
	// PlanTemplates.
	ProductID param.Field[string] `query:"productId"`
}

// URLQuery serializes [PlanTemplateListParams]'s query parameters as `url.Values`.
func (r PlanTemplateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanTemplateDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
