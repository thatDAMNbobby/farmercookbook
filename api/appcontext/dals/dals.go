package dals

import (
	"github.com/thatDAMNbobby/farmercookbook/appcontext/clients"
	"github.com/thatDAMNbobby/farmercookbook/config"
)

type DALs struct {
	Config  config.Config
	Clients clients.Clients
}

func New(config *config.Config, clients *clients.Clients) *DALs {
	return &DALs{
		Config:  *config,
		Clients: *clients,
	}
}
