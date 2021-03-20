package main

import (
	"github.com/davex98/nobl9-backend/random-generator/application"
	"github.com/davex98/nobl9-backend/random-generator/common/log"
	"github.com/davex98/nobl9-backend/random-generator/common/server"
	"github.com/davex98/nobl9-backend/random-generator/controllers"
	"github.com/davex98/nobl9-backend/random-generator/infrastracture"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

func main() {
	log.Init()
	client := http.Client{Timeout: time.Second * 10}
	randomService := infrastracture.NewRandomService(client)
	numberService := application.NewNumberService(randomService)
	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return controllers.HandlerFromMux(controllers.NewHttpServer(numberService), router)
	})
}
