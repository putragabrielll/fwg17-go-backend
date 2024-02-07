package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func PTRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllPT)
	rg.GET("/:id", adminController.IdPT)
	rg.POST("/", adminController.CreatePT)
	rg.PATCH("/:id", adminController.UpdatePT)
	rg.DELETE("/:id", adminController.DeletePT)
}