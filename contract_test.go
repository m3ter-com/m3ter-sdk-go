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

func TestContractNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.New(context.TODO(), m3ter.ContractNewParams{
		AccountID:                 m3ter.F("x"),
		EndDate:                   m3ter.F(time.Now()),
		Name:                      m3ter.F("x"),
		StartDate:                 m3ter.F(time.Now()),
		ApplyContractPeriodLimits: m3ter.F(true),
		BillGroupingKeyID:         m3ter.F("billGroupingKeyId"),
		Code:                      m3ter.F("S?oC\"$]C] ]]]]]5]"),
		CustomFields: m3ter.F(map[string]m3ter.ContractNewParamsCustomFieldsUnion{
			"foo": shared.UnionString("string"),
		}),
		Description:         m3ter.F("description"),
		PurchaseOrderNumber: m3ter.F("purchaseOrderNumber"),
		UsageFilters: m3ter.F([]m3ter.ContractNewParamsUsageFilter{{
			DimensionCode: m3ter.F("x"),
			Mode:          m3ter.F(m3ter.ContractNewParamsUsageFiltersModeInclude),
			Value:         m3ter.F("x"),
		}}),
		Version: m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractGet(t *testing.T) {
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
	_, err := client.Contracts.Get(
		context.TODO(),
		"id",
		m3ter.ContractGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.Update(
		context.TODO(),
		"id",
		m3ter.ContractUpdateParams{
			AccountID:                 m3ter.F("x"),
			EndDate:                   m3ter.F(time.Now()),
			Name:                      m3ter.F("x"),
			StartDate:                 m3ter.F(time.Now()),
			ApplyContractPeriodLimits: m3ter.F(true),
			BillGroupingKeyID:         m3ter.F("billGroupingKeyId"),
			Code:                      m3ter.F("S?oC\"$]C] ]]]]]5]"),
			CustomFields: m3ter.F(map[string]m3ter.ContractUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			Description:         m3ter.F("description"),
			PurchaseOrderNumber: m3ter.F("purchaseOrderNumber"),
			UsageFilters: m3ter.F([]m3ter.ContractUpdateParamsUsageFilter{{
				DimensionCode: m3ter.F("x"),
				Mode:          m3ter.F(m3ter.ContractUpdateParamsUsageFiltersModeInclude),
				Value:         m3ter.F("x"),
			}}),
			Version: m3ter.F(int64(0)),
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

func TestContractListWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.List(context.TODO(), m3ter.ContractListParams{
		AccountID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Codes:     m3ter.F([]string{"string"}),
		IDs:       m3ter.F([]string{"string"}),
		NextToken: m3ter.F("nextToken"),
		PageSize:  m3ter.F(int64(1)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractDelete(t *testing.T) {
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
	_, err := client.Contracts.Delete(
		context.TODO(),
		"id",
		m3ter.ContractDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContractEndDateBillingEntitiesWithOptionalParams(t *testing.T) {
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
	_, err := client.Contracts.EndDateBillingEntities(
		context.TODO(),
		"id",
		m3ter.ContractEndDateBillingEntitiesParams{
			BillingEntities: m3ter.F([]m3ter.ContractEndDateBillingEntitiesParamsBillingEntity{m3ter.ContractEndDateBillingEntitiesParamsBillingEntityContract}),
			EndDate:         m3ter.F(time.Now()),
			ApplyToChildren: m3ter.F(true),
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
