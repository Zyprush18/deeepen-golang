package handlertask

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/servicetask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HandlerTask struct {
	svc servicetask.TaskSevice
}

func NewHandler(s servicetask.TaskSevice) HandlerTask  {
	return HandlerTask{svc: s}
}

func (h *HandlerTask) Index(c *gin.Context)  {
	userid := c.MustGet("user_id").(int)
	resp,err := h.svc.GatAll(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"success",
		"data":resp,
	})
}

func (h *HandlerTask) Store(c *gin.Context)  {
	userid := c.MustGet("user_id").(int)

	req := new(request.Task)

	if err:=c.ShouldBindJSON(req);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Failed Create a new task: body request is missing or failed validation",
			"error":err.Error(),
		})
		return
	}

	if err:= h.svc.Create(req, userid);err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":"Success Created",
	})
}

func (h *HandlerTask) Show(c *gin.Context)  {
	userid := c.MustGet("user_id").(int)

	getid := c.Param("id")
	id, err := strconv.Atoi(getid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type param",
		})
		return
	}

	resp, err := h.svc.Show(id, userid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message":"Not Found Data",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":"success",
		"data":resp,
	})
}

func (h *HandlerTask) Update(c *gin.Context)  {
	userid := c.MustGet("user_id").(int)

	getid := c.Param("id")
	id, err := strconv.Atoi(getid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type param",
		})
		return
	}

	req := new(request.Task)
	if err:= c.ShouldBindJSON(req);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Body Request Is Missing",
			"error":err.Error(),
		})
		return
	}

	if err:= h.svc.Update(id, userid, req) ;err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message":"Not Found Data",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"success update",
	})
}

func (h *HandlerTask) Delete(c *gin.Context)  {
	userid := c.MustGet("user_id").(int)

	getid := c.Param("id")
	id, err := strconv.Atoi(getid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid type param",
		})
		return
	}

	if err:= h.svc.Delete(id,userid);err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message":"Not Found Data",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":"Success Delete",
	})
}