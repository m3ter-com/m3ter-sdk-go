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

func TestCustomFieldGet(t *testing.T) {
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
	_, err := client.CustomFields.Get(context.TODO(), m3ter.CustomFieldGetParams{})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomFieldUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CustomFields.Update(context.TODO(), m3ter.CustomFieldUpdateParams{
		Account: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsAccountUnion{}),
		AccountPlan: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsAccountPlanUnion{
			"New CF Test": shared.UnionString("Test Value"),
		}),
		Aggregation:         m3ter.F(map[string]m3ter.CustomFieldUpdateParamsAggregationUnion{}),
		CompoundAggregation: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsCompoundAggregationUnion{}),
		Contract: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsContractUnion{
			"foo": shared.UnionString("string"),
		}),
		Meter: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsMeterUnion{}),
		Organization: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsOrganizationUnion{
			"Org Example 2": shared.UnionString("Sample text 2."),
			"Org Example 1": shared.UnionString("Sample text 1."),
		}),
		Plan:         m3ter.F(map[string]m3ter.CustomFieldUpdateParamsPlanUnion{}),
		PlanTemplate: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsPlanTemplateUnion{}),
		Product: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsProductUnion{
			"Product CF Example": shared.UnionFloat(42.000000),
		}),
		Version: m3ter.F(int64(6)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
