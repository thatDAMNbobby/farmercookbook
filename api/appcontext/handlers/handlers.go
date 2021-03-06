package handlers

import (
	"github.com/thatDAMNbobby/farmercookbook/appcontext/clients"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/dals"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/services"
	"github.com/thatDAMNbobby/farmercookbook/config"
	"github.com/thatDAMNbobby/farmercookbook/handlers/health"
	"github.com/thatDAMNbobby/farmercookbook/handlers/index"
	"github.com/thatDAMNbobby/farmercookbook/handlers/search"
)

type Handlers struct {
	Health health.Health
	Search search.Search
	Index  index.Index
}

func New(config *config.Config, clients *clients.Clients, dals *dals.DALs, services *services.Services) *Handlers {
	return &Handlers{
		Health: health.New(&health.Deps{
			Health: services.Health,
			Render: clients.Render,
		}),
		Search: search.New(&search.Deps{
			Search: services.Search,
			Render: clients.Render,
		}),
		Index: index.New(&index.Deps{
			IndexService: services.Index,
			Render:       clients.Render,
		}),
	}
}
