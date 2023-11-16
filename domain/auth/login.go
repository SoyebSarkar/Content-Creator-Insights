package auth

import (
	"errors"
	"fmt"

	users "github.com/SoyebSarkar/content-creator-insight/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (login *Login) LoginValidate() error {
	if login.Email == "" {
		return errors.New("invalid email")
	}
	if login.Password == "" {
		return errors.New("invalid password")
	}
	return nil
}

func (login *Login) CheckValidPassword() (bool, error) {
	dbPassword, err := users.GetUserPasswordFromDB(login.Email)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("database error")
	}
	passwordByte := []byte(login.Password)
	passwordByte = append(passwordByte, []byte(login.Email)...)
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), passwordByte)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("wrong password")
	}
	return true, nil
}
