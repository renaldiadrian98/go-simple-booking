package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userid int, email string, roleId int) (string, error) {
	var err error

	// Creating Access Token
	secretKey := os.Getenv("SECRETKEY")
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["email"] = email
	atClaims["role_id"] = roleId
	atClaims["expire"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
