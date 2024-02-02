package adminUsersController

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	modelsUsers "github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL USERS
func ListAllUsers(c *gin.Context){ // contex => c "inisial aja"
	users, err := modelsUsers.ListAllUsers()
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
	users, err := modelsUsers.FindUsersId(id)
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
	err := c.Bind(&usersData) // menggunakan pointer
	if err != nil {
		msg := "Invalid Email!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	usersData.Role = "customer"


	createUser, _ := modelsUsers.CreateUsers(usersData)
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
	err := c.Bind(&usersData) // menggunakan pointer
	if err != nil {
		msg := "Invalid Email!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	usersData.Id = id // mengarahkan isi dari usersData dengan value "id" di ambil dari id di atas.

	updatedUsers, err := modelsUsers.UpdateUsers(usersData)
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
	users, err := modelsUsers.DeleteUsersId(id)
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