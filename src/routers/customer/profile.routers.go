package customerRouters

import (
	"github.com/gin-gonic/gin"
	customerController "github.com/putragabrielll/go-backend/src/controllers/customer"
)

func profileRouter(rg *gin.RouterGroup) {
	rg.GET("/", customerController.ProfileUser)
	rg.PATCH("/", customerController.UpdateUser)
}
