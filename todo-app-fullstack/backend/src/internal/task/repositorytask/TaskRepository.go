package repositorytask

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/entity"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/response"
	"gorm.io/gorm"
)

type TaskRepo interface {
	GetAll(id int)([]response.Task, error)
	Store(req *request.Task) error
	Show(id,userid int) (*response.Task, error)
	Update(id,userid int, req *request.Task) error
	Delete(id,userid int) error
}

type database struct {
	db  *gorm.DB
}

func NewConnection(d *gorm.DB) database  {
	return database{db: d}
}

func (d *database) GetAll(id int)([]response.Task, error) {
	var modeltask []entity.Task
	if err:= d.db.Model(&entity.Task{}).Where("user_id = ?",id).Find(&modeltask).Error;err != nil {
		return nil, err
	}

	return response.ParseAllTask(modeltask),nil
}

func (d *database) Store(req *request.Task) error {
	if err:=d.db.Table("tasks").Create(req).Error;err != nil {
		return err
	}

	return nil
}

func (d *database) Show(id,userid int) (*response.Task, error) {
	var modeltask entity.Task
	if err := d.db.Model(&entity.Task{}).Preload("ListTasks").Where("id = ? AND user_id = ?", id, userid).First(&modeltask).Error;err != nil {
		return nil, err
	}

	return &response.Task{
		Id: modeltask.Id,
		Title: modeltask.Title,
		SubTitle: modeltask.SubTitle,
		Description: modeltask.Description,
		UserId: modeltask.UserId,
		ListTasks: response.ParseAllLIstTask(modeltask.ListTasks),
		Model: response.Model{
			CreatedAt: modeltask.CreatedAt,
			UpdatedAt: modeltask.UpdatedAt,
		},
	},nil
}


func (d *database) Update(id,userid int, req *request.Task) error  {
	if err:= d.db.Table("tasks").Where("id = ? and user_id = ?", id,userid).Updates(&req).Error;err != nil {
		return  err
	}

	return nil
}

func (d *database) Delete(id,userid int) error  {
	if err:= d.db.Model(&entity.Task{}).Where("id = ?  AND user_id = ?",id,userid).Delete(&entity.Task{}).Error;err != nil {
		return err
	}

	return  nil
}