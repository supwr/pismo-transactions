package transaction

import "errors"

var (
	ErrOperationTypeNotFound = errors.New("Operation Type not found")
	ErrAccountNotFound       = errors.New("Account not found")
	ErrInsuficientFunds      = errors.New("Insuficient funds")
)
