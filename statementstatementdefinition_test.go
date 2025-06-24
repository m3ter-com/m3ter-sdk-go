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

func TestStatementStatementDefinitionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Statements.StatementDefinitions.New(context.TODO(), m3ter.StatementStatementDefinitionNewParams{
		AggregationFrequency: m3ter.F(m3ter.StatementStatementDefinitionNewParamsAggregationFrequencyDay),
		Dimensions: m3ter.F([]m3ter.StatementStatementDefinitionNewParamsDimension{{
			Filter:     m3ter.F([]string{"string"}),
			Name:       m3ter.F("x"),
			Attributes: m3ter.F([]string{"string"}),
			MeterID:    m3ter.F("meterId"),
		}}),
		IncludePricePerUnit: m3ter.F(true),
		Measures: m3ter.F([]m3ter.StatementStatementDefinitionNewParamsMeasure{{
			Aggregations: m3ter.F([]m3ter.StatementStatementDefinitionNewParamsMeasuresAggregation{m3ter.StatementStatementDefinitionNewParamsMeasuresAggregationSum}),
			MeterID:      m3ter.F("meterId"),
			Name:         m3ter.F("name"),
		}}),
		Name:    m3ter.F("name"),
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

func TestStatementStatementDefinitionGet(t *testing.T) {
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
	_, err := client.Statements.StatementDefinitions.Get(
		context.TODO(),
		"id",
		m3ter.StatementStatementDefinitionGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStatementStatementDefinitionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Statements.StatementDefinitions.Update(
		context.TODO(),
		"id",
		m3ter.StatementStatementDefinitionUpdateParams{
			AggregationFrequency: m3ter.F(m3ter.StatementStatementDefinitionUpdateParamsAggregationFrequencyDay),
			Dimensions: m3ter.F([]m3ter.StatementStatementDefinitionUpdateParamsDimension{{
				Filter:     m3ter.F([]string{"string"}),
				Name:       m3ter.F("x"),
				Attributes: m3ter.F([]string{"string"}),
				MeterID:    m3ter.F("meterId"),
			}}),
			IncludePricePerUnit: m3ter.F(true),
			Measures: m3ter.F([]m3ter.StatementStatementDefinitionUpdateParamsMeasure{{
				Aggregations: m3ter.F([]m3ter.StatementStatementDefinitionUpdateParamsMeasuresAggregation{m3ter.StatementStatementDefinitionUpdateParamsMeasuresAggregationSum}),
				MeterID:      m3ter.F("meterId"),
				Name:         m3ter.F("name"),
			}}),
			Name:    m3ter.F("name"),
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

func TestStatementStatementDefinitionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Statements.StatementDefinitions.List(context.TODO(), m3ter.StatementStatementDefinitionListParams{
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

func TestStatementStatementDefinitionDelete(t *testing.T) {
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
	_, err := client.Statements.StatementDefinitions.Delete(
		context.TODO(),
		"id",
		m3ter.StatementStatementDefinitionDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
