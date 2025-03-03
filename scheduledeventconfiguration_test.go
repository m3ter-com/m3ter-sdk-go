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

func TestScheduledEventConfigurationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ScheduledEventConfigurations.New(context.TODO(), m3ter.ScheduledEventConfigurationNewParams{
		OrgID:   m3ter.F("orgId"),
		Entity:  m3ter.F("Bill"),
		Field:   m3ter.F("endDate"),
		Name:    m3ter.F("scheduled.bill.enddateEvent"),
		Offset:  m3ter.F(int64(5)),
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

func TestScheduledEventConfigurationGet(t *testing.T) {
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
	_, err := client.ScheduledEventConfigurations.Get(
		context.TODO(),
		"id",
		m3ter.ScheduledEventConfigurationGetParams{
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

func TestScheduledEventConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ScheduledEventConfigurations.Update(
		context.TODO(),
		"id",
		m3ter.ScheduledEventConfigurationUpdateParams{
			OrgID:   m3ter.F("orgId"),
			Entity:  m3ter.F("Bill"),
			Field:   m3ter.F("endDate"),
			Name:    m3ter.F("scheduled.bill.enddateEvent"),
			Offset:  m3ter.F(int64(5)),
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

func TestScheduledEventConfigurationListWithOptionalParams(t *testing.T) {
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
	_, err := client.ScheduledEventConfigurations.List(context.TODO(), m3ter.ScheduledEventConfigurationListParams{
		OrgID:     m3ter.F("orgId"),
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

func TestScheduledEventConfigurationDelete(t *testing.T) {
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
	_, err := client.ScheduledEventConfigurations.Delete(
		context.TODO(),
		"id",
		m3ter.ScheduledEventConfigurationDeleteParams{
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
