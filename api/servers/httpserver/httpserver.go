package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/handlers"
	"github.com/thatDAMNbobby/farmercookbook/config"
	"github.com/thatDAMNbobby/farmercookbook/servers/httpserver/routes/health"
	"github.com/thatDAMNbobby/farmercookbook/servers/httpserver/routes/search"
	"github.com/thatDAMNbobby/farmercookbook/servers/runnable"
	"github.com/thatDAMNbobby/farmercookbook/servers/server"
	"net/http"
)

type Deps struct {
	Handlers *handlers.Handlers
}

type Config struct {
	Port int
}

type impl struct {
	deps      *Deps
	config    *config.ServerConfig
	tcpServer *http.Server
}

func New(deps *Deps, config *config.ServerConfig) runnable.Runnable {
	return &impl{
		deps:   deps,
		config: config,
	}
}

func newRouter(handlers *handlers.Handlers) http.Handler {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	search.Register(apiRouter, "/search", handlers)

	v1Router := apiRouter.PathPrefix("/v1").Subrouter()
	health.Register(v1Router, "/health", handlers)

	return router
}

func (impl *impl) Start(ctx context.Context) {
	router := newRouter(impl.deps.Handlers)

	go func() {
		entry := log.WithField("port", impl.config.Port)
		impl.tcpServer = server.NewTCPServerWithRouter(ctx, impl.config.Port, router)
		entry.Info("Starting TCP server")
		err := impl.tcpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			entry.Error(err)
		}
	}()
}

func (impl *impl) Stop(ctx context.Context) {
	log.Info("Stopping HTTP servers")

	log.Info("Stopping TCP server")
	err := impl.tcpServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Failed to shutdown TCP server: ", err)
	}

	log.Info("Gracefully stopped HTTP server")
}
