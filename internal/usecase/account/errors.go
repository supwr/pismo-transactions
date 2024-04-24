package account

import "errors"

var (
	ErrAccountAlreadyExists = errors.New("There's already an account with this document")
)
