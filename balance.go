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

// BalanceService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options      []option.RequestOption
	Transactions *BalanceTransactionService
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r *BalanceService) {
	r = &BalanceService{}
	r.Options = opts
	r.Transactions = NewBalanceTransactionService(opts...)
	return
}

// Create a new Balance for the given end customer Account.
//
// This endpoint allows you to create a new Balance for a specific end customer
// Account. The Balance details should be provided in the request body.
func (r *BalanceService) New(ctx context.Context, orgID string, body BalanceNewParams, opts ...option.RequestOption) (res *Balance, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific Balance.
//
// This endpoint returns the details of the specified Balance.
func (r *BalanceService) Get(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Balance, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Balance.
//
// This endpoint allows you to update the details of a specific Balance. The
// updated Balance details should be provided in the request body.
func (r *BalanceService) Update(ctx context.Context, orgID string, id string, body BalanceUpdateParams, opts ...option.RequestOption) (res *Balance, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a list of all Balances for your Organization.
//
// This endpoint returns a list of all Balances associated with your organization.
// You can filter the Balances by the end customer's Account UUID and end dates,
// and paginate through them using the `pageSize` and `nextToken` parameters.
func (r *BalanceService) List(ctx context.Context, orgID string, query BalanceListParams, opts ...option.RequestOption) (res *BalanceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances", orgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific Balance.
//
// This endpoint allows you to delete a specific Balance with the given UUID.
func (r *BalanceService) Delete(ctx context.Context, orgID string, id string, opts ...option.RequestOption) (res *Balance, err error) {
	opts = append(r.Options[:], opts...)
	if orgID == "" {
		err = errors.New("missing required orgId parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s", orgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Balance struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// The unique identifier (UUID) for the end customer Account the Balance belongs
	// to.
	AccountID string `json:"accountId"`
	// The financial value that the Balance holds.
	Amount float64 `json:"amount"`
	// A description for the bill line items for charges drawn-down against the
	// Balance.
	BalanceDrawDownDescription string `json:"balanceDrawDownDescription"`
	// A unique short code assigned to the Balance.
	Code string `json:"code"`
	// The unique identifier (UUID) for the user who created the Balance.
	CreatedBy string `json:"createdBy"`
	// The currency code used for the Balance amount. For example: USD, GBP or EUR.
	Currency string `json:"currency"`
	// A description of the Balance.
	Description string `json:"description"`
	// The date and time _(in ISO 8601 format)_ when the Balance was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the Balance was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The date _(in ISO 8601 format)_ after which the Balance will no longer be
	// active.
	EndDate time.Time `json:"endDate" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance.
	LastModifiedBy string `json:"lastModifiedBy"`
	// A list of line item charge types that can draw-down against the Balance amount
	// at billing.
	LineItemTypes []BalanceLineItemType `json:"lineItemTypes"`
	// The official name of the Balance.
	Name string `json:"name"`
	// A description for overage charges.
	OverageDescription string `json:"overageDescription"`
	// The percentage surcharge applied to overage charges _(usage above the Balance)_.
	OverageSurchargePercent float64 `json:"overageSurchargePercent"`
	// A list of Product IDs whose consumption charges due at billing can be drawn-down
	// against the Balance amount.
	ProductIDs []string `json:"productIds"`
	// The maximum amount that can be carried over past the Balance end date and
	// draw-down against for billing if there is an unused Balance amount remaining
	// when the Balance end date is reached.
	RolloverAmount float64 `json:"rolloverAmount"`
	// The end date _(in ISO 8601 format)_ for the rollover grace period, which is the
	// period that unused Balance amounts can be carried over beyond the specified
	// Balance `endDate` and continue to be drawn-down against for billing.
	RolloverEndDate time.Time `json:"rolloverEndDate" format:"date-time"`
	// The date _(in ISO 8601 format)_ when the Balance becomes active.
	StartDate time.Time   `json:"startDate" format:"date-time"`
	JSON      balanceJSON `json:"-"`
}

// balanceJSON contains the JSON metadata for the struct [Balance]
type balanceJSON struct {
	ID                         apijson.Field
	Version                    apijson.Field
	AccountID                  apijson.Field
	Amount                     apijson.Field
	BalanceDrawDownDescription apijson.Field
	Code                       apijson.Field
	CreatedBy                  apijson.Field
	Currency                   apijson.Field
	Description                apijson.Field
	DtCreated                  apijson.Field
	DtLastModified             apijson.Field
	EndDate                    apijson.Field
	LastModifiedBy             apijson.Field
	LineItemTypes              apijson.Field
	Name                       apijson.Field
	OverageDescription         apijson.Field
	OverageSurchargePercent    apijson.Field
	ProductIDs                 apijson.Field
	RolloverAmount             apijson.Field
	RolloverEndDate            apijson.Field
	StartDate                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceJSON) RawJSON() string {
	return r.raw
}

// Available line item types for Balances
type BalanceLineItemType string

const (
	BalanceLineItemTypeStandingCharge            BalanceLineItemType = "STANDING_CHARGE"
	BalanceLineItemTypeUsage                     BalanceLineItemType = "USAGE"
	BalanceLineItemTypeMinimumSpend              BalanceLineItemType = "MINIMUM_SPEND"
	BalanceLineItemTypeCounterRunningTotalCharge BalanceLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BalanceLineItemTypeCounterAdjustmentDebit    BalanceLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
)

func (r BalanceLineItemType) IsKnown() bool {
	switch r {
	case BalanceLineItemTypeStandingCharge, BalanceLineItemTypeUsage, BalanceLineItemTypeMinimumSpend, BalanceLineItemTypeCounterRunningTotalCharge, BalanceLineItemTypeCounterAdjustmentDebit:
		return true
	}
	return false
}

type BalanceListResponse = interface{}

type BalanceNewParams struct {
	// The unique identifier (UUID) for the end customer Account.
	AccountID param.Field[string] `json:"accountId,required"`
	// The currency code used for the Balance amount. For example: USD, GBP or EUR.
	Currency param.Field[string] `json:"currency,required"`
	// The date _(in ISO 8601 format)_ after which the Balance will no longer be active
	// for the Account.
	//
	// **Note:** You can use the `rolloverEndDate` request parameter to define an
	// extended grace period for continued draw-down against the Balance if any amount
	// remains when the specified `endDate` is reached.
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date-time"`
	// The date _(in ISO 8601 format)_ when the Balance becomes active.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// A description for the bill line items for draw-down charges against the Balance.
	// _(Optional)._
	BalanceDrawDownDescription param.Field[string] `json:"balanceDrawDownDescription"`
	// Unique short code for the Balance.
	Code param.Field[string] `json:"code"`
	// A description of the Balance.
	Description param.Field[string] `json:"description"`
	// Specify the line item charge types that can draw-down at billing against the
	// Balance amount. Options are:
	//
	// - `"MINIMUM_SPEND"`
	// - `"STANDING_CHARGE"`
	// - `"USAGE"`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	//
	// **NOTE:** If no charge types are specified, by default _all types_ can draw-down
	// against the Balance amount at billing.
	LineItemTypes param.Field[[]BalanceNewParamsLineItemType] `json:"lineItemTypes"`
	// The official name for the Balance.
	Name param.Field[string] `json:"name"`
	// A description for Bill line items overage charges.
	OverageDescription param.Field[string] `json:"overageDescription"`
	// Define a surcharge level, as a percentage of regular usage rating, applied to
	// overages _(usage charges that exceed the Balance amount)_. For example, if the
	// regular usage rate is $10 per unit of usage consumed and
	// `overageSurchargePercent` is set at 10%, then any usage charged above the
	// original Balance amount is charged at $11 per unit of usage.
	OverageSurchargePercent param.Field[float64] `json:"overageSurchargePercent"`
	// Specify the Products whose consumption charges due at billing can be drawn-down
	// against the Balance amount.
	//
	// **Note:** If you don't specify any Products for Balance draw-down, by default
	// the consumption charges for any Product the Account consumes will be drawn-down
	// against the Balance amount.
	ProductIDs param.Field[[]string] `json:"productIds"`
	// The maximum amount that can be carried over past the Balance end date for
	// draw-down at billing if there is any unused Balance amount when the end date is
	// reached. Works with `rolloverEndDate` to define the amount and duration of a
	// Balance "grace period". _(Optional)_
	//
	// **Notes:**
	//
	//   - If you leave `rolloverAmount` empty and only enter a `rolloverEndDate`, any
	//     amount left over after the Balance end date is reached will be drawn-down
	//     against up to the specified `rolloverEndDate`.
	//   - You must enter a `rolloverEndDate`. If you only enter a `rolloverAmount`
	//     without entering a `rolloverEndDate`, you'll receive an error when trying to
	//     create or update the Balance.
	//   - If you don't want to grant any grace period for outstanding Balance amounts,
	//     then do not use `rolloverAmount` and `rolloverEndDate`.
	RolloverAmount param.Field[float64] `json:"rolloverAmount"`
	// The end date _(in ISO 8601 format)_ for the grace period during which unused
	// Balance amounts can be carried over and drawn-down against at billing.
	//
	// **Note:** Use `rolloverAmount` if you want to specify a maximum amount that can
	// be carried over and made available for draw-down.
	RolloverEndDate param.Field[time.Time] `json:"rolloverEndDate" format:"date-time"`
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

func (r BalanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Available line item types for Balances
type BalanceNewParamsLineItemType string

const (
	BalanceNewParamsLineItemTypeStandingCharge            BalanceNewParamsLineItemType = "STANDING_CHARGE"
	BalanceNewParamsLineItemTypeUsage                     BalanceNewParamsLineItemType = "USAGE"
	BalanceNewParamsLineItemTypeMinimumSpend              BalanceNewParamsLineItemType = "MINIMUM_SPEND"
	BalanceNewParamsLineItemTypeCounterRunningTotalCharge BalanceNewParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BalanceNewParamsLineItemTypeCounterAdjustmentDebit    BalanceNewParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
)

func (r BalanceNewParamsLineItemType) IsKnown() bool {
	switch r {
	case BalanceNewParamsLineItemTypeStandingCharge, BalanceNewParamsLineItemTypeUsage, BalanceNewParamsLineItemTypeMinimumSpend, BalanceNewParamsLineItemTypeCounterRunningTotalCharge, BalanceNewParamsLineItemTypeCounterAdjustmentDebit:
		return true
	}
	return false
}

type BalanceUpdateParams struct {
	// The unique identifier (UUID) for the end customer Account.
	AccountID param.Field[string] `json:"accountId,required"`
	// The currency code used for the Balance amount. For example: USD, GBP or EUR.
	Currency param.Field[string] `json:"currency,required"`
	// The date _(in ISO 8601 format)_ after which the Balance will no longer be active
	// for the Account.
	//
	// **Note:** You can use the `rolloverEndDate` request parameter to define an
	// extended grace period for continued draw-down against the Balance if any amount
	// remains when the specified `endDate` is reached.
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date-time"`
	// The date _(in ISO 8601 format)_ when the Balance becomes active.
	StartDate param.Field[time.Time] `json:"startDate,required" format:"date-time"`
	// A description for the bill line items for draw-down charges against the Balance.
	// _(Optional)._
	BalanceDrawDownDescription param.Field[string] `json:"balanceDrawDownDescription"`
	// Unique short code for the Balance.
	Code param.Field[string] `json:"code"`
	// A description of the Balance.
	Description param.Field[string] `json:"description"`
	// Specify the line item charge types that can draw-down at billing against the
	// Balance amount. Options are:
	//
	// - `"MINIMUM_SPEND"`
	// - `"STANDING_CHARGE"`
	// - `"USAGE"`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	//
	// **NOTE:** If no charge types are specified, by default _all types_ can draw-down
	// against the Balance amount at billing.
	LineItemTypes param.Field[[]BalanceUpdateParamsLineItemType] `json:"lineItemTypes"`
	// The official name for the Balance.
	Name param.Field[string] `json:"name"`
	// A description for Bill line items overage charges.
	OverageDescription param.Field[string] `json:"overageDescription"`
	// Define a surcharge level, as a percentage of regular usage rating, applied to
	// overages _(usage charges that exceed the Balance amount)_. For example, if the
	// regular usage rate is $10 per unit of usage consumed and
	// `overageSurchargePercent` is set at 10%, then any usage charged above the
	// original Balance amount is charged at $11 per unit of usage.
	OverageSurchargePercent param.Field[float64] `json:"overageSurchargePercent"`
	// Specify the Products whose consumption charges due at billing can be drawn-down
	// against the Balance amount.
	//
	// **Note:** If you don't specify any Products for Balance draw-down, by default
	// the consumption charges for any Product the Account consumes will be drawn-down
	// against the Balance amount.
	ProductIDs param.Field[[]string] `json:"productIds"`
	// The maximum amount that can be carried over past the Balance end date for
	// draw-down at billing if there is any unused Balance amount when the end date is
	// reached. Works with `rolloverEndDate` to define the amount and duration of a
	// Balance "grace period". _(Optional)_
	//
	// **Notes:**
	//
	//   - If you leave `rolloverAmount` empty and only enter a `rolloverEndDate`, any
	//     amount left over after the Balance end date is reached will be drawn-down
	//     against up to the specified `rolloverEndDate`.
	//   - You must enter a `rolloverEndDate`. If you only enter a `rolloverAmount`
	//     without entering a `rolloverEndDate`, you'll receive an error when trying to
	//     create or update the Balance.
	//   - If you don't want to grant any grace period for outstanding Balance amounts,
	//     then do not use `rolloverAmount` and `rolloverEndDate`.
	RolloverAmount param.Field[float64] `json:"rolloverAmount"`
	// The end date _(in ISO 8601 format)_ for the grace period during which unused
	// Balance amounts can be carried over and drawn-down against at billing.
	//
	// **Note:** Use `rolloverAmount` if you want to specify a maximum amount that can
	// be carried over and made available for draw-down.
	RolloverEndDate param.Field[time.Time] `json:"rolloverEndDate" format:"date-time"`
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

func (r BalanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Available line item types for Balances
type BalanceUpdateParamsLineItemType string

const (
	BalanceUpdateParamsLineItemTypeStandingCharge            BalanceUpdateParamsLineItemType = "STANDING_CHARGE"
	BalanceUpdateParamsLineItemTypeUsage                     BalanceUpdateParamsLineItemType = "USAGE"
	BalanceUpdateParamsLineItemTypeMinimumSpend              BalanceUpdateParamsLineItemType = "MINIMUM_SPEND"
	BalanceUpdateParamsLineItemTypeCounterRunningTotalCharge BalanceUpdateParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BalanceUpdateParamsLineItemTypeCounterAdjustmentDebit    BalanceUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
)

func (r BalanceUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case BalanceUpdateParamsLineItemTypeStandingCharge, BalanceUpdateParamsLineItemTypeUsage, BalanceUpdateParamsLineItemTypeMinimumSpend, BalanceUpdateParamsLineItemTypeCounterRunningTotalCharge, BalanceUpdateParamsLineItemTypeCounterAdjustmentDebit:
		return true
	}
	return false
}

type BalanceListParams struct {
	// The unique identifier (UUID) for the end customer's account.
	AccountID param.Field[string] `query:"accountId"`
	// Only include Balances with end dates earlier than this date.
	EndDateEnd param.Field[string] `query:"endDateEnd"`
	// Only include Balances with end dates equal to or later than this date.
	EndDateStart param.Field[string] `query:"endDateStart"`
	// The `nextToken` for retrieving the next page of Balances. It is used to fetch
	// the next page of Balances in a paginated list.
	NextToken param.Field[string] `query:"nextToken"`
	// The maximum number of Balances to return per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [BalanceListParams]'s query parameters as `url.Values`.
func (r BalanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
