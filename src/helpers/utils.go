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
	} else if strings.HasPrefix(err.Error(), "Key: 'RLUsers.Email'") { 
		c.JSON(http.StatusBadRequest, &services.ResponseBack{
			Success: false,
			Message: ms,
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
	} else {
		c.JSON(http.StatusInternalServerError, &services.ResponseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
}