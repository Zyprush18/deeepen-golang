package helper

import (
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type Model struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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