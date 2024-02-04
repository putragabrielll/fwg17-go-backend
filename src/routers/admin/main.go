package adminRouters

import "github.com/gin-gonic/gin"

func AdminRouter(r *gin.RouterGroup){
	usersRouter(r.Group("/users"))
}