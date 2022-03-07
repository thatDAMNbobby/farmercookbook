package requests

import "github.com/olivere/elastic/v7"

type SearchRequest struct {
	Query      elastic.Query
	Type       string
	Source     []string
	From       int
	Size       int
	Indices    []string
	Aggs       map[string]elastic.Query
	Sort       []elastic.Sorter
	PostFilter elastic.Query
	Explain    bool
}
