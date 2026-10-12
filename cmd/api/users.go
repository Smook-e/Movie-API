package main

import (
	"net/http"
	"context"
	"github.com/Smook-e/Movie-API/internal/data"
	"github.com/danielgtaylor/huma/v2"
)

func (app *application) createUser(ctx context.Context, input *data.UserCreateInput) (*data.UserCreateOutput, error) {
	user := &data.User{
		Name:      input.Body.Name,
		Email:     input.Body.Email,
		Activated: false,
	}
	if !data.ValidateEmail(input.Body.Email) {
		return nil, huma.NewError(http.StatusBadRequest, "Please provide a valid email address")
	}
	err := user.Password.Set(input.Body.Password)
	if err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "failed to set password")
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		if err == data.ErrDuplicateEmail {
			return nil, huma.NewError(http.StatusConflict, "email already exists")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to insert user")
	}
	output := &data.UserCreateOutput{}
	output.Body.User = *user
	return output, nil
}

func (app *application) getUser(ctx context.Context, input *data.UserGetIDInput) (*data.UserCreateOutput, error) {
	user, err := app.models.Users.Get(input.ID)
	if err != nil {
		if err == data.ErrRecordNotFound {
			return nil, huma.NewError(http.StatusNotFound, "user not found")
		}
		return nil, huma.NewError(http.StatusInternalServerError, "failed to get user")
	}
	output := &data.UserCreateOutput{}
	output.Body.User = *user
	return output, nil
}

func (app *application) listUsers(ctx context.Context, _ *struct{}) (*data.UserGetAllOutput, error) {
	users, err := app.models.Users.GetAll()
	if err != nil {
		return nil, huma.NewError(http.StatusInternalServerError, "failed to list users")
	}
	output := &data.UserGetAllOutput{}
	output.Body.Users = users
	return output, nil
}
