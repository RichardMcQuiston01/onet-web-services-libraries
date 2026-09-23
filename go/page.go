package onet

import (
	"net/url"
	"strconv"
)

// PageParams controls which page of results to return.
// Both fields are optional; the API defaults to Start=1, End=Start+19.
// The maximum page window is 2000 results (End ≤ Start+1999).
type PageParams struct {
	Start int // 1-based index of the first result
	End   int // inclusive index of the last result
}

func (p PageParams) apply(q url.Values) {
	if p.Start > 0 {
		q.Set("start", strconv.Itoa(p.Start))
	}
	if p.End > 0 {
		q.Set("end", strconv.Itoa(p.End))
	}
}

// PageMeta holds the pagination fields present in all list responses.
type PageMeta struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Total int    `json:"total"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
}
