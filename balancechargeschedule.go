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

// BalanceChargeScheduleService contains methods and other services that help with
// interacting with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceChargeScheduleService] method instead.
type BalanceChargeScheduleService struct {
	Options []option.RequestOption
}

// NewBalanceChargeScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBalanceChargeScheduleService(opts ...option.RequestOption) (r *BalanceChargeScheduleService) {
	r = &BalanceChargeScheduleService{}
	r.Options = opts
	return
}

// Create a new BalanceChargeSchedule.
func (r *BalanceChargeScheduleService) New(ctx context.Context, balanceID string, params BalanceChargeScheduleNewParams, opts ...option.RequestOption) (res *BalanceChargeScheduleNewResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/balancechargeschedules", params.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a BalanceChargeSchedule for the given UUID.
func (r *BalanceChargeScheduleService) Get(ctx context.Context, balanceID string, id string, query BalanceChargeScheduleGetParams, opts ...option.RequestOption) (res *BalanceChargeScheduleGetResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/balancechargeschedules/%s", query.OrgID, balanceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a BalanceChargeSchedule for the given UUID.
func (r *BalanceChargeScheduleService) Update(ctx context.Context, balanceID string, id string, params BalanceChargeScheduleUpdateParams, opts ...option.RequestOption) (res *BalanceChargeScheduleUpdateResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/balancechargeschedules/%s", params.OrgID, balanceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Retrieve a list of BalanceChargeSchedule entities
func (r *BalanceChargeScheduleService) List(ctx context.Context, balanceID string, params BalanceChargeScheduleListParams, opts ...option.RequestOption) (res *pagination.Cursor[BalanceChargeScheduleListResponse], err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/balancechargeschedules", params.OrgID, balanceID)
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

// Retrieve a list of BalanceChargeSchedule entities
func (r *BalanceChargeScheduleService) ListAutoPaging(ctx context.Context, balanceID string, params BalanceChargeScheduleListParams, opts ...option.RequestOption) *pagination.CursorAutoPager[BalanceChargeScheduleListResponse] {
	return pagination.NewCursorAutoPager(r.List(ctx, balanceID, params, opts...))
}

// Delete the BalanceChargeSchedule for the given UUID.
func (r *BalanceChargeScheduleService) Delete(ctx context.Context, balanceID string, id string, body BalanceChargeScheduleDeleteParams, opts ...option.RequestOption) (res *BalanceChargeScheduleDeleteResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/balancechargeschedules/%s", body.OrgID, balanceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Previews the Charges this Schedule would create, without persisting them. You
// can use this call to obtain a preview of the Charges a Schedule you plan to
// create for a Balance would generate.
func (r *BalanceChargeScheduleService) Preview(ctx context.Context, balanceID string, params BalanceChargeSchedulePreviewParams, opts ...option.RequestOption) (res *BalanceChargeSchedulePreviewResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/balances/%s/balancechargeschedules/preview", params.OrgID, balanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BalanceChargeScheduleNewResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription string `json:"chargeDescription,required"`
	// The amount of each Charge created by the Balance Charge Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Charge Schedule was
	// created for.
	BalanceID string `json:"balanceId"`
	// Specifies a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// generated for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If blank, then the relevant Epoch date from your
	// Organization's configuration is used.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency BalanceChargeScheduleNewResponseBillFrequency `json:"billFrequency"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// Specifies how Charges created by the Balance Charge Schedule are billed - either
	// in arrears or in advance:
	//
	// - If `false` then billing is in arrears.
	// - If `true` then billing is in advance.
	BillInAdvance bool `json:"billInAdvance"`
	// Unique short code for the Balance Charge Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created the Balance Charge
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency of the Charges created by the Balance Charge Schedule.
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
	CustomFields map[string]BalanceChargeScheduleNewResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance Charge
	// Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Charge Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Charge will be created by
	// the Balance Charge Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the previous Charge was generated
	// by the Balance Charge Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit price. If the Charge was created with `amount` only, then will be null.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units. If the Charge was created with `amount` only, then will be
	// null.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                `json:"version"`
	JSON    balanceChargeScheduleNewResponseJSON `json:"-"`
}

// balanceChargeScheduleNewResponseJSON contains the JSON metadata for the struct
// [BalanceChargeScheduleNewResponse]
type balanceChargeScheduleNewResponseJSON struct {
	ID                     apijson.Field
	ChargeDescription      apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	BillEpoch              apijson.Field
	BillFrequency          apijson.Field
	BillFrequencyInterval  apijson.Field
	BillInAdvance          apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	PreviousRun            apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BalanceChargeScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceChargeScheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleNewResponseBillFrequency string

const (
	BalanceChargeScheduleNewResponseBillFrequencyDaily    BalanceChargeScheduleNewResponseBillFrequency = "DAILY"
	BalanceChargeScheduleNewResponseBillFrequencyWeekly   BalanceChargeScheduleNewResponseBillFrequency = "WEEKLY"
	BalanceChargeScheduleNewResponseBillFrequencyMonthly  BalanceChargeScheduleNewResponseBillFrequency = "MONTHLY"
	BalanceChargeScheduleNewResponseBillFrequencyAnnually BalanceChargeScheduleNewResponseBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleNewResponseBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleNewResponseBillFrequencyDaily, BalanceChargeScheduleNewResponseBillFrequencyWeekly, BalanceChargeScheduleNewResponseBillFrequencyMonthly, BalanceChargeScheduleNewResponseBillFrequencyAnnually:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceChargeScheduleNewResponseCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleNewResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceChargeScheduleNewResponseCustomFieldsUnion)(nil)).Elem(),
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

type BalanceChargeScheduleGetResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription string `json:"chargeDescription,required"`
	// The amount of each Charge created by the Balance Charge Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Charge Schedule was
	// created for.
	BalanceID string `json:"balanceId"`
	// Specifies a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// generated for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If blank, then the relevant Epoch date from your
	// Organization's configuration is used.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency BalanceChargeScheduleGetResponseBillFrequency `json:"billFrequency"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// Specifies how Charges created by the Balance Charge Schedule are billed - either
	// in arrears or in advance:
	//
	// - If `false` then billing is in arrears.
	// - If `true` then billing is in advance.
	BillInAdvance bool `json:"billInAdvance"`
	// Unique short code for the Balance Charge Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created the Balance Charge
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency of the Charges created by the Balance Charge Schedule.
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
	CustomFields map[string]BalanceChargeScheduleGetResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance Charge
	// Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Charge Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Charge will be created by
	// the Balance Charge Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the previous Charge was generated
	// by the Balance Charge Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit price. If the Charge was created with `amount` only, then will be null.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units. If the Charge was created with `amount` only, then will be
	// null.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                `json:"version"`
	JSON    balanceChargeScheduleGetResponseJSON `json:"-"`
}

// balanceChargeScheduleGetResponseJSON contains the JSON metadata for the struct
// [BalanceChargeScheduleGetResponse]
type balanceChargeScheduleGetResponseJSON struct {
	ID                     apijson.Field
	ChargeDescription      apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	BillEpoch              apijson.Field
	BillFrequency          apijson.Field
	BillFrequencyInterval  apijson.Field
	BillInAdvance          apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	PreviousRun            apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BalanceChargeScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceChargeScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleGetResponseBillFrequency string

const (
	BalanceChargeScheduleGetResponseBillFrequencyDaily    BalanceChargeScheduleGetResponseBillFrequency = "DAILY"
	BalanceChargeScheduleGetResponseBillFrequencyWeekly   BalanceChargeScheduleGetResponseBillFrequency = "WEEKLY"
	BalanceChargeScheduleGetResponseBillFrequencyMonthly  BalanceChargeScheduleGetResponseBillFrequency = "MONTHLY"
	BalanceChargeScheduleGetResponseBillFrequencyAnnually BalanceChargeScheduleGetResponseBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleGetResponseBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleGetResponseBillFrequencyDaily, BalanceChargeScheduleGetResponseBillFrequencyWeekly, BalanceChargeScheduleGetResponseBillFrequencyMonthly, BalanceChargeScheduleGetResponseBillFrequencyAnnually:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceChargeScheduleGetResponseCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleGetResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceChargeScheduleGetResponseCustomFieldsUnion)(nil)).Elem(),
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

type BalanceChargeScheduleUpdateResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription string `json:"chargeDescription,required"`
	// The amount of each Charge created by the Balance Charge Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Charge Schedule was
	// created for.
	BalanceID string `json:"balanceId"`
	// Specifies a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// generated for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If blank, then the relevant Epoch date from your
	// Organization's configuration is used.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency BalanceChargeScheduleUpdateResponseBillFrequency `json:"billFrequency"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// Specifies how Charges created by the Balance Charge Schedule are billed - either
	// in arrears or in advance:
	//
	// - If `false` then billing is in arrears.
	// - If `true` then billing is in advance.
	BillInAdvance bool `json:"billInAdvance"`
	// Unique short code for the Balance Charge Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created the Balance Charge
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency of the Charges created by the Balance Charge Schedule.
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
	CustomFields map[string]BalanceChargeScheduleUpdateResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance Charge
	// Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Charge Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Charge will be created by
	// the Balance Charge Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the previous Charge was generated
	// by the Balance Charge Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit price. If the Charge was created with `amount` only, then will be null.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units. If the Charge was created with `amount` only, then will be
	// null.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                   `json:"version"`
	JSON    balanceChargeScheduleUpdateResponseJSON `json:"-"`
}

// balanceChargeScheduleUpdateResponseJSON contains the JSON metadata for the
// struct [BalanceChargeScheduleUpdateResponse]
type balanceChargeScheduleUpdateResponseJSON struct {
	ID                     apijson.Field
	ChargeDescription      apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	BillEpoch              apijson.Field
	BillFrequency          apijson.Field
	BillFrequencyInterval  apijson.Field
	BillInAdvance          apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	PreviousRun            apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BalanceChargeScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceChargeScheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleUpdateResponseBillFrequency string

const (
	BalanceChargeScheduleUpdateResponseBillFrequencyDaily    BalanceChargeScheduleUpdateResponseBillFrequency = "DAILY"
	BalanceChargeScheduleUpdateResponseBillFrequencyWeekly   BalanceChargeScheduleUpdateResponseBillFrequency = "WEEKLY"
	BalanceChargeScheduleUpdateResponseBillFrequencyMonthly  BalanceChargeScheduleUpdateResponseBillFrequency = "MONTHLY"
	BalanceChargeScheduleUpdateResponseBillFrequencyAnnually BalanceChargeScheduleUpdateResponseBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleUpdateResponseBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleUpdateResponseBillFrequencyDaily, BalanceChargeScheduleUpdateResponseBillFrequencyWeekly, BalanceChargeScheduleUpdateResponseBillFrequencyMonthly, BalanceChargeScheduleUpdateResponseBillFrequencyAnnually:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceChargeScheduleUpdateResponseCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleUpdateResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceChargeScheduleUpdateResponseCustomFieldsUnion)(nil)).Elem(),
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

type BalanceChargeScheduleListResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription string `json:"chargeDescription,required"`
	// The amount of each Charge created by the Balance Charge Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Charge Schedule was
	// created for.
	BalanceID string `json:"balanceId"`
	// Specifies a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// generated for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If blank, then the relevant Epoch date from your
	// Organization's configuration is used.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency BalanceChargeScheduleListResponseBillFrequency `json:"billFrequency"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// Specifies how Charges created by the Balance Charge Schedule are billed - either
	// in arrears or in advance:
	//
	// - If `false` then billing is in arrears.
	// - If `true` then billing is in advance.
	BillInAdvance bool `json:"billInAdvance"`
	// Unique short code for the Balance Charge Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created the Balance Charge
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency of the Charges created by the Balance Charge Schedule.
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
	CustomFields map[string]BalanceChargeScheduleListResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance Charge
	// Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Charge Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Charge will be created by
	// the Balance Charge Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the previous Charge was generated
	// by the Balance Charge Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit price. If the Charge was created with `amount` only, then will be null.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units. If the Charge was created with `amount` only, then will be
	// null.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                 `json:"version"`
	JSON    balanceChargeScheduleListResponseJSON `json:"-"`
}

// balanceChargeScheduleListResponseJSON contains the JSON metadata for the struct
// [BalanceChargeScheduleListResponse]
type balanceChargeScheduleListResponseJSON struct {
	ID                     apijson.Field
	ChargeDescription      apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	BillEpoch              apijson.Field
	BillFrequency          apijson.Field
	BillFrequencyInterval  apijson.Field
	BillInAdvance          apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	PreviousRun            apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BalanceChargeScheduleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceChargeScheduleListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleListResponseBillFrequency string

const (
	BalanceChargeScheduleListResponseBillFrequencyDaily    BalanceChargeScheduleListResponseBillFrequency = "DAILY"
	BalanceChargeScheduleListResponseBillFrequencyWeekly   BalanceChargeScheduleListResponseBillFrequency = "WEEKLY"
	BalanceChargeScheduleListResponseBillFrequencyMonthly  BalanceChargeScheduleListResponseBillFrequency = "MONTHLY"
	BalanceChargeScheduleListResponseBillFrequencyAnnually BalanceChargeScheduleListResponseBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleListResponseBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleListResponseBillFrequencyDaily, BalanceChargeScheduleListResponseBillFrequencyWeekly, BalanceChargeScheduleListResponseBillFrequencyMonthly, BalanceChargeScheduleListResponseBillFrequencyAnnually:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceChargeScheduleListResponseCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleListResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceChargeScheduleListResponseCustomFieldsUnion)(nil)).Elem(),
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

type BalanceChargeScheduleDeleteResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription string `json:"chargeDescription,required"`
	// The amount of each Charge created by the Balance Charge Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Charge Schedule was
	// created for.
	BalanceID string `json:"balanceId"`
	// Specifies a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// generated for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If blank, then the relevant Epoch date from your
	// Organization's configuration is used.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency BalanceChargeScheduleDeleteResponseBillFrequency `json:"billFrequency"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// Specifies how Charges created by the Balance Charge Schedule are billed - either
	// in arrears or in advance:
	//
	// - If `false` then billing is in arrears.
	// - If `true` then billing is in advance.
	BillInAdvance bool `json:"billInAdvance"`
	// Unique short code for the Balance Charge Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created the Balance Charge
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency of the Charges created by the Balance Charge Schedule.
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
	CustomFields map[string]BalanceChargeScheduleDeleteResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance Charge
	// Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Charge Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Charge will be created by
	// the Balance Charge Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the previous Charge was generated
	// by the Balance Charge Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit price. If the Charge was created with `amount` only, then will be null.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units. If the Charge was created with `amount` only, then will be
	// null.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                   `json:"version"`
	JSON    balanceChargeScheduleDeleteResponseJSON `json:"-"`
}

// balanceChargeScheduleDeleteResponseJSON contains the JSON metadata for the
// struct [BalanceChargeScheduleDeleteResponse]
type balanceChargeScheduleDeleteResponseJSON struct {
	ID                     apijson.Field
	ChargeDescription      apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	BillEpoch              apijson.Field
	BillFrequency          apijson.Field
	BillFrequencyInterval  apijson.Field
	BillInAdvance          apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	PreviousRun            apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BalanceChargeScheduleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceChargeScheduleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleDeleteResponseBillFrequency string

const (
	BalanceChargeScheduleDeleteResponseBillFrequencyDaily    BalanceChargeScheduleDeleteResponseBillFrequency = "DAILY"
	BalanceChargeScheduleDeleteResponseBillFrequencyWeekly   BalanceChargeScheduleDeleteResponseBillFrequency = "WEEKLY"
	BalanceChargeScheduleDeleteResponseBillFrequencyMonthly  BalanceChargeScheduleDeleteResponseBillFrequency = "MONTHLY"
	BalanceChargeScheduleDeleteResponseBillFrequencyAnnually BalanceChargeScheduleDeleteResponseBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleDeleteResponseBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleDeleteResponseBillFrequencyDaily, BalanceChargeScheduleDeleteResponseBillFrequencyWeekly, BalanceChargeScheduleDeleteResponseBillFrequencyMonthly, BalanceChargeScheduleDeleteResponseBillFrequencyAnnually:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceChargeScheduleDeleteResponseCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleDeleteResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceChargeScheduleDeleteResponseCustomFieldsUnion)(nil)).Elem(),
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

type BalanceChargeSchedulePreviewResponse struct {
	// The UUID of the entity.
	ID string `json:"id,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription string `json:"chargeDescription,required"`
	// The amount of each Charge created by the Balance Charge Schedule.
	Amount float64 `json:"amount"`
	// The unique identifier (UUID) for the Balance this Balance Charge Schedule was
	// created for.
	BalanceID string `json:"balanceId"`
	// Specifies a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// generated for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If blank, then the relevant Epoch date from your
	// Organization's configuration is used.
	BillEpoch time.Time `json:"billEpoch" format:"date"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency BalanceChargeSchedulePreviewResponseBillFrequency `json:"billFrequency"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval int64 `json:"billFrequencyInterval"`
	// Specifies how Charges created by the Balance Charge Schedule are billed - either
	// in arrears or in advance:
	//
	// - If `false` then billing is in arrears.
	// - If `true` then billing is in advance.
	BillInAdvance bool `json:"billInAdvance"`
	// Unique short code for the Balance Charge Schedule.
	Code string `json:"code"`
	// The unique identifier (UUID) of the user who created the Balance Charge
	// Schedule.
	CreatedBy string `json:"createdBy"`
	// The currency of the Charges created by the Balance Charge Schedule.
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
	CustomFields map[string]BalanceChargeSchedulePreviewResponseCustomFieldsUnion `json:"customFields"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the Balance Charge Schedule was
	// last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The unique identifier (UUID) for the user who last modified the Balance Charge
	// Schedule.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The name of the Balance Charge Schedule.
	Name string `json:"name"`
	// The date and time (_in ISO-8601 format_) when the next Charge will be created by
	// the Balance Charge Schedule.
	NextRun time.Time `json:"nextRun" format:"date-time"`
	// The date and time (_in ISO-8601 format_) when the previous Charge was generated
	// by the Balance Charge Schedule.
	PreviousRun time.Time `json:"previousRun" format:"date-time"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" format:"date-time"`
	// The service period start date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" format:"date-time"`
	// Unit price. If the Charge was created with `amount` only, then will be null.
	UnitPrice float64 `json:"unitPrice"`
	// Number of units. If the Charge was created with `amount` only, then will be
	// null.
	Units float64 `json:"units"`
	// The version number:
	//
	//   - **Create:** On initial Create to insert a new entity, the version is set at 1
	//     in the response.
	//   - **Update:** On successful Update, the version is incremented by 1 in the
	//     response.
	Version int64                                    `json:"version"`
	JSON    balanceChargeSchedulePreviewResponseJSON `json:"-"`
}

// balanceChargeSchedulePreviewResponseJSON contains the JSON metadata for the
// struct [BalanceChargeSchedulePreviewResponse]
type balanceChargeSchedulePreviewResponseJSON struct {
	ID                     apijson.Field
	ChargeDescription      apijson.Field
	Amount                 apijson.Field
	BalanceID              apijson.Field
	BillEpoch              apijson.Field
	BillFrequency          apijson.Field
	BillFrequencyInterval  apijson.Field
	BillInAdvance          apijson.Field
	Code                   apijson.Field
	CreatedBy              apijson.Field
	Currency               apijson.Field
	CustomFields           apijson.Field
	DtCreated              apijson.Field
	DtLastModified         apijson.Field
	LastModifiedBy         apijson.Field
	Name                   apijson.Field
	NextRun                apijson.Field
	PreviousRun            apijson.Field
	ServicePeriodEndDate   apijson.Field
	ServicePeriodStartDate apijson.Field
	UnitPrice              apijson.Field
	Units                  apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BalanceChargeSchedulePreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r balanceChargeSchedulePreviewResponseJSON) RawJSON() string {
	return r.raw
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeSchedulePreviewResponseBillFrequency string

const (
	BalanceChargeSchedulePreviewResponseBillFrequencyDaily    BalanceChargeSchedulePreviewResponseBillFrequency = "DAILY"
	BalanceChargeSchedulePreviewResponseBillFrequencyWeekly   BalanceChargeSchedulePreviewResponseBillFrequency = "WEEKLY"
	BalanceChargeSchedulePreviewResponseBillFrequencyMonthly  BalanceChargeSchedulePreviewResponseBillFrequency = "MONTHLY"
	BalanceChargeSchedulePreviewResponseBillFrequencyAnnually BalanceChargeSchedulePreviewResponseBillFrequency = "ANNUALLY"
)

func (r BalanceChargeSchedulePreviewResponseBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeSchedulePreviewResponseBillFrequencyDaily, BalanceChargeSchedulePreviewResponseBillFrequencyWeekly, BalanceChargeSchedulePreviewResponseBillFrequencyMonthly, BalanceChargeSchedulePreviewResponseBillFrequencyAnnually:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type BalanceChargeSchedulePreviewResponseCustomFieldsUnion interface {
	ImplementsBalanceChargeSchedulePreviewResponseCustomFieldsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BalanceChargeSchedulePreviewResponseCustomFieldsUnion)(nil)).Elem(),
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

type BalanceChargeScheduleNewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency param.Field[BalanceChargeScheduleNewParamsBillFrequency] `json:"billFrequency,required"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval param.Field[int64] `json:"billFrequencyInterval,required"`
	// Used to specify how Charges created by the Balance Charge Schedule are billed -
	// either in arrears or in advance:
	//
	// - Use `false` for billing in arrears.
	// - Use `true` for billing in advance.
	BillInAdvance param.Field[bool] `json:"billInAdvance,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription param.Field[string] `json:"chargeDescription,required"`
	// Unique short code for the Balance Charge Schedule.
	Code param.Field[string] `json:"code,required"`
	// The currency of the Charges created by the Balance Charge Schedule.
	Currency param.Field[string] `json:"currency,required"`
	// The name of the Balance Charge Schedule.
	Name param.Field[string] `json:"name,required"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date (_in ISO-8601 format)_ of the Balance Charge
	// Schedule.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// The amount of each Charge created by the Balance Charge Schedule. Must be
	// omitted if `units` and `unitPrice` are provided.
	Amount param.Field[float64] `json:"amount"`
	// Specify a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// created for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If left blank, then the relevant Epoch date from your
	// Organization's configuration will be used as the billing cycle date instead.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
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
	CustomFields param.Field[map[string]BalanceChargeScheduleNewParamsCustomFieldsUnion] `json:"customFields"`
	// Unit price for Charge. Must be provided when `units` is used. Must be omitted
	// when `amount` is used.
	UnitPrice param.Field[float64] `json:"unitPrice"`
	// Number of units defined for the Charges created by the Schedule. Required when
	// `unitPrice` is provided. Must be omitted when `amount` is used.
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

func (r BalanceChargeScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleNewParamsBillFrequency string

const (
	BalanceChargeScheduleNewParamsBillFrequencyDaily    BalanceChargeScheduleNewParamsBillFrequency = "DAILY"
	BalanceChargeScheduleNewParamsBillFrequencyWeekly   BalanceChargeScheduleNewParamsBillFrequency = "WEEKLY"
	BalanceChargeScheduleNewParamsBillFrequencyMonthly  BalanceChargeScheduleNewParamsBillFrequency = "MONTHLY"
	BalanceChargeScheduleNewParamsBillFrequencyAnnually BalanceChargeScheduleNewParamsBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleNewParamsBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleNewParamsBillFrequencyDaily, BalanceChargeScheduleNewParamsBillFrequencyWeekly, BalanceChargeScheduleNewParamsBillFrequencyMonthly, BalanceChargeScheduleNewParamsBillFrequencyAnnually:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type BalanceChargeScheduleNewParamsCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleNewParamsCustomFieldsUnion()
}

type BalanceChargeScheduleGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type BalanceChargeScheduleUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency param.Field[BalanceChargeScheduleUpdateParamsBillFrequency] `json:"billFrequency,required"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval param.Field[int64] `json:"billFrequencyInterval,required"`
	// Used to specify how Charges created by the Balance Charge Schedule are billed -
	// either in arrears or in advance:
	//
	// - Use `false` for billing in arrears.
	// - Use `true` for billing in advance.
	BillInAdvance param.Field[bool] `json:"billInAdvance,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription param.Field[string] `json:"chargeDescription,required"`
	// Unique short code for the Balance Charge Schedule.
	Code param.Field[string] `json:"code,required"`
	// The currency of the Charges created by the Balance Charge Schedule.
	Currency param.Field[string] `json:"currency,required"`
	// The name of the Balance Charge Schedule.
	Name param.Field[string] `json:"name,required"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date (_in ISO-8601 format)_ of the Balance Charge
	// Schedule.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// The amount of each Charge created by the Balance Charge Schedule. Must be
	// omitted if `units` and `unitPrice` are provided.
	Amount param.Field[float64] `json:"amount"`
	// Specify a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// created for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If left blank, then the relevant Epoch date from your
	// Organization's configuration will be used as the billing cycle date instead.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
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
	CustomFields param.Field[map[string]BalanceChargeScheduleUpdateParamsCustomFieldsUnion] `json:"customFields"`
	// Unit price for Charge. Must be provided when `units` is used. Must be omitted
	// when `amount` is used.
	UnitPrice param.Field[float64] `json:"unitPrice"`
	// Number of units defined for the Charges created by the Schedule. Required when
	// `unitPrice` is provided. Must be omitted when `amount` is used.
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

func (r BalanceChargeScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeScheduleUpdateParamsBillFrequency string

const (
	BalanceChargeScheduleUpdateParamsBillFrequencyDaily    BalanceChargeScheduleUpdateParamsBillFrequency = "DAILY"
	BalanceChargeScheduleUpdateParamsBillFrequencyWeekly   BalanceChargeScheduleUpdateParamsBillFrequency = "WEEKLY"
	BalanceChargeScheduleUpdateParamsBillFrequencyMonthly  BalanceChargeScheduleUpdateParamsBillFrequency = "MONTHLY"
	BalanceChargeScheduleUpdateParamsBillFrequencyAnnually BalanceChargeScheduleUpdateParamsBillFrequency = "ANNUALLY"
)

func (r BalanceChargeScheduleUpdateParamsBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeScheduleUpdateParamsBillFrequencyDaily, BalanceChargeScheduleUpdateParamsBillFrequencyWeekly, BalanceChargeScheduleUpdateParamsBillFrequencyMonthly, BalanceChargeScheduleUpdateParamsBillFrequencyAnnually:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type BalanceChargeScheduleUpdateParamsCustomFieldsUnion interface {
	ImplementsBalanceChargeScheduleUpdateParamsCustomFieldsUnion()
}

type BalanceChargeScheduleListParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string]   `path:"orgId,required"`
	IDs   param.Field[[]string] `query:"ids"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of BalanceChargeSchedules to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
}

// URLQuery serializes [BalanceChargeScheduleListParams]'s query parameters as
// `url.Values`.
func (r BalanceChargeScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BalanceChargeScheduleDeleteParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type BalanceChargeSchedulePreviewParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// Represents standard scheduling frequencies options for a job.
	BillFrequency param.Field[BalanceChargeSchedulePreviewParamsBillFrequency] `json:"billFrequency,required"`
	// How often Bills are issued. For example, if billFrequency is `MONTHLY` and
	// `billFrequencyInterval` is 3, Bills are issued every three months.
	BillFrequencyInterval param.Field[int64] `json:"billFrequencyInterval,required"`
	// Used to specify how Charges created by the Balance Charge Schedule are billed -
	// either in arrears or in advance:
	//
	// - Use `false` for billing in arrears.
	// - Use `true` for billing in advance.
	BillInAdvance param.Field[bool] `json:"billInAdvance,required"`
	// The description for Charges created by the Balance Charge Schedule. Used on
	// Bills for Charge line items.
	ChargeDescription param.Field[string] `json:"chargeDescription,required"`
	// Unique short code for the Balance Charge Schedule.
	Code param.Field[string] `json:"code,required"`
	// The currency of the Charges created by the Balance Charge Schedule.
	Currency param.Field[string] `json:"currency,required"`
	// The name of the Balance Charge Schedule.
	Name param.Field[string] `json:"name,required"`
	// The service period end date (_in ISO-8601 format_) of the Balance Charge
	// Schedule.
	ServicePeriodEndDate param.Field[time.Time] `json:"servicePeriodEndDate,required" format:"date-time"`
	// The service period start date (_in ISO-8601 format)_ of the Balance Charge
	// Schedule.
	ServicePeriodStartDate param.Field[time.Time] `json:"servicePeriodStartDate,required" format:"date-time"`
	// nextToken for multi page retrievals
	NextToken param.Field[string] `query:"nextToken"`
	// Number of Charges to retrieve per page
	PageSize param.Field[int64] `query:"pageSize"`
	// The amount of each Charge created by the Balance Charge Schedule. Must be
	// omitted if `units` and `unitPrice` are provided.
	Amount param.Field[float64] `json:"amount"`
	// Specify a billing cycle date (_in ISO-8601 format_) for when the first Bill is
	// created for Balance Charges created by the Schedule, and also acts as a
	// reference for when in the Schedule period subsequent Bills are created for the
	// defined `billFrequency`. If left blank, then the relevant Epoch date from your
	// Organization's configuration will be used as the billing cycle date instead.
	BillEpoch param.Field[time.Time] `json:"billEpoch" format:"date"`
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
	CustomFields param.Field[map[string]BalanceChargeSchedulePreviewParamsCustomFieldsUnion] `json:"customFields"`
	// Unit price for Charge. Must be provided when `units` is used. Must be omitted
	// when `amount` is used.
	UnitPrice param.Field[float64] `json:"unitPrice"`
	// Number of units defined for the Charges created by the Schedule. Required when
	// `unitPrice` is provided. Must be omitted when `amount` is used.
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

func (r BalanceChargeSchedulePreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [BalanceChargeSchedulePreviewParams]'s query parameters as
// `url.Values`.
func (r BalanceChargeSchedulePreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Represents standard scheduling frequencies options for a job.
type BalanceChargeSchedulePreviewParamsBillFrequency string

const (
	BalanceChargeSchedulePreviewParamsBillFrequencyDaily    BalanceChargeSchedulePreviewParamsBillFrequency = "DAILY"
	BalanceChargeSchedulePreviewParamsBillFrequencyWeekly   BalanceChargeSchedulePreviewParamsBillFrequency = "WEEKLY"
	BalanceChargeSchedulePreviewParamsBillFrequencyMonthly  BalanceChargeSchedulePreviewParamsBillFrequency = "MONTHLY"
	BalanceChargeSchedulePreviewParamsBillFrequencyAnnually BalanceChargeSchedulePreviewParamsBillFrequency = "ANNUALLY"
)

func (r BalanceChargeSchedulePreviewParamsBillFrequency) IsKnown() bool {
	switch r {
	case BalanceChargeSchedulePreviewParamsBillFrequencyDaily, BalanceChargeSchedulePreviewParamsBillFrequencyWeekly, BalanceChargeSchedulePreviewParamsBillFrequencyMonthly, BalanceChargeSchedulePreviewParamsBillFrequencyAnnually:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type BalanceChargeSchedulePreviewParamsCustomFieldsUnion interface {
	ImplementsBalanceChargeSchedulePreviewParamsCustomFieldsUnion()
}
