package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func productsRouter(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllProducts)
	rg.GET("/:id", adminController.IdProducts)
	rg.POST("/", adminController.CreateProducts)
	rg.PATCH("/:id", adminController.UpdateProducts)
	rg.DELETE("/:id", adminController.DeleteProducts)
}