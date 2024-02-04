package adminController

import (
	"net/http"
	"strconv"

	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL USERS
func ListAllUsers(c *gin.Context){ // contex => c "inisial aja"
	users, err := models.ListAllUsers()
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all users!",
		Results: users,
	})
}


// GET USERS BY id
func IdUsers(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.FindUsersId(id)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail users!",
		Results: users,
	})
}


// CREATE USERS
func CreateUsers(c *gin.Context){
	usersData := services.Person{} // menggunakan tipe data yg ada di model users.
	c.ShouldBind(&usersData) // menggunakan pointer
	// if err != nil {
	// 	msg := "Invalid Email!"
	// 	helpers.Utils(err, msg, c) // Error Handle
	// 	return
	// }
	paswdhash := []byte(usersData.Password) // proses hashing password
	hasedPasswd, _ := argonize.Hash(paswdhash)

	usersData.Password = hasedPasswd.String()
	usersData.Role = "customer"


	createUser, err := models.CreateUsers(usersData)
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


// UPDATE USERS
func UpdateUsers(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	usersData := services.Person{} // menggunakan tipe data yg ada di model users.
	c.ShouldBind(&usersData) // menggunakan pointer
	// if err != nil {
	// 	msg := "Invalid Email!"
	// 	helpers.Utils(err, msg, c) // Error Handle
	// 	return
	// }
	paswdhash := []byte(usersData.Password) // proses hashing password
	hasedPasswd, _ := argonize.Hash(paswdhash)
	
	usersData.Password = hasedPasswd.String()
	usersData.Id = id // mengarahkan isi dari usersData dengan value "id" di ambil dari id di atas.

	updatedUsers, err := models.UpdateUsers(usersData)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Users updated successfully!",
		Results: updatedUsers,
	})
}


// DELETE USERS
func DeleteUsers(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.DeleteUsersId(id)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete users successfully!",
		Results: users,
	})
}