package service

import (
	"go-oracle/config"
	"go-oracle/src/service/cache"
	"go-oracle/src/service/minio"
	"go-oracle/src/service/repository"

	"github.com/go-kit/kit/log"
)

// Service ...
type Service struct {
	Logger log.Logger
	Config *config.Config
	DB     repository.Repository
	Store  minio.MinioService
	Cache  cache.CacheService
}
