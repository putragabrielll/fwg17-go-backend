package globalRouters

import (
	"github.com/gin-gonic/gin"
	globalController "github.com/putragabrielll/go-backend/src/controllers/global"
)


func sizeRouters(rg *gin.RouterGroup){
	rg.GET("/", globalController.ListAllPS)
}