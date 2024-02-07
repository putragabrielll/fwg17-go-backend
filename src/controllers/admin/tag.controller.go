package adminController

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
	"math"
	"net/http"
	"strconv"
)

// SELECT ALL TAGS
func ListAllTags(c *gin.Context) {
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "id")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountTags(filter)
	page_total := math.Ceil(float64(countData) / float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	products, err := models.ListAllTags(filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Tags not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all tags!",
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


// GET TAGS BY id
func IdTags(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	tags, err := models.FindTags(id)
	if err != nil {
		msg := "Tags not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail tags!",
		Results: tags,
	})
}


// CREATE TAGS
func CreateTags(c *gin.Context){
	tagsData := services.Tags{}
	err := c.ShouldBind(&tagsData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	
	createTags, _ := models.CreateTags(tagsData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create products successfully!",
		Results: createTags,
	})
}


// UPDATE TAGS
func UpdateTags(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	tagsData := services.Tags{}
	err := c.ShouldBind(&tagsData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	tagsData.Id = id

	updateTags, err := models.UpdateTags(tagsData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update categories successfully!",
		Results: updateTags,
	})
}


// DELETE TAGS
func DeleteTags(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	tags, err := models.DeleteTags(id)
	if err != nil {
		msg := "Tags not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete tags successfully!",
		Results: tags,
	})
}