package adminController

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL PROMO
func ListAllPromo(c *gin.Context) {
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "id")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountAllPromo(filter)
	page_total := math.Ceil(float64(countData) / float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	products, err := models.ListAllPromo(filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Promo not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all products!",
		PageInfo: services.PageInfo{
			CurrentPage: page,
			TotalPage:   int(page_total),
			NextPage:    page_next,
			PrevPage:    page_prev,
			TotalData:   countData,
		},
		Results: products,
	})
}


// GET PROMO BY id
func IdPromo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.FindPromoId(id)
	if err != nil {
		msg := "Promo not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail promo!",
		Results: users,
	})
}


// CREATE PROMO
func CreatePromo(c *gin.Context){
	promoData := services.Promo{}
	err := c.ShouldBind(&promoData)
	fmt.Println(err)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	
	createdPromo, _ := models.CreatePromo(promoData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create products successfully!",
		Results: createdPromo,
	})
}


// UPDATE PROMO
func UpdatePromo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	promoData := services.Promo{}
	err := c.ShouldBind(&promoData) // untuk memasukkan data dari form ke struck Person{}
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	promoData.Id = id

	updatedPromo, err := models.UpdatePromo(promoData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update products successfully!",
		Results: updatedPromo,
	})
}


// DELETE PROMO
func DeletePromo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.DeletePromo(id)
	if err != nil {
		msg := "Promo not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete promo successfully!",
		Results: users,
	})
}
