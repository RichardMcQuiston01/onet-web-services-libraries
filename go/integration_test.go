package onet_test

import (
	"context"
	"os"
	"testing"

	onet "github.com/RichardMcQuiston01/onet-web-services-go"
)

// Integration tests hit the live O*NET API. They are skipped in short mode
// and when ONET_API_KEY is not set.
//
// Run with: ONET_API_KEY=<key> go test -run Integration -v

func integrationClient(t *testing.T) *onet.Client {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}
	key := os.Getenv("ONET_API_KEY")
	if key == "" {
		t.Skip("ONET_API_KEY not set")
	}
	return onet.NewClient(key)
}

func TestIntegrationAbout(t *testing.T) {
	c := integrationClient(t)
	got, err := c.About(context.Background())
	if err != nil {
		t.Fatalf("About: %v", err)
	}
	if got.APIVersion == "" {
		t.Error("APIVersion is empty")
	}
	t.Logf("api_version=%s onet_version=%s release_date=%s",
		got.APIVersion, got.OnetVersion, got.ReleaseDate)
}
