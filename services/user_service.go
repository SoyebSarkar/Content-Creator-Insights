package service

import (
	"errors"
	"fmt"

	users "github.com/SoyebSarkar/content-creator-insight/domain/user"
)

func CreateUser(user users.User) error {
	if err := user.Validate(); err != nil {
		fmt.Print("-->", err)
		return err
	}
	err, isExist := user.IsUserExist()
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("email id already exist")
	}
	if err := user.Save(); err != nil {
		return err
	}

	return nil
}
