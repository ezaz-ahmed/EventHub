package models

import (
	"fmt"

	"github.com/ezaz-ahmed/EventHub/db"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {

	fmt.Println(user.Email, user.Password)

	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	user.ID = id

	return err
}
