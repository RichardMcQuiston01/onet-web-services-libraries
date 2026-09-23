package onet

import (
	"context"
	"net/url"
)

// OccupationList is a paginated list of occupations.
type OccupationList struct {
	PageMeta
	Occupation []Occupation `json:"occupation"`
}

// OccupationDetail is a full occupation record returned by OnlineOccupation.
type OccupationDetail struct {
	Code        string          `json:"code"`
	Title       string          `json:"title"`
	AlsoCalled  []string        `json:"also_called,omitempty"`
	WhatTheyDo  string          `json:"what_they_do,omitempty"`
	OnTheJob    []string        `json:"on_the_job,omitempty"`
	CareerVideo string          `json:"career_video,omitempty"`
	Tags        *OccupationTags `json:"tags,omitempty"`
}

// OnlineSearch searches O*NET occupations by keyword.
// keyword is required; page is optional (zero value uses API defaults).
func (c *Client) OnlineSearch(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out OccupationList
	if err := c.do(ctx, "online/search", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OnlineOccupations returns a paginated list of all O*NET occupations.
func (c *Client) OnlineOccupations(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	if err := c.do(ctx, "online/occupations/", q, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// OnlineOccupation returns a single occupation by its O*NET-SOC code.
func (c *Client) OnlineOccupation(ctx context.Context, code string) (*OccupationDetail, error) {
	var out OccupationDetail
	if err := c.do(ctx, "online/occupations/"+code+"/", nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// occupationResource is the shared path helper for summary and details endpoints.
func (c *Client) occupationResource(ctx context.Context, code, kind, topic string, out any) error {
	return c.do(ctx, "online/occupations/"+code+"/"+kind+"/"+topic, nil, out)
}
