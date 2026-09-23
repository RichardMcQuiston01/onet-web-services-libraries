package onet

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestAbout(t *testing.T) {
	want := AboutResponse{
		APIVersion:  "2.0",
		OnetVersion: "28.3",
		ReleaseDate: "2024-07-01",
	}

	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about/" {
			t.Errorf("path = %q, want /about/", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want)
	})
	defer close()

	got, err := c.About(context.Background())
	if err != nil {
		t.Fatalf("About: %v", err)
	}
	if *got != want {
		t.Errorf("About() = %+v, want %+v", *got, want)
	}
}

func TestAbout_PropagatesAPIError(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(429)
		w.Write([]byte(`{"error":"rate limit exceeded"}`))
	})
	defer close()

	_, err := c.About(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("error type = %T, want *APIError", err)
	}
	if apiErr.StatusCode != 429 {
		t.Errorf("StatusCode = %d, want 429", apiErr.StatusCode)
	}
}
