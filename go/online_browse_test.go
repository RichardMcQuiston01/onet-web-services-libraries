package onet

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestOnlineBrightOutlookCategories(t *testing.T) {
	want := BrightOutlookCategoriesResponse{
		Category: []BrightOutlookCategory{{Value: "grow_fast", Title: "Grow Fast"}},
	}
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/online/bright_outlook/" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.OnlineBrightOutlookCategories(context.Background())
	if err != nil {
		t.Fatalf("OnlineBrightOutlookCategories: %v", err)
	}
	if len(got.Category) != 1 || got.Category[0].Value != "grow_fast" {
		t.Errorf("Category = %+v", got.Category)
	}
}

func TestOnlineCrosswalkPaths(t *testing.T) {
	crosswalks := []struct {
		name   string
		callFn func(*Client) error
	}{
		{"military", func(c *Client) error { _, err := c.OnlineCrosswalkMilitary(context.Background(), "infantry", PageParams{}); return err }},
		{"education", func(c *Client) error { _, err := c.OnlineCrosswalkEducation(context.Background(), "engineering", PageParams{}); return err }},
		{"occupation_handbook", func(c *Client) error { _, err := c.OnlineCrosswalkOccupationHandbook(context.Background(), "nurse", PageParams{}); return err }},
		{"SOC", func(c *Client) error { _, err := c.OnlineCrosswalkSOC(context.Background(), "15-1252", PageParams{}); return err }},
		{"DOT", func(c *Client) error { _, err := c.OnlineCrosswalkDOT(context.Background(), "programmer", PageParams{}); return err }},
		{"RAPIDS", func(c *Client) error { _, err := c.OnlineCrosswalkRAPIDS(context.Background(), "plumber", PageParams{}); return err }},
		{"ESCO", func(c *Client) error { _, err := c.OnlineCrosswalkESCO(context.Background(), "developer", PageParams{}); return err }},
	}

	for _, tt := range crosswalks {
		t.Run(tt.name, func(t *testing.T) {
			wantSuffix := "/online/crosswalks/" + tt.name
			c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				if !strings.HasSuffix(r.URL.Path, wantSuffix) {
					t.Errorf("path = %q, want suffix %q", r.URL.Path, wantSuffix)
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

func TestOnlineJobDutiesResults_JoinsTaskIDs(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("tasks"); got != "t1,t2,t3" {
			t.Errorf("tasks = %q, want %q", got, "t1,t2,t3")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.OnlineJobDutiesResults(context.Background(), "15-1252.00", []string{"t1", "t2", "t3"})
	if err != nil {
		t.Fatalf("OnlineJobDutiesResults: %v", err)
	}
}

func TestOnlineONETDataPaths(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
		callFn   func(*Client) error
	}{
		{"abilities list", "/online/onet_data/abilities/", func(c *Client) error { _, err := c.OnlineONETAbilities(context.Background()); return err }},
		{"abilities detail", "/online/onet_data/abilities/1.A.1.a", func(c *Client) error { _, err := c.OnlineONETAbility(context.Background(), "1.A.1.a"); return err }},
		{"interests list", "/online/onet_data/interests/", func(c *Client) error { _, err := c.OnlineONETInterests(context.Background()); return err }},
		{"work_styles list", "/online/onet_data/work_styles/", func(c *Client) error { _, err := c.OnlineONETWorkStyles(context.Background()); return err }},
		{"skills_basic list", "/online/onet_data/skills_basic/", func(c *Client) error { _, err := c.OnlineONETBasicSkills(context.Background()); return err }},
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
