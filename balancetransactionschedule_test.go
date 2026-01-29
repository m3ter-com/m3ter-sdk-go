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

func TestBalanceTransactionScheduleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.TransactionSchedules.New(
		context.TODO(),
		"balanceId",
		m3ter.BalanceTransactionScheduleNewParams{
			ScheduleRequest: m3ter.ScheduleRequestParam{
				Amount:                 m3ter.F(0.000000),
				Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
				EndDate:                m3ter.F(time.Now()),
				Frequency:              m3ter.F(m3ter.ScheduleRequestFrequencyDaily),
				FrequencyInterval:      m3ter.F(int64(1)),
				Name:                   m3ter.F("x"),
				StartDate:              m3ter.F(time.Now()),
				TransactionDescription: m3ter.F("x"),
				TransactionTypeID:      m3ter.F("transactionTypeId"),
				CurrencyPaid:           m3ter.F("currencyPaid"),
				CustomFields: m3ter.F(map[string]m3ter.ScheduleRequestCustomFieldsUnionParam{
					"foo": shared.UnionString("string"),
				}),
				Paid:    m3ter.F(0.000000),
				Version: m3ter.F(int64(0)),
			},
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

func TestBalanceTransactionScheduleGet(t *testing.T) {
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
	_, err := client.Balances.TransactionSchedules.Get(
		context.TODO(),
		"balanceId",
		"id",
		m3ter.BalanceTransactionScheduleGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBalanceTransactionScheduleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.TransactionSchedules.Update(
		context.TODO(),
		"balanceId",
		"id",
		m3ter.BalanceTransactionScheduleUpdateParams{
			ScheduleRequest: m3ter.ScheduleRequestParam{
				Amount:                 m3ter.F(0.000000),
				Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
				EndDate:                m3ter.F(time.Now()),
				Frequency:              m3ter.F(m3ter.ScheduleRequestFrequencyDaily),
				FrequencyInterval:      m3ter.F(int64(1)),
				Name:                   m3ter.F("x"),
				StartDate:              m3ter.F(time.Now()),
				TransactionDescription: m3ter.F("x"),
				TransactionTypeID:      m3ter.F("transactionTypeId"),
				CurrencyPaid:           m3ter.F("currencyPaid"),
				CustomFields: m3ter.F(map[string]m3ter.ScheduleRequestCustomFieldsUnionParam{
					"foo": shared.UnionString("string"),
				}),
				Paid:    m3ter.F(0.000000),
				Version: m3ter.F(int64(0)),
			},
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

func TestBalanceTransactionScheduleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.TransactionSchedules.List(
		context.TODO(),
		"balanceId",
		m3ter.BalanceTransactionScheduleListParams{
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

func TestBalanceTransactionScheduleDelete(t *testing.T) {
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
	_, err := client.Balances.TransactionSchedules.Delete(
		context.TODO(),
		"balanceId",
		"id",
		m3ter.BalanceTransactionScheduleDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBalanceTransactionSchedulePreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.TransactionSchedules.Preview(
		context.TODO(),
		"balanceId",
		m3ter.BalanceTransactionSchedulePreviewParams{
			ScheduleRequest: m3ter.ScheduleRequestParam{
				Amount:                 m3ter.F(0.000000),
				Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
				EndDate:                m3ter.F(time.Now()),
				Frequency:              m3ter.F(m3ter.ScheduleRequestFrequencyDaily),
				FrequencyInterval:      m3ter.F(int64(1)),
				Name:                   m3ter.F("x"),
				StartDate:              m3ter.F(time.Now()),
				TransactionDescription: m3ter.F("x"),
				TransactionTypeID:      m3ter.F("transactionTypeId"),
				CurrencyPaid:           m3ter.F("currencyPaid"),
				CustomFields: m3ter.F(map[string]m3ter.ScheduleRequestCustomFieldsUnionParam{
					"foo": shared.UnionString("string"),
				}),
				Paid:    m3ter.F(0.000000),
				Version: m3ter.F(int64(0)),
			},
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
