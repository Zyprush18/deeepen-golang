package servicetask

import (
	"fmt"
	"time"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/repositorytask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/response"
)

type TaskSevice interface {
	GatAll(id int) ([]response.Task, error)
	Create(req *request.Task, id int) error
	Show(id, userid int) (*response.Task, error)
	Update(id, userid int, req *request.Task) error
	Delete(id, userid int) error
}

type takeTaskRepo struct {
	repo repositorytask.TaskRepo
}

func NewService(r repositorytask.TaskRepo) TaskSevice {
	return &takeTaskRepo{repo: r}
}

func (r *takeTaskRepo) GatAll(id int) ([]response.Task, error) {
	return r.repo.GetAll(id)
}

func (r *takeTaskRepo) Create(req *request.Task, id int) error {
	now := time.Now()
	req.UserId = uint(id)
	req.Model = request.Model{
		CreatedAt: now,
	}

	return r.repo.Store(req)
}

func (r *takeTaskRepo) Show(id, userid int) (*response.Task, error) {
	return r.repo.Show(id, userid)
}

func (r *takeTaskRepo) Update(id, userid int, req *request.Task) error {
	now := time.Now()
	fmt.Println(req)
	if req != nil {
		req.Model = request.Model{
			UpdatedAt: now,
		}
	}

	return r.repo.Update(id, userid, req)
}

func (r *takeTaskRepo) Delete(id, userid int) error  {
	return r.repo.Delete(id,userid)
}