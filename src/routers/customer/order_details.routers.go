package customerRouters

import (
	"github.com/gin-gonic/gin"
	customerController "github.com/putragabrielll/go-backend/src/controllers/customer"
)


func ordersDetailsRouters(rg *gin.RouterGroup){
	rg.GET("/", customerController.ListAllOD)
	// rg.GET("/:id", customerController.IdOD)
	rg.POST("/", customerController.CreateOD)
	rg.PATCH("/:id", customerController.UpdateOD)
	// rg.DELETE("/:id", customerController.DeleteOD)
}