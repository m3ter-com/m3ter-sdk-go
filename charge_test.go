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
)

func TestChargeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.New(context.TODO(), m3ter.ChargeNewParams{
		AccountID:              m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
		Currency:               m3ter.F("x"),
		EntityType:             m3ter.F(m3ter.ChargeNewParamsEntityTypeAdHoc),
		LineItemType:           m3ter.F(m3ter.ChargeNewParamsLineItemTypeBalanceFee),
		Name:                   m3ter.F("x"),
		ServicePeriodEndDate:   m3ter.F(time.Now()),
		ServicePeriodStartDate: m3ter.F(time.Now()),
		AccountingProductID:    m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Amount:                 m3ter.F(0.000000),
		BillDate:               m3ter.F("2022-01-04"),
		ContractID:             m3ter.F("contractId"),
		Description:            m3ter.F("description"),
		EntityID:               m3ter.F("entityId"),
		Notes:                  m3ter.F("notes"),
		UnitPrice:              m3ter.F(0.000000),
		Units:                  m3ter.F(0.000000),
		Version:                m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChargeGet(t *testing.T) {
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
	_, err := client.Charges.Get(
		context.TODO(),
		"id",
		m3ter.ChargeGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChargeUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.Update(
		context.TODO(),
		"id",
		m3ter.ChargeUpdateParams{
			AccountID:              m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
			Currency:               m3ter.F("x"),
			EntityType:             m3ter.F(m3ter.ChargeUpdateParamsEntityTypeAdHoc),
			LineItemType:           m3ter.F(m3ter.ChargeUpdateParamsLineItemTypeBalanceFee),
			Name:                   m3ter.F("x"),
			ServicePeriodEndDate:   m3ter.F(time.Now()),
			ServicePeriodStartDate: m3ter.F(time.Now()),
			AccountingProductID:    m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Amount:                 m3ter.F(0.000000),
			BillDate:               m3ter.F("2022-01-04"),
			ContractID:             m3ter.F("contractId"),
			Description:            m3ter.F("description"),
			EntityID:               m3ter.F("entityId"),
			Notes:                  m3ter.F("notes"),
			UnitPrice:              m3ter.F(0.000000),
			Units:                  m3ter.F(0.000000),
			Version:                m3ter.F(int64(0)),
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

func TestChargeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Charges.List(context.TODO(), m3ter.ChargeListParams{
		AccountID:  m3ter.F("accountId"),
		BillDate:   m3ter.F(time.Now()),
		EntityID:   m3ter.F("entityId"),
		EntityType: m3ter.F(m3ter.ChargeListParamsEntityTypeAdHoc),
		IDs:        m3ter.F([]string{"string"}),
		NextToken:  m3ter.F("nextToken"),
		PageSize:   m3ter.F(int64(1)),
		ScheduleID: m3ter.F("scheduleId"),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChargeDelete(t *testing.T) {
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
	_, err := client.Charges.Delete(
		context.TODO(),
		"id",
		m3ter.ChargeDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
