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
	"github.com/m3ter-com/m3ter-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// ProductService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProductService] method instead.
type ProductService struct {
	Options []option.RequestOption
}

// NewProductService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProductService(opts ...option.RequestOption) (r *ProductService) {
	r = &ProductService{}
	r.Options = opts
	return
}

// Create a new Product.
//
// This endpoint creates a new Product within the specified Organization. The
// details of the Product are provided in the request body.
func (r *ProductService) New(ctx context.Context, orgID string, body ProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/products", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Product with the given UUID.
//
// This endpoint retrieves the details of a specific Product within a specified
// Organization, using the Product UUID.
func (r *ProductService) Get(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/products/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Product with the given UUID.
//
// This endpoint updates the details of a specific Product within a specified
// Organization, using the Product UUID. The updated details are provided in the
// request body.
//
// **Note:** If you have created Custom Fields for a Product, when you use this
// endpoint to update the Product use the `customFields` parameter to preserve
// those Custom Fields. If you omit them from the update request, they will be
// lost.
func (r *ProductService) Update(ctx context.Context, orgID string, id string, body ProductUpdateParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/products/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of Products.
//
// This endpoint retrieves a list of all the Products within a specified
// Organization. The list can be paginated, and supports filtering by specific
// Product IDs.
func (r *ProductService) List(ctx context.Context, orgID string, query ProductListParams, opts ...option.RequestOption) (res *ProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/products", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a Product with the given UUID.
//
// This endpoint deletes a specific Product within a specified Organization, using
// the Product UUID.
func (r *ProductService) Delete(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Product, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/products/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Product struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// A unique short code to identify the Product. It should not contain control
	// chracters or spaces.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this Product.
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
	CustomFields map[string]ProductCustomFieldsUnion `json:"customFields"`
	// The date and time _(in ISO-8601 format)_ when the Product was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the Product was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this Product.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Descriptive name for the Product providing context and information.
	Name string      `json:"name"`
	JSON productJSON `json:"-"`
}

// productJSON contains the JSON metadata for the struct [Product]
type productJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	CustomFields   apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Product) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r productJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type ProductCustomFieldsUnion interface {
	ImplementsProductCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProductCustomFieldsUnion)(nil)).Elem(),
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

type ProductListResponse = interface{}

type ProductNewParams struct {
	// A unique short code to identify the Product. It should not contain control
	// chracters or spaces.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name for the Product providing context and information.
	Name param.Field[string] `json:"name,required"`
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
	CustomFields param.Field[map[string]ProductNewParamsCustomFieldsUnion] `json:"customFields"`
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

func (r ProductNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ProductNewParamsCustomFieldsUnion interface {
	ImplementsProductNewParamsCustomFieldsUnion()
}

type ProductUpdateParams struct {
	// A unique short code to identify the Product. It should not contain control
	// chracters or spaces.
	Code param.Field[string] `json:"code,required"`
	// Descriptive name for the Product providing context and information.
	Name param.Field[string] `json:"name,required"`
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
	CustomFields param.Field[map[string]ProductUpdateParamsCustomFieldsUnion] `json:"customFields"`
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

func (r ProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ProductUpdateParamsCustomFieldsUnion interface {
	ImplementsProductUpdateParamsCustomFieldsUnion()
}

type ProductListParams struct {
	// List of specific Product UUIDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Products in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Products to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [ProductListParams]'s query parameters as `url.Values`.
func (r ProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
