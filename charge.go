// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/apiquery"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/packages/pagination"
)

// ChargeService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChargeService] method instead.
type ChargeService struct {
	Options []option.RequestOption
}

// NewChargeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChargeService(opts ...option.RequestOption) (r *ChargeService) {
	r = &ChargeService{}
	r.Options = opts
	return
}

// Create a new Charge.
//
// **NOTES:**
//
//   - To create an ad-hoc Charge on an Account, use the `accountId` request
//     parameter.
//   - To create a balance fee Charge for a Balance, use the `entityId` request
//     parameter to specify which Balance on an Account the Charge is for.
//   - To define the value of the Charge amount that is billed, you can simply
//     specify an `amount` or use a number of `units` together with a `unitPrice` for
//     a calculated value = units x unit price. But you cannot specify _both an
//     amount and units/unit price_.
func (r *ChargeService) New(ctx context.Context, params ChargeNewParams, opts ...option.RequestOption) (res *ChargeNewResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/charges", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a Charge for the given UUID.
func (r *ChargeService) Get(ctx context.Context, id string, query ChargeGetParams, opts ...option.RequestOption) (res *ChargeGetResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/charges/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Charge for the given UUID.
//
// **NOTE:** When you update a Charge on an Account, you can provide either a
// Charge `amount` or Charge `units` together with a `unitPrice`, but _not both_.
func (r *ChargeService) Update(ctx context.Context, id string, params ChargeUpdateParams, opts ...option.RequestOption) (res *ChargeUpdateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/charges/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Charge entities
func (r *ChargeService) List(ctx context.Context, params ChargeListParams, opts ...option.RequestOption) (res *pagination.Cursor[ChargeListResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/charges", params.OrgID)
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

// Retrieve a list of Charge entities
func (r *ChargeService) ListAutoPaging(ctx context.Context, params ChargeListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[ChargeListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Charge for the given UUID.
func (r *ChargeService) Delete(ctx context.Context, id string, body ChargeDeleteParams, opts ...option.RequestOption) (res *ChargeDeleteResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/charges/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Response containing a Charge entity
type ChargeNewResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The ID of the Account the Charge was created for.
	AccountID string `json:"accountId"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID string `json:"accountingProductId"`
	// The Charge amount. If `amount` has been defined, then `units` and `unitPrice`
	// cannot be used.
	Amount float64 `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate time.Time `json:"billDate" format:"date"`
	// The ID of the Bill created for this Charge.
	BillID string `json:"billId"`
	// The unique short code of the Charge.
	Code string `json:"code"`
	// The ID of a Contract on the Account that the Charge has been added to.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) of the user who created the Charge.
	CreatedBy string `json:"createdBy"`
	// Charge currency.
	Currency string `json:"currency"`
	// The description added to the Bill line item for the Charge.
	Description string `json:"description"`
	// The date and time (_in ISO-8601 format_) when the Charge was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO 8601 format_) when the Charge was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	EntityID string `json:"entityId"`
	// The entity type the Charge has been created for.
	EntityType ChargeNewResponseEntityType `json:"entityType"`
	// The unique identifier (UUID) of the user who last modified the Charge.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The line item type used for billing a Charge.
	LineItemType ChargeNewResponseLineItemType `json:"lineItemType"`
	// Name of the Charge. Added to the Bill line item description for Charge.
	Name string `json:"name"`
	// Information about the Charge for accounting purposes, such as the reason it was
	// created. This information will not be added to the created Bill line item for
	// the Charge.
	Notes string `json:"notes"`
	// The ID of the Balance Charge Schedule that created the Charge.
	ScheduleID string `json:"scheduleId"`
	// The service period end date (_in ISO-8601 format_) for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge .
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit Price for the Charge. Provided together with `units`:
	//
	// - Null if the Charge was created with `amount` only.
	// - If `units` and `unitPrice` are provided, `amount` cannot be used.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units of the Charge. Provided together with `unitPrice`. If `units`
	// and `unitPrice` are provided, `amount` cannot be used.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                 `json:"version"`
	JSON    chargeNewResponseJSON `json:"-"`
}

// chargeNewResponseJSON contains the JSON metadata for the struct
// [ChargeNewResponse]
type chargeNewResponseJSON struct {
	ID                     apijson.Field
	AccountID              apijson.Field
	AccountingProductID    apijson.Field
	Amount                 apijson.Field
	BillDate               apijson.Field
	BillID                 apijson.Field
	Code                   apijson.Field
	ContractID             apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	EntityID               apijson.Field
	EntityType             apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	Name                   apijson.Field
	Notes                  apijson.Field
	ScheduleID             apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ChargeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeNewResponseJSON) RawJSON() string {
	return r.raw
}

// The entity type the Charge has been created for.
type ChargeNewResponseEntityType string

const (
	ChargeNewResponseEntityTypeAdHoc   ChargeNewResponseEntityType = "AD_HOC"
	ChargeNewResponseEntityTypeBalance ChargeNewResponseEntityType = "BALANCE"
)

func (r ChargeNewResponseEntityType) IsKnown() bool {
	switch r {
	case ChargeNewResponseEntityTypeAdHoc, ChargeNewResponseEntityTypeBalance:
		return true
	}
	return false
}

// The line item type used for billing a Charge.
type ChargeNewResponseLineItemType string

const (
	ChargeNewResponseLineItemTypeBalanceFee ChargeNewResponseLineItemType = "BALANCE_FEE"
	ChargeNewResponseLineItemTypeAdHoc      ChargeNewResponseLineItemType = "AD_HOC"
)

func (r ChargeNewResponseLineItemType) IsKnown() bool {
	switch r {
	case ChargeNewResponseLineItemTypeBalanceFee, ChargeNewResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

// Response containing a Charge entity
type ChargeGetResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The ID of the Account the Charge was created for.
	AccountID string `json:"accountId"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID string `json:"accountingProductId"`
	// The Charge amount. If `amount` has been defined, then `units` and `unitPrice`
	// cannot be used.
	Amount float64 `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate time.Time `json:"billDate" format:"date"`
	// The ID of the Bill created for this Charge.
	BillID string `json:"billId"`
	// The unique short code of the Charge.
	Code string `json:"code"`
	// The ID of a Contract on the Account that the Charge has been added to.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) of the user who created the Charge.
	CreatedBy string `json:"createdBy"`
	// Charge currency.
	Currency string `json:"currency"`
	// The description added to the Bill line item for the Charge.
	Description string `json:"description"`
	// The date and time (_in ISO-8601 format_) when the Charge was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO 8601 format_) when the Charge was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	EntityID string `json:"entityId"`
	// The entity type the Charge has been created for.
	EntityType ChargeGetResponseEntityType `json:"entityType"`
	// The unique identifier (UUID) of the user who last modified the Charge.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The line item type used for billing a Charge.
	LineItemType ChargeGetResponseLineItemType `json:"lineItemType"`
	// Name of the Charge. Added to the Bill line item description for Charge.
	Name string `json:"name"`
	// Information about the Charge for accounting purposes, such as the reason it was
	// created. This information will not be added to the created Bill line item for
	// the Charge.
	Notes string `json:"notes"`
	// The ID of the Balance Charge Schedule that created the Charge.
	ScheduleID string `json:"scheduleId"`
	// The service period end date (_in ISO-8601 format_) for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge .
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit Price for the Charge. Provided together with `units`:
	//
	// - Null if the Charge was created with `amount` only.
	// - If `units` and `unitPrice` are provided, `amount` cannot be used.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units of the Charge. Provided together with `unitPrice`. If `units`
	// and `unitPrice` are provided, `amount` cannot be used.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                 `json:"version"`
	JSON    chargeGetResponseJSON `json:"-"`
}

// chargeGetResponseJSON contains the JSON metadata for the struct
// [ChargeGetResponse]
type chargeGetResponseJSON struct {
	ID                     apijson.Field
	AccountID              apijson.Field
	AccountingProductID    apijson.Field
	Amount                 apijson.Field
	BillDate               apijson.Field
	BillID                 apijson.Field
	Code                   apijson.Field
	ContractID             apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	EntityID               apijson.Field
	EntityType             apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	Name                   apijson.Field
	Notes                  apijson.Field
	ScheduleID             apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ChargeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeGetResponseJSON) RawJSON() string {
	return r.raw
}

// The entity type the Charge has been created for.
type ChargeGetResponseEntityType string

const (
	ChargeGetResponseEntityTypeAdHoc   ChargeGetResponseEntityType = "AD_HOC"
	ChargeGetResponseEntityTypeBalance ChargeGetResponseEntityType = "BALANCE"
)

func (r ChargeGetResponseEntityType) IsKnown() bool {
	switch r {
	case ChargeGetResponseEntityTypeAdHoc, ChargeGetResponseEntityTypeBalance:
		return true
	}
	return false
}

// The line item type used for billing a Charge.
type ChargeGetResponseLineItemType string

const (
	ChargeGetResponseLineItemTypeBalanceFee ChargeGetResponseLineItemType = "BALANCE_FEE"
	ChargeGetResponseLineItemTypeAdHoc      ChargeGetResponseLineItemType = "AD_HOC"
)

func (r ChargeGetResponseLineItemType) IsKnown() bool {
	switch r {
	case ChargeGetResponseLineItemTypeBalanceFee, ChargeGetResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

// Response containing a Charge entity
type ChargeUpdateResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The ID of the Account the Charge was created for.
	AccountID string `json:"accountId"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID string `json:"accountingProductId"`
	// The Charge amount. If `amount` has been defined, then `units` and `unitPrice`
	// cannot be used.
	Amount float64 `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate time.Time `json:"billDate" format:"date"`
	// The ID of the Bill created for this Charge.
	BillID string `json:"billId"`
	// The unique short code of the Charge.
	Code string `json:"code"`
	// The ID of a Contract on the Account that the Charge has been added to.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) of the user who created the Charge.
	CreatedBy string `json:"createdBy"`
	// Charge currency.
	Currency string `json:"currency"`
	// The description added to the Bill line item for the Charge.
	Description string `json:"description"`
	// The date and time (_in ISO-8601 format_) when the Charge was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO 8601 format_) when the Charge was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	EntityID string `json:"entityId"`
	// The entity type the Charge has been created for.
	EntityType ChargeUpdateResponseEntityType `json:"entityType"`
	// The unique identifier (UUID) of the user who last modified the Charge.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The line item type used for billing a Charge.
	LineItemType ChargeUpdateResponseLineItemType `json:"lineItemType"`
	// Name of the Charge. Added to the Bill line item description for Charge.
	Name string `json:"name"`
	// Information about the Charge for accounting purposes, such as the reason it was
	// created. This information will not be added to the created Bill line item for
	// the Charge.
	Notes string `json:"notes"`
	// The ID of the Balance Charge Schedule that created the Charge.
	ScheduleID string `json:"scheduleId"`
	// The service period end date (_in ISO-8601 format_) for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge .
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit Price for the Charge. Provided together with `units`:
	//
	// - Null if the Charge was created with `amount` only.
	// - If `units` and `unitPrice` are provided, `amount` cannot be used.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units of the Charge. Provided together with `unitPrice`. If `units`
	// and `unitPrice` are provided, `amount` cannot be used.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                    `json:"version"`
	JSON    chargeUpdateResponseJSON `json:"-"`
}

// chargeUpdateResponseJSON contains the JSON metadata for the struct
// [ChargeUpdateResponse]
type chargeUpdateResponseJSON struct {
	ID                     apijson.Field
	AccountID              apijson.Field
	AccountingProductID    apijson.Field
	Amount                 apijson.Field
	BillDate               apijson.Field
	BillID                 apijson.Field
	Code                   apijson.Field
	ContractID             apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	EntityID               apijson.Field
	EntityType             apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	Name                   apijson.Field
	Notes                  apijson.Field
	ScheduleID             apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ChargeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The entity type the Charge has been created for.
type ChargeUpdateResponseEntityType string

const (
	ChargeUpdateResponseEntityTypeAdHoc   ChargeUpdateResponseEntityType = "AD_HOC"
	ChargeUpdateResponseEntityTypeBalance ChargeUpdateResponseEntityType = "BALANCE"
)

func (r ChargeUpdateResponseEntityType) IsKnown() bool {
	switch r {
	case ChargeUpdateResponseEntityTypeAdHoc, ChargeUpdateResponseEntityTypeBalance:
		return true
	}
	return false
}

// The line item type used for billing a Charge.
type ChargeUpdateResponseLineItemType string

const (
	ChargeUpdateResponseLineItemTypeBalanceFee ChargeUpdateResponseLineItemType = "BALANCE_FEE"
	ChargeUpdateResponseLineItemTypeAdHoc      ChargeUpdateResponseLineItemType = "AD_HOC"
)

func (r ChargeUpdateResponseLineItemType) IsKnown() bool {
	switch r {
	case ChargeUpdateResponseLineItemTypeBalanceFee, ChargeUpdateResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

type ChargeListResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The ID of the Account the Charge was created for.
	AccountID string `json:"accountId"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID string `json:"accountingProductId"`
	// The Charge amount. If `amount` has been defined, then `units` and `unitPrice`
	// cannot be used.
	Amount float64 `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate time.Time `json:"billDate" format:"date"`
	// The ID of the Bill created for this Charge.
	BillID string `json:"billId"`
	// The unique short code of the Charge.
	Code string `json:"code"`
	// The ID of a Contract on the Account that the Charge has been added to.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) of the user who created the Charge.
	CreatedBy string `json:"createdBy"`
	// Charge currency.
	Currency string `json:"currency"`
	// The description added to the Bill line item for the Charge.
	Description string `json:"description"`
	// The date and time (_in ISO-8601 format_) when the Charge was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO 8601 format_) when the Charge was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	EntityID string `json:"entityId"`
	// The entity type the Charge has been created for.
	EntityType ChargeListResponseEntityType `json:"entityType"`
	// The unique identifier (UUID) of the user who last modified the Charge.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The line item type used for billing a Charge.
	LineItemType ChargeListResponseLineItemType `json:"lineItemType"`
	// Name of the Charge. Added to the Bill line item description for Charge.
	Name string `json:"name"`
	// Information about the Charge for accounting purposes, such as the reason it was
	// created. This information will not be added to the created Bill line item for
	// the Charge.
	Notes string `json:"notes"`
	// The ID of the Balance Charge Schedule that created the Charge.
	ScheduleID string `json:"scheduleId"`
	// The service period end date (_in ISO-8601 format_) for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge .
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit Price for the Charge. Provided together with `units`:
	//
	// - Null if the Charge was created with `amount` only.
	// - If `units` and `unitPrice` are provided, `amount` cannot be used.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units of the Charge. Provided together with `unitPrice`. If `units`
	// and `unitPrice` are provided, `amount` cannot be used.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                  `json:"version"`
	JSON    chargeListResponseJSON `json:"-"`
}

// chargeListResponseJSON contains the JSON metadata for the struct
// [ChargeListResponse]
type chargeListResponseJSON struct {
	ID                     apijson.Field
	AccountID              apijson.Field
	AccountingProductID    apijson.Field
	Amount                 apijson.Field
	BillDate               apijson.Field
	BillID                 apijson.Field
	Code                   apijson.Field
	ContractID             apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	EntityID               apijson.Field
	EntityType             apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	Name                   apijson.Field
	Notes                  apijson.Field
	ScheduleID             apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ChargeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeListResponseJSON) RawJSON() string {
	return r.raw
}

// The entity type the Charge has been created for.
type ChargeListResponseEntityType string

const (
	ChargeListResponseEntityTypeAdHoc   ChargeListResponseEntityType = "AD_HOC"
	ChargeListResponseEntityTypeBalance ChargeListResponseEntityType = "BALANCE"
)

func (r ChargeListResponseEntityType) IsKnown() bool {
	switch r {
	case ChargeListResponseEntityTypeAdHoc, ChargeListResponseEntityTypeBalance:
		return true
	}
	return false
}

// The line item type used for billing a Charge.
type ChargeListResponseLineItemType string

const (
	ChargeListResponseLineItemTypeBalanceFee ChargeListResponseLineItemType = "BALANCE_FEE"
	ChargeListResponseLineItemTypeAdHoc      ChargeListResponseLineItemType = "AD_HOC"
)

func (r ChargeListResponseLineItemType) IsKnown() bool {
	switch r {
	case ChargeListResponseLineItemTypeBalanceFee, ChargeListResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

// Response containing a Charge entity
type ChargeDeleteResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The ID of the Account the Charge was created for.
	AccountID string `json:"accountId"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID string `json:"accountingProductId"`
	// The Charge amount. If `amount` has been defined, then `units` and `unitPrice`
	// cannot be used.
	Amount float64 `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate time.Time `json:"billDate" format:"date"`
	// The ID of the Bill created for this Charge.
	BillID string `json:"billId"`
	// The unique short code of the Charge.
	Code string `json:"code"`
	// The ID of a Contract on the Account that the Charge has been added to.
	ContractID string `json:"contractId"`
	// The unique identifier (UUID) of the user who created the Charge.
	CreatedBy string `json:"createdBy"`
	// Charge currency.
	Currency string `json:"currency"`
	// The description added to the Bill line item for the Charge.
	Description string `json:"description"`
	// The date and time (_in ISO-8601 format_) when the Charge was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO 8601 format_) when the Charge was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	EntityID string `json:"entityId"`
	// The entity type the Charge has been created for.
	EntityType ChargeDeleteResponseEntityType `json:"entityType"`
	// The unique identifier (UUID) of the user who last modified the Charge.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The line item type used for billing a Charge.
	LineItemType ChargeDeleteResponseLineItemType `json:"lineItemType"`
	// Name of the Charge. Added to the Bill line item description for Charge.
	Name string `json:"name"`
	// Information about the Charge for accounting purposes, such as the reason it was
	// created. This information will not be added to the created Bill line item for
	// the Charge.
	Notes string `json:"notes"`
	// The ID of the Balance Charge Schedule that created the Charge.
	ScheduleID string `json:"scheduleId"`
	// The service period end date (_in ISO-8601 format_) for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge .
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit Price for the Charge. Provided together with `units`:
	//
	// - Null if the Charge was created with `amount` only.
	// - If `units` and `unitPrice` are provided, `amount` cannot be used.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units of the Charge. Provided together with `unitPrice`. If `units`
	// and `unitPrice` are provided, `amount` cannot be used.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                    `json:"version"`
	JSON    chargeDeleteResponseJSON `json:"-"`
}

// chargeDeleteResponseJSON contains the JSON metadata for the struct
// [ChargeDeleteResponse]
type chargeDeleteResponseJSON struct {
	ID                     apijson.Field
	AccountID              apijson.Field
	AccountingProductID    apijson.Field
	Amount                 apijson.Field
	BillDate               apijson.Field
	BillID                 apijson.Field
	Code                   apijson.Field
	ContractID             apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	Description            apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	EntityID               apijson.Field
	EntityType             apijson.Field
	LastModifiedBy         apijson.Field
	LineItemType           apijson.Field
	Name                   apijson.Field
	Notes                  apijson.Field
	ScheduleID             apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ChargeDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r chargeDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The entity type the Charge has been created for.
type ChargeDeleteResponseEntityType string

const (
	ChargeDeleteResponseEntityTypeAdHoc   ChargeDeleteResponseEntityType = "AD_HOC"
	ChargeDeleteResponseEntityTypeBalance ChargeDeleteResponseEntityType = "BALANCE"
)

func (r ChargeDeleteResponseEntityType) IsKnown() bool {
	switch r {
	case ChargeDeleteResponseEntityTypeAdHoc, ChargeDeleteResponseEntityTypeBalance:
		return true
	}
	return false
}

// The line item type used for billing a Charge.
type ChargeDeleteResponseLineItemType string

const (
	ChargeDeleteResponseLineItemTypeBalanceFee ChargeDeleteResponseLineItemType = "BALANCE_FEE"
	ChargeDeleteResponseLineItemTypeAdHoc      ChargeDeleteResponseLineItemType = "AD_HOC"
)

func (r ChargeDeleteResponseLineItemType) IsKnown() bool {
	switch r {
	case ChargeDeleteResponseLineItemTypeBalanceFee, ChargeDeleteResponseLineItemTypeAdHoc:
		return true
	}
	return false
}

type ChargeNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The ID of the Account the Charge is being created for.
	AccountID param.Field[string] `json:"accountId,required"`
	// Unique short code for the Charge.
	Code param.Field[string] `json:"code,required"`
	// Charge currency.
	Currency param.Field[string] `json:"currency,required"`
	// The entity type the Charge has been created for.
	EntityType param.Field[ChargeNewParamsEntityType] `json:"entityType,required"`
	// Available line item types that can be used for billing a Charge.
	LineItemType param.Field[ChargeNewParamsLineItemType] `json:"lineItemType,required"`
	// Name of the Charge. Added to the Bill line item description for this Charge.
	Name param.Field[string] `json:"name,required"`
	// The service period end date (_in ISO-8601 format_)for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Amount of the Charge. If `amount` is provided, then `units` and `unitPrice` must
	// be omitted.
	Amount param.Field[float64] `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate param.Field[string] `json:"billDate"`
	// The ID of a Contract on the Account that the Charge will be added to.
	ContractID param.Field[string] `json:"contractId"`
	// The description added to the Bill line item for the Charge.
	Description param.Field[string] `json:"description"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	//
	// **NOTE:** If `entityType` is `BALANCE`, you must provide the `entityId` of the
	// Balance the Charge is for.
	EntityID param.Field[string] `json:"entityId"`
	// Used to enter information about the Charge for accounting purposes, such as the
	// reason it was created. This information will not be added to a Bill line item
	// for the Charge.
	Notes param.Field[string] `json:"notes"`
	// Unit price. If `amount` is omitted, then provide together with `units`. When
	// `amount` is provided, `unitPrice` must be omitted.
	UnitPrice param.Field[float64] `json:"unitPrice"`
	// Number of units of the Charge. If `amount` is omitted, then provide together
	// with `unitPrice`. When `amount` is provided, `units` must be omitted.
	Units param.Field[float64] `json:"units"`
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

func (r ChargeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity type the Charge has been created for.
type ChargeNewParamsEntityType string

const (
	ChargeNewParamsEntityTypeAdHoc   ChargeNewParamsEntityType = "AD_HOC"
	ChargeNewParamsEntityTypeBalance ChargeNewParamsEntityType = "BALANCE"
)

func (r ChargeNewParamsEntityType) IsKnown() bool {
	switch r {
	case ChargeNewParamsEntityTypeAdHoc, ChargeNewParamsEntityTypeBalance:
		return true
	}
	return false
}

// Available line item types that can be used for billing a Charge.
type ChargeNewParamsLineItemType string

const (
	ChargeNewParamsLineItemTypeBalanceFee ChargeNewParamsLineItemType = "BALANCE_FEE"
	ChargeNewParamsLineItemTypeAdHoc      ChargeNewParamsLineItemType = "AD_HOC"
)

func (r ChargeNewParamsLineItemType) IsKnown() bool {
	switch r {
	case ChargeNewParamsLineItemTypeBalanceFee, ChargeNewParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type ChargeGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type ChargeUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The ID of the Account the Charge is being created for.
	AccountID param.Field[string] `json:"accountId,required"`
	// Unique short code for the Charge.
	Code param.Field[string] `json:"code,required"`
	// Charge currency.
	Currency param.Field[string] `json:"currency,required"`
	// The entity type the Charge has been created for.
	EntityType param.Field[ChargeUpdateParamsEntityType] `json:"entityType,required"`
	// Available line item types that can be used for billing a Charge.
	LineItemType param.Field[ChargeUpdateParamsLineItemType] `json:"lineItemType,required"`
	// Name of the Charge. Added to the Bill line item description for this Charge.
	Name param.Field[string] `json:"name,required"`
	// The service period end date (_in ISO-8601 format_)for the Charge.
	//
	// **NOTE:** End date is exclusive.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) for the Charge.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// The Accounting Product ID assigned to the Charge.
	AccountingProductID param.Field[string] `json:"accountingProductId"`
	// Amount of the Charge. If `amount` is provided, then `units` and `unitPrice` must
	// be omitted.
	Amount param.Field[float64] `json:"amount"`
	// The date when the Charge will be added to a Bill.
	BillDate param.Field[string] `json:"billDate"`
	// The ID of a Contract on the Account that the Charge will be added to.
	ContractID param.Field[string] `json:"contractId"`
	// The description added to the Bill line item for the Charge.
	Description param.Field[string] `json:"description"`
	// The ID of the Charge linked entity. For example, the ID of an Account Balance if
	// a Balance Charge.
	//
	// **NOTE:** If `entityType` is `BALANCE`, you must provide the `entityId` of the
	// Balance the Charge is for.
	EntityID param.Field[string] `json:"entityId"`
	// Used to enter information about the Charge for accounting purposes, such as the
	// reason it was created. This information will not be added to a Bill line item
	// for the Charge.
	Notes param.Field[string] `json:"notes"`
	// Unit price. If `amount` is omitted, then provide together with `units`. When
	// `amount` is provided, `unitPrice` must be omitted.
	UnitPrice param.Field[float64] `json:"unitPrice"`
	// Number of units of the Charge. If `amount` is omitted, then provide together
	// with `unitPrice`. When `amount` is provided, `units` must be omitted.
	Units param.Field[float64] `json:"units"`
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

func (r ChargeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The entity type the Charge has been created for.
type ChargeUpdateParamsEntityType string

const (
	ChargeUpdateParamsEntityTypeAdHoc   ChargeUpdateParamsEntityType = "AD_HOC"
	ChargeUpdateParamsEntityTypeBalance ChargeUpdateParamsEntityType = "BALANCE"
)

func (r ChargeUpdateParamsEntityType) IsKnown() bool {
	switch r {
	case ChargeUpdateParamsEntityTypeAdHoc, ChargeUpdateParamsEntityTypeBalance:
		return true
	}
	return false
}

// Available line item types that can be used for billing a Charge.
type ChargeUpdateParamsLineItemType string

const (
	ChargeUpdateParamsLineItemTypeBalanceFee ChargeUpdateParamsLineItemType = "BALANCE_FEE"
	ChargeUpdateParamsLineItemTypeAdHoc      ChargeUpdateParamsLineItemType = "AD_HOC"
)

func (r ChargeUpdateParamsLineItemType) IsKnown() bool {
	switch r {
	case ChargeUpdateParamsLineItemTypeBalanceFee, ChargeUpdateParamsLineItemTypeAdHoc:
		return true
	}
	return false
}

type ChargeListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// List Charge items for the Account UUID
	AccountID param.Field[string] `query:"accountId"`
	// List Charge items for the Bill Date
	BillDate param.Field[time.Time] `query:"billDate" format:"date"`
	// List Charge items for the Entity UUID
	EntityID param.Field[string] `query:"entityId"`
	// List Charge items for the EntityType
	EntityType param.Field[ChargeListParamsEntityType] `query:"entityType"`
	// List of Charge UUIDs to retrieve
	IDs param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Charges to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// List Charge items for the Schedule UUID
	ScheduleID param.Field[string] `query:"scheduleId"`
}

// URLQuery serializes [ChargeListParams]'s query parameters as `url.Values`.
func (r ChargeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// List Charge items for the EntityType
type ChargeListParamsEntityType string

const (
	ChargeListParamsEntityTypeAdHoc   ChargeListParamsEntityType = "AD_HOC"
	ChargeListParamsEntityTypeBalance ChargeListParamsEntityType = "BALANCE"
)

func (r ChargeListParamsEntityType) IsKnown() bool {
	switch r {
	case ChargeListParamsEntityTypeAdHoc, ChargeListParamsEntityTypeBalance:
		return true
	}
	return false
}

type ChargeDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
