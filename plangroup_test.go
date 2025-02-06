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

func TestPlanGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PlanGroups.New(
		context.TODO(),
		"orgId",
		m3ter.PlanGroupNewParams{
			Currency:  m3ter.F("xxx"),
			Name:      m3ter.F("x"),
			AccountID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Code:      m3ter.F("JS!?Q0]r] ]$]"),
			CustomFields: m3ter.F(map[string]m3ter.PlanGroupNewParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			MinimumSpend:                      m3ter.F(0.000000),
			MinimumSpendAccountingProductID:   m3ter.F("minimumSpendAccountingProductId"),
			MinimumSpendBillInAdvance:         m3ter.F(true),
			MinimumSpendDescription:           m3ter.F("minimumSpendDescription"),
			StandingCharge:                    m3ter.F(0.000000),
			StandingChargeAccountingProductID: m3ter.F("standingChargeAccountingProductId"),
			StandingChargeBillInAdvance:       m3ter.F(true),
			StandingChargeDescription:         m3ter.F("standingChargeDescription"),
			Version:                           m3ter.F(int64(0)),
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

func TestPlanGroupGet(t *testing.T) {
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
	_, err := client.PlanGroups.Get(
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

func TestPlanGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PlanGroups.Update(
		context.TODO(),
		"orgId",
		"id",
		m3ter.PlanGroupUpdateParams{
			Currency:  m3ter.F("xxx"),
			Name:      m3ter.F("x"),
			AccountID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Code:      m3ter.F("JS!?Q0]r] ]$]"),
			CustomFields: m3ter.F(map[string]m3ter.PlanGroupUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			MinimumSpend:                      m3ter.F(0.000000),
			MinimumSpendAccountingProductID:   m3ter.F("minimumSpendAccountingProductId"),
			MinimumSpendBillInAdvance:         m3ter.F(true),
			MinimumSpendDescription:           m3ter.F("minimumSpendDescription"),
			StandingCharge:                    m3ter.F(0.000000),
			StandingChargeAccountingProductID: m3ter.F("standingChargeAccountingProductId"),
			StandingChargeBillInAdvance:       m3ter.F(true),
			StandingChargeDescription:         m3ter.F("standingChargeDescription"),
			Version:                           m3ter.F(int64(0)),
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

func TestPlanGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.PlanGroups.List(
		context.TODO(),
		"orgId",
		m3ter.PlanGroupListParams{
			AccountID: m3ter.F([]string{"string"}),
			IDs:       m3ter.F([]string{"string"}),
			NextToken: m3ter.F("nextToken"),
			PageSize:  m3ter.F(int64(1)),
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

func TestPlanGroupDelete(t *testing.T) {
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
	_, err := client.PlanGroups.Delete(
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
