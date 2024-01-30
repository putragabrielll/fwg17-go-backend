package main

import (
	"github.com/gin-gonic/gin"
	"github.com/putragabrielll/go-backend/src/routers"
)

func main(){
	r := gin.Default() // router => r "inisial aja"
	routers.Combine(r)
	r.Run() // jika ingin ganti PORT bisa menggunakan ":5050" .
}