package main

import (
	"github.com/mrandrew1/todo-app"
	"github.com/mrandrew1/todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %v", err.Error())
	}
}
