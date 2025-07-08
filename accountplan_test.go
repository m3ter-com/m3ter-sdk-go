// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/m3ter-sdk-go/internal/testutil"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
)

func TestAccountPlanNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := m3ter.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithAPISecret("My API Secret"),
		option.WithToken("My Token"),
		option.WithOrgID("My Org ID"),
	)
	_, err := client.AccountPlans.New(context.TODO(), m3ter.AccountPlanNewParams{
		AccountID:        m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		StartDate:        m3ter.F(time.Now()),
		BillEpoch:        m3ter.F(time.Now()),
		ChildBillingMode: m3ter.F(m3ter.AccountPlanNewParamsChildBillingModeParentSummary),
		Code:             m3ter.F("S?oC\"$]C] ]]]]]5]"),
		ContractID:       m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		CustomFields: m3ter.F(map[string]m3ter.AccountPlanNewParamsCustomFieldsUnion{
			"foo": shared.UnionString("string"),
		}),
		EndDate:     m3ter.F(time.Now()),
		PlanGroupID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		PlanID:      m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Version:     m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPlanGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := m3ter.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithAPISecret("My API Secret"),
		option.WithToken("My Token"),
		option.WithOrgID("My Org ID"),
	)
	_, err := client.AccountPlans.Get(
		context.TODO(),
		"id",
		m3ter.AccountPlanGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPlanUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := m3ter.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithAPISecret("My API Secret"),
		option.WithToken("My Token"),
		option.WithOrgID("My Org ID"),
	)
	_, err := client.AccountPlans.Update(
		context.TODO(),
		"id",
		m3ter.AccountPlanUpdateParams{
			AccountID:        m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			StartDate:        m3ter.F(time.Now()),
			BillEpoch:        m3ter.F(time.Now()),
			ChildBillingMode: m3ter.F(m3ter.AccountPlanUpdateParamsChildBillingModeParentSummary),
			Code:             m3ter.F("S?oC\"$]C] ]]]]]5]"),
			ContractID:       m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			CustomFields: m3ter.F(map[string]m3ter.AccountPlanUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			EndDate:     m3ter.F(time.Now()),
			PlanGroupID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			PlanID:      m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Version:     m3ter.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPlanListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := m3ter.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithAPISecret("My API Secret"),
		option.WithToken("My Token"),
		option.WithOrgID("My Org ID"),
	)
	_, err := client.AccountPlans.List(context.TODO(), m3ter.AccountPlanListParams{
		Account:    m3ter.F("account"),
		Contract:   m3ter.F("contract"),
		Date:       m3ter.F("date"),
		IDs:        m3ter.F([]string{"string"}),
		Includeall: m3ter.F(true),
		NextToken:  m3ter.F("nextToken"),
		PageSize:   m3ter.F(int64(1)),
		Plan:       m3ter.F("plan"),
		Product:    m3ter.F("product"),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountPlanDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := m3ter.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithAPISecret("My API Secret"),
		option.WithToken("My Token"),
		option.WithOrgID("My Org ID"),
	)
	_, err := client.AccountPlans.Delete(
		context.TODO(),
		"id",
		m3ter.AccountPlanDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
