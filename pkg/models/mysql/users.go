package mysql

import (
	"database/sql"
	"strings"

	"github.com/bavlayan/GoToDo/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	sql_query := `INSERT INTO tbl_users (id, name, email, hashed_password) VALUES(uuid(), ?, ?, ?)`
	_, err = m.DB.Exec(sql_query, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

func (m *UserModel) Authenticate(email, password string) (string, error) {
	var id string
	var hashedPassword []byte
	row := m.DB.QueryRow("SELECT id, hashed_password FROM tbl_users WHERE email = ?", email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return "", models.ErrInvalidCredentials
	} else if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return "", models.ErrInvalidCredentials
	} else if err != nil {
		return "", err
	}

	return id, nil
}

func (m *UserModel) Get(id string) (*models.User, error) {
	return nil, nil
}
