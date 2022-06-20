package utils

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port   string
	Router *Router
}

func NewServer(port string) *Server {
	return &Server{
		Port:   port,
		Router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.Router.Rules[path]
	if !exist {
		s.Router.Rules[path] = make(map[string]http.HandlerFunc)
	}
	s.Router.Rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.Router)
	fmt.Printf("Server running in port %s\n", s.Port)
	err := http.ListenAndServe(s.Port, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}
