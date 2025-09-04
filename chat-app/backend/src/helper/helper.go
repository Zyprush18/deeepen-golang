package helper

import (
	"errors"
	"log"
	"time"

	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ctxkey string

const UserId ctxkey = "userid"
const ToUserId ctxkey = "touserid"

type Messages struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  string `json:"error,omitempty"`
	Token   string `json:"token,omitempty"`
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

	claim := jwt.RegisteredClaims{
		ID:        data.Uuid,
		Subject:   data.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(jwtKey)
}

func ParseToken(jwttoken string) (*jwt.Token, error) {
	jwtKey := []byte("chatrealt1me")

	return jwt.ParseWithClaims(jwttoken, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return jwtKey, nil
	})

}

func CheckDuplicate(err error) bool  {
	var pgErr  = &pgconn.PgError{
		Code: "23505",
	}
	return errors.As(err, &pgErr)
}
