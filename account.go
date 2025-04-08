// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
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

// AccountService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Create a new Account within the Organization.
func (r *AccountService) New(ctx context.Context, params AccountNewParams, opts ...option.RequestOption) (res *AccountResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the Account with the given Account UUID.
func (r *AccountService) Get(ctx context.Context, id string, query AccountGetParams, opts ...option.RequestOption) (res *AccountResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts/%s", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Account with the given Account UUID.
//
// **Note:** If you have created Custom Fields for an Account, when you use this
// endpoint to update the Account, use the `customFields` parameter to preserve
// those Custom Fields. If you omit them from the update request, they will be
// lost.
func (r *AccountService) Update(ctx context.Context, id string, params AccountUpdateParams, opts ...option.RequestOption) (res *AccountResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts/%s", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Accounts that can be filtered by Account ID or Account Code.
func (r *AccountService) List(ctx context.Context, params AccountListParams, opts ...option.RequestOption) (res *pagination.Cursor[AccountResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts", params.OrgID)
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

// Retrieve a list of Accounts that can be filtered by Account ID or Account Code.
func (r *AccountService) ListAutoPaging(ctx context.Context, params AccountListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[AccountResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, params, opts...))
}

// Delete the Account with the given UUID. This may fail if there are any
// AccountPlans that reference the Account being deleted.
func (r *AccountService) Delete(ctx context.Context, id string, body AccountDeleteParams, opts ...option.RequestOption) (res *AccountResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts/%s", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Apply the specified end-date to billing entities associated with an Account.
//
// **NOTE:**
//
//   - When you successfully end-date billing entities, the version number of each
//     entity is incremented.
func (r *AccountService) EndDateBillingEntities(ctx context.Context, id string, params AccountEndDateBillingEntitiesParams, opts ...option.RequestOption) (res *AccountEndDateBillingEntitiesResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts/%s/enddatebillingentities", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of Accounts that are children of the specified Account.
func (r *AccountService) GetChildren(ctx context.Context, id string, params AccountGetChildrenParams, opts ...option.RequestOption) (res *AccountResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts/%s/children", params.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Search for Account entities.
//
// This endpoint executes a search query for Accounts based on the user specified
// search criteria. The search query is customizable, allowing for complex nested
// conditions and sorting. The returned list of Accounts can be paginated for
// easier management.
func (r *AccountService) Search(ctx context.Context, params AccountSearchParams, opts ...option.RequestOption) (res *AccountSearchResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/accounts/search", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type AccountResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
	// Contact address.
	Address Address `json:"address"`
	// Specify whether to auto-generate statements once Bills are approved or locked.
	//
	// - **None**. Statements will not be auto-generated.
	// - **JSON**. Statements are auto-generated in JSON format.
	// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
	AutoGenerateStatementMode AccountResponseAutoGenerateStatementMode `json:"autoGenerateStatementMode"`
	// Defines first bill date for Account Bills. For example, if the Plan attached to
	// the Account is set for monthly billing frequency and you set the first bill date
	// to be January 1st, Bills are created every month starting on that date.
	//
	// Optional attribute - if not defined, then first bill date is determined by the
	// Epoch settings at Organizational level.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Code of the Account. This is a unique short code used for the Account.
	Code string `json:"code"`
	// Configuration data for the Account
	ConfigData map[string]interface{} `json:"configData"`
	// The ID of the user who created the account.
	CreatedBy string `json:"createdBy"`
	// The order in which any Prepayment or Balance amounts on the Account are to be
	// drawn-down against for billing. Four options:
	//
	//   - `"PREPAYMENT","BALANCE"`. Draw-down against Prepayment credit before Balance
	//     credit.
	//   - `"BALANCE","PREPAYMENT"`. Draw-down against Balance credit before Prepayment
	//     credit.
	//   - `"PREPAYMENT"`. Only draw-down against Prepayment credit.
	//   - `"BALANCE"`. Only draw-down against Balance credit.
	CreditApplicationOrder []AccountResponseCreditApplicationOrder `json:"creditApplicationOrder"`
	// Account level billing currency, such as USD or GBP. Optional attribute:
	//
	//   - If you define an Account currency, this will be used for bills.
	//   - If you do not define a currency, the billing currency defined at
	//     Organizational will be used.
	//
	// **Note:** If you've attached a Plan to the Account that uses a different
	// currency to the billing currency, then you must add the relevant currency
	// conversion rate at Organization level to ensure the billing process can convert
	// line items calculated using the Plan currency into the selected billing
	// currency. If you don't add these conversion rates, then bills will fail for the
	// Account.
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
	CustomFields map[string]AccountResponseCustomFieldsUnion `json:"customFields"`
	// The number of days after the Bill generation date shown on Bills as the due
	// date.
	DaysBeforeBillDue int64 `json:"daysBeforeBillDue"`
	// The DateTime when the Account was created _(in ISO 8601 format)_.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime when the Account was last modified _(in ISO 8601 format)_.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// Contact email for the Account.
	EmailAddress string `json:"emailAddress"`
	// The ID of the user who last modified the Account.
	LastModifiedBy string `json:"lastModifiedBy"`
	// Name of the Account.
	Name string `json:"name"`
	// Parent Account ID, or null if this account does not have a parent.
	ParentAccountID string `json:"parentAccountId"`
	// Purchase Order Number of the Account.
	//
	// Optional attribute - allows you to set a purchase order number that comes
	// through into invoicing. For example, your financial systems might require this
	// as a reference for clearing payments.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
	// The UUID of the statement definition used when Bill statements are generated for
	// the Account. If no statement definition is specified for the Account, the
	// statement definition specified at Organizational level is used.
	//
	// Bill statements can be used as informative backing sheets to invoices. Based on
	// the usage breakdown defined in the statement definition, generated statements
	// give a breakdown of usage charges on Account Bills, which helps customers better
	// understand usage charges incurred over the billing period.
	//
	// See
	// [Working with Bill Statements](https://www.m3ter.com/docs/guides/running-viewing-and-managing-bills/working-with-bill-statements)
	// in the m3ter documentation for more details.
	StatementDefinitionID string              `json:"statementDefinitionId"`
	JSON                  accountResponseJSON `json:"-"`
}

// accountResponseJSON contains the JSON metadata for the struct [AccountResponse]
type accountResponseJSON struct {
	ID                        apijson.Field
	Version                   apijson.Field
	Address                   apijson.Field
	AutoGenerateStatementMode apijson.Field
	BillEpoch                 apijson.Field
	Code                      apijson.Field
	ConfigData                apijson.Field
	CreatedBy                 apijson.Field
	CreditApplicationOrder    apijson.Field
	Currency                  apijson.Field
	CustomFields              apijson.Field
	DaysBeforeBillDue         apijson.Field
	DtCreated                 apijson.Field
	DtLastModified            apijson.Field
	EmailAddress              apijson.Field
	LastModifiedBy            apijson.Field
	Name                      apijson.Field
	ParentAccountID           apijson.Field
	PurchaseOrderNumber       apijson.Field
	StatementDefinitionID     apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountResponseJSON) RawJSON() string {
	return r.raw
}

// Specify whether to auto-generate statements once Bills are approved or locked.
//
// - **None**. Statements will not be auto-generated.
// - **JSON**. Statements are auto-generated in JSON format.
// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
type AccountResponseAutoGenerateStatementMode string

const (
	AccountResponseAutoGenerateStatementModeNone       AccountResponseAutoGenerateStatementMode = "NONE"
	AccountResponseAutoGenerateStatementModeJson       AccountResponseAutoGenerateStatementMode = "JSON"
	AccountResponseAutoGenerateStatementModeJsonAndCsv AccountResponseAutoGenerateStatementMode = "JSON_AND_CSV"
)

func (r AccountResponseAutoGenerateStatementMode) IsKnown() bool {
	switch r {
	case AccountResponseAutoGenerateStatementModeNone, AccountResponseAutoGenerateStatementModeJson, AccountResponseAutoGenerateStatementModeJsonAndCsv:
		return true
	}
	return false
}

type AccountResponseCreditApplicationOrder string

const (
	AccountResponseCreditApplicationOrderPrepayment AccountResponseCreditApplicationOrder = "PREPAYMENT"
	AccountResponseCreditApplicationOrderBalance    AccountResponseCreditApplicationOrder = "BALANCE"
)

func (r AccountResponseCreditApplicationOrder) IsKnown() bool {
	switch r {
	case AccountResponseCreditApplicationOrderPrepayment, AccountResponseCreditApplicationOrderBalance:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type AccountResponseCustomFieldsUnion interface {
	ImplementsAccountResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountResponseCustomFieldsUnion)(nil)).Elem(),
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

// Contact address.
type Address struct {
	AddressLine1 string      `json:"addressLine1"`
	AddressLine2 string      `json:"addressLine2"`
	AddressLine3 string      `json:"addressLine3"`
	AddressLine4 string      `json:"addressLine4"`
	Country      string      `json:"country"`
	Locality     string      `json:"locality"`
	PostCode     string      `json:"postCode"`
	Region       string      `json:"region"`
	JSON         addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	AddressLine1 apijson.Field
	AddressLine2 apijson.Field
	AddressLine3 apijson.Field
	AddressLine4 apijson.Field
	Country      apijson.Field
	Locality     apijson.Field
	PostCode     apijson.Field
	Region       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

// Contact address.
type AddressParam struct {
	AddressLine1 param.Field[string] `json:"addressLine1"`
	AddressLine2 param.Field[string] `json:"addressLine2"`
	AddressLine3 param.Field[string] `json:"addressLine3"`
	AddressLine4 param.Field[string] `json:"addressLine4"`
	Country      param.Field[string] `json:"country"`
	Locality     param.Field[string] `json:"locality"`
	PostCode     param.Field[string] `json:"postCode"`
	Region       param.Field[string] `json:"region"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEndDateBillingEntitiesResponse struct {
	// A dictionary with keys as identifiers of billing entities and values as lists
	// containing details of the entities for which the update failed.
	FailedEntities AccountEndDateBillingEntitiesResponseFailedEntities `json:"failedEntities"`
	// A message indicating the status of the operation.
	StatusMessage string `json:"statusMessage"`
	// A dictionary with keys as identifiers of billing entities and values as lists
	// containing details of the updated entities.
	UpdatedEntities AccountEndDateBillingEntitiesResponseUpdatedEntities `json:"updatedEntities"`
	JSON            accountEndDateBillingEntitiesResponseJSON            `json:"-"`
}

// accountEndDateBillingEntitiesResponseJSON contains the JSON metadata for the
// struct [AccountEndDateBillingEntitiesResponse]
type accountEndDateBillingEntitiesResponseJSON struct {
	FailedEntities  apijson.Field
	StatusMessage   apijson.Field
	UpdatedEntities apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountEndDateBillingEntitiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEndDateBillingEntitiesResponseJSON) RawJSON() string {
	return r.raw
}

// A dictionary with keys as identifiers of billing entities and values as lists
// containing details of the entities for which the update failed.
type AccountEndDateBillingEntitiesResponseFailedEntities struct {
	Accountplan     shared.SetString                                        `json:"ACCOUNTPLAN"`
	Contract        shared.SetString                                        `json:"CONTRACT"`
	CounterPricings shared.SetString                                        `json:"COUNTER_PRICINGS"`
	Prepayment      shared.SetString                                        `json:"PREPAYMENT"`
	Pricings        shared.SetString                                        `json:"PRICINGS"`
	JSON            accountEndDateBillingEntitiesResponseFailedEntitiesJSON `json:"-"`
}

// accountEndDateBillingEntitiesResponseFailedEntitiesJSON contains the JSON
// metadata for the struct [AccountEndDateBillingEntitiesResponseFailedEntities]
type accountEndDateBillingEntitiesResponseFailedEntitiesJSON struct {
	Accountplan     apijson.Field
	Contract        apijson.Field
	CounterPricings apijson.Field
	Prepayment      apijson.Field
	Pricings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountEndDateBillingEntitiesResponseFailedEntities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEndDateBillingEntitiesResponseFailedEntitiesJSON) RawJSON() string {
	return r.raw
}

// A dictionary with keys as identifiers of billing entities and values as lists
// containing details of the updated entities.
type AccountEndDateBillingEntitiesResponseUpdatedEntities struct {
	Accountplan     shared.SetString                                         `json:"ACCOUNTPLAN"`
	Contract        shared.SetString                                         `json:"CONTRACT"`
	CounterPricings shared.SetString                                         `json:"COUNTER_PRICINGS"`
	Prepayment      shared.SetString                                         `json:"PREPAYMENT"`
	Pricings        shared.SetString                                         `json:"PRICINGS"`
	JSON            accountEndDateBillingEntitiesResponseUpdatedEntitiesJSON `json:"-"`
}

// accountEndDateBillingEntitiesResponseUpdatedEntitiesJSON contains the JSON
// metadata for the struct [AccountEndDateBillingEntitiesResponseUpdatedEntities]
type accountEndDateBillingEntitiesResponseUpdatedEntitiesJSON struct {
	Accountplan     apijson.Field
	Contract        apijson.Field
	CounterPricings apijson.Field
	Prepayment      apijson.Field
	Pricings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountEndDateBillingEntitiesResponseUpdatedEntities) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEndDateBillingEntitiesResponseUpdatedEntitiesJSON) RawJSON() string {
	return r.raw
}

type AccountSearchResponse struct {
	Data      []AccountResponse         `json:"data"`
	NextToken string                    `json:"nextToken"`
	JSON      accountSearchResponseJSON `json:"-"`
}

// accountSearchResponseJSON contains the JSON metadata for the struct
// [AccountSearchResponse]
type accountSearchResponseJSON struct {
	Data        apijson.Field
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSearchResponseJSON) RawJSON() string {
	return r.raw
}

type AccountNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Code of the Account. This is a unique short code used for the Account.
	Code param.Field[string] `json:"code,required"`
	// Contact email for the Account.
	EmailAddress param.Field[string] `json:"emailAddress,required" format:"email"`
	// Name of the Account.
	Name param.Field[string] `json:"name,required"`
	// Contact address.
	Address param.Field[AddressParam] `json:"address"`
	// Specify whether to auto-generate statements once Bills are approved or locked.
	//
	// - **None**. Statements will not be auto-generated.
	// - **JSON**. Statements are auto-generated in JSON format.
	// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
	AutoGenerateStatementMode param.Field[AccountNewParamsAutoGenerateStatementMode] `json:"autoGenerateStatementMode"`
	// Optional setting to define a _billing cycle date_, which sets the date of the
	// first Bill and acts as a reference for when in the applied billing frequency
	// period subsequent bills are created:
	//
	//   - For example, if you attach a Plan to an Account where the Plan is configured
	//     for monthly billing frequency and you've defined the period the Plan will
	//     apply to the Account to be from January 1st, 2022 until January 1st, 2023. You
	//     then set a `billEpoch` date of February 15th, 2022. The first Bill will be
	//     created for the Account on February 15th, and subsequent Bills created on the
	//     15th of the months following for the remainder of the billing period - March
	//     15th, April 15th, and so on.
	//   - If not defined, then the relevant Epoch date set for the billing frequency
	//     period at Organization level will be used instead.
	//   - The date is in ISO-8601 format.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
	// Configuration data for the Account Supported settings:
	//
	// - SendBillsToThirdParties ("true"/"false")
	ConfigData param.Field[map[string]interface{}] `json:"configData"`
	// Define the order in which any Prepayment or Balance amounts on the Account are
	// to be drawn-down against for billing. Four options:
	//
	//   - `"PREPAYMENT","BALANCE"`. Draw-down against Prepayment credit before Balance
	//     credit.
	//   - `"BALANCE","PREPAYMENT"`. Draw-down against Balance credit before Prepayment
	//     credit.
	//   - `"PREPAYMENT"`. Only draw-down against Prepayment credit.
	//   - `"BALANCE"`. Only draw-down against Balance credit.
	//
	// **NOTES:**
	//
	//   - Any setting you define here overrides the setting for credit application order
	//     at Organization level.
	//   - If the Account belongs to a Parent/Child Account hierarchy, then the
	//     `creditApplicationOrder` settings are not available, and the draw-down order
	//     defaults always to Prepayment then Balance order.
	CreditApplicationOrder param.Field[[]AccountNewParamsCreditApplicationOrder] `json:"creditApplicationOrder"`
	// Account level billing currency, such as USD or GBP. Optional attribute:
	//
	//   - If you define an Account currency, this will be used for bills.
	//   - If you do not define a currency, the billing currency defined at
	//     Organizational level will be used.
	//
	// **Note:** If you've attached a Plan to the Account that uses a different
	// currency to the billing currency, then you must add the relevant currency
	// conversion rate at Organization level to ensure the billing process can convert
	// line items calculated using the Plan currency into the selected billing
	// currency. If you don't add these conversion rates, then bills will fail for the
	// Account.
	Currency param.Field[string] `json:"currency"`
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
	CustomFields param.Field[map[string]AccountNewParamsCustomFieldsUnion] `json:"customFields"`
	// Enter the number of days after the Bill generation date that you want to show on
	// Bills as the due date.
	//
	// **Note:** If you define `daysBeforeBillDue` at individual Account level, this
	// will take precedence over any `daysBeforeBillDue` setting defined at
	// Organization level.
	DaysBeforeBillDue param.Field[int64] `json:"daysBeforeBillDue"`
	// Parent Account ID, or null if this Account does not have a parent.
	ParentAccountID param.Field[string] `json:"parentAccountId"`
	// Purchase Order Number of the Account.
	//
	// Optional attribute - allows you to set a purchase order number that comes
	// through into invoicing. For example, your financial systems might require this
	// as a reference for clearing payments.
	PurchaseOrderNumber param.Field[string] `json:"purchaseOrderNumber"`
	// The UUID of the statement definition used when Bill statements are generated for
	// the Account. If no statement definition is specified for the Account, the
	// statement definition specified at Organizational level is used.
	//
	// Bill statements can be used as informative backing sheets to invoices. Based on
	// the usage breakdown defined in the statement definition, generated statements
	// give a breakdown of usage charges on Account Bills, which helps customers better
	// understand usage charges incurred over the billing period.
	//
	// See
	// [Working with Bill Statements](https://www.m3ter.com/docs/guides/running-viewing-and-managing-bills/working-with-bill-statements)
	// in the m3ter documentation for more details.
	StatementDefinitionID param.Field[string] `json:"statementDefinitionId"`
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

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify whether to auto-generate statements once Bills are approved or locked.
//
// - **None**. Statements will not be auto-generated.
// - **JSON**. Statements are auto-generated in JSON format.
// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
type AccountNewParamsAutoGenerateStatementMode string

const (
	AccountNewParamsAutoGenerateStatementModeNone       AccountNewParamsAutoGenerateStatementMode = "NONE"
	AccountNewParamsAutoGenerateStatementModeJson       AccountNewParamsAutoGenerateStatementMode = "JSON"
	AccountNewParamsAutoGenerateStatementModeJsonAndCsv AccountNewParamsAutoGenerateStatementMode = "JSON_AND_CSV"
)

func (r AccountNewParamsAutoGenerateStatementMode) IsKnown() bool {
	switch r {
	case AccountNewParamsAutoGenerateStatementModeNone, AccountNewParamsAutoGenerateStatementModeJson, AccountNewParamsAutoGenerateStatementModeJsonAndCsv:
		return true
	}
	return false
}

type AccountNewParamsCreditApplicationOrder string

const (
	AccountNewParamsCreditApplicationOrderPrepayment AccountNewParamsCreditApplicationOrder = "PREPAYMENT"
	AccountNewParamsCreditApplicationOrderBalance    AccountNewParamsCreditApplicationOrder = "BALANCE"
)

func (r AccountNewParamsCreditApplicationOrder) IsKnown() bool {
	switch r {
	case AccountNewParamsCreditApplicationOrderPrepayment, AccountNewParamsCreditApplicationOrderBalance:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type AccountNewParamsCustomFieldsUnion interface {
	ImplementsAccountNewParamsCustomFieldsUnion()
}

type AccountGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type AccountUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Code of the Account. This is a unique short code used for the Account.
	Code param.Field[string] `json:"code,required"`
	// Contact email for the Account.
	EmailAddress param.Field[string] `json:"emailAddress,required" format:"email"`
	// Name of the Account.
	Name param.Field[string] `json:"name,required"`
	// Contact address.
	Address param.Field[AddressParam] `json:"address"`
	// Specify whether to auto-generate statements once Bills are approved or locked.
	//
	// - **None**. Statements will not be auto-generated.
	// - **JSON**. Statements are auto-generated in JSON format.
	// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
	AutoGenerateStatementMode param.Field[AccountUpdateParamsAutoGenerateStatementMode] `json:"autoGenerateStatementMode"`
	// Optional setting to define a _billing cycle date_, which sets the date of the
	// first Bill and acts as a reference for when in the applied billing frequency
	// period subsequent bills are created:
	//
	//   - For example, if you attach a Plan to an Account where the Plan is configured
	//     for monthly billing frequency and you've defined the period the Plan will
	//     apply to the Account to be from January 1st, 2022 until January 1st, 2023. You
	//     then set a `billEpoch` date of February 15th, 2022. The first Bill will be
	//     created for the Account on February 15th, and subsequent Bills created on the
	//     15th of the months following for the remainder of the billing period - March
	//     15th, April 15th, and so on.
	//   - If not defined, then the relevant Epoch date set for the billing frequency
	//     period at Organization level will be used instead.
	//   - The date is in ISO-8601 format.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
	// Configuration data for the Account Supported settings:
	//
	// - SendBillsToThirdParties ("true"/"false")
	ConfigData param.Field[map[string]interface{}] `json:"configData"`
	// Define the order in which any Prepayment or Balance amounts on the Account are
	// to be drawn-down against for billing. Four options:
	//
	//   - `"PREPAYMENT","BALANCE"`. Draw-down against Prepayment credit before Balance
	//     credit.
	//   - `"BALANCE","PREPAYMENT"`. Draw-down against Balance credit before Prepayment
	//     credit.
	//   - `"PREPAYMENT"`. Only draw-down against Prepayment credit.
	//   - `"BALANCE"`. Only draw-down against Balance credit.
	//
	// **NOTES:**
	//
	//   - Any setting you define here overrides the setting for credit application order
	//     at Organization level.
	//   - If the Account belongs to a Parent/Child Account hierarchy, then the
	//     `creditApplicationOrder` settings are not available, and the draw-down order
	//     defaults always to Prepayment then Balance order.
	CreditApplicationOrder param.Field[[]AccountUpdateParamsCreditApplicationOrder] `json:"creditApplicationOrder"`
	// Account level billing currency, such as USD or GBP. Optional attribute:
	//
	//   - If you define an Account currency, this will be used for bills.
	//   - If you do not define a currency, the billing currency defined at
	//     Organizational level will be used.
	//
	// **Note:** If you've attached a Plan to the Account that uses a different
	// currency to the billing currency, then you must add the relevant currency
	// conversion rate at Organization level to ensure the billing process can convert
	// line items calculated using the Plan currency into the selected billing
	// currency. If you don't add these conversion rates, then bills will fail for the
	// Account.
	Currency param.Field[string] `json:"currency"`
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
	CustomFields param.Field[map[string]AccountUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// Enter the number of days after the Bill generation date that you want to show on
	// Bills as the due date.
	//
	// **Note:** If you define `daysBeforeBillDue` at individual Account level, this
	// will take precedence over any `daysBeforeBillDue` setting defined at
	// Organization level.
	DaysBeforeBillDue param.Field[int64] `json:"daysBeforeBillDue"`
	// Parent Account ID, or null if this Account does not have a parent.
	ParentAccountID param.Field[string] `json:"parentAccountId"`
	// Purchase Order Number of the Account.
	//
	// Optional attribute - allows you to set a purchase order number that comes
	// through into invoicing. For example, your financial systems might require this
	// as a reference for clearing payments.
	PurchaseOrderNumber param.Field[string] `json:"purchaseOrderNumber"`
	// The UUID of the statement definition used when Bill statements are generated for
	// the Account. If no statement definition is specified for the Account, the
	// statement definition specified at Organizational level is used.
	//
	// Bill statements can be used as informative backing sheets to invoices. Based on
	// the usage breakdown defined in the statement definition, generated statements
	// give a breakdown of usage charges on Account Bills, which helps customers better
	// understand usage charges incurred over the billing period.
	//
	// See
	// [Working with Bill Statements](https://www.m3ter.com/docs/guides/running-viewing-and-managing-bills/working-with-bill-statements)
	// in the m3ter documentation for more details.
	StatementDefinitionID param.Field[string] `json:"statementDefinitionId"`
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

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify whether to auto-generate statements once Bills are approved or locked.
//
// - **None**. Statements will not be auto-generated.
// - **JSON**. Statements are auto-generated in JSON format.
// - **JSON and CSV**. Statements are auto-generated in both JSON and CSV formats.
type AccountUpdateParamsAutoGenerateStatementMode string

const (
	AccountUpdateParamsAutoGenerateStatementModeNone       AccountUpdateParamsAutoGenerateStatementMode = "NONE"
	AccountUpdateParamsAutoGenerateStatementModeJson       AccountUpdateParamsAutoGenerateStatementMode = "JSON"
	AccountUpdateParamsAutoGenerateStatementModeJsonAndCsv AccountUpdateParamsAutoGenerateStatementMode = "JSON_AND_CSV"
)

func (r AccountUpdateParamsAutoGenerateStatementMode) IsKnown() bool {
	switch r {
	case AccountUpdateParamsAutoGenerateStatementModeNone, AccountUpdateParamsAutoGenerateStatementModeJson, AccountUpdateParamsAutoGenerateStatementModeJsonAndCsv:
		return true
	}
	return false
}

type AccountUpdateParamsCreditApplicationOrder string

const (
	AccountUpdateParamsCreditApplicationOrderPrepayment AccountUpdateParamsCreditApplicationOrder = "PREPAYMENT"
	AccountUpdateParamsCreditApplicationOrderBalance    AccountUpdateParamsCreditApplicationOrder = "BALANCE"
)

func (r AccountUpdateParamsCreditApplicationOrder) IsKnown() bool {
	switch r {
	case AccountUpdateParamsCreditApplicationOrderPrepayment, AccountUpdateParamsCreditApplicationOrderBalance:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type AccountUpdateParamsCustomFieldsUnion interface {
	ImplementsAccountUpdateParamsCustomFieldsUnion()
}

type AccountListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// List of Account Codes to retrieve. These are unique short codes for each
	// Account.
	Codes param.Field[[]string] `query:"codes"`
	// List of Account IDs to retrieve.
	IDs param.Field[[]string] `query:"ids"`
	// `nextToken` for multi-page retrievals.
	NextToken param.Field[string] `query:"nextToken"`
	// Number of accounts to retrieve per page.
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type AccountEndDateBillingEntitiesParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Defines which billing entities associated with the Account will have the
	// specified end-date applied. For example, if you want the specified end-date to
	// be applied to all Prepayments/Commitments created for the Account use
	// `"PREPAYMENT"`.
	BillingEntities param.Field[[]AccountEndDateBillingEntitiesParamsBillingEntity] `json:"billingEntities,required"`
	// The end date and time applied to the specified billing entities _(in ISO 8601
	// format)_.
	EndDate param.Field[time.Time] `json:"endDate,required" format:"date-time"`
	// A Boolean TRUE/FALSE flag. For Parent Accounts, set to TRUE if you want the
	// specified end-date to be applied to any billing entities associated with Child
	// Accounts. _(Optional)_
	ApplyToChildren param.Field[bool] `json:"applyToChildren"`
}

func (r AccountEndDateBillingEntitiesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEndDateBillingEntitiesParamsBillingEntity string

const (
	AccountEndDateBillingEntitiesParamsBillingEntityContract        AccountEndDateBillingEntitiesParamsBillingEntity = "CONTRACT"
	AccountEndDateBillingEntitiesParamsBillingEntityAccountplan     AccountEndDateBillingEntitiesParamsBillingEntity = "ACCOUNTPLAN"
	AccountEndDateBillingEntitiesParamsBillingEntityPrepayment      AccountEndDateBillingEntitiesParamsBillingEntity = "PREPAYMENT"
	AccountEndDateBillingEntitiesParamsBillingEntityPricings        AccountEndDateBillingEntitiesParamsBillingEntity = "PRICINGS"
	AccountEndDateBillingEntitiesParamsBillingEntityCounterPricings AccountEndDateBillingEntitiesParamsBillingEntity = "COUNTER_PRICINGS"
)

func (r AccountEndDateBillingEntitiesParamsBillingEntity) IsKnown() bool {
	switch r {
	case AccountEndDateBillingEntitiesParamsBillingEntityContract, AccountEndDateBillingEntitiesParamsBillingEntityAccountplan, AccountEndDateBillingEntitiesParamsBillingEntityPrepayment, AccountEndDateBillingEntitiesParamsBillingEntityPricings, AccountEndDateBillingEntitiesParamsBillingEntityCounterPricings:
		return true
	}
	return false
}

type AccountGetChildrenParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID     param.Field[string] `path:"orgId,required"`
	NextToken param.Field[string] `query:"nextToken"`
	PageSize  param.Field[int64]  `query:"pageSize"`
}

// URLQuery serializes [AccountGetChildrenParams]'s query parameters as
// `url.Values`.
func (r AccountGetChildrenParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountSearchParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// `fromDocument` for multi page retrievals.
	FromDocument param.Field[int64] `query:"fromDocument"`
	// Search Operator to be used while querying search.
	Operator param.Field[AccountSearchParamsOperator] `query:"operator"`
	// Number of Accounts to retrieve per page.
	//
	// **NOTE:** If not defined, default is 10.
	PageSize param.Field[int64] `query:"pageSize"`
	// Query for data using special syntax:
	//
	// - Query parameters should be delimited using the $ (dollar sign).
	// - Allowed comparators are:
	//   - (greater than) >
	//   - (greater than or equal to) >=
	//   - (equal to) :
	//   - (less than) <
	//   - (less than or equal to) <=
	//   - (match phrase/prefix) ~
	//   - Allowed parameters are: name, code, currency, purchaseOrderNumber,
	//     parentAccountId, codes, id, createdBy, dtCreated, lastModifiedBy, ids.
	//   - Query example:
	//   - searchQuery=name~Premium On$currency:USD.
	//   - This query is translated into: find accounts whose name contains the
	//     phrase/prefix 'Premium On' AND the account currency is USD.
	//
	// **Note:** Using the ~ match phrase/prefix comparator. For best results, we
	// recommend treating this as a "starts with" comparator for your search query.
	SearchQuery param.Field[string] `query:"searchQuery"`
	// Name of the parameter on which sorting is performed. Use any field available on
	// the Account entity to sort by, such as `name`, `code`, and so on.
	SortBy param.Field[string] `query:"sortBy"`
	// Sorting order.
	SortOrder param.Field[AccountSearchParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [AccountSearchParams]'s query parameters as `url.Values`.
func (r AccountSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Search Operator to be used while querying search.
type AccountSearchParamsOperator string

const (
	AccountSearchParamsOperatorAnd AccountSearchParamsOperator = "AND"
	AccountSearchParamsOperatorOr  AccountSearchParamsOperator = "OR"
)

func (r AccountSearchParamsOperator) IsKnown() bool {
	switch r {
	case AccountSearchParamsOperatorAnd, AccountSearchParamsOperatorOr:
		return true
	}
	return false
}

// Sorting order.
type AccountSearchParamsSortOrder string

const (
	AccountSearchParamsSortOrderAsc  AccountSearchParamsSortOrder = "ASC"
	AccountSearchParamsSortOrderDesc AccountSearchParamsSortOrder = "DESC"
)

func (r AccountSearchParamsSortOrder) IsKnown() bool {
	switch r {
	case AccountSearchParamsSortOrderAsc, AccountSearchParamsSortOrderDesc:
		return true
	}
	return false
}
