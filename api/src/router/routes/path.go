package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Path represents all routes from the API
type Path struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// Configure config all routes in API
func Configure(r *mux.Router) *mux.Router {

	allPaths := userPaths

	for _, path := range allPaths {
		r.HandleFunc(path.URI, path.Function).Methods(path.Method)
	}

	return r
}
