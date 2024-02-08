package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func ordersRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllOrders)
	rg.GET("/:id", adminController.IdOrders)
	rg.POST("/", adminController.CreateOrders)
	rg.PATCH("/:id", adminController.UpdateOrders)
	rg.DELETE("/:id", adminController.DeleteOrders)
}