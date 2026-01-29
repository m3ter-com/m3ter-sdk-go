// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
)

// BalanceTransactionScheduleService contains methods and other services that help
// with interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceTransactionScheduleService] method instead.
type BalanceTransactionScheduleService struct {
	Options []option.RequestOption
}

// NewBalanceTransactionScheduleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBalanceTransactionScheduleService(opts ...option.RequestOption) (r *BalanceTransactionScheduleService) {
	r = &BalanceTransactionScheduleService{}
	r.Options = opts
	return
}

// Create a new BalanceTransactionSchedule.
func (r *BalanceTransactionScheduleService) New(ctx context.Context, balanceID string, params BalanceTransactionScheduleNewParams, opts ...option.RequestOption) (res *ScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/balancetransactionschedules", params.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a BalanceTransactionSchedule for the given UUID.
func (r *BalanceTransactionScheduleService) Get(ctx context.Context, balanceID string, id string, query BalanceTransactionScheduleGetParams, opts ...option.RequestOption) (res *ScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&query.OrgID, precfg.OrgID)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/balancetransactionschedules/%s", query.OrgID, balanceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a BalanceTransactionSchedule for the given UUID.
func (r *BalanceTransactionScheduleService) Update(ctx context.Context, balanceID string, id string, params BalanceTransactionScheduleUpdateParams, opts ...option.RequestOption) (res *ScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/balancetransactionschedules/%s", params.OrgID, balanceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of BalanceTransactionSchedule entities.
func (r *BalanceTransactionScheduleService) List(ctx context.Context, balanceID string, params BalanceTransactionScheduleListParams, opts ...option.RequestOption) (res *pagination.Cursor[ScheduleResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/balancetransactionschedules", params.OrgID, balanceID)
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

// Retrieve a list of BalanceTransactionSchedule entities.
func (r *BalanceTransactionScheduleService) ListAutoPaging(ctx context.Context, balanceID string, params BalanceTransactionScheduleListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ScheduleResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, balanceID, params, opts...))
}

// Delete the BalanceTransactionSchedule for the given UUID.
func (r *BalanceTransactionScheduleService) Delete(ctx context.Context, balanceID string, id string, body BalanceTransactionScheduleDeleteParams, opts ...option.RequestOption) (res *ScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&body.OrgID, precfg.OrgID)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/balancetransactionschedules/%s", body.OrgID, balanceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Previews the BalanceTransactions this Schedule would create, without persisting
// them. You can use this call to obtain a preview of the Transactions a Schedule
// you plan to create for a Balance would generate.
func (r *BalanceTransactionScheduleService) Preview(ctx context.Context, balanceID string, params BalanceTransactionSchedulePreviewParams, opts ...option.RequestOption) (res *ScheduleResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	requestconfig.UseDefaultParam(&params.OrgID, precfg.OrgID)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/balancetransactionschedules/preview", params.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BalanceTransactionScheduleNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID           param.Field[string]  `path:"orgId,required"`
	ScheduleRequest ScheduleRequestParam `json:"schedule_request,required"`
}

func (r BalanceTransactionScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ScheduleRequest)
}

type BalanceTransactionScheduleGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type BalanceTransactionScheduleUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID           param.Field[string]  `path:"orgId,required"`
	ScheduleRequest ScheduleRequestParam `json:"schedule_request,required"`
}

func (r BalanceTransactionScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ScheduleRequest)
}

type BalanceTransactionScheduleListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string]   `path:"orgId,required"`
	IDs   param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of BalanceTransactionSchedules to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [BalanceTransactionScheduleListParams]'s query parameters as
// `url.Values`.
func (r BalanceTransactionScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BalanceTransactionScheduleDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type BalanceTransactionSchedulePreviewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID           param.Field[string]  `path:"orgId,required"`
	ScheduleRequest ScheduleRequestParam `json:"schedule_request,required"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of BalanceTransactions to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

func (r BalanceTransactionSchedulePreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ScheduleRequest)
}

// URLQuery serializes [BalanceTransactionSchedulePreviewParams]'s query parameters
// as `url.Values`.
func (r BalanceTransactionSchedulePreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
