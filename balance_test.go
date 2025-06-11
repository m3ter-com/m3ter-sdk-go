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
		option.WithOrgID("My Org ID"),
	)
	_, err := client.Balances.New(context.TODO(), m3ter.BalanceNewParams{
		AccountID:                       m3ter.F("x"),
		Currency:                        m3ter.F("x"),
		EndDate:                         m3ter.F(time.Now()),
		StartDate:                       m3ter.F(time.Now()),
		BalanceDrawDownDescription:      m3ter.F("balanceDrawDownDescription"),
		Code:                            m3ter.F("JS!?Q0]r] ]$]"),
		ConsumptionsAccountingProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		ContractID:                      m3ter.F("contractId"),
		CustomFields: m3ter.F(map[string]m3ter.BalanceNewParamsCustomFieldsUnion{
			"foo": shared.UnionString("string"),
		}),
		Description:             m3ter.F("description"),
		FeesAccountingProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		LineItemTypes:           m3ter.F([]m3ter.BalanceNewParamsLineItemType{m3ter.BalanceNewParamsLineItemTypeStandingCharge}),
		Name:                    m3ter.F("name"),
		OverageDescription:      m3ter.F("overageDescription"),
		OverageSurchargePercent: m3ter.F(0.000000),
		ProductIDs:              m3ter.F([]string{"string"}),
		RolloverAmount:          m3ter.F(0.000000),
		RolloverEndDate:         m3ter.F(time.Now()),
		Version:                 m3ter.F(int64(0)),
	})
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
		option.WithOrgID("My Org ID"),
	)
	_, err := client.Balances.Get(
		context.TODO(),
		"id",
		m3ter.BalanceGetParams{},
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
		option.WithOrgID("My Org ID"),
	)
	_, err := client.Balances.Update(
		context.TODO(),
		"id",
		m3ter.BalanceUpdateParams{
			AccountID:                       m3ter.F("x"),
			Currency:                        m3ter.F("x"),
			EndDate:                         m3ter.F(time.Now()),
			StartDate:                       m3ter.F(time.Now()),
			BalanceDrawDownDescription:      m3ter.F("balanceDrawDownDescription"),
			Code:                            m3ter.F("JS!?Q0]r] ]$]"),
			ConsumptionsAccountingProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			ContractID:                      m3ter.F("contractId"),
			CustomFields: m3ter.F(map[string]m3ter.BalanceUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			Description:             m3ter.F("description"),
			FeesAccountingProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			LineItemTypes:           m3ter.F([]m3ter.BalanceUpdateParamsLineItemType{m3ter.BalanceUpdateParamsLineItemTypeStandingCharge}),
			Name:                    m3ter.F("name"),
			OverageDescription:      m3ter.F("overageDescription"),
			OverageSurchargePercent: m3ter.F(0.000000),
			ProductIDs:              m3ter.F([]string{"string"}),
			RolloverAmount:          m3ter.F(0.000000),
			RolloverEndDate:         m3ter.F(time.Now()),
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
		option.WithOrgID("My Org ID"),
	)
	_, err := client.Balances.List(context.TODO(), m3ter.BalanceListParams{
		AccountID:    m3ter.F("accountId"),
		Contract:     m3ter.F("contract"),
		EndDateEnd:   m3ter.F("endDateEnd"),
		EndDateStart: m3ter.F("endDateStart"),
		NextToken:    m3ter.F("nextToken"),
		PageSize:     m3ter.F(int64(1)),
	})
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
		option.WithOrgID("My Org ID"),
	)
	_, err := client.Balances.Delete(
		context.TODO(),
		"id",
		m3ter.BalanceDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
