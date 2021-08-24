package repository

import (
	"github.com/go-kit/kit/log"
	"xorm.io/xorm"
)

type dao struct {
	db     *xorm.Engine
	logger log.Logger
}

func NewRepository(db *xorm.Engine, logger log.Logger) Repository {
	return &dao{db, logger}
}
