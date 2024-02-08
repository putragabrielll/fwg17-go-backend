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

// SELECT ALL ORDERS DETAILS
func ListAllOD(c *gin.Context){
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	sortby := c.DefaultQuery("sortby", "createdAt")
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.Countod(filter)
	page_total := math.Ceil(float64(countData) / float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	detailOrder, err := models.ListAllod(filter, sortby, order, limit, offset)
	if err != nil {
		msg := "Products tags not found"
		helpers.Utils(err, msg, c)
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
		Results: detailOrder,
	})
}

// GET ORDERS DETAILS BY id
func IdOD(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	doData, err := models.Findod(id)
	if err != nil {
		msg := "Details order not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail order!",
		Results: doData,
	})
}

// CREATE ORDERS DETAILS
func CreateOD(c *gin.Context){
	odData := services.OrdersDetails{}
	err := c.ShouldBind(&odData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	// --------------
	product, err := models.FindProductsId(odData.ProductId)
	size, err := models.FindPSId(odData.ProductSizeId)
	variant, err := models.FindPVId(odData.ProductVariantId)
	odData.SubTotal = (int(product.Price) + int(size.AdditionalPrice) + int(variant.AdditionalPrice)) * odData.Qty
	// --------------

	createOD, err := models.Createod(odData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create order details successfully!",
		Results: createOD,
	})
}

// UPDATE ORDERS DETAILS
func UpdateOD(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	odData := services.OrdersDetails{}
	err := c.ShouldBind(&odData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	// --------------
	product, err := models.FindProductsId(odData.ProductId)
	size, err := models.FindPSId(odData.ProductSizeId)
	variant, err := models.FindPVId(odData.ProductVariantId)
	odData.SubTotal = (int(product.Price) + int(size.AdditionalPrice) + int(variant.AdditionalPrice)) * odData.Qty
	odData.Id = id
	// --------------

	updatedPC, err := models.Updateod(odData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c) // Error Handle
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update order details successfully!",
		Results: updatedPC,
	})
}

// DELETE ORDERS DETAILS
func DeleteOD(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	pcData, err := models.Deleteod(id)
	if err != nil {
		msg := "Order details not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete order details successfully!",
		Results: pcData,
	})
}