package service

import (
	"github.com/sai7xp/go-microservice-template/cat-service/internal/dao"
)

type CatService interface {
	// all the methods
}

type catService struct {
	dao dao.DaoService
}

func NewCatService(dao dao.DaoService) CatService {
	return &catService{
		dao: dao,
	}
}
