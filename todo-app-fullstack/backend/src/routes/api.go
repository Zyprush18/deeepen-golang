package routes

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/config"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/handler"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/repository"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/auth/service"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/listtask/handlerlisttask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/listtask/repositorylisttask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/listtask/servicelisttask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/handlertask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/repositorytask"
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/internal/task/servicetask"
	"github.com/gin-gonic/gin"
)

func NewRoute(r *gin.Engine)  {
	initdb:= config.Config()

	api := r.Group("/api/v1")

	// auth
	authrepo := repository.NewConnection(initdb)
	authservice := service.NewServiceAuth(&authrepo)
	authhandler := handler.NewHandlerAuth(authservice)

	api.POST("/register", authhandler.Register)

	
	// task
	taskrepo := repositorytask.NewConnection(initdb)
	taskservice := servicetask.NewService(&taskrepo)
	taskhandler := handlertask.NewHandler(taskservice)

	api.GET("/task", taskhandler.Index)
	api.POST("/task/add", taskhandler.Store)
	api.GET("/task/:id", taskhandler.Show)
	api.PUT("/task/:id/update", taskhandler.Update)
	api.DELETE("/task/:id/delete", taskhandler.Delete)

	// list task
	listtaskrepo := repositorylisttask.NewConnection(initdb)
	listtaskservice := servicelisttask.NewService(&listtaskrepo)
	listtaskhandler := handlerlisttask.NewHandler(listtaskservice)

	api.POST("/list/add", listtaskhandler.Store)
	api.PUT("/list/:id/update", listtaskhandler.Update)
	api.DELETE("/list/:id/delete", listtaskhandler.Delete)

	r.Run()
}