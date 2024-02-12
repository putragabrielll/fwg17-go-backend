package customerController

import (
	"net/http"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)


// GET ORDERS BY id
func IdOrders(c *gin.Context){
	claims := jwt.ExtractClaims(c)
	id := int(claims["id"].(float64))

	ordersData, err := models.CustomerFindOrders(id)
	if err != nil {
		msg := "Orders not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}

	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all orders!",
		Results: ordersData,
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