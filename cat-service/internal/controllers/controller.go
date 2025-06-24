package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sai7xp/go-microservice-template/cat-service/internal/service"
	"github.com/sai7xp/go-microservice-template/common-utils/router"
	"go.uber.org/zap"
)

type Controller struct {
	logger  *zap.Logger
	service service.CatService
	router  *router.ServiceRouter
}

func NewController(logger *zap.Logger, srv service.CatService) *Controller {
	return &Controller{
		logger:  logger,
		service: srv,
	}
}

func (c *Controller) Router(serviceName string) *router.ServiceRouter {
	return &router.ServiceRouter{
		ServiceName:       serviceName,
		CommonMiddlewares: []mux.MiddlewareFunc{},
		Routes: []router.RouteDetails{
			{
				Endpoint: "/hello",
				Method:   http.MethodGet,
				Handler: func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("Hi, I'm a Cat"))
				},
			},
		},
	}
}
