package authController

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)




func Login(c *gin.Context){
	loginauth := services.RLUsers{}
	err := c.ShouldBind(&loginauth) // untuk memasukkan data dari form ke struck Person{}
	if err != nil {
		msg := "Format Email not Support!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	finduser, err := models.FindUsersByEmail(loginauth.Email)
	if err != nil {
		msg := "Email not register!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	checkpas, _ := argonize.DecodeHashStr(finduser.Password) // cek password user
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
	err := c.ShouldBind(&usersData) // untuk memasukkan data dari form ke struck Person{}
	if err != nil {
		msg := "Invalid Email!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}

	paswdhash := []byte(usersData.Password) // proses hashing password
	hasedPasswd, _ := argonize.Hash(paswdhash)

	usersData.Password = hasedPasswd.String()
	usersData.Role = "customer"


	createUser, err := models.RegisterUsers(usersData)
	if err != nil  {
		msg := "Email Already exists!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create users successfully!",
		Results: createUser,
	})
}


func ForgotPassword(c *gin.Context){
	userReset := services.FormReset{}
	c.ShouldBind(&userReset) // untuk memasukkan data dari form ke struck Person{}
	// if err != nil {
	// 	msg := "Format Email not Support!"
	// 	helpers.Utils(err, msg, c) // Error Handle
	// 	return
	// }
	if userReset.Email != "" {
		finduser, err := models.FindUsersByEmail(userReset.Email)
		if err != nil {
			msg := "Email not register!"
			helpers.Utils(err, msg, c) // Error Handler
			return
		}
		if finduser.Id != 0 {
			// https://www.geeksforgeeks.org/how-to-get-random-permutation-of-integers-in-golang/
			getOTP := rand.Perm(9)
			userReset.Otp = strings.Trim(strings.Replace(fmt.Sprint(getOTP[0:6]), " ", "", -1), "[]") // make ot
			models.CreateRP(userReset) 
			// START SEND OTP TO EMAIL
				fmt.Println(userReset) // get otp
			// END SEND EMAIL
			c.JSON(http.StatusOK, &services.ResponseBack{
				Success: true,
				Message: "OTP success send to your email!",
			})
			return
		}
	} 
	if userReset.Otp != "" {
		findEmail, _ := models.FindRPByOTP(userReset.Otp) // pengecekan apakah OTP nya valid
		if findEmail.Id != 0 {
			if (userReset.Password == userReset.ConfirmPassword) {

				paswdhash := []byte(userReset.Password) // proses hashing password
				hasedPasswd, _ := argonize.Hash(paswdhash)

				findUser, _ := models.FindUsersByEmail(findEmail.Email)
				dataUpdate := services.Person{
					Id: findUser.Id,
					Password: hasedPasswd.String(),
				}
				updatedUsers, _ := models.UpdateUsers(dataUpdate)
				message := fmt.Sprintf("Reset password for %v success!", updatedUsers.Email)
				c.JSON(http.StatusOK, &services.ResponseBack{
					Success: true,
					Message: message,
				})
				models.DeleteOTP(findEmail.Id)
				return
			} else {
				// msg := "confirmPassword"
				// var err error
				// helpers.Utils(err, msg, c) // Error Handler
				// return
				c.JSON(http.StatusBadRequest, &services.ResponseBack{
					Success: false,
					Message: "Confirm password does not match!",
				})
				return
			}
		}
	} else {
		// msg := ""
		// var err error
		// helpers.Utils(err, msg, c) // Error Handler
		// return
		c.JSON(http.StatusInternalServerError, &services.ResponseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
}