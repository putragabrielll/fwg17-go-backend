package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)

func usersRouter(r *gin.RouterGroup){
	r.GET("/", adminController.ListAllUsers)
	r.GET("/:id", adminController.IdUsers)
	r.POST("/", adminController.CreateUsers)
	r.PATCH("/:id", adminController.UpdateUsers)
	r.DELETE("/:id", adminController.DeleteUsers)
}