package handlers

import (
	"errors"
	"goStorage/internal/storage"
	"net/http"
)

func KeyValueGetHandler(w http.ResponseWriter, r *http.Request) {
	key := GetKey(r)

	value, err := storage.Get(key)
	if errors.Is(err, storage.ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	HandleCommonError(err, w)

	w.Write([]byte(value))
}
