package routers

import (
	"github.com/gin-gonic/gin"
	adminRouters "github.com/putragabrielll/go-backend/src/routers/admin"
	authRouters "github.com/putragabrielll/go-backend/src/routers/auth"
)

func Combine(r *gin.Engine){
	//------------ AUTH ------------
	authRouters.AuthLogin(r.Group("/login"))
	authRouters.AuthRegister(r.Group("/register"))
	authRouters.AuthForgotPassword(r.Group("/forgot-password"))

	adminRouters.AdminRouter(r.Group("/admin"))
	// admin.AdminRouter(r.Group("/customer"))
}