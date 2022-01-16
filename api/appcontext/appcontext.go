package appcontext

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type AppContext struct{}

func New() *AppContext {

	return &AppContext{}
}

func (impl *AppContext) Start(ctx context.Context) {}

func (impl *AppContext) Stop(ctx context.Context) {
	log.Println("Stopping appcontext")
	log.Println("appcontext stopped")
}
