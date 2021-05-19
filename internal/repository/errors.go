package repository

import "errors"

var ErrInsufficientFunds = errors.New("Insufficient funds")
var ErrIdempontenceKeyNoFound = errors.New("Idempotency key not found")
var ErrIdempotenceKeyAttemptLimit = errors.New("Exceeded the limit of attempts for the idempotence key")
