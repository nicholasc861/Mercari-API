package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/nicholasc861/mercari-api/app/handler"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
	Route {
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route {
		"GetProductsByKeyword",
		"GET",
		"/products/{keyword}",
		handler.GetProductsByKeyword,
	},
	Route {
		"GetProductById",
		"GET",
		"/product/{id}",
		handler.GetProductById,
	},
	Route {
		"GetUserById",
		"GET",
		"/user/{id}",
		handler.GetUserById,
	},
}
