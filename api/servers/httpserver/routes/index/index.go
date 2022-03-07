package index

import (
	"github.com/gorilla/mux"
	"github.com/thatDAMNbobby/farmercookbook/appcontext/handlers"
	"net/http"
)

func Register(parentRouter *mux.Router, pathPrefix string, handers *handlers.Handlers) {
	router := parentRouter.PathPrefix(pathPrefix).Subrouter()
	router.HandleFunc("/upsert", http.HandlerFunc(handers.Index.Upsert)).Methods("POST", "PUT")
	router.HandleFunc("/delete", http.HandlerFunc(handers.Index.Upsert)).Methods("DELETE")
}
