package adminController

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type responseList struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
	// Results interface{}	`json:"results"`

}

func DataUsers(c *gin.Context){ // contex => c "inisial aja"
	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List all users!",
		// Results: ,
	})
}