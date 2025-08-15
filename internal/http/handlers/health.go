package handlers

import (
	"net/http"
	"time"

	"github.com/bhushankp/app-api.git/internal/pkg/respond"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildTime = "unknown"
)

func Health(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"status":     "ok",
		"version":    Version,
		"commit":     Commit,
		"build_time": BuildTime,
		"ts":         time.Now().UTC().Format(time.RFC3339),
	}
	respond.JSON(w, http.StatusOK, resp)
}
