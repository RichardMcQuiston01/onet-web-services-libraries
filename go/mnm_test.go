package onet

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestMNMSearch(t *testing.T) {
	want := CareerList{
		PageMeta: PageMeta{Start: 1, End: 1, Total: 1},
		Career:   []Career{{Code: "15-1252.00", Title: "Software Developers"}},
	}
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mnm/search" {
			t.Errorf("path = %q, want /mnm/search", r.URL.Path)
		}
		if r.URL.Query().Get("keyword") != "software" {
			t.Errorf("keyword = %q", r.URL.Query().Get("keyword"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.MNMSearch(context.Background(), "software", PageParams{})
	if err != nil {
		t.Fatalf("MNMSearch: %v", err)
	}
	if len(got.Career) != 1 || got.Career[0].Code != "15-1252.00" {
		t.Errorf("Career = %+v", got.Career)
	}
}

func TestMPPSearch_UsesMPPPrefix(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/mpp/search" {
			t.Errorf("path = %q, want /mpp/search", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.MPPSearch(context.Background(), "software", PageParams{})
	if err != nil {
		t.Fatalf("MPPSearch: %v", err)
	}
}

func TestMNMCareerResourcePaths(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
		callFn   func(*Client) error
	}{
		{"career", "/mnm/careers/15-1252.00/", func(c *Client) error { _, err := c.MNMCareer(context.Background(), "15-1252.00"); return err }},
		{"knowledge", "/mnm/careers/15-1252.00/knowledge", func(c *Client) error { _, err := c.MNMCareerKnowledge(context.Background(), "15-1252.00"); return err }},
		{"skills", "/mnm/careers/15-1252.00/skills", func(c *Client) error { _, err := c.MNMCareerSkills(context.Background(), "15-1252.00"); return err }},
		{"abilities", "/mnm/careers/15-1252.00/abilities", func(c *Client) error { _, err := c.MNMCareerAbilities(context.Background(), "15-1252.00"); return err }},
		{"personality", "/mnm/careers/15-1252.00/personality", func(c *Client) error { _, err := c.MNMCareerPersonality(context.Background(), "15-1252.00"); return err }},
		{"technology", "/mnm/careers/15-1252.00/technology", func(c *Client) error { _, err := c.MNMCareerTechnology(context.Background(), "15-1252.00"); return err }},
		{"education", "/mnm/careers/15-1252.00/education", func(c *Client) error { _, err := c.MNMCareerEducation(context.Background(), "15-1252.00"); return err }},
		{"job_outlook", "/mnm/careers/15-1252.00/job_outlook", func(c *Client) error { _, err := c.MNMCareerJobOutlook(context.Background(), "15-1252.00"); return err }},
		{"state", "/mnm/careers/15-1252.00/state", func(c *Client) error { _, err := c.MNMCareerState(context.Background(), "15-1252.00"); return err }},
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

func TestMNMInterestProfilerResults_JoinsAnswers(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("answers"); got != "3,4,5,3,2" {
			t.Errorf("answers = %q, want %q", got, "3,4,5,3,2")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.MNMInterestProfilerResults(context.Background(), []string{"3", "4", "5", "3", "2"})
	if err != nil {
		t.Fatalf("MNMInterestProfilerResults: %v", err)
	}
}

func TestMNMJobPreparationZone_Path(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/mnm/job_preparation/3") {
			t.Errorf("path = %q, want suffix /mnm/job_preparation/3", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.MNMJobPreparationZone(context.Background(), 3)
	if err != nil {
		t.Fatalf("MNMJobPreparationZone: %v", err)
	}
}

func TestVeteransSearch(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/veterans/military" {
			t.Errorf("path = %q, want /veterans/military", r.URL.Path)
		}
		if r.URL.Query().Get("keyword") != "infantry" {
			t.Errorf("keyword = %q", r.URL.Query().Get("keyword"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.VeteransSearch(context.Background(), "infantry", PageParams{})
	if err != nil {
		t.Fatalf("VeteransSearch: %v", err)
	}
}
