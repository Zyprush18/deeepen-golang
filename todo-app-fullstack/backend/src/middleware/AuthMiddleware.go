package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Zyprush18/deeepen-golang/todo-app-fullstack/backend/src/helper"
	"github.com/gin-gonic/gin"
)

type middleware struct {
	jwtkey string
}
func NewMiddleware(jkey string) middleware  {
	return middleware{jwtkey: jkey}
}

func (m *middleware) AuthMiddleware() gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		getToken := ctx.Request.Header.Get("Authorization")
		parsetoken := strings.Split(getToken, " ") 
		if getToken == "" || parsetoken[1] == "" || parsetoken[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message":"Unauthorized",
			})
			ctx.Abort()
			return 
		}

		// cek token
		token, err := helper.ParseTokenJwt(parsetoken[1],m.jwtkey)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			ctx.Abort()
			return 
		}

		iduser,err := strconv.Atoi(token.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			ctx.Abort()
			return 
		}
		
		ctx.Set("user_id", iduser)

		ctx.Next()
	}
}