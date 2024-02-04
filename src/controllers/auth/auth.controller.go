package authController

import (
	// "log"
	"net/http"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	modelsUsers "github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)




func Login(c *gin.Context){
	loginauth := services.RLUsers{}
	err := c.ShouldBind(&loginauth)
	if err != nil {
		msg := "Invalid Email!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	finduser, err := modelsUsers.FindUsersByEmail(loginauth.Email)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	checkpas, _ := argonize.DecodeHashStr(finduser.Password)
	paswdcheck := []byte(loginauth.Password)
	if !checkpas.IsValidPassword(paswdcheck) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Wrong Email or Password!",
		})
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Login success!",
		Results: finduser,
	})
}


func Register(c *gin.Context){
	usersData := services.RLUsers{} // menggunakan tipe data yg ada di model users.
	c.ShouldBind(&usersData) // menggunakan pointer

	paswdhash := []byte(usersData.Password) // proses hashing password
	hasedPasswd, _ := argonize.Hash(paswdhash)

	usersData.Password = hasedPasswd.String()
	usersData.Role = "customer"


	createUser, _ := modelsUsers.RegisterUsers(usersData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create users successfully!",
		Results: createUser,
	})
}


func ForgotPassword(c *gin.Context){
	
}