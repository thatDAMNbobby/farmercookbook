package health

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type Pingable interface {
	Ping(ctx context.Context) error
	GetName() string
}

type Deps struct {
	ElasticSearch Pingable
}

type Health interface {
	Ping(ctx context.Context) *HealthResponse
}

type impl struct {
	deps *Deps
}

func New(deps *Deps) Health {
	return &impl{deps: deps}
}

type HealthResponse struct {
	SearchESAvailable bool `json:"searchES"`
}

func RunPing(ctx context.Context, pingable Pingable) bool {
	err := pingable.Ping(ctx)

	if err != nil {
		log.WithField("name", pingable.GetName()).Error("Ping failed: ", err)
		return false
	}

	return true
}

func (impl *impl) Ping(ctx context.Context) *HealthResponse {
	return &HealthResponse{
		SearchESAvailable: RunPing(ctx, impl.deps.ElasticSearch),
	}
}
