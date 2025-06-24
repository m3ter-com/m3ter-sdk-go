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

// CurrencyService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCurrencyService] method instead.
type CurrencyService struct {
	Options []option.RequestOption
}

// NewCurrencyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCurrencyService(opts ...option.RequestOption) (r *CurrencyService) {
	r = &CurrencyService{}
	r.Options = opts
	return
}

// Creates a new Currency for the specified Organization.
//
// Used to create a Currency that your Organization will start to use.
func (r *CurrencyService) New(ctx context.Context, params CurrencyNewParams, opts ...option.RequestOption) (res *CurrencyResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/currency", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the specified Currency with the given UUID. Used to obtain the details
// of a specified existing Currency in your Organization.
func (r *CurrencyService) Get(ctx context.Context, id string, query CurrencyGetParams, opts ...option.RequestOption) (res *CurrencyResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/currency/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Currency with the given UUID.
//
// Used to update the attributes of the specified Currency for the specified
// Organization.
func (r *CurrencyService) Update(ctx context.Context, id string, params CurrencyUpdateParams, opts ...option.RequestOption) (res *CurrencyResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/currency/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Currencies.
//
// Retrieves a list of Currencies for the specified Organization. This endpoint
// supports pagination and includes various query parameters to filter the
// Currencies based on Currency ID, and short codes.
func (r *CurrencyService) List(ctx context.Context, params CurrencyListParams, opts ...option.RequestOption) (res *pagination.Cursor[CurrencyResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/currency", params.OrgID)
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

// Retrieve a list of Currencies.
//
// Retrieves a list of Currencies for the specified Organization. This endpoint
// supports pagination and includes various query parameters to filter the
// Currencies based on Currency ID, and short codes.
func (r *CurrencyService) ListAutoPaging(ctx context.Context, params CurrencyListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CurrencyResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Currency with the given UUID.
//
// Used to remove an existing Currency from your Organization that is no longer
// required.
func (r *CurrencyService) Delete(ctx context.Context, id string, body CurrencyDeleteParams, opts ...option.RequestOption) (res *CurrencyResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/picklists/currency/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CurrencyResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// TRUE / FALSE flag indicating whether the data entity is archived. An entity can
	// be archived if it is obsolete.
	Archived bool `json:"archived"`
	// The short code of the data entity.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this Currency.
	CreatedBy string `json:"createdBy"`
	// The date and time _(in ISO-8601 format)_ when the Currency was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO-8601 format)_ when the Currency was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) of the user who last modified this Currency.
	LastModifiedBy string `json:"lastModifiedBy"`
	// This indicates the maximum number of decimal places to use for this Currency.
	MaxDecimalPlaces int64 `json:"maxDecimalPlaces"`
	// The name of the data entity.
	Name         string                       `json:"name"`
	RoundingMode CurrencyResponseRoundingMode `json:"roundingMode"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                `json:"version"`
	JSON    currencyResponseJSON `json:"-"`
}

// currencyResponseJSON contains the JSON metadata for the struct
// [CurrencyResponse]
type currencyResponseJSON struct {
	ID               apijson.Field
	Archived         apijson.Field
	Code             apijson.Field
	CreatedBy        apijson.Field
	DtCreated        apijson.Field
	DtLastModified   apijson.Field
	LastModifiedBy   apijson.Field
	MaxDecimalPlaces apijson.Field
	Name             apijson.Field
	RoundingMode     apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CurrencyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyResponseJSON) RawJSON() string {
	return r.raw
}

type CurrencyResponseRoundingMode string

const (
	CurrencyResponseRoundingModeUp          CurrencyResponseRoundingMode = "UP"
	CurrencyResponseRoundingModeDown        CurrencyResponseRoundingMode = "DOWN"
	CurrencyResponseRoundingModeCeiling     CurrencyResponseRoundingMode = "CEILING"
	CurrencyResponseRoundingModeFloor       CurrencyResponseRoundingMode = "FLOOR"
	CurrencyResponseRoundingModeHalfUp      CurrencyResponseRoundingMode = "HALF_UP"
	CurrencyResponseRoundingModeHalfDown    CurrencyResponseRoundingMode = "HALF_DOWN"
	CurrencyResponseRoundingModeHalfEven    CurrencyResponseRoundingMode = "HALF_EVEN"
	CurrencyResponseRoundingModeUnnecessary CurrencyResponseRoundingMode = "UNNECESSARY"
)

func (r CurrencyResponseRoundingMode) IsKnown() bool {
	switch r {
	case CurrencyResponseRoundingModeUp, CurrencyResponseRoundingModeDown, CurrencyResponseRoundingModeCeiling, CurrencyResponseRoundingModeFloor, CurrencyResponseRoundingModeHalfUp, CurrencyResponseRoundingModeHalfDown, CurrencyResponseRoundingModeHalfEven, CurrencyResponseRoundingModeUnnecessary:
		return true
	}
	return false
}

type CurrencyNewParams struct {
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
	// Indicates the maximum number of decimal places to use for this Currency.
	MaxDecimalPlaces param.Field[int64]                         `json:"maxDecimalPlaces"`
	RoundingMode     param.Field[CurrencyNewParamsRoundingMode] `json:"roundingMode"`
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

func (r CurrencyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CurrencyNewParamsRoundingMode string

const (
	CurrencyNewParamsRoundingModeUp          CurrencyNewParamsRoundingMode = "UP"
	CurrencyNewParamsRoundingModeDown        CurrencyNewParamsRoundingMode = "DOWN"
	CurrencyNewParamsRoundingModeCeiling     CurrencyNewParamsRoundingMode = "CEILING"
	CurrencyNewParamsRoundingModeFloor       CurrencyNewParamsRoundingMode = "FLOOR"
	CurrencyNewParamsRoundingModeHalfUp      CurrencyNewParamsRoundingMode = "HALF_UP"
	CurrencyNewParamsRoundingModeHalfDown    CurrencyNewParamsRoundingMode = "HALF_DOWN"
	CurrencyNewParamsRoundingModeHalfEven    CurrencyNewParamsRoundingMode = "HALF_EVEN"
	CurrencyNewParamsRoundingModeUnnecessary CurrencyNewParamsRoundingMode = "UNNECESSARY"
)

func (r CurrencyNewParamsRoundingMode) IsKnown() bool {
	switch r {
	case CurrencyNewParamsRoundingModeUp, CurrencyNewParamsRoundingModeDown, CurrencyNewParamsRoundingModeCeiling, CurrencyNewParamsRoundingModeFloor, CurrencyNewParamsRoundingModeHalfUp, CurrencyNewParamsRoundingModeHalfDown, CurrencyNewParamsRoundingModeHalfEven, CurrencyNewParamsRoundingModeUnnecessary:
		return true
	}
	return false
}

type CurrencyGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type CurrencyUpdateParams struct {
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
	// Indicates the maximum number of decimal places to use for this Currency.
	MaxDecimalPlaces param.Field[int64]                            `json:"maxDecimalPlaces"`
	RoundingMode     param.Field[CurrencyUpdateParamsRoundingMode] `json:"roundingMode"`
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

func (r CurrencyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CurrencyUpdateParamsRoundingMode string

const (
	CurrencyUpdateParamsRoundingModeUp          CurrencyUpdateParamsRoundingMode = "UP"
	CurrencyUpdateParamsRoundingModeDown        CurrencyUpdateParamsRoundingMode = "DOWN"
	CurrencyUpdateParamsRoundingModeCeiling     CurrencyUpdateParamsRoundingMode = "CEILING"
	CurrencyUpdateParamsRoundingModeFloor       CurrencyUpdateParamsRoundingMode = "FLOOR"
	CurrencyUpdateParamsRoundingModeHalfUp      CurrencyUpdateParamsRoundingMode = "HALF_UP"
	CurrencyUpdateParamsRoundingModeHalfDown    CurrencyUpdateParamsRoundingMode = "HALF_DOWN"
	CurrencyUpdateParamsRoundingModeHalfEven    CurrencyUpdateParamsRoundingMode = "HALF_EVEN"
	CurrencyUpdateParamsRoundingModeUnnecessary CurrencyUpdateParamsRoundingMode = "UNNECESSARY"
)

func (r CurrencyUpdateParamsRoundingMode) IsKnown() bool {
	switch r {
	case CurrencyUpdateParamsRoundingModeUp, CurrencyUpdateParamsRoundingModeDown, CurrencyUpdateParamsRoundingModeCeiling, CurrencyUpdateParamsRoundingModeFloor, CurrencyUpdateParamsRoundingModeHalfUp, CurrencyUpdateParamsRoundingModeHalfDown, CurrencyUpdateParamsRoundingModeHalfEven, CurrencyUpdateParamsRoundingModeUnnecessary:
		return true
	}
	return false
}

type CurrencyListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Filter by archived flag. A True / False flag indicating whether to return
	// Currencies that are archived _(obsolete)_.
	//
	// - TRUE - return archived Currencies.
	// - FALSE - archived Currencies are not returned.
	Archived param.Field[bool] `query:"archived"`
	// An optional parameter to retrieve specific Currencies based on their short
	// codes.
	Codes param.Field[[]string] `query:"codes"`
	// An optional parameter to filter the list based on specific Currency unique
	// identifiers (UUIDs).
	IDs param.Field[[]string] `query:"ids"`
	// The `nextToken` for multi-page retrievals. It is used to fetch the next page of
	// Currencies in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// Specifies the maximum number of Currencies to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [CurrencyListParams]'s query parameters as `url.Values`.
func (r CurrencyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CurrencyDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
