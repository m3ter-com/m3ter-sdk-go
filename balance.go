// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
	"github.com/tidwall/gjson"
)

// BalanceService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	Options              []option.RequestOption
	Transactions         *BalanceTransactionService
	ChargeSchedules      *BalanceChargeScheduleService
	TransactionSchedules *BalanceTransactionScheduleService
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r *BalanceService) {
	r = &BalanceService{}
	r.Options = opts
	r.Transactions = NewBalanceTransactionService(opts...)
	r.ChargeSchedules = NewBalanceChargeScheduleService(opts...)
	r.TransactionSchedules = NewBalanceTransactionScheduleService(opts...)
	return
}

// Create a new Balance for the given end customer Account.
//
// This endpoint allows you to create a new Balance for a specific end customer
// Account. The Balance details should be provided in the request body.
func (r *BalanceService) New(ctx context.Context, params BalanceNewParams, opts ...option.RequestOption) (res *Balance, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a specific Balance.
//
// This endpoint returns the details of the specified Balance.
func (r *BalanceService) Get(ctx context.Context, id string, query BalanceGetParams, opts ...option.RequestOption) (res *Balance, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Balance.
//
// This endpoint allows you to update the details of a specific Balance. The
// updated Balance details should be provided in the request body.
func (r *BalanceService) Update(ctx context.Context, id string, params BalanceUpdateParams, opts ...option.RequestOption) (res *Balance, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of all Balances for your Organization.
//
// This endpoint returns a list of all Balances associated with your organization.
// You can filter the Balances by the end customer's Account UUID and end dates,
// and paginate through them using the `pageSize` and `nextToken` parameters.
//
// **NOTE:** If a Balance has a rollover amount configured and you want to use the
// `endDateStart` or `endDateEnd` query parameters, the `rolloverEndDate` is used
// as the end date for the Balance.
func (r *BalanceService) List(ctx context.Context, params BalanceListParams, opts ...option.RequestOption) (res *pagination.Cursor[Balance], err error) {
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
	path := fmt.Sprintf("organizations/%s/balances", params.OrgID)
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

// Retrieve a list of all Balances for your Organization.
//
// This endpoint returns a list of all Balances associated with your organization.
// You can filter the Balances by the end customer's Account UUID and end dates,
// and paginate through them using the `pageSize` and `nextToken` parameters.
//
// **NOTE:** If a Balance has a rollover amount configured and you want to use the
// `endDateStart` or `endDateEnd` query parameters, the `rolloverEndDate` is used
// as the end date for the Balance.
func (r *BalanceService) ListAutoPaging(ctx context.Context, params BalanceListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[Balance] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete a specific Balance.
//
// This endpoint allows you to delete a specific Balance with the given UUID.
func (r *BalanceService) Delete(ctx context.Context, id string, body BalanceDeleteParams, opts ...option.RequestOption) (res *Balance, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Balance struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// The unique identifier (UUID) for the end customer Account the Balance belongs
	// to.
	AccountID string `json:"accountId"`
	// Allow balance amounts to fall below zero. This feature is enabled on request.
	// Please get in touch with m3ter Support or your m3ter contact if you would like
	// it enabling for your organization(s).
	AllowOverdraft bool `json:"allowOverdraft"`
	// The financial value that the Balance holds.
	Amount float64 `json:"amount"`
	// A description for the bill line items for charges drawn-down against the
	// Balance.
	BalanceDrawDownDescription string `json:"balanceDrawDownDescription"`
	// A unique short code assigned to the Balance.
	Code string `json:"code"`
	// Product ID that any Balance Consumed line items will be attributed to for
	// accounting purposes.(_Optional_)
	ConsumptionsAccountingProductID string `json:"consumptionsAccountingProductId"`
	// The unique identifier (UUID) for a Contract on the Account the Balance has been
	// added to.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) for the user who created the Balance.
	CreatedBy string `json:"createdBy"`
	// The currency code used for the Balance amount. For example: USD, GBP or EUR.
	Currency string `json:"currency"`
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
	CustomFields map[string]BalanceCustomFieldsUnion `json:"customFields"`
	// A description of the Balance.
	Description string `json:"description"`
	// The date and time _(in ISO 8601 format)_ when the Balance was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the Balance was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The date _(in ISO 8601 format)_ after which the Balance will no longer be
	// active.
	EndDate time.Time `json:"endDate" format:"date-time"`
	// Product ID that any Balance Fees line items will be attributed to for accounting
	// purposes.(_Optional_)
	FeesAccountingProductID string `json:"feesAccountingProductId"`
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
	StartDate time.Time `json:"startDate" format:"date-time"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64       `json:"version"`
	JSON    balanceJSON `json:"-"`
}

// balanceJSON contains the JSON metadata for the struct [Balance]
type balanceJSON struct {
	ID                              apijson.Field
	AccountID                       apijson.Field
	AllowOverdraft                  apijson.Field
	Amount                          apijson.Field
	BalanceDrawDownDescription      apijson.Field
	Code                            apijson.Field
	ConsumptionsAccountingProductID apijson.Field
	ContractID                      apijson.Field
	CreatedBy                       apijson.Field
	Currency                        apijson.Field
	CustomFields                    apijson.Field
	Description                     apijson.Field
	DtCreated                       apijson.Field
	DtLastModified                  apijson.Field
	EndDate                         apijson.Field
	FeesAccountingProductID         apijson.Field
	LastModifiedBy                  apijson.Field
	LineItemTypes                   apijson.Field
	Name                            apijson.Field
	OverageDescription              apijson.Field
	OverageSurchargePercent         apijson.Field
	ProductIDs                      apijson.Field
	RolloverAmount                  apijson.Field
	RolloverEndDate                 apijson.Field
	StartDate                       apijson.Field
	Version                         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *Balance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceCustomFieldsUnion interface {
	ImplementsBalanceCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceCustomFieldsUnion)(nil)).Elem(),
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

// Available line item types for Balances
type BalanceLineItemType string

const (
	BalanceLineItemTypeStandingCharge            BalanceLineItemType = "STANDING_CHARGE"
	BalanceLineItemTypeUsage                     BalanceLineItemType = "USAGE"
	BalanceLineItemTypeMinimumSpend              BalanceLineItemType = "MINIMUM_SPEND"
	BalanceLineItemTypeCounterRunningTotalCharge BalanceLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BalanceLineItemTypeCounterAdjustmentDebit    BalanceLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BalanceLineItemTypeAdHoc                     BalanceLineItemType = "AD_HOC"
)

func (r BalanceLineItemType) IsKnown() bool {
	switch r {
	case BalanceLineItemTypeStandingCharge, BalanceLineItemTypeUsage, BalanceLineItemTypeMinimumSpend, BalanceLineItemTypeCounterRunningTotalCharge, BalanceLineItemTypeCounterAdjustmentDebit, BalanceLineItemTypeAdHoc:
		return true
	}
	return false
}

type BalanceNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The unique identifier (UUID) for the end customer Account.
	AccountID param.Field[string] `json:"accountId" api:"required"`
	// Unique short code for the Balance.
	Code param.Field[string] `json:"code" api:"required"`
	// The currency code used for the Balance amount. For example: USD, GBP or EUR.
	Currency param.Field[string] `json:"currency" api:"required"`
	// The date _(in ISO 8601 format)_ after which the Balance will no longer be active
	// for the Account.
	//
	// **Note:** You can use the `rolloverEndDate` request parameter to define an
	// extended grace period for continued draw-down against the Balance if any amount
	// remains when the specified `endDate` is reached.
	EndDate param.Field[time.Time] `json:"endDate" api:"required" format:"date-time"`
	// The official name for the Balance.
	Name param.Field[string] `json:"name" api:"required"`
	// The date _(in ISO 8601 format)_ when the Balance becomes active.
	StartDate param.Field[time.Time] `json:"startDate" api:"required" format:"date-time"`
	// Allow balance amounts to fall below zero. This feature is enabled on request.
	// Please get in touch with m3ter Support or your m3ter contact if you would like
	// it enabling for your organization(s).
	AllowOverdraft param.Field[bool] `json:"allowOverdraft"`
	// A description for the bill line items for draw-down charges against the Balance.
	// _(Optional)._
	BalanceDrawDownDescription param.Field[string] `json:"balanceDrawDownDescription"`
	// Product ID that any Balance Consumed line items will be attributed to for
	// accounting purposes.(_Optional_)
	ConsumptionsAccountingProductID param.Field[string] `json:"consumptionsAccountingProductId"`
	// The unique identifier (UUID) of a Contract on the Account that the Balance will
	// be added to.
	ContractID param.Field[string] `json:"contractId"`
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
	CustomFields param.Field[map[string]BalanceNewParamsCustomFieldsUnion] `json:"customFields"`
	// A description of the Balance.
	Description param.Field[string] `json:"description"`
	// Product ID that any Balance Fees line items will be attributed to for accounting
	// purposes.(_Optional_)
	FeesAccountingProductID param.Field[string] `json:"feesAccountingProductId"`
	// Specify the line item charge types that can draw-down at billing against the
	// Balance amount. Options are:
	//
	// - `"MINIMUM_SPEND"`
	// - `"STANDING_CHARGE"`
	// - `"USAGE"`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	// - `AD_HOC`
	//
	// **NOTE:** If no charge types are specified, by default _all types_ can draw-down
	// against the Balance amount at billing.
	LineItemTypes param.Field[[]BalanceNewParamsLineItemType] `json:"lineItemTypes"`
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

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type BalanceNewParamsCustomFieldsUnion interface {
	ImplementsBalanceNewParamsCustomFieldsUnion()
}

// Available line item types for Balances
type BalanceNewParamsLineItemType string

const (
	BalanceNewParamsLineItemTypeStandingCharge            BalanceNewParamsLineItemType = "STANDING_CHARGE"
	BalanceNewParamsLineItemTypeUsage                     BalanceNewParamsLineItemType = "USAGE"
	BalanceNewParamsLineItemTypeMinimumSpend              BalanceNewParamsLineItemType = "MINIMUM_SPEND"
	BalanceNewParamsLineItemTypeCounterRunningTotalCharge BalanceNewParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BalanceNewParamsLineItemTypeCounterAdjustmentDebit    BalanceNewParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BalanceNewParamsLineItemTypeAdHoc                     BalanceNewParamsLineItemType = "AD_HOC"
)

func (r BalanceNewParamsLineItemType) IsKnown() bool {
	switch r {
	case BalanceNewParamsLineItemTypeStandingCharge, BalanceNewParamsLineItemTypeUsage, BalanceNewParamsLineItemTypeMinimumSpend, BalanceNewParamsLineItemTypeCounterRunningTotalCharge, BalanceNewParamsLineItemTypeCounterAdjustmentDebit, BalanceNewParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type BalanceGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}

type BalanceUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The unique identifier (UUID) for the end customer Account.
	AccountID param.Field[string] `json:"accountId" api:"required"`
	// Unique short code for the Balance.
	Code param.Field[string] `json:"code" api:"required"`
	// The currency code used for the Balance amount. For example: USD, GBP or EUR.
	Currency param.Field[string] `json:"currency" api:"required"`
	// The date _(in ISO 8601 format)_ after which the Balance will no longer be active
	// for the Account.
	//
	// **Note:** You can use the `rolloverEndDate` request parameter to define an
	// extended grace period for continued draw-down against the Balance if any amount
	// remains when the specified `endDate` is reached.
	EndDate param.Field[time.Time] `json:"endDate" api:"required" format:"date-time"`
	// The official name for the Balance.
	Name param.Field[string] `json:"name" api:"required"`
	// The date _(in ISO 8601 format)_ when the Balance becomes active.
	StartDate param.Field[time.Time] `json:"startDate" api:"required" format:"date-time"`
	// Allow balance amounts to fall below zero. This feature is enabled on request.
	// Please get in touch with m3ter Support or your m3ter contact if you would like
	// it enabling for your organization(s).
	AllowOverdraft param.Field[bool] `json:"allowOverdraft"`
	// A description for the bill line items for draw-down charges against the Balance.
	// _(Optional)._
	BalanceDrawDownDescription param.Field[string] `json:"balanceDrawDownDescription"`
	// Product ID that any Balance Consumed line items will be attributed to for
	// accounting purposes.(_Optional_)
	ConsumptionsAccountingProductID param.Field[string] `json:"consumptionsAccountingProductId"`
	// The unique identifier (UUID) of a Contract on the Account that the Balance will
	// be added to.
	ContractID param.Field[string] `json:"contractId"`
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
	CustomFields param.Field[map[string]BalanceUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// A description of the Balance.
	Description param.Field[string] `json:"description"`
	// Product ID that any Balance Fees line items will be attributed to for accounting
	// purposes.(_Optional_)
	FeesAccountingProductID param.Field[string] `json:"feesAccountingProductId"`
	// Specify the line item charge types that can draw-down at billing against the
	// Balance amount. Options are:
	//
	// - `"MINIMUM_SPEND"`
	// - `"STANDING_CHARGE"`
	// - `"USAGE"`
	// - `"COUNTER_RUNNING_TOTAL_CHARGE"`
	// - `"COUNTER_ADJUSTMENT_DEBIT"`
	// - `AD_HOC`
	//
	// **NOTE:** If no charge types are specified, by default _all types_ can draw-down
	// against the Balance amount at billing.
	LineItemTypes param.Field[[]BalanceUpdateParamsLineItemType] `json:"lineItemTypes"`
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

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type BalanceUpdateParamsCustomFieldsUnion interface {
	ImplementsBalanceUpdateParamsCustomFieldsUnion()
}

// Available line item types for Balances
type BalanceUpdateParamsLineItemType string

const (
	BalanceUpdateParamsLineItemTypeStandingCharge            BalanceUpdateParamsLineItemType = "STANDING_CHARGE"
	BalanceUpdateParamsLineItemTypeUsage                     BalanceUpdateParamsLineItemType = "USAGE"
	BalanceUpdateParamsLineItemTypeMinimumSpend              BalanceUpdateParamsLineItemType = "MINIMUM_SPEND"
	BalanceUpdateParamsLineItemTypeCounterRunningTotalCharge BalanceUpdateParamsLineItemType = "COUNTER_RUNNING_TOTAL_CHARGE"
	BalanceUpdateParamsLineItemTypeCounterAdjustmentDebit    BalanceUpdateParamsLineItemType = "COUNTER_ADJUSTMENT_DEBIT"
	BalanceUpdateParamsLineItemTypeAdHoc                     BalanceUpdateParamsLineItemType = "AD_HOC"
)

func (r BalanceUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case BalanceUpdateParamsLineItemTypeStandingCharge, BalanceUpdateParamsLineItemTypeUsage, BalanceUpdateParamsLineItemTypeMinimumSpend, BalanceUpdateParamsLineItemTypeCounterRunningTotalCharge, BalanceUpdateParamsLineItemTypeCounterAdjustmentDebit, BalanceUpdateParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type BalanceListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The unique identifier (UUID) for the end customer's account.
	AccountID param.Field[string] `query:"accountId"`
	Contract  param.Field[string] `query:"contract"`
	// Filter Balances by contract id. Use ” with accountId to fetch unlinked
	// balances.
	ContractID param.Field[string] `query:"contractId"`
	// Only include Balances with end dates earlier than this date. If a Balance has a
	// rollover amount configured, then the `rolloverEndDate` will be used as the end
	// date.
	EndDateEnd param.Field[string] `query:"endDateEnd"`
	// Only include Balances with end dates equal to or later than this date. If a
	// Balance has a rollover amount configured, then the `rolloverEndDate` will be
	// used as the end date.
	EndDateStart param.Field[string] `query:"endDateStart"`
	// A list of unique identifiers (UUIDs) for specific Balances to retrieve.
	IDs param.Field[[]string] `query:"ids"`
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

type BalanceDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
