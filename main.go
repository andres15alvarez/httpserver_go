package main

import "os"

func main() {
	server := NewServer(":" + os.Getenv("PORT"))
	server.Handle("GET", "/", server.AddMiddleware(HandlerHome, Logger()))
	server.Handle("POST", "/user", server.AddMiddleware(HandlerCreateUser, CheckAuth(), Logger()))
	server.Handle("GET", "/user", server.AddMiddleware(HandlerGetUsers, CheckAuth(), Logger()))
	server.Listen()
}
