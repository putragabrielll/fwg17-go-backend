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

// SELECT ALL PRODUCT TAGS
func ListAllPT(c *gin.Context){
	filterby := c.DefaultQuery("filterby", "product")
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "product")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.Countpt(filterby, filter)
	page_total := math.Ceil(float64(countData)/float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	productTags, err := models.ListAllpt(filterby, filter, sortby, order, limit, offset)
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
			TotalPage: int(page_total),
			NextPage: page_next,
			PrevPage: page_prev,
			TotalData: countData,
		},
		Results: productTags,
	})
}


// GET PRODUCT TAGS BY id
func IdPT(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptData, err := models.FindPT(id)
	if err != nil {
		msg := "Product tags not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail product tags!",
		Results: ptData,
	})
}


// CREATE PRODUCT TAGS
func CreatePT(c *gin.Context){
	ptData := services.Pro_Tags{}
	err := c.ShouldBind(&ptData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	
	createdpt, _ := models.CreatePT(ptData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create product tags successfully!",
		Results: createdpt,
	})
}


// UPDATE PRODUCT TAGS
func UpdatePT(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptsData := services.Pro_Tags{}
	err := c.ShouldBind(&ptsData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	ptsData.Id = id

	updatedPT, err := models.UpdatePT(ptsData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update product tags successfully!",
		Results: updatedPT,
	})
}


// DELETE PRODUCT TAGS
func DeletePT(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptData, err := models.DeletePT(id)
	if err != nil {
		msg := "Product tags not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete product tags successfully!",
		Results: ptData,
	})
}