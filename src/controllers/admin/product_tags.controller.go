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


func ListAllPT(c *gin.Context){
	filterby := c.DefaultQuery("filterby", "product")
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "product")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountCategories(filterby, filter)
	page_total := math.Ceil(float64(countData)/float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1


	products, err := models.ListAllCtgr(filterby, filter, sortby, order, limit, offset)
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