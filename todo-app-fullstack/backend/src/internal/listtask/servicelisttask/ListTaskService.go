package servicelisttask

import (
	"time"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/listtask/repositorylisttask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
)

type ListTaskService interface {
	Create(req *request.ListTask) error
	Update(id int,req *request.ListTask) error
	Delete(id int) error
}

type takeTaskRepo struct {
	repo repositorylisttask.ListTaskRepo
}

func NewService(r repositorylisttask.ListTaskRepo) ListTaskService  {
	return &takeTaskRepo{repo: r}
}
func (r *takeTaskRepo) Create(req *request.ListTask) error  {
	now:= time.Now()
	req.CreatedAt = now
	return  r.repo.Create(req)
}

func (r *takeTaskRepo) Update(id int,req *request.ListTask) error  {
	now := time.Now()
	if req != nil {
		req.UpdatedAt = now
	}

	return r.repo.Update(id, req)
}

func (r *takeTaskRepo) Delete(id int) error {
	return r.repo.Delete(id)
}