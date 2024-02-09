package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/routers"
	"github.com/putragabrielll/go-backend/src/services"
)


func noLink(c *gin.Context){
	c.JSON(http.StatusNotFound, &services.ResponseBack{
		Success: false,
		Message: "Resource not found!",
	})
}

func main(){
	r := gin.Default() // router => r "inisial aja"
	routers.Combine(r)
	r.NoRoute(noLink) // jika link tidak ada
	r.Run("127.0.0.1:8080") // jika ingin ganti PORT bisa menggunakan ":5050" .
}