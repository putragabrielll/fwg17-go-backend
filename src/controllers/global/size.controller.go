package globalController

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/helpers"
	"github.com/putragabrielll/go-backend/src/models"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT ALL product size
func ListAllPS(c *gin.Context){
	ListPS, err := models.ListAllps()
	if err != nil {
		msg := "Product size not found!"
		helpers.Utils(err, msg, c) // Error Handler
		return
	}
	
	c.JSON(http.StatusOK, &services.ResponseList{
		Success: true,
		Message: "List all Size!",
		Results: ListPS,
	})
}