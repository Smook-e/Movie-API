package main

import (
	"net/http"
	"context"
	"github.com/Smook-e/Movie-API/internal/data"
	"github.com/danielgtaylor/huma/v2"
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)
func (app *application) createMovie(ctx context.Context, input *data.CreateMovieInput) (*data.CreateMovieOutput, error) {
	

	movie := &data.Movie{
		Runtime:   data.Runtime(input.Body.Runtime),
		Title:     input.Body.Title,
		Year:      input.Body.Year,
		Genres:    input.Body.Genres,
	}
	err := app.models.Movies.Insert(movie)
	if err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "failed to insert movie")
	}

	return &data.CreateMovieOutput{Body: *movie}, nil
}

func (app *application) getMovie(ctx context.Context, input *data.GetMovieInput) (*data.GetMovieOutput, error) {
	movie, err := app.models.Movies.Get(input.ID)
	if err != nil {
		if err == data.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "movie not found")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to get movie")
	}

	return &data.GetMovieOutput{Body: *movie}, nil
}

func (app *application) listMovies(ctx context.Context, input *data.ListMoviesInput) (*data.ListMoviesOutput, error) {
	movies, metadata, err := app.models.Movies.GetAll(input.Title, input.Genres, input.Filter)
	if err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "failed to list movies")
	}

	output := &data.ListMoviesOutput{}
	output.Body.Movies = movies
	output.Body.Metadata = metadata

	return output, nil
}

func (app *application) updateMovie(ctx context.Context, input *data.UpdateMovieInput) (*data.CreateMovieOutput, error) {

	movie ,err := app.models.Movies.Get(input.ID)
	if err != nil {
		if err == data.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "movie not found")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to get movie")
	}

	movie.Title = input.Body.Title
	movie.Year = input.Body.Year
	movie.Runtime = data.Runtime(input.Body.Runtime)
	movie.Genres = input.Body.Genres

	err = app.models.Movies.Update(movie)
	if err != nil {
		if err == data.ErrEditConflict {
			return nil, huma.NewError(http.StatusConflict, "edit conflict. Please try again")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to update movie")
	}

	return &data.CreateMovieOutput{Body: *movie}, nil
}
func (app *application) patchMovie(ctx context.Context, input *data.PatchMovieInput) (*data.CreateMovieOutput, error) {
	movie ,err := app.models.Movies.Get(input.ID)
	if err != nil {
		if err == data.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "movie not found")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to get movie")
	}

	if input.Body.Title != nil {
		movie.Title = *input.Body.Title
	}
	if input.Body.Year != nil {
		movie.Year = *input.Body.Year
	}
	if input.Body.Runtime != nil {
		movie.Runtime = data.Runtime(*input.Body.Runtime)
	}
	if input.Body.Genres != nil {
		movie.Genres = input.Body.Genres
	}

	err = app.models.Movies.Update(movie)
	if err != nil {
		if err == data.ErrEditConflict {
			return nil, huma.NewError(http.StatusConflict, "edit conflict. Please try again")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to update movie")
	}

	return &data.CreateMovieOutput{Body: *movie}, nil
}

func (app *application) deleteMovie(ctx context.Context, input *data.GetMovieInput) ( *struct{},error) {
	err := app.models.Movies.Delete(input.ID)
	if err != nil {
		if err == data.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "movie not found")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to delete movie")
	}

	return nil, nil
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