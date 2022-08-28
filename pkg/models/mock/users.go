package mock

import (
	"time"

	"github.com/bavlayan/GoToDo/pkg/models"
)

var mocUser = &models.User{
	ID:      "2c1ef52b-24d7-46c0-ab4c-4125453378ba",
	Name:    "Test",
	Email:   "test@todoitems.com",
	Created: time.Now(),
}

type UserModel struct{}

func (u *UserModel) Insert(name, email, password string) error {
	switch email {
	case "duplicatemail@todoitems.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (u *UserModel) Authenticate(email, password string) (string, error) {
	switch email {
	case "test@todoitems.com":
		return "2c1ef52b-24d7-46c0-ab4c-4125453378ba", nil
	default:
		return "", models.ErrInvalidCredentials
	}
}

func (u *UserModel) Get(id string) (*models.User, error) {
	switch id {
	case "2c1ef52b-24d7-46c0-ab4c-4125453378ba":
		return mocUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
