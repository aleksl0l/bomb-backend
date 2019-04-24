package http

// TODO: Add possibility to check needed methods for views

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type MyRouter struct {
	*mux.Router
}

func NewRouter() *MyRouter {
	return &MyRouter{mux.NewRouter()}
}

func (r *MyRouter) HandleView(s string, f func(*json.Decoder) (int, interface{})) *mux.Route {
	return r.HandleFunc(s, func(w http.ResponseWriter, r *http.Request) {
		statusCode, response := f(ParseJSON(r.Body))
		genericResponse(statusCode, w, response)
	})
}
