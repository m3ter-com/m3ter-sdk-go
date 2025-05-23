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

func TestPricingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pricings.New(context.TODO(), m3ter.PricingNewParams{
		OrgID: m3ter.F("orgId"),
		PricingBands: m3ter.F([]shared.PricingBandParam{{
			FixedPrice:   m3ter.F(0.000000),
			LowerLimit:   m3ter.F(0.000000),
			UnitPrice:    m3ter.F(0.000000),
			ID:           m3ter.F("id"),
			CreditTypeID: m3ter.F("creditTypeId"),
		}}),
		StartDate:                 m3ter.F(time.Now()),
		AccountingProductID:       m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		AggregationID:             m3ter.F("aggregationId"),
		Code:                      m3ter.F("JS!?Q0]r] ]$]"),
		CompoundAggregationID:     m3ter.F("compoundAggregationId"),
		Cumulative:                m3ter.F(true),
		Description:               m3ter.F("description"),
		EndDate:                   m3ter.F(time.Now()),
		MinimumSpend:              m3ter.F(0.000000),
		MinimumSpendBillInAdvance: m3ter.F(true),
		MinimumSpendDescription:   m3ter.F("minimumSpendDescription"),
		OveragePricingBands: m3ter.F([]shared.PricingBandParam{{
			FixedPrice:   m3ter.F(0.000000),
			LowerLimit:   m3ter.F(0.000000),
			UnitPrice:    m3ter.F(0.000000),
			ID:           m3ter.F("id"),
			CreditTypeID: m3ter.F("creditTypeId"),
		}}),
		PlanID:         m3ter.F("planId"),
		PlanTemplateID: m3ter.F("planTemplateId"),
		Segment: m3ter.F(map[string]string{
			"foo": "string",
		}),
		TiersSpanPlan: m3ter.F(true),
		Type:          m3ter.F(m3ter.PricingNewParamsTypeDebit),
		Version:       m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPricingGet(t *testing.T) {
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
	_, err := client.Pricings.Get(
		context.TODO(),
		"id",
		m3ter.PricingGetParams{
			OrgID: m3ter.F("orgId"),
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

func TestPricingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pricings.Update(
		context.TODO(),
		"id",
		m3ter.PricingUpdateParams{
			OrgID: m3ter.F("orgId"),
			PricingBands: m3ter.F([]shared.PricingBandParam{{
				FixedPrice:   m3ter.F(0.000000),
				LowerLimit:   m3ter.F(0.000000),
				UnitPrice:    m3ter.F(0.000000),
				ID:           m3ter.F("id"),
				CreditTypeID: m3ter.F("creditTypeId"),
			}}),
			StartDate:                 m3ter.F(time.Now()),
			AccountingProductID:       m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			AggregationID:             m3ter.F("aggregationId"),
			Code:                      m3ter.F("JS!?Q0]r] ]$]"),
			CompoundAggregationID:     m3ter.F("compoundAggregationId"),
			Cumulative:                m3ter.F(true),
			Description:               m3ter.F("description"),
			EndDate:                   m3ter.F(time.Now()),
			MinimumSpend:              m3ter.F(0.000000),
			MinimumSpendBillInAdvance: m3ter.F(true),
			MinimumSpendDescription:   m3ter.F("minimumSpendDescription"),
			OveragePricingBands: m3ter.F([]shared.PricingBandParam{{
				FixedPrice:   m3ter.F(0.000000),
				LowerLimit:   m3ter.F(0.000000),
				UnitPrice:    m3ter.F(0.000000),
				ID:           m3ter.F("id"),
				CreditTypeID: m3ter.F("creditTypeId"),
			}}),
			PlanID:         m3ter.F("planId"),
			PlanTemplateID: m3ter.F("planTemplateId"),
			Segment: m3ter.F(map[string]string{
				"foo": "string",
			}),
			TiersSpanPlan: m3ter.F(true),
			Type:          m3ter.F(m3ter.PricingUpdateParamsTypeDebit),
			Version:       m3ter.F(int64(0)),
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

func TestPricingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pricings.List(context.TODO(), m3ter.PricingListParams{
		OrgID:          m3ter.F("orgId"),
		Date:           m3ter.F("date"),
		IDs:            m3ter.F([]string{"string"}),
		NextToken:      m3ter.F("nextToken"),
		PageSize:       m3ter.F(int64(1)),
		PlanID:         m3ter.F("planId"),
		PlanTemplateID: m3ter.F("planTemplateId"),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPricingDelete(t *testing.T) {
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
	_, err := client.Pricings.Delete(
		context.TODO(),
		"id",
		m3ter.PricingDeleteParams{
			OrgID: m3ter.F("orgId"),
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
