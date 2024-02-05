package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func productsRouter(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllProducts)
	// rg.GET("/:id", adminController.IdUsers)
	// rg.POST("/", adminController.CreateUsers)
	// rg.PATCH("/:id", adminController.UpdateUsers)
	// rg.DELETE("/:id", adminController.DeleteUsers)
}