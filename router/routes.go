package router

import (
	"net/http"

	"github.com/ramrodo/tech-assessment-loan-startup/handler"
)

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array
type Routes []Route

var routes = Routes{
	Route{
		"CreditAssignment",
		"POST",
		"/credit-assignment",
		handler.CreditAssignment,
	},
}
