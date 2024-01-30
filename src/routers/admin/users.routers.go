package admin

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)

func usersRouter(r *gin.RouterGroup){
	r.GET("/", adminController.DataUsers)
}