package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/Zyprush18/deeepen-golang/chat-app/src/model/response"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ctxkey string

const Name ctxkey = "name"

type Messages struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  string `json:"error,omitempty"`
	Token   string `json:"token,omitempty"`
}

type CustomClaim struct {
	Name string
	jwt.RegisteredClaims
}

func HashPass(pass string) string {
	hashpw, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		log.Fatal("Failed Hashin Password: ", err)
	}

	return string(hashpw)
}

func CheckHashPw(pass, reqpass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(reqpass)); err != nil {
		return false
	}

	return true
}

func GenerateToken(data *response.Auth) (string, error) {
	jwtKey := []byte("chatrealt1me")

	claim := CustomClaim{
		Name: data.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprintf("%d", data.ID),
			Subject:   data.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(jwtKey)
}

func ParseToken(jwttoken string) (*jwt.Token, error) {
	jwtKey := []byte("chatrealt1me")

	return jwt.ParseWithClaims(jwttoken, &CustomClaim{}, func(t *jwt.Token) (any, error) {
		return jwtKey, nil
	})

}
