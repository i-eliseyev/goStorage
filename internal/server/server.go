package server

import (
	"github.com/gorilla/mux"
	"goStorage/internal/handlers"
	"net/http"
)

func Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/v1/{key}", handlers.KeyValuePutHandler).Methods(http.MethodPut)
	r.HandleFunc("/v1/{key}", handlers.KeyValueGetHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/{key}", handlers.KeyValueDeleteHandler).Methods(http.MethodDelete)

	return http.ListenAndServe(":8080", r)
}
