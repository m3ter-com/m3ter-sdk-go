// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// CustomFieldService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomFieldService] method instead.
type CustomFieldService struct {
	Options []option.RequestOption
}

// NewCustomFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomFieldService(opts ...option.RequestOption) (r *CustomFieldService) {
	r = &CustomFieldService{}
	r.Options = opts
	return
}

// Retrieve all Custom Fields added at Organizational level for the entities that
// support them.
func (r *CustomFieldService) Get(ctx context.Context, query CustomFieldGetParams, opts ...option.RequestOption) (res *CustomFieldsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/customfields", query.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Custom Fields added at Organization level to entities that support them.
func (r *CustomFieldService) Update(ctx context.Context, params CustomFieldUpdateParams, opts ...option.RequestOption) (res *CustomFieldsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/customfields", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type CustomFieldsResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// CustomFields added to Account entities.
	Account map[string]CustomFieldsResponseAccountUnion `json:"account"`
	// CustomFields added to accountPlan entities.
	AccountPlan map[string]CustomFieldsResponseAccountPlanUnion `json:"accountPlan"`
	// CustomFields added to simple Aggregation entities.
	Aggregation map[string]CustomFieldsResponseAggregationUnion `json:"aggregation"`
	// CustomFields added to Compound Aggregation entities.
	CompoundAggregation map[string]CustomFieldsResponseCompoundAggregationUnion `json:"compoundAggregation"`
	// The id of the user who created this custom field.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the Organization was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when a custom field was last modified - created, modified, or
	// deleted - for the Organization _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this custom field.
	LastModifiedBy string `json:"lastModifiedBy"`
	// CustomFields added to Meter entities.
	Meter map[string]CustomFieldsResponseMeterUnion `json:"meter"`
	// CustomFields added to the Organization.
	Organization map[string]CustomFieldsResponseOrganizationUnion `json:"organization"`
	// CustomFields added to Plan entities.
	Plan map[string]CustomFieldsResponsePlanUnion `json:"plan"`
	// CustomFields added to planTemplate entities.
	PlanTemplate map[string]CustomFieldsResponsePlanTemplateUnion `json:"planTemplate"`
	// CustomFields added to Product entities.
	Product map[string]CustomFieldsResponseProductUnion `json:"product"`
	JSON    customFieldsResponseJSON                    `json:"-"`
}

// customFieldsResponseJSON contains the JSON metadata for the struct
// [CustomFieldsResponse]
type customFieldsResponseJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	Account             apijson.Field
	AccountPlan         apijson.Field
	Aggregation         apijson.Field
	CompoundAggregation apijson.Field
	CreatedBy           apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	LastModifiedBy      apijson.Field
	Meter               apijson.Field
	Organization        apijson.Field
	Plan                apijson.Field
	PlanTemplate        apijson.Field
	Product             apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomFieldsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customFieldsResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseAccountUnion interface {
	ImplementsCustomFieldsResponseAccountUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseAccountUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseAccountPlanUnion interface {
	ImplementsCustomFieldsResponseAccountPlanUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseAccountPlanUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseAggregationUnion interface {
	ImplementsCustomFieldsResponseAggregationUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseAggregationUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseCompoundAggregationUnion interface {
	ImplementsCustomFieldsResponseCompoundAggregationUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseCompoundAggregationUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseMeterUnion interface {
	ImplementsCustomFieldsResponseMeterUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseMeterUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseOrganizationUnion interface {
	ImplementsCustomFieldsResponseOrganizationUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseOrganizationUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponsePlanUnion interface {
	ImplementsCustomFieldsResponsePlanUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponsePlanUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponsePlanTemplateUnion interface {
	ImplementsCustomFieldsResponsePlanTemplateUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponsePlanTemplateUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type CustomFieldsResponseProductUnion interface {
	ImplementsCustomFieldsResponseProductUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomFieldsResponseProductUnion)(nil)).Elem(),
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

type CustomFieldGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type CustomFieldUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Updates to Account entity CustomFields.
	Account param.Field[map[string]CustomFieldUpdateParamsAccountUnion] `json:"account"`
	// Updates to accountPlan entity CustomFields.
	AccountPlan param.Field[map[string]CustomFieldUpdateParamsAccountPlanUnion] `json:"accountPlan"`
	// Updates to simple Aggregation entity CustomFields.
	Aggregation param.Field[map[string]CustomFieldUpdateParamsAggregationUnion] `json:"aggregation"`
	// Updates to Compound Aggregation entity CustomFields.
	CompoundAggregation param.Field[map[string]CustomFieldUpdateParamsCompoundAggregationUnion] `json:"compoundAggregation"`
	// Updates to Meter entity CustomFields.
	Meter param.Field[map[string]CustomFieldUpdateParamsMeterUnion] `json:"meter"`
	// Updates to Organization CustomFields.
	Organization param.Field[map[string]CustomFieldUpdateParamsOrganizationUnion] `json:"organization"`
	// Updates to Plan entity CustomFields.
	Plan param.Field[map[string]CustomFieldUpdateParamsPlanUnion] `json:"plan"`
	// Updates to planTemplate entity CustomFields.
	PlanTemplate param.Field[map[string]CustomFieldUpdateParamsPlanTemplateUnion] `json:"planTemplate"`
	// Updates to Product entity CustomFields.
	Product param.Field[map[string]CustomFieldUpdateParamsProductUnion] `json:"product"`
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

func (r CustomFieldUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsAccountUnion interface {
	ImplementsCustomFieldUpdateParamsAccountUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsAccountPlanUnion interface {
	ImplementsCustomFieldUpdateParamsAccountPlanUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsAggregationUnion interface {
	ImplementsCustomFieldUpdateParamsAggregationUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsCompoundAggregationUnion interface {
	ImplementsCustomFieldUpdateParamsCompoundAggregationUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsMeterUnion interface {
	ImplementsCustomFieldUpdateParamsMeterUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsOrganizationUnion interface {
	ImplementsCustomFieldUpdateParamsOrganizationUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsPlanUnion interface {
	ImplementsCustomFieldUpdateParamsPlanUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsPlanTemplateUnion interface {
	ImplementsCustomFieldUpdateParamsPlanTemplateUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type CustomFieldUpdateParamsProductUnion interface {
	ImplementsCustomFieldUpdateParamsProductUnion()
}
