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

// DebitReasonService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDebitReasonService] method instead.
type DebitReasonService struct {
	Options []option.RequestOption
}

// NewDebitReasonService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDebitReasonService(opts ...option.RequestOption) (r *DebitReasonService) {
	r = &DebitReasonService{}
	r.Options = opts
	return
}

// Create a new Debit Reason for your Organization. When you've created a Debit
// Reason, it becomes available as a debit type for adding Debit line items to
// Bills. See [Debits](https://www.m3ter.com/docs/api#tag/Debits).
func (r *DebitReasonService) New(ctx context.Context, params DebitReasonNewParams, opts ...option.RequestOption) (res *DebitReasonResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/debitreasons", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Debit Reason with the given UUID.
func (r *DebitReasonService) Get(ctx context.Context, id string, query DebitReasonGetParams, opts ...option.RequestOption) (res *DebitReasonResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/debitreasons/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Debit Reason with the given UUID.
func (r *DebitReasonService) Update(ctx context.Context, id string, params DebitReasonUpdateParams, opts ...option.RequestOption) (res *DebitReasonResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/debitreasons/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of the Debit Reason entities created for your Organization. You
// can filter the list returned for the call by Debit Reason ID, Debit Reason short
// code, or by Archive status.
func (r *DebitReasonService) List(ctx context.Context, params DebitReasonListParams, opts ...option.RequestOption) (res *pagination.Cursor[DebitReasonResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/debitreasons", params.OrgID)
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

// Retrieve a list of the Debit Reason entities created for your Organization. You
// can filter the list returned for the call by Debit Reason ID, Debit Reason short
// code, or by Archive status.
func (r *DebitReasonService) ListAutoPaging(ctx context.Context, params DebitReasonListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[DebitReasonResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Debit Reason with the given UUID.
func (r *DebitReasonService) Delete(ctx context.Context, id string, body DebitReasonDeleteParams, opts ...option.RequestOption) (res *DebitReasonResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/debitreasons/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DebitReasonResponse struct {
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
	// The id of the user who created this debit reason.
	CreatedBy string `json:"createdBy"`
	// The DateTime when the debit reason was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the debit reason was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this debit reason.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the data entity.
	Name string                  `json:"name"`
	JSON debitReasonResponseJSON `json:"-"`
}

// debitReasonResponseJSON contains the JSON metadata for the struct
// [DebitReasonResponse]
type debitReasonResponseJSON struct {
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

func (r *DebitReasonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r debitReasonResponseJSON) RawJSON() string {
	return r.raw
}

type DebitReasonNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
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

func (r DebitReasonNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DebitReasonGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type DebitReasonUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
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

func (r DebitReasonUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DebitReasonListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Filter using the boolean archived flag. DebitReasons can be archived if they are
	// obsolete.
	//
	// - TRUE includes DebitReasons that have been archived.
	// - FALSE excludes archived DebitReasons.
	Archived param.Field[bool] `query:"archived"`
	// List of Debit Reason short codes to retrieve.
	Codes param.Field[[]string] `query:"codes"`
	// List of Debit Reason IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Debit Reasons to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [DebitReasonListParams]'s query parameters as `url.Values`.
func (r DebitReasonListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DebitReasonDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
