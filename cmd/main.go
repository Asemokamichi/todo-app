package main

import (
	"log"

	"github.com/Asemokamichi/todo-app"
	"github.com/Asemokamichi/todo-app/pkg/handler"
	"github.com/Asemokamichi/todo-app/pkg/repository"
	"github.com/Asemokamichi/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := todo.Server{}
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
