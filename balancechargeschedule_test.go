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

func TestBalanceChargeScheduleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.ChargeSchedules.New(
		context.TODO(),
		"balanceId",
		m3ter.BalanceChargeScheduleNewParams{
			BillFrequency:          m3ter.F(m3ter.BalanceChargeScheduleNewParamsBillFrequencyDaily),
			BillFrequencyInterval:  m3ter.F(int64(1)),
			BillInAdvance:          m3ter.F(true),
			ChargeDescription:      m3ter.F("x"),
			Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
			Currency:               m3ter.F("currency"),
			Name:                   m3ter.F("x"),
			ServicePeriodEndDate:   m3ter.F(time.Now()),
			ServicePeriodStartDate: m3ter.F(time.Now()),
			Amount:                 m3ter.F(0.000000),
			BillEpoch:              m3ter.F(time.Now()),
			CustomFields: m3ter.F(map[string]m3ter.BalanceChargeScheduleNewParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			UnitPrice: m3ter.F(0.000000),
			Units:     m3ter.F(0.000000),
			Version:   m3ter.F(int64(0)),
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

func TestBalanceChargeScheduleGet(t *testing.T) {
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
	_, err := client.Balances.ChargeSchedules.Get(
		context.TODO(),
		"balanceId",
		"id",
		m3ter.BalanceChargeScheduleGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBalanceChargeScheduleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.ChargeSchedules.Update(
		context.TODO(),
		"balanceId",
		"id",
		m3ter.BalanceChargeScheduleUpdateParams{
			BillFrequency:          m3ter.F(m3ter.BalanceChargeScheduleUpdateParamsBillFrequencyDaily),
			BillFrequencyInterval:  m3ter.F(int64(1)),
			BillInAdvance:          m3ter.F(true),
			ChargeDescription:      m3ter.F("x"),
			Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
			Currency:               m3ter.F("currency"),
			Name:                   m3ter.F("x"),
			ServicePeriodEndDate:   m3ter.F(time.Now()),
			ServicePeriodStartDate: m3ter.F(time.Now()),
			Amount:                 m3ter.F(0.000000),
			BillEpoch:              m3ter.F(time.Now()),
			CustomFields: m3ter.F(map[string]m3ter.BalanceChargeScheduleUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			UnitPrice: m3ter.F(0.000000),
			Units:     m3ter.F(0.000000),
			Version:   m3ter.F(int64(0)),
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

func TestBalanceChargeScheduleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.ChargeSchedules.List(
		context.TODO(),
		"balanceId",
		m3ter.BalanceChargeScheduleListParams{
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

func TestBalanceChargeScheduleDelete(t *testing.T) {
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
	_, err := client.Balances.ChargeSchedules.Delete(
		context.TODO(),
		"balanceId",
		"id",
		m3ter.BalanceChargeScheduleDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBalanceChargeSchedulePreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.ChargeSchedules.Preview(
		context.TODO(),
		"balanceId",
		m3ter.BalanceChargeSchedulePreviewParams{
			BillFrequency:          m3ter.F(m3ter.BalanceChargeSchedulePreviewParamsBillFrequencyDaily),
			BillFrequencyInterval:  m3ter.F(int64(1)),
			BillInAdvance:          m3ter.F(true),
			ChargeDescription:      m3ter.F("x"),
			Code:                   m3ter.F("S?oC\"$]C] ]]]]]5]"),
			Currency:               m3ter.F("currency"),
			Name:                   m3ter.F("x"),
			ServicePeriodEndDate:   m3ter.F(time.Now()),
			ServicePeriodStartDate: m3ter.F(time.Now()),
			NextToken:              m3ter.F("nextToken"),
			PageSize:               m3ter.F(int64(1)),
			Amount:                 m3ter.F(0.000000),
			BillEpoch:              m3ter.F(time.Now()),
			CustomFields: m3ter.F(map[string]m3ter.BalanceChargeSchedulePreviewParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			UnitPrice: m3ter.F(0.000000),
			Units:     m3ter.F(0.000000),
			Version:   m3ter.F(int64(0)),
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
