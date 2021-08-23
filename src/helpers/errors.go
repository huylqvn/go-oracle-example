package helpers

import "net/http"

var (
	ErrInvalidParams = errInvalidParams{}
)

type errInvalidParams struct{}

func (e errInvalidParams) Error() string {
	return "invalid parameters"
}
func (e errInvalidParams) StatusCode() int {
	return http.StatusBadRequest
}
