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

func TestAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.New(context.TODO(), m3ter.AccountNewParams{
		OrgID:        m3ter.F("orgId"),
		Code:         m3ter.F("JS!?Q0]r] ]$]"),
		EmailAddress: m3ter.F("dev@stainlessapi.com"),
		Name:         m3ter.F("x"),
		Address: m3ter.F(m3ter.AccountNewParamsAddress{
			AddressLine1: m3ter.F("addressLine1"),
			AddressLine2: m3ter.F("addressLine2"),
			AddressLine3: m3ter.F("addressLine3"),
			AddressLine4: m3ter.F("addressLine4"),
			Country:      m3ter.F("country"),
			Locality:     m3ter.F("locality"),
			PostCode:     m3ter.F("postCode"),
			Region:       m3ter.F("region"),
		}),
		AutoGenerateStatementMode: m3ter.F(m3ter.AccountNewParamsAutoGenerateStatementModeNone),
		BillEpoch:                 m3ter.F(time.Now()),
		ConfigData: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		CreditApplicationOrder: m3ter.F([]m3ter.AccountNewParamsCreditApplicationOrder{m3ter.AccountNewParamsCreditApplicationOrderPrepayment}),
		Currency:               m3ter.F("USD"),
		CustomFields: m3ter.F(map[string]m3ter.AccountNewParamsCustomFieldsUnion{
			"foo": shared.UnionString("string"),
		}),
		DaysBeforeBillDue:     m3ter.F(int64(1)),
		ParentAccountID:       m3ter.F("parentAccountId"),
		PurchaseOrderNumber:   m3ter.F("purchaseOrderNumber"),
		StatementDefinitionID: m3ter.F("statementDefinitionId"),
		Version:               m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGet(t *testing.T) {
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
	_, err := client.Accounts.Get(
		context.TODO(),
		"id",
		m3ter.AccountGetParams{
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

func TestAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Update(
		context.TODO(),
		"id",
		m3ter.AccountUpdateParams{
			OrgID:        m3ter.F("orgId"),
			Code:         m3ter.F("JS!?Q0]r] ]$]"),
			EmailAddress: m3ter.F("dev@stainlessapi.com"),
			Name:         m3ter.F("x"),
			Address: m3ter.F(m3ter.AccountUpdateParamsAddress{
				AddressLine1: m3ter.F("addressLine1"),
				AddressLine2: m3ter.F("addressLine2"),
				AddressLine3: m3ter.F("addressLine3"),
				AddressLine4: m3ter.F("addressLine4"),
				Country:      m3ter.F("country"),
				Locality:     m3ter.F("locality"),
				PostCode:     m3ter.F("postCode"),
				Region:       m3ter.F("region"),
			}),
			AutoGenerateStatementMode: m3ter.F(m3ter.AccountUpdateParamsAutoGenerateStatementModeNone),
			BillEpoch:                 m3ter.F(time.Now()),
			ConfigData: m3ter.F(map[string]interface{}{
				"foo": "bar",
			}),
			CreditApplicationOrder: m3ter.F([]m3ter.AccountUpdateParamsCreditApplicationOrder{m3ter.AccountUpdateParamsCreditApplicationOrderPrepayment}),
			Currency:               m3ter.F("USD"),
			CustomFields: m3ter.F(map[string]m3ter.AccountUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			DaysBeforeBillDue:     m3ter.F(int64(1)),
			ParentAccountID:       m3ter.F("parentAccountId"),
			PurchaseOrderNumber:   m3ter.F("purchaseOrderNumber"),
			StatementDefinitionID: m3ter.F("statementDefinitionId"),
			Version:               m3ter.F(int64(0)),
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

func TestAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.List(context.TODO(), m3ter.AccountListParams{
		OrgID:     m3ter.F("orgId"),
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

func TestAccountDelete(t *testing.T) {
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
	_, err := client.Accounts.Delete(
		context.TODO(),
		"id",
		m3ter.AccountDeleteParams{
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

func TestAccountEndDateBillingEntitiesWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.EndDateBillingEntities(
		context.TODO(),
		"id",
		m3ter.AccountEndDateBillingEntitiesParams{
			OrgID:           m3ter.F("orgId"),
			BillingEntities: m3ter.F([]m3ter.AccountEndDateBillingEntitiesParamsBillingEntity{m3ter.AccountEndDateBillingEntitiesParamsBillingEntityContract}),
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

func TestAccountGetChildrenWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.GetChildren(
		context.TODO(),
		"id",
		m3ter.AccountGetChildrenParams{
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

func TestAccountSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Search(context.TODO(), m3ter.AccountSearchParams{
		OrgID:        m3ter.F("orgId"),
		FromDocument: m3ter.F(int64(0)),
		Operator:     m3ter.F(m3ter.AccountSearchParamsOperatorAnd),
		PageSize:     m3ter.F(int64(1)),
		SearchQuery:  m3ter.F("searchQuery"),
		SortBy:       m3ter.F("sortBy"),
		SortOrder:    m3ter.F(m3ter.AccountSearchParamsSortOrderAsc),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
