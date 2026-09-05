package api

import (
	"net/http"
)

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func RegisterAPIRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/api/storages", apiStoragesHandler)

}
