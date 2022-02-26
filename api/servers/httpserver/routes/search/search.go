package search

import (
	"github.com/gorilla/mux"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/handlers"
	"net/http"
)

func Register(parentRouter *mux.Router, pathPrefix string, handlers *handlers.Handlers) {
	router := parentRouter.PathPrefix(pathPrefix).Subrouter()

	router.HandleFunc("/", http.HandlerFunc(handlers.Search.Query)).Methods("GET")
}
