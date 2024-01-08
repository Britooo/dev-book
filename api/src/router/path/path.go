package path

import "net/http"

// Path represents all routes from the API
type Path struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}
