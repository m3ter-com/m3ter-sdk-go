// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

// StatementService contains methods and other services that help with interacting
// with the m3ter API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementService] method instead.
type StatementService struct {
	Options              []option.RequestOption
	StatementJobs        *StatementStatementJobService
	StatementDefinitions *StatementStatementDefinitionService
}

// NewStatementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStatementService(opts ...option.RequestOption) (r *StatementService) {
	r = &StatementService{}
	r.Options = opts
	r.StatementJobs = NewStatementStatementJobService(opts...)
	r.StatementDefinitions = NewStatementStatementDefinitionService(opts...)
	return
}

// Generate a specific Bill Statement for the provided Bill UUID in CSV format.
//
// Bill Statements are backing sheets to the invoices sent to your customers. Bill
// Statements provide a breakdown of the usage responsible for the usage charge
// line items shown on invoices.
//
// The response to this call returns a pre-signed `downloadUrl`, which you then use
// with a `GET` call to obtain the Bill statement in CSV format.
func (r *StatementService) NewCsv(ctx context.Context, id string, body StatementNewCsvParams, opts ...option.RequestOption) (res *ObjectURLResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/statement/csv", body.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieve a specific Bill Statement for the given Bill UUID in CSV format.
//
// Bill Statements are backing sheets to the invoices sent to your customers. Bill
// Statements provide a breakdown of the usage responsible for the usage charge
// line items shown on invoices.
//
// The response includes a pre-signed `downloadUrl`, which must be used with a
// separate `GET` call to download the actual Bill Statement. This ensures secure
// access to the requested information.
func (r *StatementService) GetCsv(ctx context.Context, id string, query StatementGetCsvParams, opts ...option.RequestOption) (res *ObjectURLResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/statement/csv", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a Bill Statement in JSON format for a given Bill ID.
//
// Bill Statements are backing sheets to the invoices sent to your customers. Bill
// Statements provide a breakdown of the usage responsible for the usage charge
// line items shown on invoices.
//
// The response to this call returns a pre-signed `downloadUrl`, which you use with
// a `GET` call to obtain the Bill Statement.
func (r *StatementService) GetJson(ctx context.Context, id string, query StatementGetJsonParams, opts ...option.RequestOption) (res *ObjectURLResponse, err error) {
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
	path := fmt.Sprintf("organizations/%s/bills/%s/statement/json", query.OrgID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ObjectURLResponse struct {
	// The pre-signed download URL.
	DownloadURL string                `json:"downloadUrl" format:"url"`
	JSON        objectURLResponseJSON `json:"-"`
}

// objectURLResponseJSON contains the JSON metadata for the struct
// [ObjectURLResponse]
type objectURLResponseJSON struct {
	DownloadURL apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectURLResponseJSON) RawJSON() string {
	return r.raw
}

type StatementNewCsvParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type StatementGetCsvParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}

type StatementGetJsonParams struct {
	// Use [option.WithOrgID] on the client to set a global default for this field.
	OrgID param.Field[string] `path:"orgId,required"`
}
