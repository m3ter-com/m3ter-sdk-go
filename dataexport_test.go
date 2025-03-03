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

func TestDataExportNewAdhocWithOptionalParams(t *testing.T) {
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
	_, err := client.DataExports.NewAdhoc(context.TODO(), m3ter.DataExportNewAdhocParams{
		OrgID: m3ter.F("orgId"),
		Body: m3ter.AdHocOperationalDataRequestParam{
			OperationalDataTypes: m3ter.F([]m3ter.AdHocOperationalDataRequestOperationalDataType{m3ter.AdHocOperationalDataRequestOperationalDataTypeBills}),
			SourceType:           m3ter.F(m3ter.AdHocOperationalDataRequestSourceTypeUsage),
			Version:              m3ter.F(int64(0)),
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
