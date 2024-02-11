package adminController

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"github.com/KEINOS/go-argonize"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/middlewares"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL USERS
func ListAllUsers(c *gin.Context) { // contex => c "inisial aja"
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "id")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountAllUsers(filter) // COUNT DATA USER
	// https://dev.to/natamacm/round-numbers-in-go-5c01
	page_total := math.Ceil(float64(countData) / float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	users, err := models.ListAllUsers(filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all users!",
		PageInfo: services.PageInfo{
			CurrentPage: page,
			TotalPage:   int(page_total),
			NextPage:    page_next,
			PrevPage:    page_prev,
			TotalData:   countData,
		},
		Results: users,
	})
}

// GET USERS BY id
func IdUsers(c *gin.Context) {
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
func CreateUsers(c *gin.Context) {
	usersData := services.Person{}  // menggunakan tipe data yg ada di model users.
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

	createUser, err := models.CreateUsers(usersData)
	if err != nil {
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
func UpdateUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	usersData := services.Person{}
	err := c.ShouldBind(&usersData) // untuk memasukkan data dari form ke struck Person{}
	// if err != nil {
	// 	msg := "Invalid Email!"
	// 	helpers.Utils(err, msg, c) // Error Handle
	// 	return
	// }
	if usersData.Password != "" {
		paswdhash := []byte(usersData.Password) // proses hashing password
		hasedPasswd, _ := argonize.Hash(paswdhash)
		usersData.Password = hasedPasswd.String()
	}

	// -------------
	cekFile, err := middlewares.UploadFile(c, "profile") // fungsi upload file
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	usersData.Picture = cekFile
	usersData.Id = id // mengarahkan isi dari usersData dengan value "id" di ambil dari id di atas.
	// -------------

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
func DeleteUsers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.DeleteUsersId(id)
	if err != nil {
		msg := "Users not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete users successfully!",
		Results: users,
	})
}
