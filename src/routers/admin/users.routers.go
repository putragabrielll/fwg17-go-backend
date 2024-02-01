package admin

import (
	"github.com/gin-gonic/gin"
	adminUsersController "github.com/putragabrielll/go-backend/src/controllers/admin"
)

func usersRouter(r *gin.RouterGroup){
	r.GET("/", adminUsersController.ListAllUsers)
	r.GET("/:id", adminUsersController.IdUsers)
	r.POST("/", adminUsersController.CreateUsers)
	// r.PATCH("/:id", adminUsersController.CreateUsers)
	// r.DELETE("/:id", adminUsersController.CreateUsers)
}