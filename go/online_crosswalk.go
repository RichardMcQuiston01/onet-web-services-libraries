package onet

import (
	"context"
	"net/url"
)

// crosswalk is the shared helper for all /online/crosswalks/{name} endpoints.
// All crosswalk endpoints accept a keyword and optional pagination, and return
// a paged list of occupations.
func (c *Client) crosswalk(ctx context.Context, name, keyword string, page PageParams) (*OccupationList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/crosswalks/"+name, q, &out)
}

// OnlineCrosswalkMilitary searches occupations by military keyword.
func (c *Client) OnlineCrosswalkMilitary(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "military", keyword, page)
}

// OnlineCrosswalkEducation searches occupations by education program keyword.
func (c *Client) OnlineCrosswalkEducation(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "education", keyword, page)
}

// OnlineCrosswalkOccupationHandbook searches occupations by Occupational Outlook Handbook keyword.
func (c *Client) OnlineCrosswalkOccupationHandbook(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "occupation_handbook", keyword, page)
}

// OnlineCrosswalkSOC searches occupations by Standard Occupational Classification keyword.
func (c *Client) OnlineCrosswalkSOC(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "SOC", keyword, page)
}

// OnlineCrosswalkDOT searches occupations by Dictionary of Occupational Titles keyword.
func (c *Client) OnlineCrosswalkDOT(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "DOT", keyword, page)
}

// OnlineCrosswalkRAPIDS searches occupations by RAPIDS apprenticeship code or keyword.
func (c *Client) OnlineCrosswalkRAPIDS(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "RAPIDS", keyword, page)
}

// OnlineCrosswalkESCO searches occupations by ESCO (European Skills/Competences) keyword.
func (c *Client) OnlineCrosswalkESCO(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	return c.crosswalk(ctx, "ESCO", keyword, page)
}
