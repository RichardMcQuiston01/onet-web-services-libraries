package onet

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func TestOnlineSearch(t *testing.T) {
	want := OccupationList{
		PageMeta:   PageMeta{Start: 1, End: 2, Total: 2},
		Occupation: []Occupation{{Code: "15-1252.00", Title: "Software Developers"}},
	}

	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/online/search" {
			t.Errorf("path = %q, want /online/search", r.URL.Path)
		}
		if got := r.URL.Query().Get("keyword"); got != "software" {
			t.Errorf("keyword = %q, want %q", got, "software")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.OnlineSearch(context.Background(), "software", PageParams{})
	if err != nil {
		t.Fatalf("OnlineSearch: %v", err)
	}
	if len(got.Occupation) != 1 || got.Occupation[0].Code != "15-1252.00" {
		t.Errorf("Occupation = %+v, want code 15-1252.00", got.Occupation)
	}
}

func TestOnlineSearch_PageParams(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("start") != "21" || q.Get("end") != "40" {
			t.Errorf("pagination params: start=%s end=%s", q.Get("start"), q.Get("end"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OccupationList{})
	})
	defer close()

	_, _ = c.OnlineSearch(context.Background(), "dev", PageParams{Start: 21, End: 40})
}

func TestOnlineOccupations(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/online/occupations/" {
			t.Errorf("path = %q, want /online/occupations/", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OccupationList{PageMeta: PageMeta{Total: 1000}})
	})
	defer close()

	got, err := c.OnlineOccupations(context.Background(), PageParams{})
	if err != nil {
		t.Fatalf("OnlineOccupations: %v", err)
	}
	if got.Total != 1000 {
		t.Errorf("Total = %d, want 1000", got.Total)
	}
}

func TestOnlineOccupation(t *testing.T) {
	want := OccupationDetail{
		Code:       "15-1252.00",
		Title:      "Software Developers",
		WhatTheyDo: "Develop software systems.",
	}

	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/online/occupations/15-1252.00/") {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.OnlineOccupation(context.Background(), "15-1252.00")
	if err != nil {
		t.Fatalf("OnlineOccupation: %v", err)
	}
	if got.Code != want.Code || got.Title != want.Title {
		t.Errorf("OnlineOccupation() = %+v, want %+v", got, want)
	}
}

func TestOnlineSummarySkills(t *testing.T) {
	want := OccupationElements{
		Code:    "15-1252.00",
		Title:   "Software Developers",
		Element: []Element{{ID: "2.A.1.a", Name: "Reading Comprehension"}},
	}

	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/online/occupations/15-1252.00/summary/skills"
		if r.URL.Path != wantPath {
			t.Errorf("path = %q, want %q", r.URL.Path, wantPath)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.OnlineSummarySkills(context.Background(), "15-1252.00")
	if err != nil {
		t.Fatalf("OnlineSummarySkills: %v", err)
	}
	if len(got.Element) != 1 || got.Element[0].ID != "2.A.1.a" {
		t.Errorf("Element = %+v", got.Element)
	}
}

func TestOnlineDetailsSkills(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		wantPath := "/online/occupations/15-1252.00/details/skills"
		if r.URL.Path != wantPath {
			t.Errorf("path = %q, want %q", r.URL.Path, wantPath)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(OccupationElements{})
	})
	defer close()

	_, err := c.OnlineDetailsSkills(context.Background(), "15-1252.00")
	if err != nil {
		t.Fatalf("OnlineDetailsSkills: %v", err)
	}
}

func TestOnlineSummaryTopicPaths(t *testing.T) {
	// Verify that each summary method hits the correct path segment.
	tests := []struct {
		name     string
		topic    string
		callFn   func(*Client) error
	}{
		{"tasks", "tasks", func(c *Client) error { _, err := c.OnlineSummaryTasks(context.Background(), "11-1011.00"); return err }},
		{"technology_skills", "technology_skills", func(c *Client) error { _, err := c.OnlineSummaryTechnologySkills(context.Background(), "11-1011.00"); return err }},
		{"work_activities", "work_activities", func(c *Client) error { _, err := c.OnlineSummaryWorkActivities(context.Background(), "11-1011.00"); return err }},
		{"detailed_work_activities", "detailed_work_activities", func(c *Client) error { _, err := c.OnlineSummaryDetailedWorkActivities(context.Background(), "11-1011.00"); return err }},
		{"work_context", "work_context", func(c *Client) error { _, err := c.OnlineSummaryWorkContext(context.Background(), "11-1011.00"); return err }},
		{"job_zone", "job_zone", func(c *Client) error { _, err := c.OnlineSummaryJobZone(context.Background(), "11-1011.00"); return err }},
		{"apprenticeship", "apprenticeship", func(c *Client) error { _, err := c.OnlineSummaryApprenticeship(context.Background(), "11-1011.00"); return err }},
		{"knowledge", "knowledge", func(c *Client) error { _, err := c.OnlineSummaryKnowledge(context.Background(), "11-1011.00"); return err }},
		{"education", "education", func(c *Client) error { _, err := c.OnlineSummaryEducation(context.Background(), "11-1011.00"); return err }},
		{"abilities", "abilities", func(c *Client) error { _, err := c.OnlineSummaryAbilities(context.Background(), "11-1011.00"); return err }},
		{"interests", "interests", func(c *Client) error { _, err := c.OnlineSummaryInterests(context.Background(), "11-1011.00"); return err }},
		{"work_styles", "work_styles", func(c *Client) error { _, err := c.OnlineSummaryWorkStyles(context.Background(), "11-1011.00"); return err }},
		{"related_occupations", "related_occupations", func(c *Client) error { _, err := c.OnlineSummaryRelatedOccupations(context.Background(), "11-1011.00"); return err }},
		{"professional_associations", "professional_associations", func(c *Client) error { _, err := c.OnlineSummaryProfessionalAssociations(context.Background(), "11-1011.00"); return err }},
		{"military_career_summaries", "military_career_summaries", func(c *Client) error { _, err := c.OnlineSummaryMilitaryCareerSummaries(context.Background(), "11-1011.00"); return err }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantSuffix := "/online/occupations/11-1011.00/summary/" + tt.topic
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
