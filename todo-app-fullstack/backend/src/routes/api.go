package routes

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/config"
	"github.com/gin-gonic/gin"
)

func NewRoute(r *gin.Engine)  {
	config.Config()

	r.Run()
}