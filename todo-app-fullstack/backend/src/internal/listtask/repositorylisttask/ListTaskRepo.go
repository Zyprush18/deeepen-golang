package repositorylisttask

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/database/entity"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"gorm.io/gorm"
)


type ListTaskRepo interface {
	Create(req *request.ListTask) error
	Update(id int, req *request.ListTask) error
	Delete(id int) error
}

type database struct {
	db *gorm.DB
}


func NewConnection(d *gorm.DB) database {
	return  database{db: d}
}

func (d *database) Create(req *request.ListTask) error {
	if err:= d.db.Table("list_tasks").Create(req).Error;err != nil {
		return err
	}

	return nil
}

func (d *database) Update(id int, req *request.ListTask) error  {
	if err:= d.db.Table("list_tasks").Where("id=?",id).Updates(req).Error;err != nil {
		return  err
	}

	return nil
}

func (d *database) Delete(id int) error {
	if err:= d.db.Model(&entity.ListTask{}).Where("id=?", id).Delete(&entity.ListTask{}).Error;err != nil {
		return err
	}

	return nil
}