package adminUsersController

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/putragabrielll/go-backend/src/helpers"
	modelsUsers "github.com/putragabrielll/go-backend/src/models"
)




type Person struct{
	Id int 				`json:"id"`
	FullName string 	`json:"fullname"`
	Email string 		`json:"email"`
	PhoneNumber string 	`json:"phoneNumber"`
	Address string 		`json:"address"`
	Picture string 		`json:"picture"`
	Role string 		`json:"role"`
	Password string 	`json:"password"`
	CreatedAt string 	`json:"createdAt"`
	UpdatedAt string 	`json:"updatedAt"`
}

type responseList struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
	Results interface{}	`json:"results"`

}
type responseBack struct{
	Success bool		`json:"success"`
	Message string		`json:"message"`
}







// SELECT ALL USERS
func ListAllUsers(c *gin.Context){ // contex => c "inisial aja"
	users, err := modelsUsers.ListAllUsers()
	// helpers.Utils(err []helpers.responseBack{}) // Error Handler

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &responseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
	

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List all users!",
		Results: users,
	})
}


// GET USERS BY id
func IdUsers(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := modelsUsers.FindUsersId(id)

	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, &responseBack{
			Success: false,
			Message: "Internal Server Error",
		})
		return 
	}
	

	c.JSON(http.StatusOK, &responseList{
		Success: true,
		Message: "List all users!",
		Results: users,
	})
}


//