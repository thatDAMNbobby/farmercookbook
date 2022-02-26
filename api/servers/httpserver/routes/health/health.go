package health

import (
	"github.com/gorilla/mux"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/handlers"
	"net/http"
)

// Register Health endpoints
func Register(parentRouter *mux.Router, pathPrefix string, handlers *handlers.Handlers) {
	router := parentRouter.PathPrefix(pathPrefix).Subrouter()

	router.Handle("/ping", http.HandlerFunc(handlers.Health.Ping)).Methods(http.MethodGet)
}
