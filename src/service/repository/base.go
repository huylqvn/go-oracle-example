package repository

import (
	"github.com/go-kit/kit/log"

	"gorm.io/gorm"
)

type dao struct {
	db     *gorm.DB
	logger log.Logger
}

func NewRepository(db *gorm.DB, logger log.Logger) Repository {
	return &dao{db, logger}
}
