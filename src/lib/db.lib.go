package lib

import (
	"os"
	"github.com/jmoiron/sqlx"
)


var DBClient *sqlx.DB

func DbConnection(){
	// dbname=aplikasiCoffeeShop sslmode=disable
	// postgresql://localhost:5432/aplikasiCoffeeShop

	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}