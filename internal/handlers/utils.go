package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetKey(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["key"]
}

func HandleCommonError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
