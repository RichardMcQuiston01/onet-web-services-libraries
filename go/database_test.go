package onet

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestDatabaseTables(t *testing.T) {
	want := DatabaseTablesResponse{
		Table: []DatabaseTableRef{
			{ID: "abilities", Title: "Abilities"},
			{ID: "skills", Title: "Skills"},
		},
	}
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/database/" {
			t.Errorf("path = %q, want /database/", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.DatabaseTables(context.Background())
	if err != nil {
		t.Fatalf("DatabaseTables: %v", err)
	}
	if len(got.Table) != 2 {
		t.Errorf("len(Table) = %d, want 2", len(got.Table))
	}
}

func TestDatabaseTableInfo(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/database/info/abilities" {
			t.Errorf("path = %q, want /database/info/abilities", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"id":"abilities","title":"Abilities","column":[]}`))
	})
	defer close()

	got, err := c.DatabaseTableInfo(context.Background(), "abilities")
	if err != nil {
		t.Fatalf("DatabaseTableInfo: %v", err)
	}
	if got.ID != "abilities" {
		t.Errorf("ID = %q, want abilities", got.ID)
	}
}

func TestDatabaseRows_Filters(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/database/rows/abilities" {
			t.Errorf("path = %q, want /database/rows/abilities", r.URL.Path)
		}
		filters := r.URL.Query()["filter"]
		if len(filters) != 2 {
			t.Errorf("len(filter) = %d, want 2", len(filters))
		}
		wantFilters := map[string]bool{
			"onetsoc_code.eq.15-1252.00": true,
			"scale_id.eq.LV":             true,
		}
		for _, f := range filters {
			if !wantFilters[f] {
				t.Errorf("unexpected filter %q", f)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.DatabaseRows(context.Background(), "abilities",
		[]FilterParam{
			{Column: "onetsoc_code", Op: FilterEq, Value: "15-1252.00"},
			{Column: "scale_id", Op: FilterEq, Value: "LV"},
		},
		nil,
		PageParams{},
	)
	if err != nil {
		t.Fatalf("DatabaseRows: %v", err)
	}
}

func TestDatabaseRows_Sort(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		sorts := r.URL.Query()["sort"]
		if len(sorts) != 2 {
			t.Errorf("len(sort) = %d, want 2", len(sorts))
		}
		if sorts[0] != "onetsoc_code" {
			t.Errorf("sort[0] = %q, want onetsoc_code", sorts[0])
		}
		if sorts[1] != "data_value.descending" {
			t.Errorf("sort[1] = %q, want data_value.descending", sorts[1])
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
	})
	defer close()

	_, err := c.DatabaseRows(context.Background(), "abilities",
		nil,
		[]SortParam{
			{Column: "onetsoc_code"},
			{Column: "data_value", Descending: true},
		},
		PageParams{},
	)
	if err != nil {
		t.Fatalf("DatabaseRows: %v", err)
	}
}

func TestFilterParam_Encode(t *testing.T) {
	tests := []struct {
		f    FilterParam
		want string
	}{
		{FilterParam{"col", FilterEq, "val"}, "col.eq.val"},
		{FilterParam{"name", FilterContains, "dev"}, "name.contains.dev"},
		{FilterParam{"code", FilterStartsWith, "15-"}, "code.startswith.15-"},
	}
	for _, tt := range tests {
		if got := tt.f.encode(); got != tt.want {
			t.Errorf("encode() = %q, want %q", got, tt.want)
		}
	}
}

func TestSortParam_Encode(t *testing.T) {
	asc := SortParam{Column: "title"}
	if got := asc.encode(); got != "title" {
		t.Errorf("encode() = %q, want title", got)
	}
	desc := SortParam{Column: "data_value", Descending: true}
	if got := desc.encode(); got != "data_value.descending" {
		t.Errorf("encode() = %q, want data_value.descending", got)
	}
}
