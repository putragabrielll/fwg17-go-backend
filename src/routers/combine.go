package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/routers/admin"
)

func Combine(r *gin.Engine){
	admin.AdminRouter(r.Group("/admin"))
	// admin.AdminRouter(r.Group("/customer"))
}