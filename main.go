package main

import (
	"context"
	"log"
	"os"
)

func main() {
	InitEnv()
	db := ConnectDB()
	ctx := context.Background()
	_, err := db.ExecContext(ctx, "SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	server := NewServer(":" + os.Getenv("PORT"))
	server.Handle("GET", "/", server.AddMiddleware(HandlerHome, Logger()))
	server.Handle("POST", "/user", server.AddMiddleware(HandlerCreateUser, CheckAuth(), Logger()))
	server.Handle("GET", "/user", server.AddMiddleware(HandlerGetUsers, CheckAuth(), Logger()))
	server.Listen()
}
