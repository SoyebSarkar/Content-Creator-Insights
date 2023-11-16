package service

import (
	"fmt"

	users "github.com/SoyebSarkar/content-creator-insight/domain/user"
)

func CreateUser(user users.User) error {
	if err := user.Validate(); err != nil {
		fmt.Print("-->", err)
		return err
	}
	if err := user.Save(); err != nil {
		return err
	}
	return nil
}
