package routers

import (
	"github.com/gin-gonic/gin"
	adminRouters "github.com/putragabrielll/go-backend/src/routers/admin"
	authRouters "github.com/putragabrielll/go-backend/src/routers/auth"
	customerRouters "github.com/putragabrielll/go-backend/src/routers/customer"
	globalRouters "github.com/putragabrielll/go-backend/src/routers/global"
)

func Combine(r *gin.Engine){
	//------------ AUTH ------------
	authRouters.AuthLogin(r.Group("/login"))
	authRouters.AuthRegister(r.Group("/register"))
	authRouters.AuthForgotPassword(r.Group("/forgot-password"))

	adminRouters.AdminRouter(r.Group("/admin"))
	customerRouters.CustomerRouter(r.Group("/customer"))

	//------------ GLOBAL ------------
	globalRouters.GlobalRouter(r.Group(""))
}