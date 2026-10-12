package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("Record not found")
	ErrEditConflict   = errors.New("Edit conflict. Please try again")
)

type Models struct {
	Movies MovieModel
	Users UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users: UserModel{DB: db},
	}
}