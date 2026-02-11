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
)

func TestPlanTemplateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PlanTemplates.New(context.TODO(), m3ter.PlanTemplateNewParams{
		BillFrequency:               m3ter.F(m3ter.PlanTemplateNewParamsBillFrequencyDaily),
		Currency:                    m3ter.F("USD"),
		Name:                        m3ter.F("string"),
		ProductID:                   m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		StandingCharge:              m3ter.F(0.000000),
		BillFrequencyInterval:       m3ter.F(int64(1)),
		Code:                        m3ter.F("string"),
		CustomFields:                m3ter.F(map[string]m3ter.PlanTemplateNewParamsCustomFieldsUnion{}),
		MinimumSpend:                m3ter.F(0.000000),
		MinimumSpendBillInAdvance:   m3ter.F(true),
		MinimumSpendDescription:     m3ter.F("string"),
		Ordinal:                     m3ter.F(int64(0)),
		StandingChargeBillInAdvance: m3ter.F(true),
		StandingChargeDescription:   m3ter.F("string"),
		StandingChargeInterval:      m3ter.F(int64(1)),
		StandingChargeOffset:        m3ter.F(int64(364)),
		Version:                     m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanTemplateGet(t *testing.T) {
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
	_, err := client.PlanTemplates.Get(
		context.TODO(),
		"id",
		m3ter.PlanTemplateGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanTemplateUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PlanTemplates.Update(
		context.TODO(),
		"id",
		m3ter.PlanTemplateUpdateParams{
			BillFrequency:               m3ter.F(m3ter.PlanTemplateUpdateParamsBillFrequencyDaily),
			Currency:                    m3ter.F("USD"),
			Name:                        m3ter.F("string"),
			ProductID:                   m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			StandingCharge:              m3ter.F(0.000000),
			BillFrequencyInterval:       m3ter.F(int64(1)),
			Code:                        m3ter.F("string"),
			CustomFields:                m3ter.F(map[string]m3ter.PlanTemplateUpdateParamsCustomFieldsUnion{}),
			MinimumSpend:                m3ter.F(0.000000),
			MinimumSpendBillInAdvance:   m3ter.F(true),
			MinimumSpendDescription:     m3ter.F("string"),
			Ordinal:                     m3ter.F(int64(0)),
			StandingChargeBillInAdvance: m3ter.F(true),
			StandingChargeDescription:   m3ter.F("string"),
			StandingChargeInterval:      m3ter.F(int64(1)),
			StandingChargeOffset:        m3ter.F(int64(364)),
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

func TestPlanTemplateListWithOptionalParams(t *testing.T) {
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
	_, err := client.PlanTemplates.List(context.TODO(), m3ter.PlanTemplateListParams{
		IDs:       m3ter.F([]string{"string"}),
		NextToken: m3ter.F("nextToken"),
		PageSize:  m3ter.F(int64(1)),
		ProductID: m3ter.F("productId"),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPlanTemplateDelete(t *testing.T) {
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
	_, err := client.PlanTemplates.Delete(
		context.TODO(),
		"id",
		m3ter.PlanTemplateDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
