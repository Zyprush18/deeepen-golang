package main

import (
	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.NewRoute(r)
}