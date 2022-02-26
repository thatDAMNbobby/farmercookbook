package health

import (
	"context"
	"github.com/thatDAMNbobby/farmercookbook/handlers"
	"net/http"
	"time"

	"github.com/thatDAMNbobby/farmercookbook/clients/render"
	"github.com/thatDAMNbobby/farmercookbook/services/health"
)

type Deps struct {
	Health health.Health
	Render render.Render
}

type impl struct {
	Deps *Deps
}

type Health interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

func New(deps *Deps) Health {
	return &impl{Deps: deps}
}

func (impl *impl) Ping(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	result := impl.Deps.Health.Ping(ctx)

	unavailableSystems := make([]string, 0)

	if !result.SearchESAvailable {
		unavailableSystems = append(unavailableSystems, "SearchES")
	}

	if len(unavailableSystems) > 0 {
		response := handlers.Status{
			Status:  http.StatusServiceUnavailable,
			Message: "Service Unavailable",
			Failed:  unavailableSystems,
		}

		//nolint
		impl.Deps.Render.JSON(w, http.StatusServiceUnavailable, response)
		return
	}

	//nolint
	impl.Deps.Render.JSON(w, http.StatusOK, handlers.Status{
		Status:  http.StatusOK,
		Message: "Healthy",
		Failed:  unavailableSystems,
	})
}
