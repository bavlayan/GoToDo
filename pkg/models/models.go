package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type TodoItems struct {
	ID          string
	Completed   bool
	CreatedDate time.Time
	Description string
	Deleted     bool
}
