package onet

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://services.onetcenter.org/"

// Client accesses the O*NET Web Services API.
// All methods are safe for concurrent use.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL overrides the default API base URL.
// Useful in tests or when targeting a non-production environment.
func WithBaseURL(u string) Option {
	return func(c *Client) {
		c.baseURL = u
	}
}

// WithHTTPClient replaces the default HTTP client.
// Use this to set timeouts, transport options, or inject a test client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.http = hc
	}
}

// NewClient returns a Client authenticated with apiKey.
// Register for an API key at https://services.onetcenter.org/reference/start/authorization.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: defaultBaseURL,
		http:    &http.Client{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// do executes a GET request to path, appends query params, and JSON-decodes
// a successful response into out. Non-2xx responses are returned as *APIError.
func (c *Client) do(ctx context.Context, path string, query url.Values, out any) error {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return fmt.Errorf("onet: invalid base URL: %w", err)
	}
	target := base.JoinPath(path)
	if len(query) > 0 {
		target.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target.String(), nil)
	if err != nil {
		return fmt.Errorf("onet: build request: %w", err)
	}
	req.Header.Set("X-API-Key", c.apiKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("onet: request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var body struct {
			Err string `json:"error"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&body)
		return &APIError{StatusCode: resp.StatusCode, Message: body.Err}
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("onet: decode response: %w", err)
	}
	return nil
}
