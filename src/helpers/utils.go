package helpers

import (
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)


type responseBack struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`

}

var c *gin.Context // untuk menggunakan gin context

func Utils(err error){
	// log.Fatalln(err)
	if err != nil {
		if strings.HasPrefix(err.Error(), "sql: no rows in result set") { // masi bermasalah
			log.Fatalln(err)
			c.JSON(http.StatusNotFound, &responseBack{
				Success: false,
				Message: "Users not found",
			})
			return
		}
		
		c.JSON(http.StatusInternalServerError, &responseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
}