package main

func main() {
	server := NewServer(":8000")
	server.Handle("GET", "/", server.AddMiddleware(HandlerHome, Logger()))
	server.Handle("POST", "/user", server.AddMiddleware(HandlerCreateUser, CheckAuth(), Logger()))
	server.Handle("GET", "/user", server.AddMiddleware(HandlerGetUsers, CheckAuth(), Logger()))
	server.Listen()
}
