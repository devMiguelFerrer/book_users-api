package users

import "github.com/devMiguelFerrer/book_users-api/utils/errors"

//User ...
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreateAt  string `json:"dateCreated"`
}

//ValidateEmail ...
func (user *User) ValidateEmail() *errors.RestError {
	return nil
}
