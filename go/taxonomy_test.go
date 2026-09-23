package onet

import (
	"context"
	"net/http"
	"testing"
)

func TestTaxonomyCrosswalkPaths(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
		callFn   func(*Client) error
	}{
		{"From2010", "/taxonomy/2010/active/15-1252.00", func(c *Client) error { _, err := c.TaxonomyFrom2010(context.Background(), "15-1252.00"); return err }},
		{"To2010", "/taxonomy/active/2010/15-1252.00", func(c *Client) error { _, err := c.TaxonomyTo2010(context.Background(), "15-1252.00"); return err }},
		{"From2019", "/taxonomy/2019/active/15-1252.00", func(c *Client) error { _, err := c.TaxonomyFrom2019(context.Background(), "15-1252.00"); return err }},
		{"To2019", "/taxonomy/active/2019/15-1252.00", func(c *Client) error { _, err := c.TaxonomyTo2019(context.Background(), "15-1252.00"); return err }},
		{"2010To2019", "/taxonomy/2010/2019/15-1252.00", func(c *Client) error { _, err := c.Taxonomy2010To2019(context.Background(), "15-1252.00"); return err }},
		{"2019To2010", "/taxonomy/2019/2010/15-1252.00", func(c *Client) error { _, err := c.Taxonomy2019To2010(context.Background(), "15-1252.00"); return err }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.wantPath {
					t.Errorf("path = %q, want %q", r.URL.Path, tt.wantPath)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{}`))
			})
			defer close()
			if err := tt.callFn(c); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
