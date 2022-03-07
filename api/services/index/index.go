package index

import (
	"context"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch"
	"github.com/thatDAMNbobby/farmercookbook/clients/elasticsearch/requests"
	"github.com/thatDAMNbobby/farmercookbook/utils"
)

type Deps struct {
	Elasticsearch elasticsearch.Elasticsearch
}

type Index interface {
	Delete(ctx context.Context, request *requests.DeleteRequest) error
	Upsert(ctx context.Context, request *requests.UpsertRequest) error
}

type impl struct {
	deps *Deps
}

func New(deps *Deps) Index {
	return &impl{deps}
}

func (impl *impl) Delete(ctx context.Context, request *requests.DeleteRequest) error {
	utils.PrintDebugJSON("delete request", request)
	err := impl.deps.Elasticsearch.Delete(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (impl *impl) Upsert(ctx context.Context, request *requests.UpsertRequest) error {
	utils.PrintDebugJSON("upsert request", request)
	err := impl.deps.Elasticsearch.Upsert(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
