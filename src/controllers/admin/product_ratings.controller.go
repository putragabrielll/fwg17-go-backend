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

// SELECT ALL PRODUCT RATINGS
func ListAllPR(c *gin.Context){
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "product")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.Countpr(filter)
	page_total := math.Ceil(float64(countData)/float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	productTags, err := models.ListAllpr(filter, sortby, order, limit, offset)
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


// GET PRODUCT RATINGS BY id
func IdPR(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptData, err := models.Findpr(id)
	if err != nil {
		msg := "Product ratings not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail product ratings!",
		Results: ptData,
	})
}


// CREATE PRODUCT RATINGS
func CreatePR(c *gin.Context){
	prData := services.Pro_Rate{}
	err := c.ShouldBind(&prData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	
	createdpt, err := models.Createpr(prData)
	fmt.Println(err)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create product ratings successfully!",
		Results: createdpt,
	})
}


// UPDATE PRODUCT RATINGS
func UpdatePR(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	prData := services.Pro_Rate{}
	err := c.ShouldBind(&prData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	prData.Id = id

	updatedPT, err := models.Updatepr(prData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update product ratings successfully!",
		Results: updatedPT,
	})
}

// DELETE PRODUCT RATINGS
func DeletePR(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptData, err := models.Deletepr(id)
	if err != nil {
		msg := "Product ratings not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete product ratings successfully!",
		Results: ptData,
	})
}