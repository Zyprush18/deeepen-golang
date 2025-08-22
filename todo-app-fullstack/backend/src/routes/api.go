package routes

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/config"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/handler"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/repository"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/service"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/handlertask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/repositorytask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/servicetask"
	"github.com/gin-gonic/gin"
)

func NewRoute(r *gin.Engine)  {
	initdb:= config.Config()

	api := r.Group("/api/v1")

	authrepo := repository.NewConnection(initdb)
	authservice := service.NewServiceAuth(&authrepo)
	authhandler := handler.NewHandlerAuth(authservice)



	api.POST("/register", authhandler.Register)
	
	taskrepo := repositorytask.NewConnection(initdb)
	taskservice := servicetask.NewService(&taskrepo)
	taskhandler := handlertask.NewHandler(taskservice)

	api.GET("/task", taskhandler.Index)
	api.POST("/task/add", taskhandler.Store)
	api.GET("/task/:id", taskhandler.Show)
	api.PUT("/task/:id/update", taskhandler.Update)
	api.DELETE("/task/:id/delete", taskhandler.Delete)


	r.Run()
}