package main

import (
	server "L0"
	"L0/pkg/handler"
	"L0/pkg/repository"
	"L0/pkg/service"
	"log"
)

func main() {
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
