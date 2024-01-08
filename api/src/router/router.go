package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// GetRouter returns a new router with all routes configured
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
