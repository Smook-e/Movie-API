package data

import (
	"time"
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
type CreateMovieOutput struct {
	Body Movie 
}

type GetMovieInput struct {
    ID int64 `path:"id" doc:"Movie ID"`
}

type GetMovieOutput struct {
	Body Movie 
}