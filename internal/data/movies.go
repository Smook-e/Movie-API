package data

import (
	"time"
	"database/sql"
	"github.com/lib/pq"

)
type Movie struct {
	ID int64 			`json:"id"`
	CreatedAt time.Time `json:"-"`
	Runtime Runtime		`json:"runtime" doc:"Runtime in minutes"`
	Title string 		`json:"title"`
	Year int32 			`json:"year"`
	Genres []string 	`json:"genres"`
	Version int32 		`json:"version,omitzero"`
}
type MovieInput struct {
	Title string 		`json:"title" minLength:"2" maxLength:"100"`
	Year int32 			`json:"year" minimum:"1888" maximum:"2100"`
	Runtime int32		`json:"runtime" minimum:"1" maximum:"600"`
	Genres []string 	`json:"genres,omitzero"`
}
type CreateMovieInput struct {
	Body MovieInput 
}
type UpdateMovieInput struct {
	Body MovieInput
	ID int64 `path:"id" doc:"Movie ID"`
}

type PatchMovieInput struct {
	Body struct {
		Title *string 		`json:"title,omitempty" minLength:"2" maxLength:"100"`
		Year *int32 		`json:"year,omitempty" minimum:"1888" maximum:"2100"`
		Runtime *int32		`json:"runtime,omitempty" minimum:"1" maximum:"600"`
		Genres []string 	`json:"genres,omitempty"`
	}
	ID int64 `path:"id" doc:"Movie ID"`
}

type CreateMovieOutput struct {
	Body Movie 
}

type GetMovieInput struct {
    ID int64 `path:"id" doc:"Movie ID"`
}

type GetMovieOutput struct {
	Body Movie 
}


type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
	INSERT INTO movies (title, year, runtime, genres)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}
	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	
	query := `
	SELECT id, created_at, title, year, runtime, genres, version
	FROM movies
	WHERE id = $1`

	var movie Movie
	err := m.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)
	if err != nil {
		switch  err{
		case sql.ErrNoRows:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	
	return &movie, nil
}

func (m MovieModel) Update(movie *Movie) error {

	query := `
	UPDATE movies
	SET title = $1, year = $2, runtime = $3, genres = $4, version = version + 1
	WHERE id = $5 AND version = $6 
	RETURNING version`

	args := []any{
		movie.Title,
		movie.Year,
		movie.Runtime,
		pq.Array(movie.Genres),
		movie.ID,
		movie.Version,
	}

	err := m.DB.QueryRow(query, args...).Scan(&movie.Version)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m MovieModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM movies
	WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}