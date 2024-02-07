package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func tagRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllTags)
	rg.GET("/:id", adminController.IdTags)
	rg.POST("/", adminController.CreateTags)
	rg.PATCH("/:id", adminController.UpdateTags)
	rg.DELETE("/:id", adminController.DeleteTags)
}