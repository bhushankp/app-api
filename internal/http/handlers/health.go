package handlers

import (
	"net/http"

	"github.com/bhushankp/app-api.git/internal/pkg/respond"
)

func Health(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "service is alive",
	})
}
