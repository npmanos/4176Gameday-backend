package server

import (
	"context"
	"net/http"
	"time"

	ihttp "github.com/npmanos/4176Gameday-backend/internal/http"
)

func openAPIHandler(openAPI []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		_, _ = w.Write(openAPI)
	}
}

// Pinger defines an interface for pinging a service to see if
// it is alive.
type Pinger interface {
	Ping(ctx context.Context) error
}

type healthServices struct {
	TBA        bool `json:"tba"`
	PostgreSQL bool `json:"postgresql"`
}

type healthStatus struct {
	Uptime   string         `json:"uptime"`
	Services healthServices `json:"services"`
	Ok       bool           `json:"ok"`
}

func healthHandler(getUptime func() time.Duration, tba, postgres Pinger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		services := healthServices{
			TBA:        tba.Ping(r.Context()) == nil,
			PostgreSQL: postgres.Ping(r.Context()) == nil,
		}

		ihttp.Respond(w, healthStatus{
			Uptime:   getUptime().String(),
			Services: services,
			Ok:       services.TBA && services.PostgreSQL,
		}, http.StatusOK)
	}
}
