package path

import (
	"api/src/controller"
	"net/http"
)

const basePath = "/users"
const byIdPath = basePath + "/{userId}"

var userPaths = []Path{
	{
		URI:          basePath,
		Method:       http.MethodPost,
		Function:     controller.Create,
		RequiresAuth: false,
	},
	{
		URI:          basePath,
		Method:       http.MethodGet,
		Function:     controller.List,
		RequiresAuth: false,
	},
	{
		URI:          byIdPath,
		Method:       http.MethodGet,
		Function:     controller.ById,
		RequiresAuth: false,
	},
	{
		URI:          byIdPath,
		Method:       http.MethodPut,
		Function:     controller.Update,
		RequiresAuth: false,
	},
	{
		URI:          byIdPath,
		Method:       http.MethodDelete,
		Function:     controller.Delete,
		RequiresAuth: false,
	},
}
