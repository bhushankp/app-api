package handlers

import (
	"net/http"

	"github.com/yourname/go-backend-sprint1/pkg/respond"
)

type DBPinger interface {
	Ping() error
}

func Ready(db DBPinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			respond.JSON(w, http.StatusServiceUnavailable, map[string]string{
				"status":  "error",
				"message": "database not ready",
			})
			return
		}
		respond.JSON(w, http.StatusOK, map[string]string{
			"status":  "ok",
			"message": "service is ready",
		})
	}
}
