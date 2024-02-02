package helpers

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)


type responseBack struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
}


func Utils(err error, ms string, c *gin.Context){
	
	if strings.HasPrefix(err.Error(), "sql: no rows") { 
		c.JSON(http.StatusNotFound, &responseBack{
			Success: false,
			Message: ms,
		})
		return
	} else if strings.HasPrefix(err.Error(), "Key: 'Person.Email'") { 
		c.JSON(http.StatusBadRequest, &responseBack{
			Success: false,
			Message: ms,
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, &responseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
}