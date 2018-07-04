package utils

import (
	"code.jogchat.internal/dgrijalva-jwt-go"
	"github.com/satori/go.uuid"
	"time"
	"code.jogchat.internal/go-schemaless/utils"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func NewUUID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

func GetToken(secret string, email string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = email
	claims["iss"] = "jogchat.com"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	signed, err := token.SignedString([]byte(secret))
	utils.CheckErr(err)
	return signed
}
