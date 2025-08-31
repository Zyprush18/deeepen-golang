package servicesauth

import (
	"errors"

	"github.com/Zyprush18/deeepen-golang/chat-app/src/handler/auth/repositoryauth"
	"github.com/Zyprush18/deeepen-golang/chat-app/src/helper"
	"github.com/Zyprush18/deeepen-golang/chat-app/src/model/request"
)

type AuthService interface {
	Register(req *request.Register) error
	Login(req *request.Login) (string, error)
}

type takeAuthRepo struct {
	repo repositoryauth.AuthRepo
}

func NewService(r repositoryauth.AuthRepo) AuthService {
	return &takeAuthRepo{repo: r}
}

func (r *takeAuthRepo) Register(req *request.Register) error {
	req.Password = helper.HashPass(req.Password)
	return r.repo.Register(req)
}

func (r *takeAuthRepo) Login(req *request.Login) (string, error) {
	data, err := r.repo.Login(req.Email)
	if err != nil {
		return "", err
	}

	if !helper.CheckHashPw(data.Password, req.Password) {
		return "", errors.New("pass incorrect")
	}

	return helper.GenerateToken(data)
}
