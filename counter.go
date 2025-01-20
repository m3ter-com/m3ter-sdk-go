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

// CounterService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCounterService] method instead.
type CounterService struct {
	Options []option.RequestOption
}

// NewCounterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCounterService(opts ...option.RequestOption) (r *CounterService) {
	r = &CounterService{}
	r.Options = opts
	return
}

// Create a new Counter.
func (r *CounterService) New(ctx context.Context, orgID string, body CounterNewParams, opts ...option.RequestOption) (res *Counter, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counters", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Counter for the given UUID.
func (r *CounterService) Get(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Counter, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counters/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Counter for the given UUID.
func (r *CounterService) Update(ctx context.Context, orgID string, id string, body CounterUpdateParams, opts ...option.RequestOption) (res *Counter, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counters/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of Counter entities that can be filtered by Product, Counter ID,
// or Codes.
func (r *CounterService) List(ctx context.Context, orgID string, query CounterListParams, opts ...option.RequestOption) (res *pagination.Cursor[Counter], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counters", orgID)
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

// Retrieve a list of Counter entities that can be filtered by Product, Counter ID,
// or Codes.
func (r *CounterService) ListAutoPaging(ctx context.Context, orgID string, query CounterListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Counter] {
	return pagination.NewCursorAutoPager(r.List(ctx, orgID, query, opts...))
}

type Counter struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Code of the Counter. A unique short code to identify the Counter.
	Code string `json:"code"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Name of the Counter.
	Name string `json:"name"`
	// UUID of the product the Counter belongs to. _(Optional)_ - if no `productId` is
	// returned, the Counter is Global. A Global Counter can be used to price Plans or
	// Plan Templates belonging to any Product.
	ProductID string `json:"productId"`
	// Label for units shown on Bill line items, and indicating to customers what they
	// are being charged for.
	Unit string      `json:"unit"`
	JSON counterJSON `json:"-"`
}

// counterJSON contains the JSON metadata for the struct [Counter]
type counterJSON struct {
	ID             apijson.Field
	Version        apijson.Field
	Code           apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Name           apijson.Field
	ProductID      apijson.Field
	Unit           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Counter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r counterJSON) RawJSON() string {
	return r.raw
}

type CounterNewParams struct {
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

func (r CounterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterUpdateParams struct {
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

func (r CounterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterListParams struct {
	// List of Counter codes to retrieve. These are unique short codes to identify each
	// Counter.
	Codes param.Field[[]string] `query:"codes"`
	// List of Counter IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// NextToken for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Counters to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// List of Products UUIDs to retrieve Counters for.
	ProductID param.Field[[]string] `query:"productId"`
}

// URLQuery serializes [CounterListParams]'s query parameters as `url.Values`.
func (r CounterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
