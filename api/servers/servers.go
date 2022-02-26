package servers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/handlers"
	"github.com/thatDAMNbobby/farmercookbook/config"
	"github.com/thatDAMNbobby/farmercookbook/servers/httpserver"
	"github.com/thatDAMNbobby/farmercookbook/servers/runnable"
)

type Deps struct {
	Handlers *handlers.Handlers
}

type impl struct {
	httpServer runnable.Runnable
}

func New(deps *Deps, config *config.ServerConfig) runnable.Runnable {
	return &impl{
		httpServer: httpserver.New(&httpserver.Deps{Handlers: deps.Handlers}, config),
	}
}

func (impl *impl) Start(ctx context.Context) {
	log.Println("Starting all servers ...")
	go impl.httpServer.Start(ctx)
}

func (impl *impl) Stop(ctx context.Context) {
	log.Println("Stopping all servers ...")
	impl.httpServer.Stop(ctx)
	log.Println("Servers stopped")
}
