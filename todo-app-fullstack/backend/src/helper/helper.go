package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func HashingPass(pass string) string  {
	hashpass,err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return string(hashpass)
}

func Duplicate(errs error) bool {
	err := pgconn.PgError{
		Code: "23505",
	}

	return errs == &err 
}


type customClaims struct {
	userid int
	jwt.RegisteredClaims
}

func GenerateToken(user response.User,jwtkey string) string {
	claim := &customClaims{
		userid: int(user.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			ID: fmt.Sprintf("%d",user.Id),
			Subject: user.Email,
			Issuer: "backend",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}


	generate := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	token,err := generate.SignedString([]byte(jwtkey))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return token
}