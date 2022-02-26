package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"time"
)

type Config struct {
	Address          string
	Sniff            bool
	IndexRetryMinS   int
	IndexRetryMaxS   int
	SearchRetryMinMS int
	SearchRetryMaxMS int
}

type Elasticsearch interface {
	Search(ctx context.Context, request *SearchRequest) (*elastic.SearchResult, error)
	Ping(ctx context.Context) error
	GetName() string
}

type impl struct {
	search  *elastic.Client
	index   *elastic.Client
	address string
}

func New(config *Config) Elasticsearch {
	searchBackoff := elastic.NewExponentialBackoff(time.Duration(config.SearchRetryMinMS)*time.Millisecond, time.Duration(config.SearchRetryMaxMS)*time.Millisecond)
	searchRetrier := elastic.NewBackoffRetrier(searchBackoff)
	search, err := elastic.NewClient(elastic.SetSniff(config.Sniff), elastic.SetURL(config.Address), elastic.SetRetrier(searchRetrier))
	if err != nil {
		panic(err)
	}

	indexBackoff := elastic.NewExponentialBackoff(time.Duration(config.IndexRetryMinS)*time.Second, time.Duration(config.IndexRetryMaxS)*time.Second)
	indexRetrier := elastic.NewBackoffRetrier(indexBackoff)
	index, err := elastic.NewClient(elastic.SetSniff(config.Sniff), elastic.SetURL(config.Address), elastic.SetRetrier(indexRetrier))
	if err != nil {
		panic(err)
	}

	return &impl{
		index:   index,
		search:  search,
		address: config.Address,
	}
}

func (impl *impl) Search(ctx context.Context, request *SearchRequest) (*elastic.SearchResult, error) {

	search := impl.search.
		Search(request.Indices...).
		Query(request.Query).
		From(request.From).
		Size(request.Size).
		SortBy(request.Sort...)

	resp, err := search.Do(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (impl *impl) Ping(ctx context.Context) error {
	_, code, err := impl.search.Ping(impl.address).Do(ctx)

	if err != nil {
		return errors.WithStack(err)
	}

	if code != 200 {
		return errors.New(fmt.Sprintf("ping failed: code=%+v", code))
	}

	return nil
}

func (impl *impl) GetName() string {
	return "elasticsearch"
}
