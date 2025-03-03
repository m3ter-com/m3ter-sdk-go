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

func TestIntegrationConfigurationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.New(context.TODO(), m3ter.IntegrationConfigurationNewParams{
		OrgID: m3ter.F("orgId"),
		ConfigData: m3ter.F(map[string]interface{}{
			"foo": "bar",
		}),
		Credentials: m3ter.F(m3ter.IntegrationConfigurationNewParamsCredentials{
			Type:        m3ter.F(m3ter.IntegrationConfigurationNewParamsCredentialsTypeHTTPBasic),
			Destination: m3ter.F(m3ter.IntegrationConfigurationNewParamsCredentialsDestinationWebhook),
			Empty:       m3ter.F(true),
			Name:        m3ter.F("name"),
			Version:     m3ter.F(int64(0)),
		}),
		Destination:              m3ter.F("destination"),
		DestinationID:            m3ter.F("destinationId"),
		EntityID:                 m3ter.F("entityId"),
		EntityType:               m3ter.F("entityType"),
		IntegrationCredentialsID: m3ter.F("integrationCredentialsId"),
		Name:                     m3ter.F("name"),
		Version:                  m3ter.F(int64(0)),
	})
	if err != nil {
		var apierr *m3ter.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntegrationConfigurationGet(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.Get(
		context.TODO(),
		"id",
		m3ter.IntegrationConfigurationGetParams{
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

func TestIntegrationConfigurationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.Update(
		context.TODO(),
		"id",
		m3ter.IntegrationConfigurationUpdateParams{
			OrgID: m3ter.F("orgId"),
			ConfigData: m3ter.F(map[string]interface{}{
				"foo": "bar",
			}),
			Credentials: m3ter.F(m3ter.IntegrationConfigurationUpdateParamsCredentials{
				Type:        m3ter.F(m3ter.IntegrationConfigurationUpdateParamsCredentialsTypeHTTPBasic),
				Destination: m3ter.F(m3ter.IntegrationConfigurationUpdateParamsCredentialsDestinationWebhook),
				Empty:       m3ter.F(true),
				Name:        m3ter.F("name"),
				Version:     m3ter.F(int64(0)),
			}),
			Destination:              m3ter.F("destination"),
			DestinationID:            m3ter.F("destinationId"),
			EntityID:                 m3ter.F("entityId"),
			EntityType:               m3ter.F("entityType"),
			IntegrationCredentialsID: m3ter.F("integrationCredentialsId"),
			Name:                     m3ter.F("name"),
			Version:                  m3ter.F(int64(0)),
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

func TestIntegrationConfigurationListWithOptionalParams(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.List(context.TODO(), m3ter.IntegrationConfigurationListParams{
		OrgID:     m3ter.F("orgId"),
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

func TestIntegrationConfigurationDelete(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.Delete(
		context.TODO(),
		"id",
		m3ter.IntegrationConfigurationDeleteParams{
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

func TestIntegrationConfigurationEnable(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.Enable(
		context.TODO(),
		"id",
		m3ter.IntegrationConfigurationEnableParams{
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

func TestIntegrationConfigurationGetByEntityWithOptionalParams(t *testing.T) {
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
	_, err := client.IntegrationConfigurations.GetByEntity(
		context.TODO(),
		"entityType",
		m3ter.IntegrationConfigurationGetByEntityParams{
			OrgID:    m3ter.F("orgId"),
			EntityID: m3ter.F("entityId"),
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
