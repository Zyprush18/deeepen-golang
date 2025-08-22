package handlerlisttask

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/listtask/servicelisttask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handlerListTask struct {
	svc servicelisttask.ListTaskService
}

func NewHandler(s servicelisttask.ListTaskService) handlerListTask {
	return handlerListTask{svc: s}
}

func (h *handlerListTask) Store(c *gin.Context) {
	req := new(request.ListTask)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Create a new list task: body request is missing or failed validation",
			"errors":  err.Error(),
		})

		return
	}

	if err := h.svc.Create(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Create a new list task",
	})

}

func (h *handlerListTask) Update(c *gin.Context) {
	getid := c.Param("id")
	id, err := strconv.Atoi(getid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type param",
		})
		return
	}

	req := new(request.ListTask)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Create a new list task: body request is missing or failed validation",
			"errors":  err.Error(),
		})
		return
	}

	if err := h.svc.Update(id, req); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found Data",
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Update data",
	})
}

func (h *handlerListTask) Delete(c *gin.Context) {
	getid := c.Param("id")
	id, err := strconv.Atoi(getid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type param",
		})
		return
	}

	if err:= h.svc.Delete(id);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not Found Data",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"success Delete data",
	})
}
