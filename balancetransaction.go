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

// BalanceTransactionService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceTransactionService] method instead.
type BalanceTransactionService struct {
	Options []option.RequestOption
}

// NewBalanceTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBalanceTransactionService(opts ...option.RequestOption) (r *BalanceTransactionService) {
	r = &BalanceTransactionService{}
	r.Options = opts
	return
}

// Add a Transaction to a Balance. This endpoint allows you to create a new
// Transaction amount for a Balance. This amount then becomes available at billing
// for draw-down to cover charges due. The Transaction details should be provided
// in the request body.
//
// Before you can add a Transaction amount, you must first set up Transaction Types
// at the Organization Level - see the
// [Transaction Type](https://www.m3ter.com/docs/api#tag/TransactionType) section
// in this API Reference for more details. You can then use this call to add an
// instance of a Transaction Type to a Balance.
//
// **Note:** If you have a customer whose payment is in a different currency to the
// Balance currency, you can use the `paid` and `paidCurrency` request parameters
// to record the amount paid and alternative currency respectively. For example,
// you might add a Transaction amount of 200 USD to a Balance on a customer Account
// where the customer actually paid you 50 units in virtual currency X.
func (r *BalanceTransactionService) New(ctx context.Context, balanceID string, params BalanceTransactionNewParams, opts ...option.RequestOption) (res *TransactionResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/transactions", params.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve all Transactions for a specific Balance.
//
// This endpoint returns a list of all Transactions associated with a specific
// Balance. You can paginate through the Transactions by using the `pageSize` and
// `nextToken` parameters.
func (r *BalanceTransactionService) List(ctx context.Context, balanceID string, params BalanceTransactionListParams, opts ...option.RequestOption) (res *pagination.Cursor[TransactionResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/transactions", params.OrgID, balanceID)
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

// Retrieve all Transactions for a specific Balance.
//
// This endpoint returns a list of all Transactions associated with a specific
// Balance. You can paginate through the Transactions by using the `pageSize` and
// `nextToken` parameters.
func (r *BalanceTransactionService) ListAutoPaging(ctx context.Context, balanceID string, params BalanceTransactionListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[TransactionResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, balanceID, params, opts...))
}

// Retrieves the Balance Transactions Summary for a given Balance.
//
// The response contains useful recorded and calculated Transaction amounts created
// for a Balance during the time it is active for the Account, including amounts
// relevant to any rollover amount configured for a Balance:
//
//   - `totalCreditAmount`. The sum of all credits amounts created for the Balance.
//   - `totalDebitAmount`. The sum of all debit amounts created for the Balance.
//   - `initialCreditAmount`. The initial credit amount created for the Balance.
//   - `expiredBalanceAmount`. The amount of the Balance remaining at the time the
//     Balance expires and which is not included in any configured Rollover amount.
//     For example, suppose a Balance reaches its end date and $1000 credit remains
//     unused. If the Balance is configured to rollover $800, then the
//     `expiredBalanceAmount` is calculated as $1000 - $800 = $200.
//   - `rolloverConsumed`. The sum of debits made against the configured rollover
//     amount. Note that this amount is dynamic relative to when the API call is made
//     until either the rollover end date is reached or the cap configured for the
//     rollover amount is reached, after which it will be unchanged. If no rollover
//     is configured for a Balance, then this is ignored.
//   - `balanceConsumed`. The sum of debits made against the Balance. Note that this
//     amount is dynamic relative to when the API call is made until either the
//     Balance end date is reached or the available Balance amount reaches zero,
//     after which it will be unchanged.
func (r *BalanceTransactionService) Summary(ctx context.Context, balanceID string, query BalanceTransactionSummaryParams, opts ...option.RequestOption) (res *BalanceTransactionSummaryResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/transactions/summary", query.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ScheduleRequestParam struct {
	// The amount of each Balance Transaction created by this Schedule.
	Amount param.Field[float64] `json:"amount" api:"required"`
	// The unique short code of the Balance Transaction Schedule.
	Code param.Field[string] `json:"code" api:"required"`
	// The end date (_in ISO-8601 format_) of the Balance Transaction Schedule.
	//
	// **NOTE:** End date is exclusive.
	EndDate param.Field[time.Time] `json:"endDate" api:"required" format:"date-time"`
	// Represents standard scheduling frequencies options for a job.
	Frequency param.Field[ScheduleRequestFrequency] `json:"frequency" api:"required"`
	// Used in conjunction with `frequency` to define how often Balance Transactions
	// are generated by the Schedule. For example, if `frequency` is MONTHLY and
	// `frequencyInterval` is 3, then Balance Transactions are generated every three
	// months.
	FrequencyInterval param.Field[int64] `json:"frequencyInterval" api:"required"`
	// The name of the Balance Transaction Schedule.
	Name param.Field[string] `json:"name" api:"required"`
	// The start date (_in ISO-8601 format_) of the Balance Transaction Schedule.
	StartDate param.Field[time.Time] `json:"startDate" api:"required" format:"date-time"`
	// The description of each Balance Transaction that will be created by this
	// Schedule.
	TransactionDescription param.Field[string] `json:"transactionDescription" api:"required"`
	// The unique identifier (UUID) of the transaction type used to create Transactions
	// by this Schedule. You can obtain a list of Transaction Types created for the
	// Organization. See the the
	// [List TransactionTypes](https://www.m3ter.com/docs/api#tag/TransactionType/operation/ListTransactionTypes)
	// endpoint of this API Reference.
	TransactionTypeID param.Field[string] `json:"transactionTypeId" api:"required"`
	// The currency code of the payment if it differs from the Balance currency. For
	// example: USD, GBP or EUR.
	CurrencyPaid param.Field[string] `json:"currencyPaid"`
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
	CustomFields param.Field[map[string]ScheduleRequestCustomFieldsUnionParam] `json:"customFields"`
	// The payment amount if the payment currency differs from the Balance currency.
	Paid param.Field[float64] `json:"paid"`
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

func (r ScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Represents standard scheduling frequencies options for a job.
type ScheduleRequestFrequency string

const (
	ScheduleRequestFrequencyDaily    ScheduleRequestFrequency = "DAILY"
	ScheduleRequestFrequencyWeekly   ScheduleRequestFrequency = "WEEKLY"
	ScheduleRequestFrequencyMonthly  ScheduleRequestFrequency = "MONTHLY"
	ScheduleRequestFrequencyAnnually ScheduleRequestFrequency = "ANNUALLY"
)

func (r ScheduleRequestFrequency) IsKnown() bool {
	switch r {
	case ScheduleRequestFrequencyDaily, ScheduleRequestFrequencyWeekly, ScheduleRequestFrequencyMonthly, ScheduleRequestFrequencyAnnually:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ScheduleRequestCustomFieldsUnionParam interface {
	ImplementsScheduleRequestCustomFieldsUnionParam()
}

type ScheduleResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// The amount of each Balance Transaction created by this Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Transaction Schedule
	// was created for.
	BalanceID string `json:"balanceId"`
	// Unique short code of the Balance Transaction Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created this Balance Transaction
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency code of the payment if it differs from the Balance currency. For
	// example: USD, GBP or EUR.
	CurrencyPaid string `json:"currencyPaid"`
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
	CustomFields map[string]ScheduleResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Transaction Schedule
	// was created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Transaction Schedule
	// was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The end date (_in ISO-8601 format_) of the Balance Transaction Schedule.
	EndDate time.Time `json:"endDate" format:"date-time"`
	// Represents standard scheduling frequencies options for a job.
	Frequency ScheduleResponseFrequency `json:"frequency"`
	// Used in conjunction with `frequency` to define how often Balance Transactions
	// are generated. For example, if `frequency` is MONTHLY and `frequencyInterval` is
	// 3, Balance Transactions are generated every three months.
	FrequencyInterval int64 `json:"frequencyInterval"`
	// The unique identifier (UUID) of the user who last modified this Balance
	// Transaction Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Transaction Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Transaction will be
	// generated by the Balance Transaction Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The payment amount if the payment currency differs from the Balance currency.
	Paid float64 `json:"paid"`
	// The date and time (_in ISO-8601 format_) when the previous Transaction was
	// generated by the Balance Transaction Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The start date (_in ISO-8601 format_) of the Balance Transaction Schedule.
	StartDate time.Time `json:"startDate" format:"date-time"`
	// The description of each Balance Transaction created by this Schedule.
	TransactionDescription string `json:"transactionDescription"`
	// The unique identifier (UUID) of the transaction type used for Transactions
	// created by this Schedule.
	TransactionTypeID string `json:"transactionTypeId"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                `json:"version"`
	JSON    scheduleResponseJSON `json:"-"`
}

// scheduleResponseJSON contains the JSON metadata for the struct
// [ScheduleResponse]
type scheduleResponseJSON struct {
	ID                     apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	CurrencyPaid           apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	EndDate                apijson.Field
	Frequency              apijson.Field
	FrequencyInterval      apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	Paid                   apijson.Field
	PreviousRun            apijson.Field
	StartDate              apijson.Field
	TransactionDescription apijson.Field
	TransactionTypeID      apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type ScheduleResponseCustomFieldsUnion interface {
	ImplementsScheduleResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScheduleResponseCustomFieldsUnion)(nil)).Elem(),
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

// Represents standard scheduling frequencies options for a job.
type ScheduleResponseFrequency string

const (
	ScheduleResponseFrequencyDaily    ScheduleResponseFrequency = "DAILY"
	ScheduleResponseFrequencyWeekly   ScheduleResponseFrequency = "WEEKLY"
	ScheduleResponseFrequencyMonthly  ScheduleResponseFrequency = "MONTHLY"
	ScheduleResponseFrequencyAnnually ScheduleResponseFrequency = "ANNUALLY"
)

func (r ScheduleResponseFrequency) IsKnown() bool {
	switch r {
	case ScheduleResponseFrequencyDaily, ScheduleResponseFrequencyWeekly, ScheduleResponseFrequencyMonthly, ScheduleResponseFrequencyAnnually:
		return true
	}
	return false
}

type TransactionResponse struct {
	// The UUID of the entity.
	ID string `json:"id" api:"required"`
	// The financial value of the transaction, as recorded in the balance.
	Amount float64 `json:"amount"`
	// The date _(in ISO 8601 format)_ when the balance transaction was applied, i.e.,
	// when the balance was affected.
	AppliedDate time.Time `json:"appliedDate" format:"date-time"`
	// The unique identifier (UUID) for the user who created the balance transaction.
	CreatedBy string `json:"createdBy"`
	// The currency code such as USD, GBP, EUR of the payment, if it differs from the
	// balance currency.
	CurrencyPaid string `json:"currencyPaid"`
	// A brief description explaining the purpose or context of the transaction.
	Description string `json:"description"`
	// The date and time _(in ISO 8601 format)_ when the balance transaction was first
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time _(in ISO 8601 format)_ when the balance transaction was last
	// modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the entity associated with the Transaction, as
	// specified by the `entityType`.
	EntityID string `json:"entityId"`
	// The type of entity associated with the Transaction - identifies who or what was
	// responsible for the Transaction being added to the Balance - such as a **User**,
	// a **Service User**, or a **Bill**.
	EntityType TransactionResponseEntityType `json:"entityType"`
	// The unique identifier (UUID) for the user who last modified the balance
	// transaction.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The actual payment amount if the payment currency differs from the Balance
	// currency.
	Paid float64 `json:"paid"`
	// The date _(in ISO 8601 format)_ when the transaction was recorded in the system.
	TransactionDate time.Time `json:"transactionDate" format:"date-time"`
	// The unique identifier (UUID) for the Transaction type. This is obtained from the
	// list of created Transaction Types within the Organization Configuration.
	TransactionTypeID string `json:"transactionTypeId"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                   `json:"version"`
	JSON    transactionResponseJSON `json:"-"`
}

// transactionResponseJSON contains the JSON metadata for the struct
// [TransactionResponse]
type transactionResponseJSON struct {
	ID                apijson.Field
	Amount            apijson.Field
	AppliedDate       apijson.Field
	CreatedBy         apijson.Field
	CurrencyPaid      apijson.Field
	Description       apijson.Field
	DtCreated         apijson.Field
	DtLastModified    apijson.Field
	EntityID          apijson.Field
	EntityType        apijson.Field
	LastModifiedBy    apijson.Field
	Paid              apijson.Field
	TransactionDate   apijson.Field
	TransactionTypeID apijson.Field
	Version           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TransactionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionResponseJSON) RawJSON() string {
	return r.raw
}

// The type of entity associated with the Transaction - identifies who or what was
// responsible for the Transaction being added to the Balance - such as a **User**,
// a **Service User**, or a **Bill**.
type TransactionResponseEntityType string

const (
	TransactionResponseEntityTypeBill        TransactionResponseEntityType = "BILL"
	TransactionResponseEntityTypeCommitment  TransactionResponseEntityType = "COMMITMENT"
	TransactionResponseEntityTypeUser        TransactionResponseEntityType = "USER"
	TransactionResponseEntityTypeServiceUser TransactionResponseEntityType = "SERVICE_USER"
	TransactionResponseEntityTypeScheduler   TransactionResponseEntityType = "SCHEDULER"
)

func (r TransactionResponseEntityType) IsKnown() bool {
	switch r {
	case TransactionResponseEntityTypeBill, TransactionResponseEntityTypeCommitment, TransactionResponseEntityTypeUser, TransactionResponseEntityTypeServiceUser, TransactionResponseEntityTypeScheduler:
		return true
	}
	return false
}

type BalanceTransactionSummaryResponse struct {
	// Amount consumed from the original balance
	BalanceConsumed float64 `json:"balanceConsumed"`
	// Amount of the balance that expired without being used
	ExpiredBalanceAmount float64 `json:"expiredBalanceAmount"`
	InitialCreditAmount  float64 `json:"initialCreditAmount"`
	// Amount consumed from rollover credit
	RolloverConsumed  float64                               `json:"rolloverConsumed"`
	TotalCreditAmount float64                               `json:"totalCreditAmount"`
	TotalDebitAmount  float64                               `json:"totalDebitAmount"`
	JSON              balanceTransactionSummaryResponseJSON `json:"-"`
}

// balanceTransactionSummaryResponseJSON contains the JSON metadata for the struct
// [BalanceTransactionSummaryResponse]
type balanceTransactionSummaryResponseJSON struct {
	BalanceConsumed      apijson.Field
	ExpiredBalanceAmount apijson.Field
	InitialCreditAmount  apijson.Field
	RolloverConsumed     apijson.Field
	TotalCreditAmount    apijson.Field
	TotalDebitAmount     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BalanceTransactionSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceTransactionSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type BalanceTransactionNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
	// The financial value of the transaction.
	Amount param.Field[float64] `json:"amount" api:"required"`
	// The date _(in ISO 8601 format)_ when the Balance transaction was applied.
	AppliedDate param.Field[time.Time] `json:"appliedDate" format:"date-time"`
	// The currency code of the payment if it differs from the Balance currency. For
	// example: USD, GBP or EUR.
	CurrencyPaid param.Field[string] `json:"currencyPaid"`
	// A brief description explaining the purpose and context of the transaction.
	Description param.Field[string] `json:"description"`
	// The payment amount if the payment currency differs from the Balance currency.
	Paid param.Field[float64] `json:"paid"`
	// The date _(in ISO 8601 format)_ when the transaction occurred.
	TransactionDate param.Field[time.Time] `json:"transactionDate" format:"date-time"`
	// The unique identifier (UUID) of the transaction type. This is obtained from the
	// list of created Transaction Types within the Organization Configuration.
	TransactionTypeID param.Field[string] `json:"transactionTypeId"`
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

func (r BalanceTransactionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BalanceTransactionListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID      param.Field[string]                                 `path:"orgId" api:"required"`
	EntityID   param.Field[string]                                 `query:"entityId"`
	EntityType param.Field[BalanceTransactionListParamsEntityType] `query:"entityType"`
	// `nextToken` for multi page retrievals. A token for retrieving the next page of
	// transactions. You'll get this from the response to your request.
	NextToken param.Field[string] `query:"nextToken"`
	// The maximum number of transactions to return per page.
	PageSize          param.Field[int64]  `query:"pageSize"`
	TransactionTypeID param.Field[string] `query:"transactionTypeId"`
}

// URLQuery serializes [BalanceTransactionListParams]'s query parameters as
// `url.Values`.
func (r BalanceTransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BalanceTransactionListParamsEntityType string

const (
	BalanceTransactionListParamsEntityTypeBill        BalanceTransactionListParamsEntityType = "BILL"
	BalanceTransactionListParamsEntityTypeCommitment  BalanceTransactionListParamsEntityType = "COMMITMENT"
	BalanceTransactionListParamsEntityTypeUser        BalanceTransactionListParamsEntityType = "USER"
	BalanceTransactionListParamsEntityTypeServiceUser BalanceTransactionListParamsEntityType = "SERVICE_USER"
	BalanceTransactionListParamsEntityTypeScheduler   BalanceTransactionListParamsEntityType = "SCHEDULER"
)

func (r BalanceTransactionListParamsEntityType) IsKnown() bool {
	switch r {
	case BalanceTransactionListParamsEntityTypeBill, BalanceTransactionListParamsEntityTypeCommitment, BalanceTransactionListParamsEntityTypeUser, BalanceTransactionListParamsEntityTypeServiceUser, BalanceTransactionListParamsEntityTypeScheduler:
		return true
	}
	return false
}

type BalanceTransactionSummaryParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId" api:"required"`
}
