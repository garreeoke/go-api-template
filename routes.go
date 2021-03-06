package main

import (
	"net/http"
)

// Route connection info
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type to store routes
type Routes []Route

// ROUTES
var routes = Routes{
	Route{
		"Healthz",
		"GET",
		"/v1/sample-go-api/healthz",
		Healthz,
	},
	Route{
		"PostExample",
		"POST",
		"/v1/sample-go-api/record",
		PostExample,
	},
}
