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

func TestBalanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.New(
		context.TODO(),
		"orgId",
		m3ter.BalanceNewParams{
			AccountID:                  m3ter.F("x"),
			Currency:                   m3ter.F("x"),
			EndDate:                    m3ter.F(time.Now()),
			StartDate:                  m3ter.F(time.Now()),
			BalanceDrawDownDescription: m3ter.F("balanceDrawDownDescription"),
			Code:                       m3ter.F("JS!?Q0]r] ]$]"),
			Description:                m3ter.F("description"),
			LineItemTypes:              m3ter.F([]m3ter.BalanceNewParamsLineItemType{m3ter.BalanceNewParamsLineItemTypeStandingCharge}),
			Name:                       m3ter.F("name"),
			OverageDescription:         m3ter.F("overageDescription"),
			OverageSurchargePercent:    m3ter.F(0.000000),
			ProductIDs:                 m3ter.F([]string{"string"}),
			RolloverAmount:             m3ter.F(0.000000),
			RolloverEndDate:            m3ter.F(time.Now()),
			Version:                    m3ter.F(int64(0)),
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

func TestBalanceGet(t *testing.T) {
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
	_, err := client.Balances.Get(
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

func TestBalanceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.Update(
		context.TODO(),
		"orgId",
		"id",
		m3ter.BalanceUpdateParams{
			AccountID:                  m3ter.F("x"),
			Currency:                   m3ter.F("x"),
			EndDate:                    m3ter.F(time.Now()),
			StartDate:                  m3ter.F(time.Now()),
			BalanceDrawDownDescription: m3ter.F("balanceDrawDownDescription"),
			Code:                       m3ter.F("JS!?Q0]r] ]$]"),
			Description:                m3ter.F("description"),
			LineItemTypes:              m3ter.F([]m3ter.BalanceUpdateParamsLineItemType{m3ter.BalanceUpdateParamsLineItemTypeStandingCharge}),
			Name:                       m3ter.F("name"),
			OverageDescription:         m3ter.F("overageDescription"),
			OverageSurchargePercent:    m3ter.F(0.000000),
			ProductIDs:                 m3ter.F([]string{"string"}),
			RolloverAmount:             m3ter.F(0.000000),
			RolloverEndDate:            m3ter.F(time.Now()),
			Version:                    m3ter.F(int64(0)),
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

func TestBalanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Balances.List(
		context.TODO(),
		"orgId",
		m3ter.BalanceListParams{
			AccountID:    m3ter.F("accountId"),
			EndDateEnd:   m3ter.F("endDateEnd"),
			EndDateStart: m3ter.F("endDateStart"),
			NextToken:    m3ter.F("nextToken"),
			PageSize:     m3ter.F(int64(1)),
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

func TestBalanceDelete(t *testing.T) {
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
	_, err := client.Balances.Delete(
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
