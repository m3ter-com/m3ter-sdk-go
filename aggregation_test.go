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

func TestAggregationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Aggregations.New(context.TODO(), m3ter.AggregationNewParams{
		Aggregation:         m3ter.F(m3ter.AggregationNewParamsAggregationSum),
		MeterID:             m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Name:                m3ter.F("x"),
		QuantityPerUnit:     m3ter.F(1.000000),
		Rounding:            m3ter.F(m3ter.AggregationNewParamsRoundingUp),
		TargetField:         m3ter.F("x"),
		Unit:                m3ter.F("x"),
		AccountingProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		Code:                m3ter.F("example_code"),
		CustomFields: m3ter.F(map[string]m3ter.AggregationNewParamsCustomFieldsUnion{
			"foo": shared.UnionString("string"),
		}),
		CustomSql:       m3ter.F("customSql"),
		DefaultValue:    m3ter.F(0.000000),
		SegmentedFields: m3ter.F([]string{"string"}),
		Segments: m3ter.F([]map[string]string{{
			"foo": "string",
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

func TestAggregationGet(t *testing.T) {
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
	_, err := client.Aggregations.Get(
		context.TODO(),
		"id",
		m3ter.AggregationGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAggregationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Aggregations.Update(
		context.TODO(),
		"id",
		m3ter.AggregationUpdateParams{
			Aggregation:         m3ter.F(m3ter.AggregationUpdateParamsAggregationSum),
			MeterID:             m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Name:                m3ter.F("x"),
			QuantityPerUnit:     m3ter.F(1.000000),
			Rounding:            m3ter.F(m3ter.AggregationUpdateParamsRoundingUp),
			TargetField:         m3ter.F("x"),
			Unit:                m3ter.F("x"),
			AccountingProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Code:                m3ter.F("example_code"),
			CustomFields: m3ter.F(map[string]m3ter.AggregationUpdateParamsCustomFieldsUnion{
				"foo": shared.UnionString("string"),
			}),
			CustomSql:       m3ter.F("customSql"),
			DefaultValue:    m3ter.F(0.000000),
			SegmentedFields: m3ter.F([]string{"string"}),
			Segments: m3ter.F([]map[string]string{{
				"foo": "string",
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

func TestAggregationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Aggregations.List(context.TODO(), m3ter.AggregationListParams{
		Codes:     m3ter.F([]string{"string"}),
		IDs:       m3ter.F([]string{"string"}),
		NextToken: m3ter.F("nextToken"),
		PageSize:  m3ter.F(int64(1)),
		ProductID: m3ter.F([]string{"string"}),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAggregationDelete(t *testing.T) {
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
	_, err := client.Aggregations.Delete(
		context.TODO(),
		"id",
		m3ter.AggregationDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
