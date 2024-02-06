package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func categoriesRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllCategories)
	rg.GET("/:id", adminController.IdCategories)
	rg.POST("/", adminController.CreateCTGR)
	rg.PATCH("/:id", adminController.UpdateCTGR)
	rg.DELETE("/:id", adminController.DeleteCTGR)
}