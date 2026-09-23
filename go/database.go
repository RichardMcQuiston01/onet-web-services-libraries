package onet

import (
	"context"
	"net/url"
)

// FilterOp is a comparison operator for DatabaseRows filter params.
type FilterOp string

const (
	FilterEq         FilterOp = "eq"
	FilterNe         FilterOp = "ne"
	FilterLt         FilterOp = "lt"
	FilterLe         FilterOp = "le"
	FilterGt         FilterOp = "gt"
	FilterGe         FilterOp = "ge"
	FilterContains   FilterOp = "contains"
	FilterStartsWith FilterOp = "startswith"
)

// FilterParam is a filter condition for DatabaseRows.
// It encodes to the API format: column.op.value (e.g., "onetsoc_code.eq.15-1252.00").
type FilterParam struct {
	Column string
	Op     FilterOp
	Value  string
}

func (f FilterParam) encode() string {
	return f.Column + "." + string(f.Op) + "." + f.Value
}

// SortParam is a sort directive for DatabaseRows.
// It encodes to the API format: column or column.descending.
type SortParam struct {
	Column     string
	Descending bool
}

func (s SortParam) encode() string {
	if s.Descending {
		return s.Column + ".descending"
	}
	return s.Column
}

// DatabaseColumn describes a column in an O*NET database table.
type DatabaseColumn struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// DatabaseTableRef is a brief reference to a database table.
type DatabaseTableRef struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// DatabaseTablesResponse is returned by DatabaseTables.
type DatabaseTablesResponse struct {
	Table []DatabaseTableRef `json:"table"`
}

// DatabaseTableInfoResponse is returned by DatabaseTableInfo.
type DatabaseTableInfoResponse struct {
	ID     string           `json:"id"`
	Title  string           `json:"title"`
	Column []DatabaseColumn `json:"column"`
}

// DatabaseRow is a single result row from DatabaseRows. Column values are keyed
// by column ID and may be strings, numbers, or null depending on the table.
type DatabaseRow map[string]any

// DatabaseRowsResponse is returned by DatabaseRows.
type DatabaseRowsResponse struct {
	PageMeta
	Row []DatabaseRow `json:"row"`
}

// DatabaseTables returns the list of available O*NET database tables.
func (c *Client) DatabaseTables(ctx context.Context) (*DatabaseTablesResponse, error) {
	var out DatabaseTablesResponse
	return &out, c.do(ctx, "database/", nil, &out)
}

// DatabaseTableInfo returns column definitions for the given table.
func (c *Client) DatabaseTableInfo(ctx context.Context, tableID string) (*DatabaseTableInfoResponse, error) {
	var out DatabaseTableInfoResponse
	return &out, c.do(ctx, "database/info/"+tableID, nil, &out)
}

// DatabaseRows queries rows from the given table with optional filters, sort, and pagination.
// Filters and sorts are applied in the order given; multiple filters are ANDed together.
func (c *Client) DatabaseRows(ctx context.Context, tableID string, filters []FilterParam, sorts []SortParam, page PageParams) (*DatabaseRowsResponse, error) {
	q := url.Values{}
	for _, f := range filters {
		q.Add("filter", f.encode())
	}
	for _, s := range sorts {
		q.Add("sort", s.encode())
	}
	page.apply(q)
	var out DatabaseRowsResponse
	return &out, c.do(ctx, "database/rows/"+tableID, q, &out)
}
