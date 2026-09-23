package onet

import "context"

// TaxonomyOccupation is an occupation returned in a taxonomy crosswalk response.
type TaxonomyOccupation struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// TaxonomyResponse is the uniform response shape for all taxonomy crosswalk endpoints.
type TaxonomyResponse struct {
	Code        string               `json:"code"`
	Title       string               `json:"title"`
	Description string               `json:"description,omitempty"`
	Occupation  []TaxonomyOccupation `json:"occupation"`
}

// taxonomyCrosswalk is the shared helper for all /taxonomy/{from}/{to}/{code} endpoints.
func (c *Client) taxonomyCrosswalk(ctx context.Context, from, to, code string) (*TaxonomyResponse, error) {
	var out TaxonomyResponse
	return &out, c.do(ctx, "taxonomy/"+from+"/"+to+"/"+code, nil, &out)
}

// TaxonomyFrom2010 returns the active O*NET-SOC occupation(s) for a 2010 code.
func (c *Client) TaxonomyFrom2010(ctx context.Context, code string) (*TaxonomyResponse, error) {
	return c.taxonomyCrosswalk(ctx, "2010", "active", code)
}

// TaxonomyTo2010 returns the 2010 O*NET-SOC code(s) for an active code.
func (c *Client) TaxonomyTo2010(ctx context.Context, code string) (*TaxonomyResponse, error) {
	return c.taxonomyCrosswalk(ctx, "active", "2010", code)
}

// TaxonomyFrom2019 returns the active O*NET-SOC occupation(s) for a 2019 code.
func (c *Client) TaxonomyFrom2019(ctx context.Context, code string) (*TaxonomyResponse, error) {
	return c.taxonomyCrosswalk(ctx, "2019", "active", code)
}

// TaxonomyTo2019 returns the 2019 O*NET-SOC code(s) for an active code.
func (c *Client) TaxonomyTo2019(ctx context.Context, code string) (*TaxonomyResponse, error) {
	return c.taxonomyCrosswalk(ctx, "active", "2019", code)
}

// Taxonomy2010To2019 returns the 2019 O*NET-SOC code(s) for a 2010 code.
func (c *Client) Taxonomy2010To2019(ctx context.Context, code string) (*TaxonomyResponse, error) {
	return c.taxonomyCrosswalk(ctx, "2010", "2019", code)
}

// Taxonomy2019To2010 returns the 2010 O*NET-SOC code(s) for a 2019 code.
func (c *Client) Taxonomy2019To2010(ctx context.Context, code string) (*TaxonomyResponse, error) {
	return c.taxonomyCrosswalk(ctx, "2019", "2010", code)
}
