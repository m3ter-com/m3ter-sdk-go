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
)

// TransactionTypeService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionTypeService] method instead.
type TransactionTypeService struct {
	Options []option.RequestOption
}

// NewTransactionTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionTypeService(opts ...option.RequestOption) (r *TransactionTypeService) {
	r = &TransactionTypeService{}
	r.Options = opts
	return
}

// Create a new TransactionType for the specified Organization. Details of the new
// TransactionType should be included in the request body.
func (r *TransactionTypeService) New(ctx context.Context, orgID string, body TransactionTypeNewParams, opts ...option.RequestOption) (res *TransactionType, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/transactiontypes", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the TransactionType with the given UUID from the specified
// Organization.
func (r *TransactionTypeService) Get(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *TransactionType, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/transactiontypes/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the TransactionType with the specified UUID for the specified
// Organization. Update details for the TransactionType should be included in the
// request body.
func (r *TransactionTypeService) Update(ctx context.Context, orgID string, id string, body TransactionTypeUpdateParams, opts ...option.RequestOption) (res *TransactionType, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/transactiontypes/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of TransactionType entities for the specified Organization. The
// list can be paginated for easier management, and supports filtering by various
// parameters.
func (r *TransactionTypeService) List(ctx context.Context, orgID string, query TransactionTypeListParams, opts ...option.RequestOption) (res *TransactionTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/transactiontypes", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the TransactionType with the given UUID from the specified Organization.
func (r *TransactionTypeService) Delete(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *TransactionType, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/transactiontypes/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type TransactionType struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// TRUE / FALSE flag indicating whether the data entity is archived. An entity can
	// be archived if it is obsolete.
	Archived bool `json:"archived"`
	// The short code of the data entity.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this TransactionType.
	CreatedBy string `json:"createdBy"`
	// The date and time _(in ISO-8601 format)_ when the TransactionType was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the TransactionType was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this TransactionType.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data entity.
	Name string              `json:"name"`
	JSON transactionTypeJSON `json:"-"`
}

// transactionTypeJSON contains the JSON metadata for the struct [TransactionType]
type transactionTypeJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	Archived       apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TransactionType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionTypeJSON) RawJSON() string {
	return r.raw
}

type TransactionTypeListResponse = interface{}

type TransactionTypeNewParams struct {
	// The name of the entity.
	Name param.Field[string] `json:"name,required"`
	// A Boolean TRUE / FALSE flag indicating whether the entity is archived. An entity
	// can be archived if it is obsolete.
	//
	// - TRUE - the entity is in the archived state.
	// - FALSE - the entity is not in the archived state.
	Archived param.Field[bool] `json:"archived"`
	// The short code for the entity.
	Code param.Field[string] `json:"code"`
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

func (r TransactionTypeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionTypeUpdateParams struct {
	// The name of the entity.
	Name param.Field[string] `json:"name,required"`
	// A Boolean TRUE / FALSE flag indicating whether the entity is archived. An entity
	// can be archived if it is obsolete.
	//
	// - TRUE - the entity is in the archived state.
	// - FALSE - the entity is not in the archived state.
	Archived param.Field[bool] `json:"archived"`
	// The short code for the entity.
	Code param.Field[string] `json:"code"`
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

func (r TransactionTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionTypeListParams struct {
	// Filter with this Boolean flag whether to include TransactionTypes that are
	// archived.
	//
	// - TRUE - include archived TransactionTypes in the list.
	// - FALSE - exclude archived TransactionTypes.
	Archived param.Field[bool] `query:"archived"`
	// A list of TransactionType short codes to retrieve.
	Codes param.Field[[]string] `query:"codes"`
	// A list of TransactionType unique identifiers (UUIDs) to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// TransactionTypes in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of TransactionTypes to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [TransactionTypeListParams]'s query parameters as
// `url.Values`.
func (r TransactionTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
