package main

import (
	"fmt"
	"service/http"
	"service/postgres/simple_connection"
	"service/todo"
)

func main() {

	//"postgres://"

	todolist := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todolist)
	httpServer := http.NewHTTPServer(httpHandlers)

	simple_connection.CheckConnetcion()

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}

}
