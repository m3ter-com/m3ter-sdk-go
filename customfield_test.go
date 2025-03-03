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
	_, err := client.CustomFields.Get(context.TODO(), m3ter.CustomFieldGetParams{
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
		OrgID: m3ter.F("orgId"),
		Account: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		AccountPlan: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		Aggregation: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		CompoundAggregation: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		Meter: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		Organization: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		Plan: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		PlanTemplate: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		Product: m3ter.F(map[string]interface{}{
			"foo": "bar",
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
