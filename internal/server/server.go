package server

import (
	"context"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/npmanos/4176Gameday-backend/internal/config"
	ihttp "github.com/npmanos/4176Gameday-backend/internal/http"
	"github.com/npmanos/4176Gameday-backend/internal/store"
	"github.com/npmanos/4176Gameday-backend/internal/tba"
	"github.com/sirupsen/logrus"
)

//go:generate go run ../cmd/pack/pack.go -package server -in openapi.yaml -out openapi.go -name openAPI
var openAPI []byte

// Server is the scouting API server
type Server struct {
	config.Server

	TBA    *tba.Service
	Store  *store.Service
	Logger *logrus.Logger
	start  time.Time
}

func (s *Server) uptime() time.Duration {
	return time.Since(s.start)
}

// Run starts the server, and returns if it runs into an error
func (s *Server) Run(ctx context.Context) error {
	router := s.registerRoutes()

	var handler http.Handler = router
	handler = ihttp.LimitBody(handler)
	handler = gziphandler.GzipHandler(handler)
	handler = ihttp.Log(handler, s.Logger)
	handler = ihttp.Auth(handler, s.JWTSecret)
	handler = ihttp.CORS(handler, s.Origin)

	httpServer := &http.Server{
		Addr:              s.Listen,
		Handler:           handler,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: time.Second * 15,
		WriteTimeout:      time.Second * 15,
		IdleTimeout:       time.Second * 30,
		MaxHeaderBytes:    4096,
	}

	s.start = time.Now()

	errs := make(chan error)
	go func() {
		s.Logger.WithField("httpAddress", s.Listen).Info("serving http")
		errs <- httpServer.ListenAndServe()
	}()

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		return httpServer.Shutdown(ctx)
	}
}
