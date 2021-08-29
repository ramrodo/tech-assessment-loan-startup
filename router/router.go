package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramrodo/tech-assessment-loan-startup/handler"
)

// NewRouter - returns a router object with routes
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var httpHandler http.Handler

		httpHandler = route.HandlerFunc
		httpHandler = handler.Logger(httpHandler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(httpHandler)
	}

	return router
}
