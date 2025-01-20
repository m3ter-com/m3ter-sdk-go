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

func TestMeterNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Meters.New(
		context.TODO(),
		"orgId",
		m3ter.MeterNewParams{
			Code: m3ter.F("JS!?Q0]r] ]$]"),
			DataFields: m3ter.F([]m3ter.MeterNewParamsDataField{{
				Category: m3ter.F(m3ter.MeterNewParamsDataFieldsCategoryWho),
				Code:     m3ter.F("{1{}}_"),
				Name:     m3ter.F("x"),
				Unit:     m3ter.F("x"),
			}}),
			DerivedFields: m3ter.F([]m3ter.MeterNewParamsDerivedField{{
				Category: m3ter.F(m3ter.MeterNewParamsDerivedFieldsCategoryWho),
				Code:     m3ter.F("{1{}}_"),
				Name:     m3ter.F("x"),
				Unit:     m3ter.F("x"),
			}}),
			Name: m3ter.F("x"),
			CustomFields: m3ter.F(map[string]interface{}{
				"foo": "bar",
			}),
			GroupID:   m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			ProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Version:   m3ter.F(int64(0)),
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

func TestMeterGet(t *testing.T) {
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
	_, err := client.Meters.Get(
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

func TestMeterUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Meters.Update(
		context.TODO(),
		"orgId",
		"id",
		m3ter.MeterUpdateParams{
			Code: m3ter.F("JS!?Q0]r] ]$]"),
			DataFields: m3ter.F([]m3ter.MeterUpdateParamsDataField{{
				Category: m3ter.F(m3ter.MeterUpdateParamsDataFieldsCategoryWho),
				Code:     m3ter.F("{1{}}_"),
				Name:     m3ter.F("x"),
				Unit:     m3ter.F("x"),
			}}),
			DerivedFields: m3ter.F([]m3ter.MeterUpdateParamsDerivedField{{
				Category: m3ter.F(m3ter.MeterUpdateParamsDerivedFieldsCategoryWho),
				Code:     m3ter.F("{1{}}_"),
				Name:     m3ter.F("x"),
				Unit:     m3ter.F("x"),
			}}),
			Name: m3ter.F("x"),
			CustomFields: m3ter.F(map[string]interface{}{
				"foo": "bar",
			}),
			GroupID:   m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			ProductID: m3ter.F("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			Version:   m3ter.F(int64(0)),
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

func TestMeterListWithOptionalParams(t *testing.T) {
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
	_, err := client.Meters.List(
		context.TODO(),
		"orgId",
		m3ter.MeterListParams{
			Codes:     m3ter.F([]string{"string"}),
			IDs:       m3ter.F([]string{"string"}),
			NextToken: m3ter.F("nextToken"),
			PageSize:  m3ter.F(int64(1)),
			ProductID: m3ter.F([]interface{}{map[string]interface{}{}}),
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
