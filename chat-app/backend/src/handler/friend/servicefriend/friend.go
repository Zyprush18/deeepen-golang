package servicefriend

import (
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/friend/repositoryfriend"
	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/request"
)

type FriendRepo interface {
	CreateFriend(req *request.Friends) error
	UpdateFriend(req *request.Friends) error
	DeleteFriend(from_id,to_id int) error
}

type takeFriendRepo struct {
	repo repositoryfriend.Friend
}

func NewService(r repositoryfriend.Friend) FriendRepo  {
	return &takeFriendRepo{repo: r}
}

func (r *takeFriendRepo) CreateFriend(req *request.Friends) error  {
	req.Status = "pending"
	return r.repo.CreateFriend(req)
}
func (r *takeFriendRepo) UpdateFriend(req *request.Friends) error  {
	req.Status = "accept"
	return r.repo.UpdateFriend(req)
}

func (r *takeFriendRepo) DeleteFriend(from_id,to_id int) error  {
	return r.repo.DeleteFriend(from_id,to_id)
}