package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is a collection of the Route structure
type Routes []Route

//NewRouter creates a new Mux Router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"HelloWorld",
		"GET",
		"/",
		HelloWorld,
	},
	Route{
		"Send",
		"POST",
		"/send",
		Send,
	},
}
