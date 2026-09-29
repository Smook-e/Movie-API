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
		OperationID: "create-movie",
		Method:      http.MethodPost,
		Path:        "/v1/movies",
		Summary:     "Create a new movie",
		Tags:        []string{"movies"},
		Description: "Create a new movie",
	}, app.createMovie)
	// router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)

	return router

}
