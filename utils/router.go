package utils

import "net/http"

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	_, exist := r.Rules[path]
	handler, methodExist := r.Rules[path][method]
	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.Method, request.URL.Path)
	if !methodExist {
		w.WriteHeader((http.StatusMethodNotAllowed))
		return
	}
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
