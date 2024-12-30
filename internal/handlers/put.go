package handlers

import (
	"goStorage/internal/storage"
	"goStorage/internal/transaction_logger"
	"io"
	"net/http"
)

func KeyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	key := GetKey(r)

	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stringValue := string(value)
	err = storage.Put(key, stringValue)
	HandleCommonError(err, w)

	transaction_logger.Logger.WritePut(key, stringValue)

	w.WriteHeader(http.StatusCreated)
}
