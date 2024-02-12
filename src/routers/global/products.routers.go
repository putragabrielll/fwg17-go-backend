package globalRouters

import (
	"github.com/gin-gonic/gin"
	globalController "github.com/putragabrielll/go-backend/src/controllers/global"
)


func productsRouter(rg *gin.RouterGroup){
	rg.GET("/", globalController.ListAllProducts)
	rg.GET("/:id", globalController.IdProducts)
}