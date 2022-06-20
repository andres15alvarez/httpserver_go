package main

import (
	"os"

	"github.com/andres15alvarez/go_http_server/handlers"
	"github.com/andres15alvarez/go_http_server/middlewares"
	"github.com/andres15alvarez/go_http_server/utils"
)

func main() {
	utils.InitEnvironment()
	server := utils.NewServer(":" + os.Getenv("PORT"))
	server.Handle("GET", "/", server.AddMiddleware(handlers.Home, middlewares.Logger()))
	server.Handle("POST", "/user", server.AddMiddleware(handlers.CreateUser, middlewares.Logger()))
	server.Handle("GET", "/user", server.AddMiddleware(handlers.ListUsers, middlewares.CheckAuth(), middlewares.Logger()))
	server.Listen()
}
