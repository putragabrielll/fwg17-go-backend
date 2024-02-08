package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func PRRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllPR)
	rg.GET("/:id", adminController.IdPR)
	rg.POST("/", adminController.CreatePR)
	rg.PATCH("/:id", adminController.UpdatePR)
	rg.DELETE("/:id", adminController.DeletePR)
}