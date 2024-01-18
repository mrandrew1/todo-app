package main

import (
	"github.com/mrandrew1/todo-app"
	"github.com/mrandrew1/todo-app/pkg/handler"
	"github.com/mrandrew1/todo-app/pkg/repository"
	"github.com/mrandrew1/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %v", err.Error())
	}
}
