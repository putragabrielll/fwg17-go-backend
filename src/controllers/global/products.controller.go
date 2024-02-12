package globalController

import (
	"math"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL PRODUCTS
func ListAllProducts(c *gin.Context){
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "id")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountAllProducts(filter)
	page_total := math.Ceil(float64(countData)/float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1


	products, err := models.ListAllProducts(filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Products not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all products!",
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


// GET PRODUCTS BY id
func IdProducts(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := models.FindProductsId(id)
	if err != nil {
		msg := "Product not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail products!",
		Results: users,
	})
}