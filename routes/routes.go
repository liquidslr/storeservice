package routes

import "net/http"

// Route structure for url mapping
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes array of Route
type Routes []Route

var routes = Routes{
	Route{
		"GetValue",
		"GET",
		"/api/value/{key}",
		GetKeyValue,
	},
	Route{
		"SetValue",
		"POST",
		"/api/set/value/",
		SetKeyValue,
	},

	Route{
		"GetAll",
		"GET",
		"/api/get/all/",
		GetAllValues,
	},
}
