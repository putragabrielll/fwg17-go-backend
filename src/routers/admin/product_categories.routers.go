package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func PCRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllPC)
	rg.GET("/:id", adminController.IdPC)
	rg.POST("/", adminController.CreatePC)
	rg.PATCH("/:id", adminController.UpdatePC)
	rg.DELETE("/:id", adminController.DeletePC)
}