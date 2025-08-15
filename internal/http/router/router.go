package router

import (
	"github.com/bhushankp/app-api.git/internal/http/handlers"
	"github.com/bhushankp/app-api.git/internal/infra/db"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	dbMock := &db.MockDB{Healthy: true}

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/healtz", handlers.Health).Methods("GET")
	v1.HandleFunc("/readyz", handlers.Ready(dbMock)).Methods("GET")
	return r
}
