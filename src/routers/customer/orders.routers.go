package customerRouters

import (
	"github.com/gin-gonic/gin"
	customerController "github.com/putragabrielll/go-backend/src/controllers/customer"
)


func ordersRouters(rg *gin.RouterGroup){
	rg.GET("/", customerController.IdOrders)
	// rg.GET("/:id", adminController.IdOrders)
	rg.POST("/", customerController.CreateOrders)
}