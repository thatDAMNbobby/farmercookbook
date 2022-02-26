package search

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch"
)

type Deps struct {
	Elasticsearch elasticsearch.Elasticsearch
}

type Search interface {
	Search(ctx context.Context, query *elasticsearch.SearchRequest) (*elastic.SearchResult, error)
}

type impl struct {
	deps *Deps
}

func New(deps *Deps) Search {
	return &impl{deps}
}

func (impl *impl) Search(ctx context.Context, query *elasticsearch.SearchRequest) (*elastic.SearchResult, error) {
	results, err := impl.deps.Elasticsearch.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return results, nil
}
