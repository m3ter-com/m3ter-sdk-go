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
		Account: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsAccountUnion{
			"foo": shared.UnionString("string"),
		}),
		AccountPlan: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsAccountPlanUnion{
			"foo": shared.UnionString("string"),
		}),
		Aggregation: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsAggregationUnion{
			"foo": shared.UnionString("string"),
		}),
		CompoundAggregation: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsCompoundAggregationUnion{
			"foo": shared.UnionString("string"),
		}),
		Contract: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsContractUnion{
			"foo": shared.UnionString("string"),
		}),
		Meter: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsMeterUnion{
			"foo": shared.UnionString("string"),
		}),
		Organization: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsOrganizationUnion{
			"foo": shared.UnionString("string"),
		}),
		Plan: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsPlanUnion{
			"foo": shared.UnionString("string"),
		}),
		PlanTemplate: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsPlanTemplateUnion{
			"foo": shared.UnionString("string"),
		}),
		Product: m3ter.F(map[string]m3ter.CustomFieldUpdateParamsProductUnion{
			"foo": shared.UnionString("string"),
		}),
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
