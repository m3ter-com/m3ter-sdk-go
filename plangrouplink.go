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

// PlanGroupLinkService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPlanGroupLinkService] method instead.
type PlanGroupLinkService struct {
	Options []option.RequestOption
}

// NewPlanGroupLinkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlanGroupLinkService(opts ...option.RequestOption) (r *PlanGroupLinkService) {
	r = &PlanGroupLinkService{}
	r.Options = opts
	return
}

// Create a new PlanGroupLink.
func (r *PlanGroupLinkService) New(ctx context.Context, params PlanGroupLinkNewParams, opts ...option.RequestOption) (res *PlanGroupLinkResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plangrouplinks", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a PlanGroupLink for the given UUID.
func (r *PlanGroupLinkService) Get(ctx context.Context, id string, query PlanGroupLinkGetParams, opts ...option.RequestOption) (res *PlanGroupLinkResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plangrouplinks/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update PlanGroupLink for the given UUID.
func (r *PlanGroupLinkService) Update(ctx context.Context, id string, params PlanGroupLinkUpdateParams, opts ...option.RequestOption) (res *PlanGroupLinkResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plangrouplinks/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of PlanGroupLink entities
func (r *PlanGroupLinkService) List(ctx context.Context, params PlanGroupLinkListParams, opts ...option.RequestOption) (res *pagination.Cursor[PlanGroupLinkResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/plangrouplinks", params.OrgID)
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

// Retrieve a list of PlanGroupLink entities
func (r *PlanGroupLinkService) ListAutoPaging(ctx context.Context, params PlanGroupLinkListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[PlanGroupLinkResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a PlanGroupLink for the given UUID.
func (r *PlanGroupLinkService) Delete(ctx context.Context, id string, body PlanGroupLinkDeleteParams, opts ...option.RequestOption) (res *PlanGroupLinkResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/plangrouplinks/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PlanGroupLinkResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The id of the user who created this plan group link.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the plan group link was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the plan group link was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this plan group link.
	LastModifiedBy string `json:"lastModifiedBy"`
	// ID of the linked PlanGroup
	PlanGroupID string `json:"planGroupId"`
	// ID of the linked Plan
	PlanID string `json:"planId"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                     `json:"version"`
	JSON    planGroupLinkResponseJSON `json:"-"`
}

// planGroupLinkResponseJSON contains the JSON metadata for the struct
// [PlanGroupLinkResponse]
type planGroupLinkResponseJSON struct {
	ID             apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	PlanGroupID    apijson.Field
	PlanID         apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlanGroupLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r planGroupLinkResponseJSON) RawJSON() string {
	return r.raw
}

type PlanGroupLinkNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID       param.Field[string] `path:"orgId,required"`
	PlanGroupID param.Field[string] `json:"planGroupId,required"`
	PlanID      param.Field[string] `json:"planId,required"`
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

func (r PlanGroupLinkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanGroupLinkGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type PlanGroupLinkUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID       param.Field[string] `path:"orgId,required"`
	PlanGroupID param.Field[string] `json:"planGroupId,required"`
	PlanID      param.Field[string] `json:"planId,required"`
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

func (r PlanGroupLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlanGroupLinkListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// list of IDs to retrieve
	IDs param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of PlanGroupLinks to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// UUID of the Plan to retrieve PlanGroupLinks for
	Plan param.Field[string] `query:"plan"`
	// UUID of the PlanGroup to retrieve PlanGroupLinks for
	PlanGroup param.Field[string] `query:"planGroup"`
}

// URLQuery serializes [PlanGroupLinkListParams]'s query parameters as
// `url.Values`.
func (r PlanGroupLinkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PlanGroupLinkDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
