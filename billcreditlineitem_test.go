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

func TestBillCreditLineItemNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.CreditLineItems.New(
		context.TODO(),
		"billId",
		m3ter.BillCreditLineItemNewParams{
			OrgID:                  m3ter.F("orgId"),
			Amount:                 m3ter.F(1.000000),
			Description:            m3ter.F("x"),
			ProductID:              m3ter.F("productId"),
			ReferencedBillID:       m3ter.F("referencedBillId"),
			ReferencedLineItemID:   m3ter.F("referencedLineItemId"),
			ServicePeriodEndDate:   m3ter.F(time.Now()),
			ServicePeriodStartDate: m3ter.F(time.Now()),
			CreditReasonID:         m3ter.F("creditReasonId"),
			LineItemType:           m3ter.F(m3ter.BillCreditLineItemNewParamsLineItemTypeStandingCharge),
			ReasonID:               m3ter.F("reasonId"),
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

func TestBillCreditLineItemGet(t *testing.T) {
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
	_, err := client.Bills.CreditLineItems.Get(
		context.TODO(),
		"billId",
		"id",
		m3ter.BillCreditLineItemGetParams{
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

func TestBillCreditLineItemUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.CreditLineItems.Update(
		context.TODO(),
		"billId",
		"id",
		m3ter.BillCreditLineItemUpdateParams{
			OrgID:                  m3ter.F("orgId"),
			Amount:                 m3ter.F(1.000000),
			Description:            m3ter.F("x"),
			ProductID:              m3ter.F("productId"),
			ReferencedBillID:       m3ter.F("referencedBillId"),
			ReferencedLineItemID:   m3ter.F("referencedLineItemId"),
			ServicePeriodEndDate:   m3ter.F(time.Now()),
			ServicePeriodStartDate: m3ter.F(time.Now()),
			CreditReasonID:         m3ter.F("creditReasonId"),
			LineItemType:           m3ter.F(m3ter.BillCreditLineItemUpdateParamsLineItemTypeStandingCharge),
			ReasonID:               m3ter.F("reasonId"),
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

func TestBillCreditLineItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Bills.CreditLineItems.List(
		context.TODO(),
		"billId",
		m3ter.BillCreditLineItemListParams{
			OrgID:     m3ter.F("orgId"),
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

func TestBillCreditLineItemDelete(t *testing.T) {
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
	_, err := client.Bills.CreditLineItems.Delete(
		context.TODO(),
		"billId",
		"id",
		m3ter.BillCreditLineItemDeleteParams{
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
