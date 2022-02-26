package appcontext

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/clients"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/dals"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/handlers"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/services"
	"github.com/thatDAMNbobby/farmercookbook/config"
)

type AppContext struct {
	Clients  *clients.Clients
	DALs     *dals.DALs
	Services *services.Services
	Handlers *handlers.Handlers
}

func New(config *config.Config) *AppContext {
	clientsImpl := clients.New(config)
	dalsImpl := dals.New(config, clientsImpl)
	servicesImpl := services.New(config, clientsImpl, dalsImpl)
	handlersImpl := handlers.New(config, clientsImpl, dalsImpl, servicesImpl)

	return &AppContext{
		Clients:  clientsImpl,
		DALs:     dalsImpl,
		Services: servicesImpl,
		Handlers: handlersImpl,
	}
}

func (impl *AppContext) Start(ctx context.Context) {}

func (impl *AppContext) Stop(ctx context.Context) {
	log.Println("Stopping appcontext")
	log.Println("appcontext stopped")
}
