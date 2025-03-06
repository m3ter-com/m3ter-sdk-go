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

// CreditReasonService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditReasonService] method instead.
type CreditReasonService struct {
	Options []option.RequestOption
}

// NewCreditReasonService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCreditReasonService(opts ...option.RequestOption) (r *CreditReasonService) {
	r = &CreditReasonService{}
	r.Options = opts
	return
}

// Create a new Credit Reason for your Organization. When you've created a Credit
// Reason, it becomes available as a credit type for adding Credit line items to
// Bills. See [Credits](https://www.m3ter.com/docs/api#tag/Credits).
func (r *CreditReasonService) New(ctx context.Context, params CreditReasonNewParams, opts ...option.RequestOption) (res *CreditReasonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/creditreasons", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Credit Reason with the given UUID.
func (r *CreditReasonService) Get(ctx context.Context, id string, query CreditReasonGetParams, opts ...option.RequestOption) (res *CreditReasonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/creditreasons/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Credit Reason with the given UUID.
func (r *CreditReasonService) Update(ctx context.Context, id string, params CreditReasonUpdateParams, opts ...option.RequestOption) (res *CreditReasonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/creditreasons/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of the Credit Reason entities created for your Organization. You
// can filter the list returned for the call by Credit Reason ID, Credit Reason
// short code, or by Archive status.
func (r *CreditReasonService) List(ctx context.Context, params CreditReasonListParams, opts ...option.RequestOption) (res *pagination.Cursor[CreditReasonResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/creditreasons", params.OrgID)
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

// Retrieve a list of the Credit Reason entities created for your Organization. You
// can filter the list returned for the call by Credit Reason ID, Credit Reason
// short code, or by Archive status.
func (r *CreditReasonService) ListAutoPaging(ctx context.Context, params CreditReasonListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CreditReasonResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Credit Reason with the given UUID.
func (r *CreditReasonService) Delete(ctx context.Context, id string, body CreditReasonDeleteParams, opts ...option.RequestOption) (res *CreditReasonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/picklists/creditreasons/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CreditReasonResponse struct {
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
	// The id of the user who created this credit reason.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the credit reason was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the credit reason was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this credit reason.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data entity.
	Name string                   `json:"name"`
	JSON creditReasonResponseJSON `json:"-"`
}

// creditReasonResponseJSON contains the JSON metadata for the struct
// [CreditReasonResponse]
type creditReasonResponseJSON struct {
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

func (r *CreditReasonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditReasonResponseJSON) RawJSON() string {
	return r.raw
}

type CreditReasonNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
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

func (r CreditReasonNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditReasonGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type CreditReasonUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
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

func (r CreditReasonUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CreditReasonListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// TRUE / FALSE archived flag to filter the list. CreditReasons can be archived
	// once they are obsolete.
	//
	// - TRUE includes archived CreditReasons.
	// - FALSE excludes CreditReasons that are archived.
	Archived param.Field[bool] `query:"archived"`
	// List of Credit Reason short codes to retrieve.
	Codes param.Field[[]string] `query:"codes"`
	// List of Credit Reason IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of credit reasons to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [CreditReasonListParams]'s query parameters as `url.Values`.
func (r CreditReasonListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CreditReasonDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
