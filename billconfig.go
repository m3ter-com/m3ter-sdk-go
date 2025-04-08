// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// BillConfigService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillConfigService] method instead.
type BillConfigService struct {
	Options []option.RequestOption
}

// NewBillConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillConfigService(opts ...option.RequestOption) (r *BillConfigService) {
	r = &BillConfigService{}
	r.Options = opts
	return
}

// Retrieve the Organization-wide BillConfig.
func (r *BillConfigService) Get(ctx context.Context, query BillConfigGetParams, opts ...option.RequestOption) (res *BillConfigResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/billconfig", query.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Organization-wide BillConfig.
//
// You can use this endpoint to set a global lock date for **all** Bills - any Bill
// with a service period end date on or before the set date will be locked and
// cannot be updated or recalculated.
func (r *BillConfigService) Update(ctx context.Context, params BillConfigUpdateParams, opts ...option.RequestOption) (res *BillConfigResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/billconfig", params.OrgID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type BillConfigResponse struct {
	// The Organization UUID. The Organization represents your company as a direct
	// customer of the m3ter service.
	ID string `json:"id"`
	// The global lock date _(in ISO 8601 format)_ when all Bills will be locked.
	//
	// For example: `"2024-03-01"`.
	BillLockDate time.Time `json:"billLockDate" format:"date"`
	// The id of the user who created this bill config.
	CreatedBy string `json:"createdBy"`
	// The DateTime _(in ISO-8601 format)_ when the bill config was first created.
	DtCreated time.Time `json:"dtCreated" format:"date-time"`
	// The DateTime _(in ISO-8601 format)_ when the bill config was last modified.
	DtLastModified time.Time `json:"dtLastModified" format:"date-time"`
	// The id of the user who last modified this bill config.
	LastModifiedBy string `json:"lastModifiedBy"`
	// The version number:
	//
	// - Default value when newly created is one.
	// - Incremented by 1 each time it is updated.
	Version int64                  `json:"version"`
	JSON    billConfigResponseJSON `json:"-"`
}

// billConfigResponseJSON contains the JSON metadata for the struct
// [BillConfigResponse]
type billConfigResponseJSON struct {
	ID             apijson.Field
	BillLockDate   apijson.Field
	CreatedBy      apijson.Field
	DtCreated      apijson.Field
	DtLastModified apijson.Field
	LastModifiedBy apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BillConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billConfigResponseJSON) RawJSON() string {
	return r.raw
}

type BillConfigGetParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type BillConfigUpdateParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
	// The global lock date when all Bills will be locked _(in ISO 8601 format)_.
	//
	// For example: `"2024-03-01"`.
	BillLockDate param.Field[time.Time] `json:"billLockDate" format:"date"`
	// The version number:
	//
	//   - Default value when newly created is one.
	//   - On Update, version is required and must match the existing version because a
	//     check is performed to ensure sequential versioning is preserved. Version is
	//     incremented by 1 and listed in the response
	Version param.Field[int64] `json:"version"`
}

func (r BillConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
