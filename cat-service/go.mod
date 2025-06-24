module github.com/sai7xp/go-microservice-template/cat-service

go 1.24.0

replace github.com/sai7xp/go-microservice-template/common-utils => ../common-utils

require (
	github.com/sai7xp/go-microservice-template/common-utils v0.0.0-20250619070154-4225c64125ef
	go.uber.org/zap v1.27.0
)

require (
	github.com/gorilla/mux v1.8.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)
