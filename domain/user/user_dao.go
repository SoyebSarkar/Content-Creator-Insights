package users

import (
	"github.com/SoyebSarkar/content-creator-insight/datasource/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	queryGetUserPassDB = "SELECT `password` FROM `users` WHERE `email` = ?"
	queryUserExist     = "SELECT COUNT(*) FROM `users` WHERE email = ?"
)

func HashPassword(password string, salt []byte) string {
	passBytes := []byte(password)
	passBytes = append(passBytes, salt...)
	hashedPass, _ := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	return string(hashedPass)
}

func GetUserPasswordFromDB(email string) (string, error) {
	var password string
	err := mysql.Client.QueryRow(queryGetUserPassDB, email).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
func CheckIfUserExist(email string) (error, bool) {
	var count int
	err := mysql.Client.QueryRow(queryUserExist, email).Scan(&count)
	if err != nil {
		return err, false
	}
	if count == 0 {
		return nil, false
	}
	return nil, true
}
