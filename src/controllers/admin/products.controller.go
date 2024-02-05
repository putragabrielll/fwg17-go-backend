package adminController

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL PRODUCTS
func ListAllProducts(c *gin.Context){
	products, err := models.ListAllProducts()
	if err != nil {
		msg := "Products not found"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all products!",
		Results: products,
	})
}