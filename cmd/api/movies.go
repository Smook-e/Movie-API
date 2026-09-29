package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/Smook-e/Movie-API/internal/data"
	"github.com/danielgtaylor/huma/v2"
	"time"
)


func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}
func (app *application) createMovie(ctx context.Context, input *data.CreateMovieInput) (*data.CreateMovieOutput, error) {
	

	movie := &data.Movie{
		ID:        1,
		CreatedAt: time.Now(),
		Runtime:   data.Runtime(input.Body.Runtime),
		Title:     input.Body.Title,
		Year:      input.Body.Year,
		Genres:    input.Body.Genres,
		Version:   1,
	}

	return &data.CreateMovieOutput{Body: *movie}, nil
}

func (app *application) getMovie(ctx context.Context, input *data.GetMovieInput) (*data.GetMovieOutput, error) {
	if input.ID < 1 {
		return nil, huma.NewError(http.StatusBadRequest, "invalid movie ID")
	}
	movie := data.Movie{ID: input.ID,
		Title: "Drive", 
		Year: 2011, 
		Runtime: 100, 
		Genres: []string{"Action", "Adventure"}, 
		Version: 1, 
		CreatedAt: time.Now(),
	}

	return &data.GetMovieOutput{Body: movie}, nil
}


// func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	
// 	id, err := app.readIDParam(r)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	movie := data.Movie{
// 		ID: id,
// 		CreatedAt: time.Now(),
// 		Title: "Casablanca",	
// 		Year: 1942,
// 		Runtime: 102,
// 		Genres: []string{"drama", "romance", "war"},
// 	}
// 	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
// 	if err != nil {
// 		app.logger.Error("error writing JSON", "error", err)
// 		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
// 		return
// 	}
// }