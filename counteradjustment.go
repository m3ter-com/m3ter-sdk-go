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

// CounterAdjustmentService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCounterAdjustmentService] method instead.
type CounterAdjustmentService struct {
	Options []option.RequestOption
}

// NewCounterAdjustmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCounterAdjustmentService(opts ...option.RequestOption) (r *CounterAdjustmentService) {
	r = &CounterAdjustmentService{}
	r.Options = opts
	return
}

// Create a new CounterAdjustment for an Account using a Counter.
//
// **Notes:**
//
//   - Use the new absolute value for the Counter for the selected date - if it was
//     15 and has increased to 20, enter 20; if it was 15 and has decreased to 10,
//     enter 10. _Do not enter_ the plus or minus value relative to the previous
//     Counter value on the Account.
//   - CounterAdjustments on Accounts are supported down to a _specific day_ of
//     granularity - you cannot create more than one CounterAdjustment for any given
//     day using the same Counter and you'll receive an error if you try to do this.
func (r *CounterAdjustmentService) New(ctx context.Context, params CounterAdjustmentNewParams, opts ...option.RequestOption) (res *CounterAdjustment, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counteradjustments", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a CounterAdjustment for the given UUID.
func (r *CounterAdjustmentService) Get(ctx context.Context, id string, query CounterAdjustmentGetParams, opts ...option.RequestOption) (res *CounterAdjustment, err error) {
	opts = append(r.Options[:], opts...)
	if query.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counteradjustments/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a CounterAdjustment for an Account.
func (r *CounterAdjustmentService) Update(ctx context.Context, id string, params CounterAdjustmentUpdateParams, opts ...option.RequestOption) (res *CounterAdjustment, err error) {
	opts = append(r.Options[:], opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counteradjustments/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of CounterAdjustments created for Accounts in your Organization.
// You can filter the list returned by date, Account ID, or Counter ID.
//
// **CONSTRAINTS:**
//
//   - The `counterId` query parameter is always required when calling this endpoint,
//     used either as a single query parameter or in combination with any of the
//     other query parameters.
//   - If you want to use the `date`, `dateStart`, or `dateEnd` query parameters, you
//     must also use the `accountId` query parameter.
func (r *CounterAdjustmentService) List(ctx context.Context, params CounterAdjustmentListParams, opts ...option.RequestOption) (res *pagination.Cursor[CounterAdjustment], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counteradjustments", params.OrgID)
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

// Retrieve a list of CounterAdjustments created for Accounts in your Organization.
// You can filter the list returned by date, Account ID, or Counter ID.
//
// **CONSTRAINTS:**
//
//   - The `counterId` query parameter is always required when calling this endpoint,
//     used either as a single query parameter or in combination with any of the
//     other query parameters.
//   - If you want to use the `date`, `dateStart`, or `dateEnd` query parameters, you
//     must also use the `accountId` query parameter.
func (r *CounterAdjustmentService) ListAutoPaging(ctx context.Context, params CounterAdjustmentListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[CounterAdjustment] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a CounterAdjustment for the given UUID.
func (r *CounterAdjustmentService) Delete(ctx context.Context, id string, body CounterAdjustmentDeleteParams, opts ...option.RequestOption) (res *CounterAdjustment, err error) {
	opts = append(r.Options[:], opts...)
	if body.OrgID.Value == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/counteradjustments/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CounterAdjustment struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The Account ID the CounterAdjustment was created for.
	AccountID string `json:"accountId"`
	// The ID of the Counter that was used to make the CounterAdjustment on the
	// Account.
	CounterID string `json:"counterId"`
	// The ID of the user who created this item.
	CreatedBy string `json:"createdBy"`
	// The date the CounterAdjustment was created for the Account _(in ISO-8601 date
	// format)_.
	Date time.Time `json:"date" format:"date"`
	// The DateTime when this item was created _(in ISO-8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when this item was last modified _(in ISO-8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the user who last modified this item.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Purchase Order Number for the Counter Adjustment. _(Optional)_
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// Integer Value of the Counter that was used to make the CounterAdjustment.
	Value int64                 `json:"value"`
	JSON  counterAdjustmentJSON `json:"-"`
}

// counterAdjustmentJSON contains the JSON metadata for the struct
// [CounterAdjustment]
type counterAdjustmentJSON struct {
	ID                  apijson.Field
	Version             apijson.Field
	AccountID           apijson.Field
	CounterID           apijson.Field
	CreatedBy           apijson.Field
	Date                apijson.Field
	DtCreated           apijson.Field
	DtLastModified      apijson.Field
	LastModifiedBy      apijson.Field
	PurchaseOrderNumber apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CounterAdjustment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r counterAdjustmentJSON) RawJSON() string {
	return r.raw
}

type CounterAdjustmentNewParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The Account ID the CounterAdjustment is created for.
	AccountID param.Field[string] `json:"accountId,required"`
	// The ID of the Counter used for the CounterAdjustment on the Account.
	CounterID param.Field[string] `json:"counterId,required"`
	// The date the CounterAdjustment is created for the Account _(in ISO-8601 date
	// format)_.
	//
	// **Note:** CounterAdjustments on Accounts are supported down to a _specific day_
	// of granularity - you cannot create more than one CounterAdjustment for any given
	// day using the same Counter and you'll receive an error if you try to do this.
	Date param.Field[string] `json:"date,required"`
	// Integer Value of the Counter used for the CounterAdjustment.
	//
	// **Note:** Use the new absolute value for the Counter for the selected date - if
	// it was 15 and has increased to 20, enter 20; if it was 15 and has decreased to
	// 10, enter 10. _Do not enter_ the plus or minus value relative to the previous
	// Counter value on the Account.
	Value param.Field[int64] `json:"value,required"`
	// Purchase Order Number for the Counter Adjustment. _(Optional)_
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

func (r CounterAdjustmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterAdjustmentGetParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}

type CounterAdjustmentUpdateParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// The Account ID the CounterAdjustment is created for.
	AccountID param.Field[string] `json:"accountId,required"`
	// The ID of the Counter used for the CounterAdjustment on the Account.
	CounterID param.Field[string] `json:"counterId,required"`
	// The date the CounterAdjustment is created for the Account _(in ISO-8601 date
	// format)_.
	//
	// **Note:** CounterAdjustments on Accounts are supported down to a _specific day_
	// of granularity - you cannot create more than one CounterAdjustment for any given
	// day using the same Counter and you'll receive an error if you try to do this.
	Date param.Field[string] `json:"date,required"`
	// Integer Value of the Counter used for the CounterAdjustment.
	//
	// **Note:** Use the new absolute value for the Counter for the selected date - if
	// it was 15 and has increased to 20, enter 20; if it was 15 and has decreased to
	// 10, enter 10. _Do not enter_ the plus or minus value relative to the previous
	// Counter value on the Account.
	Value param.Field[int64] `json:"value,required"`
	// Purchase Order Number for the Counter Adjustment. _(Optional)_
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

func (r CounterAdjustmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CounterAdjustmentListParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
	// List CounterAdjustment items for the Account UUID.
	AccountID param.Field[string] `query:"accountId"`
	// List CounterAdjustment items for the Counter UUID.
	CounterID param.Field[string] `query:"counterId"`
	// List CounterAdjustment items for the given date.
	Date      param.Field[string] `query:"date"`
	DateEnd   param.Field[string] `query:"dateEnd"`
	DateStart param.Field[string] `query:"dateStart"`
	// Only include CounterAdjustments with end dates earlier than this date.
	EndDateEnd param.Field[string] `query:"endDateEnd"`
	// Only include CounterAdjustments with end dates equal to or later than this date.
	EndDateStart param.Field[string] `query:"endDateStart"`
	// nextToken for multi page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of CounterAdjustments to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [CounterAdjustmentListParams]'s query parameters as
// `url.Values`.
func (r CounterAdjustmentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CounterAdjustmentDeleteParams struct {
	OrgID param.Field[string] `path:"orgId,required"`
}
