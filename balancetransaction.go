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
func (r *BalanceTransactionService) Summary(ctx context.Context, balanceID string, query BalanceTransactionSummaryParams, opts ...option.RequestOption) (res *BalanceTransactionSummaryResponse, err error) {
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
	if balanceID == "" {
		err = errors.New("missing required balanceId parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/balances/%s/transactions/summary", query.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64 `json:"version,required"`
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
	TransactionTypeID string                  `json:"transactionTypeId"`
	JSON              transactionResponseJSON `json:"-"`
}

// transactionResponseJSON contains the JSON metadata for the struct
// [TransactionResponse]
type transactionResponseJSON struct {
	ID                apijson.Field
	Version           apijson.Field
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
)

func (r TransactionResponseEntityType) IsKnown() bool {
	switch r {
	case TransactionResponseEntityTypeBill, TransactionResponseEntityTypeCommitment, TransactionResponseEntityTypeUser, TransactionResponseEntityTypeServiceUser:
		return true
	}
	return false
}

type BalanceTransactionSummaryResponse struct {
	InitialCreditAmount float64                               `json:"initialCreditAmount"`
	TotalCreditAmount   float64                               `json:"totalCreditAmount"`
	TotalDebitAmount    float64                               `json:"totalDebitAmount"`
	JSON                balanceTransactionSummaryResponseJSON `json:"-"`
}

// balanceTransactionSummaryResponseJSON contains the JSON metadata for the struct
// [BalanceTransactionSummaryResponse]
type balanceTransactionSummaryResponseJSON struct {
	InitialCreditAmount apijson.Field
	TotalCreditAmount   apijson.Field
	TotalDebitAmount    apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *BalanceTransactionSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceTransactionSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type BalanceTransactionNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The financial value of the transaction.
	Amount param.Field[float64] `json:"amount,required"`
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
	OrgID param.Field[string] `path:"orgId,required"`
	// `nextToken` for multi page retrievals. A token for retrieving the next page of
	// transactions. You'll get this from the response to your request.
	NextToken param.Field[string] `query:"nextToken"`
	// The maximum number of transactions to return per page.
	PageSize          param.Field[int64]  `query:"pageSize"`
	ScheduleID        param.Field[string] `query:"scheduleId"`
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

type BalanceTransactionSummaryParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
