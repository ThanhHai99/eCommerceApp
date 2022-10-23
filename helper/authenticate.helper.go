package helper

import (
	"eCommerce/config"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

var (
	jwtConfigs = config.JwtConfigs{}
	_          = env.Parse(&jwtConfigs)
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	passwordTrim := html.EscapeString(strings.TrimSpace(password))
	hashedString, err1 := bcrypt.GenerateFromPassword([]byte(passwordTrim), bcrypt.DefaultCost)
	if err1 != nil {
		hashedString = []byte("")
	}
	return string(hashedString), err1
}

func CreateAccessToken(username string) (string, error) {
	jwtSecretKey := []byte(jwtConfigs.SecretKey)
	claims := &Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err1 := token.SignedString(jwtSecretKey)
	if err1 != nil {
		return "", err1
	}
	return accessToken, nil
}

func IsAccessTokenValid(accessToken string) bool {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImhhaGEiLCJleHAiOjE2NjY1MTYxNjJ9.aQfi1xbfxZY1SG7MkWi0Q0Wpycd5JBdf3bW7-AtvOkI"

	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

	return true

}
