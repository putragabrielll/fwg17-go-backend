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

// SELECT ALL ORDERS
func ListAllOrders(c *gin.Context){
	table := c.DefaultQuery("table", "user")
	column := c.DefaultQuery("column", "fullName")
	filter := c.DefaultQuery("filter", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	order := c.DefaultQuery("order", "ASC")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	offset := (page - 1) * limit

	countData, _ := models.CountOrders(table, column, filter)
	page_total := math.Ceil(float64(countData) / float64(limit))
	page_next := page + 1
	if !(page_next <= int(page_total)) {
		page_next = int(0)
	}
	page_prev := page - 1

	ordersData, err := models.ListOrders(table, column, filter, order, limit, offset)
	if err != nil {
		msg := "Orders not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseAll{
		Success: true,
		Message: "List all orders!",
		PageInfo: services.PageInfo{
			CurrentPage: page,
			TotalPage:   int(page_total),
			NextPage:    page_next,
			PrevPage:    page_prev,
			TotalData:   countData,
		},
		Results: ordersData,
	})
}

// GET ORDERS BY id
func IdOrders(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	ptData, err := models.FindOrders(id)
	if err != nil {
		msg := "Orders not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Detail orders!",
		Results: ptData,
	})
}

// CREATE ORDERS
func CreateOrders(c *gin.Context){
	ordersData := services.Orders{}
	err := c.ShouldBind(&ordersData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	// ------------
	findUser, err := models.FindUsersId(ordersData.UsersId)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c)
		return
	}
	ordersData.TaxAmount = int64(float64(ordersData.Total) * 0.1)
	ordersData.Status = "on-progress"
	ordersData.DeliveryAddress = findUser.Address.Val
	ordersData.FullName = findUser.FullName.Val
	ordersData.Email = findUser.Email
	// ------------
	createorder, err := models.CreateOrders(ordersData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Create orders successfully!",
		Results: createorder,
	})
}

// UPDATE ORDERS
func UpdateOrders(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	orderData := services.Orders{}
	err := c.ShouldBind(&orderData)
	if err != nil {
		msg := "Data not be null!"
		helpers.Utils(err, msg, c)
		return
	}
	// ------------
	findUser, err := models.FindUsersId(orderData.UsersId)
	if err != nil {
		msg := "Users not found"
		helpers.Utils(err, msg, c)
		return
	}
	orderData.TaxAmount = int64(float64(orderData.Total) * 0.1)
	orderData.Status = "on-progress"
	orderData.DeliveryAddress = findUser.Address.Val
	orderData.FullName = findUser.FullName.Val
	orderData.Email = findUser.Email
	orderData.Id = id
	// ------------

	updatedPT, _ := models.UpdateOrders(orderData)
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Update orders successfully!",
		Results: updatedPT,
	})
}

// DELETE ORDERS
func DeleteOrders(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	orders, err := models.DeleteOrders(id)
	if err != nil {
		msg := "Orders not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "Delete Orders successfully!",
		Results: orders,
	})
}
