package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/services"
	"net/http"
	"strings"
)

func Utils(err error, ms string, c *gin.Context) {

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
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Tags.ProductId'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Product Id not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Tags.TagsId'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Tags Id not be null!",
		})
		return
	} else if strings.HasSuffix(err.Error(), `violates foreign key constraint "productTags_productId_fkey"`) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Product Id is not regis!",
		})
		return
	} else if strings.HasSuffix(err.Error(), `violates foreign key constraint "productTags_tagsId_fkey"`) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Tags Id is not regis!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Cate.ProductId'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Product Id not be null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Cate.CategoriesId'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Categories Id not be null!",
		})
		return
	} else if strings.HasSuffix(err.Error(), `violates foreign key constraint "productCategories_productId_fkey"`) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Product Id is not regis!",
		})
		return
	} else if strings.HasSuffix(err.Error(), `violates foreign key constraint "productCategories_categoriesId_fkey"`) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Categories Id is not regis!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Rate.ProductId'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Product Id is not null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Rate.Rate'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Rate is not null!",
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Pro_Rate.UsersId'") {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Users Id is not null!",
		})
		return
	} else if strings.HasSuffix(err.Error(), `violates foreign key constraint "productRatings_productId_fkey"`) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Product Id is not regis!",
		})
		return
	} else if strings.HasSuffix(err.Error(), `violates foreign key constraint "productRatings_usersId_fkey"`) {
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: "Users Id is not regis!",
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
