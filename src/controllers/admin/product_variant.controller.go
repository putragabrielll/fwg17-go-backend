package adminController

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL PRODUCT VARIANT
func ListAllVariant(c *gin.Context){
	ListPS, err := models.ListAllpv()
	if err != nil {
		msg := "Product variant not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all product variant!",
		Results: ListPS,
	})
}


// GET product variant BY id
func IdPV(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	GetPV, err := models.FindPVId(id)
	if err != nil {
		msg := "Product variant not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail Product variant!",
		Results: GetPV,
	})
}


// CREATE PRODUCT VARIANT
func CreatePV(c *gin.Context){
	pvData := services.Pv{}
	err := c.ShouldBind(&pvData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	
	createdPV, _ := models.CreatePv(pvData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create products successfully!",
		Results: createdPV,
	})
}


// UPDATE PRODUCT VARIAN
func UpdatePV(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	pvData := services.Pv{}
	err := c.ShouldBind(&pvData) // untuk memasukkan data dari form ke struck Person{}
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	pvData.Id = id

	updatePV, err := models.UpdatePv(pvData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update products variant successfully!",
		Results: updatePV,
	})
}


// DELETE PRODUCT VARIAT
func DeletePV(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	deletePV, err := models.DeletePv(id)
	if err != nil {
		msg := "Products variant not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete products variant successfully!",
		Results: deletePV,
	})
}