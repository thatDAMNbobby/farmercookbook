package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch/requests"
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
	Search(ctx context.Context, request *requests.SearchRequest) (*elastic.SearchResult, error)
	Delete(ctx context.Context, request *requests.DeleteRequest) error
	Upsert(ctx context.Context, request *requests.UpsertRequest) error
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

func (impl *impl) Search(ctx context.Context, request *requests.SearchRequest) (*elastic.SearchResult, error) {

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

	fmt.Printf("%s: %d hits\n", request.Indices, resp.TotalHits())
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

func (impl *impl) Delete(ctx context.Context, request *requests.DeleteRequest) error {
	_, err := impl.index.Delete().Index(request.Index).Id(request.Id).Do(ctx)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (impl *impl) Upsert(ctx context.Context, request *requests.UpsertRequest) error {
	var err error
	if request.Id == nil {
		_, err = impl.index.Index().Index(request.Index).BodyJson(request.Body).Do(ctx)
	}
	if request.Id != nil {
		_, err = impl.index.Index().Index(request.Index).Id(*request.Id).BodyJson(request.Body).Do(ctx)
	}
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
