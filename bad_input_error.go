package ego

import "net/http"

type BadInputError struct {
	message string
}

func (e BadInputError) Error() string {
	return e.message
}

func NewBadInputError(message string) BadInputError {
	return BadInputError{
		message: message,
	}
}

func (e BadInputError) ToGqlErrorCode() string {
	return "BAD_USER_INPUT"
}

func (e BadInputError) ToControllerStatusCode() int {
	return http.StatusBadRequest
}
