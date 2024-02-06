package helpers

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/services"
)


func Utils(err error, ms string, c *gin.Context){
	
	if strings.HasPrefix(err.Error(), "sql: no rows") { 
		c.JSON(http.StatusNotFound, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Person.Email'") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Person.Password'") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Password not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'RLUsers.Email'") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'RLUsers.Password'") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Password not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'FormReset.Email'") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if strings.HasPrefix(err.Error(), "invalid hash format") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if strings.HasSuffix(err.Error(), `unique constraint "users_email_key"`) { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if ms == "confirmPassword" {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Confirm password does not match!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Products.Name'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Name products not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Products.Price'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Price not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Products.Image'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Image not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Products.Description'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Description not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Products.Qty'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Qty not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Promo.Name'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Promo Name not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Promo.Code'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Promo Code not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Promo.Percentage'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Percentage not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Promo.MaximumPromo'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Maximum Promo not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Promo.MinimumAmount'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Minimum Amount not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pv.Name'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Variant name not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pv.AdditionalPrice'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Additional Price not be null!",
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, &services.ResponseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
}