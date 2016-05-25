package route

import (
	"net/http"

	"app/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controller.TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controller.TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		controller.TodoShow,
	},
}
