package adminRouters

import (
	"github.com/gin-gonic/gin"
	adminController "github.com/putragabrielll/go-backend/src/controllers/admin"
)


func promoRouters(rg *gin.RouterGroup){
	rg.GET("/", adminController.ListAllPromo)
	rg.GET("/:id", adminController.IdPromo)
	rg.POST("/", adminController.CreatePromo)
	rg.PATCH("/:id", adminController.UpdatePromo)
	rg.DELETE("/:id", adminController.DeletePromo)
}