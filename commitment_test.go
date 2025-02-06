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

func TestCommitmentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Commitments.New(
		context.TODO(),
		"orgId",
		m3ter.CommitmentNewParams{
			AccountID:                  m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Amount:                     m3ter.F(1.000000),
			Currency:                   m3ter.F("x"),
			EndDate:                    m3ter.F(time.Now()),
			StartDate:                  m3ter.F(time.Now()),
			AccountingProductID:        m3ter.F("accountingProductId"),
			AmountFirstBill:            m3ter.F(0.000000),
			AmountPrePaid:              m3ter.F(0.000000),
			BillEpoch:                  m3ter.F(time.Now()),
			BillingInterval:            m3ter.F(int64(1)),
			BillingOffset:              m3ter.F(int64(0)),
			BillingPlanID:              m3ter.F("billingPlanId"),
			ChildBillingMode:           m3ter.F(m3ter.CommitmentNewParamsChildBillingModeParentSummary),
			CommitmentFeeBillInAdvance: m3ter.F(true),
			CommitmentFeeDescription:   m3ter.F("commitmentFeeDescription"),
			CommitmentUsageDescription: m3ter.F("commitmentUsageDescription"),
			ContractID:                 m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			FeeDates: m3ter.F([]m3ter.CommitmentNewParamsFeeDate{{
				Amount:                 m3ter.F(1.000000),
				Date:                   m3ter.F(time.Now()),
				ServicePeriodEndDate:   m3ter.F(time.Now()),
				ServicePeriodStartDate: m3ter.F(time.Now()),
			}}),
			LineItemTypes:           m3ter.F([]m3ter.CommitmentNewParamsLineItemType{m3ter.CommitmentNewParamsLineItemTypeStandingCharge}),
			OverageDescription:      m3ter.F("overageDescription"),
			OverageSurchargePercent: m3ter.F(0.000000),
			ProductIDs:              m3ter.F([]string{"string"}),
			SeparateOverageUsage:    m3ter.F(true),
			Version:                 m3ter.F(int64(0)),
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

func TestCommitmentGet(t *testing.T) {
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
	_, err := client.Commitments.Get(
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

func TestCommitmentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Commitments.Update(
		context.TODO(),
		"orgId",
		"id",
		m3ter.CommitmentUpdateParams{
			AccountID:                  m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Amount:                     m3ter.F(1.000000),
			Currency:                   m3ter.F("x"),
			EndDate:                    m3ter.F(time.Now()),
			StartDate:                  m3ter.F(time.Now()),
			AccountingProductID:        m3ter.F("accountingProductId"),
			AmountFirstBill:            m3ter.F(0.000000),
			AmountPrePaid:              m3ter.F(0.000000),
			BillEpoch:                  m3ter.F(time.Now()),
			BillingInterval:            m3ter.F(int64(1)),
			BillingOffset:              m3ter.F(int64(0)),
			BillingPlanID:              m3ter.F("billingPlanId"),
			ChildBillingMode:           m3ter.F(m3ter.CommitmentUpdateParamsChildBillingModeParentSummary),
			CommitmentFeeBillInAdvance: m3ter.F(true),
			CommitmentFeeDescription:   m3ter.F("commitmentFeeDescription"),
			CommitmentUsageDescription: m3ter.F("commitmentUsageDescription"),
			ContractID:                 m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			FeeDates: m3ter.F([]m3ter.CommitmentUpdateParamsFeeDate{{
				Amount:                 m3ter.F(1.000000),
				Date:                   m3ter.F(time.Now()),
				ServicePeriodEndDate:   m3ter.F(time.Now()),
				ServicePeriodStartDate: m3ter.F(time.Now()),
			}}),
			LineItemTypes:           m3ter.F([]m3ter.CommitmentUpdateParamsLineItemType{m3ter.CommitmentUpdateParamsLineItemTypeStandingCharge}),
			OverageDescription:      m3ter.F("overageDescription"),
			OverageSurchargePercent: m3ter.F(0.000000),
			ProductIDs:              m3ter.F([]string{"string"}),
			SeparateOverageUsage:    m3ter.F(true),
			Version:                 m3ter.F(int64(0)),
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

func TestCommitmentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Commitments.List(
		context.TODO(),
		"orgId",
		m3ter.CommitmentListParams{
			AccountID:    m3ter.F("accountId"),
			ContractID:   m3ter.F("contractId"),
			Date:         m3ter.F("date"),
			EndDateEnd:   m3ter.F("endDateEnd"),
			EndDateStart: m3ter.F("endDateStart"),
			IDs:          m3ter.F([]string{"string"}),
			NextToken:    m3ter.F("nextToken"),
			PageSize:     m3ter.F(int64(1)),
			ProductID:    m3ter.F("productId"),
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

func TestCommitmentDelete(t *testing.T) {
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
	_, err := client.Commitments.Delete(
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

func TestCommitmentSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Commitments.Search(
		context.TODO(),
		"orgId",
		m3ter.CommitmentSearchParams{
			FromDocument: m3ter.F(int64(0)),
			Operator:     m3ter.F(m3ter.CommitmentSearchParamsOperatorAnd),
			PageSize:     m3ter.F(int64(1)),
			SearchQuery:  m3ter.F("searchQuery"),
			SortBy:       m3ter.F("sortBy"),
			SortOrder:    m3ter.F(m3ter.CommitmentSearchParamsSortOrderAsc),
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
