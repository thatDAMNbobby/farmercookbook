package services

import (
	"github.com/thatDAMNbobby/farmercookbook/appcontext/clients"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/dals"
	"github.com/thatDAMNbobby/farmercookbook/clients/render"
	"github.com/thatDAMNbobby/farmercookbook/config"
	"github.com/thatDAMNbobby/farmercookbook/services/health"
	"github.com/thatDAMNbobby/farmercookbook/services/search"
)

type Services struct {
	Health health.Health
	Search search.Search
	Render render.Render
}

func New(config *config.Config, clients *clients.Clients, dals *dals.DALs) *Services {
	return &Services{
		Health: health.New(
			&health.Deps{
				ElasticSearch: clients.Elasticsearch,
			}),
		Search: search.New(
			&search.Deps{
				ElasticSearch: clients.Elasticsearch,
				DAL:           dals.Recipe,
			}),
	}
}
