// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/m3ter-com/m3ter-sdk-go"
	"github.com/m3ter-com/m3ter-sdk-go/internal/testutil"
	"github.com/m3ter-com/m3ter-sdk-go/option"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
)

func TestPlanNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Plans.New(
		context.TODO(),
		"orgId",
		m3ter.PlanNewParams{
			Code:           m3ter.F("JS!?Q0]r] ]$]"),
			Name:           m3ter.F("x"),
			PlanTemplateID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			AccountID:      m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Bespoke:        m3ter.F(true),
			CustomFields: m3ter.F(map[string]m3ter.PlanNewParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			MinimumSpend:                m3ter.F(0.000000),
			MinimumSpendBillInAdvance:   m3ter.F(true),
			MinimumSpendDescription:     m3ter.F("minimumSpendDescription"),
			Ordinal:                     m3ter.F(int64(0)),
			StandingCharge:              m3ter.F(0.000000),
			StandingChargeBillInAdvance: m3ter.F(true),
			StandingChargeDescription:   m3ter.F("standingChargeDescription"),
			Version:                     m3ter.F(int64(0)),
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

func TestPlanGet(t *testing.T) {
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
	)
	_, err := client.Plans.Get(
		context.TODO(),
		"orgId",
		"id",
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanUpdateWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Plans.Update(
		context.TODO(),
		"orgId",
		"id",
		m3ter.PlanUpdateParams{
			Code:           m3ter.F("JS!?Q0]r] ]$]"),
			Name:           m3ter.F("x"),
			PlanTemplateID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			AccountID:      m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Bespoke:        m3ter.F(true),
			CustomFields: m3ter.F(map[string]m3ter.PlanUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			MinimumSpend:                m3ter.F(0.000000),
			MinimumSpendBillInAdvance:   m3ter.F(true),
			MinimumSpendDescription:     m3ter.F("minimumSpendDescription"),
			Ordinal:                     m3ter.F(int64(0)),
			StandingCharge:              m3ter.F(0.000000),
			StandingChargeBillInAdvance: m3ter.F(true),
			StandingChargeDescription:   m3ter.F("standingChargeDescription"),
			Version:                     m3ter.F(int64(0)),
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

func TestPlanListWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Plans.List(
		context.TODO(),
		"orgId",
		m3ter.PlanListParams{
			AccountID: m3ter.F([]string{"string"}),
			IDs:       m3ter.F([]string{"string"}),
			NextToken: m3ter.F("nextToken"),
			PageSize:  m3ter.F(int64(1)),
			ProductID: m3ter.F("productId"),
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

func TestPlanDelete(t *testing.T) {
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
	)
	_, err := client.Plans.Delete(
		context.TODO(),
		"orgId",
		"id",
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
