package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type RouteDetails struct {
	Endpoint    string
	Method      string
	Middlewares []mux.MiddlewareFunc
	Handler     http.HandlerFunc
}

type ServerConfig struct {
	Host string `env:"HOST" validate:"required"`
	Port string `env:"PORT" validate:"required,numeric"`
}

type ServiceRouter struct {
	ServiceName       string
	CommonMiddlewares []mux.MiddlewareFunc
	Routes            []RouteDetails
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func (s *ServiceRouter) StartServer() {
	cfg := ServerConfig{
		Host: getenv("HOST", "localhost"),
		Port: getenv("PORT", "8080"),
	}

	router := mux.NewRouter()
	subRouterForEachService := router.PathPrefix("/" + s.ServiceName).Subrouter()

	for _, eachRoute := range s.Routes {
		r := subRouterForEachService.NewRoute().Subrouter()
		r.Use(s.CommonMiddlewares...)
		r.Use(eachRoute.Middlewares...)
		r.HandleFunc(eachRoute.Endpoint, eachRoute.Handler).Methods(eachRoute.Method)
	}

	httpServer := &http.Server{
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: router,
	}

	log.Println("Server start at port ", cfg.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("failed to start server at port ", cfg.Port)
	}
}
