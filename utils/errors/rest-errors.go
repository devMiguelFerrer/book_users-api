package errors

import "net/http"

//RestError ...
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

//NewBadRequestError ...
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError ...
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
