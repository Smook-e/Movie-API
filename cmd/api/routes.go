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

	huma.Register(api, huma.Operation{
		OperationID: "get-movie",
		Method:      http.MethodGet,
		Path:        "/v1/movies/{id}",
		Summary:     "Get a movie by ID",
		Tags:        []string{"movies"},
		Description: "Retrieve a movie by its ID",
	}, app.getMovie)

	huma.Register(api, huma.Operation{
		OperationID: "list-movies",
		Method:      http.MethodGet,
		Path:        "/v1/movies",
		Summary:     "List movies",
		Tags:        []string{"movies"},
		Description: "List movies with optional filtering and sorting",
	}, app.listMovies)
	
	huma.Register(api, huma.Operation{
		OperationID: "create-movie",
		Method:      http.MethodPost,
		Path:        "/v1/movies",
		Summary:     "Create a new movie",
		Tags:        []string{"movies"},
		Description: "Create a new movie",
		DefaultStatus: http.StatusCreated,
	}, app.createMovie)

	huma.Register(api, huma.Operation{
		OperationID: "update-movie",
		Method:      http.MethodPut,
		Path:        "/v1/movies/{id}",
		Summary:     "Update a movie",
		Tags:        []string{"movies"},
		Description: "Update a movie by its ID",
		Errors: 	 []int{http.StatusNotFound, http.StatusConflict},
	}, app.updateMovie)
	
	huma.Register(api, huma.Operation{
		OperationID: "patch-movies",
		Method:      http.MethodPatch,
		Path:        "/v1/movies/{id}",
		Summary:     "Partially update a movie",
		Tags:        []string{"movies"},
		Description: "Partially update a movie by its ID",
		Errors: 	 []int{http.StatusNotFound, http.StatusConflict},
	}, app.patchMovie)
	
	huma.Register(api, huma.Operation{
		OperationID: "delete-movie",
		Method:      http.MethodDelete,
		Path:        "/v1/movies/{id}",
		Summary:     "Delete a movie",
		Tags:        []string{"movies"},
		Description: "Delete a movie by its ID",
		DefaultStatus: http.StatusNoContent,
	}, app.deleteMovie)



	return app.recoverPanic(router)

}
