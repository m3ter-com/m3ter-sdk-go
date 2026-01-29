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

func TestLookupTableLookupTableRevisionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.LookupTables.LookupTableRevisions.New(
		context.TODO(),
		"lookupTableId",
		m3ter.LookupTableLookupTableRevisionNewParams{
			LookupTableRevisionRequest: m3ter.LookupTableRevisionRequestParam{
				Fields: m3ter.F([]m3ter.LookupTableRevisionRequestFieldParam{{
					Type: m3ter.F(m3ter.LookupTableRevisionRequestFieldsTypeString),
					Name: m3ter.F("lookupfield"),
				}, {
					Type: m3ter.F(m3ter.LookupTableRevisionRequestFieldsTypeString),
					Name: m3ter.F("lookupfield"),
				}}),
				Keys: m3ter.F([]string{"foo", "bar", "baz"}),
				Name: m3ter.F("x"),
				CustomFields: m3ter.F(map[string]m3ter.LookupTableRevisionRequestCustomFieldsUnionParam{
					"foo": shared.UnionString("string"),
				}),
				StartDate: m3ter.F(time.Now()),
				Version:   m3ter.F(int64(0)),
			},
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

func TestLookupTableLookupTableRevisionGet(t *testing.T) {
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
	_, err := client.LookupTables.LookupTableRevisions.Get(
		context.TODO(),
		"lookupTableId",
		"id",
		m3ter.LookupTableLookupTableRevisionGetParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLookupTableLookupTableRevisionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.LookupTables.LookupTableRevisions.Update(
		context.TODO(),
		"lookupTableId",
		"id",
		m3ter.LookupTableLookupTableRevisionUpdateParams{
			LookupTableRevisionRequest: m3ter.LookupTableRevisionRequestParam{
				Fields: m3ter.F([]m3ter.LookupTableRevisionRequestFieldParam{{
					Type: m3ter.F(m3ter.LookupTableRevisionRequestFieldsTypeString),
					Name: m3ter.F("lookupfield"),
				}, {
					Type: m3ter.F(m3ter.LookupTableRevisionRequestFieldsTypeString),
					Name: m3ter.F("lookupfield"),
				}}),
				Keys: m3ter.F([]string{"foo", "bar", "baz"}),
				Name: m3ter.F("x"),
				CustomFields: m3ter.F(map[string]m3ter.LookupTableRevisionRequestCustomFieldsUnionParam{
					"foo": shared.UnionString("string"),
				}),
				StartDate: m3ter.F(time.Now()),
				Version:   m3ter.F(int64(0)),
			},
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

func TestLookupTableLookupTableRevisionListWithOptionalParams(t *testing.T) {
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
	_, err := client.LookupTables.LookupTableRevisions.List(
		context.TODO(),
		"lookupTableId",
		m3ter.LookupTableLookupTableRevisionListParams{
			IDs:       m3ter.F([]string{"string"}),
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

func TestLookupTableLookupTableRevisionDelete(t *testing.T) {
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
	_, err := client.LookupTables.LookupTableRevisions.Delete(
		context.TODO(),
		"lookupTableId",
		"id",
		m3ter.LookupTableLookupTableRevisionDeleteParams{},
	)
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLookupTableLookupTableRevisionUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.LookupTables.LookupTableRevisions.UpdateStatus(
		context.TODO(),
		"lookupTableId",
		"id",
		m3ter.LookupTableLookupTableRevisionUpdateStatusParams{
			LookupTableRevisionStatusRequest: m3ter.LookupTableRevisionStatusRequestParam{
				Status:  m3ter.F(m3ter.LookupTableRevisionStatusRequestStatusDraft),
				Version: m3ter.F(int64(0)),
			},
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
