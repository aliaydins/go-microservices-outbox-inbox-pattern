package order

import "errors"

var (
	ErrOrderNotCreated      = errors.New("Error occured when order saving to database ")
	ErrOutboxItemNotCreated = errors.New("Error occured when outbox item saving to database ")
)
