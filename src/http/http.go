package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"go-oracle/src/endpoints"
	"go-oracle/src/service"
)

// NewHTTPHandler ...
func NewHTTPHandler(s service.Service, endpoints endpoints.Endpoints, logger log.Logger, useCORS bool) http.Handler {
	r := chi.NewRouter()

	if useCORS {
		// cors := cors.New(cors.Options{
		// 	AllowedOrigins:   []string{"*"},
		// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// 	AllowCredentials: true,
		// })
		// r.Use(cors.Handler)
		cors := cors.AllowAll()
		r.Use(cors.Handler)
	}
	r.Use(middleware.Recoverer)

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", httptransport.NewServer(
			endpoints.HealthCheck,
			DecodeNullRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	return r
}

func DecodeNullRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
