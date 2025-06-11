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

func TestUsageGetFailedIngestDownloadURLWithOptionalParams(t *testing.T) {
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
	_, err := client.Usage.GetFailedIngestDownloadURL(context.TODO(), m3ter.UsageGetFailedIngestDownloadURLParams{
		File: m3ter.F("file"),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUsageQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Usage.Query(context.TODO(), m3ter.UsageQueryParams{
		AccountIDs: m3ter.F([]string{"string"}),
		Aggregations: m3ter.F([]m3ter.UsageQueryParamsAggregation{{
			FieldCode: m3ter.F("x"),
			FieldType: m3ter.F(m3ter.UsageQueryParamsAggregationsFieldTypeDimension),
			Function:  m3ter.F(m3ter.UsageQueryParamsAggregationsFunctionSum),
			MeterID:   m3ter.F("x"),
		}}),
		DimensionFilters: m3ter.F([]m3ter.UsageQueryParamsDimensionFilter{{
			FieldCode: m3ter.F("x"),
			MeterID:   m3ter.F("x"),
			Values:    m3ter.F([]string{"string"}),
		}}),
		EndDate: m3ter.F(time.Now()),
		Groups: m3ter.F([]m3ter.DataExplorerGroupParam{{
			GroupType: m3ter.F(m3ter.DataExplorerGroupGroupTypeAccount),
		}}),
		Limit:     m3ter.F(int64(1)),
		MeterIDs:  m3ter.F([]string{"string"}),
		StartDate: m3ter.F(time.Now()),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUsageSubmit(t *testing.T) {
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
	_, err := client.Usage.Submit(context.TODO(), m3ter.UsageSubmitParams{
		SubmitMeasurementsRequest: m3ter.SubmitMeasurementsRequestParam{
			Measurements: m3ter.F([]m3ter.MeasurementRequestParam{{
				Account: m3ter.F("Acme Corp"),
				Meter:   m3ter.F("string"),
				Ts:      m3ter.F(time.Now()),
				Cost: m3ter.F(map[string]float64{
					"property1": 0.000000,
					"property2": 0.000000,
				}),
				Ets: m3ter.F(time.Now()),
				Income: m3ter.F(map[string]float64{
					"property1": 0.000000,
					"property2": 0.000000,
				}),
				Measure: m3ter.F(map[string]float64{
					"property1": 0.000000,
					"property2": 0.000000,
				}),
				Metadata: m3ter.F(map[string]string{
					"property1": "string",
					"property2": "string",
				}),
				Other: m3ter.F(map[string]string{
					"property1": "string",
					"property2": "string",
				}),
				Uid: m3ter.F("string"),
				What: m3ter.F(map[string]string{
					"property1": "string",
					"property2": "string",
				}),
				Where: m3ter.F(map[string]string{
					"property1": "string",
					"property2": "string",
				}),
				Who: m3ter.F(map[string]string{
					"property1": "string",
					"property2": "string",
				}),
			}}),
		},
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
