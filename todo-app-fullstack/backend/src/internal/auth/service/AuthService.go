package service

import (
	"time"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/repository"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
)

type AuthService interface {
	Register(req *request.Register) error
	Login(req *request.Login) (string, error)
}

type takeAuthRepo struct {
	repo repository.AuthRepo
	jwtKey string
}

func NewServiceAuth(r repository.AuthRepo, jkwy string) AuthService  {
	return &takeAuthRepo{repo: r,jwtKey: jkwy}
}

func (r *takeAuthRepo) Register(req *request.Register) error  {
	now := time.Now()
	requser := &request.User{
		Username: req.Username,
		Email: req.Email,
		Password: helper.HashingPass(req.Password),
		Model: request.Model{
			CreatedAt: now,
		},
	}

	return r.repo.Register(requser)
}

func (r *takeAuthRepo) Login(req *request.Login) (string, error)  {
	resp, err:= r.repo.Login(req.Email)
	if err != nil {
		return "", err
	}

	token := helper.GenerateToken(resp, r.jwtKey)

	return token,nil
}