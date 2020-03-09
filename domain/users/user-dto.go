package users

import (
	"fmt"

	"github.com/devMiguelFerrer/book_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user User) Get() *errors.RestError {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	user.ID = result.ID
	return nil
}

func (user User) Save() *errors.RestError {
	return nil
}
