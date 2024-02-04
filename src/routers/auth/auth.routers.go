package authRouters

import (
	"github.com/gin-gonic/gin"
	authController "github.com/putragabrielll/go-backend/src/controllers/auth"
)


func AuthLogin(r *gin.RouterGroup){
	r.POST("/", authController.Login)
}

func AuthRegister(r *gin.RouterGroup){
	r.POST("/", authController.Register)
}

func AuthForgotPassword(r *gin.RouterGroup){
	r.POST("/", authController.ForgotPassword)
}