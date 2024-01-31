package adminController

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Person struct{
	Id int 				`json:"id"`
	FullName string 	`json:"fullname"`
	Email string 		`json:"email"`
	PhoneNumber string 	`json:"phoneNumber"`
	Address string 		`json:"address"`
	Picture string 		`json:"picture"`
	Password string 	`json:"password"`
	CreatedAt string 	`json:"createdAt"`
}

type responseList struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
	Results interface{}	`json:"results"`

}

func ListAllUsers(c *gin.Context){ // contex => c "inisial aja"
	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List all users!",
		Results: Person{},
	})
}