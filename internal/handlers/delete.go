package handlers

import (
	"goStorage/internal/storage"
	"goStorage/internal/transaction_logger"
	"net/http"
)

func KeyValueDeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := GetKey(r)

	err := storage.Delete(key)
	HandleCommonError(err, w)

	transaction_logger.Logger.WriteDelete(key)

	w.WriteHeader(http.StatusOK)
}
