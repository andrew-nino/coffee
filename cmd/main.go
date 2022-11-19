package main

import (
	"coffee-app"
	"coffee-app/pkg/handler"
	"coffee-app/pkg/repository"
	"coffee-app/pkg/service"
	"log"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	
	srv := new(coffee.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server : %s", err.Error())
	}
}
