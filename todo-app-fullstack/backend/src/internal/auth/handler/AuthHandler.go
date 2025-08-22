package handler

import (
	"net/http"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/service"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/model/request"
	"github.com/gin-gonic/gin"
)

type handlerAuth struct {
	svc service.AuthService
}

func NewHandlerAuth(s service.AuthService) handlerAuth  {
	return handlerAuth{svc: s}
}

func (h *handlerAuth) Register(c *gin.Context)  {
	req := new(request.Register)

	if err:= c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed Request",
			"error": err.Error(),
		})
		return
	}


	if err:= h.svc.Register(req);err != nil {
		if helper.Duplicate(err) {
			c.JSON(http.StatusConflict, gin.H{
				"message":"email already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Something Went Wrong",
		})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"message":"Success Register",
	})

}