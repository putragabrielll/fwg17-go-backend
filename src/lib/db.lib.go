package lib

import (
	"log"
	"github.com/jmoiron/sqlx"
)


func LibDb(){
	// dbname=aplikasiCoffeeShop sslmode=disable
	// postgresql://localhost:5432/aplikasiCoffeeShop
	
	db, err := sqlx.Connect("postgres", "dbname=aplikasiCoffeeShop sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}