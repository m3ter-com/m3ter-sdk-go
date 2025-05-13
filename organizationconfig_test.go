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

func TestOrganizationConfigGet(t *testing.T) {
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
	_, err := client.OrganizationConfig.Get(context.TODO(), m3ter.OrganizationConfigGetParams{
		OrgID: m3ter.F("orgId"),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.OrganizationConfig.Update(context.TODO(), m3ter.OrganizationConfigUpdateParams{
		OrgID: m3ter.F("orgId"),
		OrganizationConfigRequest: m3ter.OrganizationConfigRequestParam{
			Currency:                        m3ter.F("USD"),
			DayEpoch:                        m3ter.F("2022-01-01"),
			DaysBeforeBillDue:               m3ter.F(int64(1)),
			MonthEpoch:                      m3ter.F("2022-01-01"),
			Timezone:                        m3ter.F("UTC"),
			WeekEpoch:                       m3ter.F("2022-01-04"),
			YearEpoch:                       m3ter.F("2022-01-01"),
			AutoApproveBillsGracePeriod:     m3ter.F(int64(2)),
			AutoApproveBillsGracePeriodUnit: m3ter.F("DAYS"),
			AutoGenerateStatementMode:       m3ter.F(m3ter.OrganizationConfigRequestAutoGenerateStatementModeNone),
			BillPrefix:                      m3ter.F("Bill-"),
			CommitmentFeeBillInAdvance:      m3ter.F(true),
			ConsolidateBills:                m3ter.F(true),
			CreditApplicationOrder:          m3ter.F([]m3ter.OrganizationConfigRequestCreditApplicationOrder{m3ter.OrganizationConfigRequestCreditApplicationOrderPrepayment}),
			CurrencyConversions: m3ter.F([]shared.CurrencyConversionParam{{
				From:       m3ter.F("EUR"),
				To:         m3ter.F("USD"),
				Multiplier: m3ter.F(1.120000),
			}}),
			DefaultStatementDefinitionID: m3ter.F("defaultStatementDefinitionId"),
			ExternalInvoiceDate:          m3ter.F("LAST_DAY_OF_ARREARS"),
			MinimumSpendBillInAdvance:    m3ter.F(true),
			ScheduledBillInterval:        m3ter.F(0.000000),
			SequenceStartNumber:          m3ter.F(int64(1000)),
			StandingChargeBillInAdvance:  m3ter.F(true),
			SuppressedEmptyBills:         m3ter.F(true),
			Version:                      m3ter.F(int64(0)),
		},
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
