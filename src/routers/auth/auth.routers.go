package authRouters

import (
	"github.com/gin-gonic/gin"
	authController "github.com/putragabrielll/go-backend/src/controllers/auth"
	"github.com/putragabrielll/go-backend/src/middlewares"
)


func AuthLogin(r *gin.RouterGroup){
	authMiddleware, _ := middlewares.Auth()
	r.POST("/", authMiddleware.LoginHandler)
}

func AuthRegister(r *gin.RouterGroup){
	r.POST("/", authController.Register)
}

func AuthForgotPassword(r *gin.RouterGroup){
	r.POST("/", authController.ForgotPassword)
}