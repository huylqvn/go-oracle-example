package app

import (
	"context"
	"go-oracle/src/service"

	"github.com/go-kit/kit/endpoint"
)

func HealthCheckHandler(s *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return true, nil
	}
}
