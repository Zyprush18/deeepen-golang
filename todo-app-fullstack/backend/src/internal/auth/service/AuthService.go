package service

import (
	"time"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/repository"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
)

type AuthService interface {
	Register(req *request.Register) error
}

type takeAuthRepo struct {
	repo repository.AuthRepo
}

func NewServiceAuth(r repository.AuthRepo) AuthService  {
	return &takeAuthRepo{repo: r}
}

func (r *takeAuthRepo) Register(req *request.Register) error  {
	now := time.Now()
	requser := &request.User{
		Username: req.Username,
		Email: req.Email,
		Password: helper.HashingPass(req.Password),
		Model: helper.Model{
			CreatedAt: now,
		},
	}

	return r.repo.Register(requser)
}