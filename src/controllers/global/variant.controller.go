package globalController

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL PRODUCT VARIANT
func ListAllVariant(c *gin.Context){
	ListPS, err := models.ListAllpv()
	if err != nil {
		msg := "Product variant not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all product variant!",
		Results: ListPS,
	})
}