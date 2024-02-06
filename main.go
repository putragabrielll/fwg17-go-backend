package main

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/routers"
)

func main(){
	r := gin.Default() // router => r "inisial aja"
	routers.Combine(r)
	r.Run("127.0.0.1:8080") // jika ingin ganti PORT bisa menggunakan ":5050" .
}