package main

import (
	"context"
	"fmt"
	"service/http"
	"service/postgres/simple_connection"
	"service/postgres/simple_sql"
	"service/todo"
)

func main() {

	//"postgres://"

	todolist := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todolist)
	httpServer := http.NewHTTPServer(httpHandlers)

	ctx := context.Background()

	conn, err := simple_connection.CreateConnetcion(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	// if err := simple_sql.InsertRow(
	// 	ctx,
	// 	conn,
	// 	"Обед",
	// 	"Покушать",
	// 	false,
	// 	time.Now(),
	// ); err != nil {
	// 	panic(err)
	// }

	if err := simple_sql.UpdateRow(ctx, conn); err != nil {
		panic(err)
	}

	if err := simple_sql.DeleteRow(ctx, conn); err != nil {
		panic(err)
	}

	fmt.Println("succeed!")

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}

}
