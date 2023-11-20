package users

import (
	"errors"

	"github.com/SoyebSarkar/content-creator-insight/datasource/mysql"
)

type User struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

const (
	queryInsertUser = "INSERT INTO `users`(`id`, `name`, `email`, `password`) VALUES (NULL,?,?,?)"
)

func (user *User) Validate() error {
	if user.Email == "" {
		return errors.New("invalid Email")
	}
	if user.Password == "" || user.ConfirmPassword == "" || user.Password != user.ConfirmPassword {
		return errors.New("invalid password")
	}
	if user.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}
func (user *User) IsUserExist() (error, bool) {
	err, isExist := CheckIfUserExist(user.Email)
	if err != nil {
		return err, false
	}
	return nil, isExist
}
func (user *User) Save() error {
	hashPassword := HashPassword(user.Password, []byte(user.Email))
	_, err := mysql.Client.Exec(queryInsertUser, user.Name, user.Email, hashPassword)
	if err != nil {
		return err
	}
	return nil
}
