package helper

import "golang.org/x/crypto/bcrypt"

func BcryptPassword(passwordSalt string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordSalt), 8)
	return string(newPassword)
}
