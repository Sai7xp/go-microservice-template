package dao

import "go.uber.org/zap"

type DaoService interface {
	// Dao Methods - SaveCat, GetCat
}

type dao struct {
	logger *zap.Logger //TODO: abstract out this logger, what if we want to use another package
	// daoClient DBService // TODO: db service from common utils
}

func NewDaoService(logger *zap.Logger) DaoService {
	return &dao{
		logger: logger,
	}
}

// TODO: dao struct will implement all the methods specified in DaoService interface
