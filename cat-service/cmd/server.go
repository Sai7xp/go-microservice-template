package server

import (
	"github.com/sai7xp/go-microservice-template/cat-service/internal/controllers"
	"github.com/sai7xp/go-microservice-template/cat-service/internal/dao"
	"github.com/sai7xp/go-microservice-template/cat-service/internal/service"
	"go.uber.org/zap"
)

func Serve() {
	logger, _ := zap.NewDevelopment()
	dao := dao.NewDaoService(logger)
	srv := service.NewCatService(dao)
	router := controllers.NewController(logger, srv).Router("cat-service")
	router.StartServer()
}
