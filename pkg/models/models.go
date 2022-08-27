package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentiasl")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type TodoItems struct {
	ID          string
	Completed   bool
	CreatedDate time.Time
	Description string
	Deleted     bool
}

type User struct {
	ID             string
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
