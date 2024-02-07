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

// SELECT ALL CATEGORIES
func ListAllCategories(c *gin.Context){
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "id")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountCategories(filter)
	page_total := math.Ceil(float64(countData)/float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1


	products, err := models.ListAllCtgr(filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Categories not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all categories!",
		PageInfo: services.PageInfo{
			CurrentPage: page,
			TotalPage: int(page_total),
			NextPage: page_next,
			PrevPage: page_prev,
			TotalData: countData,
		},
		Results: products,
	})
}


// GET CATEGORIES BY id
func IdCategories(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	categories, err := models.FindCategories(id)
	if err != nil {
		msg := "Categories not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail categories!",
		Results: categories,
	})
}


// CREATE CATEGORIES
func CreateCTGR(c *gin.Context){
	categoriesData := services.Categories{}
	err := c.ShouldBind(&categoriesData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	
	createdCategories, _ := models.CreateCat(categoriesData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create products successfully!",
		Results: createdCategories,
	})
}


// UPDATE CATEGORIES
func UpdateCTGR(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	categoriesData := services.Categories{}
	err := c.ShouldBind(&categoriesData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	categoriesData.Id = id

	updatedUsers, err := models.UpdateCat(categoriesData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update categories successfully!",
		Results: updatedUsers,
	})
}


// DELETE CATEGORIES
func DeleteCTGR(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	cateData, err := models.DeleteCat(id)
	if err != nil {
		msg := "Categories not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete categories successfully!",
		Results: cateData,
	})
}