package helpers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type responseBack struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`

}

func Utils(err []responseBack, c *gin.Context){

	if err != nil {
		c.JSON(http.StatusInternalServerError, &responseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
}