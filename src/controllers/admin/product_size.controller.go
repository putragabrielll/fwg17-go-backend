package adminController

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL product size
func ListAllPS(c *gin.Context){
	ListPS, err := models.ListAllps()
	if err != nil {
		msg := "Product size not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all products!",
		Results: ListPS,
	})
}


// GET product size BY id
func IdPS(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	GetPS, err := models.FindPSId(id)
	if err != nil {
		msg := "Product size not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail Product size!",
		Results: GetPS,
	})
}


// UPDATE PRODUCT SIZE
func UpdatePS(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	dataPS := services.Ps{}
	err := c.ShouldBind(&dataPS) // untuk memasukkan data dari form ke struck Person{}
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	dataPS.Id = id
	updatedUsers, err := models.UpdatePS(dataPS)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update products size successfully!",
		Results: updatedUsers,
	})
}