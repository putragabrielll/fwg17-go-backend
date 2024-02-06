package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func pvRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllVariant)
	rg.GET("/:id", adminController.IdPV)
	rg.POST("/", adminController.CreatePV)
	rg.PATCH("/:id", adminController.UpdatePV)
	rg.DELETE("/:id", adminController.DeletePV)
}