package customerRouters

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/middlewares"
)


func CustomerRouter(r *gin.RouterGroup){
	authMiddleware, _ := middlewares.Auth()
	r.Use(authMiddleware.MiddlewareFunc())

	profileRouter(r.Group("/profile"))
	// profileRouter(r.Group("/profile"))
}