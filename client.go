// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"net/http"
	"os"

	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the m3ter API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options              []option.RequestOption
	Authentication       *AuthenticationService
	Accounts             *AccountService
	AccountPlans         *AccountPlanService
	Aggregations         *AggregationService
	Balances             *BalanceService
	BillConfig           *BillConfigService
	Commitments          *CommitmentService
	CompoundAggregations *CompoundAggregationService
	Contracts            *ContractService
	Counters             *CounterService
	CounterAdjustments   *CounterAdjustmentService
	CounterPricings      *CounterPricingService
	CreditReasons        *CreditReasonService
	Currencies           *CurrencyService
	DebitReasons         *DebitReasonService
	Meters               *MeterService
	OrganizationConfig   *OrganizationConfigService
	Plans                *PlanService
	PlanGroups           *PlanGroupService
	PlanGroupLinks       *PlanGroupLinkService
	PlanTemplates        *PlanTemplateService
	Pricings             *PricingService
	Products             *ProductService
	TransactionTypes     *TransactionTypeService
	DataExports          *DataExportService
}

// NewClient generates a new client with the default option read from the
// environment (M3TER_API_KEY, M3TER_API_SECRET, M3TER_API_TOKEN). The option
// passed in as arguments are applied after these default arguments, and all option
// will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("M3TER_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_SECRET"); ok {
		defaults = append(defaults, option.WithAPISecret(o))
	}
	if o, ok := os.LookupEnv("M3TER_API_TOKEN"); ok {
		defaults = append(defaults, option.WithToken(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Authentication = NewAuthenticationService(opts...)
	r.Accounts = NewAccountService(opts...)
	r.AccountPlans = NewAccountPlanService(opts...)
	r.Aggregations = NewAggregationService(opts...)
	r.Balances = NewBalanceService(opts...)
	r.BillConfig = NewBillConfigService(opts...)
	r.Commitments = NewCommitmentService(opts...)
	r.CompoundAggregations = NewCompoundAggregationService(opts...)
	r.Contracts = NewContractService(opts...)
	r.Counters = NewCounterService(opts...)
	r.CounterAdjustments = NewCounterAdjustmentService(opts...)
	r.CounterPricings = NewCounterPricingService(opts...)
	r.CreditReasons = NewCreditReasonService(opts...)
	r.Currencies = NewCurrencyService(opts...)
	r.DebitReasons = NewDebitReasonService(opts...)
	r.Meters = NewMeterService(opts...)
	r.OrganizationConfig = NewOrganizationConfigService(opts...)
	r.Plans = NewPlanService(opts...)
	r.PlanGroups = NewPlanGroupService(opts...)
	r.PlanGroupLinks = NewPlanGroupLinkService(opts...)
	r.PlanTemplates = NewPlanTemplateService(opts...)
	r.Pricings = NewPricingService(opts...)
	r.Products = NewProductService(opts...)
	r.TransactionTypes = NewTransactionTypeService(opts...)
	r.DataExports = NewDataExportService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
