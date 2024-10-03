// Copyright (c) 2024 space-code
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package utils

import "math"

// ListQuery represents the parameters for querying a paginated list.
type ListQuery struct {
	Size    int            `query:"size" json:"size,omitempty"`
	Page    int            `query:"page" json:"page,omitempty"`
	OrderBy string         `query:"orderBy" json:"orderBy,omitempty"`
	Filters []*FilterModel `query:"filters" json:"filters,omitempty"`
}

// FilterModel defines the structure for filtering query results.
type FilterModel struct {
	Field      string `query:"field" json:"field"`
	Value      string `query:"value" json:"value"`
	Comparison string `query:"comparison" json:"comparison"`
}

// ListResult encapsulates the result of a paginated query.
type ListResult[T interface{}] struct {
	Size       int   `json:"size,omitempty" bson:"size"`
	Page       int   `json:"page,omitempty" bson:"page"`
	TotalItems int64 `json:"totalItems,omitempty" bson:"totalItems"`
	TotalPage  int   `json:"totalPage,omitempty" bson:"totalPage"`
	Items      []T   `json:"items,omitempty" bson:"items"`
}

// NewListResult constructs a new ListResult instance.
func NewListResult[T any](items []T, size int, page int, totalItems int64) *ListResult[T] {
	listResult := &ListResult[T]{Items: items, Size: size, Page: page, TotalItems: totalItems}

	listResult.TotalPage = getTotalPages(totalItems, size)

	return listResult
}

// getTotalPages calculates the total number of pages.
func getTotalPages(totalCount int64, size int) int {
	d := float64(totalCount) / float64(size)
	return int(math.Ceil(d))
}

// GetPage returns the current page number from the ListQuery.
func (q *ListQuery) GetPage() int {
	return q.Page
}

// GetSize returns the number of items per page from the ListQuery.
func (q *ListQuery) GetSize() int {
	return q.Size
}

// GetOffset calculates the offset for the database query based on the current page and page size.
func (q *ListQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

// GetLimit returns the maximum number of items to retrieve, based on the page size.
func (q *ListQuery) GetLimit() int {
	return q.Size
}

// GetOrderBy returns the field by which the results should be ordered.
func (q *ListQuery) GetOrderBy() string {
	return q.OrderBy
}
