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

func TestCounterPricingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CounterPricings.New(context.TODO(), m3ter.CounterPricingNewParams{
		CounterID: m3ter.F("x"),
		PricingBands: m3ter.F([]shared.PricingBandParam{{
			FixedPrice:   m3ter.F(0.000000),
			LowerLimit:   m3ter.F(0.000000),
			UnitPrice:    m3ter.F(0.000000),
			ID:           m3ter.F("id"),
			CreditTypeID: m3ter.F("creditTypeId"),
		}}),
		StartDate:                 m3ter.F(time.Now()),
		AccountingProductID:       m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Code:                      m3ter.F("JS!?Q0]r] ]$]"),
		Cumulative:                m3ter.F(true),
		Description:               m3ter.F("description"),
		EndDate:                   m3ter.F(time.Now()),
		PlanID:                    m3ter.F("planId"),
		PlanTemplateID:            m3ter.F("planTemplateId"),
		ProRateAdjustmentCredit:   m3ter.F(true),
		ProRateAdjustmentDebit:    m3ter.F(true),
		ProRateRunningTotal:       m3ter.F(true),
		RunningTotalBillInAdvance: m3ter.F(true),
		Version:                   m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterPricingGet(t *testing.T) {
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
	_, err := client.CounterPricings.Get(
		context.TODO(),
		"id",
		m3ter.CounterPricingGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCounterPricingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CounterPricings.Update(
		context.TODO(),
		"id",
		m3ter.CounterPricingUpdateParams{
			CounterID: m3ter.F("x"),
			PricingBands: m3ter.F([]shared.PricingBandParam{{
				FixedPrice:   m3ter.F(0.000000),
				LowerLimit:   m3ter.F(0.000000),
				UnitPrice:    m3ter.F(0.000000),
				ID:           m3ter.F("id"),
				CreditTypeID: m3ter.F("creditTypeId"),
			}}),
			StartDate:                 m3ter.F(time.Now()),
			AccountingProductID:       m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Code:                      m3ter.F("JS!?Q0]r] ]$]"),
			Cumulative:                m3ter.F(true),
			Description:               m3ter.F("description"),
			EndDate:                   m3ter.F(time.Now()),
			PlanID:                    m3ter.F("planId"),
			PlanTemplateID:            m3ter.F("planTemplateId"),
			ProRateAdjustmentCredit:   m3ter.F(true),
			ProRateAdjustmentDebit:    m3ter.F(true),
			ProRateRunningTotal:       m3ter.F(true),
			RunningTotalBillInAdvance: m3ter.F(true),
			Version:                   m3ter.F(int64(0)),
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

func TestCounterPricingListWithOptionalParams(t *testing.T) {
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
	_, err := client.CounterPricings.List(context.TODO(), m3ter.CounterPricingListParams{
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

func TestCounterPricingDelete(t *testing.T) {
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
	_, err := client.CounterPricings.Delete(
		context.TODO(),
		"id",
		m3ter.CounterPricingDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
