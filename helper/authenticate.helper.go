package helper

import (
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

func HashPassword(password string) (string, error) {
	passwordTrim := html.EscapeString(strings.TrimSpace(password))
	hashedString, err := bcrypt.GenerateFromPassword([]byte(passwordTrim), bcrypt.DefaultCost)
	if err != nil {
		hashedString = []byte("")
	}
	return string(hashedString), err
}

//func CreateToken(password string) string {
//
//}
