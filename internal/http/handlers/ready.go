// internal/http/handlers/ready.go
package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/bhushankp/app-api.git/internal/pkg/respond"
)

type DBPinger interface {
	PingContext(ctx context.Context) error
}

func Ready(p DBPinger, timeout time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := ""
		if v := r.Context().Value("reqID"); v != nil {
			if s, ok := v.(string); ok {
				reqID = s
			}
		}
		if p == nil {
			respond.Error(w, http.StatusServiceUnavailable, "db-pinger-not-configured", reqID)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		if err := p.PingContext(ctx); err != nil {
			respond.Error(w, http.StatusServiceUnavailable, "db-unavailable", reqID)
			return
		}
		respond.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}
}
