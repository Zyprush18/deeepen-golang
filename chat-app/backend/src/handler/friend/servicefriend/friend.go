package servicefriend

// import (
// 	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/handler/friend/repositoryfriend"
// 	"github.com/Zyprush18/deeepen-golang/chat-app/backend/src/model/response"
// )

// type FriendRepo interface {
// 	GetAll(id int)(*response.UserFriend, error)
// }

// type takeFriendRepo struct {
// 	repo repositoryfriend.Friend
// }

// func NewService(r repositoryfriend.Friend) FriendRepo  {
// 	return &takeFriendRepo{repo: r}
// }

// func (r *takeFriendRepo) GetAll(id int)(*response.UserFriend, error)  {
// 	return r.repo.GetFriend(id)
// }