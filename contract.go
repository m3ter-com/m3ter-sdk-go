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

// ContractService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractService] method instead.
type ContractService struct {
	Options []option.RequestOption
}

// NewContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractService(opts ...option.RequestOption) (r *ContractService) {
	r = &ContractService{}
	r.Options = opts
	return
}

// Create a new Contract.
//
// Creates a new Contract for the specified Account. The Contract includes
// information such as the associated Account along with start and end dates.
func (r *ContractService) New(ctx context.Context, orgID string, body ContractNewParams, opts ...option.RequestOption) (res *Contract, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/contracts", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the Contract with the given UUID. Used to obtain the details of a
// Contract.
func (r *ContractService) Get(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Contract, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/contracts/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Contract with the given UUID.
//
// This endpoint updates the details of the Contract with the specified ID. Used to
// modify details of an existing Contract such as the start or end dates.
//
// **Note:** If you have created Custom Fields for a Contract, when you use this
// endpoint to update the Contract use the `customFields` parameter to preserve
// those Custom Fields. If you omit them from the update request, they will be
// lost.
func (r *ContractService) Update(ctx context.Context, orgID string, id string, body ContractUpdateParams, opts ...option.RequestOption) (res *Contract, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/contracts/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of Contracts by Organization ID. Supports pagination and
// includes various query parameters to filter the Contracts returned based on
// Contract IDs or short codes.
func (r *ContractService) List(ctx context.Context, orgID string, query ContractListParams, opts ...option.RequestOption) (res *pagination.Cursor[Contract], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/contracts", orgID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieves a list of Contracts by Organization ID. Supports pagination and
// includes various query parameters to filter the Contracts returned based on
// Contract IDs or short codes.
func (r *ContractService) ListAutoPaging(ctx context.Context, orgID string, query ContractListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Contract] {
	return pagination.NewCursorAutoPager(r.List(ctx, orgID, query, opts...))
}

// Deletes the Contract with the specified UUID. Used to remove an existing
// Contract from an Account.
//
// **Note:** This call will fail if there are any AccountPlans or Commitments that
// have been added to the Contract.
func (r *ContractService) Delete(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Contract, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/contracts/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Contract struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The unique identifier (UUID) of the Account associated with this Contract.
	AccountID string `json:"accountId"`
	// The short code of the Contract.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this Contract.
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
	CustomFields map[string]ContractCustomFieldsUnion `json:"customFields"`
	// The description of the Contract, which provides context and information.
	Description string `json:"description"`
	// The date and time _(in ISO-8601 format)_ when the Contract was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the Contract was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The exclusive end date of the Contract _(in ISO-8601 format)_. This means the
	// Contract is active until midnight on the day **_before_** this date.
	EndDate time.Time `json:"endDate" format:"date"`
	// The unique identifier (UUID) of the user who last modified this Contract.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Contract.
	Name string `json:"name"`
	// The Purchase Order Number associated with the Contract.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// The start date for the Contract _(in ISO-8601 format)_. This date is inclusive,
	// meaning the Contract is active from this date onward.
	StartDate time.Time    `json:"startDate" format:"date"`
	JSON      contractJSON `json:"-"`
}

// contractJSON contains the JSON metadata for the struct [Contract]
type contractJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	AccountID           apijson.Field
	Code                apijson.Field
	CreatedBy           apijson.Field
	CustomFields        apijson.Field
	Description         apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	EndDate             apijson.Field
	LastModifiedBy      apijson.Field
	Name                apijson.Field
	PurchaseOrderNumber apijson.Field
	StartDate           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Contract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type ContractCustomFieldsUnion interface {
	ImplementsContractCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ContractCustomFieldsUnion)(nil)).Elem(),
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

type ContractNewParams struct {
	// The unique identifier (UUID) of the Account associated with this Contract.
	AccountID param.Field[string] `json:"accountId,required"`
	// The exclusive end date of the Contract _(in ISO-8601 format)_. This means the
	// Contract is active until midnight on the day **_before_** this date.
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date"`
	// The name of the Contract.
	Name param.Field[string] `json:"name,required"`
	// The start date for the Contract _(in ISO-8601 format)_. This date is inclusive,
	// meaning the Contract is active from this date onward.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date"`
	// The short code of the Contract.
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
	CustomFields param.Field[map[string]ContractNewParamsCustomFieldsUnion] `json:"customFields"`
	// The description of the Contract, which provides context and information.
	Description param.Field[string] `json:"description"`
	// The Purchase Order Number associated with the Contract.
	PurchaseOrderNumber param.Field[string] `json:"purchaseOrderNumber"`
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

func (r ContractNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ContractNewParamsCustomFieldsUnion interface {
	ImplementsContractNewParamsCustomFieldsUnion()
}

type ContractUpdateParams struct {
	// The unique identifier (UUID) of the Account associated with this Contract.
	AccountID param.Field[string] `json:"accountId,required"`
	// The exclusive end date of the Contract _(in ISO-8601 format)_. This means the
	// Contract is active until midnight on the day **_before_** this date.
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date"`
	// The name of the Contract.
	Name param.Field[string] `json:"name,required"`
	// The start date for the Contract _(in ISO-8601 format)_. This date is inclusive,
	// meaning the Contract is active from this date onward.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date"`
	// The short code of the Contract.
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
	CustomFields param.Field[map[string]ContractUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// The description of the Contract, which provides context and information.
	Description param.Field[string] `json:"description"`
	// The Purchase Order Number associated with the Contract.
	PurchaseOrderNumber param.Field[string] `json:"purchaseOrderNumber"`
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

func (r ContractUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ContractUpdateParamsCustomFieldsUnion interface {
	ImplementsContractUpdateParamsCustomFieldsUnion()
}

type ContractListParams struct {
	AccountID param.Field[string] `query:"accountId"`
	// An optional parameter to retrieve specific Contracts based on their short codes.
	Codes param.Field[[]string] `query:"codes"`
	// An optional parameter to filter the list based on specific Contract unique
	// identifiers (UUIDs).
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Contracts in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Contracts to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ContractListParams]'s query parameters as `url.Values`.
func (r ContractListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
