package onet

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(handler)
	c := NewClient("test-key", WithBaseURL(srv.URL+"/"))
	return c, srv.Close
}

func TestDo_SetsAuthHeader(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("X-API-Key"); got != "test-key" {
			t.Errorf("X-API-Key = %q, want %q", got, "test-key")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{})
	})
	defer close()

	var out any
	_ = c.do(context.Background(), "test", nil, &out)
}

func TestDo_UsesGET(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{})
	})
	defer close()

	var out any
	_ = c.do(context.Background(), "test", nil, &out)
}

func TestDo_DecodesJSON(t *testing.T) {
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"foo": "bar"})
	})
	defer close()

	var out struct {
		Foo string `json:"foo"`
	}
	if err := c.do(context.Background(), "test", nil, &out); err != nil {
		t.Fatalf("do: %v", err)
	}
	if out.Foo != "bar" {
		t.Errorf("Foo = %q, want %q", out.Foo, "bar")
	}
}

func TestDo_QueryParams(t *testing.T) {
	var gotQuery url.Values
	c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		gotQuery = r.URL.Query()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{})
	})
	defer close()

	q := url.Values{"keyword": {"go developer"}, "start": {"5"}}
	var out any
	_ = c.do(context.Background(), "search", q, &out)

	if got := gotQuery.Get("keyword"); got != "go developer" {
		t.Errorf("keyword = %q, want %q", got, "go developer")
	}
	if got := gotQuery.Get("start"); got != "5" {
		t.Errorf("start = %q, want %q", got, "5")
	}
}

func TestDo_APIErrors(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		wantMsg    string
	}{
		{"422 bad param", 422, `{"error":"must have required property 'keyword'"}`, "must have required property 'keyword'"},
		{"429 rate limit", 429, `{"error":"rate limit exceeded"}`, "rate limit exceeded"},
		{"404 not found", 404, `{"error":"not found"}`, "not found"},
		{"500 server error", 500, `{"error":"internal server error"}`, "internal server error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, close := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.body))
			})
			defer close()

			var out any
			err := c.do(context.Background(), "test", nil, &out)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			apiErr, ok := err.(*APIError)
			if !ok {
				t.Fatalf("error type = %T, want *APIError", err)
			}
			if apiErr.StatusCode != tt.statusCode {
				t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, tt.statusCode)
			}
			if apiErr.Message != tt.wantMsg {
				t.Errorf("Message = %q, want %q", apiErr.Message, tt.wantMsg)
			}
		})
	}
}

func TestAPIError_Error(t *testing.T) {
	e := &APIError{StatusCode: 422, Message: "invalid keyword"}
	want := "onet: HTTP 422: invalid keyword"
	if got := e.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestPageParams_Apply(t *testing.T) {
	tests := []struct {
		name      string
		params    PageParams
		wantStart string
		wantEnd   string
	}{
		{"both set", PageParams{Start: 1, End: 20}, "1", "20"},
		{"start only", PageParams{Start: 5}, "5", ""},
		{"zero values omitted", PageParams{}, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := url.Values{}
			tt.params.apply(q)
			if got := q.Get("start"); got != tt.wantStart {
				t.Errorf("start = %q, want %q", got, tt.wantStart)
			}
			if got := q.Get("end"); got != tt.wantEnd {
				t.Errorf("end = %q, want %q", got, tt.wantEnd)
			}
		})
	}
}
