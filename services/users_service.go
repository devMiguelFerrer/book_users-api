package services

import (
	"net/http"

	"github.com/devMiguelFerrer/book_users-api/domain/users"
	"github.com/devMiguelFerrer/book_users-api/utils/errors"
)

//CreateUser ...
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.ValidateEmail(); err != nil {
		return &user, nil
	}
	return nil, &errors.RestError{
		Message: "internal error",
		Status:  http.StatusInternalServerError,
		Error:   "internal_error",
	}
}
