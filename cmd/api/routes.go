package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
    "github.com/danielgtaylor/huma/v2/adapters/humago"

)

func (app *application) routes() http.Handler {
	
	router := http.NewServeMux()
    api := humago.New(router, huma.DefaultConfig("Movie API", "1.0.0"))

	huma.Register(api, huma.Operation{
		OperationID: "healthcheck",
		Method:      http.MethodGet,
		Path:        "/v1/healthcheck",
		Summary:     "Check the health of the API",
		Tags:        []string{"health"},
		Description: "Check the health of the API",
	}, app.healthcheck)

	app.RegisterMovieRoutes(&api)
	app.RegisterUserRoutes(&api)

	return app.recoverPanic(router)

}

func (app *application) RegisterMovieRoutes(api *huma.API) {
	huma.Register(*api, huma.Operation{
		OperationID: "get-movie",
		Method:      http.MethodGet,
		Path:        "/v1/movies/{id}",
		Summary:     "Get a movie by ID",
		Tags:        []string{"movies"},
		Description: "Retrieve a movie by its ID",
	}, app.getMovie)

	huma.Register(*api, huma.Operation{
		OperationID: "list-movies",
		Method:      http.MethodGet,
		Path:        "/v1/movies",
		Summary:     "List movies",
		Tags:        []string{"movies"},
		Description: "List movies with optional filtering and sorting",
	}, app.listMovies)
	
	huma.Register(*api, huma.Operation{
		OperationID: "create-movie",
		Method:      http.MethodPost,
		Path:        "/v1/movies",
		Summary:     "Create a new movie",
		Tags:        []string{"movies"},
		Description: "Create a new movie",
		DefaultStatus: http.StatusCreated,
	}, app.createMovie)

	huma.Register(*api, huma.Operation{
		OperationID: "update-movie",
		Method:      http.MethodPut,
		Path:        "/v1/movies/{id}",
		Summary:     "Update a movie",
		Tags:        []string{"movies"},
		Description: "Update a movie by its ID",
		Errors: 	 []int{http.StatusNotFound, http.StatusConflict},
	}, app.updateMovie)
	
	huma.Register(*api, huma.Operation{
		OperationID: "patch-movies",
		Method:      http.MethodPatch,
		Path:        "/v1/movies/{id}",
		Summary:     "Partially update a movie",
		Tags:        []string{"movies"},
		Description: "Partially update a movie by its ID",
		Errors: 	 []int{http.StatusNotFound, http.StatusConflict},
	}, app.patchMovie)
	
	huma.Register(*api, huma.Operation{
		OperationID: "delete-movie",
		Method:      http.MethodDelete,
		Path:        "/v1/movies/{id}",
		Summary:     "Delete a movie",
		Tags:        []string{"movies"},
		Description: "Delete a movie by its ID",
		DefaultStatus: http.StatusNoContent,
	}, app.deleteMovie)

}

func (app *application) RegisterUserRoutes(api *huma.API) {
	huma.Register(*api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/v1/users",
		Summary:     "Create a new user",
		Tags:        []string{"users"},
		Description: "Create a new user",
		DefaultStatus: http.StatusCreated,
	}, app.createUser)

	huma.Register(*api, huma.Operation{
		OperationID: "get-user",
		Method:	  http.MethodGet,
		Path:        "/v1/users/{id}",
		Summary:     "Get a user by ID",
		Tags:        []string{"users"},
		Description: "Retrieve a user by its ID",
	}, app.getUser)

	huma.Register(*api, huma.Operation{
		OperationID: "list-users",
		Method:	  	 http.MethodGet,
		Path:        "/v1/users",
		Summary:     "List users",
		Tags:        []string{"users"},
		Description: "List users",
	}, app.listUsers)
}