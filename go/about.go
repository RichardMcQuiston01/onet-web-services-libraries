package onet

import "context"

// AboutResponse contains version information for the O*NET Web Services API.
type AboutResponse struct {
	APIVersion  string `json:"api_version"`
	OnetVersion string `json:"onet_version"`
	ReleaseDate string `json:"release_date"`
}

// About returns version and release information for the API.
func (c *Client) About(ctx context.Context) (*AboutResponse, error) {
	var out AboutResponse
	if err := c.do(ctx, "about/", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
