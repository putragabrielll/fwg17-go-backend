package adminController

import (
	"math"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)


// SELECT ALL PRODUCT CATEGORIES
func ListAllPC(c *gin.Context){
	filterby := c.DefaultQuery("filterby", "product")
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "product")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.Countpc(filterby, filter)
	page_total := math.Ceil(float64(countData) / float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	productCate, err := models.ListAllpc(filterby, filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Products tags not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all products tags!",
		PageInfo: services.PageInfo{
			CurrentPage: page,
			TotalPage:   int(page_total),
			NextPage:    page_next,
			PrevPage:    page_prev,
			TotalData:   countData,
		},
		Results: productCate,
	})
}


// GET PRODUCT CATEGORIES BY id
func IdPC(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptData, err := models.FindPT(id)
	if err != nil {
		msg := "Product categories not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail product categories!",
		Results: ptData,
	})
}


// CREATE PRODUCT CATEGORIES
func CreatePC(c *gin.Context){
	pcData := services.Pro_Cate{}
	err := c.ShouldBind(&pcData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}

	createdpc, err := models.CreatePC(pcData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create product tags successfully!",
		Results: createdpc,
	})
}


// UPDATE PRODUCT CATEGORIES
func UpdatePC(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	pcData := services.Pro_Cate{}
	err := c.ShouldBind(&pcData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	pcData.Id = id

	updatedPC, err := models.UpdatePC(pcData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update product tags successfully!",
		Results: updatedPC,
	})
}


// DELETE PRODUCT CATEGORIES
func DeletePC(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	pcData, err := models.DeletePC(id)
	if err != nil {
		msg := "Product categories not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete product categories successfully!",
		Results: pcData,
	})
}