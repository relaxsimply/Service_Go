package main

import (
	"fmt"
	"service/http"
	"service/todo"
)

func main() {
	todolist := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todolist)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}

}
