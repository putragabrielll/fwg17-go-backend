package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/routers"
	"github.com/putragabrielll/go-backend/src/services"
	"net/http"
)

func noLink(c *gin.Context) {
	c.JSON(http.StatusNotFound, &services.ResponseBack{
		Success: false,
		Message: "Resource not found!",
	})
}

func main() {
	r := gin.Default() // router => r "inisial aja"
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	}))
	routers.Combine(r)
	r.NoRoute(noLink)       // jika link tidak ada
	r.Run(":8080") // jika ingin ganti PORT bisa menggunakan ":5050" .
}
