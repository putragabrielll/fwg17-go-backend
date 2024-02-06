package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func psRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllPS)
	rg.GET("/:id", adminController.IdPS)
	rg.PATCH("/:id", adminController.UpdatePS)
}