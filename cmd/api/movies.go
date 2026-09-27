package main

import (
	"fmt"
	"net/http"
	"github.com/Smook-e/Movie-API/internal/data"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	
	id, err := app.readIDParam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "Casablanca",	
		Year: 1942,
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Error("error writing JSON", "error", err)
		http.Error(w, "Server encountered an error", http.StatusInternalServerError)
		return
	}
}