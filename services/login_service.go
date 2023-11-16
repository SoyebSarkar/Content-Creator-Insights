package service

import "github.com/SoyebSarkar/content-creator-insight/domain/auth"

func LoginValidate(login *auth.Login) (bool, error) {
	if err := login.LoginValidate(); err != nil {
		return false, err
	}
	result, err := login.CheckValidPassword()
	if err != nil {
		return false, err
	}
	return result, nil
}
