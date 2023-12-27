package models

import (
	"fmt"
	"github.com/eliofery/golang-restapi/database"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	op := "user.Save"

	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	u.ID = int(userId)

	return nil
}
