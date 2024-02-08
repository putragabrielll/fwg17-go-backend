package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func ordersDetailsRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllOD)
	rg.GET("/:id", adminController.IdOD)
	rg.POST("/", adminController.CreateOD)
	rg.PATCH("/:id", adminController.UpdateOD)
	rg.DELETE("/:id", adminController.DeleteOD)
}