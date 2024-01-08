package router

import "github.com/gorilla/mux"

// GetRouter returns a new router with all routes configured
func GetRouter() *mux.Router {
	return mux.NewRouter()
}
