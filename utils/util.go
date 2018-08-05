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
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	signed, err := token.SignedString([]byte(secret))
	utils.CheckErr(err)
	return signed
}

// TODO; figure out a way to use generic type
func List2Map(list []string) map[string]bool {
	result := map[string]bool{}
	for _, entry := range list {
		result[entry] = true
	}
	return result
}
